// Peripheral: FMC_Bank1E_Periph  Flexible Memory Controller Bank1E.
// Instances:
//  FMC_Bank1E  mmap.FMC_Bank1E_R_BASE
// Registers:
//  0x00 32  BWTR[7] NOR/PSRAM write timing registers.
// Import:
//  stm32/o/f469xx/mmap
package fmc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	EADDSET  BWTR = 0x0F << 0  //+ ADDSET[3:0] bits (Address setup phase duration).
	EADDHLD  BWTR = 0x0F << 4  //+ ADDHLD[3:0] bits (Address-hold phase duration).
	EDATAST  BWTR = 0xFF << 8  //+ DATAST [3:0] bits (Data-phase duration).
	EBUSTURN BWTR = 0x0F << 16 //+ BUSTURN[3:0] bits (Bus turnaround duration).
	EACCMOD  BWTR = 0x03 << 28 //+ ACCMOD[1:0] bits (Access mode).
)

const (
	EADDSETn  = 0
	EADDHLDn  = 4
	EDATASTn  = 8
	EBUSTURNn = 16
	EACCMODn  = 28
)
