// +build f469xx
package otm8009a

import (
	"delay"
	"stm32/hal/dsi"
	"stm32/hal/gpio"
)

type orientation uint8

const (
	LcdClock uint32 = 27429

	LCD_ORIENTATION_LANDSCAPE orientation = 0x00
	LCD_ORIENTATION_PORTRAIT              = 0x01

	//  pixel data format ie color coding transmitted on DSI Data lane in DSI packets
	FORMAT_RGB888 uint32 = 0x00 // Pixel format chosen is RGB888 : 24 bpp
	FORMAT_RBG565        = 0x02 // Pixel format chosen is RGB565 : 16 bpp

	// Width and Height in Portrait mode
	PORTRAIT_WIDTH  uint32 = 480 // LCD PIXEL WIDTH
	PORTRAIT_HEIGHT        = 800 // LCD PIXEL HEIGHT

	// Width and Height in Landscape mode
	LANDSCAPE_WIDTH  uint32 = 800 // LCD PIXEL WIDTH
	LANDSCAPE_HEIGHT        = 480 // LCD PIXEL HEIGHT

	// Timing parameters work for either orientation
	TIMING_HSYNC uint32 = 2  // Horizontal synchronization
	TIMING_HBP          = 34 // Horizontal back porch
	TIMING_HFP          = 34 // Horizontal front porch
	TIMING_VSYNC        = 1  // Vertical synchronization
	TIMING_VBP          = 15 // Vertical back porch
	TIMING_VFP          = 16 // Vertical front porch
)

func InitDisplay(orient orientation) {
	var laneClkSpeed uint32 = 62500
	cfg := dsi.Config{
		NumLanes:       dsi.DSI_TWO_DATA_LANES,
		TXEscapeClkDiv: laneClkSpeed / 15620,
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

	var lcd_x_size uint32 = LANDSCAPE_WIDTH
	var lcd_y_size uint32 = LANDSCAPE_HEIGHT
	if orient == LCD_ORIENTATION_PORTRAIT {
		lcd_x_size = PORTRAIT_WIDTH
		lcd_y_size = PORTRAIT_HEIGHT
	}

	HACT := lcd_x_size
	VACT := lcd_y_size

	vidCfg := dsi.VideoConfig{
		VirtualChannelID:             0,
		ColorCoding:                  FORMAT_RGB888,
		Mode:                         dsi.DSI_VID_MODE_BURST,
		PacketSize:                   HACT,
		NumberOfChunks:               0,
		NullPacketSize:               0xFFF,
		HSPolarity:                   dsi.DSI_VSYNC_ACTIVE_HIGH,
		VSPolarity:                   dsi.DSI_HSYNC_ACTIVE_HIGH,
		DEPolarity:                   dsi.DSI_DATA_ENABLE_ACTIVE_HIGH,
		HorizontalSyncActive:         (TIMING_HSYNC * laneClkSpeed) / LcdClock,
		HorizontalBackPorch:          (TIMING_HBP * laneClkSpeed) / LcdClock,
		HorizontalLine:               ((HACT + TIMING_HSYNC + TIMING_HBP + TIMING_HFP) * laneClkSpeed) / LcdClock,
		VerticalSyncActive:           TIMING_VSYNC,
		VerticalBackPorch:            TIMING_VBP,
		VerticalFrontPorch:           TIMING_VFP,
		VerticalActive:               VACT,
		LPCommandEnable:              dsi.DSI_LP_COMMAND_ENABLE,
		LPLargestPacketSize:          16,
		LPVACTLargestPacketSize:      0,
		LPHorizontalFrontPorchEnable: dsi.DSI_LP_HFP_ENABLE,
		LPHorizontalBackPorchEnable:  dsi.DSI_LP_HBP_ENABLE,
		LPVerticalActiveEnable:       dsi.DSI_LP_VACT_ENABLE,
		LPVerticalFrontPorchEnable:   dsi.DSI_LP_VFP_ENABLE,
		LPVerticalBackPorchEnable:    dsi.DSI_LP_VBP_ENABLE,
		LPVerticalSyncActiveEnable:   dsi.DSI_LP_VSYNC_ENABLE,
	}

	dsi.DSIPeriph.ConfigVideoMode(vidCfg)
}
