package ltdc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"/Users/ian/code/src/github.com/ianmcmahon/emgo/egpath/src/stm32/o/f469xx/mmap"
)

type LTDC_Periph struct {
	_     [2]uint32
	SSCR  RSSCR
	BPCR  RBPCR
	AWCR  RAWCR
	TWCR  RTWCR
	GCR   RGCR
	_     [2]uint32
	SRCR  RSRCR
	_     uint32
	BCCR  RBCCR
	_     uint32
	IER   RIER
	ISR   RISR
	ICR   RICR
	LIPCR RLIPCR
	CPSR  RCPSR
	CDSR  RCDSR
}

func (p *LTDC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var LTDC = (*LTDC_Periph)(unsafe.Pointer(uintptr(mmap.LTDC_BASE)))

type SSCR uint32

func (b SSCR) Field(mask SSCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SSCR) J(v int) SSCR {
	return SSCR(bits.MakeField32(v, uint32(mask)))
}

type RSSCR struct{ mmio.U32 }

func (r *RSSCR) Bits(mask SSCR) SSCR    { return SSCR(r.U32.Bits(uint32(mask))) }
func (r *RSSCR) StoreBits(mask, b SSCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSSCR) SetBits(mask SSCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSSCR) ClearBits(mask SSCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSSCR) Load() SSCR             { return SSCR(r.U32.Load()) }
func (r *RSSCR) Store(b SSCR)           { r.U32.Store(uint32(b)) }

func (r *RSSCR) AtomicStoreBits(mask, b SSCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSSCR) AtomicSetBits(mask SSCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSSCR) AtomicClearBits(mask SSCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSSCR struct{ mmio.UM32 }

func (rm RMSSCR) Load() SSCR   { return SSCR(rm.UM32.Load()) }
func (rm RMSSCR) Store(b SSCR) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) VSH() RMSSCR {
	return RMSSCR{mmio.UM32{&p.SSCR.U32, uint32(VSH)}}
}

func (p *LTDC_Periph) HSW() RMSSCR {
	return RMSSCR{mmio.UM32{&p.SSCR.U32, uint32(HSW)}}
}

type BPCR uint32

func (b BPCR) Field(mask BPCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BPCR) J(v int) BPCR {
	return BPCR(bits.MakeField32(v, uint32(mask)))
}

type RBPCR struct{ mmio.U32 }

func (r *RBPCR) Bits(mask BPCR) BPCR    { return BPCR(r.U32.Bits(uint32(mask))) }
func (r *RBPCR) StoreBits(mask, b BPCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBPCR) SetBits(mask BPCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RBPCR) ClearBits(mask BPCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RBPCR) Load() BPCR             { return BPCR(r.U32.Load()) }
func (r *RBPCR) Store(b BPCR)           { r.U32.Store(uint32(b)) }

func (r *RBPCR) AtomicStoreBits(mask, b BPCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RBPCR) AtomicSetBits(mask BPCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RBPCR) AtomicClearBits(mask BPCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMBPCR struct{ mmio.UM32 }

func (rm RMBPCR) Load() BPCR   { return BPCR(rm.UM32.Load()) }
func (rm RMBPCR) Store(b BPCR) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) AVBP() RMBPCR {
	return RMBPCR{mmio.UM32{&p.BPCR.U32, uint32(AVBP)}}
}

func (p *LTDC_Periph) AHBP() RMBPCR {
	return RMBPCR{mmio.UM32{&p.BPCR.U32, uint32(AHBP)}}
}

type AWCR uint32

func (b AWCR) Field(mask AWCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AWCR) J(v int) AWCR {
	return AWCR(bits.MakeField32(v, uint32(mask)))
}

type RAWCR struct{ mmio.U32 }

func (r *RAWCR) Bits(mask AWCR) AWCR    { return AWCR(r.U32.Bits(uint32(mask))) }
func (r *RAWCR) StoreBits(mask, b AWCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAWCR) SetBits(mask AWCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAWCR) ClearBits(mask AWCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAWCR) Load() AWCR             { return AWCR(r.U32.Load()) }
func (r *RAWCR) Store(b AWCR)           { r.U32.Store(uint32(b)) }

func (r *RAWCR) AtomicStoreBits(mask, b AWCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAWCR) AtomicSetBits(mask AWCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAWCR) AtomicClearBits(mask AWCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAWCR struct{ mmio.UM32 }

func (rm RMAWCR) Load() AWCR   { return AWCR(rm.UM32.Load()) }
func (rm RMAWCR) Store(b AWCR) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) AAH() RMAWCR {
	return RMAWCR{mmio.UM32{&p.AWCR.U32, uint32(AAH)}}
}

func (p *LTDC_Periph) AAW() RMAWCR {
	return RMAWCR{mmio.UM32{&p.AWCR.U32, uint32(AAW)}}
}

type TWCR uint32

func (b TWCR) Field(mask TWCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TWCR) J(v int) TWCR {
	return TWCR(bits.MakeField32(v, uint32(mask)))
}

type RTWCR struct{ mmio.U32 }

func (r *RTWCR) Bits(mask TWCR) TWCR    { return TWCR(r.U32.Bits(uint32(mask))) }
func (r *RTWCR) StoreBits(mask, b TWCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTWCR) SetBits(mask TWCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RTWCR) ClearBits(mask TWCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTWCR) Load() TWCR             { return TWCR(r.U32.Load()) }
func (r *RTWCR) Store(b TWCR)           { r.U32.Store(uint32(b)) }

func (r *RTWCR) AtomicStoreBits(mask, b TWCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTWCR) AtomicSetBits(mask TWCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTWCR) AtomicClearBits(mask TWCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTWCR struct{ mmio.UM32 }

func (rm RMTWCR) Load() TWCR   { return TWCR(rm.UM32.Load()) }
func (rm RMTWCR) Store(b TWCR) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) TOTALH() RMTWCR {
	return RMTWCR{mmio.UM32{&p.TWCR.U32, uint32(TOTALH)}}
}

func (p *LTDC_Periph) TOTALW() RMTWCR {
	return RMTWCR{mmio.UM32{&p.TWCR.U32, uint32(TOTALW)}}
}

type GCR uint32

func (b GCR) Field(mask GCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask GCR) J(v int) GCR {
	return GCR(bits.MakeField32(v, uint32(mask)))
}

type RGCR struct{ mmio.U32 }

func (r *RGCR) Bits(mask GCR) GCR     { return GCR(r.U32.Bits(uint32(mask))) }
func (r *RGCR) StoreBits(mask, b GCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RGCR) SetBits(mask GCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RGCR) ClearBits(mask GCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RGCR) Load() GCR             { return GCR(r.U32.Load()) }
func (r *RGCR) Store(b GCR)           { r.U32.Store(uint32(b)) }

func (r *RGCR) AtomicStoreBits(mask, b GCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RGCR) AtomicSetBits(mask GCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RGCR) AtomicClearBits(mask GCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMGCR struct{ mmio.UM32 }

func (rm RMGCR) Load() GCR   { return GCR(rm.UM32.Load()) }
func (rm RMGCR) Store(b GCR) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) LTDCEN() RMGCR {
	return RMGCR{mmio.UM32{&p.GCR.U32, uint32(LTDCEN)}}
}

func (p *LTDC_Periph) DBW() RMGCR {
	return RMGCR{mmio.UM32{&p.GCR.U32, uint32(DBW)}}
}

func (p *LTDC_Periph) DGW() RMGCR {
	return RMGCR{mmio.UM32{&p.GCR.U32, uint32(DGW)}}
}

func (p *LTDC_Periph) DRW() RMGCR {
	return RMGCR{mmio.UM32{&p.GCR.U32, uint32(DRW)}}
}

func (p *LTDC_Periph) DEN() RMGCR {
	return RMGCR{mmio.UM32{&p.GCR.U32, uint32(DEN)}}
}

func (p *LTDC_Periph) PCPOL() RMGCR {
	return RMGCR{mmio.UM32{&p.GCR.U32, uint32(PCPOL)}}
}

func (p *LTDC_Periph) DEPOL() RMGCR {
	return RMGCR{mmio.UM32{&p.GCR.U32, uint32(DEPOL)}}
}

func (p *LTDC_Periph) VSPOL() RMGCR {
	return RMGCR{mmio.UM32{&p.GCR.U32, uint32(VSPOL)}}
}

func (p *LTDC_Periph) HSPOL() RMGCR {
	return RMGCR{mmio.UM32{&p.GCR.U32, uint32(HSPOL)}}
}

type SRCR uint32

func (b SRCR) Field(mask SRCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SRCR) J(v int) SRCR {
	return SRCR(bits.MakeField32(v, uint32(mask)))
}

type RSRCR struct{ mmio.U32 }

func (r *RSRCR) Bits(mask SRCR) SRCR    { return SRCR(r.U32.Bits(uint32(mask))) }
func (r *RSRCR) StoreBits(mask, b SRCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSRCR) SetBits(mask SRCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSRCR) ClearBits(mask SRCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSRCR) Load() SRCR             { return SRCR(r.U32.Load()) }
func (r *RSRCR) Store(b SRCR)           { r.U32.Store(uint32(b)) }

func (r *RSRCR) AtomicStoreBits(mask, b SRCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSRCR) AtomicSetBits(mask SRCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSRCR) AtomicClearBits(mask SRCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSRCR struct{ mmio.UM32 }

func (rm RMSRCR) Load() SRCR   { return SRCR(rm.UM32.Load()) }
func (rm RMSRCR) Store(b SRCR) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) IMR() RMSRCR {
	return RMSRCR{mmio.UM32{&p.SRCR.U32, uint32(IMR)}}
}

func (p *LTDC_Periph) VBR() RMSRCR {
	return RMSRCR{mmio.UM32{&p.SRCR.U32, uint32(VBR)}}
}

type BCCR uint32

func (b BCCR) Field(mask BCCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BCCR) J(v int) BCCR {
	return BCCR(bits.MakeField32(v, uint32(mask)))
}

type RBCCR struct{ mmio.U32 }

func (r *RBCCR) Bits(mask BCCR) BCCR    { return BCCR(r.U32.Bits(uint32(mask))) }
func (r *RBCCR) StoreBits(mask, b BCCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBCCR) SetBits(mask BCCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RBCCR) ClearBits(mask BCCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RBCCR) Load() BCCR             { return BCCR(r.U32.Load()) }
func (r *RBCCR) Store(b BCCR)           { r.U32.Store(uint32(b)) }

func (r *RBCCR) AtomicStoreBits(mask, b BCCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RBCCR) AtomicSetBits(mask BCCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RBCCR) AtomicClearBits(mask BCCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMBCCR struct{ mmio.UM32 }

func (rm RMBCCR) Load() BCCR   { return BCCR(rm.UM32.Load()) }
func (rm RMBCCR) Store(b BCCR) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) BCBLUE() RMBCCR {
	return RMBCCR{mmio.UM32{&p.BCCR.U32, uint32(BCBLUE)}}
}

func (p *LTDC_Periph) BCGREEN() RMBCCR {
	return RMBCCR{mmio.UM32{&p.BCCR.U32, uint32(BCGREEN)}}
}

func (p *LTDC_Periph) BCRED() RMBCCR {
	return RMBCCR{mmio.UM32{&p.BCCR.U32, uint32(BCRED)}}
}

type IER uint32

func (b IER) Field(mask IER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IER) J(v int) IER {
	return IER(bits.MakeField32(v, uint32(mask)))
}

type RIER struct{ mmio.U32 }

func (r *RIER) Bits(mask IER) IER     { return IER(r.U32.Bits(uint32(mask))) }
func (r *RIER) StoreBits(mask, b IER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIER) SetBits(mask IER)      { r.U32.SetBits(uint32(mask)) }
func (r *RIER) ClearBits(mask IER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RIER) Load() IER             { return IER(r.U32.Load()) }
func (r *RIER) Store(b IER)           { r.U32.Store(uint32(b)) }

func (r *RIER) AtomicStoreBits(mask, b IER) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RIER) AtomicSetBits(mask IER)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RIER) AtomicClearBits(mask IER)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMIER struct{ mmio.UM32 }

func (rm RMIER) Load() IER   { return IER(rm.UM32.Load()) }
func (rm RMIER) Store(b IER) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) LIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(LIE)}}
}

func (p *LTDC_Periph) FUIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(FUIE)}}
}

func (p *LTDC_Periph) TERRIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(TERRIE)}}
}

func (p *LTDC_Periph) RRIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(RRIE)}}
}

type ISR uint32

func (b ISR) Field(mask ISR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ISR) J(v int) ISR {
	return ISR(bits.MakeField32(v, uint32(mask)))
}

type RISR struct{ mmio.U32 }

func (r *RISR) Bits(mask ISR) ISR     { return ISR(r.U32.Bits(uint32(mask))) }
func (r *RISR) StoreBits(mask, b ISR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RISR) SetBits(mask ISR)      { r.U32.SetBits(uint32(mask)) }
func (r *RISR) ClearBits(mask ISR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RISR) Load() ISR             { return ISR(r.U32.Load()) }
func (r *RISR) Store(b ISR)           { r.U32.Store(uint32(b)) }

func (r *RISR) AtomicStoreBits(mask, b ISR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RISR) AtomicSetBits(mask ISR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RISR) AtomicClearBits(mask ISR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMISR struct{ mmio.UM32 }

func (rm RMISR) Load() ISR   { return ISR(rm.UM32.Load()) }
func (rm RMISR) Store(b ISR) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) LIF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(LIF)}}
}

func (p *LTDC_Periph) FUIF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(FUIF)}}
}

func (p *LTDC_Periph) TERRIF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TERRIF)}}
}

func (p *LTDC_Periph) RRIF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(RRIF)}}
}

type ICR uint32

func (b ICR) Field(mask ICR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ICR) J(v int) ICR {
	return ICR(bits.MakeField32(v, uint32(mask)))
}

type RICR struct{ mmio.U32 }

func (r *RICR) Bits(mask ICR) ICR     { return ICR(r.U32.Bits(uint32(mask))) }
func (r *RICR) StoreBits(mask, b ICR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RICR) SetBits(mask ICR)      { r.U32.SetBits(uint32(mask)) }
func (r *RICR) ClearBits(mask ICR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RICR) Load() ICR             { return ICR(r.U32.Load()) }
func (r *RICR) Store(b ICR)           { r.U32.Store(uint32(b)) }

func (r *RICR) AtomicStoreBits(mask, b ICR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RICR) AtomicSetBits(mask ICR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RICR) AtomicClearBits(mask ICR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMICR struct{ mmio.UM32 }

func (rm RMICR) Load() ICR   { return ICR(rm.UM32.Load()) }
func (rm RMICR) Store(b ICR) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) CLIF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CLIF)}}
}

func (p *LTDC_Periph) CFUIF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CFUIF)}}
}

func (p *LTDC_Periph) CTERRIF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CTERRIF)}}
}

func (p *LTDC_Periph) CRRIF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CRRIF)}}
}

type LIPCR uint32

func (b LIPCR) Field(mask LIPCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask LIPCR) J(v int) LIPCR {
	return LIPCR(bits.MakeField32(v, uint32(mask)))
}

type RLIPCR struct{ mmio.U32 }

func (r *RLIPCR) Bits(mask LIPCR) LIPCR   { return LIPCR(r.U32.Bits(uint32(mask))) }
func (r *RLIPCR) StoreBits(mask, b LIPCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RLIPCR) SetBits(mask LIPCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RLIPCR) ClearBits(mask LIPCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RLIPCR) Load() LIPCR             { return LIPCR(r.U32.Load()) }
func (r *RLIPCR) Store(b LIPCR)           { r.U32.Store(uint32(b)) }

func (r *RLIPCR) AtomicStoreBits(mask, b LIPCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RLIPCR) AtomicSetBits(mask LIPCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RLIPCR) AtomicClearBits(mask LIPCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMLIPCR struct{ mmio.UM32 }

func (rm RMLIPCR) Load() LIPCR   { return LIPCR(rm.UM32.Load()) }
func (rm RMLIPCR) Store(b LIPCR) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) LIPOS() RMLIPCR {
	return RMLIPCR{mmio.UM32{&p.LIPCR.U32, uint32(LIPOS)}}
}

type CPSR uint32

func (b CPSR) Field(mask CPSR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CPSR) J(v int) CPSR {
	return CPSR(bits.MakeField32(v, uint32(mask)))
}

type RCPSR struct{ mmio.U32 }

func (r *RCPSR) Bits(mask CPSR) CPSR    { return CPSR(r.U32.Bits(uint32(mask))) }
func (r *RCPSR) StoreBits(mask, b CPSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCPSR) SetBits(mask CPSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCPSR) ClearBits(mask CPSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCPSR) Load() CPSR             { return CPSR(r.U32.Load()) }
func (r *RCPSR) Store(b CPSR)           { r.U32.Store(uint32(b)) }

func (r *RCPSR) AtomicStoreBits(mask, b CPSR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCPSR) AtomicSetBits(mask CPSR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCPSR) AtomicClearBits(mask CPSR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCPSR struct{ mmio.UM32 }

func (rm RMCPSR) Load() CPSR   { return CPSR(rm.UM32.Load()) }
func (rm RMCPSR) Store(b CPSR) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) CYPOS() RMCPSR {
	return RMCPSR{mmio.UM32{&p.CPSR.U32, uint32(CYPOS)}}
}

func (p *LTDC_Periph) CXPOS() RMCPSR {
	return RMCPSR{mmio.UM32{&p.CPSR.U32, uint32(CXPOS)}}
}

type CDSR uint32

func (b CDSR) Field(mask CDSR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CDSR) J(v int) CDSR {
	return CDSR(bits.MakeField32(v, uint32(mask)))
}

type RCDSR struct{ mmio.U32 }

func (r *RCDSR) Bits(mask CDSR) CDSR    { return CDSR(r.U32.Bits(uint32(mask))) }
func (r *RCDSR) StoreBits(mask, b CDSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCDSR) SetBits(mask CDSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCDSR) ClearBits(mask CDSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCDSR) Load() CDSR             { return CDSR(r.U32.Load()) }
func (r *RCDSR) Store(b CDSR)           { r.U32.Store(uint32(b)) }

func (r *RCDSR) AtomicStoreBits(mask, b CDSR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCDSR) AtomicSetBits(mask CDSR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCDSR) AtomicClearBits(mask CDSR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCDSR struct{ mmio.UM32 }

func (rm RMCDSR) Load() CDSR   { return CDSR(rm.UM32.Load()) }
func (rm RMCDSR) Store(b CDSR) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) VDES() RMCDSR {
	return RMCDSR{mmio.UM32{&p.CDSR.U32, uint32(VDES)}}
}

func (p *LTDC_Periph) HDES() RMCDSR {
	return RMCDSR{mmio.UM32{&p.CDSR.U32, uint32(HDES)}}
}

func (p *LTDC_Periph) VSYNCS() RMCDSR {
	return RMCDSR{mmio.UM32{&p.CDSR.U32, uint32(VSYNCS)}}
}

func (p *LTDC_Periph) HSYNCS() RMCDSR {
	return RMCDSR{mmio.UM32{&p.CDSR.U32, uint32(HSYNCS)}}
}
