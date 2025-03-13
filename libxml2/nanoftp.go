package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

// llgo:type C
type FtpListCallback func(unsafe.Pointer, *int8, *int8, *int8, *int8, c.Ulong, c.Int, c.Int, *int8, c.Int, c.Int, c.Int)

// llgo:type C
type FtpDataCallback func(unsafe.Pointer, *int8, c.Int)

/*
 * Init
 */
//go:linkname NanoFTPInit C.xmlNanoFTPInit
func NanoFTPInit()

//go:linkname NanoFTPCleanup C.xmlNanoFTPCleanup
func NanoFTPCleanup()

/*
 * Creating/freeing contexts.
 */
//go:linkname NanoFTPNewCtxt C.xmlNanoFTPNewCtxt
func NanoFTPNewCtxt(URL *int8) unsafe.Pointer

//go:linkname NanoFTPFreeCtxt C.xmlNanoFTPFreeCtxt
func NanoFTPFreeCtxt(ctx unsafe.Pointer)

//go:linkname NanoFTPConnectTo C.xmlNanoFTPConnectTo
func NanoFTPConnectTo(server *int8, port c.Int) unsafe.Pointer

/*
 * Opening/closing session connections.
 */
//go:linkname NanoFTPOpen C.xmlNanoFTPOpen
func NanoFTPOpen(URL *int8) unsafe.Pointer

//go:linkname NanoFTPConnect C.xmlNanoFTPConnect
func NanoFTPConnect(ctx unsafe.Pointer) c.Int

//go:linkname NanoFTPClose C.xmlNanoFTPClose
func NanoFTPClose(ctx unsafe.Pointer) c.Int

//go:linkname NanoFTPQuit C.xmlNanoFTPQuit
func NanoFTPQuit(ctx unsafe.Pointer) c.Int

//go:linkname NanoFTPScanProxy C.xmlNanoFTPScanProxy
func NanoFTPScanProxy(URL *int8)

//go:linkname NanoFTPProxy C.xmlNanoFTPProxy
func NanoFTPProxy(host *int8, port c.Int, user *int8, passwd *int8, type_ c.Int)

//go:linkname NanoFTPUpdateURL C.xmlNanoFTPUpdateURL
func NanoFTPUpdateURL(ctx unsafe.Pointer, URL *int8) c.Int

/*
 * Rather internal commands.
 */
//go:linkname NanoFTPGetResponse C.xmlNanoFTPGetResponse
func NanoFTPGetResponse(ctx unsafe.Pointer) c.Int

//go:linkname NanoFTPCheckResponse C.xmlNanoFTPCheckResponse
func NanoFTPCheckResponse(ctx unsafe.Pointer) c.Int

/*
 * CD/DIR/GET handlers.
 */
//go:linkname NanoFTPCwd C.xmlNanoFTPCwd
func NanoFTPCwd(ctx unsafe.Pointer, directory *int8) c.Int

//go:linkname NanoFTPDele C.xmlNanoFTPDele
func NanoFTPDele(ctx unsafe.Pointer, file *int8) c.Int

//go:linkname NanoFTPGetConnection C.xmlNanoFTPGetConnection
func NanoFTPGetConnection(ctx unsafe.Pointer) c.Int

//go:linkname NanoFTPCloseConnection C.xmlNanoFTPCloseConnection
func NanoFTPCloseConnection(ctx unsafe.Pointer) c.Int

//go:linkname NanoFTPList C.xmlNanoFTPList
func NanoFTPList(ctx unsafe.Pointer, callback FtpListCallback, userData unsafe.Pointer, filename *int8) c.Int

//go:linkname NanoFTPGetSocket C.xmlNanoFTPGetSocket
func NanoFTPGetSocket(ctx unsafe.Pointer, filename *int8) c.Int

//go:linkname NanoFTPGet C.xmlNanoFTPGet
func NanoFTPGet(ctx unsafe.Pointer, callback FtpDataCallback, userData unsafe.Pointer, filename *int8) c.Int

//go:linkname NanoFTPRead C.xmlNanoFTPRead
func NanoFTPRead(ctx unsafe.Pointer, dest unsafe.Pointer, len c.Int) c.Int
