// Peripheral: IWDG_Periph  Independent WATCHDOG.
// Instances:
//  IWDG  mmap.IWDG_BASE
// Registers:
//  0x00 32  KR  Key register.
//  0x04 32  PR  Prescaler register.
//  0x08 32  RLR Reload register.
//  0x0C 32  SR  Status register.
// Import:
//  /Users/ian/code/src/github.com/ianmcmahon/emgo/egpath/src/stm32/o/f469xx/mmap
package iwdg

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	KEY KR = 0xFFFF << 0 //+ Key value (write only, read 0000h).
)

const (
	KEYn = 0
)

const (
	RL RLR = 0xFFF << 0 //+ Watchdog counter reload value.
)

const (
	RLn = 0
)

const (
	PVU SR = 0x01 << 0 //+ Watchdog prescaler value update.
	RVU SR = 0x01 << 1 //+ Watchdog counter reload value update.
)

const (
	PVUn = 0
	RVUn = 1
)
