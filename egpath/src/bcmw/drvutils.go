package bcmw

import (
	"delay"
	"encoding/binary/le"
	"fmt"
	"rtos"

	"sdcard"
	"sdcard/sdio"
)

func (d *Driver) debug(f string, args ...interface{}) {
	if d.error() {
		return
	}
	fmt.Printf(f, args...)
}

func (d *Driver) error() bool {
	return d.sd.Err(false) != nil ||
		d.ioStatus&^sdcard.IO_CURRENT_STATE != 0 ||
		d.err != 0
}

func (d *Driver) firstErr() error {
	if err := d.sd.Err(false); err != nil {
		return err
	}
	if d.ioStatus&^sdcard.IO_CURRENT_STATE != 0 {
		return ErrIOStatus
	}
	if d.err != 0 {
		return d.err
	}
	return nil
}

func (d *Driver) cmd52(f, addr int, flags sdcard.IORWFlags, val byte) byte {
	if d.error() {
		return 0
	}
	val, d.ioStatus = d.sd.SendCmd(sdcard.CMD52(f, addr, flags, val)).R5()
	return val
}

func (d *Driver) sdioSetBlockSize(f, blksiz int) {
	if d.error() {
		return
	}
	d.cmd52(cia, f<<8+sdio.FBR_BLKSIZE0, sdcard.Write, byte(blksiz))
	d.cmd52(cia, f<<8+sdio.FBR_BLKSIZE1, sdcard.Write, byte(blksiz>>8))
}

func (d *Driver) sdioEnableFunc(f, timeoutms int) {
	if d.error() {
		return
	}
	r := d.cmd52(cia, sdio.CCCR_IOEN, sdcard.Read, 0)
	m := byte(1 << uint(f))
	d.cmd52(cia, sdio.CCCR_IOEN, sdcard.Write, r|m)
	for retry := timeoutms >> 1; retry > 0; retry-- {
		r := d.cmd52(cia, sdio.CCCR_IORDY, sdcard.Read, 0)
		if d.error() || r&m != 0 {
			return
		}
		delay.Millisec(2)
	}
	d.err = ErrTimeout
}

func (d *Driver) sdioDisableFunc(f int) {
	if d.error() {
		return
	}
	r := d.cmd52(cia, sdio.CCCR_IOEN, sdcard.Read, 0)
	r &^= 1 << uint(f)
	d.cmd52(cia, sdio.CCCR_IOEN, sdcard.Write, r)
}

func (d *Driver) sdiodRead8(addr int) byte {
	if d.error() {
		return 0
	}
	return d.cmd52(backplane, addr, sdcard.Read, 0)
}

func (d *Driver) sdiodWrite8(addr int, v byte) {
	if d.error() {
		return
	}
	d.cmd52(backplane, addr, sdcard.Write, v)
}

// The backplaneSetWindow, backplaneRead32, backplaneWrite32 are methods
// that allow to access core registers in the way specific to Sonics Silicon
// Backplane. More info: http://www.gc-linux.org/wiki/Wii:WLAN

func (d *Driver) backplaneSetWindow(addr uint32) {
	if d.error() {
		return
	}
	addr &^= 0x7FFF
	win := d.backplaneWindow
	if win == addr {
		return
	}
	d.backplaneWindow = addr
	for n := 0; n < 3; n++ {
		addr >>= 8
		win >>= 8
		if a := byte(addr); a != byte(win) {
			d.sdiodWrite8(sbsdioFunc1SBAddrLow+n, a)
		}
	}
	if d.error() {
		d.backplaneWindow = 0
	}
}

func (d *Driver) backplaneRead8(addr uint32) byte {
	d.backplaneSetWindow(addr)
	return d.sdiodRead8(int(addr & 0x7FFF))
}

func (d *Driver) backplaneWrite8(addr uint32, b byte) {
	d.backplaneSetWindow(addr)
	d.sdiodWrite8(int(addr&0x7FFF), b)
}

func waitDataReady(sd sdcard.Host) bool {
	return sd.Wait(rtos.Nanosec() + 1e9)
}

func (d *Driver) backplaneRead32(addr uint32) uint32 {
	d.backplaneSetWindow(addr)
	if d.error() {
		return 0
	}
	sd := d.sd
	if !waitDataReady(sd) {
		d.err = ErrTimeout
		return 0
	}
	var buf [1]uint64
	sd.SetupData(sdcard.Recv|sdcard.IO|sdcard.Block4, buf[:], 4)
	_, d.ioStatus = sd.SendCmd(sdcard.CMD53(
		backplane, int(addr&0x7FFF|sbsdioAccess32bit), sdcard.Read, 4,
	)).R5()
	return le.Decode32(sdcard.AsData(buf[:]).Bytes())
}

func (d *Driver) backplaneWrite32(addr, v uint32) {
	d.backplaneSetWindow(addr)
	if d.error() {
		return
	}
	sd := d.sd
	if !waitDataReady(sd) {
		d.err = ErrTimeout
		return
	}
	var buf [1]uint64
	le.Encode32(sdcard.AsData(buf[:]).Bytes(), v)
	sd.SetupData(sdcard.Send|sdcard.IO|sdcard.Block4, buf[:], 4)
	_, d.ioStatus = sd.SendCmd(sdcard.CMD53(
		backplane, int(addr&0x7FFF|sbsdioAccess32bit), sdcard.Write, 4,
	)).R5()
}

func (d *Driver) backplaneWrite(addr uint32, data []byte) {
	if d.error() {
		return
	}
	head, aligned, tail := uint64slice(data)
	d.debug(
		"data: %d (%x), head: %d, aligned: %d, tail: %d\n",
		len(data), &data[0], len(head), len(aligned), len(tail),
	)
	for _, b := range head {
		d.backplaneWrite8(addr, b)
		addr++
	}
	if d.error() {
		return
	}
	sd := d.sd
	/*
		for {
			nbl := len(aligned) >> 3
			if nbl >= 0x1FF {
				nbl = 0x1FF
			}
			d.backplaneSetWindow(addr)
			if !waitDataReady(sd) {
				d.err = ErrTimeout
				return
			}
			sd.SetupData(sdcard.Send|sdcard.IO|sdcard.Block64, aligned, nbl*64)
			_, d.ioStatus = sd.SendCmd(sdcard.CMD53(
				backplane, int(addr&0x7FFF), sdcard.BlockWrite|sdcard.IncAddr, nbl,
			)).R5()
			if d.error() {
				return
			}
			aligned = aligned[nbl*8:]
			addr += uint32(nbl) * 64
		}
	*/
	for i := 0; i < len(aligned); i++ {
		d.backplaneSetWindow(addr)
		if !waitDataReady(sd) {
			d.err = ErrTimeout
			return
		}
		sd.SetupData(sdcard.Send|sdcard.IO|sdcard.Block8, aligned[i:], 8)
		_, d.ioStatus = sd.SendCmd(sdcard.CMD53(
			backplane, int(addr&0x7FFF), sdcard.Write|sdcard.IncAddr, 8,
		)).R5()
		if d.error() {
			return
		}
		addr += 8
	}
	for _, b := range tail {
		d.backplaneWrite8(addr, b)
		addr++
	}
}

/*

func (d *Driver) wbb(addr uint32, buf []uint64) {
	sd := d.sd
	for len(buf) >= 8 {
		d.sdiodSetBackplaneWindow(addr)
		if d.timeout = !waitDataReady(sd); d.timeout {
			return
		}
		nbl := len(buf) >> 3
		if nbl >= 0x1FF {
			nbl = 0x1FF
		}
		sd.SetupData(sdcard.Send|sdcard.IO|sdcard.Block64, buf, nbl*64)
		_, d.ioStatus = sd.SendCmd(sdcard.CMD53(
			backplane, int(addr&0x7FFF), sdcard.BlockWrite|sdcard.IncAddr, nbl,
		)).R5()
		if d.error() {
			return
		}
		buf = buf[nbl*8:]
		addr += uint32(nbl) * 64
	}
	if len(buf) == 0 {
		return
	}
	d.sdiodSetBackplaneWindow(addr)
	// STM32x MCUs don't handle well multibyte transfers so configure the SDMMC
	// driver to use 8-byte block mode len(buf) times.
	for i := 0; i < len(buf); i++ {
		if d.timeout = !waitDataReady(sd); d.timeout {
			return
		}
		sd.SetupData(sdcard.Send|sdcard.IO|sdcard.Block8, buf[i:], 8)
		_, d.ioStatus = sd.SendCmd(sdcard.CMD53(
			backplane, int(addr&0x7FFF|sbsdioAccess32bit),
			sdcard.Write|sdcard.IncAddr, 8,
		)).R5()
		if d.error() {
			return
		}
		addr += 8
	}
}

*/

func (d *Driver) chipIsCoreUp(core int) bool {
	if d.error() {
		return false
	}
	base := wrapBase[core]
	r := d.backplaneRead32(base + agentIOCtl)
	if r&(ioCtlFGC|ioCtlClk) != ioCtlClk {
		return false
	}
	return d.backplaneRead32(base+agentResetCtl)&1 == 0
}

func (d *Driver) chipCoreDisable(core int, prereset, reset uint32) {
	if d.error() {
		return
	}
	base := wrapBase[core]
	if d.backplaneRead32(base+agentResetCtl)&1 == 0 {
		goto configure // Already in reset state.
	}
	d.backplaneWrite32(base+agentIOCtl, ioCtlFGC|ioCtlClk|prereset)
	d.backplaneRead32(base + agentIOCtl)
	d.backplaneWrite32(base+agentResetCtl, 1)
	delay.Millisec(1)
	if d.backplaneRead32(base+agentResetCtl)&1 == 0 {
		if d.err == 0 {
			d.err = ErrTimeout
		}
		return
	}
configure:
	d.backplaneWrite32(base+agentIOCtl, ioCtlFGC|ioCtlClk|reset)
	d.backplaneRead32(base + agentIOCtl)
}

func (d *Driver) chipCoreReset(core int, prereset, reset, postreset uint32) {
	if d.error() {
		return
	}
	d.chipCoreDisable(core, prereset, reset)
	base := wrapBase[core]
	for retry := 3; ; retry-- {
		d.backplaneWrite32(base+agentResetCtl, 0)
		delay.Millisec(1)
		r := d.backplaneRead32(base + agentResetCtl)
		if d.error() {
			return
		}
		if r&1 == 0 {
			break
		}
		if retry == 1 {
			d.err = ErrTimeout
			return
		}
	}
	d.backplaneWrite32(base+agentIOCtl, ioCtlClk|postreset)
	d.backplaneRead32(base + agentIOCtl)
}
