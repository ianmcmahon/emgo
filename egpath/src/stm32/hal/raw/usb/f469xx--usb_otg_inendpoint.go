// +build f469xx

// Peripheral: USB_OTG_INEndpoint_Periph  USB_OTG_IN_Endpoint-Specific_Register.
// Instances:
// Registers:
//  0x00 32  DIEPCTL  dev IN Endpoint Control Reg    900h + (ep_num * 20h) + 00h.
//  0x08 32  DIEPINT  dev IN Endpoint Itr Reg        900h + (ep_num * 20h) + 08h.
//  0x10 32  DIEPTSIZ IN Endpoint Txfer Size         900h + (ep_num * 20h) + 10h.
//  0x14 32  DIEPDMA  IN Endpoint DMA Address Reg    900h + (ep_num * 20h) + 14h.
//  0x18 32  DTXFSTS  IN Endpoint Tx FIFO Status Reg 900h + (ep_num * 20h) + 18h.
// Import:
//  stm32/o/f469xx/mmap
package usb

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	MPSIZ          DIEPCTL = 0x7FF << 0 //+ Maximum packet size.
	USBAEP         DIEPCTL = 0x01 << 15 //+ USB active endpoint.
	EONUM_DPID     DIEPCTL = 0x01 << 16 //+ Even/odd frame.
	NAKSTS         DIEPCTL = 0x01 << 17 //+ NAK status.
	EPTYP          DIEPCTL = 0x03 << 18 //+ Endpoint type.
	STALL          DIEPCTL = 0x01 << 21 //+ STALL handshake.
	TXFNUM         DIEPCTL = 0x0F << 22 //+ TxFIFO number.
	CNAK           DIEPCTL = 0x01 << 26 //+ Clear NAK.
	SNAK           DIEPCTL = 0x01 << 27 //+ Set NAK.
	SD0PID_SEVNFRM DIEPCTL = 0x01 << 28 //+ Set DATA0 PID.
	SODDFRM        DIEPCTL = 0x01 << 29 //+ Set odd frame.
	EPDIS          DIEPCTL = 0x01 << 30 //+ Endpoint disable.
	EPENA          DIEPCTL = 0x01 << 31 //+ Endpoint enable.
)

const (
	MPSIZn          = 0
	USBAEPn         = 15
	EONUM_DPIDn     = 16
	NAKSTSn         = 17
	EPTYPn          = 18
	STALLn          = 21
	TXFNUMn         = 22
	CNAKn           = 26
	SNAKn           = 27
	SD0PID_SEVNFRMn = 28
	SODDFRMn        = 29
	EPDISn          = 30
	EPENAn          = 31
)

const (
	XFRC       DIEPINT = 0x01 << 0  //+ Transfer completed interrupt.
	EPDISD     DIEPINT = 0x01 << 1  //+ Endpoint disabled interrupt.
	TOC        DIEPINT = 0x01 << 3  //+ Timeout condition.
	ITTXFE     DIEPINT = 0x01 << 4  //+ IN token received when TxFIFO is empty.
	INEPNE     DIEPINT = 0x01 << 6  //+ IN endpoint NAK effective.
	TXFE       DIEPINT = 0x01 << 7  //+ Transmit FIFO empty.
	TXFIFOUDRN DIEPINT = 0x01 << 8  //+ Transmit Fifo Underrun.
	BNA        DIEPINT = 0x01 << 9  //+ Buffer not available interrupt.
	PKTDRPSTS  DIEPINT = 0x01 << 11 //+ Packet dropped status.
	BERR       DIEPINT = 0x01 << 12 //+ Babble error interrupt.
	NAK        DIEPINT = 0x01 << 13 //+ NAK interrupt.
)

const (
	XFRCn       = 0
	EPDISDn     = 1
	TOCn        = 3
	ITTXFEn     = 4
	INEPNEn     = 6
	TXFEn       = 7
	TXFIFOUDRNn = 8
	BNAn        = 9
	PKTDRPSTSn  = 11
	BERRn       = 12
	NAKn        = 13
)

const (
	XFRSIZ DIEPTSIZ = 0x7FFFF << 0 //+ Transfer size.
	PKTCNT DIEPTSIZ = 0x3FF << 19  //+ Packet count.
	MULCNT DIEPTSIZ = 0x03 << 29   //+ Packet count.
)

const (
	XFRSIZn = 0
	PKTCNTn = 19
	MULCNTn = 29
)

const (
	DMAADDR DIEPDMA = 0xFFFFFFFF << 0 //+ DMA address.
)

const (
	DMAADDRn = 0
)

const (
	INEPTFSAV DTXFSTS = 0xFFFF << 0 //+ IN endpoint TxFIFO space available.
)

const (
	INEPTFSAVn = 0
)
