// +build f469xx
package dsi

import (
	"stm32/hal/raw/dsi"
	"stm32/hal/raw/rcc"
	"sync"

	"github.com/ianmcmahon/emgo/egroot/src/delay"
)

type halStatus uint8
type state uint8
type dsiError uint32

const (
	HSE_VALUE uint32 = 8000000 // this is specific to the f469

	HAL_OK      halStatus = 0x00
	HAL_ERROR             = 0x01
	HAL_BUSY              = 0x02
	HAL_TIMEOUT           = 0x03

	HAL_DSI_STATE_RESET   state = 0x00
	HAL_DSI_STATE_READY         = 0x01
	HAL_DSI_STATE_ERROR         = 0x02
	HAL_DSI_STATE_BUSY          = 0x03
	HAL_DSI_STATE_TIMEOUT       = 0x04

	HAL_DSI_ERROR_NONE dsiError = 0
	HAL_DSI_ERROR_ACK           = 0x00000001 /*!< acknowledge errors          */
	HAL_DSI_ERROR_PHY           = 0x00000002 /*!< PHY related errors          */
	HAL_DSI_ERROR_TX            = 0x00000004 /*!< transmission error          */
	HAL_DSI_ERROR_RX            = 0x00000008 /*!< reception error             */
	HAL_DSI_ERROR_ECC           = 0x00000010 /*!< ECC errors                  */
	HAL_DSI_ERROR_CRC           = 0x00000020 /*!< CRC error                   */
	HAL_DSI_ERROR_PSE           = 0x00000040 /*!< Packet Size error           */
	HAL_DSI_ERROR_EOT           = 0x00000080 /*!< End Of Transmission error   */
	HAL_DSI_ERROR_OVF           = 0x00000100 /*!< FIFO overflow error         */
	HAL_DSI_ERROR_GEN           = 0x00000200 /*!< Generic FIFO related errors */

	DSI_PLL_IN_DIV1 dsi.WRPCR = 0x00000001
	DSI_PLL_IN_DIV2           = 0x00000002
	DSI_PLL_IN_DIV3           = 0x00000003
	DSI_PLL_IN_DIV4           = 0x00000004
	DSI_PLL_IN_DIV5           = 0x00000005
	DSI_PLL_IN_DIV6           = 0x00000006
	DSI_PLL_IN_DIV7           = 0x00000007

	DSI_PLL_OUT_DIV1 dsi.WRPCR = 0x00000000
	DSI_PLL_OUT_DIV2           = 0x00000001
	DSI_PLL_OUT_DIV4           = 0x00000002
	DSI_PLL_OUT_DIV8           = 0x00000003

	DSI_ONE_DATA_LANE  dsi.PCONFR = 0x00000000
	DSI_TWO_DATA_LANES            = 0x00000001

	DSI_PHY_CLK_CONTROL    dsi.CLCR = dsi.DPCC
	DSI_AUTO_CLOCK_CONTROL          = dsi.ACR
)

type DSI struct {
	Instance  *dsi.DSI_Periph
	State     state
	ErrorCode dsiError
	ErrorMsk  uint32
	mutex     sync.Mutex
}

type Config struct {
	AutoClockLaneControl dsi.CLCR
	TXEscapeClkDiv       uint32
	NumLanes             dsi.PCONFR
}

type PLLConfig struct {
	LoopDivFactor   dsi.WRPCR
	InputDivFactor  dsi.WRPCR
	OutputDivFactor dsi.WRPCR
}

var DSIPeriph DSI

func (d *DSI) Init(cfg Config, pllCfg PLLConfig) halStatus {
	if d.Instance == nil {
		d.Instance = dsi.DSI
		d.State = HAL_DSI_STATE_RESET
	}

	if d.State == HAL_DSI_STATE_RESET {
		/* Initialize the low level hardware */
		d.MspInit()
	}
	d.State = HAL_DSI_STATE_BUSY
	d.EnableRegulator()
	// Wait for Regulator Ready Flag
	// todo: needs a timeout
	for d.Instance.RRS().Load() == 0 {
		delay.Millisec(1)
	}

	// Set PLL Division factors
	d.Instance.WRPCR.ClearBits(dsi.PLL_NDIV | dsi.PLL_IDF | dsi.PLL_ODF)
	d.Instance.WRPCR.StoreBits(dsi.PLL_NDIV, (pllCfg.LoopDivFactor << dsi.PLL_NDIVn))
	d.Instance.WRPCR.StoreBits(dsi.PLL_IDF, (pllCfg.InputDivFactor << dsi.PLL_IDFn))
	d.Instance.WRPCR.StoreBits(dsi.PLL_ODF, (pllCfg.OutputDivFactor << dsi.PLL_ODFn))
	d.EnablePLL()
	// Wait for PLL Lock Status
	for d.Instance.PLLLS().Load() == 0 {
		delay.Millisec(1)
	}

	/*************************** Set the PHY parameters ***************************/

	/* D-PHY clock and digital enable*/
	d.Instance.PCTLR.SetBits(dsi.CKE | dsi.DEN)

	/* Clock lane configuration */
	d.Instance.CLCR.ClearBits(dsi.DPCC | dsi.ACR)
	d.Instance.CLCR.SetBits(dsi.DPCC | cfg.AutoClockLaneControl)

	/* Configure the number of active data lanes */
	d.Instance.PCONFR.StoreBits(dsi.NL, cfg.NumLanes<<dsi.NLn)

	/************************ Set the DSI clock parameters ************************/

	/* Set the TX escape clock division factor */
	d.Instance.CCR.StoreBits(dsi.TXECKDIV, dsi.CCR(cfg.TXEscapeClkDiv<<dsi.TXECKDIVn))

	/* Calculate the bit period in high-speed mode in unit of 0.25 ns (UIX4) */
	/* The equation is : UIX4 = IntegerPart( (1000/F_PHY_Mhz) * 4 )          */
	/* Where : F_PHY_Mhz = (NDIV * HSE_Mhz) / (IDF * ODF)                    */
	tempIDF := pllCfg.InputDivFactor
	if tempIDF == 0 {
		tempIDF = 1
	}
	unitIntervalx4 := (4000000 * uint32(tempIDF) * (1 << uint32(pllCfg.OutputDivFactor))) / ((HSE_VALUE / 1000) * uint32(pllCfg.LoopDivFactor))

	/* Set the bit period in high-speed mode */
	d.Instance.WPCR[0].StoreBits(dsi.HSTXDCL, dsi.WPCR(unitIntervalx4))

	/* Disable all error interrupts and reset the Error Mask */
	d.Instance.IER[0].Store(0)
	d.Instance.IER[1].Store(0)
	d.ErrorMsk = 0
	d.ErrorCode = HAL_DSI_ERROR_NONE
	d.State = HAL_DSI_STATE_READY

	return HAL_OK
}

func (d *DSI) DeInit() halStatus {
	if d.Instance == nil {
		return HAL_ERROR
	}

	d.State = HAL_DSI_STATE_BUSY
	d.DisableWrapper()
	d.DisableHost()

	// D-PHY clock and digital disable
	d.Instance.PCTLR.ClearBits(dsi.CKE | dsi.DEN)

	d.DisablePLL()
	d.DisableRegulator()

	/* DeInit the low level hardware */
	// this is unimplemented in st's hal, but stubbed so it can be overridden
	// d.MspDeInit()

	/* Initialise the error code */
	d.ErrorCode = HAL_DSI_ERROR_NONE

	/* Initialize the DSI state*/
	d.State = HAL_DSI_STATE_RESET

	d.mutex.Unlock()

	return HAL_OK
}

func (d *DSI) MspInit() {
	// LTDC Clock Enable
	rcc.RCC.LTDCEN().Set()
	delay.Loop(2)

	// Toggle SW Reset of LTDC IP
	rcc.RCC.LTDCRST().Set()
	rcc.RCC.LTDCRST().Clear()

	// Enable DMA2D Clk
	rcc.RCC.DMA2DEN().Set()
	delay.Loop(2)

	// Toggle SW Reset of DMA2D IP
	rcc.RCC.DMA2DRST().Set()
	rcc.RCC.DMA2DRST().Clear()

	// Enable DSI Clk
	rcc.RCC.DSIEN().Set()
	delay.Loop(2)

	// Toggle SW Reset of DSI IP
	rcc.RCC.DSIRST().Set()
	rcc.RCC.DSIRST().Clear()

	// TODO: setup interrupts:
	//   HAL_NVIC_SetPriority(LTDC_IRQn, 3, 0);
	//   HAL_NVIC_EnableIRQ(LTDC_IRQn);
	//
	//   /** @brief NVIC configuration for DMA2D interrupt that is now enabled */
	//   HAL_NVIC_SetPriority(DMA2D_IRQn, 3, 0);
	//   HAL_NVIC_EnableIRQ(DMA2D_IRQn);
	//
	//   /** @brief NVIC configuration for DSI interrupt that is now enabled */
	//   HAL_NVIC_SetPriority(DSI_IRQn, 3, 0);
	//   HAL_NVIC_EnableIRQ(DSI_IRQn);
}

func (d *DSI) GetError() dsiError {
	return d.ErrorCode
}

func (d *DSI) EnableHost() {
	d.Instance.CR.SetBits(dsi.CR_EN)
}

func (d *DSI) DisableHost() {
	d.Instance.CR.ClearBits(dsi.CR_EN)
}

func (d *DSI) EnableWrapper() {
	d.Instance.WCR.SetBits(dsi.DSIEN)
}

func (d *DSI) DisableWrapper() {
	d.Instance.WCR.ClearBits(dsi.DSIEN)
}

func (d *DSI) EnablePLL() {
	d.Instance.WRPCR.SetBits(dsi.PLLEN)
}

func (d *DSI) DisablePLL() {
	d.Instance.WRPCR.ClearBits(dsi.PLLEN)
}

func (d *DSI) EnableRegulator() {
	d.Instance.WRPCR.SetBits(dsi.REGEN)
}

func (d *DSI) DisableRegulator() {
	d.Instance.WRPCR.ClearBits(dsi.REGEN)
}
