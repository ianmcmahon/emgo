// +build f40_41xxx f411xe f469xx

package i2c

import (
	"unsafe"

	"stm32/hal/raw/mmap"
)

//emgo:const
var (
	I2C1 = (*Periph)(unsafe.Pointer(mmap.I2C1_BASE))
	I2C2 = (*Periph)(unsafe.Pointer(mmap.I2C2_BASE))
	I2C3 = (*Periph)(unsafe.Pointer(mmap.I2C3_BASE))
)
