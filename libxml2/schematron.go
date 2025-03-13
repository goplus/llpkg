package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type SchematronValidOptions c.Int

const (
	SCHEMATRONOUTQUIET  SchematronValidOptions = 1
	SCHEMATRONOUTTEXT   SchematronValidOptions = 2
	SCHEMATRONOUTXML    SchematronValidOptions = 4
	SCHEMATRONOUTERROR  SchematronValidOptions = 8
	SCHEMATRONOUTFILE   SchematronValidOptions = 256
	SCHEMATRONOUTBUFFER SchematronValidOptions = 512
	SCHEMATRONOUTIO     SchematronValidOptions = 1024
)

type X_xmlSchematron struct {
	Unused [8]uint8
}
type Schematron X_xmlSchematron
type SchematronPtr *Schematron

// llgo:type C
type SchematronValidityErrorFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})

// llgo:type C
type SchematronValidityWarningFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})

type X_xmlSchematronParserCtxt struct {
	Unused [8]uint8
}
type SchematronParserCtxt X_xmlSchematronParserCtxt
type SchematronParserCtxtPtr *SchematronParserCtxt

type X_xmlSchematronValidCtxt struct {
	Unused [8]uint8
}
type SchematronValidCtxt X_xmlSchematronValidCtxt
type SchematronValidCtxtPtr *SchematronValidCtxt

/*
 * Interfaces for parsing.
 */
//go:linkname SchematronNewParserCtxt C.xmlSchematronNewParserCtxt
func SchematronNewParserCtxt(URL *int8) SchematronParserCtxtPtr

//go:linkname SchematronNewMemParserCtxt C.xmlSchematronNewMemParserCtxt
func SchematronNewMemParserCtxt(buffer *int8, size c.Int) SchematronParserCtxtPtr

//go:linkname SchematronNewDocParserCtxt C.xmlSchematronNewDocParserCtxt
func SchematronNewDocParserCtxt(doc DocPtr) SchematronParserCtxtPtr

//go:linkname SchematronFreeParserCtxt C.xmlSchematronFreeParserCtxt
func SchematronFreeParserCtxt(ctxt SchematronParserCtxtPtr)

/*****
XMLPUBFUN void
	    xmlSchematronSetParserErrors(xmlSchematronParserCtxtPtr ctxt,
					 xmlSchematronValidityErrorFunc err,
					 xmlSchematronValidityWarningFunc warn,
					 void *ctx);
XMLPUBFUN int
		xmlSchematronGetParserErrors(xmlSchematronParserCtxtPtr ctxt,
					xmlSchematronValidityErrorFunc * err,
					xmlSchematronValidityWarningFunc * warn,
					void **ctx);
XMLPUBFUN int
		xmlSchematronIsValid	(xmlSchematronValidCtxtPtr ctxt);
 *****/
//go:linkname SchematronParse C.xmlSchematronParse
func SchematronParse(ctxt SchematronParserCtxtPtr) SchematronPtr

//go:linkname SchematronFree C.xmlSchematronFree
func SchematronFree(schema SchematronPtr)

/*
 * Interfaces for validating
 */
//go:linkname SchematronSetValidStructuredErrors C.xmlSchematronSetValidStructuredErrors
func SchematronSetValidStructuredErrors(ctxt SchematronValidCtxtPtr, serror StructuredErrorFunc, ctx unsafe.Pointer)

/******
XMLPUBFUN void
	    xmlSchematronSetValidErrors	(xmlSchematronValidCtxtPtr ctxt,
					 xmlSchematronValidityErrorFunc err,
					 xmlSchematronValidityWarningFunc warn,
					 void *ctx);
XMLPUBFUN int
	    xmlSchematronGetValidErrors	(xmlSchematronValidCtxtPtr ctxt,
					 xmlSchematronValidityErrorFunc *err,
					 xmlSchematronValidityWarningFunc *warn,
					 void **ctx);
XMLPUBFUN int
	    xmlSchematronSetValidOptions(xmlSchematronValidCtxtPtr ctxt,
					 int options);
XMLPUBFUN int
	    xmlSchematronValidCtxtGetOptions(xmlSchematronValidCtxtPtr ctxt);
XMLPUBFUN int
            xmlSchematronValidateOneElement (xmlSchematronValidCtxtPtr ctxt,
			                 xmlNodePtr elem);
 *******/
//go:linkname SchematronNewValidCtxt C.xmlSchematronNewValidCtxt
func SchematronNewValidCtxt(schema SchematronPtr, options c.Int) SchematronValidCtxtPtr

//go:linkname SchematronFreeValidCtxt C.xmlSchematronFreeValidCtxt
func SchematronFreeValidCtxt(ctxt SchematronValidCtxtPtr)

//go:linkname SchematronValidateDoc C.xmlSchematronValidateDoc
func SchematronValidateDoc(ctxt SchematronValidCtxtPtr, instance DocPtr) c.Int
