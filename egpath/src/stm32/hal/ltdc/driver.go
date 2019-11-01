// +build f469xx
package ltdc

import (
	"sync"

	"github.com/ianmcmahon/emgo/egpath/src/stm32/o/f469xx/irq"
	"github.com/ianmcmahon/emgo/egroot/src/rtos"

	"stm32/hal/dsi"
	"stm32/hal/raw/ltdc"
)

type halStatus uint8
type state uint8
type ltdcError uint32

const (
	HAL_OK      halStatus = 0x00
	HAL_ERROR             = 0x01
	HAL_BUSY              = 0x02
	HAL_TIMEOUT           = 0x03

	HAL_LTDC_STATE_RESET   state = 0x00 ///LTDC not yet initialized or disabled
	HAL_LTDC_STATE_READY         = 0x01 // LTDC initialized and ready for use
	HAL_LTDC_STATE_BUSY          = 0x02 // LTDC internal process is ongoing
	HAL_LTDC_STATE_TIMEOUT       = 0x03 // LTDC Timeout state
	HAL_LTDC_STATE_ERROR         = 0x04 // LTDC state error

	HAL_LTDC_ERROR_NONE    ltdcError = 0x00000000 // LTDC No error
	HAL_LTDC_ERROR_TE                = 0x00000001 // LTDC Transfer error
	HAL_LTDC_ERROR_FU                = 0x00000002 // LTDC FIFO Underrun
	HAL_LTDC_ERROR_TIMEOUT           = 0x00000020 // LTDC Timeout error

	LTDC_HSPOLARITY_AL ltdc.GCR = ltdc.GCR(0) // Horizontal Synchronization is active low.
	LTDC_HSPOLARITY_AH          = ltdc.HSPOL  // Horizontal Synchronization is active high.

	LTDC_VSPOLARITY_AL ltdc.GCR = ltdc.GCR(0) // Vertical Synchronization is active low.
	LTDC_VSPOLARITY_AH          = ltdc.VSPOL  // Vertical Synchronization is active high.

	LTDC_DEPOLARITY_AL ltdc.GCR = ltdc.GCR(0) // Data Enable, is active low.
	LTDC_DEPOLARITY_AH          = ltdc.DEPOL  // Data Enable, is active high.

	LTDC_PCPOLARITY_IPC  ltdc.GCR = ltdc.GCR(0) // input pixel clock.
	LTDC_PCPOLARITY_IIPC          = ltdc.PCPOL  // inverted input pixel clock.

	LTDC_HORIZONTALSYNC = ltdc.HSW >> 16 // Horizontal synchronization width.
	LTDC_VERTICALSYNC   = ltdc.VSH       // Vertical synchronization height.

	LTDC_COLOR = 0x000000FF // Color mask

	LTDC_BLENDING_FACTOR1_CA    = 0x00000400 // Blending factor : Cte Alpha
	LTDC_BLENDING_FACTOR1_PAxCA = 0x00000600 // Blending factor : Cte Alpha x Pixel Alpha

	LTDC_BLENDING_FACTOR2_CA    = 0x00000005 // Blending factor : Cte Alpha
	LTDC_BLENDING_FACTOR2_PAxCA = 0x00000007 // Blending factor : Cte Alpha x Pixel Alpha

	LTDC_PIXEL_FORMAT_ARGB8888 = 0x00000000 // ARGB8888 LTDC pixel format
	LTDC_PIXEL_FORMAT_RGB888   = 0x00000001 // RGB888 LTDC pixel format
	LTDC_PIXEL_FORMAT_RGB565   = 0x00000002 // RGB565 LTDC pixel format
	LTDC_PIXEL_FORMAT_ARGB1555 = 0x00000003 // ARGB1555 LTDC pixel format
	LTDC_PIXEL_FORMAT_ARGB4444 = 0x00000004 // ARGB4444 LTDC pixel format
	LTDC_PIXEL_FORMAT_L8       = 0x00000005 // L8 LTDC pixel format
	LTDC_PIXEL_FORMAT_AL44     = 0x00000006 // AL44 LTDC pixel format
	LTDC_PIXEL_FORMAT_AL88     = 0x00000007 // AL88 LTDC pixel format

	//LTDC_STOPPOSITION  = ltdc.WHSPPOS >> 16 // LTDC Layer stop position
	//LTDC_STARTPOSITION = ltdc.WHSTPOS       // LTDC Layer start position

	//LTDC_COLOR_FRAME_BUFFER = ltdc.CFBLL   // LTDC Layer Line length
	//LTDC_LINE_NUMBER        = ltdc.CFBLNBR // LTDC Layer Line number
)

type LTDC struct {
	Instance            *ltdc.LTDC_Periph
	State               state
	ErrorCode           ltdcError
	mutex               sync.Mutex
	LayerCfg            [2]*LayerConfig
	Layers              [2]*ltdc.LTDC_Layer_Periph
	ErrorCallback       func()
	LineEventCallback   func()
	ReloadEventCallback func()
}

type Color struct {
	Blue     uint8
	Green    uint8
	Red      uint8
	Reserved uint8
}

func (c Color) ToMem() uint32 {
	return uint32(c.Reserved)<<24 | uint32(c.Red)<<16 | uint32(c.Green)<<8 | uint32(c.Blue)
}

type Config struct {
	HSPolarity         ltdc.GCR // configures the horizontal synchronization polarity.
	VSPolarity         ltdc.GCR // configures the vertical synchronization polarity.
	DEPolarity         ltdc.GCR // configures the data enable polarity.
	PCPolarity         ltdc.GCR // configures the pixel clock polarity.
	HorizontalSync     uint32   // configures the number of Horizontal synchronization width.
	VerticalSync       uint32   // configures the number of Vertical synchronization height.
	AccumulatedHBP     uint32   // configures the accumulated horizontal back porch width.
	AccumulatedVBP     uint32   // configures the accumulated vertical back porch height.
	AccumulatedActiveW uint32   // configures the accumulated active width.
	AccumulatedActiveH uint32   // configures the accumulated active height.
	TotalWidth         uint32   // configures the total width.
	TotalHeight        uint32   // configures the total height.
	Backcolor          Color    // Configures the background color.
	LayerCfg           LayerConfig
}

type LayerConfig struct {
	WindowX0        uint32 // Configures the Window Horizontal Start Position.
	WindowX1        uint32 // Configures the Window Horizontal Stop Position.
	WindowY0        uint32 // Configures the Window vertical Start Position.
	WindowY1        uint32 // Configures the Window vertical Stop Position.
	PixelFormat     uint32 // Specifies the pixel format.
	Alpha           uint32 // Specifies the constant alpha used for blending.
	Alpha0          uint32 // Configures the default alpha value.
	BlendingFactor1 uint32 // Select the blending factor 1.
	BlendingFactor2 uint32 // Select the blending factor 2.
	FBStartAddress  uint32 // Configures the color frame buffer address
	ImageWidth      uint32 // Configures the color frame buffer line length.
	ImageHeight     uint32 // Specifies the number of line in frame buffer.
	Backcolor       Color  // Configures the layer background color.
}

var LTDCPeriph LTDC

func (p *LTDC) Init(cfg Config) halStatus {
	if p.Instance == nil {
		p.Instance = ltdc.LTDC
		p.State = HAL_LTDC_STATE_RESET
	}

	if p.State == HAL_LTDC_STATE_RESET {
		p.mutex.Unlock()
		p.MspInit()
	}

	p.State = HAL_LTDC_STATE_BUSY

	// Configure HS, VS, DE, PC polarity
	p.Instance.HSPOL().Store(ltdc.GCR(cfg.HSPolarity))
	p.Instance.VSPOL().Store(ltdc.GCR(cfg.VSPolarity))
	p.Instance.DEPOL().Store(ltdc.GCR(cfg.DEPolarity))
	p.Instance.PCPOL().Store(ltdc.GCR(cfg.PCPolarity))

	// Set Synchronization time
	p.Instance.VSH().Store(ltdc.SSCR(cfg.VerticalSync))
	p.Instance.HSW().Store(ltdc.SSCR(cfg.HorizontalSync))

	// Set Accumulated Back Porch
	p.Instance.AVBP().Store(ltdc.BPCR(cfg.AccumulatedVBP))
	p.Instance.AHBP().Store(ltdc.BPCR(cfg.AccumulatedHBP))

	// Set Accumulated Active Width
	p.Instance.AAH().Store(ltdc.AWCR(cfg.AccumulatedActiveH))
	p.Instance.AAW().Store(ltdc.AWCR(cfg.AccumulatedActiveW))

	// Set Total Width
	p.Instance.TOTALH().Store(ltdc.TWCR(cfg.TotalHeight))
	p.Instance.TOTALW().Store(ltdc.TWCR(cfg.TotalWidth))

	// Set Background Color
	p.Instance.BCBLUE().Store(ltdc.BCCR(cfg.Backcolor.Blue))
	p.Instance.BCRED().Store(ltdc.BCCR(cfg.Backcolor.Red))
	p.Instance.BCGREEN().Store(ltdc.BCCR(cfg.Backcolor.Green))

	// Enable the transfer Error interrupt
	p.Instance.TERRIE().Set()

	// Enable the FIFO underrun interrupt
	p.Instance.FUIE().Set()

	// Enable LTDC by setting LTDCEN bit
	p.Instance.LTDCEN().Set()

	p.ErrorCode = HAL_LTDC_ERROR_NONE

	p.State = HAL_LTDC_STATE_READY

	return HAL_OK
}

func (p *LTDC) MspInit() halStatus {
	// enable interrupts
	rtos.IRQ(irq.LTDC).SetPrio(3)
	rtos.IRQ(irq.LTDC).Enable()
	rtos.IRQ(irq.LTDC).UseHandler(p.IRQHandler)

	return HAL_OK
}

func (p *LTDC) InitFromVideoConfig(cfg *Config, vidCfg *dsi.VideoConfig) halStatus {
	// Retrieve signal polarities from DSI
	// The following polarity is inverted:
	//		LTDC_DEPOLARITY_AL <-> LTDC_DEPOLARITY_AH

	// Note 1: Code in line w/ Current LTDC specification (brought fwd from hal driver irm 6/21/19)
	switch vidCfg.DEPolarity {
	case dsi.DSI_DATA_ENABLE_ACTIVE_HIGH:
		cfg.DEPolarity = LTDC_DEPOLARITY_AL
	case dsi.DSI_DATA_ENABLE_ACTIVE_LOW:
		cfg.DEPolarity = LTDC_DEPOLARITY_AH
	}

	switch vidCfg.VSPolarity {
	case dsi.DSI_VSYNC_ACTIVE_HIGH:
		cfg.VSPolarity = LTDC_VSPOLARITY_AL
	case dsi.DSI_VSYNC_ACTIVE_LOW:
		cfg.VSPolarity = LTDC_VSPOLARITY_AH
	}

	switch vidCfg.HSPolarity {
	case dsi.DSI_HSYNC_ACTIVE_HIGH:
		cfg.HSPolarity = LTDC_HSPOLARITY_AL
	case dsi.DSI_HSYNC_ACTIVE_LOW:
		cfg.HSPolarity = LTDC_HSPOLARITY_AH
	}

	// Retrieve vertical timing parameters from DSI
	cfg.VerticalSync = uint32(vidCfg.VerticalSyncActive) - 1
	cfg.AccumulatedVBP = uint32(vidCfg.VerticalSyncActive) + uint32(vidCfg.VerticalBackPorch) - 1
	cfg.AccumulatedActiveH = uint32(vidCfg.VerticalSyncActive) + uint32(vidCfg.VerticalBackPorch) + uint32(vidCfg.VerticalActive) - 1
	cfg.TotalHeight = uint32(vidCfg.VerticalSyncActive) + uint32(vidCfg.VerticalBackPorch) + uint32(vidCfg.VerticalActive) + uint32(vidCfg.VerticalFrontPorch) - 1

	return HAL_OK
}

func (p *LTDC) ConfigLayer(cfg *LayerConfig, layerIdx uint16) halStatus {
	p.mutex.Lock()
	p.State = HAL_LTDC_STATE_BUSY

	p.LayerCfg[layerIdx] = cfg
	p.SetConfig(cfg, layerIdx)
	p.Instance.IMR().Set()

	p.State = HAL_LTDC_STATE_READY
	p.mutex.Unlock()
	return HAL_OK
}

func (p *LTDC) SetConfig(cfg *LayerConfig, layerIdx uint16) {
	layer := p.Layers[layerIdx]
	hbp := uint32(p.Instance.AHBP().Load())
	vbp := uint32(p.Instance.AVBP().Load())
	/* Configures the horizontal start and stop position */
	layer.WHPCR.StoreBits(0xFFF, ltdc.WHPCR(cfg.WindowX0+hbp+1))
	layer.WHPCR.StoreBits(0xFFF<<16, ltdc.WHPCR((cfg.WindowX1+hbp+1)<<16))

	/* Configures the vertical start and stop position */
	layer.WVPCR.StoreBits(0xFFF, ltdc.WVPCR(cfg.WindowY0+vbp+1))
	layer.WVPCR.StoreBits(0xFFF<<16, ltdc.WVPCR((cfg.WindowY1+vbp+1)<<16))

	/* Specifies the pixel format */
	layer.PFCR.StoreBits(0x7, ltdc.PFCR(cfg.PixelFormat))

	/* Configures the default color values */
	layer.DCCR.Store(ltdc.DCCR(cfg.Backcolor.ToMem()))

	/* Specifies the constant alpha value */
	layer.CACR.StoreBits(0xFF, ltdc.CACR(cfg.Alpha))

	/* Specifies the blending factors */
	layer.BFCR.StoreBits(0x0707, ltdc.BFCR(cfg.BlendingFactor1|cfg.BlendingFactor2))

	/* Configures the color frame buffer start address */
	layer.CFBAR.Store(ltdc.CFBAR(cfg.FBStartAddress))

	bpp := uint32(1)
	switch cfg.PixelFormat {
	case LTDC_PIXEL_FORMAT_ARGB8888:
		bpp = 4
	case LTDC_PIXEL_FORMAT_RGB888:
		bpp = 3
	case LTDC_PIXEL_FORMAT_ARGB4444:
		bpp = 2
	case LTDC_PIXEL_FORMAT_RGB565:
		bpp = 2
	case LTDC_PIXEL_FORMAT_ARGB1555:
		bpp = 2
	case LTDC_PIXEL_FORMAT_AL88:
		bpp = 2
	}

	/* Configures the color frame buffer pitch in byte */
	layer.CFBLR.StoreBits(0x1FFF, ltdc.CFBLR((cfg.WindowX1-cfg.WindowX0)*bpp+3))
	layer.CFBLR.StoreBits(0x1FFF<<16, ltdc.CFBLR(cfg.ImageWidth*bpp))

	/* Configures the frame buffer line number */
	layer.CFBLNR.StoreBits(0x7FF, ltdc.CFBLNR(cfg.ImageHeight))

	/* Enable LTDC_Layer by setting LEN bit */
	layer.CR.SetBits(1)
}

func (p *LTDC) IRQHandler() {
	// Transfer Error Interrupt management
	if p.Instance.TERRIF().Load() != 0 {
		// Disable the transfer Error interrupt
		p.Instance.TERRIE().Clear()
		// Clear the transfer error flag
		p.Instance.TERRIF().Clear()
		// Update error code
		p.ErrorCode |= HAL_LTDC_ERROR_TE
		// Change LTDC state
		p.State = HAL_LTDC_STATE_ERROR
		// Process unlocked
		p.mutex.Unlock()
		// Transfer error Callback
		p.ErrorCallback()
	}
	// FIFO underrun Interrupt management
	if p.Instance.FUIF().Load() != 0 {
		// Disable the FIFO underrun interrupt
		p.Instance.FUIE().Clear()
		// Clear the FIFO underrun flag
		p.Instance.FUIF().Clear()
		// Update error code
		p.ErrorCode |= HAL_LTDC_ERROR_FU
		// Change LTDC state
		p.State = HAL_LTDC_STATE_ERROR
		// Process unlocked
		p.mutex.Unlock()
		// Transfer error Callback
		p.ErrorCallback()
	}
	// Line Interrupt management
	if p.Instance.LIF().Load() != 0 {
		// Disable the Line interrupt
		p.Instance.LIE().Clear()
		// Clear the Line interrupt flag
		p.Instance.LIF().Clear()
		// Change LTDC state
		p.State = HAL_LTDC_STATE_READY
		// Process unlocked
		p.mutex.Unlock()
		// Line interrupt Callback
		p.LineEventCallback()
	}
	// Register reload Interrupt management
	if p.Instance.RRIF().Load() != 0 {
		// Disable the register reload interrupt
		p.Instance.RRIE().Clear()
		// Clear the register reload flag
		p.Instance.RRIF().Clear()
		// Change LTDC state
		p.State = HAL_LTDC_STATE_READY
		// Process unlocked
		p.mutex.Unlock()
		// Register reload interrupt Callback
		p.ReloadEventCallback()
	}
}
