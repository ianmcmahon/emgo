// +build f469xx
package dsi

import (
	"delay"
	"rtos"
	"stm32/hal/raw/dsi"
	"stm32/hal/raw/rcc"
	"stm32/o/f469xx/irq"
	"sync"
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
	HAL_DSI_ERROR_ACK           = 0x00000001 //// acknowledge errors
	HAL_DSI_ERROR_PHY           = 0x00000002 //// PHY related errors
	HAL_DSI_ERROR_TX            = 0x00000004 //// transmission error
	HAL_DSI_ERROR_RX            = 0x00000008 //// reception error
	HAL_DSI_ERROR_ECC           = 0x00000010 //// ECC errors
	HAL_DSI_ERROR_CRC           = 0x00000020 //// CRC error
	HAL_DSI_ERROR_PSE           = 0x00000040 //// Packet Size error
	HAL_DSI_ERROR_EOT           = 0x00000080 //// End Of Transmission error
	HAL_DSI_ERROR_OVF           = 0x00000100 //// FIFO overflow error
	HAL_DSI_ERROR_GEN           = 0x00000200 //// Generic FIFO related errors

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

	DSI_VID_MODE_NB_PULSES dsi.VMCR = 0
	DSI_VID_MODE_NB_EVENTS          = 1
	DSI_VID_MODE_BURST              = 2

	DSI_HSYNC_ACTIVE_HIGH       dsi.LPCR = dsi.LPCR(0)
	DSI_HSYNC_ACTIVE_LOW                 = dsi.HSP
	DSI_VSYNC_ACTIVE_HIGH                = dsi.LPCR(0)
	DSI_VSYNC_ACTIVE_LOW                 = dsi.VSP
	DSI_DATA_ENABLE_ACTIVE_HIGH          = dsi.LPCR(0)
	DSI_DATA_ENABLE_ACTIVE_LOW           = dsi.DEP

	DSI_LP_COMMAND_DISABLE dsi.VMCR = dsi.VMCR(0)
	DSI_LP_COMMAND_ENABLE           = dsi.VMCR_LPCE

	DSI_LP_HFP_DISABLE dsi.VMCR = dsi.VMCR(0)
	DSI_LP_HFP_ENABLE           = dsi.LPHFPE

	DSI_LP_HBP_DISABLE dsi.VMCR = dsi.VMCR(0)
	DSI_LP_HBP_ENABLE           = dsi.VMCR_LPHBPE

	DSI_LP_VACT_DISABLE dsi.VMCR = dsi.VMCR(0)
	DSI_LP_VACT_ENABLE           = dsi.VMCR_LPVAE

	DSI_LP_VFP_DISABLE dsi.VMCR = dsi.VMCR(0)
	DSI_LP_VFP_ENABLE           = dsi.VMCR_LPVFPE

	DSI_LP_VBP_DISABLE dsi.VMCR = dsi.VMCR(0)
	DSI_LP_VBP_ENABLE           = dsi.VMCR_LPVBPE

	DSI_LP_VSYNC_DISABLE dsi.VMCR = dsi.VMCR(0)
	DSI_LP_VSYNC_ENABLE           = dsi.VMCR_LPVSAE

	DSI_RGB565 dsi.LCOLCR = 0x00
	DSI_RGB666            = 0x03
	DSI_RGB888            = 0x05

	DSI_DCS_SHORT_PKT_WRITE_P0 dsi.GHCR = 0x00000005 // DCS short write, no parameters
	DSI_DCS_SHORT_PKT_WRITE_P1          = 0x00000015 // DCS short write, one parameter
	DSI_GEN_SHORT_PKT_WRITE_P0          = 0x00000003 // Generic short write, no parameters
	DSI_GEN_SHORT_PKT_WRITE_P1          = 0x00000013 // Generic short write, one parameter
	DSI_GEN_SHORT_PKT_WRITE_P2          = 0x00000023 // Generic short write, two parameters

	DSI_DCS_LONG_PKT_WRITE dsi.GHCR = 0x00000039 // DCS long write
	DSI_GEN_LONG_PKT_WRITE          = 0x00000029 // Generic long write
)

type DSI struct {
	Instance              *dsi.DSI_Periph
	Cfg                   Config
	VidCfg                VideoConfig
	State                 state
	ErrorCode             dsiError
	ErrorMsk              uint32
	mutex                 sync.Mutex
	TearingEffectCallback func()
	EndOfRefreshCallback  func()
}

type Config struct {
	AutoClockLaneControl dsi.CLCR
	TXEscapeClkDiv       uint32
	NumLanes             dsi.PCONFR
}

type VideoConfig struct {
	VirtualChannelID             dsi.LVCIDR // Virtual channel ID
	ColorCoding                  dsi.LCOLCR // Color coding for LTDC interface
	LooselyPacked                dsi.LCOLCR // Enable or disable loosely packed stream (needed only when using 18-bit configuration).
	Mode                         dsi.VMCR   // Video mode type
	PacketSize                   dsi.VPCR   // Video packet size
	NumberOfChunks               dsi.VCCR   // Number of chunks
	NullPacketSize               dsi.VNPCR  // Null packet size
	HSPolarity                   dsi.LPCR   // HSYNC pin polarity
	VSPolarity                   dsi.LPCR   // VSYNC pin polarity
	DEPolarity                   dsi.LPCR   // Data Enable pin polarity
	HorizontalSyncActive         dsi.VHSACR // Horizontal synchronism active duration (in lane byte clock cycles)
	HorizontalBackPorch          dsi.VHBPCR // Horizontal back-porch duration (in lane byte clock cycles)
	HorizontalLine               dsi.VLCR   // Horizontal line duration (in lane byte clock cycles)
	VerticalSyncActive           dsi.VVSACR // Vertical synchronism active duration
	VerticalBackPorch            dsi.VVBPCR // Vertical back-porch duration
	VerticalFrontPorch           dsi.VVFPCR // Vertical front-porch duration
	VerticalActive               dsi.VVACR  // Vertical active duration
	LPCommandEnable              dsi.VMCR   // Low-power command enable
	LPLargestPacketSize          dsi.LPMCR  // The size, in bytes, of the low power largest packet that can fit in a line during VSA, VBP and VFP regions
	LPVACTLargestPacketSize      dsi.VMCR   // The size, in bytes, of the low power largest packet that can fit in a line during VACT region
	LPHorizontalFrontPorchEnable dsi.VMCR   // Low-power horizontal front-porch enable This parameter can be any value of @ref DSI_LP_HFP
	LPHorizontalBackPorchEnable  dsi.VMCR   // Low-power horizontal back-porch enable This parameter can be any value of @ref DSI_LP_HBP
	LPVerticalActiveEnable       dsi.VMCR   // Low-power vertical active enable This parameter can be any value of @ref DSI_LP_VACT
	LPVerticalFrontPorchEnable   dsi.VMCR   // Low-power vertical front-porch enable This parameter can be any value of @ref DSI_LP_VFP
	LPVerticalBackPorchEnable    dsi.VMCR   // Low-power vertical back-porch enable This parameter can be any value of @ref DSI_LP_VBP
	LPVerticalSyncActiveEnable   dsi.VMCR   // Low-power vertical sync active enable This parameter can be any value of @ref DSI_LP_VSYNC
	FrameBTAAcknowledgeEnable    dsi.VMCR   // Frame bus-turn-around acknowledge enable This parameter can be any value of @ref DSI_FBTA_acknowledge
}

type PhyTimings struct {
	ClockLaneHS2LPTime  uint32 // The maximum time that the D-PHY clock lane takes to go from high-speed to low-power transmission
	ClockLaneLP2HSTime  uint32 // The maximum time that the D-PHY clock lane takes to go from low-power to high-speed transmission
	DataLaneHS2LPTime   uint32 // The maximum time that the D-PHY data lanes takes to go from high-speed to low-power transmission
	DataLaneLP2HSTime   uint32 // The maximum time that the D-PHY data lanes takes to go from low-power to high-speed transmission
	DataLaneMaxReadTime uint32 // The maximum time required to perform a read command
	StopWaitTime        uint32 // The minimum wait period to request a High-Speed transmission after the Stop state
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
	d.Cfg = cfg

	if d.State == HAL_DSI_STATE_RESET {
		// Initialize the low level hardware
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

	//************************** Set the PHY parameters **************************

	// D-PHY clock and digital enable
	d.Instance.PCTLR.SetBits(dsi.CKE | dsi.DEN)

	// Clock lane configuration
	d.Instance.CLCR.ClearBits(dsi.DPCC | dsi.ACR)
	d.Instance.CLCR.SetBits(dsi.DPCC | cfg.AutoClockLaneControl)

	// Configure the number of active data lanes
	d.Instance.PCONFR.StoreBits(dsi.NL, cfg.NumLanes<<dsi.NLn)

	//*********************** Set the DSI clock parameters ***********************

	// Set the TX escape clock division factor
	d.Instance.CCR.StoreBits(dsi.TXECKDIV, dsi.CCR(cfg.TXEscapeClkDiv<<dsi.TXECKDIVn))

	// Calculate the bit period in high-speed mode in unit of 0.25 ns (UIX4)
	// The equation is : UIX4 = IntegerPart( (1000/F_PHY_Mhz) * 4 )
	// Where : F_PHY_Mhz = (NDIV * HSE_Mhz) / (IDF * ODF)
	tempIDF := pllCfg.InputDivFactor
	if tempIDF == 0 {
		tempIDF = 1
	}
	unitIntervalx4 := (4000000 * uint32(tempIDF) * (1 << uint32(pllCfg.OutputDivFactor))) / ((HSE_VALUE / 1000) * uint32(pllCfg.LoopDivFactor))

	// Set the bit period in high-speed mode
	d.Instance.WPCR[0].StoreBits(dsi.HSTXDCL, dsi.WPCR(unitIntervalx4))

	// Disable all error interrupts and reset the Error Mask
	d.Instance.IER[0].Store(0)
	d.Instance.IER[1].Store(0)
	d.ErrorMsk = 0
	d.ErrorCode = HAL_DSI_ERROR_NONE
	d.State = HAL_DSI_STATE_READY

	return HAL_OK
}

func (d *DSI) Start() halStatus {
	d.mutex.Lock()
	d.Instance.CR_EN().Set()
	d.Instance.DSIEN().Set()
	d.mutex.Unlock()
	return HAL_OK
}

func (d *DSI) Stop() halStatus {
	d.mutex.Lock()
	d.Instance.CR_EN().Clear()
	d.Instance.DSIEN().Clear()
	d.mutex.Unlock()
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

	// DeInit the low level hardware
	d.MspDeInit()

	// Initialise the error code
	d.ErrorCode = HAL_DSI_ERROR_NONE

	// Initialize the DSI state
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

	/*
		1906 __STATIC_INLINE void NVIC_SetPriority(IRQn_Type IRQn, uint32_t priority) {
		1908   if ((int32_t)(IRQn) < 0) {
		1910     SCB->SHPR[(((uint32_t)(int32_t)IRQn) & 0xFUL)-4UL] = (uint8_t)((priority << (8U - __NVIC_PRIO_BITS)) & (uint32_t)0xFFUL);
		1911   } else {
		1914     NVIC->IP[((uint32_t)(int32_t)IRQn)]                = (uint8_t)((priority << (8U - __NVIC_PRIO_BITS)) & (uint32_t)0xFFUL);
		1915   }
		1916 }

		1621 /
		1622   \brief   Enable External Interrupt
		1623   \details Enables a device-specific interrupt in the NVIC interrupt controller.
		1624   \param [in]      IRQn  External interrupt number. Value cannot be negative.
		1625  /
		1626 __STATIC_INLINE void NVIC_EnableIRQ(IRQn_Type IRQn) {
		1628   NVIC->ISER[(((uint32_t)(int32_t)IRQn) >> 5UL)] = (uint32_t)(1UL << (((uint32_t)(int32_t)IRQn) & 0x1FUL));
		1629 }
	*/

	rtos.IRQ(irq.LTDC).SetPrio(3)
	rtos.IRQ(irq.LTDC).Enable()

	rtos.IRQ(irq.DMA2D).SetPrio(3)
	rtos.IRQ(irq.DMA2D).Enable()

	rtos.IRQ(irq.DSI).SetPrio(3)
	rtos.IRQ(irq.DSI).Enable()
	rtos.IRQ(irq.DSI).UseHandler(d.IRQHandler)
}

func (d *DSI) MspDeInit() {
	// todo: implement
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

func (d *DSI) ConfigVideoMode(vidCfg VideoConfig) halStatus {
	d.mutex.Lock()

	// Select video mode by resetting CMDM and DSIM bits
	d.Instance.CMDM().Clear()
	d.Instance.DSIM().Clear()

	// Configure the video mode transmission type
	d.Instance.VMCR_VMT().Store(vidCfg.Mode)

	// Configure the video packet size
	d.Instance.VPCR_VPSIZE().Store(vidCfg.PacketSize)

	// Set the chunks number to be transmitted through the DSI link
	d.Instance.VCCR_NUMC().Store(vidCfg.NumberOfChunks)

	// Set the size of the null packet
	d.Instance.VNPCR_NPSIZE().Store(vidCfg.NullPacketSize)

	// Select the virtual channel for the LTDC interface traffic
	d.Instance.LVCIDR_VCID().Store(vidCfg.VirtualChannelID)

	// Configure the polarity of control signals
	d.Instance.DEP().Store(vidCfg.DEPolarity)
	d.Instance.VSP().Store(vidCfg.VSPolarity)
	d.Instance.HSP().Store(vidCfg.HSPolarity)

	// Select the color coding for the host
	d.Instance.LCOLCR_COLC().Store(vidCfg.ColorCoding)

	// Select the color coding for the wrapper
	// TODO: this doesn't appear to be storing properly
	d.Instance.COLMUX().Store(dsi.WCFGR(vidCfg.ColorCoding))

	// Enable/disable the loosely packed variant to 18-bit configuration
	if vidCfg.ColorCoding == DSI_RGB666 {
		d.Instance.LCOLCR_LPE().Store(vidCfg.LooselyPacked)
	}

	// Set the Horizontal Synchronization Active (HSA) in lane byte clock cycles
	d.Instance.VHSACR_HSA().Store(vidCfg.HorizontalSyncActive)

	// Set the Horizontal Back Porch (HBP) in lane byte clock cycles
	d.Instance.VHBPCR_HBP().Store(vidCfg.HorizontalBackPorch)

	// Set the total line time (HLINE=HSA+HBP+HACT+HFP) in lane byte clock cycles
	d.Instance.VLCR_HLINE().Store(vidCfg.HorizontalLine)

	// Set the Vertical Synchronization Active (VSA)
	d.Instance.VVSACR_VSA().Store(vidCfg.VerticalSyncActive)

	// Set the Vertical Back Porch (VBP)
	d.Instance.VVBPCR_VBP().Store(vidCfg.VerticalBackPorch)

	// Set the Vertical Front Porch (VFP)
	d.Instance.VVFPCR_VFP().Store(vidCfg.VerticalFrontPorch)

	// Set the Vertical Active period
	d.Instance.VVACR_VA().Store(vidCfg.VerticalActive)

	// Configure the command transmission mode
	d.Instance.VMCR_LPCE().Store(vidCfg.LPCommandEnable)

	// Low power largest packet size
	d.Instance.LPMCR_LPSIZE().Store(vidCfg.LPLargestPacketSize)

	// Low power VACT largest packet size
	d.Instance.LPMCR_VLPSIZE().Store(dsi.LPMCR(vidCfg.LPVACTLargestPacketSize))

	// Enable LP transition in HFP period
	d.Instance.LPHFPE().Store(vidCfg.LPHorizontalFrontPorchEnable)

	// Enable LP transition in HBP period
	d.Instance.VMCR_LPHBPE().Store(vidCfg.LPHorizontalBackPorchEnable)

	// Enable LP transition in VACT period
	d.Instance.VMCR_LPVAE().Store(vidCfg.LPVerticalActiveEnable)

	// Enable LP transition in VFP period
	d.Instance.VMCR_LPVFPE().Store(vidCfg.LPVerticalFrontPorchEnable)

	// Enable LP transition in VBP period
	d.Instance.VMCR_LPVBPE().Store(vidCfg.LPVerticalBackPorchEnable)

	// Enable LP transition in vertical sync period
	d.Instance.VMCR_LPVSAE().Store(vidCfg.LPVerticalSyncActiveEnable)

	// Enable the request for an acknowledge response at the end of a frame
	d.Instance.VMCR_FBTAAE().Store(vidCfg.FrameBTAAcknowledgeEnable)

	// Process unlocked
	d.mutex.Unlock()

	return HAL_OK
}

func (d *DSI) ConfigPhyTimer(cfg *PhyTimings) halStatus {
	/* Process locked */
	d.mutex.Lock()

	maxTime := cfg.ClockLaneHS2LPTime
	if cfg.ClockLaneLP2HSTime > cfg.ClockLaneHS2LPTime {
		maxTime = cfg.ClockLaneLP2HSTime
	}

	/* Clock lane timer configuration */

	/* In Automatic Clock Lane control mode, the DSI Host can turn off the clock lane between two
	   High-Speed transmission.
	   To do so, the DSI Host calculates the time required for the clock lane to change from HighSpeed
	   to Low-Power and from Low-Power to High-Speed.
	   This timings are configured by the HS2LP_TIME and LP2HS_TIME in the DSI Host Clock Lane Timer Configuration Register (DSI_CLTCR).
	   But the DSI Host is not calculating LP2HS_TIME + HS2LP_TIME but 2 x HS2LP_TIME.

	   Workaround : Configure HS2LP_TIME and LP2HS_TIME with the same value being the max of HS2LP_TIME or LP2HS_TIME.
	*/
	d.Instance.CLTCR_LP2HS_TIME().Store(dsi.CLTCR(maxTime))
	d.Instance.CLTCR_HS2LP_TIME().Store(dsi.CLTCR(maxTime))

	/* Data lane timer configuration */
	d.Instance.MRD_TIME().Store(dsi.DLTCR(cfg.DataLaneMaxReadTime))
	d.Instance.DLTCR_LP2HS_TIME().Store(dsi.DLTCR(cfg.DataLaneLP2HSTime))
	d.Instance.DLTCR_HS2LP_TIME().Store(dsi.DLTCR(cfg.DataLaneHS2LPTime))

	/* Configure the wait period to request HS transmission after a stop state */
	d.Instance.SW_TIME().Store(dsi.PCONFR(cfg.StopWaitTime))

	d.mutex.Unlock()

	return HAL_OK
}

func VPCR(v uint32) dsi.VPCR {
	return dsi.VPCR(v)
}

func VLCR(v uint32) dsi.VLCR {
	return dsi.VLCR(v)
}

func VHSACR(v uint32) dsi.VHSACR {
	return dsi.VHSACR(v)
}

func VHBPCR(v uint32) dsi.VHBPCR {
	return dsi.VHBPCR(v)
}

func VVACR(v uint32) dsi.VVACR {
	return dsi.VVACR(v)
}

func (p *DSI) ShortWrite(channelID uint32, mode dsi.GHCR, data0, data1 uint8) {
	p.mutex.Lock()

	// wait for command fifo empty
	for p.Instance.CMDFE().Load() == 0 {
		// todo: needs timeout BAD
		delay.Millisec(1)
	}

	p.ConfigPacketHeader(dsi.GHCR(channelID), mode, data0, data1)

	p.mutex.Unlock()
}

func (p *DSI) LongWrite(channelID uint32, mode dsi.GHCR, data []uint8) {
	p.mutex.Lock()

	// wait for command fifo empty
	for p.Instance.CMDFE().Load() == 0 {
		// todo: needs timeout BAD
		delay.Millisec(1)
	}

	var acc uint32 = 0
	cmd := data[len(data)-1] // cmd is the last byte in the array
	acc |= uint32(cmd)       // cmd should be LSB of first packet
	for i, b := range data {
		if i == len(data)-1 { // last byte, cmd byte, needs skipping
			break
		}
		acc |= uint32(b) << ((uint32(i) + 1) % 4) * 8

		if (i+1)%4 == 3 {
			p.Instance.GPDR.Store(dsi.GPDR(acc))
			acc = 0
		}
	}

	pktLen := len(data)

	p.ConfigPacketHeader(dsi.GHCR(channelID), mode, uint8(pktLen&0xFF), uint8((pktLen&0xFF00)>>8))

	p.mutex.Unlock()
}

func (p *DSI) ConfigPacketHeader(channelID dsi.GHCR, mode dsi.GHCR, data0, data1 uint8) {
	header := mode<<dsi.DTn | channelID<<dsi.GHCR_VCIDn
	header |= dsi.GHCR(data1)<<dsi.WCMSBn | dsi.GHCR(data0)<<dsi.WCLSBn
	p.Instance.GHCR.Store(header)
}

func (p *DSI) IRQHandler() {
	// Tearing Effect Interrupt management
	if p.Instance.TEIF().Load() != 0 {
		// Clear the Tearing Effect Interrupt Flag
		p.Instance.TEIF().Clear()

		// Tearing Effect Callback
		p.TearingEffectCallback()
	}

	// End of Refresh Interrupt management
	if p.Instance.ERIF().Load() != 0 {
		// Clear the End of Refresh Interrupt Flag
		p.Instance.ERIF().Clear()

		// End of Refresh Callback
		p.EndOfRefreshCallback()
	}
}
