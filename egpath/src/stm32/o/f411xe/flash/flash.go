// Peripheral: FLASH_Periph  FLASH Registers.
// Instances:
//  FLASH  mmap.FLASH_R_BASE
// Registers:
//  0x00 32  ACR      Access control register.
//  0x04 32  KEYR     Key register.
//  0x08 32  OPTKEYR  Option key register.
//  0x0C 32  SR       Status register.
//  0x10 32  CR       Control register.
//  0x14 32  OPTCR[2] Option control registers.
// Import:
//  stm32/o/f411xe/mmap
package flash

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	LATENCY       ACR = 0x0F << 0 //+
	LATENCY_0WS   ACR = 0x00 << 0
	LATENCY_1WS   ACR = 0x01 << 0
	LATENCY_2WS   ACR = 0x02 << 0
	LATENCY_3WS   ACR = 0x03 << 0
	LATENCY_4WS   ACR = 0x04 << 0
	LATENCY_5WS   ACR = 0x05 << 0
	LATENCY_6WS   ACR = 0x06 << 0
	LATENCY_7WS   ACR = 0x07 << 0
	PRFTEN        ACR = 0x01 << 8  //+
	ICEN          ACR = 0x01 << 9  //+
	DCEN          ACR = 0x01 << 10 //+
	ICRST         ACR = 0x01 << 11 //+
	DCRST         ACR = 0x01 << 12 //+
	BYTE0_ADDRESS ACR = 0x10008F << 10
	BYTE2_ADDRESS ACR = 0x40023C03 << 0
)

const (
	LATENCYn = 0
	PRFTENn  = 8
	ICENn    = 9
	DCENn    = 10
	ICRSTn   = 11
	DCRSTn   = 12
)

const (
	EOP    SR = 0x01 << 0  //+
	SOP    SR = 0x01 << 1  //+
	WRPERR SR = 0x01 << 4  //+
	PGAERR SR = 0x01 << 5  //+
	PGPERR SR = 0x01 << 6  //+
	PGSERR SR = 0x01 << 7  //+
	RDERR  SR = 0x01 << 8  //+
	BSY    SR = 0x01 << 16 //+
)

const (
	EOPn    = 0
	SOPn    = 1
	WRPERRn = 4
	PGAERRn = 5
	PGPERRn = 6
	PGSERRn = 7
	RDERRn  = 8
	BSYn    = 16
)

const (
	PG    CR = 0x01 << 0  //+
	SER   CR = 0x01 << 1  //+
	MER   CR = 0x01 << 2  //+
	SNB   CR = 0x1F << 3  //+
	PSIZE CR = 0x03 << 8  //+
	STRT  CR = 0x01 << 16 //+
	EOPIE CR = 0x01 << 24 //+
	LOCK  CR = 0x01 << 31 //+
)

const (
	PGn    = 0
	SERn   = 1
	MERn   = 2
	SNBn   = 3
	PSIZEn = 8
	STRTn  = 16
	EOPIEn = 24
	LOCKn  = 31
)

const (
	OPTLOCK    OPTCR = 0x01 << 0   //+
	OPTSTRT    OPTCR = 0x01 << 1   //+
	BOR_LEV    OPTCR = 0x03 << 2   //+
	WDG_SW     OPTCR = 0x01 << 5   //+
	nRST_STOP  OPTCR = 0x01 << 6   //+
	nRST_STDBY OPTCR = 0x01 << 7   //+
	RDP        OPTCR = 0xFF << 8   //+
	nWRP       OPTCR = 0xFFF << 16 //+
	nWRP_10    OPTCR = 0x400 << 16
	nWRP_11    OPTCR = 0x800 << 16
)

const (
	OPTLOCKn    = 0
	OPTSTRTn    = 1
	BOR_LEVn    = 2
	WDG_SWn     = 5
	nRST_STOPn  = 6
	nRST_STDBYn = 7
	RDPn        = 8
	nWRPn       = 16
)
