package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const SCHEMAS_ANYATTR_SKIP = 1
const SCHEMAS_ANYATTR_LAX = 2
const SCHEMAS_ANYATTR_STRICT = 3
const SCHEMAS_ANY_SKIP = 1
const SCHEMAS_ANY_LAX = 2
const SCHEMAS_ANY_STRICT = 3
const SCHEMAS_ATTR_USE_PROHIBITED = 0
const SCHEMAS_ATTR_USE_REQUIRED = 1
const SCHEMAS_ATTR_USE_OPTIONAL = 2
const SCHEMAS_FACET_UNKNOWN = 0
const SCHEMAS_FACET_PRESERVE = 1
const SCHEMAS_FACET_REPLACE = 2
const SCHEMAS_FACET_COLLAPSE = 3

type SchemaValType c.Int

const (
	SCHEMASUNKNOWN       SchemaValType = 0
	SCHEMASSTRING        SchemaValType = 1
	SCHEMASNORMSTRING    SchemaValType = 2
	SCHEMASDECIMAL       SchemaValType = 3
	SCHEMASTIME          SchemaValType = 4
	SCHEMASGDAY          SchemaValType = 5
	SCHEMASGMONTH        SchemaValType = 6
	SCHEMASGMONTHDAY     SchemaValType = 7
	SCHEMASGYEAR         SchemaValType = 8
	SCHEMASGYEARMONTH    SchemaValType = 9
	SCHEMASDATE          SchemaValType = 10
	SCHEMASDATETIME      SchemaValType = 11
	SCHEMASDURATION      SchemaValType = 12
	SCHEMASFLOAT         SchemaValType = 13
	SCHEMASDOUBLE        SchemaValType = 14
	SCHEMASBOOLEAN       SchemaValType = 15
	SCHEMASTOKEN         SchemaValType = 16
	SCHEMASLANGUAGE      SchemaValType = 17
	SCHEMASNMTOKEN       SchemaValType = 18
	SCHEMASNMTOKENS      SchemaValType = 19
	SCHEMASNAME          SchemaValType = 20
	SCHEMASQNAME         SchemaValType = 21
	SCHEMASNCNAME        SchemaValType = 22
	SCHEMASID            SchemaValType = 23
	SCHEMASIDREF         SchemaValType = 24
	SCHEMASIDREFS        SchemaValType = 25
	SCHEMASENTITY        SchemaValType = 26
	SCHEMASENTITIES      SchemaValType = 27
	SCHEMASNOTATION      SchemaValType = 28
	SCHEMASANYURI        SchemaValType = 29
	SCHEMASINTEGER       SchemaValType = 30
	SCHEMASNPINTEGER     SchemaValType = 31
	SCHEMASNINTEGER      SchemaValType = 32
	SCHEMASNNINTEGER     SchemaValType = 33
	SCHEMASPINTEGER      SchemaValType = 34
	SCHEMASINT           SchemaValType = 35
	SCHEMASUINT          SchemaValType = 36
	SCHEMASLONG          SchemaValType = 37
	SCHEMASULONG         SchemaValType = 38
	SCHEMASSHORT         SchemaValType = 39
	SCHEMASUSHORT        SchemaValType = 40
	SCHEMASBYTE          SchemaValType = 41
	SCHEMASUBYTE         SchemaValType = 42
	SCHEMASHEXBINARY     SchemaValType = 43
	SCHEMASBASE64BINARY  SchemaValType = 44
	SCHEMASANYTYPE       SchemaValType = 45
	SCHEMASANYSIMPLETYPE SchemaValType = 46
)

type SchemaTypeType c.Int

const (
	SCHEMATYPEBASIC           SchemaTypeType = 1
	SCHEMATYPEANY             SchemaTypeType = 2
	SCHEMATYPEFACET           SchemaTypeType = 3
	SCHEMATYPESIMPLE          SchemaTypeType = 4
	SCHEMATYPECOMPLEX         SchemaTypeType = 5
	SCHEMATYPESEQUENCE        SchemaTypeType = 6
	SCHEMATYPECHOICE          SchemaTypeType = 7
	SCHEMATYPEALL             SchemaTypeType = 8
	SCHEMATYPESIMPLECONTENT   SchemaTypeType = 9
	SCHEMATYPECOMPLEXCONTENT  SchemaTypeType = 10
	SCHEMATYPEUR              SchemaTypeType = 11
	SCHEMATYPERESTRICTION     SchemaTypeType = 12
	SCHEMATYPEEXTENSION       SchemaTypeType = 13
	SCHEMATYPEELEMENT         SchemaTypeType = 14
	SCHEMATYPEATTRIBUTE       SchemaTypeType = 15
	SCHEMATYPEATTRIBUTEGROUP  SchemaTypeType = 16
	SCHEMATYPEGROUP           SchemaTypeType = 17
	SCHEMATYPENOTATION        SchemaTypeType = 18
	SCHEMATYPELIST            SchemaTypeType = 19
	SCHEMATYPEUNION           SchemaTypeType = 20
	SCHEMATYPEANYATTRIBUTE    SchemaTypeType = 21
	SCHEMATYPEIDCUNIQUE       SchemaTypeType = 22
	SCHEMATYPEIDCKEY          SchemaTypeType = 23
	SCHEMATYPEIDCKEYREF       SchemaTypeType = 24
	SCHEMATYPEPARTICLE        SchemaTypeType = 25
	SCHEMATYPEATTRIBUTEUSE    SchemaTypeType = 26
	SCHEMAFACETMININCLUSIVE   SchemaTypeType = 1000
	SCHEMAFACETMINEXCLUSIVE   SchemaTypeType = 1001
	SCHEMAFACETMAXINCLUSIVE   SchemaTypeType = 1002
	SCHEMAFACETMAXEXCLUSIVE   SchemaTypeType = 1003
	SCHEMAFACETTOTALDIGITS    SchemaTypeType = 1004
	SCHEMAFACETFRACTIONDIGITS SchemaTypeType = 1005
	SCHEMAFACETPATTERN        SchemaTypeType = 1006
	SCHEMAFACETENUMERATION    SchemaTypeType = 1007
	SCHEMAFACETWHITESPACE     SchemaTypeType = 1008
	SCHEMAFACETLENGTH         SchemaTypeType = 1009
	SCHEMAFACETMAXLENGTH      SchemaTypeType = 1010
	SCHEMAFACETMINLENGTH      SchemaTypeType = 1011
	SCHEMAEXTRAQNAMEREF       SchemaTypeType = 2000
	SCHEMAEXTRAATTRUSEPROHIB  SchemaTypeType = 2001
)

type SchemaContentType c.Int

const (
	SCHEMACONTENTUNKNOWN         SchemaContentType = 0
	SCHEMACONTENTEMPTY           SchemaContentType = 1
	SCHEMACONTENTELEMENTS        SchemaContentType = 2
	SCHEMACONTENTMIXED           SchemaContentType = 3
	SCHEMACONTENTSIMPLE          SchemaContentType = 4
	SCHEMACONTENTMIXEDORELEMENTS SchemaContentType = 5
	SCHEMACONTENTBASIC           SchemaContentType = 6
	SCHEMACONTENTANY             SchemaContentType = 7
)

type X_xmlSchemaVal struct {
	Unused [8]uint8
}
type SchemaVal X_xmlSchemaVal
type SchemaValPtr *SchemaVal

type X_xmlSchemaType struct {
	Type              SchemaTypeType
	Next              *X_xmlSchemaType
	Name              *Char
	Id                *Char
	Ref               *Char
	RefNs             *Char
	Annot             SchemaAnnotPtr
	Subtypes          SchemaTypePtr
	Attributes        SchemaAttributePtr
	Node              NodePtr
	MinOccurs         c.Int
	MaxOccurs         c.Int
	Flags             c.Int
	ContentType       SchemaContentType
	Base              *Char
	BaseNs            *Char
	BaseType          SchemaTypePtr
	Facets            SchemaFacetPtr
	Redef             *X_xmlSchemaType
	Recurse           c.Int
	AttributeUses     *SchemaAttributeLinkPtr
	AttributeWildcard SchemaWildcardPtr
	BuiltInType       c.Int
	MemberTypes       SchemaTypeLinkPtr
	FacetSet          SchemaFacetLinkPtr
	RefPrefix         *Char
	ContentTypeDef    SchemaTypePtr
	ContModel         RegexpPtr
	TargetNamespace   *Char
	AttrUses          unsafe.Pointer
}
type SchemaType X_xmlSchemaType
type SchemaTypePtr *SchemaType

type X_xmlSchemaFacet struct {
	Type       SchemaTypeType
	Next       *X_xmlSchemaFacet
	Value      *Char
	Id         *Char
	Annot      SchemaAnnotPtr
	Node       NodePtr
	Fixed      c.Int
	Whitespace c.Int
	Val        SchemaValPtr
	Regexp     RegexpPtr
}
type SchemaFacet X_xmlSchemaFacet
type SchemaFacetPtr *SchemaFacet

type X_xmlSchemaAnnot struct {
	Next    *X_xmlSchemaAnnot
	Content NodePtr
}
type SchemaAnnot X_xmlSchemaAnnot
type SchemaAnnotPtr *SchemaAnnot

type X_xmlSchemaAttribute struct {
	Type            SchemaTypeType
	Next            *X_xmlSchemaAttribute
	Name            *Char
	Id              *Char
	Ref             *Char
	RefNs           *Char
	TypeName        *Char
	TypeNs          *Char
	Annot           SchemaAnnotPtr
	Base            SchemaTypePtr
	Occurs          c.Int
	DefValue        *Char
	Subtypes        SchemaTypePtr
	Node            NodePtr
	TargetNamespace *Char
	Flags           c.Int
	RefPrefix       *Char
	DefVal          SchemaValPtr
	RefDecl         SchemaAttributePtr
}
type SchemaAttribute X_xmlSchemaAttribute
type SchemaAttributePtr *SchemaAttribute

type X_xmlSchemaAttributeLink struct {
	Next *X_xmlSchemaAttributeLink
	Attr *X_xmlSchemaAttribute
}
type SchemaAttributeLink X_xmlSchemaAttributeLink
type SchemaAttributeLinkPtr *SchemaAttributeLink

type X_xmlSchemaWildcardNs struct {
	Next  *X_xmlSchemaWildcardNs
	Value *Char
}
type SchemaWildcardNs X_xmlSchemaWildcardNs
type SchemaWildcardNsPtr *SchemaWildcardNs

type X_xmlSchemaWildcard struct {
	Type            SchemaTypeType
	Id              *Char
	Annot           SchemaAnnotPtr
	Node            NodePtr
	MinOccurs       c.Int
	MaxOccurs       c.Int
	ProcessContents c.Int
	Any             c.Int
	NsSet           SchemaWildcardNsPtr
	NegNsSet        SchemaWildcardNsPtr
	Flags           c.Int
}
type SchemaWildcard X_xmlSchemaWildcard
type SchemaWildcardPtr *SchemaWildcard

type X_xmlSchemaAttributeGroup struct {
	Type              SchemaTypeType
	Next              *X_xmlSchemaAttribute
	Name              *Char
	Id                *Char
	Ref               *Char
	RefNs             *Char
	Annot             SchemaAnnotPtr
	Attributes        SchemaAttributePtr
	Node              NodePtr
	Flags             c.Int
	AttributeWildcard SchemaWildcardPtr
	RefPrefix         *Char
	RefItem           SchemaAttributeGroupPtr
	TargetNamespace   *Char
	AttrUses          unsafe.Pointer
}
type SchemaAttributeGroup X_xmlSchemaAttributeGroup
type SchemaAttributeGroupPtr *SchemaAttributeGroup

type X_xmlSchemaTypeLink struct {
	Next *X_xmlSchemaTypeLink
	Type SchemaTypePtr
}
type SchemaTypeLink X_xmlSchemaTypeLink
type SchemaTypeLinkPtr *SchemaTypeLink

type X_xmlSchemaFacetLink struct {
	Next  *X_xmlSchemaFacetLink
	Facet SchemaFacetPtr
}
type SchemaFacetLink X_xmlSchemaFacetLink
type SchemaFacetLinkPtr *SchemaFacetLink

type X_xmlSchemaElement struct {
	Type            SchemaTypeType
	Next            *X_xmlSchemaType
	Name            *Char
	Id              *Char
	Ref             *Char
	RefNs           *Char
	Annot           SchemaAnnotPtr
	Subtypes        SchemaTypePtr
	Attributes      SchemaAttributePtr
	Node            NodePtr
	MinOccurs       c.Int
	MaxOccurs       c.Int
	Flags           c.Int
	TargetNamespace *Char
	NamedType       *Char
	NamedTypeNs     *Char
	SubstGroup      *Char
	SubstGroupNs    *Char
	Scope           *Char
	Value           *Char
	RefDecl         *X_xmlSchemaElement
	ContModel       RegexpPtr
	ContentType     SchemaContentType
	RefPrefix       *Char
	DefVal          SchemaValPtr
	Idcs            unsafe.Pointer
}
type SchemaElement X_xmlSchemaElement
type SchemaElementPtr *SchemaElement

type X_xmlSchemaNotation struct {
	Type            SchemaTypeType
	Name            *Char
	Annot           SchemaAnnotPtr
	Identifier      *Char
	TargetNamespace *Char
}
type SchemaNotation X_xmlSchemaNotation
type SchemaNotationPtr *SchemaNotation

/**
 * _xmlSchema:
 *
 * A Schemas definition
 */

type X_xmlSchema struct {
	Name            *Char
	TargetNamespace *Char
	Version         *Char
	Id              *Char
	Doc             DocPtr
	Annot           SchemaAnnotPtr
	Flags           c.Int
	TypeDecl        HashTablePtr
	AttrDecl        HashTablePtr
	AttrgrpDecl     HashTablePtr
	ElemDecl        HashTablePtr
	NotaDecl        HashTablePtr
	SchemasImports  HashTablePtr
	X_private       unsafe.Pointer
	GroupDecl       HashTablePtr
	Dict            DictPtr
	Includes        unsafe.Pointer
	Preserve        c.Int
	Counter         c.Int
	IdcDef          HashTablePtr
	Volatiles       unsafe.Pointer
}

//go:linkname SchemaFreeType C.xmlSchemaFreeType
func SchemaFreeType(type_ SchemaTypePtr)

//go:linkname SchemaFreeWildcard C.xmlSchemaFreeWildcard
func SchemaFreeWildcard(wildcard SchemaWildcardPtr)
