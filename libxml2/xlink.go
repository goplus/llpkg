package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type XlinkHRef *Char
type XlinkRole *Char
type XlinkTitle *Char
type XlinkType c.Int

const (
	XLINKTYPENONE        XlinkType = 0
	XLINKTYPESIMPLE      XlinkType = 1
	XLINKTYPEEXTENDED    XlinkType = 2
	XLINKTYPEEXTENDEDSET XlinkType = 3
)

type XlinkShow c.Int

const (
	XLINKSHOWNONE    XlinkShow = 0
	XLINKSHOWNEW     XlinkShow = 1
	XLINKSHOWEMBED   XlinkShow = 2
	XLINKSHOWREPLACE XlinkShow = 3
)

type XlinkActuate c.Int

const (
	XLINKACTUATENONE      XlinkActuate = 0
	XLINKACTUATEAUTO      XlinkActuate = 1
	XLINKACTUATEONREQUEST XlinkActuate = 2
)

// llgo:type C
type XlinkNodeDetectFunc func(unsafe.Pointer, NodePtr)

// llgo:type C
type XlinkSimpleLinkFunk func(unsafe.Pointer, NodePtr, XlinkHRef, XlinkRole, XlinkTitle)

// llgo:type C
type XlinkExtendedLinkFunk func(unsafe.Pointer, NodePtr, c.Int, *XlinkHRef, *XlinkRole, c.Int, *XlinkRole, *XlinkRole, *XlinkShow, *XlinkActuate, c.Int, *XlinkTitle, **Char)

// llgo:type C
type XlinkExtendedLinkSetFunk func(unsafe.Pointer, NodePtr, c.Int, *XlinkHRef, *XlinkRole, c.Int, *XlinkTitle, **Char)

type X_xlinkHandler struct {
	Simple   unsafe.Pointer
	Extended unsafe.Pointer
	Set      unsafe.Pointer
}
type XlinkHandler X_xlinkHandler
type XlinkHandlerPtr *XlinkHandler

/*
 * The default detection routine, can be overridden, they call the default
 * detection callbacks.
 */
//go:linkname XlinkGetDefaultDetect C.xlinkGetDefaultDetect
func XlinkGetDefaultDetect() XlinkNodeDetectFunc

//go:linkname XlinkSetDefaultDetect C.xlinkSetDefaultDetect
func XlinkSetDefaultDetect(func_ XlinkNodeDetectFunc)

/*
 * Routines to set/get the default handlers.
 */
//go:linkname XlinkGetDefaultHandler C.xlinkGetDefaultHandler
func XlinkGetDefaultHandler() XlinkHandlerPtr

//go:linkname XlinkSetDefaultHandler C.xlinkSetDefaultHandler
func XlinkSetDefaultHandler(handler XlinkHandlerPtr)

/*
 * Link detection module itself.
 */
//go:linkname XlinkIsLink C.xlinkIsLink
func XlinkIsLink(doc DocPtr, node NodePtr) XlinkType
