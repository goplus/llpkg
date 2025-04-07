package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const XPATH_POINT = 5
const XPATH_RANGE = 6
const XPATH_LOCATIONSET = 7

type X_xmlXPathContext struct {
	Doc                DocPtr
	Node               NodePtr
	NbVariablesUnused  c.Int
	MaxVariablesUnused c.Int
	VarHash            HashTablePtr
	NbTypes            c.Int
	MaxTypes           c.Int
	Types              XPathTypePtr
	NbFuncsUnused      c.Int
	MaxFuncsUnused     c.Int
	FuncHash           HashTablePtr
	NbAxis             c.Int
	MaxAxis            c.Int
	Axis               XPathAxisPtr
	Namespaces         *NsPtr
	NsNr               c.Int
	User               unsafe.Pointer
	ContextSize        c.Int
	ProximityPosition  c.Int
	Xptr               c.Int
	Here               NodePtr
	Origin             NodePtr
	NsHash             HashTablePtr
	VarLookupFunc      unsafe.Pointer
	VarLookupData      unsafe.Pointer
	Extra              unsafe.Pointer
	Function           *Char
	FunctionURI        *Char
	FuncLookupFunc     unsafe.Pointer
	FuncLookupData     unsafe.Pointer
	TmpNsList          *NsPtr
	TmpNsNr            c.Int
	UserData           unsafe.Pointer
	Error              unsafe.Pointer
	LastError          Error
	DebugNode          NodePtr
	Dict               DictPtr
	Flags              c.Int
	Cache              unsafe.Pointer
	OpLimit            c.Ulong
	OpCount            c.Ulong
	Depth              c.Int
}
type XPathContext X_xmlXPathContext
type XPathContextPtr *XPathContext

type X_xmlXPathParserContext struct {
	Cur        *Char
	Base       *Char
	Error      c.Int
	Context    XPathContextPtr
	Value      XPathObjectPtr
	ValueNr    c.Int
	ValueMax   c.Int
	ValueTab   *XPathObjectPtr
	Comp       XPathCompExprPtr
	Xptr       c.Int
	Ancestor   NodePtr
	ValueFrame c.Int
}
type XPathParserContext X_xmlXPathParserContext
type XPathParserContextPtr *XPathParserContext
type XPathError c.Int

const (
	XPATHEXPRESSIONOK__1           XPathError = 0
	XPATHNUMBERERROR__1            XPathError = 1
	XPATHUNFINISHEDLITERALERROR__1 XPathError = 2
	XPATHSTARTLITERALERROR__1      XPathError = 3
	XPATHVARIABLEREFERROR__1       XPathError = 4
	XPATHUNDEFVARIABLEERROR__1     XPathError = 5
	XPATHINVALIDPREDICATEERROR__1  XPathError = 6
	XPATHEXPRERROR__1              XPathError = 7
	XPATHUNCLOSEDERROR__1          XPathError = 8
	XPATHUNKNOWNFUNCERROR__1       XPathError = 9
	XPATHINVALIDOPERAND__1         XPathError = 10
	XPATHINVALIDTYPE__1            XPathError = 11
	XPATHINVALIDARITY__1           XPathError = 12
	XPATHINVALIDCTXTSIZE__1        XPathError = 13
	XPATHINVALIDCTXTPOSITION__1    XPathError = 14
	XPATHMEMORYERROR__1            XPathError = 15
	XPTRSYNTAXERROR__1             XPathError = 16
	XPTRRESOURCEERROR__1           XPathError = 17
	XPTRSUBRESOURCEERROR__1        XPathError = 18
	XPATHUNDEFPREFIXERROR__1       XPathError = 19
	XPATHENCODINGERROR__1          XPathError = 20
	XPATHINVALIDCHARERROR__1       XPathError = 21
	XPATHINVALIDCTXT               XPathError = 22
	XPATHSTACKERROR                XPathError = 23
	XPATHFORBIDVARIABLEERROR       XPathError = 24
	XPATHOPLIMITEXCEEDED           XPathError = 25
	XPATHRECURSIONLIMITEXCEEDED    XPathError = 26
)

type X_xmlNodeSet struct {
	NodeNr  c.Int
	NodeMax c.Int
	NodeTab *NodePtr
}
type NodeSet X_xmlNodeSet
type NodeSetPtr *NodeSet
type XPathObjectType c.Int

const (
	XPATHUNDEFINED XPathObjectType = 0
	XPATHNODESET   XPathObjectType = 1
	XPATHBOOLEAN   XPathObjectType = 2
	XPATHNUMBER    XPathObjectType = 3
	XPATHSTRING    XPathObjectType = 4
	XPATHUSERS     XPathObjectType = 8
	XPATHXSLTTREE  XPathObjectType = 9
)

type X_xmlXPathObject struct {
	Type       XPathObjectType
	Nodesetval NodeSetPtr
	Boolval    c.Int
	Floatval   float64
	Stringval  *Char
	User       unsafe.Pointer
	Index      c.Int
	User2      unsafe.Pointer
	Index2     c.Int
}
type XPathObject X_xmlXPathObject
type XPathObjectPtr *XPathObject

// llgo:type C
type XPathConvertFunc func(XPathObjectPtr, c.Int) c.Int

type X_xmlXPathType struct {
	Name *Char
	Func unsafe.Pointer
}
type XPathType X_xmlXPathType
type XPathTypePtr *XPathType

type X_xmlXPathVariable struct {
	Name  *Char
	Value XPathObjectPtr
}
type XPathVariable X_xmlXPathVariable
type XPathVariablePtr *XPathVariable

// llgo:type C
type XPathEvalFunc func(XPathParserContextPtr, c.Int)

type X_xmlXPathFunct struct {
	Name *Char
	Func unsafe.Pointer
}
type XPathFunct X_xmlXPathFunct
type XPathFuncPtr *XPathFunct

// llgo:type C
type XPathAxisFunc func(XPathParserContextPtr, XPathObjectPtr) XPathObjectPtr

type X_xmlXPathAxis struct {
	Name *Char
	Func unsafe.Pointer
}
type XPathAxis X_xmlXPathAxis
type XPathAxisPtr *XPathAxis

// llgo:type C
type XPathFunction func(XPathParserContextPtr, c.Int)

// llgo:type C
type XPathVariableLookupFunc func(unsafe.Pointer, *Char, *Char) XPathObjectPtr

// llgo:type C
type XPathFuncLookupFunc func(unsafe.Pointer, *Char, *Char) XPathFunction

type X_xmlXPathCompExpr struct {
	Unused [8]uint8
}
type XPathCompExpr X_xmlXPathCompExpr
type XPathCompExprPtr *XPathCompExpr

//go:linkname XPathFreeObject C.xmlXPathFreeObject
func XPathFreeObject(obj XPathObjectPtr)

//go:linkname XPathNodeSetCreate C.xmlXPathNodeSetCreate
func XPathNodeSetCreate(val NodePtr) NodeSetPtr

//go:linkname XPathFreeNodeSetList C.xmlXPathFreeNodeSetList
func XPathFreeNodeSetList(obj XPathObjectPtr)

//go:linkname XPathFreeNodeSet C.xmlXPathFreeNodeSet
func XPathFreeNodeSet(obj NodeSetPtr)

//go:linkname XPathObjectCopy C.xmlXPathObjectCopy
func XPathObjectCopy(val XPathObjectPtr) XPathObjectPtr

//go:linkname XPathCmpNodes C.xmlXPathCmpNodes
func XPathCmpNodes(node1 NodePtr, node2 NodePtr) c.Int

/**
 * Conversion functions to basic types.
 */
//go:linkname XPathCastNumberToBoolean C.xmlXPathCastNumberToBoolean
func XPathCastNumberToBoolean(val float64) c.Int

// llgo:link (*Char).XPathCastStringToBoolean C.xmlXPathCastStringToBoolean
func (recv_ *Char) XPathCastStringToBoolean() c.Int {
	return 0
}

//go:linkname XPathCastNodeSetToBoolean C.xmlXPathCastNodeSetToBoolean
func XPathCastNodeSetToBoolean(ns NodeSetPtr) c.Int

//go:linkname XPathCastToBoolean C.xmlXPathCastToBoolean
func XPathCastToBoolean(val XPathObjectPtr) c.Int

//go:linkname XPathCastBooleanToNumber C.xmlXPathCastBooleanToNumber
func XPathCastBooleanToNumber(val c.Int) float64

// llgo:link (*Char).XPathCastStringToNumber C.xmlXPathCastStringToNumber
func (recv_ *Char) XPathCastStringToNumber() float64 {
	return 0
}

//go:linkname XPathCastNodeToNumber C.xmlXPathCastNodeToNumber
func XPathCastNodeToNumber(node NodePtr) float64

//go:linkname XPathCastNodeSetToNumber C.xmlXPathCastNodeSetToNumber
func XPathCastNodeSetToNumber(ns NodeSetPtr) float64

//go:linkname XPathCastToNumber C.xmlXPathCastToNumber
func XPathCastToNumber(val XPathObjectPtr) float64

//go:linkname XPathCastBooleanToString C.xmlXPathCastBooleanToString
func XPathCastBooleanToString(val c.Int) *Char

//go:linkname XPathCastNumberToString C.xmlXPathCastNumberToString
func XPathCastNumberToString(val float64) *Char

//go:linkname XPathCastNodeToString C.xmlXPathCastNodeToString
func XPathCastNodeToString(node NodePtr) *Char

//go:linkname XPathCastNodeSetToString C.xmlXPathCastNodeSetToString
func XPathCastNodeSetToString(ns NodeSetPtr) *Char

//go:linkname XPathCastToString C.xmlXPathCastToString
func XPathCastToString(val XPathObjectPtr) *Char

//go:linkname XPathConvertBoolean C.xmlXPathConvertBoolean
func XPathConvertBoolean(val XPathObjectPtr) XPathObjectPtr

//go:linkname XPathConvertNumber C.xmlXPathConvertNumber
func XPathConvertNumber(val XPathObjectPtr) XPathObjectPtr

//go:linkname XPathConvertString C.xmlXPathConvertString
func XPathConvertString(val XPathObjectPtr) XPathObjectPtr

/**
 * Context handling.
 */
//go:linkname XPathNewContext C.xmlXPathNewContext
func XPathNewContext(doc DocPtr) XPathContextPtr

//go:linkname XPathFreeContext C.xmlXPathFreeContext
func XPathFreeContext(ctxt XPathContextPtr)

//go:linkname XPathSetErrorHandler C.xmlXPathSetErrorHandler
func XPathSetErrorHandler(ctxt XPathContextPtr, handler StructuredErrorFunc, context unsafe.Pointer)

//go:linkname XPathContextSetCache C.xmlXPathContextSetCache
func XPathContextSetCache(ctxt XPathContextPtr, active c.Int, value c.Int, options c.Int) c.Int

/**
 * Evaluation functions.
 */
//go:linkname XPathOrderDocElems C.xmlXPathOrderDocElems
func XPathOrderDocElems(doc DocPtr) c.Long

//go:linkname XPathSetContextNode C.xmlXPathSetContextNode
func XPathSetContextNode(node NodePtr, ctx XPathContextPtr) c.Int

//go:linkname XPathNodeEval C.xmlXPathNodeEval
func XPathNodeEval(node NodePtr, str *Char, ctx XPathContextPtr) XPathObjectPtr

// llgo:link (*Char).XPathEval C.xmlXPathEval
func (recv_ *Char) XPathEval(ctx XPathContextPtr) XPathObjectPtr {
	return nil
}

// llgo:link (*Char).XPathEvalExpression C.xmlXPathEvalExpression
func (recv_ *Char) XPathEvalExpression(ctxt XPathContextPtr) XPathObjectPtr {
	return nil
}

//go:linkname XPathEvalPredicate C.xmlXPathEvalPredicate
func XPathEvalPredicate(ctxt XPathContextPtr, res XPathObjectPtr) c.Int

/**
 * Separate compilation/evaluation entry points.
 */
// llgo:link (*Char).XPathCompile C.xmlXPathCompile
func (recv_ *Char) XPathCompile() XPathCompExprPtr {
	return nil
}

//go:linkname XPathCtxtCompile C.xmlXPathCtxtCompile
func XPathCtxtCompile(ctxt XPathContextPtr, str *Char) XPathCompExprPtr

//go:linkname XPathCompiledEval C.xmlXPathCompiledEval
func XPathCompiledEval(comp XPathCompExprPtr, ctx XPathContextPtr) XPathObjectPtr

//go:linkname XPathCompiledEvalToBoolean C.xmlXPathCompiledEvalToBoolean
func XPathCompiledEvalToBoolean(comp XPathCompExprPtr, ctxt XPathContextPtr) c.Int

//go:linkname XPathFreeCompExpr C.xmlXPathFreeCompExpr
func XPathFreeCompExpr(comp XPathCompExprPtr)

//go:linkname XPathInit C.xmlXPathInit
func XPathInit()

//go:linkname XPathIsNaN C.xmlXPathIsNaN
func XPathIsNaN(val float64) c.Int

//go:linkname XPathIsInf C.xmlXPathIsInf
func XPathIsInf(val float64) c.Int
