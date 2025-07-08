package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const DFU_SUBCLASS = 0x01
const DFU_RT_PROTOCOL = 0x01
const DFU_MODE_PROTOCOL = 0x02
const DFU_DETACH = 0x00
const DFU_DNLOAD = 0x01
const DFU_UPLOAD = 0x02
const DFU_GETSTATUS = 0x03
const DFU_CLRSTATUS = 0x04
const DFU_GETSTATE = 0x05
const DFU_ABORT = 0x06
const DFU_FUNC_DESC = 0x21
const DFU_ATTR_WILL_DETACH = 0x08
const DFU_ATTR_MANIFESTATION_TOLERANT = 0x04
const DFU_ATTR_CAN_UPLOAD = 0x02
const DFU_ATTR_CAN_DNLOAD = 0x01
const DFU_VERSION = 0x0110

/** Run-Time Functional Descriptor */

type DfuRuntimeDescriptor struct {
	BLength         c.Uint8T
	BDescriptorType c.Uint8T
	BmAttributes    c.Uint8T
	WDetachTimeOut  c.Uint16T
	WTransferSize   c.Uint16T
	BcdDFUVersion   c.Uint16T
}
type DfuStatus c.Int

const (
	StatusOK        DfuStatus = 0
	ErrTARGET       DfuStatus = 1
	ErrFILE         DfuStatus = 2
	ErrWRITE        DfuStatus = 3
	ErrERASE        DfuStatus = 4
	ErrCHECK_ERASED DfuStatus = 5
	ErrPROG         DfuStatus = 6
	ErrVERIFY       DfuStatus = 7
	ErrADDRESS      DfuStatus = 8
	ErrNOTDONE      DfuStatus = 9
	ErrFIRMWARE     DfuStatus = 10
	ErrVENDOR       DfuStatus = 11
	ErrUSB          DfuStatus = 12
	ErrPOR          DfuStatus = 13
	ErrUNKNOWN      DfuStatus = 14
	ErrSTALLEDPKT   DfuStatus = 15
)

type DfuState c.Int

const (
	AppIDLE              DfuState = 0
	AppDETACH            DfuState = 1
	DfuIDLE              DfuState = 2
	DfuDNLOAD_SYNC       DfuState = 3
	DfuDNBUSY            DfuState = 4
	DfuDNLOAD_IDLE       DfuState = 5
	DfuMANIFEST_SYNC     DfuState = 6
	DfuMANIFEST          DfuState = 7
	DfuMANIFEST_WAIT_RST DfuState = 8
	DfuUPLOAD_IDLE       DfuState = 9
	DfuERROR             DfuState = 10
)

/*
 These callbacks are made public so the ACM driver can call them to handle the switch to DFU.
*/
// llgo:link (*UsbSetupPacket).DfuClassHandleReq C.dfu_class_handle_req
func (recv_ *UsbSetupPacket) DfuClassHandleReq(data_len *c.Int32T, data **c.Uint8T) c.Int {
	return 0
}

// llgo:link UsbDcStatusCode.DfuStatusCb C.dfu_status_cb
func (recv_ UsbDcStatusCode) DfuStatusCb(param *c.Uint8T) {
}

//go:linkname UsbDfuInit C.usb_dfu_init
func UsbDfuInit() c.Int

// llgo:link (*UsbSetupPacket).DfuCustomHandleReq C.dfu_custom_handle_req
func (recv_ *UsbSetupPacket) DfuCustomHandleReq(data_len *c.Int32T, data **c.Uint8T) c.Int {
	return 0
}

// llgo:type C
type UsbDfuDetachRoutineT func(c.Int)

//go:linkname UsbDfuSetDetachCb C.usb_dfu_set_detach_cb
func UsbDfuSetDetachCb(cb UsbDfuDetachRoutineT)

//go:linkname UsbDfuForceDetach C.usb_dfu_force_detach
func UsbDfuForceDetach()
