// +build f469xx

// Peripheral: SPI_Periph  Serial Peripheral Interface.
// Instances:
//  I2S2ext  mmap.I2S2ext_BASE
//  SPI2     mmap.SPI2_BASE
//  SPI3     mmap.SPI3_BASE
//  I2S3ext  mmap.I2S3ext_BASE
//  SPI1     mmap.SPI1_BASE
//  SPI4     mmap.SPI4_BASE
//  SPI5     mmap.SPI5_BASE
//  SPI6     mmap.SPI6_BASE
// Registers:
//  0x00 32  CR1     Control register 1 (not used in I2S mode).
//  0x04 32  CR2     Control register 2.
//  0x08 32  SR      Status register.
//  0x0C 32  DR      Data register.
//  0x10 32  CRCPR   CRC polynomial register (not used in I2S mode).
//  0x14 32  RXCRCR  RX CRC register (not used in I2S mode).
//  0x18 32  TXCRCR  TX CRC register (not used in I2S mode).
//  0x1C 32  I2SCFGR SPI_I2S configuration register.
//  0x20 32  I2SPR   SPI_I2S prescaler register.
// Import:
//  stm32/o/f469xx/mmap
package spi

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	CPHA     CR1 = 0x01 << 0  //+ Clock Phase.
	CPOL     CR1 = 0x01 << 1  //+ Clock Polarity.
	MSTR     CR1 = 0x01 << 2  //+ Master Selection.
	BR       CR1 = 0x07 << 3  //+ BR[2:0] bits (Baud Rate Control).
	SPE      CR1 = 0x01 << 6  //+ SPI Enable.
	LSBFIRST CR1 = 0x01 << 7  //+ Frame Format.
	SSI      CR1 = 0x01 << 8  //+ Internal slave select.
	SSM      CR1 = 0x01 << 9  //+ Software slave management.
	RXONLY   CR1 = 0x01 << 10 //+ Receive only.
	DFF      CR1 = 0x01 << 11 //+ Data Frame Format.
	CRCNEXT  CR1 = 0x01 << 12 //+ Transmit CRC next.
	CRCEN    CR1 = 0x01 << 13 //+ Hardware CRC calculation enable.
	BIDIOE   CR1 = 0x01 << 14 //+ Output enable in bidirectional mode.
	BIDIMODE CR1 = 0x01 << 15 //+ Bidirectional data mode enable.
)

const (
	CPHAn     = 0
	CPOLn     = 1
	MSTRn     = 2
	BRn       = 3
	SPEn      = 6
	LSBFIRSTn = 7
	SSIn      = 8
	SSMn      = 9
	RXONLYn   = 10
	DFFn      = 11
	CRCNEXTn  = 12
	CRCENn    = 13
	BIDIOEn   = 14
	BIDIMODEn = 15
)

const (
	RXDMAEN CR2 = 0x01 << 0 //+ Rx Buffer DMA Enable.
	TXDMAEN CR2 = 0x01 << 1 //+ Tx Buffer DMA Enable.
	SSOE    CR2 = 0x01 << 2 //+ SS Output Enable.
	FRF     CR2 = 0x01 << 4 //+ Frame Format.
	ERRIE   CR2 = 0x01 << 5 //+ Error Interrupt Enable.
	RXNEIE  CR2 = 0x01 << 6 //+ RX buffer Not Empty Interrupt Enable.
	TXEIE   CR2 = 0x01 << 7 //+ Tx buffer Empty Interrupt Enable.
)

const (
	RXDMAENn = 0
	TXDMAENn = 1
	SSOEn    = 2
	FRFn     = 4
	ERRIEn   = 5
	RXNEIEn  = 6
	TXEIEn   = 7
)

const (
	RXNE   SR = 0x01 << 0 //+ Receive buffer Not Empty.
	TXE    SR = 0x01 << 1 //+ Transmit buffer Empty.
	CHSIDE SR = 0x01 << 2 //+ Channel side.
	UDR    SR = 0x01 << 3 //+ Underrun flag.
	CRCERR SR = 0x01 << 4 //+ CRC Error flag.
	MODF   SR = 0x01 << 5 //+ Mode fault.
	OVR    SR = 0x01 << 6 //+ Overrun flag.
	BSY    SR = 0x01 << 7 //+ Busy flag.
	FRE    SR = 0x01 << 8 //+ Frame format error flag.
)

const (
	RXNEn   = 0
	TXEn    = 1
	CHSIDEn = 2
	UDRn    = 3
	CRCERRn = 4
	MODFn   = 5
	OVRn    = 6
	BSYn    = 7
	FREn    = 8
)

const (
	CRCPOLY CRCPR = 0xFFFF << 0 //+ CRC polynomial register.
)

const (
	CRCPOLYn = 0
)

const (
	RXCRC RXCRCR = 0xFFFF << 0 //+ Rx CRC Register.
)

const (
	RXCRCn = 0
)

const (
	TXCRC TXCRCR = 0xFFFF << 0 //+ Tx CRC Register.
)

const (
	TXCRCn = 0
)

const (
	CHLEN   I2SCFGR = 0x01 << 0  //+ Channel length (number of bits per audio channel).
	DATLEN  I2SCFGR = 0x03 << 1  //+ DATLEN[1:0] bits (Data length to be transferred).
	CKPOL   I2SCFGR = 0x01 << 3  //+ steady state clock polarity.
	I2SSTD  I2SCFGR = 0x03 << 4  //+ I2SSTD[1:0] bits (I2S standard selection).
	PCMSYNC I2SCFGR = 0x01 << 7  //+ PCM frame synchronization.
	I2SCFG  I2SCFGR = 0x03 << 8  //+ I2SCFG[1:0] bits (I2S configuration mode).
	I2SE    I2SCFGR = 0x01 << 10 //+ I2S Enable.
	I2SMOD  I2SCFGR = 0x01 << 11 //+ I2S mode selection.
	ASTRTEN I2SCFGR = 0x01 << 12 //+ Asynchronous start enable.
)

const (
	CHLENn   = 0
	DATLENn  = 1
	CKPOLn   = 3
	I2SSTDn  = 4
	PCMSYNCn = 7
	I2SCFGn  = 8
	I2SEn    = 10
	I2SMODn  = 11
	ASTRTENn = 12
)

const (
	I2SDIV I2SPR = 0xFF << 0 //+ I2S Linear prescaler.
	ODD    I2SPR = 0x01 << 8 //+ Odd factor for the prescaler.
	MCKOE  I2SPR = 0x01 << 9 //+ Master Clock Output Enable.
)

const (
	I2SDIVn = 0
	ODDn    = 8
	MCKOEn  = 9
)
