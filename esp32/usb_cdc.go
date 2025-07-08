package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const CDC_SRN_1_20 = 0x0120
const ACM_SUBCLASS = 0x02
const ECM_SUBCLASS = 0x06
const EEM_SUBCLASS = 0x0c
const AT_CMD_V250_PROTOCOL = 0x01
const EEM_PROTOCOL = 0x07
const DATA_INTERFACE_CLASS = 0x0A
const CS_INTERFACE = 0x24
const CS_ENDPOINT = 0x25
const HEADER_FUNC_DESC = 0x00
const CALL_MANAGEMENT_FUNC_DESC = 0x01
const ACM_FUNC_DESC = 0x02
const UNION_FUNC_DESC = 0x06
const ETHERNET_FUNC_DESC = 0x0F
const CDC_SEND_ENC_CMD = 0x00
const CDC_GET_ENC_RSP = 0x01
const SET_LINE_CODING = 0x20
const GET_LINE_CODING = 0x21
const SET_CONTROL_LINE_STATE = 0x22
const SET_CONTROL_LINE_STATE_RTS = 0x02
const SET_CONTROL_LINE_STATE_DTR = 0x01
const SERIAL_STATE_OVERRUN = 0x40
const SERIAL_STATE_PARITY = 0x20
const SERIAL_STATE_FRAMING = 0x10
const SERIAL_STATE_RING = 0x08
const SERIAL_STATE_BREAK = 0x04
const SERIAL_STATE_TX_CARRIER = 0x02
const SERIAL_STATE_RX_CARRIER = 0x01
const SET_ETHERNET_MULTICAST_FILTERS = 0x40
const SET_ETHERNET_PM_FILTER = 0x41
const GET_ETHERNET_PM_FILTER = 0x42
const SET_ETHERNET_PACKET_FILTER = 0x43
const GET_ETHERNET_STATISTIC = 0x44
const PACKET_TYPE_MULTICAST = 0x10
const PACKET_TYPE_BROADCAST = 0x08
const PACKET_TYPE_DIRECTED = 0x04
const PACKET_TYPE_ALL_MULTICAST = 0x02
const PACKET_TYPE_PROMISCUOUS = 0x01

/** Header Functional Descriptor */

type CdcHeaderDescriptor struct {
	BFunctionLength    c.Uint8T
	BDescriptorType    c.Uint8T
	BDescriptorSubtype c.Uint8T
	BcdCDC             c.Uint16T
}

/** Union Interface Functional Descriptor */

type CdcUnionDescriptor struct {
	BFunctionLength        c.Uint8T
	BDescriptorType        c.Uint8T
	BDescriptorSubtype     c.Uint8T
	BControlInterface      c.Uint8T
	BSubordinateInterface0 c.Uint8T
}

/** Call Management Functional Descriptor */

type CdcCmDescriptor struct {
	BFunctionLength    c.Uint8T
	BDescriptorType    c.Uint8T
	BDescriptorSubtype c.Uint8T
	BmCapabilities     c.Uint8T
	BDataInterface     c.Uint8T
}

/** Abstract Control Management Functional Descriptor */

type CdcAcmDescriptor struct {
	BFunctionLength    c.Uint8T
	BDescriptorType    c.Uint8T
	BDescriptorSubtype c.Uint8T
	BmCapabilities     c.Uint8T
}

/** Data structure for GET_LINE_CODING / SET_LINE_CODING class requests */

type CdcAcmLineCoding struct {
	DwDTERate   c.Uint32T
	BCharFormat c.Uint8T
	BParityType c.Uint8T
	BDataBits   c.Uint8T
}

/** Data structure for the notification about SerialState */

type CdcAcmNotification struct {
	BmRequestType     c.Uint8T
	BNotificationType c.Uint8T
	WValue            c.Uint16T
	WIndex            c.Uint16T
	WLength           c.Uint16T
	Data              c.Uint16T
}

/** Ethernet Networking Functional Descriptor */

type CdcEcmDescriptor struct {
	BFunctionLength      c.Uint8T
	BDescriptorType      c.Uint8T
	BDescriptorSubtype   c.Uint8T
	IMACAddress          c.Uint8T
	BmEthernetStatistics c.Uint32T
	WMaxSegmentSize      c.Uint16T
	WNumberMCFilters     c.Uint16T
	BNumberPowerFilters  c.Uint8T
}
