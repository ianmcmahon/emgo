// Peripheral: DMA_Stream_Periph  DMA Controller.
// Instances:
//  DMA1_Stream0  mmap.DMA1_Stream0_BASE
//  DMA1_Stream1  mmap.DMA1_Stream1_BASE
//  DMA1_Stream2  mmap.DMA1_Stream2_BASE
//  DMA1_Stream3  mmap.DMA1_Stream3_BASE
//  DMA1_Stream4  mmap.DMA1_Stream4_BASE
//  DMA1_Stream5  mmap.DMA1_Stream5_BASE
//  DMA1_Stream6  mmap.DMA1_Stream6_BASE
//  DMA1_Stream7  mmap.DMA1_Stream7_BASE
//  DMA2_Stream0  mmap.DMA2_Stream0_BASE
//  DMA2_Stream1  mmap.DMA2_Stream1_BASE
//  DMA2_Stream2  mmap.DMA2_Stream2_BASE
//  DMA2_Stream3  mmap.DMA2_Stream3_BASE
//  DMA2_Stream4  mmap.DMA2_Stream4_BASE
//  DMA2_Stream5  mmap.DMA2_Stream5_BASE
//  DMA2_Stream6  mmap.DMA2_Stream6_BASE
//  DMA2_Stream7  mmap.DMA2_Stream7_BASE
// Registers:
//  0x00 32  CR   DMA stream x configuration register.
//  0x04 32  NDTR DMA stream x number of data register.
//  0x08 32  PAR  DMA stream x peripheral address register.
//  0x0C 32  M0AR DMA stream x memory 0 address register.
//  0x10 32  M1AR DMA stream x memory 1 address register.
//  0x14 32  FCR  DMA stream x FIFO control register.
// Import:
//  stm32/o/f40_41xxx/mmap
package dma

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	CHSEL    CR_Bits = 0x07 << 25 //+
	CHSEL_0  CR_Bits = 0x01 << 25
	CHSEL_1  CR_Bits = 0x02 << 25
	CHSEL_2  CR_Bits = 0x04 << 25
	MBURST   CR_Bits = 0x03 << 23 //+
	MBURST_0 CR_Bits = 0x01 << 23
	MBURST_1 CR_Bits = 0x02 << 23
	PBURST   CR_Bits = 0x03 << 21 //+
	PBURST_0 CR_Bits = 0x01 << 21
	PBURST_1 CR_Bits = 0x02 << 21
	ACK      CR_Bits = 0x01 << 20 //+
	CT       CR_Bits = 0x01 << 19 //+
	DBM      CR_Bits = 0x01 << 18 //+
	PL       CR_Bits = 0x03 << 16 //+
	PL_0     CR_Bits = 0x01 << 16
	PL_1     CR_Bits = 0x02 << 16
	PINCOS   CR_Bits = 0x01 << 15 //+
	MSIZE    CR_Bits = 0x03 << 13 //+
	MSIZE_0  CR_Bits = 0x01 << 13
	MSIZE_1  CR_Bits = 0x02 << 13
	PSIZE    CR_Bits = 0x03 << 11 //+
	PSIZE_0  CR_Bits = 0x01 << 11
	PSIZE_1  CR_Bits = 0x02 << 11
	MINC     CR_Bits = 0x01 << 10 //+
	PINC     CR_Bits = 0x01 << 9  //+
	CIRC     CR_Bits = 0x01 << 8  //+
	DIR      CR_Bits = 0x03 << 6  //+
	DIR_0    CR_Bits = 0x01 << 6
	DIR_1    CR_Bits = 0x02 << 6
	PFCTRL   CR_Bits = 0x01 << 5 //+
	TCIE     CR_Bits = 0x01 << 4 //+
	HTIE     CR_Bits = 0x01 << 3 //+
	TEIE     CR_Bits = 0x01 << 2 //+
	DMEIE    CR_Bits = 0x01 << 1 //+
	EN       CR_Bits = 0x01 << 0 //+
)

const (
	CHSELn  = 25
	MBURSTn = 23
	PBURSTn = 21
	ACKn    = 20
	CTn     = 19
	DBMn    = 18
	PLn     = 16
	PINCOSn = 15
	MSIZEn  = 13
	PSIZEn  = 11
	MINCn   = 10
	PINCn   = 9
	CIRCn   = 8
	DIRn    = 6
	PFCTRLn = 5
	TCIEn   = 4
	HTIEn   = 3
	TEIEn   = 2
	DMEIEn  = 1
	ENn     = 0
)

const (
	FEIE  FCR_Bits = 0x01 << 7 //+
	FS    FCR_Bits = 0x07 << 3 //+
	FS_0  FCR_Bits = 0x01 << 3
	FS_1  FCR_Bits = 0x02 << 3
	FS_2  FCR_Bits = 0x04 << 3
	DMDIS FCR_Bits = 0x01 << 2 //+
	FTH   FCR_Bits = 0x03 << 0 //+
	FTH_0 FCR_Bits = 0x01 << 0
	FTH_1 FCR_Bits = 0x02 << 0
)

const (
	FEIEn  = 7
	FSn    = 3
	DMDISn = 2
	FTHn   = 0
)
