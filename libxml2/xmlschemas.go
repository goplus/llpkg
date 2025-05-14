package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type SchemaValidError c.Int

const (
	SCHEMASERROK             SchemaValidError = 0
	SCHEMASERRNOROOT         SchemaValidError = 1
	SCHEMASERRUNDECLAREDELEM SchemaValidError = 2
	SCHEMASERRNOTTOPLEVEL    SchemaValidError = 3
	SCHEMASERRMISSING        SchemaValidError = 4
	SCHEMASERRWRONGELEM      SchemaValidError = 5
	SCHEMASERRNOTYPE         SchemaValidError = 6
	SCHEMASERRNOROLLBACK     SchemaValidError = 7
	SCHEMASERRISABSTRACT     SchemaValidError = 8
	SCHEMASERRNOTEMPTY       SchemaValidError = 9
	SCHEMASERRELEMCONT       SchemaValidError = 10
	SCHEMASERRHAVEDEFAULT    SchemaValidError = 11
	SCHEMASERRNOTNILLABLE    SchemaValidError = 12
	SCHEMASERREXTRACONTENT   SchemaValidError = 13
	SCHEMASERRINVALIDATTR    SchemaValidError = 14
	SCHEMASERRINVALIDELEM    SchemaValidError = 15
	SCHEMASERRNOTDETERMINIST SchemaValidError = 16
	SCHEMASERRCONSTRUCT      SchemaValidError = 17
	SCHEMASERRINTERNAL       SchemaValidError = 18
	SCHEMASERRNOTSIMPLE      SchemaValidError = 19
	SCHEMASERRATTRUNKNOWN    SchemaValidError = 20
	SCHEMASERRATTRINVALID    SchemaValidError = 21
	SCHEMASERRVALUE          SchemaValidError = 22
	SCHEMASERRFACET          SchemaValidError = 23
	SCHEMASERR_              SchemaValidError = 24
	SCHEMASERRXXX            SchemaValidError = 25
)

type SchemaValidOption c.Int

const SCHEMAVALVCICREATE SchemaValidOption = 1

type Schema X_xmlSchema
type SchemaPtr *Schema

// llgo:type C
type SchemaValidityErrorFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})

// llgo:type C
type SchemaValidityWarningFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})

type X_xmlSchemaParserCtxt struct {
	Unused [8]uint8
}
type SchemaParserCtxt X_xmlSchemaParserCtxt
type SchemaParserCtxtPtr *SchemaParserCtxt

type X_xmlSchemaValidCtxt struct {
	Unused [8]uint8
}
type SchemaValidCtxt X_xmlSchemaValidCtxt
type SchemaValidCtxtPtr *SchemaValidCtxt

// llgo:type C
type SchemaValidityLocatorFunc func(unsafe.Pointer, **int8, *c.Ulong) c.Int

/*
 * Interfaces for parsing.
 */
//go:linkname SchemaNewParserCtxt C.xmlSchemaNewParserCtxt
func SchemaNewParserCtxt(URL *int8) SchemaParserCtxtPtr

//go:linkname SchemaNewMemParserCtxt C.xmlSchemaNewMemParserCtxt
func SchemaNewMemParserCtxt(buffer *int8, size c.Int) SchemaParserCtxtPtr

//go:linkname SchemaNewDocParserCtxt C.xmlSchemaNewDocParserCtxt
func SchemaNewDocParserCtxt(doc DocPtr) SchemaParserCtxtPtr

//go:linkname SchemaFreeParserCtxt C.xmlSchemaFreeParserCtxt
func SchemaFreeParserCtxt(ctxt SchemaParserCtxtPtr)

//go:linkname SchemaSetParserErrors C.xmlSchemaSetParserErrors
func SchemaSetParserErrors(ctxt SchemaParserCtxtPtr, err SchemaValidityErrorFunc, warn SchemaValidityWarningFunc, ctx unsafe.Pointer)

//go:linkname SchemaSetParserStructuredErrors C.xmlSchemaSetParserStructuredErrors
func SchemaSetParserStructuredErrors(ctxt SchemaParserCtxtPtr, serror StructuredErrorFunc, ctx unsafe.Pointer)

//go:linkname SchemaGetParserErrors C.xmlSchemaGetParserErrors
func SchemaGetParserErrors(ctxt SchemaParserCtxtPtr, err SchemaValidityErrorFunc, warn SchemaValidityWarningFunc, ctx *unsafe.Pointer) c.Int

//go:linkname SchemaIsValid C.xmlSchemaIsValid
func SchemaIsValid(ctxt SchemaValidCtxtPtr) c.Int

//go:linkname SchemaParse C.xmlSchemaParse
func SchemaParse(ctxt SchemaParserCtxtPtr) SchemaPtr

//go:linkname SchemaFree C.xmlSchemaFree
func SchemaFree(schema SchemaPtr)

//go:linkname SchemaDump C.xmlSchemaDump
func SchemaDump(output *c.FILE, schema SchemaPtr)

/*
 * Interfaces for validating
 */
//go:linkname SchemaSetValidErrors C.xmlSchemaSetValidErrors
func SchemaSetValidErrors(ctxt SchemaValidCtxtPtr, err SchemaValidityErrorFunc, warn SchemaValidityWarningFunc, ctx unsafe.Pointer)

//go:linkname SchemaSetValidStructuredErrors C.xmlSchemaSetValidStructuredErrors
func SchemaSetValidStructuredErrors(ctxt SchemaValidCtxtPtr, serror StructuredErrorFunc, ctx unsafe.Pointer)

//go:linkname SchemaGetValidErrors C.xmlSchemaGetValidErrors
func SchemaGetValidErrors(ctxt SchemaValidCtxtPtr, err SchemaValidityErrorFunc, warn SchemaValidityWarningFunc, ctx *unsafe.Pointer) c.Int

//go:linkname SchemaSetValidOptions C.xmlSchemaSetValidOptions
func SchemaSetValidOptions(ctxt SchemaValidCtxtPtr, options c.Int) c.Int

//go:linkname SchemaValidateSetFilename C.xmlSchemaValidateSetFilename
func SchemaValidateSetFilename(vctxt SchemaValidCtxtPtr, filename *int8)

//go:linkname SchemaValidCtxtGetOptions C.xmlSchemaValidCtxtGetOptions
func SchemaValidCtxtGetOptions(ctxt SchemaValidCtxtPtr) c.Int

//go:linkname SchemaNewValidCtxt C.xmlSchemaNewValidCtxt
func SchemaNewValidCtxt(schema SchemaPtr) SchemaValidCtxtPtr

//go:linkname SchemaFreeValidCtxt C.xmlSchemaFreeValidCtxt
func SchemaFreeValidCtxt(ctxt SchemaValidCtxtPtr)

//go:linkname SchemaValidateDoc C.xmlSchemaValidateDoc
func SchemaValidateDoc(ctxt SchemaValidCtxtPtr, instance DocPtr) c.Int

//go:linkname SchemaValidateOneElement C.xmlSchemaValidateOneElement
func SchemaValidateOneElement(ctxt SchemaValidCtxtPtr, elem NodePtr) c.Int

//go:linkname SchemaValidateStream C.xmlSchemaValidateStream
func SchemaValidateStream(ctxt SchemaValidCtxtPtr, input ParserInputBufferPtr, enc CharEncoding, sax SAXHandlerPtr, user_data unsafe.Pointer) c.Int

//go:linkname SchemaValidateFile C.xmlSchemaValidateFile
func SchemaValidateFile(ctxt SchemaValidCtxtPtr, filename *int8, options c.Int) c.Int

//go:linkname SchemaValidCtxtGetParserCtxt C.xmlSchemaValidCtxtGetParserCtxt
func SchemaValidCtxtGetParserCtxt(ctxt SchemaValidCtxtPtr) ParserCtxtPtr

type X_xmlSchemaSAXPlug struct {
	Unused [8]uint8
}
type SchemaSAXPlugStruct X_xmlSchemaSAXPlug
type SchemaSAXPlugPtr *SchemaSAXPlugStruct

//go:linkname SchemaSAXPlug C.xmlSchemaSAXPlug
func SchemaSAXPlug(ctxt SchemaValidCtxtPtr, sax *SAXHandlerPtr, user_data *unsafe.Pointer) SchemaSAXPlugPtr

//go:linkname SchemaSAXUnplug C.xmlSchemaSAXUnplug
func SchemaSAXUnplug(plug SchemaSAXPlugPtr) c.Int

//go:linkname SchemaValidateSetLocator C.xmlSchemaValidateSetLocator
func SchemaValidateSetLocator(vctxt SchemaValidCtxtPtr, f SchemaValidityLocatorFunc, ctxt unsafe.Pointer)
