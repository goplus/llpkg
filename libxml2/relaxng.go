package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type X_xmlRelaxNG struct {
	Unused [8]uint8
}
type RelaxNG X_xmlRelaxNG
type RelaxNGPtr *RelaxNG

// llgo:type C
type RelaxNGValidityErrorFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})

// llgo:type C
type RelaxNGValidityWarningFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})

type X_xmlRelaxNGParserCtxt struct {
	Unused [8]uint8
}
type RelaxNGParserCtxt X_xmlRelaxNGParserCtxt
type RelaxNGParserCtxtPtr *RelaxNGParserCtxt

type X_xmlRelaxNGValidCtxt struct {
	Unused [8]uint8
}
type RelaxNGValidCtxt X_xmlRelaxNGValidCtxt
type RelaxNGValidCtxtPtr *RelaxNGValidCtxt
type RelaxNGValidErr c.Int

const (
	RELAXNGOK              RelaxNGValidErr = 0
	RELAXNGERRMEMORY       RelaxNGValidErr = 1
	RELAXNGERRTYPE         RelaxNGValidErr = 2
	RELAXNGERRTYPEVAL      RelaxNGValidErr = 3
	RELAXNGERRDUPID        RelaxNGValidErr = 4
	RELAXNGERRTYPECMP      RelaxNGValidErr = 5
	RELAXNGERRNOSTATE      RelaxNGValidErr = 6
	RELAXNGERRNODEFINE     RelaxNGValidErr = 7
	RELAXNGERRLISTEXTRA    RelaxNGValidErr = 8
	RELAXNGERRLISTEMPTY    RelaxNGValidErr = 9
	RELAXNGERRINTERNODATA  RelaxNGValidErr = 10
	RELAXNGERRINTERSEQ     RelaxNGValidErr = 11
	RELAXNGERRINTEREXTRA   RelaxNGValidErr = 12
	RELAXNGERRELEMNAME     RelaxNGValidErr = 13
	RELAXNGERRATTRNAME     RelaxNGValidErr = 14
	RELAXNGERRELEMNONS     RelaxNGValidErr = 15
	RELAXNGERRATTRNONS     RelaxNGValidErr = 16
	RELAXNGERRELEMWRONGNS  RelaxNGValidErr = 17
	RELAXNGERRATTRWRONGNS  RelaxNGValidErr = 18
	RELAXNGERRELEMEXTRANS  RelaxNGValidErr = 19
	RELAXNGERRATTREXTRANS  RelaxNGValidErr = 20
	RELAXNGERRELEMNOTEMPTY RelaxNGValidErr = 21
	RELAXNGERRNOELEM       RelaxNGValidErr = 22
	RELAXNGERRNOTELEM      RelaxNGValidErr = 23
	RELAXNGERRATTRVALID    RelaxNGValidErr = 24
	RELAXNGERRCONTENTVALID RelaxNGValidErr = 25
	RELAXNGERREXTRACONTENT RelaxNGValidErr = 26
	RELAXNGERRINVALIDATTR  RelaxNGValidErr = 27
	RELAXNGERRDATAELEM     RelaxNGValidErr = 28
	RELAXNGERRVALELEM      RelaxNGValidErr = 29
	RELAXNGERRLISTELEM     RelaxNGValidErr = 30
	RELAXNGERRDATATYPE     RelaxNGValidErr = 31
	RELAXNGERRVALUE        RelaxNGValidErr = 32
	RELAXNGERRLIST         RelaxNGValidErr = 33
	RELAXNGERRNOGRAMMAR    RelaxNGValidErr = 34
	RELAXNGERREXTRADATA    RelaxNGValidErr = 35
	RELAXNGERRLACKDATA     RelaxNGValidErr = 36
	RELAXNGERRINTERNAL     RelaxNGValidErr = 37
	RELAXNGERRELEMWRONG    RelaxNGValidErr = 38
	RELAXNGERRTEXTWRONG    RelaxNGValidErr = 39
)

type RelaxNGParserFlag c.Int

const (
	RELAXNGPNONE    RelaxNGParserFlag = 0
	RELAXNGPFREEDOC RelaxNGParserFlag = 1
	RELAXNGPCRNG    RelaxNGParserFlag = 2
)

//go:linkname RelaxNGInitTypes C.xmlRelaxNGInitTypes
func RelaxNGInitTypes() c.Int

//go:linkname RelaxNGCleanupTypes C.xmlRelaxNGCleanupTypes
func RelaxNGCleanupTypes()

/*
 * Interfaces for parsing.
 */
//go:linkname RelaxNGNewParserCtxt C.xmlRelaxNGNewParserCtxt
func RelaxNGNewParserCtxt(URL *int8) RelaxNGParserCtxtPtr

//go:linkname RelaxNGNewMemParserCtxt C.xmlRelaxNGNewMemParserCtxt
func RelaxNGNewMemParserCtxt(buffer *int8, size c.Int) RelaxNGParserCtxtPtr

//go:linkname RelaxNGNewDocParserCtxt C.xmlRelaxNGNewDocParserCtxt
func RelaxNGNewDocParserCtxt(doc DocPtr) RelaxNGParserCtxtPtr

//go:linkname RelaxParserSetFlag C.xmlRelaxParserSetFlag
func RelaxParserSetFlag(ctxt RelaxNGParserCtxtPtr, flag c.Int) c.Int

//go:linkname RelaxNGFreeParserCtxt C.xmlRelaxNGFreeParserCtxt
func RelaxNGFreeParserCtxt(ctxt RelaxNGParserCtxtPtr)

//go:linkname RelaxNGSetParserErrors C.xmlRelaxNGSetParserErrors
func RelaxNGSetParserErrors(ctxt RelaxNGParserCtxtPtr, err RelaxNGValidityErrorFunc, warn RelaxNGValidityWarningFunc, ctx unsafe.Pointer)

//go:linkname RelaxNGGetParserErrors C.xmlRelaxNGGetParserErrors
func RelaxNGGetParserErrors(ctxt RelaxNGParserCtxtPtr, err RelaxNGValidityErrorFunc, warn RelaxNGValidityWarningFunc, ctx *unsafe.Pointer) c.Int

//go:linkname RelaxNGSetParserStructuredErrors C.xmlRelaxNGSetParserStructuredErrors
func RelaxNGSetParserStructuredErrors(ctxt RelaxNGParserCtxtPtr, serror StructuredErrorFunc, ctx unsafe.Pointer)

//go:linkname RelaxNGParse C.xmlRelaxNGParse
func RelaxNGParse(ctxt RelaxNGParserCtxtPtr) RelaxNGPtr

//go:linkname RelaxNGFree C.xmlRelaxNGFree
func RelaxNGFree(schema RelaxNGPtr)

//go:linkname RelaxNGDump C.xmlRelaxNGDump
func RelaxNGDump(output *c.FILE, schema RelaxNGPtr)

//go:linkname RelaxNGDumpTree C.xmlRelaxNGDumpTree
func RelaxNGDumpTree(output *c.FILE, schema RelaxNGPtr)

/*
 * Interfaces for validating
 */
//go:linkname RelaxNGSetValidErrors C.xmlRelaxNGSetValidErrors
func RelaxNGSetValidErrors(ctxt RelaxNGValidCtxtPtr, err RelaxNGValidityErrorFunc, warn RelaxNGValidityWarningFunc, ctx unsafe.Pointer)

//go:linkname RelaxNGGetValidErrors C.xmlRelaxNGGetValidErrors
func RelaxNGGetValidErrors(ctxt RelaxNGValidCtxtPtr, err RelaxNGValidityErrorFunc, warn RelaxNGValidityWarningFunc, ctx *unsafe.Pointer) c.Int

//go:linkname RelaxNGSetValidStructuredErrors C.xmlRelaxNGSetValidStructuredErrors
func RelaxNGSetValidStructuredErrors(ctxt RelaxNGValidCtxtPtr, serror StructuredErrorFunc, ctx unsafe.Pointer)

//go:linkname RelaxNGNewValidCtxt C.xmlRelaxNGNewValidCtxt
func RelaxNGNewValidCtxt(schema RelaxNGPtr) RelaxNGValidCtxtPtr

//go:linkname RelaxNGFreeValidCtxt C.xmlRelaxNGFreeValidCtxt
func RelaxNGFreeValidCtxt(ctxt RelaxNGValidCtxtPtr)

//go:linkname RelaxNGValidateDoc C.xmlRelaxNGValidateDoc
func RelaxNGValidateDoc(ctxt RelaxNGValidCtxtPtr, doc DocPtr) c.Int

/*
 * Interfaces for progressive validation when possible
 */
//go:linkname RelaxNGValidatePushElement C.xmlRelaxNGValidatePushElement
func RelaxNGValidatePushElement(ctxt RelaxNGValidCtxtPtr, doc DocPtr, elem NodePtr) c.Int

//go:linkname RelaxNGValidatePushCData C.xmlRelaxNGValidatePushCData
func RelaxNGValidatePushCData(ctxt RelaxNGValidCtxtPtr, data *Char, len c.Int) c.Int

//go:linkname RelaxNGValidatePopElement C.xmlRelaxNGValidatePopElement
func RelaxNGValidatePopElement(ctxt RelaxNGValidCtxtPtr, doc DocPtr, elem NodePtr) c.Int

//go:linkname RelaxNGValidateFullElement C.xmlRelaxNGValidateFullElement
func RelaxNGValidateFullElement(ctxt RelaxNGValidCtxtPtr, doc DocPtr, elem NodePtr) c.Int
