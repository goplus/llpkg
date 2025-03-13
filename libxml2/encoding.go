package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type CharEncError c.Int

const (
	ENCERRSUCCESS  CharEncError = 0
	ENCERRSPACE    CharEncError = -1
	ENCERRINPUT    CharEncError = -2
	ENCERRPARTIAL  CharEncError = -3
	ENCERRINTERNAL CharEncError = -4
	ENCERRMEMORY   CharEncError = -5
)

type CharEncoding c.Int

const (
	CHARENCODINGERROR    CharEncoding = -1
	CHARENCODINGNONE     CharEncoding = 0
	CHARENCODINGUTF8     CharEncoding = 1
	CHARENCODINGUTF16LE  CharEncoding = 2
	CHARENCODINGUTF16BE  CharEncoding = 3
	CHARENCODINGUCS4LE   CharEncoding = 4
	CHARENCODINGUCS4BE   CharEncoding = 5
	CHARENCODINGEBCDIC   CharEncoding = 6
	CHARENCODINGUCS42143 CharEncoding = 7
	CHARENCODINGUCS43412 CharEncoding = 8
	CHARENCODINGUCS2     CharEncoding = 9
	CHARENCODING88591    CharEncoding = 10
	CHARENCODING88592    CharEncoding = 11
	CHARENCODING88593    CharEncoding = 12
	CHARENCODING88594    CharEncoding = 13
	CHARENCODING88595    CharEncoding = 14
	CHARENCODING88596    CharEncoding = 15
	CHARENCODING88597    CharEncoding = 16
	CHARENCODING88598    CharEncoding = 17
	CHARENCODING88599    CharEncoding = 18
	CHARENCODING2022JP   CharEncoding = 19
	CHARENCODINGSHIFTJIS CharEncoding = 20
	CHARENCODINGEUCJP    CharEncoding = 21
	CHARENCODINGASCII    CharEncoding = 22
)

// llgo:type C
type CharEncodingInputFunc func(*int8, *c.Int, *int8, *c.Int) c.Int

// llgo:type C
type CharEncodingOutputFunc func(*int8, *c.Int, *int8, *c.Int) c.Int

type X_xmlCharEncodingHandler struct {
	Name   *int8
	Input  unsafe.Pointer
	Output unsafe.Pointer
}
type CharEncodingHandler X_xmlCharEncodingHandler
type CharEncodingHandlerPtr *CharEncodingHandler

/*
 * Interfaces for encoding handlers.
 */
//go:linkname InitCharEncodingHandlers C.xmlInitCharEncodingHandlers
func InitCharEncodingHandlers()

//go:linkname CleanupCharEncodingHandlers C.xmlCleanupCharEncodingHandlers
func CleanupCharEncodingHandlers()

//go:linkname RegisterCharEncodingHandler C.xmlRegisterCharEncodingHandler
func RegisterCharEncodingHandler(handler CharEncodingHandlerPtr)

// llgo:link CharEncoding.LookupCharEncodingHandler C.xmlLookupCharEncodingHandler
func (recv_ CharEncoding) LookupCharEncodingHandler(out *CharEncodingHandlerPtr) c.Int {
	return 0
}

//go:linkname OpenCharEncodingHandler C.xmlOpenCharEncodingHandler
func OpenCharEncodingHandler(name *int8, output c.Int, out *CharEncodingHandlerPtr) c.Int

// llgo:link CharEncoding.GetCharEncodingHandler C.xmlGetCharEncodingHandler
func (recv_ CharEncoding) GetCharEncodingHandler() CharEncodingHandlerPtr {
	return nil
}

//go:linkname FindCharEncodingHandler C.xmlFindCharEncodingHandler
func FindCharEncodingHandler(name *int8) CharEncodingHandlerPtr

//go:linkname NewCharEncodingHandler C.xmlNewCharEncodingHandler
func NewCharEncodingHandler(name *int8, input CharEncodingInputFunc, output CharEncodingOutputFunc) CharEncodingHandlerPtr

/*
 * Interfaces for encoding names and aliases.
 */
//go:linkname AddEncodingAlias C.xmlAddEncodingAlias
func AddEncodingAlias(name *int8, alias *int8) c.Int

//go:linkname DelEncodingAlias C.xmlDelEncodingAlias
func DelEncodingAlias(alias *int8) c.Int

//go:linkname GetEncodingAlias C.xmlGetEncodingAlias
func GetEncodingAlias(alias *int8) *int8

//go:linkname CleanupEncodingAliases C.xmlCleanupEncodingAliases
func CleanupEncodingAliases()

//go:linkname ParseCharEncoding C.xmlParseCharEncoding
func ParseCharEncoding(name *int8) CharEncoding

// llgo:link CharEncoding.GetCharEncodingName C.xmlGetCharEncodingName
func (recv_ CharEncoding) GetCharEncodingName() *int8 {
	return nil
}

/*
 * Interfaces directly used by the parsers.
 */
//go:linkname DetectCharEncoding C.xmlDetectCharEncoding
func DetectCharEncoding(in *int8, len c.Int) CharEncoding

/** DOC_ENABLE */
// llgo:link (*CharEncodingHandler).CharEncOutFunc C.xmlCharEncOutFunc
func (recv_ *CharEncodingHandler) CharEncOutFunc(out *X_xmlBuffer, in *X_xmlBuffer) c.Int {
	return 0
}

// llgo:link (*CharEncodingHandler).CharEncInFunc C.xmlCharEncInFunc
func (recv_ *CharEncodingHandler) CharEncInFunc(out *X_xmlBuffer, in *X_xmlBuffer) c.Int {
	return 0
}

// llgo:link (*CharEncodingHandler).CharEncFirstLine C.xmlCharEncFirstLine
func (recv_ *CharEncodingHandler) CharEncFirstLine(out *X_xmlBuffer, in *X_xmlBuffer) c.Int {
	return 0
}

// llgo:link (*CharEncodingHandler).CharEncCloseFunc C.xmlCharEncCloseFunc
func (recv_ *CharEncodingHandler) CharEncCloseFunc() c.Int {
	return 0
}

/*
 * Export a few useful functions
 */
//go:linkname UTF8Toisolat1 C.UTF8Toisolat1
func UTF8Toisolat1(out *int8, outlen *c.Int, in *int8, inlen *c.Int) c.Int

//go:linkname Isolat1ToUTF8 C.isolat1ToUTF8
func Isolat1ToUTF8(out *int8, outlen *c.Int, in *int8, inlen *c.Int) c.Int
