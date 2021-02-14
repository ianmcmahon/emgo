// +build f469xx
package display

import (
	"delay"
	"stm32/hal/display/otm8009a"
	"stm32/hal/dsi"
	"stm32/hal/gpio"
	"stm32/hal/ltdc"
	"stm32/hal/raw/rcc"
)

type orientation uint32

const (
	LcdClock uint32 = 27429

	OTM8009A_FORMAT_RGB888 = 0x00 // Pixel format chosen is RGB888 : 24 bpp
	OTM8009A_FORMAT_RBG565 = 0x02 // Pixel format chosen is RGB565 : 16 bpp

	LCD_ORIENTATION_LANDSCAPE orientation = 0x00
	LCD_ORIENTATION_PORTRAIT              = 0x01

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

	RCC_PLLSAIDIVR_2  = 0x0000
	RCC_PLLSAIDIVR_4  = 0x0001
	RCC_PLLSAIDIVR_8  = 0x0002
	RCC_PLLSAIDIVR_16 = 0x0003
)

var (
	lcd_x_size uint32 = LANDSCAPE_WIDTH
	lcd_y_size uint32 = LANDSCAPE_HEIGHT
)

func Reset() {
	gpio.H.EnableClock(true)
	xres := gpio.H.Pin(7)
	xres.Setup(&gpio.Config{Mode: gpio.Out, Driver: gpio.OpenDrain, Speed: gpio.High})
	xres.Clear()
	delay.Millisec(20)
	xres.Set()
	delay.Millisec(10)
}

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

	dsi.DSIPeriph.Init(cfg, pllCfg)

	if orient == LCD_ORIENTATION_PORTRAIT {
		lcd_x_size = PORTRAIT_WIDTH
		lcd_y_size = PORTRAIT_HEIGHT
	}

	HACT := lcd_x_size
	VACT := lcd_y_size

	vidCfg := dsi.VideoConfig{
		VirtualChannelID:             0,
		ColorCoding:                  dsi.DSI_RGB888,
		Mode:                         dsi.DSI_VID_MODE_BURST,
		PacketSize:                   dsi.VPCR(HACT),
		NumberOfChunks:               0,
		NullPacketSize:               0xFFF,
		HSPolarity:                   dsi.DSI_VSYNC_ACTIVE_HIGH,
		VSPolarity:                   dsi.DSI_HSYNC_ACTIVE_HIGH,
		DEPolarity:                   dsi.DSI_DATA_ENABLE_ACTIVE_HIGH,
		HorizontalSyncActive:         dsi.VHSACR((TIMING_HSYNC * laneClkSpeed) / LcdClock),
		HorizontalBackPorch:          dsi.VHBPCR(TIMING_HBP * laneClkSpeed / LcdClock),
		HorizontalLine:               dsi.VLCR((HACT + TIMING_HSYNC + TIMING_HBP + TIMING_HFP) * laneClkSpeed / LcdClock),
		VerticalSyncActive:           TIMING_VSYNC,
		VerticalBackPorch:            TIMING_VBP,
		VerticalFrontPorch:           TIMING_VFP,
		VerticalActive:               dsi.VVACR(VACT),
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

	ltdcCfg := ltdc.Config{
		HorizontalSync:     TIMING_HSYNC - 1,
		AccumulatedHBP:     TIMING_HSYNC + TIMING_HBP - 1,
		AccumulatedActiveW: lcd_x_size + TIMING_HSYNC + TIMING_HBP - 1,
		TotalWidth:         lcd_x_size + TIMING_HSYNC + TIMING_HBP + TIMING_HFP - 1,
		LayerCfg: ltdc.LayerConfig{
			ImageWidth:  lcd_x_size,
			ImageHeight: lcd_y_size,
		},
		PCPolarity: ltdc.LTDC_PCPOLARITY_IPC,
		Backcolor:  ltdc.Color{0x33, 0x55, 0x77, 0x00},
	}

	// Configure PLL for LTDC
	// stm32f4xx_hal does this through a big func in rcc_ex that handles every possible case
	// for now let's set it up manually
	// PLLSAI_VCO Input = HSE_VALUE/PLL_M = 8MHz/8 = 1MHz
	// PLLSAI_VCO Output = PLLSAI_VCO Input * PLLSAIN = 384MHz
	// PLLLCLOCK = PLLSAI_VCO Output/PLLSAIR = 384/7 = 54.857Mhz

	// disable pllsai clock
	rcc.RCC.PLLSAION().Clear()
	// Wait for PLL to disable
	for rcc.RCC.PLLSAIRDY().Load() != 0 {
		delay.Millisec(1)
	}
	rcc.RCC.PLLSAIN().Store(384)
	rcc.RCC.PLLSAIR().Store(7)
	rcc.RCC.PLLSAIDIVR().Store(RCC_PLLSAIDIVR_2)

	// enable pllsai clock
	rcc.RCC.PLLSAION().Set()
	// Wait for PLL to lock
	for rcc.RCC.PLLSAIRDY().Load() == 0 {
		delay.Millisec(1)
	}

	ltdc.LTDCPeriph.InitFromVideoConfig(&ltdcCfg, &vidCfg)

	ltdc.LTDCPeriph.Init(ltdcCfg)

	dsi.DSIPeriph.Start()

	// Initialize SDRAM
	// BSP_SDRAM_Init()

	// Initialize font (prolly not)
	// BSP_LCD_SetFont(&LCD_DEFAULT_FONT)

	// Initialize OTM8009A
	otm8009a.Init(otm8009a.OTM8009A_FORMAT_RGB888, uint32(orient))
}

func GetXSize() uint32 {
	return lcd_x_size
}

func GetYSize() uint32 {
	return lcd_y_size
}

func SetXSize(x uint32) {
	// todo: implementme
}

func SetYSize(x uint32) {
	// todo: implementme
}

/**
 * Initializes the LCD layers.
 * LayerIndex: Layer foreground or background
 * FB_Address: Layer frame buffer
 */
func LayerDefaultInit(LayerIndex uint16, FB_Address uint32) {
	Layercfg := ltdc.LayerConfig{
		WindowX0:        0,
		WindowX1:        GetXSize(),
		WindowY0:        0,
		WindowY1:        GetYSize(),
		PixelFormat:     ltdc.LTDC_PIXEL_FORMAT_ARGB8888,
		FBStartAddress:  FB_Address,
		Alpha:           0xFF,
		Alpha0:          0,
		Backcolor:       ltdc.Color{0, 0, 0, 0},
		BlendingFactor1: ltdc.LTDC_BLENDING_FACTOR1_PAxCA,
		BlendingFactor2: ltdc.LTDC_BLENDING_FACTOR2_PAxCA,
		ImageWidth:      GetXSize(),
		ImageHeight:     GetYSize(),
	}

	ltdc.LTDCPeriph.ConfigLayer(&Layercfg, LayerIndex)
}
