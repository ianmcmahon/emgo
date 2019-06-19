package fmc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f469xx/mmap"
)

type FMC_Bank5_6_Periph struct {
	SDCR  [2]RSDCR
	SDTR  [2]RSDTR
	SDCMR RSDCMR
	SDRTR RSDRTR
	SDSR  RSDSR
}

func (p *FMC_Bank5_6_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var FMC_Bank5_6 = (*FMC_Bank5_6_Periph)(unsafe.Pointer(uintptr(mmap.FMC_Bank5_6_R_BASE)))

type SDCR uint32

func (b SDCR) Field(mask SDCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SDCR) J(v int) SDCR {
	return SDCR(bits.MakeField32(v, uint32(mask)))
}

type RSDCR struct{ mmio.U32 }

func (r *RSDCR) Bits(mask SDCR) SDCR    { return SDCR(r.U32.Bits(uint32(mask))) }
func (r *RSDCR) StoreBits(mask, b SDCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSDCR) SetBits(mask SDCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSDCR) ClearBits(mask SDCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSDCR) Load() SDCR             { return SDCR(r.U32.Load()) }
func (r *RSDCR) Store(b SDCR)           { r.U32.Store(uint32(b)) }

func (r *RSDCR) AtomicStoreBits(mask, b SDCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSDCR) AtomicSetBits(mask SDCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSDCR) AtomicClearBits(mask SDCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSDCR struct{ mmio.UM32 }

func (rm RMSDCR) Load() SDCR   { return SDCR(rm.UM32.Load()) }
func (rm RMSDCR) Store(b SDCR) { rm.UM32.Store(uint32(b)) }

func (p *FMC_Bank5_6_Periph) NC(n int) RMSDCR {
	return RMSDCR{mmio.UM32{&p.SDCR[n].U32, uint32(NC)}}
}

func (p *FMC_Bank5_6_Periph) NR(n int) RMSDCR {
	return RMSDCR{mmio.UM32{&p.SDCR[n].U32, uint32(NR)}}
}

func (p *FMC_Bank5_6_Periph) SDMWID(n int) RMSDCR {
	return RMSDCR{mmio.UM32{&p.SDCR[n].U32, uint32(SDMWID)}}
}

func (p *FMC_Bank5_6_Periph) NB(n int) RMSDCR {
	return RMSDCR{mmio.UM32{&p.SDCR[n].U32, uint32(NB)}}
}

func (p *FMC_Bank5_6_Periph) CAS(n int) RMSDCR {
	return RMSDCR{mmio.UM32{&p.SDCR[n].U32, uint32(CAS)}}
}

func (p *FMC_Bank5_6_Periph) WP(n int) RMSDCR {
	return RMSDCR{mmio.UM32{&p.SDCR[n].U32, uint32(WP)}}
}

func (p *FMC_Bank5_6_Periph) SDCLK(n int) RMSDCR {
	return RMSDCR{mmio.UM32{&p.SDCR[n].U32, uint32(SDCLK)}}
}

func (p *FMC_Bank5_6_Periph) RBURST(n int) RMSDCR {
	return RMSDCR{mmio.UM32{&p.SDCR[n].U32, uint32(RBURST)}}
}

func (p *FMC_Bank5_6_Periph) RPIPE(n int) RMSDCR {
	return RMSDCR{mmio.UM32{&p.SDCR[n].U32, uint32(RPIPE)}}
}

type SDTR uint32

func (b SDTR) Field(mask SDTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SDTR) J(v int) SDTR {
	return SDTR(bits.MakeField32(v, uint32(mask)))
}

type RSDTR struct{ mmio.U32 }

func (r *RSDTR) Bits(mask SDTR) SDTR    { return SDTR(r.U32.Bits(uint32(mask))) }
func (r *RSDTR) StoreBits(mask, b SDTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSDTR) SetBits(mask SDTR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSDTR) ClearBits(mask SDTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSDTR) Load() SDTR             { return SDTR(r.U32.Load()) }
func (r *RSDTR) Store(b SDTR)           { r.U32.Store(uint32(b)) }

func (r *RSDTR) AtomicStoreBits(mask, b SDTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSDTR) AtomicSetBits(mask SDTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSDTR) AtomicClearBits(mask SDTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSDTR struct{ mmio.UM32 }

func (rm RMSDTR) Load() SDTR   { return SDTR(rm.UM32.Load()) }
func (rm RMSDTR) Store(b SDTR) { rm.UM32.Store(uint32(b)) }

func (p *FMC_Bank5_6_Periph) TMRD(n int) RMSDTR {
	return RMSDTR{mmio.UM32{&p.SDTR[n].U32, uint32(TMRD)}}
}

func (p *FMC_Bank5_6_Periph) TXSR(n int) RMSDTR {
	return RMSDTR{mmio.UM32{&p.SDTR[n].U32, uint32(TXSR)}}
}

func (p *FMC_Bank5_6_Periph) TRAS(n int) RMSDTR {
	return RMSDTR{mmio.UM32{&p.SDTR[n].U32, uint32(TRAS)}}
}

func (p *FMC_Bank5_6_Periph) TRC(n int) RMSDTR {
	return RMSDTR{mmio.UM32{&p.SDTR[n].U32, uint32(TRC)}}
}

func (p *FMC_Bank5_6_Periph) TWR(n int) RMSDTR {
	return RMSDTR{mmio.UM32{&p.SDTR[n].U32, uint32(TWR)}}
}

func (p *FMC_Bank5_6_Periph) TRP(n int) RMSDTR {
	return RMSDTR{mmio.UM32{&p.SDTR[n].U32, uint32(TRP)}}
}

func (p *FMC_Bank5_6_Periph) TRCD(n int) RMSDTR {
	return RMSDTR{mmio.UM32{&p.SDTR[n].U32, uint32(TRCD)}}
}

type SDCMR uint32

func (b SDCMR) Field(mask SDCMR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SDCMR) J(v int) SDCMR {
	return SDCMR(bits.MakeField32(v, uint32(mask)))
}

type RSDCMR struct{ mmio.U32 }

func (r *RSDCMR) Bits(mask SDCMR) SDCMR   { return SDCMR(r.U32.Bits(uint32(mask))) }
func (r *RSDCMR) StoreBits(mask, b SDCMR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSDCMR) SetBits(mask SDCMR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSDCMR) ClearBits(mask SDCMR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSDCMR) Load() SDCMR             { return SDCMR(r.U32.Load()) }
func (r *RSDCMR) Store(b SDCMR)           { r.U32.Store(uint32(b)) }

func (r *RSDCMR) AtomicStoreBits(mask, b SDCMR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSDCMR) AtomicSetBits(mask SDCMR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSDCMR) AtomicClearBits(mask SDCMR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSDCMR struct{ mmio.UM32 }

func (rm RMSDCMR) Load() SDCMR   { return SDCMR(rm.UM32.Load()) }
func (rm RMSDCMR) Store(b SDCMR) { rm.UM32.Store(uint32(b)) }

func (p *FMC_Bank5_6_Periph) MODE() RMSDCMR {
	return RMSDCMR{mmio.UM32{&p.SDCMR.U32, uint32(MODE)}}
}

func (p *FMC_Bank5_6_Periph) CTB2() RMSDCMR {
	return RMSDCMR{mmio.UM32{&p.SDCMR.U32, uint32(CTB2)}}
}

func (p *FMC_Bank5_6_Periph) CTB1() RMSDCMR {
	return RMSDCMR{mmio.UM32{&p.SDCMR.U32, uint32(CTB1)}}
}

func (p *FMC_Bank5_6_Periph) NRFS() RMSDCMR {
	return RMSDCMR{mmio.UM32{&p.SDCMR.U32, uint32(NRFS)}}
}

func (p *FMC_Bank5_6_Periph) MRD() RMSDCMR {
	return RMSDCMR{mmio.UM32{&p.SDCMR.U32, uint32(MRD)}}
}

type SDRTR uint32

func (b SDRTR) Field(mask SDRTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SDRTR) J(v int) SDRTR {
	return SDRTR(bits.MakeField32(v, uint32(mask)))
}

type RSDRTR struct{ mmio.U32 }

func (r *RSDRTR) Bits(mask SDRTR) SDRTR   { return SDRTR(r.U32.Bits(uint32(mask))) }
func (r *RSDRTR) StoreBits(mask, b SDRTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSDRTR) SetBits(mask SDRTR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSDRTR) ClearBits(mask SDRTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSDRTR) Load() SDRTR             { return SDRTR(r.U32.Load()) }
func (r *RSDRTR) Store(b SDRTR)           { r.U32.Store(uint32(b)) }

func (r *RSDRTR) AtomicStoreBits(mask, b SDRTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSDRTR) AtomicSetBits(mask SDRTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSDRTR) AtomicClearBits(mask SDRTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSDRTR struct{ mmio.UM32 }

func (rm RMSDRTR) Load() SDRTR   { return SDRTR(rm.UM32.Load()) }
func (rm RMSDRTR) Store(b SDRTR) { rm.UM32.Store(uint32(b)) }

func (p *FMC_Bank5_6_Periph) CRE() RMSDRTR {
	return RMSDRTR{mmio.UM32{&p.SDRTR.U32, uint32(CRE)}}
}

func (p *FMC_Bank5_6_Periph) COUNT() RMSDRTR {
	return RMSDRTR{mmio.UM32{&p.SDRTR.U32, uint32(COUNT)}}
}

func (p *FMC_Bank5_6_Periph) REIE() RMSDRTR {
	return RMSDRTR{mmio.UM32{&p.SDRTR.U32, uint32(REIE)}}
}

type SDSR uint32

func (b SDSR) Field(mask SDSR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SDSR) J(v int) SDSR {
	return SDSR(bits.MakeField32(v, uint32(mask)))
}

type RSDSR struct{ mmio.U32 }

func (r *RSDSR) Bits(mask SDSR) SDSR    { return SDSR(r.U32.Bits(uint32(mask))) }
func (r *RSDSR) StoreBits(mask, b SDSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSDSR) SetBits(mask SDSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSDSR) ClearBits(mask SDSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSDSR) Load() SDSR             { return SDSR(r.U32.Load()) }
func (r *RSDSR) Store(b SDSR)           { r.U32.Store(uint32(b)) }

func (r *RSDSR) AtomicStoreBits(mask, b SDSR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSDSR) AtomicSetBits(mask SDSR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSDSR) AtomicClearBits(mask SDSR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSDSR struct{ mmio.UM32 }

func (rm RMSDSR) Load() SDSR   { return SDSR(rm.UM32.Load()) }
func (rm RMSDSR) Store(b SDSR) { rm.UM32.Store(uint32(b)) }

func (p *FMC_Bank5_6_Periph) RE() RMSDSR {
	return RMSDSR{mmio.UM32{&p.SDSR.U32, uint32(RE)}}
}

func (p *FMC_Bank5_6_Periph) MODES1() RMSDSR {
	return RMSDSR{mmio.UM32{&p.SDSR.U32, uint32(MODES1)}}
}

func (p *FMC_Bank5_6_Periph) MODES2() RMSDSR {
	return RMSDSR{mmio.UM32{&p.SDSR.U32, uint32(MODES2)}}
}

func (p *FMC_Bank5_6_Periph) BUSY() RMSDSR {
	return RMSDSR{mmio.UM32{&p.SDSR.U32, uint32(BUSY)}}
}
