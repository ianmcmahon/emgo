package main

import (
	"delay"
	"fmt"

	"stm32/hal/display/otm8009a"
	"stm32/hal/gpio"
	"stm32/hal/system"
	"stm32/hal/system/timer/systick"
)

var green, orange, red, blue gpio.Pin

func init() {
	system.Setup168(8)
	systick.Setup(2e6)

	gpio.D.EnableClock(false)
	gpio.G.EnableClock(false)
	gpio.K.EnableClock(false)
	green = gpio.G.Pin(6)
	orange = gpio.D.Pin(4)
	red = gpio.D.Pin(5)
	blue = gpio.K.Pin(3)

	cfg := gpio.Config{Mode: gpio.Out, Speed: gpio.Low}

	green.Setup(&cfg)
	orange.Setup(&cfg)
	red.Setup(&cfg)
	blue.Setup(&cfg)

	otm8009a.InitDisplay(otm8009a.LCD_ORIENTATION_LANDSCAPE)
}

func wait() {
	//delay.Loop(1e7)
	delay.Millisec(500)
}

func main() {
	for {
		blue.Clear()
		green.Clear()
		orange.Set()
		red.Set()

		//		fmt.Println("blink")

		wait()
		red.Clear()
		blue.Set()
		orange.Clear()
		green.Set()
		wait()
		fmt.Println("BLINK")
	}
}
