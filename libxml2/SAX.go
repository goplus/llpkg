package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

//go:linkname GetPublicId C.getPublicId
func GetPublicId(ctx unsafe.Pointer) *Char

//go:linkname GetSystemId C.getSystemId
func GetSystemId(ctx unsafe.Pointer) *Char

//go:linkname SetDocumentLocator C.setDocumentLocator
func SetDocumentLocator(ctx unsafe.Pointer, loc SAXLocatorPtr)

//go:linkname GetLineNumber C.getLineNumber
func GetLineNumber(ctx unsafe.Pointer) c.Int

//go:linkname GetColumnNumber C.getColumnNumber
func GetColumnNumber(ctx unsafe.Pointer) c.Int

//go:linkname IsStandalone C.isStandalone
func IsStandalone(ctx unsafe.Pointer) c.Int

//go:linkname HasInternalSubset C.hasInternalSubset
func HasInternalSubset(ctx unsafe.Pointer) c.Int

//go:linkname HasExternalSubset C.hasExternalSubset
func HasExternalSubset(ctx unsafe.Pointer) c.Int

//go:linkname InternalSubset C.internalSubset
func InternalSubset(ctx unsafe.Pointer, name *Char, ExternalID *Char, SystemID *Char)

//go:linkname ExternalSubset C.externalSubset
func ExternalSubset(ctx unsafe.Pointer, name *Char, ExternalID *Char, SystemID *Char)

//go:linkname GetEntity C.getEntity
func GetEntity(ctx unsafe.Pointer, name *Char) EntityPtr

//go:linkname GetParameterEntity__1 C.getParameterEntity
func GetParameterEntity__1(ctx unsafe.Pointer, name *Char) EntityPtr

//go:linkname ResolveEntity C.resolveEntity
func ResolveEntity(ctx unsafe.Pointer, publicId *Char, systemId *Char) ParserInputPtr

//go:linkname EntityDecl C.entityDecl
func EntityDecl(ctx unsafe.Pointer, name *Char, type_ c.Int, publicId *Char, systemId *Char, content *Char)

//go:linkname AttributeDecl C.attributeDecl
func AttributeDecl(ctx unsafe.Pointer, elem *Char, fullname *Char, type_ c.Int, def c.Int, defaultValue *Char, tree EnumerationPtr)

//go:linkname ElementDecl C.elementDecl
func ElementDecl(ctx unsafe.Pointer, name *Char, type_ c.Int, content ElementContentPtr)

//go:linkname NotationDecl C.notationDecl
func NotationDecl(ctx unsafe.Pointer, name *Char, publicId *Char, systemId *Char)

//go:linkname UnparsedEntityDecl C.unparsedEntityDecl
func UnparsedEntityDecl(ctx unsafe.Pointer, name *Char, publicId *Char, systemId *Char, notationName *Char)

//go:linkname StartDocument C.startDocument
func StartDocument(ctx unsafe.Pointer)

//go:linkname EndDocument C.endDocument
func EndDocument(ctx unsafe.Pointer)

//go:linkname GetAttribute C.attribute
func GetAttribute(ctx unsafe.Pointer, fullname *Char, value *Char)

//go:linkname StartElement C.startElement
func StartElement(ctx unsafe.Pointer, fullname *Char, atts **Char)

//go:linkname EndElement C.endElement
func EndElement(ctx unsafe.Pointer, name *Char)

//go:linkname Reference C.reference
func Reference(ctx unsafe.Pointer, name *Char)

//go:linkname Characters C.characters
func Characters(ctx unsafe.Pointer, ch *Char, len c.Int)

//go:linkname IgnorableWhitespace C.ignorableWhitespace
func IgnorableWhitespace(ctx unsafe.Pointer, ch *Char, len c.Int)

//go:linkname ProcessingInstruction C.processingInstruction
func ProcessingInstruction(ctx unsafe.Pointer, target *Char, data *Char)

//go:linkname GlobalNamespace C.globalNamespace
func GlobalNamespace(ctx unsafe.Pointer, href *Char, prefix *Char)

//go:linkname SetNamespace C.setNamespace
func SetNamespace(ctx unsafe.Pointer, name *Char)

//go:linkname GetNamespace C.getNamespace
func GetNamespace(ctx unsafe.Pointer) NsPtr

//go:linkname CheckNamespace C.checkNamespace
func CheckNamespace(ctx unsafe.Pointer, nameSpace *Char) c.Int

//go:linkname NamespaceDecl C.namespaceDecl
func NamespaceDecl(ctx unsafe.Pointer, href *Char, prefix *Char)

//go:linkname Comment C.comment
func Comment(ctx unsafe.Pointer, value *Char)

//go:linkname CdataBlock C.cdataBlock
func CdataBlock(ctx unsafe.Pointer, value *Char, len c.Int)

// llgo:link (*SAXHandlerV1).InitxmlDefaultSAXHandler C.initxmlDefaultSAXHandler
func (recv_ *SAXHandlerV1) InitxmlDefaultSAXHandler(warning c.Int) {
}

// llgo:link (*SAXHandlerV1).InithtmlDefaultSAXHandler C.inithtmlDefaultSAXHandler
func (recv_ *SAXHandlerV1) InithtmlDefaultSAXHandler() {
}
