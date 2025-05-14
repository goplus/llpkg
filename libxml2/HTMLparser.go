package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type HtmlParserCtxt ParserCtxt
type HtmlParserCtxtPtr ParserCtxtPtr
type HtmlParserNodeInfo ParserNodeInfo
type HtmlSAXHandler SAXHandler
type HtmlSAXHandlerPtr SAXHandlerPtr
type HtmlParserInput ParserInput
type HtmlParserInputPtr ParserInputPtr
type HtmlDocPtr DocPtr
type HtmlNodePtr NodePtr

type X_htmlElemDesc struct {
	Name          *int8
	StartTag      int8
	EndTag        int8
	SaveEndTag    int8
	Empty         int8
	Depr          int8
	Dtd           int8
	Isinline      int8
	Desc          *int8
	Subelts       **int8
	Defaultsubelt *int8
	AttrsOpt      **int8
	AttrsDepr     **int8
	AttrsReq      **int8
}
type HtmlElemDesc X_htmlElemDesc
type HtmlElemDescPtr *HtmlElemDesc

type X_htmlEntityDesc struct {
	Value c.Uint
	Name  *int8
	Desc  *int8
}
type HtmlEntityDesc X_htmlEntityDesc
type HtmlEntityDescPtr *HtmlEntityDesc

//go:linkname X__htmlDefaultSAXHandler C.__htmlDefaultSAXHandler
func X__htmlDefaultSAXHandler() *SAXHandlerV1

/*
 * There is only few public functions.
 */
//go:linkname HtmlInitAutoClose C.htmlInitAutoClose
func HtmlInitAutoClose()

// llgo:link (*Char).HtmlTagLookup C.htmlTagLookup
func (recv_ *Char) HtmlTagLookup() *HtmlElemDesc {
	return nil
}

// llgo:link (*Char).HtmlEntityLookup C.htmlEntityLookup
func (recv_ *Char) HtmlEntityLookup() *HtmlEntityDesc {
	return nil
}

//go:linkname HtmlEntityValueLookup C.htmlEntityValueLookup
func HtmlEntityValueLookup(value c.Uint) *HtmlEntityDesc

//go:linkname HtmlIsAutoClosed C.htmlIsAutoClosed
func HtmlIsAutoClosed(doc HtmlDocPtr, elem HtmlNodePtr) c.Int

//go:linkname HtmlAutoCloseTag C.htmlAutoCloseTag
func HtmlAutoCloseTag(doc HtmlDocPtr, name *Char, elem HtmlNodePtr) c.Int

//go:linkname HtmlParseEntityRef C.htmlParseEntityRef
func HtmlParseEntityRef(ctxt HtmlParserCtxtPtr, str **Char) *HtmlEntityDesc

//go:linkname HtmlParseCharRef C.htmlParseCharRef
func HtmlParseCharRef(ctxt HtmlParserCtxtPtr) c.Int

//go:linkname HtmlParseElement C.htmlParseElement
func HtmlParseElement(ctxt HtmlParserCtxtPtr)

//go:linkname HtmlNewParserCtxt C.htmlNewParserCtxt
func HtmlNewParserCtxt() HtmlParserCtxtPtr

// llgo:link (*HtmlSAXHandler).HtmlNewSAXParserCtxt C.htmlNewSAXParserCtxt
func (recv_ *HtmlSAXHandler) HtmlNewSAXParserCtxt(userData unsafe.Pointer) HtmlParserCtxtPtr {
	return nil
}

//go:linkname HtmlCreateMemoryParserCtxt C.htmlCreateMemoryParserCtxt
func HtmlCreateMemoryParserCtxt(buffer *int8, size c.Int) HtmlParserCtxtPtr

//go:linkname HtmlParseDocument C.htmlParseDocument
func HtmlParseDocument(ctxt HtmlParserCtxtPtr) c.Int

// llgo:link (*Char).HtmlSAXParseDoc C.htmlSAXParseDoc
func (recv_ *Char) HtmlSAXParseDoc(encoding *int8, sax HtmlSAXHandlerPtr, userData unsafe.Pointer) HtmlDocPtr {
	return nil
}

// llgo:link (*Char).HtmlParseDoc C.htmlParseDoc
func (recv_ *Char) HtmlParseDoc(encoding *int8) HtmlDocPtr {
	return nil
}

//go:linkname HtmlCreateFileParserCtxt C.htmlCreateFileParserCtxt
func HtmlCreateFileParserCtxt(filename *int8, encoding *int8) HtmlParserCtxtPtr

//go:linkname HtmlSAXParseFile C.htmlSAXParseFile
func HtmlSAXParseFile(filename *int8, encoding *int8, sax HtmlSAXHandlerPtr, userData unsafe.Pointer) HtmlDocPtr

//go:linkname HtmlParseFile C.htmlParseFile
func HtmlParseFile(filename *int8, encoding *int8) HtmlDocPtr

//go:linkname UTF8ToHtml C.UTF8ToHtml
func UTF8ToHtml(out *int8, outlen *c.Int, in *int8, inlen *c.Int) c.Int

//go:linkname HtmlEncodeEntities C.htmlEncodeEntities
func HtmlEncodeEntities(out *int8, outlen *c.Int, in *int8, inlen *c.Int, quoteChar c.Int) c.Int

// llgo:link (*Char).HtmlIsScriptAttribute C.htmlIsScriptAttribute
func (recv_ *Char) HtmlIsScriptAttribute() c.Int {
	return 0
}

//go:linkname HtmlHandleOmittedElem C.htmlHandleOmittedElem
func HtmlHandleOmittedElem(val c.Int) c.Int

/**
 * Interfaces for the Push mode.
 */
//go:linkname HtmlCreatePushParserCtxt C.htmlCreatePushParserCtxt
func HtmlCreatePushParserCtxt(sax HtmlSAXHandlerPtr, user_data unsafe.Pointer, chunk *int8, size c.Int, filename *int8, enc CharEncoding) HtmlParserCtxtPtr

//go:linkname HtmlParseChunk C.htmlParseChunk
func HtmlParseChunk(ctxt HtmlParserCtxtPtr, chunk *int8, size c.Int, terminate c.Int) c.Int

//go:linkname HtmlFreeParserCtxt C.htmlFreeParserCtxt
func HtmlFreeParserCtxt(ctxt HtmlParserCtxtPtr)

type HtmlParserOption c.Int

const (
	HTMLPARSERECOVER   HtmlParserOption = 1
	HTMLPARSENODEFDTD  HtmlParserOption = 4
	HTMLPARSENOERROR   HtmlParserOption = 32
	HTMLPARSENOWARNING HtmlParserOption = 64
	HTMLPARSEPEDANTIC  HtmlParserOption = 128
	HTMLPARSENOBLANKS  HtmlParserOption = 256
	HTMLPARSENONET     HtmlParserOption = 2048
	HTMLPARSENOIMPLIED HtmlParserOption = 8192
	HTMLPARSECOMPACT   HtmlParserOption = 65536
	HTMLPARSEIGNOREENC HtmlParserOption = 2097152
)

//go:linkname HtmlCtxtReset C.htmlCtxtReset
func HtmlCtxtReset(ctxt HtmlParserCtxtPtr)

//go:linkname HtmlCtxtUseOptions C.htmlCtxtUseOptions
func HtmlCtxtUseOptions(ctxt HtmlParserCtxtPtr, options c.Int) c.Int

// llgo:link (*Char).HtmlReadDoc C.htmlReadDoc
func (recv_ *Char) HtmlReadDoc(URL *int8, encoding *int8, options c.Int) HtmlDocPtr {
	return nil
}

//go:linkname HtmlReadFile C.htmlReadFile
func HtmlReadFile(URL *int8, encoding *int8, options c.Int) HtmlDocPtr

//go:linkname HtmlReadMemory C.htmlReadMemory
func HtmlReadMemory(buffer *int8, size c.Int, URL *int8, encoding *int8, options c.Int) HtmlDocPtr

//go:linkname HtmlReadFd C.htmlReadFd
func HtmlReadFd(fd c.Int, URL *int8, encoding *int8, options c.Int) HtmlDocPtr

//go:linkname HtmlReadIO C.htmlReadIO
func HtmlReadIO(ioread InputReadCallback, ioclose InputCloseCallback, ioctx unsafe.Pointer, URL *int8, encoding *int8, options c.Int) HtmlDocPtr

//go:linkname HtmlCtxtParseDocument C.htmlCtxtParseDocument
func HtmlCtxtParseDocument(ctxt HtmlParserCtxtPtr, input ParserInputPtr) HtmlDocPtr

//go:linkname HtmlCtxtReadDoc C.htmlCtxtReadDoc
func HtmlCtxtReadDoc(ctxt ParserCtxtPtr, cur *Char, URL *int8, encoding *int8, options c.Int) HtmlDocPtr

//go:linkname HtmlCtxtReadFile C.htmlCtxtReadFile
func HtmlCtxtReadFile(ctxt ParserCtxtPtr, filename *int8, encoding *int8, options c.Int) HtmlDocPtr

//go:linkname HtmlCtxtReadMemory C.htmlCtxtReadMemory
func HtmlCtxtReadMemory(ctxt ParserCtxtPtr, buffer *int8, size c.Int, URL *int8, encoding *int8, options c.Int) HtmlDocPtr

//go:linkname HtmlCtxtReadFd C.htmlCtxtReadFd
func HtmlCtxtReadFd(ctxt ParserCtxtPtr, fd c.Int, URL *int8, encoding *int8, options c.Int) HtmlDocPtr

//go:linkname HtmlCtxtReadIO C.htmlCtxtReadIO
func HtmlCtxtReadIO(ctxt ParserCtxtPtr, ioread InputReadCallback, ioclose InputCloseCallback, ioctx unsafe.Pointer, URL *int8, encoding *int8, options c.Int) HtmlDocPtr

type HtmlStatus c.Int

const (
	HTMLNA         HtmlStatus = 0
	HTMLINVALID    HtmlStatus = 1
	HTMLDEPRECATED HtmlStatus = 2
	HTMLVALID      HtmlStatus = 4
	HTMLREQUIRED   HtmlStatus = 12
)

/* Using htmlElemDesc rather than name here, to emphasise the fact
   that otherwise there's a lookup overhead
*/
// llgo:link (*HtmlElemDesc).HtmlAttrAllowed C.htmlAttrAllowed
func (recv_ *HtmlElemDesc) HtmlAttrAllowed(*Char, c.Int) HtmlStatus {
	return 0
}

// llgo:link (*HtmlElemDesc).HtmlElementAllowedHere C.htmlElementAllowedHere
func (recv_ *HtmlElemDesc) HtmlElementAllowedHere(*Char) c.Int {
	return 0
}

// llgo:link (*HtmlElemDesc).HtmlElementStatusHere C.htmlElementStatusHere
func (recv_ *HtmlElemDesc) HtmlElementStatusHere(*HtmlElemDesc) HtmlStatus {
	return 0
}

//go:linkname HtmlNodeStatus C.htmlNodeStatus
func HtmlNodeStatus(HtmlNodePtr, c.Int) HtmlStatus
