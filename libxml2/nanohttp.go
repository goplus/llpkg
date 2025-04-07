package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

//go:linkname NanoHTTPInit C.xmlNanoHTTPInit
func NanoHTTPInit()

//go:linkname NanoHTTPCleanup C.xmlNanoHTTPCleanup
func NanoHTTPCleanup()

//go:linkname NanoHTTPScanProxy C.xmlNanoHTTPScanProxy
func NanoHTTPScanProxy(URL *int8)

//go:linkname NanoHTTPFetch C.xmlNanoHTTPFetch
func NanoHTTPFetch(URL *int8, filename *int8, contentType **int8) c.Int

//go:linkname NanoHTTPMethod C.xmlNanoHTTPMethod
func NanoHTTPMethod(URL *int8, method *int8, input *int8, contentType **int8, headers *int8, ilen c.Int) unsafe.Pointer

//go:linkname NanoHTTPMethodRedir C.xmlNanoHTTPMethodRedir
func NanoHTTPMethodRedir(URL *int8, method *int8, input *int8, contentType **int8, redir **int8, headers *int8, ilen c.Int) unsafe.Pointer

//go:linkname NanoHTTPOpen C.xmlNanoHTTPOpen
func NanoHTTPOpen(URL *int8, contentType **int8) unsafe.Pointer

//go:linkname NanoHTTPOpenRedir C.xmlNanoHTTPOpenRedir
func NanoHTTPOpenRedir(URL *int8, contentType **int8, redir **int8) unsafe.Pointer

//go:linkname NanoHTTPReturnCode C.xmlNanoHTTPReturnCode
func NanoHTTPReturnCode(ctx unsafe.Pointer) c.Int

//go:linkname NanoHTTPAuthHeader C.xmlNanoHTTPAuthHeader
func NanoHTTPAuthHeader(ctx unsafe.Pointer) *int8

//go:linkname NanoHTTPRedir C.xmlNanoHTTPRedir
func NanoHTTPRedir(ctx unsafe.Pointer) *int8

//go:linkname NanoHTTPContentLength C.xmlNanoHTTPContentLength
func NanoHTTPContentLength(ctx unsafe.Pointer) c.Int

//go:linkname NanoHTTPEncoding C.xmlNanoHTTPEncoding
func NanoHTTPEncoding(ctx unsafe.Pointer) *int8

//go:linkname NanoHTTPMimeType C.xmlNanoHTTPMimeType
func NanoHTTPMimeType(ctx unsafe.Pointer) *int8

//go:linkname NanoHTTPRead C.xmlNanoHTTPRead
func NanoHTTPRead(ctx unsafe.Pointer, dest unsafe.Pointer, len c.Int) c.Int

//go:linkname NanoHTTPSave C.xmlNanoHTTPSave
func NanoHTTPSave(ctxt unsafe.Pointer, filename *int8) c.Int

//go:linkname NanoHTTPClose C.xmlNanoHTTPClose
func NanoHTTPClose(ctx unsafe.Pointer)
