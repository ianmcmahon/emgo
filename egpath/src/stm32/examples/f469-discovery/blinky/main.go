package main

import (
	"delay"
	"fmt"

	"stm32/hal/dsi"
	"stm32/hal/gpio"
	"stm32/hal/system"
	"stm32/hal/system/timer/systick"
)

var green, orange, red, blue gpio.Pin
var dsiHandle *dsi.DSI

func initDisplay() {
	cfg := dsi.Config{
		NumLanes:       dsi.DSI_TWO_DATA_LANES,
		TXEscapeClkDiv: 62500 / 15620,
	}
	pllCfg := dsi.PLLConfig{
		LoopDivFactor:   125,
		InputDivFactor:  dsi.DSI_PLL_IN_DIV2,
		OutputDivFactor: dsi.DSI_PLL_OUT_DIV1,
	}

	// gotta reset the LCD first dummy
	gpio.H.EnableClock(true)
	xres := gpio.H.Pin(7)
	xres.Setup(&gpio.Config{Mode: gpio.Out, Driver: gpio.OpenDrain, Speed: gpio.High})
	xres.Clear()
	delay.Millisec(20)
	xres.Set()
	delay.Millisec(10)

	dsi.DSIPeriph.Init(cfg, pllCfg)

}

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

	initDisplay()
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
