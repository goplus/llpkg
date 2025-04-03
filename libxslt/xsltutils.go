package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/libxml2"
	"unsafe"
)

/*
 * Our own version of namespaced attributes lookup.
 */
//go:linkname GetNsProp C.xsltGetNsProp
func GetNsProp(node libxml2.NodePtr, name *libxml2.Char, nameSpace *libxml2.Char) *libxml2.Char

//go:linkname GetCNsProp C.xsltGetCNsProp
func GetCNsProp(style StylesheetPtr, node libxml2.NodePtr, name *libxml2.Char, nameSpace *libxml2.Char) *libxml2.Char

//go:linkname GetUTF8Char C.xsltGetUTF8Char
func GetUTF8Char(utf *int8, len *c.Int) c.Int

type DebugTraceCodes c.Int

const (
	TRACEALL            DebugTraceCodes = -1
	TRACENONE           DebugTraceCodes = 0
	TRACECOPYTEXT       DebugTraceCodes = 1
	TRACEPROCESSNODE    DebugTraceCodes = 2
	TRACEAPPLYTEMPLATE  DebugTraceCodes = 4
	TRACECOPY           DebugTraceCodes = 8
	TRACECOMMENT        DebugTraceCodes = 16
	TRACEPI             DebugTraceCodes = 32
	TRACECOPYOF         DebugTraceCodes = 64
	TRACEVALUEOF        DebugTraceCodes = 128
	TRACECALLTEMPLATE   DebugTraceCodes = 256
	TRACEAPPLYTEMPLATES DebugTraceCodes = 512
	TRACECHOOSE         DebugTraceCodes = 1024
	TRACEIF             DebugTraceCodes = 2048
	TRACEFOREACH        DebugTraceCodes = 4096
	TRACESTRIPSPACES    DebugTraceCodes = 8192
	TRACETEMPLATES      DebugTraceCodes = 16384
	TRACEKEYS           DebugTraceCodes = 32768
	TRACEVARIABLES      DebugTraceCodes = 65536
)

// llgo:link DebugTraceCodes.DebugSetDefaultTrace C.xsltDebugSetDefaultTrace
func (recv_ DebugTraceCodes) DebugSetDefaultTrace() {
}

//go:linkname DebugGetDefaultTrace C.xsltDebugGetDefaultTrace
func DebugGetDefaultTrace() DebugTraceCodes

//go:linkname PrintErrorContext C.xsltPrintErrorContext
func PrintErrorContext(ctxt TransformContextPtr, style StylesheetPtr, node libxml2.NodePtr)

//go:linkname Message C.xsltMessage
func Message(ctxt TransformContextPtr, node libxml2.NodePtr, inst libxml2.NodePtr)

//go:linkname SetGenericErrorFunc C.xsltSetGenericErrorFunc
func SetGenericErrorFunc(ctx unsafe.Pointer, handler libxml2.GenericErrorFunc)

//go:linkname SetGenericDebugFunc C.xsltSetGenericDebugFunc
func SetGenericDebugFunc(ctx unsafe.Pointer, handler libxml2.GenericErrorFunc)

//go:linkname SetTransformErrorFunc C.xsltSetTransformErrorFunc
func SetTransformErrorFunc(ctxt TransformContextPtr, ctx unsafe.Pointer, handler libxml2.GenericErrorFunc)

//go:linkname TransformError C.xsltTransformError
func TransformError(ctxt TransformContextPtr, style StylesheetPtr, node libxml2.NodePtr, msg *int8, __llgo_va_list ...interface{})

//go:linkname SetCtxtParseOptions C.xsltSetCtxtParseOptions
func SetCtxtParseOptions(ctxt TransformContextPtr, options c.Int) c.Int

/*
 * Sorting.
 */
//go:linkname DocumentSortFunction C.xsltDocumentSortFunction
func DocumentSortFunction(list libxml2.NodeSetPtr)

//go:linkname SetSortFunc C.xsltSetSortFunc
func SetSortFunc(handler SortFunc)

//go:linkname SetCtxtSortFunc C.xsltSetCtxtSortFunc
func SetCtxtSortFunc(ctxt TransformContextPtr, handler SortFunc)

//go:linkname DefaultSortFunction C.xsltDefaultSortFunction
func DefaultSortFunction(ctxt TransformContextPtr, sorts *libxml2.NodePtr, nbsorts c.Int)

//go:linkname DoSortFunction C.xsltDoSortFunction
func DoSortFunction(ctxt TransformContextPtr, sorts *libxml2.NodePtr, nbsorts c.Int)

//go:linkname ComputeSortResult C.xsltComputeSortResult
func ComputeSortResult(ctxt TransformContextPtr, sort libxml2.NodePtr) *libxml2.XPathObjectPtr

/*
 * QNames handling.
 */
//go:linkname SplitQName C.xsltSplitQName
func SplitQName(dict libxml2.DictPtr, name *libxml2.Char, prefix **libxml2.Char) *libxml2.Char

//go:linkname GetQNameURI C.xsltGetQNameURI
func GetQNameURI(node libxml2.NodePtr, name **libxml2.Char) *libxml2.Char

//go:linkname GetQNameURI2 C.xsltGetQNameURI2
func GetQNameURI2(style StylesheetPtr, node libxml2.NodePtr, name **libxml2.Char) *libxml2.Char

/*
 * Output, reuse libxml I/O buffers.
 */
//go:linkname SaveResultTo C.xsltSaveResultTo
func SaveResultTo(buf libxml2.OutputBufferPtr, result libxml2.DocPtr, style StylesheetPtr) c.Int

//go:linkname SaveResultToFilename C.xsltSaveResultToFilename
func SaveResultToFilename(URI *int8, result libxml2.DocPtr, style StylesheetPtr, compression c.Int) c.Int

//go:linkname SaveResultToFile C.xsltSaveResultToFile
func SaveResultToFile(file *c.FILE, result libxml2.DocPtr, style StylesheetPtr) c.Int

//go:linkname SaveResultToFd C.xsltSaveResultToFd
func SaveResultToFd(fd c.Int, result libxml2.DocPtr, style StylesheetPtr) c.Int

//go:linkname SaveResultToString C.xsltSaveResultToString
func SaveResultToString(doc_txt_ptr **libxml2.Char, doc_txt_len *c.Int, result libxml2.DocPtr, style StylesheetPtr) c.Int

/*
 * XPath interface
 */
//go:linkname XPathCompile C.xsltXPathCompile
func XPathCompile(style StylesheetPtr, str *libxml2.Char) libxml2.XPathCompExprPtr

//go:linkname XPathCompileFlags C.xsltXPathCompileFlags
func XPathCompileFlags(style StylesheetPtr, str *libxml2.Char, flags c.Int) libxml2.XPathCompExprPtr

type DebugStatusCodes c.Int

const (
	DEBUGNONE       DebugStatusCodes = 0
	DEBUGINIT       DebugStatusCodes = 1
	DEBUGSTEP       DebugStatusCodes = 2
	DEBUGSTEPOUT    DebugStatusCodes = 3
	DEBUGNEXT       DebugStatusCodes = 4
	DEBUGSTOP       DebugStatusCodes = 5
	DEBUGCONT       DebugStatusCodes = 6
	DEBUGRUN        DebugStatusCodes = 7
	DEBUGRUNRESTART DebugStatusCodes = 8
	DEBUGQUIT       DebugStatusCodes = 9
)

// llgo:type C
type HandleDebuggerCallback func(libxml2.NodePtr, libxml2.NodePtr, TemplatePtr, TransformContextPtr)

// llgo:type C
type AddCallCallback func(TemplatePtr, libxml2.NodePtr) c.Int

// llgo:type C
type DropCallCallback func()

//go:linkname GetDebuggerStatus C.xsltGetDebuggerStatus
func GetDebuggerStatus() c.Int
