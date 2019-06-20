// Peripheral: I2C_Periph  Inter-integrated Circuit Interface.
// Instances:
//  I2C1  mmap.I2C1_BASE
//  I2C2  mmap.I2C2_BASE
//  I2C3  mmap.I2C3_BASE
// Registers:
//  0x00 32  CR1   Control register 1.
//  0x04 32  CR2   Control register 2.
//  0x08 32  OAR1  Own address register 1.
//  0x0C 32  OAR2  Own address register 2.
//  0x10 32  DR    Data register.
//  0x14 32  SR1   Status register 1.
//  0x18 32  SR2   Status register 2.
//  0x1C 32  CCR   Clock control register.
//  0x20 32  TRISE TRISE register.
//  0x24 32  FLTR  FLTR register.
// Import:
//  stm32/o/f411xe/mmap
package i2c

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	PE        CR1 = 0x01 << 0  //+ Peripheral Enable.
	SMBUS     CR1 = 0x01 << 1  //+ SMBus Mode.
	SMBTYPE   CR1 = 0x01 << 3  //+ SMBus Type.
	ENARP     CR1 = 0x01 << 4  //+ ARP Enable.
	ENPEC     CR1 = 0x01 << 5  //+ PEC Enable.
	ENGC      CR1 = 0x01 << 6  //+ General Call Enable.
	NOSTRETCH CR1 = 0x01 << 7  //+ Clock Stretching Disable (Slave mode).
	START     CR1 = 0x01 << 8  //+ Start Generation.
	STOP      CR1 = 0x01 << 9  //+ Stop Generation.
	ACK       CR1 = 0x01 << 10 //+ Acknowledge Enable.
	POS       CR1 = 0x01 << 11 //+ Acknowledge/PEC Position (for data reception).
	PEC       CR1 = 0x01 << 12 //+ Packet Error Checking.
	ALERT     CR1 = 0x01 << 13 //+ SMBus Alert.
	SWRST     CR1 = 0x01 << 15 //+ Software Reset.
)

const (
	PEn        = 0
	SMBUSn     = 1
	SMBTYPEn   = 3
	ENARPn     = 4
	ENPECn     = 5
	ENGCn      = 6
	NOSTRETCHn = 7
	STARTn     = 8
	STOPn      = 9
	ACKn       = 10
	POSn       = 11
	PECn       = 12
	ALERTn     = 13
	SWRSTn     = 15
)

const (
	FREQ    CR2 = 0x3F << 0  //+ FREQ[5:0] bits (Peripheral Clock Frequency).
	ITERREN CR2 = 0x01 << 8  //+ Error Interrupt Enable.
	ITEVTEN CR2 = 0x01 << 9  //+ Event Interrupt Enable.
	ITBUFEN CR2 = 0x01 << 10 //+ Buffer Interrupt Enable.
	DMAEN   CR2 = 0x01 << 11 //+ DMA Requests Enable.
	LAST    CR2 = 0x01 << 12 //+ DMA Last Transfer.
)

const (
	FREQn    = 0
	ITERRENn = 8
	ITEVTENn = 9
	ITBUFENn = 10
	DMAENn   = 11
	LASTn    = 12
)

const (
	ADD0    OAR1 = 0x01 << 0  //+ Bit 0.
	ADD1    OAR1 = 0x01 << 1  //+ Bit 1.
	ADD2    OAR1 = 0x01 << 2  //+ Bit 2.
	ADD3    OAR1 = 0x01 << 3  //+ Bit 3.
	ADD4    OAR1 = 0x01 << 4  //+ Bit 4.
	ADD5    OAR1 = 0x01 << 5  //+ Bit 5.
	ADD6    OAR1 = 0x01 << 6  //+ Bit 6.
	ADD7    OAR1 = 0x01 << 7  //+ Bit 7.
	ADD8    OAR1 = 0x01 << 8  //+ Bit 8.
	ADD9    OAR1 = 0x01 << 9  //+ Bit 9.
	ADDMODE OAR1 = 0x01 << 15 //+ Addressing Mode (Slave mode).
)

const (
	ADD0n    = 0
	ADD1n    = 1
	ADD2n    = 2
	ADD3n    = 3
	ADD4n    = 4
	ADD5n    = 5
	ADD6n    = 6
	ADD7n    = 7
	ADD8n    = 8
	ADD9n    = 9
	ADDMODEn = 15
)

const (
	ENDUAL    OAR2 = 0x01 << 0 //+ Dual addressing mode enable.
	SECADD1_7 OAR2 = 0x7F << 1 //+ Interface address.
)

const (
	ENDUALn    = 0
	SECADD1_7n = 1
)

const (
	SB       SR1 = 0x01 << 0  //+ Start Bit (Master mode).
	ADDR     SR1 = 0x01 << 1  //+ Address sent (master mode)/matched (slave mode).
	BTF      SR1 = 0x01 << 2  //+ Byte Transfer Finished.
	ADD10    SR1 = 0x01 << 3  //+ 10-bit header sent (Master mode).
	STOPF    SR1 = 0x01 << 4  //+ Stop detection (Slave mode).
	RXNE     SR1 = 0x01 << 6  //+ Data Register not Empty (receivers).
	TXE      SR1 = 0x01 << 7  //+ Data Register Empty (transmitters).
	BERR     SR1 = 0x01 << 8  //+ Bus Error.
	ARLO     SR1 = 0x01 << 9  //+ Arbitration Lost (master mode).
	AF       SR1 = 0x01 << 10 //+ Acknowledge Failure.
	OVR      SR1 = 0x01 << 11 //+ Overrun/Underrun.
	PECERR   SR1 = 0x01 << 12 //+ PEC Error in reception.
	TIMEOUT  SR1 = 0x01 << 14 //+ Timeout or Tlow Error.
	SMBALERT SR1 = 0x01 << 15 //+ SMBus Alert.
)

const (
	SBn       = 0
	ADDRn     = 1
	BTFn      = 2
	ADD10n    = 3
	STOPFn    = 4
	RXNEn     = 6
	TXEn      = 7
	BERRn     = 8
	ARLOn     = 9
	AFn       = 10
	OVRn      = 11
	PECERRn   = 12
	TIMEOUTn  = 14
	SMBALERTn = 15
)

const (
	MSL        SR2 = 0x01 << 0 //+ Master/Slave.
	BUSY       SR2 = 0x01 << 1 //+ Bus Busy.
	TRA        SR2 = 0x01 << 2 //+ Transmitter/Receiver.
	GENCALL    SR2 = 0x01 << 4 //+ General Call Address (Slave mode).
	SMBDEFAULT SR2 = 0x01 << 5 //+ SMBus Device Default Address (Slave mode).
	SMBHOST    SR2 = 0x01 << 6 //+ SMBus Host Header (Slave mode).
	DUALF      SR2 = 0x01 << 7 //+ Dual Flag (Slave mode).
	PECVAL     SR2 = 0xFF << 8 //+ Packet Error Checking Register.
)

const (
	MSLn        = 0
	BUSYn       = 1
	TRAn        = 2
	GENCALLn    = 4
	SMBDEFAULTn = 5
	SMBHOSTn    = 6
	DUALFn      = 7
	PECVALn     = 8
)

const (
	CCRVAL CCR = 0xFFF << 0 //+ Clock Control Register in Fast/Standard mode (Master mode).
	DUTY   CCR = 0x01 << 14 //+ Fast Mode Duty Cycle.
	FS     CCR = 0x01 << 15 //+ I2C Master Mode Selection.
)

const (
	CCRVALn = 0
	DUTYn   = 14
	FSn     = 15
)

const (
	DNF   FLTR = 0x0F << 0 //+ Digital Noise Filter.
	ANOFF FLTR = 0x01 << 4 //+ Analog Noise Filter OFF.
)

const (
	DNFn   = 0
	ANOFFn = 4
)
