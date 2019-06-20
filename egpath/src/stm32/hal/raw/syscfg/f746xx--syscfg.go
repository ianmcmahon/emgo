// +build f746xx

// Peripheral: SYSCFG_Periph  System configuration controller.
// Instances:
//  SYSCFG  mmap.SYSCFG_BASE
// Registers:
//  0x00 32  MEMRMP    Memory remap register.
//  0x04 32  PMC       Peripheral mode configuration register.
//  0x08 32  EXTICR[4] External interrupt configuration registers.
//  0x20 32  CMPCR     Compensation cell control register.
// Import:
//  stm32/o/f746xx/mmap
package syscfg

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	MEM_BOOT MEMRMP = 0x01 << 0  //+ Boot information after Reset.
	SWP_FMC  MEMRMP = 0x03 << 10 //+ FMC Memory Mapping swapping.
)

const (
	MEM_BOOTn = 0
	SWP_FMCn  = 10
)

const (
	ADCxDC2      PMC = 0x07 << 16 //+ Refer to AN4073 on how to use this bit.
	ADC1DC2      PMC = 0x01 << 16 //  Refer to AN4073 on how to use this bit.
	ADC2DC2      PMC = 0x02 << 16 //  Refer to AN4073 on how to use this bit.
	ADC3DC2      PMC = 0x04 << 16 //  Refer to AN4073 on how to use this bit.
	MII_RMII_SEL PMC = 0x01 << 23 //+ Ethernet PHY interface selection.
)

const (
	ADCxDC2n      = 16
	MII_RMII_SELn = 23
)

const (
	EXTI0    EXTICR = 0x0F << 0  //+ EXTI 0 configuration.
	EXTI1    EXTICR = 0x0F << 4  //+ EXTI 1 configuration.
	EXTI2    EXTICR = 0x0F << 8  //+ EXTI 2 configuration.
	EXTI3    EXTICR = 0x0F << 12 //+ EXTI 3 configuration.
	EXTI0_PA EXTICR = 0x00 << 12 //  PA[0] pin.
	EXTI0_PB EXTICR = 0x01 << 0  //  PB[0] pin.
	EXTI0_PC EXTICR = 0x02 << 0  //  PC[0] pin.
	EXTI0_PD EXTICR = 0x03 << 0  //  PD[0] pin.
	EXTI0_PE EXTICR = 0x04 << 0  //  PE[0] pin.
	EXTI0_PF EXTICR = 0x05 << 0  //  PF[0] pin.
	EXTI0_PG EXTICR = 0x06 << 0  //  PG[0] pin.
	EXTI0_PH EXTICR = 0x07 << 0  //  PH[0] pin.
	EXTI0_PI EXTICR = 0x08 << 0  //  PI[0] pin.
	EXTI0_PJ EXTICR = 0x09 << 0  //  PJ[0] pin.
	EXTI0_PK EXTICR = 0x0A << 0  //  PK[0] pin.
	EXTI1_PA EXTICR = 0x00 << 12 //  PA[1] pin.
	EXTI1_PB EXTICR = 0x01 << 4  //  PB[1] pin.
	EXTI1_PC EXTICR = 0x02 << 4  //  PC[1] pin.
	EXTI1_PD EXTICR = 0x03 << 4  //  PD[1] pin.
	EXTI1_PE EXTICR = 0x04 << 4  //  PE[1] pin.
	EXTI1_PF EXTICR = 0x05 << 4  //  PF[1] pin.
	EXTI1_PG EXTICR = 0x06 << 4  //  PG[1] pin.
	EXTI1_PH EXTICR = 0x07 << 4  //  PH[1] pin.
	EXTI1_PI EXTICR = 0x08 << 4  //  PI[1] pin.
	EXTI1_PJ EXTICR = 0x09 << 4  //  PJ[1] pin.
	EXTI1_PK EXTICR = 0x0A << 4  //  PK[1] pin.
	EXTI2_PA EXTICR = 0x00 << 12 //  PA[2] pin.
	EXTI2_PB EXTICR = 0x01 << 8  //  PB[2] pin.
	EXTI2_PC EXTICR = 0x02 << 8  //  PC[2] pin.
	EXTI2_PD EXTICR = 0x03 << 8  //  PD[2] pin.
	EXTI2_PE EXTICR = 0x04 << 8  //  PE[2] pin.
	EXTI2_PF EXTICR = 0x05 << 8  //  PF[2] pin.
	EXTI2_PG EXTICR = 0x06 << 8  //  PG[2] pin.
	EXTI2_PH EXTICR = 0x07 << 8  //  PH[2] pin.
	EXTI2_PI EXTICR = 0x08 << 8  //  PI[2] pin.
	EXTI2_PJ EXTICR = 0x09 << 8  //  PJ[2] pin.
	EXTI2_PK EXTICR = 0x0A << 8  //  PK[2] pin.
	EXTI3_PA EXTICR = 0x00 << 12 //  PA[3] pin.
	EXTI3_PB EXTICR = 0x01 << 12 //  PB[3] pin.
	EXTI3_PC EXTICR = 0x02 << 12 //  PC[3] pin.
	EXTI3_PD EXTICR = 0x03 << 12 //  PD[3] pin.
	EXTI3_PE EXTICR = 0x04 << 12 //  PE[3] pin.
	EXTI3_PF EXTICR = 0x05 << 12 //  PF[3] pin.
	EXTI3_PG EXTICR = 0x06 << 12 //  PG[3] pin.
	EXTI3_PH EXTICR = 0x07 << 12 //  PH[3] pin.
	EXTI3_PI EXTICR = 0x08 << 12 //  PI[3] pin.
	EXTI3_PJ EXTICR = 0x09 << 12 //  PJ[3] pin.
	EXTI3_PK EXTICR = 0x0A << 12 //  PK[3] pin.
)

const (
	EXTI0n = 0
	EXTI1n = 4
	EXTI2n = 8
	EXTI3n = 12
)

const (
	CMP_PD CMPCR = 0x01 << 0 //+ Compensation cell power-down.
	READY  CMPCR = 0x01 << 8 //+ Compensation cell ready flag.
)

const (
	CMP_PDn = 0
	READYn  = 8
)
