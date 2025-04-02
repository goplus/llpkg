package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type CatalogPrefer c.Int

const (
	CATAPREFERNONE   CatalogPrefer = 0
	CATAPREFERPUBLIC CatalogPrefer = 1
	CATAPREFERSYSTEM CatalogPrefer = 2
)

type CatalogAllow c.Int

const (
	CATAALLOWNONE     CatalogAllow = 0
	CATAALLOWGLOBAL   CatalogAllow = 1
	CATAALLOWDOCUMENT CatalogAllow = 2
	CATAALLOWALL      CatalogAllow = 3
)

type X_xmlCatalog struct {
	Unused [8]uint8
}
type Catalog X_xmlCatalog
type CatalogPtr *Catalog

/*
 * Operations on a given catalog.
 */
//go:linkname NewCatalog C.xmlNewCatalog
func NewCatalog(sgml c.Int) CatalogPtr

//go:linkname LoadACatalog C.xmlLoadACatalog
func LoadACatalog(filename *int8) CatalogPtr

//go:linkname LoadSGMLSuperCatalog C.xmlLoadSGMLSuperCatalog
func LoadSGMLSuperCatalog(filename *int8) CatalogPtr

//go:linkname ConvertSGMLCatalog C.xmlConvertSGMLCatalog
func ConvertSGMLCatalog(catal CatalogPtr) c.Int

//go:linkname ACatalogAdd C.xmlACatalogAdd
func ACatalogAdd(catal CatalogPtr, type_ *Char, orig *Char, replace *Char) c.Int

//go:linkname ACatalogRemove C.xmlACatalogRemove
func ACatalogRemove(catal CatalogPtr, value *Char) c.Int

//go:linkname ACatalogResolve C.xmlACatalogResolve
func ACatalogResolve(catal CatalogPtr, pubID *Char, sysID *Char) *Char

//go:linkname ACatalogResolveSystem C.xmlACatalogResolveSystem
func ACatalogResolveSystem(catal CatalogPtr, sysID *Char) *Char

//go:linkname ACatalogResolvePublic C.xmlACatalogResolvePublic
func ACatalogResolvePublic(catal CatalogPtr, pubID *Char) *Char

//go:linkname ACatalogResolveURI C.xmlACatalogResolveURI
func ACatalogResolveURI(catal CatalogPtr, URI *Char) *Char

//go:linkname ACatalogDump C.xmlACatalogDump
func ACatalogDump(catal CatalogPtr, out *c.FILE)

//go:linkname FreeCatalog C.xmlFreeCatalog
func FreeCatalog(catal CatalogPtr)

//go:linkname CatalogIsEmpty C.xmlCatalogIsEmpty
func CatalogIsEmpty(catal CatalogPtr) c.Int

/*
 * Global operations.
 */
//go:linkname InitializeCatalog C.xmlInitializeCatalog
func InitializeCatalog()

//go:linkname LoadCatalog C.xmlLoadCatalog
func LoadCatalog(filename *int8) c.Int

//go:linkname LoadCatalogs C.xmlLoadCatalogs
func LoadCatalogs(paths *int8)

//go:linkname CatalogCleanup C.xmlCatalogCleanup
func CatalogCleanup()

//go:linkname CatalogDump C.xmlCatalogDump
func CatalogDump(out *c.FILE)

// llgo:link (*Char).CatalogResolve C.xmlCatalogResolve
func (recv_ *Char) CatalogResolve(sysID *Char) *Char {
	return nil
}

// llgo:link (*Char).CatalogResolveSystem C.xmlCatalogResolveSystem
func (recv_ *Char) CatalogResolveSystem() *Char {
	return nil
}

// llgo:link (*Char).CatalogResolvePublic C.xmlCatalogResolvePublic
func (recv_ *Char) CatalogResolvePublic() *Char {
	return nil
}

// llgo:link (*Char).CatalogResolveURI C.xmlCatalogResolveURI
func (recv_ *Char) CatalogResolveURI() *Char {
	return nil
}

// llgo:link (*Char).CatalogAdd C.xmlCatalogAdd
func (recv_ *Char) CatalogAdd(orig *Char, replace *Char) c.Int {
	return 0
}

// llgo:link (*Char).CatalogRemove C.xmlCatalogRemove
func (recv_ *Char) CatalogRemove() c.Int {
	return 0
}

//go:linkname ParseCatalogFile C.xmlParseCatalogFile
func ParseCatalogFile(filename *int8) DocPtr

//go:linkname CatalogConvert C.xmlCatalogConvert
func CatalogConvert() c.Int

/*
 * Strictly minimal interfaces for per-document catalogs used
 * by the parser.
 */
//go:linkname CatalogFreeLocal C.xmlCatalogFreeLocal
func CatalogFreeLocal(catalogs unsafe.Pointer)

//go:linkname CatalogAddLocal C.xmlCatalogAddLocal
func CatalogAddLocal(catalogs unsafe.Pointer, URL *Char) unsafe.Pointer

//go:linkname CatalogLocalResolve C.xmlCatalogLocalResolve
func CatalogLocalResolve(catalogs unsafe.Pointer, pubID *Char, sysID *Char) *Char

//go:linkname CatalogLocalResolveURI C.xmlCatalogLocalResolveURI
func CatalogLocalResolveURI(catalogs unsafe.Pointer, URI *Char) *Char

/*
 * Preference settings.
 */
//go:linkname CatalogSetDebug C.xmlCatalogSetDebug
func CatalogSetDebug(level c.Int) c.Int

// llgo:link CatalogPrefer.CatalogSetDefaultPrefer C.xmlCatalogSetDefaultPrefer
func (recv_ CatalogPrefer) CatalogSetDefaultPrefer() CatalogPrefer {
	return 0
}

// llgo:link CatalogAllow.CatalogSetDefaults C.xmlCatalogSetDefaults
func (recv_ CatalogAllow) CatalogSetDefaults() {
}

//go:linkname CatalogGetDefaults C.xmlCatalogGetDefaults
func CatalogGetDefaults() CatalogAllow

/* DEPRECATED interfaces */
// llgo:link (*Char).CatalogGetSystem C.xmlCatalogGetSystem
func (recv_ *Char) CatalogGetSystem() *Char {
	return nil
}

// llgo:link (*Char).CatalogGetPublic C.xmlCatalogGetPublic
func (recv_ *Char) CatalogGetPublic() *Char {
	return nil
}
