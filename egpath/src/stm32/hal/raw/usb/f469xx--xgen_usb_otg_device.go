// +build f469xx
package usb

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f469xx/mmap"
)

type USB_OTG_Device_Periph struct {
	DCFG       RDCFG
	DCTL       RDCTL
	DSTS       RDSTS
	_          uint32
	DIEPMSK    RDIEPMSK
	DOEPMSK    RDOEPMSK
	DAINT      RDAINT
	DAINTMSK   RDAINTMSK
	_          [2]uint32
	DVBUSDIS   RDVBUSDIS
	DVBUSPULSE RDVBUSPULSE
	DTHRCTL    RDTHRCTL
	DIEPEMPMSK RDIEPEMPMSK
	DEACHINT   RDEACHINT
	DEACHMSK   RDEACHMSK
	_          uint32
	DINEP1MSK  RDINEP1MSK
	_          [15]uint32
	DOUTEP1MSK RDOUTEP1MSK
}

func (p *USB_OTG_Device_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type DCFG uint32

func (b DCFG) Field(mask DCFG) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DCFG) J(v int) DCFG {
	return DCFG(bits.MakeField32(v, uint32(mask)))
}

type RDCFG struct{ mmio.U32 }

func (r *RDCFG) Bits(mask DCFG) DCFG    { return DCFG(r.U32.Bits(uint32(mask))) }
func (r *RDCFG) StoreBits(mask, b DCFG) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDCFG) SetBits(mask DCFG)      { r.U32.SetBits(uint32(mask)) }
func (r *RDCFG) ClearBits(mask DCFG)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDCFG) Load() DCFG             { return DCFG(r.U32.Load()) }
func (r *RDCFG) Store(b DCFG)           { r.U32.Store(uint32(b)) }

func (r *RDCFG) AtomicStoreBits(mask, b DCFG) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDCFG) AtomicSetBits(mask DCFG)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDCFG) AtomicClearBits(mask DCFG)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDCFG struct{ mmio.UM32 }

func (rm RMDCFG) Load() DCFG   { return DCFG(rm.UM32.Load()) }
func (rm RMDCFG) Store(b DCFG) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) DSPD() RMDCFG {
	return RMDCFG{mmio.UM32{&p.DCFG.U32, uint32(DSPD)}}
}

func (p *USB_OTG_Device_Periph) NZLSOHSK() RMDCFG {
	return RMDCFG{mmio.UM32{&p.DCFG.U32, uint32(NZLSOHSK)}}
}

func (p *USB_OTG_Device_Periph) DAD() RMDCFG {
	return RMDCFG{mmio.UM32{&p.DCFG.U32, uint32(DAD)}}
}

func (p *USB_OTG_Device_Periph) PFIVL() RMDCFG {
	return RMDCFG{mmio.UM32{&p.DCFG.U32, uint32(PFIVL)}}
}

func (p *USB_OTG_Device_Periph) PERSCHIVL() RMDCFG {
	return RMDCFG{mmio.UM32{&p.DCFG.U32, uint32(PERSCHIVL)}}
}

type DCTL uint32

func (b DCTL) Field(mask DCTL) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DCTL) J(v int) DCTL {
	return DCTL(bits.MakeField32(v, uint32(mask)))
}

type RDCTL struct{ mmio.U32 }

func (r *RDCTL) Bits(mask DCTL) DCTL    { return DCTL(r.U32.Bits(uint32(mask))) }
func (r *RDCTL) StoreBits(mask, b DCTL) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDCTL) SetBits(mask DCTL)      { r.U32.SetBits(uint32(mask)) }
func (r *RDCTL) ClearBits(mask DCTL)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDCTL) Load() DCTL             { return DCTL(r.U32.Load()) }
func (r *RDCTL) Store(b DCTL)           { r.U32.Store(uint32(b)) }

func (r *RDCTL) AtomicStoreBits(mask, b DCTL) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDCTL) AtomicSetBits(mask DCTL)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDCTL) AtomicClearBits(mask DCTL)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDCTL struct{ mmio.UM32 }

func (rm RMDCTL) Load() DCTL   { return DCTL(rm.UM32.Load()) }
func (rm RMDCTL) Store(b DCTL) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) RWUSIG() RMDCTL {
	return RMDCTL{mmio.UM32{&p.DCTL.U32, uint32(RWUSIG)}}
}

func (p *USB_OTG_Device_Periph) SDIS() RMDCTL {
	return RMDCTL{mmio.UM32{&p.DCTL.U32, uint32(SDIS)}}
}

func (p *USB_OTG_Device_Periph) GINSTS() RMDCTL {
	return RMDCTL{mmio.UM32{&p.DCTL.U32, uint32(GINSTS)}}
}

func (p *USB_OTG_Device_Periph) GONSTS() RMDCTL {
	return RMDCTL{mmio.UM32{&p.DCTL.U32, uint32(GONSTS)}}
}

func (p *USB_OTG_Device_Periph) TCTL() RMDCTL {
	return RMDCTL{mmio.UM32{&p.DCTL.U32, uint32(TCTL)}}
}

func (p *USB_OTG_Device_Periph) SGINAK() RMDCTL {
	return RMDCTL{mmio.UM32{&p.DCTL.U32, uint32(SGINAK)}}
}

func (p *USB_OTG_Device_Periph) CGINAK() RMDCTL {
	return RMDCTL{mmio.UM32{&p.DCTL.U32, uint32(CGINAK)}}
}

func (p *USB_OTG_Device_Periph) SGONAK() RMDCTL {
	return RMDCTL{mmio.UM32{&p.DCTL.U32, uint32(SGONAK)}}
}

func (p *USB_OTG_Device_Periph) CGONAK() RMDCTL {
	return RMDCTL{mmio.UM32{&p.DCTL.U32, uint32(CGONAK)}}
}

func (p *USB_OTG_Device_Periph) POPRGDNE() RMDCTL {
	return RMDCTL{mmio.UM32{&p.DCTL.U32, uint32(POPRGDNE)}}
}

type DSTS uint32

func (b DSTS) Field(mask DSTS) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DSTS) J(v int) DSTS {
	return DSTS(bits.MakeField32(v, uint32(mask)))
}

type RDSTS struct{ mmio.U32 }

func (r *RDSTS) Bits(mask DSTS) DSTS    { return DSTS(r.U32.Bits(uint32(mask))) }
func (r *RDSTS) StoreBits(mask, b DSTS) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDSTS) SetBits(mask DSTS)      { r.U32.SetBits(uint32(mask)) }
func (r *RDSTS) ClearBits(mask DSTS)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDSTS) Load() DSTS             { return DSTS(r.U32.Load()) }
func (r *RDSTS) Store(b DSTS)           { r.U32.Store(uint32(b)) }

func (r *RDSTS) AtomicStoreBits(mask, b DSTS) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDSTS) AtomicSetBits(mask DSTS)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDSTS) AtomicClearBits(mask DSTS)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDSTS struct{ mmio.UM32 }

func (rm RMDSTS) Load() DSTS   { return DSTS(rm.UM32.Load()) }
func (rm RMDSTS) Store(b DSTS) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) SUSPSTS() RMDSTS {
	return RMDSTS{mmio.UM32{&p.DSTS.U32, uint32(SUSPSTS)}}
}

func (p *USB_OTG_Device_Periph) ENUMSPD() RMDSTS {
	return RMDSTS{mmio.UM32{&p.DSTS.U32, uint32(ENUMSPD)}}
}

func (p *USB_OTG_Device_Periph) EERR() RMDSTS {
	return RMDSTS{mmio.UM32{&p.DSTS.U32, uint32(EERR)}}
}

func (p *USB_OTG_Device_Periph) FNSOF() RMDSTS {
	return RMDSTS{mmio.UM32{&p.DSTS.U32, uint32(FNSOF)}}
}

type DIEPMSK uint32

func (b DIEPMSK) Field(mask DIEPMSK) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DIEPMSK) J(v int) DIEPMSK {
	return DIEPMSK(bits.MakeField32(v, uint32(mask)))
}

type RDIEPMSK struct{ mmio.U32 }

func (r *RDIEPMSK) Bits(mask DIEPMSK) DIEPMSK { return DIEPMSK(r.U32.Bits(uint32(mask))) }
func (r *RDIEPMSK) StoreBits(mask, b DIEPMSK) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDIEPMSK) SetBits(mask DIEPMSK)      { r.U32.SetBits(uint32(mask)) }
func (r *RDIEPMSK) ClearBits(mask DIEPMSK)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDIEPMSK) Load() DIEPMSK             { return DIEPMSK(r.U32.Load()) }
func (r *RDIEPMSK) Store(b DIEPMSK)           { r.U32.Store(uint32(b)) }

func (r *RDIEPMSK) AtomicStoreBits(mask, b DIEPMSK) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDIEPMSK) AtomicSetBits(mask DIEPMSK)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDIEPMSK) AtomicClearBits(mask DIEPMSK)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDIEPMSK struct{ mmio.UM32 }

func (rm RMDIEPMSK) Load() DIEPMSK   { return DIEPMSK(rm.UM32.Load()) }
func (rm RMDIEPMSK) Store(b DIEPMSK) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) XFRCM() RMDIEPMSK {
	return RMDIEPMSK{mmio.UM32{&p.DIEPMSK.U32, uint32(XFRCM)}}
}

func (p *USB_OTG_Device_Periph) EPDM() RMDIEPMSK {
	return RMDIEPMSK{mmio.UM32{&p.DIEPMSK.U32, uint32(EPDM)}}
}

func (p *USB_OTG_Device_Periph) TOM() RMDIEPMSK {
	return RMDIEPMSK{mmio.UM32{&p.DIEPMSK.U32, uint32(TOM)}}
}

func (p *USB_OTG_Device_Periph) ITTXFEMSK() RMDIEPMSK {
	return RMDIEPMSK{mmio.UM32{&p.DIEPMSK.U32, uint32(ITTXFEMSK)}}
}

func (p *USB_OTG_Device_Periph) INEPNMM() RMDIEPMSK {
	return RMDIEPMSK{mmio.UM32{&p.DIEPMSK.U32, uint32(INEPNMM)}}
}

func (p *USB_OTG_Device_Periph) INEPNEM() RMDIEPMSK {
	return RMDIEPMSK{mmio.UM32{&p.DIEPMSK.U32, uint32(INEPNEM)}}
}

func (p *USB_OTG_Device_Periph) TXFURM() RMDIEPMSK {
	return RMDIEPMSK{mmio.UM32{&p.DIEPMSK.U32, uint32(TXFURM)}}
}

func (p *USB_OTG_Device_Periph) BIM() RMDIEPMSK {
	return RMDIEPMSK{mmio.UM32{&p.DIEPMSK.U32, uint32(BIM)}}
}

type DOEPMSK uint32

func (b DOEPMSK) Field(mask DOEPMSK) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DOEPMSK) J(v int) DOEPMSK {
	return DOEPMSK(bits.MakeField32(v, uint32(mask)))
}

type RDOEPMSK struct{ mmio.U32 }

func (r *RDOEPMSK) Bits(mask DOEPMSK) DOEPMSK { return DOEPMSK(r.U32.Bits(uint32(mask))) }
func (r *RDOEPMSK) StoreBits(mask, b DOEPMSK) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDOEPMSK) SetBits(mask DOEPMSK)      { r.U32.SetBits(uint32(mask)) }
func (r *RDOEPMSK) ClearBits(mask DOEPMSK)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDOEPMSK) Load() DOEPMSK             { return DOEPMSK(r.U32.Load()) }
func (r *RDOEPMSK) Store(b DOEPMSK)           { r.U32.Store(uint32(b)) }

func (r *RDOEPMSK) AtomicStoreBits(mask, b DOEPMSK) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDOEPMSK) AtomicSetBits(mask DOEPMSK)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDOEPMSK) AtomicClearBits(mask DOEPMSK)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDOEPMSK struct{ mmio.UM32 }

func (rm RMDOEPMSK) Load() DOEPMSK   { return DOEPMSK(rm.UM32.Load()) }
func (rm RMDOEPMSK) Store(b DOEPMSK) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) XFRCM() RMDOEPMSK {
	return RMDOEPMSK{mmio.UM32{&p.DOEPMSK.U32, uint32(XFRCM)}}
}

func (p *USB_OTG_Device_Periph) EPDM() RMDOEPMSK {
	return RMDOEPMSK{mmio.UM32{&p.DOEPMSK.U32, uint32(EPDM)}}
}

func (p *USB_OTG_Device_Periph) STUPM() RMDOEPMSK {
	return RMDOEPMSK{mmio.UM32{&p.DOEPMSK.U32, uint32(STUPM)}}
}

func (p *USB_OTG_Device_Periph) OTEPDM() RMDOEPMSK {
	return RMDOEPMSK{mmio.UM32{&p.DOEPMSK.U32, uint32(OTEPDM)}}
}

func (p *USB_OTG_Device_Periph) OTEPSPRM() RMDOEPMSK {
	return RMDOEPMSK{mmio.UM32{&p.DOEPMSK.U32, uint32(OTEPSPRM)}}
}

func (p *USB_OTG_Device_Periph) B2BSTUP() RMDOEPMSK {
	return RMDOEPMSK{mmio.UM32{&p.DOEPMSK.U32, uint32(B2BSTUP)}}
}

func (p *USB_OTG_Device_Periph) OPEM() RMDOEPMSK {
	return RMDOEPMSK{mmio.UM32{&p.DOEPMSK.U32, uint32(OPEM)}}
}

func (p *USB_OTG_Device_Periph) BOIM() RMDOEPMSK {
	return RMDOEPMSK{mmio.UM32{&p.DOEPMSK.U32, uint32(BOIM)}}
}

type DAINT uint32

func (b DAINT) Field(mask DAINT) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DAINT) J(v int) DAINT {
	return DAINT(bits.MakeField32(v, uint32(mask)))
}

type RDAINT struct{ mmio.U32 }

func (r *RDAINT) Bits(mask DAINT) DAINT   { return DAINT(r.U32.Bits(uint32(mask))) }
func (r *RDAINT) StoreBits(mask, b DAINT) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDAINT) SetBits(mask DAINT)      { r.U32.SetBits(uint32(mask)) }
func (r *RDAINT) ClearBits(mask DAINT)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDAINT) Load() DAINT             { return DAINT(r.U32.Load()) }
func (r *RDAINT) Store(b DAINT)           { r.U32.Store(uint32(b)) }

func (r *RDAINT) AtomicStoreBits(mask, b DAINT) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDAINT) AtomicSetBits(mask DAINT)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDAINT) AtomicClearBits(mask DAINT)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDAINT struct{ mmio.UM32 }

func (rm RMDAINT) Load() DAINT   { return DAINT(rm.UM32.Load()) }
func (rm RMDAINT) Store(b DAINT) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) IEPINT() RMDAINT {
	return RMDAINT{mmio.UM32{&p.DAINT.U32, uint32(IEPINT)}}
}

func (p *USB_OTG_Device_Periph) OEPINT() RMDAINT {
	return RMDAINT{mmio.UM32{&p.DAINT.U32, uint32(OEPINT)}}
}

type DAINTMSK uint32

func (b DAINTMSK) Field(mask DAINTMSK) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DAINTMSK) J(v int) DAINTMSK {
	return DAINTMSK(bits.MakeField32(v, uint32(mask)))
}

type RDAINTMSK struct{ mmio.U32 }

func (r *RDAINTMSK) Bits(mask DAINTMSK) DAINTMSK { return DAINTMSK(r.U32.Bits(uint32(mask))) }
func (r *RDAINTMSK) StoreBits(mask, b DAINTMSK)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDAINTMSK) SetBits(mask DAINTMSK)       { r.U32.SetBits(uint32(mask)) }
func (r *RDAINTMSK) ClearBits(mask DAINTMSK)     { r.U32.ClearBits(uint32(mask)) }
func (r *RDAINTMSK) Load() DAINTMSK              { return DAINTMSK(r.U32.Load()) }
func (r *RDAINTMSK) Store(b DAINTMSK)            { r.U32.Store(uint32(b)) }

func (r *RDAINTMSK) AtomicStoreBits(mask, b DAINTMSK) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDAINTMSK) AtomicSetBits(mask DAINTMSK)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDAINTMSK) AtomicClearBits(mask DAINTMSK)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDAINTMSK struct{ mmio.UM32 }

func (rm RMDAINTMSK) Load() DAINTMSK   { return DAINTMSK(rm.UM32.Load()) }
func (rm RMDAINTMSK) Store(b DAINTMSK) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) IEPM() RMDAINTMSK {
	return RMDAINTMSK{mmio.UM32{&p.DAINTMSK.U32, uint32(IEPM)}}
}

func (p *USB_OTG_Device_Periph) OEPM() RMDAINTMSK {
	return RMDAINTMSK{mmio.UM32{&p.DAINTMSK.U32, uint32(OEPM)}}
}

type DVBUSDIS uint32

func (b DVBUSDIS) Field(mask DVBUSDIS) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DVBUSDIS) J(v int) DVBUSDIS {
	return DVBUSDIS(bits.MakeField32(v, uint32(mask)))
}

type RDVBUSDIS struct{ mmio.U32 }

func (r *RDVBUSDIS) Bits(mask DVBUSDIS) DVBUSDIS { return DVBUSDIS(r.U32.Bits(uint32(mask))) }
func (r *RDVBUSDIS) StoreBits(mask, b DVBUSDIS)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDVBUSDIS) SetBits(mask DVBUSDIS)       { r.U32.SetBits(uint32(mask)) }
func (r *RDVBUSDIS) ClearBits(mask DVBUSDIS)     { r.U32.ClearBits(uint32(mask)) }
func (r *RDVBUSDIS) Load() DVBUSDIS              { return DVBUSDIS(r.U32.Load()) }
func (r *RDVBUSDIS) Store(b DVBUSDIS)            { r.U32.Store(uint32(b)) }

func (r *RDVBUSDIS) AtomicStoreBits(mask, b DVBUSDIS) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDVBUSDIS) AtomicSetBits(mask DVBUSDIS)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDVBUSDIS) AtomicClearBits(mask DVBUSDIS)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDVBUSDIS struct{ mmio.UM32 }

func (rm RMDVBUSDIS) Load() DVBUSDIS   { return DVBUSDIS(rm.UM32.Load()) }
func (rm RMDVBUSDIS) Store(b DVBUSDIS) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) VBUSDT() RMDVBUSDIS {
	return RMDVBUSDIS{mmio.UM32{&p.DVBUSDIS.U32, uint32(VBUSDT)}}
}

type DVBUSPULSE uint32

func (b DVBUSPULSE) Field(mask DVBUSPULSE) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DVBUSPULSE) J(v int) DVBUSPULSE {
	return DVBUSPULSE(bits.MakeField32(v, uint32(mask)))
}

type RDVBUSPULSE struct{ mmio.U32 }

func (r *RDVBUSPULSE) Bits(mask DVBUSPULSE) DVBUSPULSE { return DVBUSPULSE(r.U32.Bits(uint32(mask))) }
func (r *RDVBUSPULSE) StoreBits(mask, b DVBUSPULSE)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDVBUSPULSE) SetBits(mask DVBUSPULSE)         { r.U32.SetBits(uint32(mask)) }
func (r *RDVBUSPULSE) ClearBits(mask DVBUSPULSE)       { r.U32.ClearBits(uint32(mask)) }
func (r *RDVBUSPULSE) Load() DVBUSPULSE                { return DVBUSPULSE(r.U32.Load()) }
func (r *RDVBUSPULSE) Store(b DVBUSPULSE)              { r.U32.Store(uint32(b)) }

func (r *RDVBUSPULSE) AtomicStoreBits(mask, b DVBUSPULSE) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RDVBUSPULSE) AtomicSetBits(mask DVBUSPULSE)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDVBUSPULSE) AtomicClearBits(mask DVBUSPULSE) { r.U32.AtomicClearBits(uint32(mask)) }

type RMDVBUSPULSE struct{ mmio.UM32 }

func (rm RMDVBUSPULSE) Load() DVBUSPULSE   { return DVBUSPULSE(rm.UM32.Load()) }
func (rm RMDVBUSPULSE) Store(b DVBUSPULSE) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) DVBUSP() RMDVBUSPULSE {
	return RMDVBUSPULSE{mmio.UM32{&p.DVBUSPULSE.U32, uint32(DVBUSP)}}
}

type DTHRCTL uint32

func (b DTHRCTL) Field(mask DTHRCTL) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DTHRCTL) J(v int) DTHRCTL {
	return DTHRCTL(bits.MakeField32(v, uint32(mask)))
}

type RDTHRCTL struct{ mmio.U32 }

func (r *RDTHRCTL) Bits(mask DTHRCTL) DTHRCTL { return DTHRCTL(r.U32.Bits(uint32(mask))) }
func (r *RDTHRCTL) StoreBits(mask, b DTHRCTL) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDTHRCTL) SetBits(mask DTHRCTL)      { r.U32.SetBits(uint32(mask)) }
func (r *RDTHRCTL) ClearBits(mask DTHRCTL)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDTHRCTL) Load() DTHRCTL             { return DTHRCTL(r.U32.Load()) }
func (r *RDTHRCTL) Store(b DTHRCTL)           { r.U32.Store(uint32(b)) }

func (r *RDTHRCTL) AtomicStoreBits(mask, b DTHRCTL) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDTHRCTL) AtomicSetBits(mask DTHRCTL)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDTHRCTL) AtomicClearBits(mask DTHRCTL)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDTHRCTL struct{ mmio.UM32 }

func (rm RMDTHRCTL) Load() DTHRCTL   { return DTHRCTL(rm.UM32.Load()) }
func (rm RMDTHRCTL) Store(b DTHRCTL) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) NONISOTHREN() RMDTHRCTL {
	return RMDTHRCTL{mmio.UM32{&p.DTHRCTL.U32, uint32(NONISOTHREN)}}
}

func (p *USB_OTG_Device_Periph) ISOTHREN() RMDTHRCTL {
	return RMDTHRCTL{mmio.UM32{&p.DTHRCTL.U32, uint32(ISOTHREN)}}
}

func (p *USB_OTG_Device_Periph) TXTHRLEN() RMDTHRCTL {
	return RMDTHRCTL{mmio.UM32{&p.DTHRCTL.U32, uint32(TXTHRLEN)}}
}

func (p *USB_OTG_Device_Periph) RXTHREN() RMDTHRCTL {
	return RMDTHRCTL{mmio.UM32{&p.DTHRCTL.U32, uint32(RXTHREN)}}
}

func (p *USB_OTG_Device_Periph) RXTHRLEN() RMDTHRCTL {
	return RMDTHRCTL{mmio.UM32{&p.DTHRCTL.U32, uint32(RXTHRLEN)}}
}

func (p *USB_OTG_Device_Periph) ARPEN() RMDTHRCTL {
	return RMDTHRCTL{mmio.UM32{&p.DTHRCTL.U32, uint32(ARPEN)}}
}

type DIEPEMPMSK uint32

func (b DIEPEMPMSK) Field(mask DIEPEMPMSK) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DIEPEMPMSK) J(v int) DIEPEMPMSK {
	return DIEPEMPMSK(bits.MakeField32(v, uint32(mask)))
}

type RDIEPEMPMSK struct{ mmio.U32 }

func (r *RDIEPEMPMSK) Bits(mask DIEPEMPMSK) DIEPEMPMSK { return DIEPEMPMSK(r.U32.Bits(uint32(mask))) }
func (r *RDIEPEMPMSK) StoreBits(mask, b DIEPEMPMSK)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDIEPEMPMSK) SetBits(mask DIEPEMPMSK)         { r.U32.SetBits(uint32(mask)) }
func (r *RDIEPEMPMSK) ClearBits(mask DIEPEMPMSK)       { r.U32.ClearBits(uint32(mask)) }
func (r *RDIEPEMPMSK) Load() DIEPEMPMSK                { return DIEPEMPMSK(r.U32.Load()) }
func (r *RDIEPEMPMSK) Store(b DIEPEMPMSK)              { r.U32.Store(uint32(b)) }

func (r *RDIEPEMPMSK) AtomicStoreBits(mask, b DIEPEMPMSK) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RDIEPEMPMSK) AtomicSetBits(mask DIEPEMPMSK)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDIEPEMPMSK) AtomicClearBits(mask DIEPEMPMSK) { r.U32.AtomicClearBits(uint32(mask)) }

type RMDIEPEMPMSK struct{ mmio.UM32 }

func (rm RMDIEPEMPMSK) Load() DIEPEMPMSK   { return DIEPEMPMSK(rm.UM32.Load()) }
func (rm RMDIEPEMPMSK) Store(b DIEPEMPMSK) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) INEPTXFEM() RMDIEPEMPMSK {
	return RMDIEPEMPMSK{mmio.UM32{&p.DIEPEMPMSK.U32, uint32(INEPTXFEM)}}
}

type DEACHINT uint32

func (b DEACHINT) Field(mask DEACHINT) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DEACHINT) J(v int) DEACHINT {
	return DEACHINT(bits.MakeField32(v, uint32(mask)))
}

type RDEACHINT struct{ mmio.U32 }

func (r *RDEACHINT) Bits(mask DEACHINT) DEACHINT { return DEACHINT(r.U32.Bits(uint32(mask))) }
func (r *RDEACHINT) StoreBits(mask, b DEACHINT)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDEACHINT) SetBits(mask DEACHINT)       { r.U32.SetBits(uint32(mask)) }
func (r *RDEACHINT) ClearBits(mask DEACHINT)     { r.U32.ClearBits(uint32(mask)) }
func (r *RDEACHINT) Load() DEACHINT              { return DEACHINT(r.U32.Load()) }
func (r *RDEACHINT) Store(b DEACHINT)            { r.U32.Store(uint32(b)) }

func (r *RDEACHINT) AtomicStoreBits(mask, b DEACHINT) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDEACHINT) AtomicSetBits(mask DEACHINT)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDEACHINT) AtomicClearBits(mask DEACHINT)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDEACHINT struct{ mmio.UM32 }

func (rm RMDEACHINT) Load() DEACHINT   { return DEACHINT(rm.UM32.Load()) }
func (rm RMDEACHINT) Store(b DEACHINT) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) IEP1INT() RMDEACHINT {
	return RMDEACHINT{mmio.UM32{&p.DEACHINT.U32, uint32(IEP1INT)}}
}

func (p *USB_OTG_Device_Periph) OEP1INT() RMDEACHINT {
	return RMDEACHINT{mmio.UM32{&p.DEACHINT.U32, uint32(OEP1INT)}}
}

type DEACHMSK uint32

func (b DEACHMSK) Field(mask DEACHMSK) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DEACHMSK) J(v int) DEACHMSK {
	return DEACHMSK(bits.MakeField32(v, uint32(mask)))
}

type RDEACHMSK struct{ mmio.U32 }

func (r *RDEACHMSK) Bits(mask DEACHMSK) DEACHMSK { return DEACHMSK(r.U32.Bits(uint32(mask))) }
func (r *RDEACHMSK) StoreBits(mask, b DEACHMSK)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDEACHMSK) SetBits(mask DEACHMSK)       { r.U32.SetBits(uint32(mask)) }
func (r *RDEACHMSK) ClearBits(mask DEACHMSK)     { r.U32.ClearBits(uint32(mask)) }
func (r *RDEACHMSK) Load() DEACHMSK              { return DEACHMSK(r.U32.Load()) }
func (r *RDEACHMSK) Store(b DEACHMSK)            { r.U32.Store(uint32(b)) }

func (r *RDEACHMSK) AtomicStoreBits(mask, b DEACHMSK) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDEACHMSK) AtomicSetBits(mask DEACHMSK)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDEACHMSK) AtomicClearBits(mask DEACHMSK)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDEACHMSK struct{ mmio.UM32 }

func (rm RMDEACHMSK) Load() DEACHMSK   { return DEACHMSK(rm.UM32.Load()) }
func (rm RMDEACHMSK) Store(b DEACHMSK) { rm.UM32.Store(uint32(b)) }

type DINEP1MSK uint32

func (b DINEP1MSK) Field(mask DINEP1MSK) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DINEP1MSK) J(v int) DINEP1MSK {
	return DINEP1MSK(bits.MakeField32(v, uint32(mask)))
}

type RDINEP1MSK struct{ mmio.U32 }

func (r *RDINEP1MSK) Bits(mask DINEP1MSK) DINEP1MSK { return DINEP1MSK(r.U32.Bits(uint32(mask))) }
func (r *RDINEP1MSK) StoreBits(mask, b DINEP1MSK)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDINEP1MSK) SetBits(mask DINEP1MSK)        { r.U32.SetBits(uint32(mask)) }
func (r *RDINEP1MSK) ClearBits(mask DINEP1MSK)      { r.U32.ClearBits(uint32(mask)) }
func (r *RDINEP1MSK) Load() DINEP1MSK               { return DINEP1MSK(r.U32.Load()) }
func (r *RDINEP1MSK) Store(b DINEP1MSK)             { r.U32.Store(uint32(b)) }

func (r *RDINEP1MSK) AtomicStoreBits(mask, b DINEP1MSK) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RDINEP1MSK) AtomicSetBits(mask DINEP1MSK)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDINEP1MSK) AtomicClearBits(mask DINEP1MSK) { r.U32.AtomicClearBits(uint32(mask)) }

type RMDINEP1MSK struct{ mmio.UM32 }

func (rm RMDINEP1MSK) Load() DINEP1MSK   { return DINEP1MSK(rm.UM32.Load()) }
func (rm RMDINEP1MSK) Store(b DINEP1MSK) { rm.UM32.Store(uint32(b)) }

type DOUTEP1MSK uint32

func (b DOUTEP1MSK) Field(mask DOUTEP1MSK) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DOUTEP1MSK) J(v int) DOUTEP1MSK {
	return DOUTEP1MSK(bits.MakeField32(v, uint32(mask)))
}

type RDOUTEP1MSK struct{ mmio.U32 }

func (r *RDOUTEP1MSK) Bits(mask DOUTEP1MSK) DOUTEP1MSK { return DOUTEP1MSK(r.U32.Bits(uint32(mask))) }
func (r *RDOUTEP1MSK) StoreBits(mask, b DOUTEP1MSK)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDOUTEP1MSK) SetBits(mask DOUTEP1MSK)         { r.U32.SetBits(uint32(mask)) }
func (r *RDOUTEP1MSK) ClearBits(mask DOUTEP1MSK)       { r.U32.ClearBits(uint32(mask)) }
func (r *RDOUTEP1MSK) Load() DOUTEP1MSK                { return DOUTEP1MSK(r.U32.Load()) }
func (r *RDOUTEP1MSK) Store(b DOUTEP1MSK)              { r.U32.Store(uint32(b)) }

func (r *RDOUTEP1MSK) AtomicStoreBits(mask, b DOUTEP1MSK) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RDOUTEP1MSK) AtomicSetBits(mask DOUTEP1MSK)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDOUTEP1MSK) AtomicClearBits(mask DOUTEP1MSK) { r.U32.AtomicClearBits(uint32(mask)) }

type RMDOUTEP1MSK struct{ mmio.UM32 }

func (rm RMDOUTEP1MSK) Load() DOUTEP1MSK   { return DOUTEP1MSK(rm.UM32.Load()) }
func (rm RMDOUTEP1MSK) Store(b DOUTEP1MSK) { rm.UM32.Store(uint32(b)) }
