package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const DEFAULT_VERSION = "1.0"
const DETECT_IDS = 2
const COMPLETE_ATTRS = 4
const SKIP_IDS = 8
const SAX2_MAGIC = 0xDEEDBEAF

// llgo:type C
type ParserInputDeallocate func(*Char)

type X_xmlParserNodeInfo struct {
	Node      *X_xmlNode
	BeginPos  c.Ulong
	BeginLine c.Ulong
	EndPos    c.Ulong
	EndLine   c.Ulong
}
type ParserNodeInfo X_xmlParserNodeInfo
type ParserNodeInfoPtr *ParserNodeInfo

type X_xmlParserNodeInfoSeq struct {
	Maximum c.Ulong
	Length  c.Ulong
	Buffer  *ParserNodeInfo
}
type ParserNodeInfoSeq X_xmlParserNodeInfoSeq
type ParserNodeInfoSeqPtr *ParserNodeInfoSeq
type ParserInputState c.Int

const (
	PARSEREOF            ParserInputState = -1
	PARSERSTART          ParserInputState = 0
	PARSERMISC           ParserInputState = 1
	PARSERPI             ParserInputState = 2
	PARSERDTD            ParserInputState = 3
	PARSERPROLOG         ParserInputState = 4
	PARSERCOMMENT        ParserInputState = 5
	PARSERSTARTTAG       ParserInputState = 6
	PARSERCONTENT        ParserInputState = 7
	PARSERCDATASECTION   ParserInputState = 8
	PARSERENDTAG         ParserInputState = 9
	PARSERENTITYDECL     ParserInputState = 10
	PARSERENTITYVALUE    ParserInputState = 11
	PARSERATTRIBUTEVALUE ParserInputState = 12
	PARSERSYSTEMLITERAL  ParserInputState = 13
	PARSEREPILOG         ParserInputState = 14
	PARSERIGNORE         ParserInputState = 15
	PARSERPUBLICLITERAL  ParserInputState = 16
	PARSERXMLDECL        ParserInputState = 17
)

type ParserMode c.Int

const (
	PARSEUNKNOWN ParserMode = 0
	PARSEDOM     ParserMode = 1
	PARSESAX     ParserMode = 2
	PARSEPUSHDOM ParserMode = 3
	PARSEPUSHSAX ParserMode = 4
	PARSEREADER  ParserMode = 5
)

type X_xmlStartTag struct {
	Unused [8]uint8
}
type StartTag X_xmlStartTag

type X_xmlParserNsData struct {
	Unused [8]uint8
}
type ParserNsData X_xmlParserNsData

type X_xmlAttrHashBucket struct {
	Unused [8]uint8
}
type AttrHashBucket X_xmlAttrHashBucket

// llgo:type C
type ResolveEntitySAXFunc func(unsafe.Pointer, *Char, *Char) ParserInputPtr

// llgo:type C
type InternalSubsetSAXFunc func(unsafe.Pointer, *Char, *Char, *Char)

// llgo:type C
type ExternalSubsetSAXFunc func(unsafe.Pointer, *Char, *Char, *Char)

// llgo:type C
type GetEntitySAXFunc func(unsafe.Pointer, *Char) EntityPtr

// llgo:type C
type GetParameterEntitySAXFunc func(unsafe.Pointer, *Char) EntityPtr

// llgo:type C
type EntityDeclSAXFunc func(unsafe.Pointer, *Char, c.Int, *Char, *Char, *Char)

// llgo:type C
type NotationDeclSAXFunc func(unsafe.Pointer, *Char, *Char, *Char)

// llgo:type C
type AttributeDeclSAXFunc func(unsafe.Pointer, *Char, *Char, c.Int, c.Int, *Char, EnumerationPtr)

// llgo:type C
type ElementDeclSAXFunc func(unsafe.Pointer, *Char, c.Int, ElementContentPtr)

// llgo:type C
type UnparsedEntityDeclSAXFunc func(unsafe.Pointer, *Char, *Char, *Char, *Char)

// llgo:type C
type SetDocumentLocatorSAXFunc func(unsafe.Pointer, SAXLocatorPtr)

// llgo:type C
type StartDocumentSAXFunc func(unsafe.Pointer)

// llgo:type C
type EndDocumentSAXFunc func(unsafe.Pointer)

// llgo:type C
type StartElementSAXFunc func(unsafe.Pointer, *Char, **Char)

// llgo:type C
type EndElementSAXFunc func(unsafe.Pointer, *Char)

// llgo:type C
type AttributeSAXFunc func(unsafe.Pointer, *Char, *Char)

// llgo:type C
type ReferenceSAXFunc func(unsafe.Pointer, *Char)

// llgo:type C
type CharactersSAXFunc func(unsafe.Pointer, *Char, c.Int)

// llgo:type C
type IgnorableWhitespaceSAXFunc func(unsafe.Pointer, *Char, c.Int)

// llgo:type C
type ProcessingInstructionSAXFunc func(unsafe.Pointer, *Char, *Char)

// llgo:type C
type CommentSAXFunc func(unsafe.Pointer, *Char)

// llgo:type C
type CdataBlockSAXFunc func(unsafe.Pointer, *Char, c.Int)

// llgo:type C
type WarningSAXFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})

// llgo:type C
type ErrorSAXFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})

// llgo:type C
type FatalErrorSAXFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})

// llgo:type C
type IsStandaloneSAXFunc func(unsafe.Pointer) c.Int

// llgo:type C
type HasInternalSubsetSAXFunc func(unsafe.Pointer) c.Int

// llgo:type C
type HasExternalSubsetSAXFunc func(unsafe.Pointer) c.Int

// llgo:type C
type StartElementNsSAX2Func func(unsafe.Pointer, *Char, *Char, *Char, c.Int, **Char, c.Int, c.Int, **Char)

// llgo:type C
type EndElementNsSAX2Func func(unsafe.Pointer, *Char, *Char, *Char)

type X_xmlSAXHandlerV1 struct {
	InternalSubset        unsafe.Pointer
	IsStandalone          unsafe.Pointer
	HasInternalSubset     unsafe.Pointer
	HasExternalSubset     unsafe.Pointer
	ResolveEntity         unsafe.Pointer
	GetEntity             unsafe.Pointer
	EntityDecl            unsafe.Pointer
	NotationDecl          unsafe.Pointer
	AttributeDecl         unsafe.Pointer
	ElementDecl           unsafe.Pointer
	UnparsedEntityDecl    unsafe.Pointer
	SetDocumentLocator    unsafe.Pointer
	StartDocument         unsafe.Pointer
	EndDocument           unsafe.Pointer
	StartElement          unsafe.Pointer
	EndElement            unsafe.Pointer
	Reference             unsafe.Pointer
	Characters            unsafe.Pointer
	IgnorableWhitespace   unsafe.Pointer
	ProcessingInstruction unsafe.Pointer
	Comment               unsafe.Pointer
	Warning               unsafe.Pointer
	Error                 unsafe.Pointer
	FatalError            unsafe.Pointer
	GetParameterEntity    unsafe.Pointer
	CdataBlock            unsafe.Pointer
	ExternalSubset        unsafe.Pointer
	Initialized           c.Uint
}
type SAXHandlerV1 X_xmlSAXHandlerV1
type SAXHandlerV1Ptr *SAXHandlerV1

// llgo:type C
type ExternalEntityLoader func(*int8, *int8, ParserCtxtPtr) ParserInputPtr

/* backward compatibility */
//go:linkname X__xmlParserVersion C.__xmlParserVersion
func X__xmlParserVersion() **int8

//go:linkname X__oldXMLWDcompatibility C.__oldXMLWDcompatibility
func X__oldXMLWDcompatibility() *c.Int

//go:linkname X__xmlParserDebugEntities C.__xmlParserDebugEntities
func X__xmlParserDebugEntities() *c.Int

//go:linkname X__xmlDefaultSAXLocator C.__xmlDefaultSAXLocator
func X__xmlDefaultSAXLocator() *SAXLocator

//go:linkname X__xmlDefaultSAXHandler C.__xmlDefaultSAXHandler
func X__xmlDefaultSAXHandler() *SAXHandlerV1

//go:linkname X__xmlDoValidityCheckingDefaultValue C.__xmlDoValidityCheckingDefaultValue
func X__xmlDoValidityCheckingDefaultValue() *c.Int

//go:linkname X__xmlGetWarningsDefaultValue C.__xmlGetWarningsDefaultValue
func X__xmlGetWarningsDefaultValue() *c.Int

//go:linkname X__xmlKeepBlanksDefaultValue C.__xmlKeepBlanksDefaultValue
func X__xmlKeepBlanksDefaultValue() *c.Int

//go:linkname X__xmlLineNumbersDefaultValue C.__xmlLineNumbersDefaultValue
func X__xmlLineNumbersDefaultValue() *c.Int

//go:linkname X__xmlLoadExtDtdDefaultValue C.__xmlLoadExtDtdDefaultValue
func X__xmlLoadExtDtdDefaultValue() *c.Int

//go:linkname X__xmlPedanticParserDefaultValue C.__xmlPedanticParserDefaultValue
func X__xmlPedanticParserDefaultValue() *c.Int

//go:linkname X__xmlSubstituteEntitiesDefaultValue C.__xmlSubstituteEntitiesDefaultValue
func X__xmlSubstituteEntitiesDefaultValue() *c.Int

//go:linkname X__xmlIndentTreeOutput C.__xmlIndentTreeOutput
func X__xmlIndentTreeOutput() *c.Int

//go:linkname X__xmlTreeIndentString C.__xmlTreeIndentString
func X__xmlTreeIndentString() **int8

//go:linkname X__xmlSaveNoEmptyTags C.__xmlSaveNoEmptyTags
func X__xmlSaveNoEmptyTags() *c.Int

/*
 * Init/Cleanup
 */
//go:linkname InitParser C.xmlInitParser
func InitParser()

//go:linkname CleanupParser C.xmlCleanupParser
func CleanupParser()

//go:linkname InitGlobals C.xmlInitGlobals
func InitGlobals()

//go:linkname CleanupGlobals C.xmlCleanupGlobals
func CleanupGlobals()

/*
 * Input functions
 */
//go:linkname ParserInputRead C.xmlParserInputRead
func ParserInputRead(in ParserInputPtr, len c.Int) c.Int

//go:linkname ParserInputGrow C.xmlParserInputGrow
func ParserInputGrow(in ParserInputPtr, len c.Int) c.Int

/*
 * Basic parsing Interfaces
 */
// llgo:link (*Char).ParseDoc C.xmlParseDoc
func (recv_ *Char) ParseDoc() DocPtr {
	return nil
}

//go:linkname ParseFile C.xmlParseFile
func ParseFile(filename *int8) DocPtr

//go:linkname ParseMemory C.xmlParseMemory
func ParseMemory(buffer *int8, size c.Int) DocPtr

//go:linkname SubstituteEntitiesDefault C.xmlSubstituteEntitiesDefault
func SubstituteEntitiesDefault(val c.Int) c.Int

//go:linkname ThrDefSubstituteEntitiesDefaultValue C.xmlThrDefSubstituteEntitiesDefaultValue
func ThrDefSubstituteEntitiesDefaultValue(v c.Int) c.Int

//go:linkname KeepBlanksDefault C.xmlKeepBlanksDefault
func KeepBlanksDefault(val c.Int) c.Int

//go:linkname ThrDefKeepBlanksDefaultValue C.xmlThrDefKeepBlanksDefaultValue
func ThrDefKeepBlanksDefaultValue(v c.Int) c.Int

//go:linkname StopParser C.xmlStopParser
func StopParser(ctxt ParserCtxtPtr)

//go:linkname PedanticParserDefault C.xmlPedanticParserDefault
func PedanticParserDefault(val c.Int) c.Int

//go:linkname ThrDefPedanticParserDefaultValue C.xmlThrDefPedanticParserDefaultValue
func ThrDefPedanticParserDefaultValue(v c.Int) c.Int

//go:linkname LineNumbersDefault C.xmlLineNumbersDefault
func LineNumbersDefault(val c.Int) c.Int

//go:linkname ThrDefLineNumbersDefaultValue C.xmlThrDefLineNumbersDefaultValue
func ThrDefLineNumbersDefaultValue(v c.Int) c.Int

//go:linkname ThrDefDoValidityCheckingDefaultValue C.xmlThrDefDoValidityCheckingDefaultValue
func ThrDefDoValidityCheckingDefaultValue(v c.Int) c.Int

//go:linkname ThrDefGetWarningsDefaultValue C.xmlThrDefGetWarningsDefaultValue
func ThrDefGetWarningsDefaultValue(v c.Int) c.Int

//go:linkname ThrDefLoadExtDtdDefaultValue C.xmlThrDefLoadExtDtdDefaultValue
func ThrDefLoadExtDtdDefaultValue(v c.Int) c.Int

//go:linkname ThrDefParserDebugEntities C.xmlThrDefParserDebugEntities
func ThrDefParserDebugEntities(v c.Int) c.Int

/*
 * Recovery mode
 */
// llgo:link (*Char).RecoverDoc C.xmlRecoverDoc
func (recv_ *Char) RecoverDoc() DocPtr {
	return nil
}

//go:linkname RecoverMemory C.xmlRecoverMemory
func RecoverMemory(buffer *int8, size c.Int) DocPtr

//go:linkname RecoverFile C.xmlRecoverFile
func RecoverFile(filename *int8) DocPtr

/*
 * Less common routines and SAX interfaces
 */
//go:linkname ParseDocument C.xmlParseDocument
func ParseDocument(ctxt ParserCtxtPtr) c.Int

//go:linkname ParseExtParsedEnt C.xmlParseExtParsedEnt
func ParseExtParsedEnt(ctxt ParserCtxtPtr) c.Int

//go:linkname SAXUserParseFile C.xmlSAXUserParseFile
func SAXUserParseFile(sax SAXHandlerPtr, user_data unsafe.Pointer, filename *int8) c.Int

//go:linkname SAXUserParseMemory C.xmlSAXUserParseMemory
func SAXUserParseMemory(sax SAXHandlerPtr, user_data unsafe.Pointer, buffer *int8, size c.Int) c.Int

//go:linkname SAXParseDoc C.xmlSAXParseDoc
func SAXParseDoc(sax SAXHandlerPtr, cur *Char, recovery c.Int) DocPtr

//go:linkname SAXParseMemory C.xmlSAXParseMemory
func SAXParseMemory(sax SAXHandlerPtr, buffer *int8, size c.Int, recovery c.Int) DocPtr

//go:linkname SAXParseMemoryWithData C.xmlSAXParseMemoryWithData
func SAXParseMemoryWithData(sax SAXHandlerPtr, buffer *int8, size c.Int, recovery c.Int, data unsafe.Pointer) DocPtr

//go:linkname SAXParseFile C.xmlSAXParseFile
func SAXParseFile(sax SAXHandlerPtr, filename *int8, recovery c.Int) DocPtr

//go:linkname SAXParseFileWithData C.xmlSAXParseFileWithData
func SAXParseFileWithData(sax SAXHandlerPtr, filename *int8, recovery c.Int, data unsafe.Pointer) DocPtr

//go:linkname SAXParseEntity C.xmlSAXParseEntity
func SAXParseEntity(sax SAXHandlerPtr, filename *int8) DocPtr

//go:linkname ParseEntity C.xmlParseEntity
func ParseEntity(filename *int8) DocPtr

//go:linkname SAXParseDTD C.xmlSAXParseDTD
func SAXParseDTD(sax SAXHandlerPtr, ExternalID *Char, SystemID *Char) DtdPtr

// llgo:link (*Char).ParseDTD C.xmlParseDTD
func (recv_ *Char) ParseDTD(SystemID *Char) DtdPtr {
	return nil
}

//go:linkname IOParseDTD C.xmlIOParseDTD
func IOParseDTD(sax SAXHandlerPtr, input ParserInputBufferPtr, enc CharEncoding) DtdPtr

//go:linkname ParseBalancedChunkMemory C.xmlParseBalancedChunkMemory
func ParseBalancedChunkMemory(doc DocPtr, sax SAXHandlerPtr, user_data unsafe.Pointer, depth c.Int, string *Char, lst *NodePtr) c.Int

//go:linkname ParseInNodeContext C.xmlParseInNodeContext
func ParseInNodeContext(node NodePtr, data *int8, datalen c.Int, options c.Int, lst *NodePtr) ParserErrors

//go:linkname ParseBalancedChunkMemoryRecover C.xmlParseBalancedChunkMemoryRecover
func ParseBalancedChunkMemoryRecover(doc DocPtr, sax SAXHandlerPtr, user_data unsafe.Pointer, depth c.Int, string *Char, lst *NodePtr, recover c.Int) c.Int

//go:linkname ParseExternalEntity C.xmlParseExternalEntity
func ParseExternalEntity(doc DocPtr, sax SAXHandlerPtr, user_data unsafe.Pointer, depth c.Int, URL *Char, ID *Char, lst *NodePtr) c.Int

//go:linkname ParseCtxtExternalEntity C.xmlParseCtxtExternalEntity
func ParseCtxtExternalEntity(ctx ParserCtxtPtr, URL *Char, ID *Char, lst *NodePtr) c.Int

/*
 * Parser contexts handling.
 */
//go:linkname NewParserCtxt C.xmlNewParserCtxt
func NewParserCtxt() ParserCtxtPtr

// llgo:link (*SAXHandler).NewSAXParserCtxt C.xmlNewSAXParserCtxt
func (recv_ *SAXHandler) NewSAXParserCtxt(userData unsafe.Pointer) ParserCtxtPtr {
	return nil
}

//go:linkname InitParserCtxt C.xmlInitParserCtxt
func InitParserCtxt(ctxt ParserCtxtPtr) c.Int

//go:linkname ClearParserCtxt C.xmlClearParserCtxt
func ClearParserCtxt(ctxt ParserCtxtPtr)

//go:linkname FreeParserCtxt C.xmlFreeParserCtxt
func FreeParserCtxt(ctxt ParserCtxtPtr)

//go:linkname SetupParserForBuffer C.xmlSetupParserForBuffer
func SetupParserForBuffer(ctxt ParserCtxtPtr, buffer *Char, filename *int8)

// llgo:link (*Char).CreateDocParserCtxt C.xmlCreateDocParserCtxt
func (recv_ *Char) CreateDocParserCtxt() ParserCtxtPtr {
	return nil
}

/*
 * Reading/setting optional parsing features.
 */
//go:linkname GetFeaturesList C.xmlGetFeaturesList
func GetFeaturesList(len *c.Int, result **int8) c.Int

//go:linkname GetFeature C.xmlGetFeature
func GetFeature(ctxt ParserCtxtPtr, name *int8, result unsafe.Pointer) c.Int

//go:linkname SetFeature C.xmlSetFeature
func SetFeature(ctxt ParserCtxtPtr, name *int8, value unsafe.Pointer) c.Int

/*
 * Interfaces for the Push mode.
 */
//go:linkname CreatePushParserCtxt C.xmlCreatePushParserCtxt
func CreatePushParserCtxt(sax SAXHandlerPtr, user_data unsafe.Pointer, chunk *int8, size c.Int, filename *int8) ParserCtxtPtr

//go:linkname ParseChunk C.xmlParseChunk
func ParseChunk(ctxt ParserCtxtPtr, chunk *int8, size c.Int, terminate c.Int) c.Int

/*
 * Special I/O mode.
 */
//go:linkname CreateIOParserCtxt C.xmlCreateIOParserCtxt
func CreateIOParserCtxt(sax SAXHandlerPtr, user_data unsafe.Pointer, ioread InputReadCallback, ioclose InputCloseCallback, ioctx unsafe.Pointer, enc CharEncoding) ParserCtxtPtr

//go:linkname NewIOInputStream C.xmlNewIOInputStream
func NewIOInputStream(ctxt ParserCtxtPtr, input ParserInputBufferPtr, enc CharEncoding) ParserInputPtr

/*
 * Node infos.
 */
//go:linkname ParserFindNodeInfo C.xmlParserFindNodeInfo
func ParserFindNodeInfo(ctxt ParserCtxtPtr, node NodePtr) *ParserNodeInfo

//go:linkname InitNodeInfoSeq C.xmlInitNodeInfoSeq
func InitNodeInfoSeq(seq ParserNodeInfoSeqPtr)

//go:linkname ClearNodeInfoSeq C.xmlClearNodeInfoSeq
func ClearNodeInfoSeq(seq ParserNodeInfoSeqPtr)

//go:linkname ParserFindNodeInfoIndex C.xmlParserFindNodeInfoIndex
func ParserFindNodeInfoIndex(seq ParserNodeInfoSeqPtr, node NodePtr) c.Ulong

//go:linkname ParserAddNodeInfo C.xmlParserAddNodeInfo
func ParserAddNodeInfo(ctxt ParserCtxtPtr, info ParserNodeInfoPtr)

/*
 * External entities handling actually implemented in xmlIO.
 */
//go:linkname SetExternalEntityLoader C.xmlSetExternalEntityLoader
func SetExternalEntityLoader(f ExternalEntityLoader)

//go:linkname GetExternalEntityLoader C.xmlGetExternalEntityLoader
func GetExternalEntityLoader() ExternalEntityLoader

//go:linkname LoadExternalEntity C.xmlLoadExternalEntity
func LoadExternalEntity(URL *int8, ID *int8, ctxt ParserCtxtPtr) ParserInputPtr

/*
 * Index lookup, actually implemented in the encoding module
 */
//go:linkname ByteConsumed C.xmlByteConsumed
func ByteConsumed(ctxt ParserCtxtPtr) c.Long

type ParserOption c.Int

const (
	PARSERECOVER    ParserOption = 1
	PARSENOENT      ParserOption = 2
	PARSEDTDLOAD    ParserOption = 4
	PARSEDTDATTR    ParserOption = 8
	PARSEDTDVALID   ParserOption = 16
	PARSENOERROR    ParserOption = 32
	PARSENOWARNING  ParserOption = 64
	PARSEPEDANTIC   ParserOption = 128
	PARSENOBLANKS   ParserOption = 256
	PARSESAX1       ParserOption = 512
	PARSEXINCLUDE   ParserOption = 1024
	PARSENONET      ParserOption = 2048
	PARSENODICT     ParserOption = 4096
	PARSENSCLEAN    ParserOption = 8192
	PARSENOCDATA    ParserOption = 16384
	PARSENOXINCNODE ParserOption = 32768
	PARSECOMPACT    ParserOption = 65536
	PARSEOLD10      ParserOption = 131072
	PARSENOBASEFIX  ParserOption = 262144
	PARSEHUGE       ParserOption = 524288
	PARSEOLDSAX     ParserOption = 1048576
	PARSEIGNOREENC  ParserOption = 2097152
	PARSEBIGLINES   ParserOption = 4194304
	PARSENOXXE      ParserOption = 8388608
)

//go:linkname CtxtReset C.xmlCtxtReset
func CtxtReset(ctxt ParserCtxtPtr)

//go:linkname CtxtResetPush C.xmlCtxtResetPush
func CtxtResetPush(ctxt ParserCtxtPtr, chunk *int8, size c.Int, filename *int8, encoding *int8) c.Int

//go:linkname CtxtSetOptions C.xmlCtxtSetOptions
func CtxtSetOptions(ctxt ParserCtxtPtr, options c.Int) c.Int

//go:linkname CtxtUseOptions C.xmlCtxtUseOptions
func CtxtUseOptions(ctxt ParserCtxtPtr, options c.Int) c.Int

//go:linkname CtxtSetErrorHandler C.xmlCtxtSetErrorHandler
func CtxtSetErrorHandler(ctxt ParserCtxtPtr, handler StructuredErrorFunc, data unsafe.Pointer)

//go:linkname CtxtSetMaxAmplification C.xmlCtxtSetMaxAmplification
func CtxtSetMaxAmplification(ctxt ParserCtxtPtr, maxAmpl c.Uint)

// llgo:link (*Char).ReadDoc C.xmlReadDoc
func (recv_ *Char) ReadDoc(URL *int8, encoding *int8, options c.Int) DocPtr {
	return nil
}

//go:linkname ReadFile C.xmlReadFile
func ReadFile(URL *int8, encoding *int8, options c.Int) DocPtr

//go:linkname ReadMemory C.xmlReadMemory
func ReadMemory(buffer *int8, size c.Int, URL *int8, encoding *int8, options c.Int) DocPtr

//go:linkname ReadFd C.xmlReadFd
func ReadFd(fd c.Int, URL *int8, encoding *int8, options c.Int) DocPtr

//go:linkname ReadIO C.xmlReadIO
func ReadIO(ioread InputReadCallback, ioclose InputCloseCallback, ioctx unsafe.Pointer, URL *int8, encoding *int8, options c.Int) DocPtr

//go:linkname CtxtParseDocument C.xmlCtxtParseDocument
func CtxtParseDocument(ctxt ParserCtxtPtr, input ParserInputPtr) DocPtr

//go:linkname CtxtReadDoc C.xmlCtxtReadDoc
func CtxtReadDoc(ctxt ParserCtxtPtr, cur *Char, URL *int8, encoding *int8, options c.Int) DocPtr

//go:linkname CtxtReadFile C.xmlCtxtReadFile
func CtxtReadFile(ctxt ParserCtxtPtr, filename *int8, encoding *int8, options c.Int) DocPtr

//go:linkname CtxtReadMemory C.xmlCtxtReadMemory
func CtxtReadMemory(ctxt ParserCtxtPtr, buffer *int8, size c.Int, URL *int8, encoding *int8, options c.Int) DocPtr

//go:linkname CtxtReadFd C.xmlCtxtReadFd
func CtxtReadFd(ctxt ParserCtxtPtr, fd c.Int, URL *int8, encoding *int8, options c.Int) DocPtr

//go:linkname CtxtReadIO C.xmlCtxtReadIO
func CtxtReadIO(ctxt ParserCtxtPtr, ioread InputReadCallback, ioclose InputCloseCallback, ioctx unsafe.Pointer, URL *int8, encoding *int8, options c.Int) DocPtr

type Feature c.Int

const (
	WITHTHREAD     Feature = 1
	WITHTREE       Feature = 2
	WITHOUTPUT     Feature = 3
	WITHPUSH       Feature = 4
	WITHREADER     Feature = 5
	WITHPATTERN    Feature = 6
	WITHWRITER     Feature = 7
	WITHSAX1       Feature = 8
	WITHFTP        Feature = 9
	WITHHTTP       Feature = 10
	WITHVALID      Feature = 11
	WITHHTML       Feature = 12
	WITHLEGACY     Feature = 13
	WITHC14N       Feature = 14
	WITHCATALOG    Feature = 15
	WITHXPATH      Feature = 16
	WITHXPTR       Feature = 17
	WITHXINCLUDE   Feature = 18
	WITHICONV      Feature = 19
	WITHISO8859X   Feature = 20
	WITHUNICODE    Feature = 21
	WITHREGEXP     Feature = 22
	WITHAUTOMATA   Feature = 23
	WITHEXPR       Feature = 24
	WITHSCHEMAS    Feature = 25
	WITHSCHEMATRON Feature = 26
	WITHMODULES    Feature = 27
	WITHDEBUG      Feature = 28
	WITHDEBUGMEM   Feature = 29
	WITHDEBUGRUN   Feature = 30
	WITHZLIB       Feature = 31
	WITHICU        Feature = 32
	WITHLZMA       Feature = 33
	WITHNONE       Feature = 99999
)

// llgo:link Feature.HasFeature C.xmlHasFeature
func (recv_ Feature) HasFeature() c.Int {
	return 0
}
