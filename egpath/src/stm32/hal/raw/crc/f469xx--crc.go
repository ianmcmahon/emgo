// +build f469xx

// Peripheral: CRC_Periph  CRC calculation unit.
// Instances:
//  CRC  mmap.CRC_BASE
// Registers:
//  0x00 32  DR  Data register.
//  0x04  8  IDR Independent data register.
//  0x08 32  CR  Control register.
// Import:
//  stm32/o/f469xx/mmap
package crc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	RESET CR = 0x01 << 0 //+ RESET bit.
)

const (
	RESETn = 0
)
