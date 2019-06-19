// +build f40_41xxx f411xe f469xx

package gpio

import (
	"stm32/hal/raw/rcc"
)

const (
	MCO      = AF0
	RTC_50HZ = AF0
	TRACE    = AF0

	TIM1 = AF1
	TIM2 = AF1

	TIM3 = AF2
	TIM4 = AF2
	TIM5 = AF2

	TIM8  = AF3
	TIM9  = AF3
	TIM10 = AF3
	TIM11 = AF3

	I2C1 = AF4
	I2C2 = AF4
	I2C3 = AF4

	SPI1 = AF5
	SPI2 = AF5
	SPI4 = AF5
	SPI5 = AF5
	SPI6 = AF5

	SPI3 = AF6
	SAI1 = AF6

	USART1 = AF7
	USART2 = AF7
	USART3 = AF7

	UART4  = AF8
	UART5  = AF8
	USART6 = AF8
	UART7  = AF8
	UART8  = AF8

	CAN1  = AF9
	CAN2  = AF9
	TIM12 = AF9
	TIM13 = AF9
	TIM14 = AF9

	OTGFS = AF10
	OTGHS = AF10

	ETH = AF11

	FSMC    = AF12
	FMC     = AF12
	SDIO    = AF12
	OTGHSFS = AF12

	DCMI = AF13

	LTDC = AF14
)

const (
	veryLow  = -1 // Not supported.
	low      = -1 //   2 MHz (CL = 50 pF, VDD > 2.7 V)
	_        = 0  //  25 MHz (CL = 50 pF, VDD > 2.7 V)
	high     = 1  //  50 MHz (CL = 40 pF, VDD > 2.7 V)
	veryHigh = 2  // 100 MHz (CL = 30 pF, VDD > 2.7 V)
)

func enreg() *rcc.RAHB1ENR   { return &rcc.RCC.AHB1ENR }
func rstreg() *rcc.RAHB1RSTR { return &rcc.RCC.AHB1RSTR }

func lpenaclk(pnum uint) {
	rcc.RCC.AHB1LPENR.AtomicSetBits(rcc.GPIOALPEN << pnum)
}
func lpdisclk(pnum uint) {
	rcc.RCC.AHB1LPENR.AtomicClearBits(rcc.GPIOALPEN << pnum)
}
