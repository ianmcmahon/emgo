// +build f469xx
// Peripheral: FMC_Bank3_Periph  Flexible Memory Controller Bank3.
// Instances:
//  FMC_Bank3  mmap.FMC_Bank3_R_BASE
// Registers:
//  0x00 32  PCR  NAND Flash control register.
//  0x04 32  SR   NAND Flash FIFO status and interrupt register.
//  0x08 32  PMEM NAND Flash Common memory space timing register.
//  0x0C 32  PATT NAND Flash Attribute memory space timing register.
//  0x14 32  ECCR NAND Flash ECC result registers.
// Import:
//  stm32/o/f469xx/mmap
package fmc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	PWAITEN PCR = 0x01 << 1  //+ Wait feature enable bit.
	PBKEN   PCR = 0x01 << 2  //+ PC Card/NAND Flash memory bank enable bit.
	PTYP    PCR = 0x01 << 3  //+ Memory type.
	PWID    PCR = 0x03 << 4  //+ PWID[1:0] bits (NAND Flash databus width).
	ECCEN   PCR = 0x01 << 6  //+ ECC computation logic enable bit.
	TCLR    PCR = 0x0F << 9  //+ TCLR[3:0] bits (CLE to RE delay).
	TAR     PCR = 0x0F << 13 //+ TAR[3:0] bits (ALE to RE delay).
	ECCPS   PCR = 0x07 << 17 //+ ECCPS[1:0] bits (ECC page size).
)

const (
	PWAITENn = 1
	PBKENn   = 2
	PTYPn    = 3
	PWIDn    = 4
	ECCENn   = 6
	TCLRn    = 9
	TARn     = 13
	ECCPSn   = 17
)

const (
	IRS   SR = 0x01 << 0 //+ Interrupt Rising Edge status.
	ILS   SR = 0x01 << 1 //+ Interrupt Level status.
	IFS   SR = 0x01 << 2 //+ Interrupt Falling Edge status.
	IREN  SR = 0x01 << 3 //+ Interrupt Rising Edge detection Enable bit.
	ILEN  SR = 0x01 << 4 //+ Interrupt Level detection Enable bit.
	IFEN  SR = 0x01 << 5 //+ Interrupt Falling Edge detection Enable bit.
	FEMPT SR = 0x01 << 6 //+ FIFO empty.
)

const (
	IRSn   = 0
	ILSn   = 1
	IFSn   = 2
	IRENn  = 3
	ILENn  = 4
	IFENn  = 5
	FEMPTn = 6
)

const (
	MEMSET2  PMEM = 0xFF << 0  //+ MEMSET2[7:0] bits (Common memory 2 setup time).
	MEMWAIT2 PMEM = 0xFF << 8  //+ MEMWAIT2[7:0] bits (Common memory 2 wait time).
	MEMHOLD2 PMEM = 0xFF << 16 //+ MEMHOLD2[7:0] bits (Common memory 2 hold time).
	MEMHIZ2  PMEM = 0xFF << 24 //+ MEMHIZ2[7:0] bits (Common memory 2 databus HiZ time).
)

const (
	MEMSET2n  = 0
	MEMWAIT2n = 8
	MEMHOLD2n = 16
	MEMHIZ2n  = 24
)

const (
	ATTSET2  PATT = 0xFF << 0  //+ ATTSET2[7:0] bits (Attribute memory 2 setup time).
	ATTWAIT2 PATT = 0xFF << 8  //+ ATTWAIT2[7:0] bits (Attribute memory 2 wait time).
	ATTHOLD2 PATT = 0xFF << 16 //+ ATTHOLD2[7:0] bits (Attribute memory 2 hold time).
	ATTHIZ2  PATT = 0xFF << 24 //+ ATTHIZ2[7:0] bits (Attribute memory 2 databus HiZ time).
)

const (
	ATTSET2n  = 0
	ATTWAIT2n = 8
	ATTHOLD2n = 16
	ATTHIZ2n  = 24
)

const (
	ECC2 ECCR = 0xFFFFFFFF << 0 //+ ECC result.
)

const (
	ECC2n = 0
)
