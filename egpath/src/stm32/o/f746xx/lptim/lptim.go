// Peripheral: LPTIM_Periph  LPTIMIMER.
// Instances:
//  LPTIM1  mmap.LPTIM1_BASE
// Registers:
//  0x00 32  ISR  Interrupt and Status register.
//  0x04 32  ICR  Interrupt Clear register.
//  0x08 32  IER  Interrupt Enable register.
//  0x0C 32  CFGR Configuration register.
//  0x10 32  CR   Control register.
//  0x14 32  CMP  Compare register.
//  0x18 32  ARR  Autoreload register.
//  0x1C 32  CNT  Counter register.
// Import:
//  stm32/o/f746xx/mmap
package lptim

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	CMPM    ISR = 0x01 << 0 //+ Compare match.
	ARRM    ISR = 0x01 << 1 //+ Autoreload match.
	EXTTRIG ISR = 0x01 << 2 //+ External trigger edge event.
	CMPOK   ISR = 0x01 << 3 //+ Compare register update OK.
	ARROK   ISR = 0x01 << 4 //+ Autoreload register update OK.
	UP      ISR = 0x01 << 5 //+ Counter direction change down to up.
	DOWN    ISR = 0x01 << 6 //+ Counter direction change up to down.
)

const (
	CMPMn    = 0
	ARRMn    = 1
	EXTTRIGn = 2
	CMPOKn   = 3
	ARROKn   = 4
	UPn      = 5
	DOWNn    = 6
)

const (
	CMPMCF    ICR = 0x01 << 0 //+ Compare match Clear Flag.
	ARRMCF    ICR = 0x01 << 1 //+ Autoreload match Clear Flag.
	EXTTRIGCF ICR = 0x01 << 2 //+ External trigger edge event Clear Flag.
	CMPOKCF   ICR = 0x01 << 3 //+ Compare register update OK Clear Flag.
	ARROKCF   ICR = 0x01 << 4 //+ Autoreload register update OK Clear Flag.
	UPCF      ICR = 0x01 << 5 //+ Counter direction change down to up Clear Flag.
	DOWNCF    ICR = 0x01 << 6 //+ Counter direction change up to down Clear Flag.
)

const (
	CMPMCFn    = 0
	ARRMCFn    = 1
	EXTTRIGCFn = 2
	CMPOKCFn   = 3
	ARROKCFn   = 4
	UPCFn      = 5
	DOWNCFn    = 6
)

const (
	CMPMIE    IER = 0x01 << 0 //+ Compare match Interrupt Enable.
	ARRMIE    IER = 0x01 << 1 //+ Autoreload match Interrupt Enable.
	EXTTRIGIE IER = 0x01 << 2 //+ External trigger edge event Interrupt Enable.
	CMPOKIE   IER = 0x01 << 3 //+ Compare register update OK Interrupt Enable.
	ARROKIE   IER = 0x01 << 4 //+ Autoreload register update OK Interrupt Enable.
	UPIE      IER = 0x01 << 5 //+ Counter direction change down to up Interrupt Enable.
	DOWNIE    IER = 0x01 << 6 //+ Counter direction change up to down Interrupt Enable.
)

const (
	CMPMIEn    = 0
	ARRMIEn    = 1
	EXTTRIGIEn = 2
	CMPOKIEn   = 3
	ARROKIEn   = 4
	UPIEn      = 5
	DOWNIEn    = 6
)

const (
	CKSEL     CFGR = 0x01 << 0  //+ Clock selector.
	CKPOL     CFGR = 0x03 << 1  //+ CKPOL[1:0] bits (Clock polarity).
	CKFLT     CFGR = 0x03 << 3  //+ CKFLT[1:0] bits (Configurable digital filter for external clock).
	TRGFLT    CFGR = 0x03 << 6  //+ TRGFLT[1:0] bits (Configurable digital filter for trigger).
	PRESC     CFGR = 0x07 << 9  //+ PRESC[2:0] bits (Clock prescaler).
	TRIGSEL   CFGR = 0x07 << 13 //+ TRIGSEL[2:0]] bits (Trigger selector).
	TRIGEN    CFGR = 0x03 << 17 //+ TRIGEN[1:0] bits (Trigger enable and polarity).
	TIMOUT    CFGR = 0x01 << 19 //+ Timout enable.
	WAVE      CFGR = 0x01 << 20 //+ Waveform shape.
	WAVPOL    CFGR = 0x01 << 21 //+ Waveform shape polarity.
	PRELOAD   CFGR = 0x01 << 22 //+ Reg update mode.
	COUNTMODE CFGR = 0x01 << 23 //+ Counter mode enable.
	ENC       CFGR = 0x01 << 24 //+ Encoder mode enable.
)

const (
	CKSELn     = 0
	CKPOLn     = 1
	CKFLTn     = 3
	TRGFLTn    = 6
	PRESCn     = 9
	TRIGSELn   = 13
	TRIGENn    = 17
	TIMOUTn    = 19
	WAVEn      = 20
	WAVPOLn    = 21
	PRELOADn   = 22
	COUNTMODEn = 23
	ENCn       = 24
)

const (
	ENABLE  CR = 0x01 << 0 //+ LPTIMer enable.
	SNGSTRT CR = 0x01 << 1 //+ Timer start in single mode.
	CNTSTRT CR = 0x01 << 2 //+ Timer start in continuous mode.
)

const (
	ENABLEn  = 0
	SNGSTRTn = 1
	CNTSTRTn = 2
)
