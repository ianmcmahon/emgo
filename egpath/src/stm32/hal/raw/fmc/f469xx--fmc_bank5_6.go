// +build f469xx
// Peripheral: FMC_Bank5_6_Periph  Flexible Memory Controller Bank5_6.
// Instances:
//  FMC_Bank5_6  mmap.FMC_Bank5_6_R_BASE
// Registers:
//  0x00 32  SDCR[2] SDRAM Control registers .
//  0x08 32  SDTR[2] SDRAM Timing registers .
//  0x10 32  SDCMR   SDRAM Command Mode register.
//  0x14 32  SDRTR   SDRAM Refresh Timer register.
//  0x18 32  SDSR    SDRAM Status register.
// Import:
//  stm32/o/f469xx/mmap
package fmc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	NC     SDCR = 0x03 << 0  //+ NC[1:0] bits (Number of column bits).
	NR     SDCR = 0x03 << 2  //+ NR[1:0] bits (Number of row bits).
	SDMWID SDCR = 0x03 << 4  //+ NR[1:0] bits (Number of row bits).
	NB     SDCR = 0x01 << 6  //+ Number of internal bank.
	CAS    SDCR = 0x03 << 7  //+ CAS[1:0] bits (CAS latency).
	WP     SDCR = 0x01 << 9  //+ Write protection.
	SDCLK  SDCR = 0x03 << 10 //+ SDRAM clock configuration.
	RBURST SDCR = 0x01 << 12 //+ Read burst.
	RPIPE  SDCR = 0x03 << 13 //+ Write protection.
)

const (
	NCn     = 0
	NRn     = 2
	SDMWIDn = 4
	NBn     = 6
	CASn    = 7
	WPn     = 9
	SDCLKn  = 10
	RBURSTn = 12
	RPIPEn  = 13
)

const (
	TMRD SDTR = 0x0F << 0  //+ TMRD[3:0] bits (Load mode register to active).
	TXSR SDTR = 0x0F << 4  //+ TXSR[3:0] bits (Exit self refresh).
	TRAS SDTR = 0x0F << 8  //+ TRAS[3:0] bits (Self refresh time).
	TRC  SDTR = 0x0F << 12 //+ TRC[2:0] bits (Row cycle delay).
	TWR  SDTR = 0x0F << 16 //+ TRC[2:0] bits (Write recovery delay).
	TRP  SDTR = 0x0F << 20 //+ TRP[2:0] bits (Row precharge delay).
	TRCD SDTR = 0x0F << 24 //+ TRP[2:0] bits (Row to column delay).
)

const (
	TMRDn = 0
	TXSRn = 4
	TRASn = 8
	TRCn  = 12
	TWRn  = 16
	TRPn  = 20
	TRCDn = 24
)

const (
	MODE SDCMR = 0x07 << 0   //+ MODE[2:0] bits (Command mode).
	CTB2 SDCMR = 0x01 << 3   //+ Command target 2.
	CTB1 SDCMR = 0x01 << 4   //+ Command target 1.
	NRFS SDCMR = 0x0F << 5   //+ NRFS[3:0] bits (Number of auto-refresh).
	MRD  SDCMR = 0x1FFF << 9 //+ MRD[12:0] bits (Mode register definition).
)

const (
	MODEn = 0
	CTB2n = 3
	CTB1n = 4
	NRFSn = 5
	MRDn  = 9
)

const (
	CRE   SDRTR = 0x01 << 0   //+ Clear refresh error flag.
	COUNT SDRTR = 0x1FFF << 1 //+ COUNT[12:0] bits (Refresh timer count).
	REIE  SDRTR = 0x01 << 14  //+ RES interupt enable.
)

const (
	CREn   = 0
	COUNTn = 1
	REIEn  = 14
)

const (
	RE     SDSR = 0x01 << 0 //+ Refresh error flag.
	MODES1 SDSR = 0x03 << 1 //+ MODES1[1:0]bits (Status mode for bank 1).
	MODES2 SDSR = 0x03 << 3 //+ MODES2[1:0]bits (Status mode for bank 2).
	BUSY   SDSR = 0x01 << 5 //+ Busy status.
)

const (
	REn     = 0
	MODES1n = 1
	MODES2n = 3
	BUSYn   = 5
)
