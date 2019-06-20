// Peripheral: QUADSPI_Periph  QUAD Serial Peripheral Interface.
// Instances:
//  QUADSPI  mmap.QSPI_R_BASE
// Registers:
//  0x00 32  CR    Control register.
//  0x04 32  DCR   Device Configuration register.
//  0x08 32  SR    Status register.
//  0x0C 32  FCR   Flag Clear register.
//  0x10 32  DLR   Data Length register.
//  0x14 32  CCR   Communication Configuration register.
//  0x18 32  AR    Address register.
//  0x1C 32  ABR   Alternate Bytes register.
//  0x20 32  DR    Data register.
//  0x24 32  PSMKR Polling Status Mask register.
//  0x28 32  PSMAR Polling Status Match register.
//  0x2C 32  PIR   Polling Interval register.
//  0x30 32  LPTR  Low Power Timeout register.
// Import:
//  stm32/o/f746xx/mmap
package quadspi

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	EN        CR = 0x01 << 0  //+ Enable.
	ABORT     CR = 0x01 << 1  //+ Abort request.
	DMAEN     CR = 0x01 << 2  //+ DMA Enable.
	TCEN      CR = 0x01 << 3  //+ Timeout Counter Enable.
	SSHIFT    CR = 0x01 << 4  //+ Sample Shift.
	DFM       CR = 0x01 << 6  //+ Dual Flash Mode.
	FSEL      CR = 0x01 << 7  //+ Flash Select.
	FTHRES    CR = 0x1F << 8  //+ FTHRES[4:0] FIFO Level.
	TEIE      CR = 0x01 << 16 //+ Transfer Error Interrupt Enable.
	TCIE      CR = 0x01 << 17 //+ Transfer Complete Interrupt Enable.
	FTIE      CR = 0x01 << 18 //+ FIFO Threshold Interrupt Enable.
	SMIE      CR = 0x01 << 19 //+ Status Match Interrupt Enable.
	TOIE      CR = 0x01 << 20 //+ TimeOut Interrupt Enable.
	APMS      CR = 0x01 << 22 //+ Bit 1.
	PMM       CR = 0x01 << 23 //+ Polling Match Mode.
	PRESCALER CR = 0xFF << 24 //+ PRESCALER[7:0] Clock prescaler.
)

const (
	ENn        = 0
	ABORTn     = 1
	DMAENn     = 2
	TCENn      = 3
	SSHIFTn    = 4
	DFMn       = 6
	FSELn      = 7
	FTHRESn    = 8
	TEIEn      = 16
	TCIEn      = 17
	FTIEn      = 18
	SMIEn      = 19
	TOIEn      = 20
	APMSn      = 22
	PMMn       = 23
	PRESCALERn = 24
)

const (
	CKMODE DCR = 0x01 << 0  //+ Mode 0 / Mode 3.
	CSHT   DCR = 0x07 << 8  //+ CSHT[2:0]: ChipSelect High Time.
	FSIZE  DCR = 0x1F << 16 //+ FSIZE[4:0]: Flash Size.
)

const (
	CKMODEn = 0
	CSHTn   = 8
	FSIZEn  = 16
)

const (
	TEF    SR = 0x01 << 0 //+ Transfer Error Flag.
	TCF    SR = 0x01 << 1 //+ Transfer Complete Flag.
	FTF    SR = 0x01 << 2 //+ FIFO Threshlod Flag.
	SMF    SR = 0x01 << 3 //+ Status Match Flag.
	TOF    SR = 0x01 << 4 //+ Timeout Flag.
	BUSY   SR = 0x01 << 5 //+ Busy.
	FLEVEL SR = 0x1F << 8 //+ FIFO Threshlod Flag.
)

const (
	TEFn    = 0
	TCFn    = 1
	FTFn    = 2
	SMFn    = 3
	TOFn    = 4
	BUSYn   = 5
	FLEVELn = 8
)

const (
	CTEF FCR = 0x01 << 0 //+ Clear Transfer Error Flag.
	CTCF FCR = 0x01 << 1 //+ Clear Transfer Complete Flag.
	CSMF FCR = 0x01 << 3 //+ Clear Status Match Flag.
	CTOF FCR = 0x01 << 4 //+ Clear Timeout Flag.
)

const (
	CTEFn = 0
	CTCFn = 1
	CSMFn = 3
	CTOFn = 4
)

const (
	DL DLR = 0xFFFFFFFF << 0 //+ DL[31:0]: Data Length.
)

const (
	DLn = 0
)

const (
	INSTRUCTION CCR = 0xFF << 0  //+ INSTRUCTION[7:0]: Instruction.
	IMODE       CCR = 0x03 << 8  //+ IMODE[1:0]: Instruction Mode.
	ADMODE      CCR = 0x03 << 10 //+ ADMODE[1:0]: Address Mode.
	ADSIZE      CCR = 0x03 << 12 //+ ADSIZE[1:0]: Address Size.
	ABMODE      CCR = 0x03 << 14 //+ ABMODE[1:0]: Alternate Bytes Mode.
	ABSIZE      CCR = 0x03 << 16 //+ ABSIZE[1:0]: Instruction Mode.
	DCYC        CCR = 0x1F << 18 //+ DCYC[4:0]: Dummy Cycles.
	DMODE       CCR = 0x03 << 24 //+ DMODE[1:0]: Data Mode.
	FMODE       CCR = 0x03 << 26 //+ FMODE[1:0]: Functional Mode.
	SIOO        CCR = 0x01 << 28 //+ SIOO: Send Instruction Only Once Mode.
	DHHC        CCR = 0x01 << 30 //+ DHHC: Delay Half Hclk Cycle.
	DDRM        CCR = 0x01 << 31 //+ DDRM: Double Data Rate Mode.
)

const (
	INSTRUCTIONn = 0
	IMODEn       = 8
	ADMODEn      = 10
	ADSIZEn      = 12
	ABMODEn      = 14
	ABSIZEn      = 16
	DCYCn        = 18
	DMODEn       = 24
	FMODEn       = 26
	SIOOn        = 28
	DHHCn        = 30
	DDRMn        = 31
)

const (
	ADDRESS AR = 0xFFFFFFFF << 0 //+ ADDRESS[31:0]: Address.
)

const (
	ADDRESSn = 0
)

const (
	ALTERNATE ABR = 0xFFFFFFFF << 0 //+ ALTERNATE[31:0]: Alternate Bytes.
)

const (
	ALTERNATEn = 0
)

const (
	DATA DR = 0xFFFFFFFF << 0 //+ DATA[31:0]: Data.
)

const (
	DATAn = 0
)

const (
	MASK PSMKR = 0xFFFFFFFF << 0 //+ MASK[31:0]: Status Mask.
)

const (
	MASKn = 0
)

const (
	MATCH PSMAR = 0xFFFFFFFF << 0 //+ MATCH[31:0]: Status Match.
)

const (
	MATCHn = 0
)

const (
	INTERVAL PIR = 0xFFFF << 0 //+ INTERVAL[15:0]: Polling Interval.
)

const (
	INTERVALn = 0
)

const (
	TIMEOUT LPTR = 0xFFFF << 0 //+ TIMEOUT[15:0]: Timeout period.
)

const (
	TIMEOUTn = 0
)
