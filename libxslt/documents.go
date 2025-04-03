package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/libxml2"
	"unsafe"
)

//go:linkname NewDocument C.xsltNewDocument
func NewDocument(ctxt TransformContextPtr, doc libxml2.DocPtr) DocumentPtr

//go:linkname LoadDocument C.xsltLoadDocument
func LoadDocument(ctxt TransformContextPtr, URI *libxml2.Char) DocumentPtr

//go:linkname FindDocument C.xsltFindDocument
func FindDocument(ctxt TransformContextPtr, doc libxml2.DocPtr) DocumentPtr

//go:linkname FreeDocuments C.xsltFreeDocuments
func FreeDocuments(ctxt TransformContextPtr)

//go:linkname LoadStyleDocument C.xsltLoadStyleDocument
func LoadStyleDocument(style StylesheetPtr, URI *libxml2.Char) DocumentPtr

//go:linkname NewStyleDocument C.xsltNewStyleDocument
func NewStyleDocument(style StylesheetPtr, doc libxml2.DocPtr) DocumentPtr

//go:linkname FreeStyleDocuments C.xsltFreeStyleDocuments
func FreeStyleDocuments(style StylesheetPtr)

type LoadType c.Int

const (
	LOADSTART      LoadType = 0
	LOADSTYLESHEET LoadType = 1
	LOADDOCUMENT   LoadType = 2
)

// llgo:type C
type DocLoaderFunc func(*libxml2.Char, libxml2.DictPtr, c.Int, unsafe.Pointer, LoadType) libxml2.DocPtr

//go:linkname SetLoaderFunc C.xsltSetLoaderFunc
func SetLoaderFunc(f DocLoaderFunc)
