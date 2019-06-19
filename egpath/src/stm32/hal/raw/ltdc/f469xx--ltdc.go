// +build f469xx
// Peripheral: LTDC_Periph  LCD-TFT Display Controller.
// Instances:
//  LTDC  mmap.LTDC_BASE
// Registers:
//  0x08 32  SSCR  Synchronization Size Configuration Register.
//  0x0C 32  BPCR  Back Porch Configuration Register.
//  0x10 32  AWCR  Active Width Configuration Register.
//  0x14 32  TWCR  Total Width Configuration Register.
//  0x18 32  GCR   Global Control Register.
//  0x24 32  SRCR  Shadow Reload Configuration Register.
//  0x2C 32  BCCR  Background Color Configuration Register.
//  0x34 32  IER   Interrupt Enable Register.
//  0x38 32  ISR   Interrupt Status Register.
//  0x3C 32  ICR   Interrupt Clear Register.
//  0x40 32  LIPCR Line Interrupt Position Configuration Register.
//  0x44 32  CPSR  Current Position Status Register.
//  0x48 32  CDSR  Current Display Status Register.
// Import:
//  stm32/o/f469xx/mmap
package ltdc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	VSH SSCR = 0x7FF << 0  //+ Vertical Synchronization Height.
	HSW SSCR = 0xFFF << 16 //+ Horizontal Synchronization Width.
)

const (
	VSHn = 0
	HSWn = 16
)

const (
	AVBP BPCR = 0x7FF << 0  //+ Accumulated Vertical Back Porch.
	AHBP BPCR = 0xFFF << 16 //+ Accumulated Horizontal Back Porch.
)

const (
	AVBPn = 0
	AHBPn = 16
)

const (
	AAH AWCR = 0x7FF << 0  //+ Accumulated Active heigh.
	AAW AWCR = 0xFFF << 16 //+ Accumulated Active Width.
)

const (
	AAHn = 0
	AAWn = 16
)

const (
	TOTALH TWCR = 0x7FF << 0  //+ Total Heigh.
	TOTALW TWCR = 0xFFF << 16 //+ Total Width.
)

const (
	TOTALHn = 0
	TOTALWn = 16
)

const (
	LTDCEN GCR = 0x01 << 0  //+ LCD-TFT controller enable bit.
	DBW    GCR = 0x07 << 4  //+ Dither Blue Width.
	DGW    GCR = 0x07 << 8  //+ Dither Green Width.
	DRW    GCR = 0x07 << 12 //+ Dither Red Width.
	DEN    GCR = 0x01 << 16 //+ Dither Enable.
	PCPOL  GCR = 0x01 << 28 //+ Pixel Clock Polarity.
	DEPOL  GCR = 0x01 << 29 //+ Data Enable Polarity.
	VSPOL  GCR = 0x01 << 30 //+ Vertical Synchronization Polarity.
	HSPOL  GCR = 0x01 << 31 //+ Horizontal Synchronization Polarity.
)

const (
	LTDCENn = 0
	DBWn    = 4
	DGWn    = 8
	DRWn    = 12
	DENn    = 16
	PCPOLn  = 28
	DEPOLn  = 29
	VSPOLn  = 30
	HSPOLn  = 31
)

const (
	IMR SRCR = 0x01 << 0 //+ Immediate Reload.
	VBR SRCR = 0x01 << 1 //+ Vertical Blanking Reload.
)

const (
	IMRn = 0
	VBRn = 1
)

const (
	BCBLUE  BCCR = 0xFF << 0  //+ Background Blue value.
	BCGREEN BCCR = 0xFF << 8  //+ Background Green value.
	BCRED   BCCR = 0xFF << 16 //+ Background Red value.
)

const (
	BCBLUEn  = 0
	BCGREENn = 8
	BCREDn   = 16
)

const (
	LIE    IER = 0x01 << 0 //+ Line Interrupt Enable.
	FUIE   IER = 0x01 << 1 //+ FIFO Underrun Interrupt Enable.
	TERRIE IER = 0x01 << 2 //+ Transfer Error Interrupt Enable.
	RRIE   IER = 0x01 << 3 //+ Register Reload interrupt enable.
)

const (
	LIEn    = 0
	FUIEn   = 1
	TERRIEn = 2
	RRIEn   = 3
)

const (
	LIF    ISR = 0x01 << 0 //+ Line Interrupt Flag.
	FUIF   ISR = 0x01 << 1 //+ FIFO Underrun Interrupt Flag.
	TERRIF ISR = 0x01 << 2 //+ Transfer Error Interrupt Flag.
	RRIF   ISR = 0x01 << 3 //+ Register Reload interrupt Flag.
)

const (
	LIFn    = 0
	FUIFn   = 1
	TERRIFn = 2
	RRIFn   = 3
)

const (
	CLIF    ICR = 0x01 << 0 //+ Clears the Line Interrupt Flag.
	CFUIF   ICR = 0x01 << 1 //+ Clears the FIFO Underrun Interrupt Flag.
	CTERRIF ICR = 0x01 << 2 //+ Clears the Transfer Error Interrupt Flag.
	CRRIF   ICR = 0x01 << 3 //+ Clears Register Reload interrupt Flag.
)

const (
	CLIFn    = 0
	CFUIFn   = 1
	CTERRIFn = 2
	CRRIFn   = 3
)

const (
	LIPOS LIPCR = 0x7FF << 0 //+ Line Interrupt Position.
)

const (
	LIPOSn = 0
)

const (
	CYPOS CPSR = 0xFFFF << 0  //+ Current Y Position.
	CXPOS CPSR = 0xFFFF << 16 //+ Current X Position.
)

const (
	CYPOSn = 0
	CXPOSn = 16
)

const (
	VDES   CDSR = 0x01 << 0 //+ Vertical Data Enable Status.
	HDES   CDSR = 0x01 << 1 //+ Horizontal Data Enable Status.
	VSYNCS CDSR = 0x01 << 2 //+ Vertical Synchronization Status.
	HSYNCS CDSR = 0x01 << 3 //+ Horizontal Synchronization Status.
)

const (
	VDESn   = 0
	HDESn   = 1
	VSYNCSn = 2
	HSYNCSn = 3
)
