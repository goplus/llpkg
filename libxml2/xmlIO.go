package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

// llgo:type C
type InputMatchCallback func(*int8) c.Int

// llgo:type C
type InputOpenCallback func(*int8) unsafe.Pointer

// llgo:type C
type InputReadCallback func(unsafe.Pointer, *int8, c.Int) c.Int

// llgo:type C
type InputCloseCallback func(unsafe.Pointer) c.Int

// llgo:type C
type OutputMatchCallback func(*int8) c.Int

// llgo:type C
type OutputOpenCallback func(*int8) unsafe.Pointer

// llgo:type C
type OutputWriteCallback func(unsafe.Pointer, *int8, c.Int) c.Int

// llgo:type C
type OutputCloseCallback func(unsafe.Pointer) c.Int

// llgo:type C
type ParserInputBufferCreateFilenameFunc func(*int8, CharEncoding) ParserInputBufferPtr

// llgo:type C
type OutputBufferCreateFilenameFunc func(*int8, CharEncodingHandlerPtr, c.Int) OutputBufferPtr

//go:linkname X__xmlParserInputBufferCreateFilenameValue C.__xmlParserInputBufferCreateFilenameValue
func X__xmlParserInputBufferCreateFilenameValue() ParserInputBufferCreateFilenameFunc

//go:linkname X__xmlOutputBufferCreateFilenameValue C.__xmlOutputBufferCreateFilenameValue
func X__xmlOutputBufferCreateFilenameValue() OutputBufferCreateFilenameFunc

/*
 * Interfaces for input
 */
//go:linkname CleanupInputCallbacks C.xmlCleanupInputCallbacks
func CleanupInputCallbacks()

//go:linkname PopInputCallbacks C.xmlPopInputCallbacks
func PopInputCallbacks() c.Int

//go:linkname RegisterDefaultInputCallbacks C.xmlRegisterDefaultInputCallbacks
func RegisterDefaultInputCallbacks()

// llgo:link CharEncoding.AllocParserInputBuffer C.xmlAllocParserInputBuffer
func (recv_ CharEncoding) AllocParserInputBuffer() ParserInputBufferPtr {
	return nil
}

//go:linkname ParserInputBufferCreateFilename C.xmlParserInputBufferCreateFilename
func ParserInputBufferCreateFilename(URI *int8, enc CharEncoding) ParserInputBufferPtr

//go:linkname ParserInputBufferCreateFile C.xmlParserInputBufferCreateFile
func ParserInputBufferCreateFile(file *c.FILE, enc CharEncoding) ParserInputBufferPtr

//go:linkname ParserInputBufferCreateFd C.xmlParserInputBufferCreateFd
func ParserInputBufferCreateFd(fd c.Int, enc CharEncoding) ParserInputBufferPtr

//go:linkname ParserInputBufferCreateMem C.xmlParserInputBufferCreateMem
func ParserInputBufferCreateMem(mem *int8, size c.Int, enc CharEncoding) ParserInputBufferPtr

//go:linkname ParserInputBufferCreateStatic C.xmlParserInputBufferCreateStatic
func ParserInputBufferCreateStatic(mem *int8, size c.Int, enc CharEncoding) ParserInputBufferPtr

//go:linkname ParserInputBufferCreateIO C.xmlParserInputBufferCreateIO
func ParserInputBufferCreateIO(ioread InputReadCallback, ioclose InputCloseCallback, ioctx unsafe.Pointer, enc CharEncoding) ParserInputBufferPtr

//go:linkname ParserInputBufferRead C.xmlParserInputBufferRead
func ParserInputBufferRead(in ParserInputBufferPtr, len c.Int) c.Int

//go:linkname ParserInputBufferGrow C.xmlParserInputBufferGrow
func ParserInputBufferGrow(in ParserInputBufferPtr, len c.Int) c.Int

//go:linkname ParserInputBufferPush C.xmlParserInputBufferPush
func ParserInputBufferPush(in ParserInputBufferPtr, len c.Int, buf *int8) c.Int

//go:linkname FreeParserInputBuffer C.xmlFreeParserInputBuffer
func FreeParserInputBuffer(in ParserInputBufferPtr)

//go:linkname ParserGetDirectory C.xmlParserGetDirectory
func ParserGetDirectory(filename *int8) *int8

//go:linkname RegisterInputCallbacks C.xmlRegisterInputCallbacks
func RegisterInputCallbacks(matchFunc InputMatchCallback, openFunc InputOpenCallback, readFunc InputReadCallback, closeFunc InputCloseCallback) c.Int

//go:linkname X__xmlParserInputBufferCreateFilename C.__xmlParserInputBufferCreateFilename
func X__xmlParserInputBufferCreateFilename(URI *int8, enc CharEncoding) ParserInputBufferPtr

/*
 * Interfaces for output
 */
//go:linkname CleanupOutputCallbacks C.xmlCleanupOutputCallbacks
func CleanupOutputCallbacks()

//go:linkname PopOutputCallbacks C.xmlPopOutputCallbacks
func PopOutputCallbacks() c.Int

//go:linkname RegisterDefaultOutputCallbacks C.xmlRegisterDefaultOutputCallbacks
func RegisterDefaultOutputCallbacks()

//go:linkname AllocOutputBuffer C.xmlAllocOutputBuffer
func AllocOutputBuffer(encoder CharEncodingHandlerPtr) OutputBufferPtr

//go:linkname OutputBufferCreateFilename C.xmlOutputBufferCreateFilename
func OutputBufferCreateFilename(URI *int8, encoder CharEncodingHandlerPtr, compression c.Int) OutputBufferPtr

//go:linkname OutputBufferCreateFile C.xmlOutputBufferCreateFile
func OutputBufferCreateFile(file *c.FILE, encoder CharEncodingHandlerPtr) OutputBufferPtr

//go:linkname OutputBufferCreateBuffer C.xmlOutputBufferCreateBuffer
func OutputBufferCreateBuffer(buffer BufferPtr, encoder CharEncodingHandlerPtr) OutputBufferPtr

//go:linkname OutputBufferCreateFd C.xmlOutputBufferCreateFd
func OutputBufferCreateFd(fd c.Int, encoder CharEncodingHandlerPtr) OutputBufferPtr

//go:linkname OutputBufferCreateIO C.xmlOutputBufferCreateIO
func OutputBufferCreateIO(iowrite OutputWriteCallback, ioclose OutputCloseCallback, ioctx unsafe.Pointer, encoder CharEncodingHandlerPtr) OutputBufferPtr

/* Couple of APIs to get the output without digging into the buffers */
//go:linkname OutputBufferGetContent C.xmlOutputBufferGetContent
func OutputBufferGetContent(out OutputBufferPtr) *Char

//go:linkname OutputBufferGetSize C.xmlOutputBufferGetSize
func OutputBufferGetSize(out OutputBufferPtr) uintptr

//go:linkname OutputBufferWrite C.xmlOutputBufferWrite
func OutputBufferWrite(out OutputBufferPtr, len c.Int, buf *int8) c.Int

//go:linkname OutputBufferWriteString C.xmlOutputBufferWriteString
func OutputBufferWriteString(out OutputBufferPtr, str *int8) c.Int

//go:linkname OutputBufferWriteEscape C.xmlOutputBufferWriteEscape
func OutputBufferWriteEscape(out OutputBufferPtr, str *Char, escaping CharEncodingOutputFunc) c.Int

//go:linkname OutputBufferFlush C.xmlOutputBufferFlush
func OutputBufferFlush(out OutputBufferPtr) c.Int

//go:linkname OutputBufferClose C.xmlOutputBufferClose
func OutputBufferClose(out OutputBufferPtr) c.Int

//go:linkname RegisterOutputCallbacks C.xmlRegisterOutputCallbacks
func RegisterOutputCallbacks(matchFunc OutputMatchCallback, openFunc OutputOpenCallback, writeFunc OutputWriteCallback, closeFunc OutputCloseCallback) c.Int

//go:linkname X__xmlOutputBufferCreateFilename C.__xmlOutputBufferCreateFilename
func X__xmlOutputBufferCreateFilename(URI *int8, encoder CharEncodingHandlerPtr, compression c.Int) OutputBufferPtr

/*  This function only exists if HTTP support built into the library  */
//go:linkname RegisterHTTPPostCallbacks C.xmlRegisterHTTPPostCallbacks
func RegisterHTTPPostCallbacks()

//go:linkname CheckHTTPInput C.xmlCheckHTTPInput
func CheckHTTPInput(ctxt ParserCtxtPtr, ret ParserInputPtr) ParserInputPtr

/*
 * A predefined entity loader disabling network accesses
 */
//go:linkname NoNetExternalEntityLoader C.xmlNoNetExternalEntityLoader
func NoNetExternalEntityLoader(URL *int8, ID *int8, ctxt ParserCtxtPtr) ParserInputPtr

// llgo:link (*Char).NormalizeWindowsPath C.xmlNormalizeWindowsPath
func (recv_ *Char) NormalizeWindowsPath() *Char {
	return nil
}

//go:linkname CheckFilename C.xmlCheckFilename
func CheckFilename(path *int8) c.Int

/**
 * Default 'file://' protocol callbacks
 */
//go:linkname FileMatch C.xmlFileMatch
func FileMatch(filename *int8) c.Int

//go:linkname FileOpen C.xmlFileOpen
func FileOpen(filename *int8) unsafe.Pointer

//go:linkname FileRead C.xmlFileRead
func FileRead(context unsafe.Pointer, buffer *int8, len c.Int) c.Int

//go:linkname FileClose C.xmlFileClose
func FileClose(context unsafe.Pointer) c.Int

/**
 * Default 'http://' protocol callbacks
 */
//go:linkname IOHTTPMatch C.xmlIOHTTPMatch
func IOHTTPMatch(filename *int8) c.Int

//go:linkname IOHTTPOpen C.xmlIOHTTPOpen
func IOHTTPOpen(filename *int8) unsafe.Pointer

//go:linkname IOHTTPOpenW C.xmlIOHTTPOpenW
func IOHTTPOpenW(post_uri *int8, compression c.Int) unsafe.Pointer

//go:linkname IOHTTPRead C.xmlIOHTTPRead
func IOHTTPRead(context unsafe.Pointer, buffer *int8, len c.Int) c.Int

//go:linkname IOHTTPClose C.xmlIOHTTPClose
func IOHTTPClose(context unsafe.Pointer) c.Int

/**
 * Default 'ftp://' protocol callbacks
 */
//go:linkname IOFTPMatch C.xmlIOFTPMatch
func IOFTPMatch(filename *int8) c.Int

//go:linkname IOFTPOpen C.xmlIOFTPOpen
func IOFTPOpen(filename *int8) unsafe.Pointer

//go:linkname IOFTPRead C.xmlIOFTPRead
func IOFTPRead(context unsafe.Pointer, buffer *int8, len c.Int) c.Int

//go:linkname IOFTPClose C.xmlIOFTPClose
func IOFTPClose(context unsafe.Pointer) c.Int

//go:linkname ParserInputBufferCreateFilenameDefault C.xmlParserInputBufferCreateFilenameDefault
func ParserInputBufferCreateFilenameDefault(func_ ParserInputBufferCreateFilenameFunc) ParserInputBufferCreateFilenameFunc

//go:linkname OutputBufferCreateFilenameDefault C.xmlOutputBufferCreateFilenameDefault
func OutputBufferCreateFilenameDefault(func_ OutputBufferCreateFilenameFunc) OutputBufferCreateFilenameFunc

//go:linkname ThrDefOutputBufferCreateFilenameDefault C.xmlThrDefOutputBufferCreateFilenameDefault
func ThrDefOutputBufferCreateFilenameDefault(func_ OutputBufferCreateFilenameFunc) OutputBufferCreateFilenameFunc

//go:linkname ThrDefParserInputBufferCreateFilenameDefault C.xmlThrDefParserInputBufferCreateFilenameDefault
func ThrDefParserInputBufferCreateFilenameDefault(func_ ParserInputBufferCreateFilenameFunc) ParserInputBufferCreateFilenameFunc
