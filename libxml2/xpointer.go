package libxml2

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type X_xmlLocationSet struct {
	LocNr  c.Int
	LocMax c.Int
	LocTab *XPathObjectPtr
}
type LocationSet X_xmlLocationSet
type LocationSetPtr *LocationSet

/*
 * Handling of location sets.
 */
//go:linkname XPtrLocationSetCreate C.xmlXPtrLocationSetCreate
func XPtrLocationSetCreate(val XPathObjectPtr) LocationSetPtr

//go:linkname XPtrFreeLocationSet C.xmlXPtrFreeLocationSet
func XPtrFreeLocationSet(obj LocationSetPtr)

//go:linkname XPtrLocationSetMerge C.xmlXPtrLocationSetMerge
func XPtrLocationSetMerge(val1 LocationSetPtr, val2 LocationSetPtr) LocationSetPtr

//go:linkname XPtrNewRange C.xmlXPtrNewRange
func XPtrNewRange(start NodePtr, startindex c.Int, end NodePtr, endindex c.Int) XPathObjectPtr

//go:linkname XPtrNewRangePoints C.xmlXPtrNewRangePoints
func XPtrNewRangePoints(start XPathObjectPtr, end XPathObjectPtr) XPathObjectPtr

//go:linkname XPtrNewRangeNodePoint C.xmlXPtrNewRangeNodePoint
func XPtrNewRangeNodePoint(start NodePtr, end XPathObjectPtr) XPathObjectPtr

//go:linkname XPtrNewRangePointNode C.xmlXPtrNewRangePointNode
func XPtrNewRangePointNode(start XPathObjectPtr, end NodePtr) XPathObjectPtr

//go:linkname XPtrNewRangeNodes C.xmlXPtrNewRangeNodes
func XPtrNewRangeNodes(start NodePtr, end NodePtr) XPathObjectPtr

//go:linkname XPtrNewLocationSetNodes C.xmlXPtrNewLocationSetNodes
func XPtrNewLocationSetNodes(start NodePtr, end NodePtr) XPathObjectPtr

//go:linkname XPtrNewLocationSetNodeSet C.xmlXPtrNewLocationSetNodeSet
func XPtrNewLocationSetNodeSet(set NodeSetPtr) XPathObjectPtr

//go:linkname XPtrNewRangeNodeObject C.xmlXPtrNewRangeNodeObject
func XPtrNewRangeNodeObject(start NodePtr, end XPathObjectPtr) XPathObjectPtr

//go:linkname XPtrNewCollapsedRange C.xmlXPtrNewCollapsedRange
func XPtrNewCollapsedRange(start NodePtr) XPathObjectPtr

//go:linkname XPtrLocationSetAdd C.xmlXPtrLocationSetAdd
func XPtrLocationSetAdd(cur LocationSetPtr, val XPathObjectPtr)

//go:linkname XPtrWrapLocationSet C.xmlXPtrWrapLocationSet
func XPtrWrapLocationSet(val LocationSetPtr) XPathObjectPtr

//go:linkname XPtrLocationSetDel C.xmlXPtrLocationSetDel
func XPtrLocationSetDel(cur LocationSetPtr, val XPathObjectPtr)

//go:linkname XPtrLocationSetRemove C.xmlXPtrLocationSetRemove
func XPtrLocationSetRemove(cur LocationSetPtr, val c.Int)

/*
 * Functions.
 */
//go:linkname XPtrNewContext C.xmlXPtrNewContext
func XPtrNewContext(doc DocPtr, here NodePtr, origin NodePtr) XPathContextPtr

// llgo:link (*Char).XPtrEval C.xmlXPtrEval
func (recv_ *Char) XPtrEval(ctx XPathContextPtr) XPathObjectPtr {
	return nil
}

//go:linkname XPtrRangeToFunction C.xmlXPtrRangeToFunction
func XPtrRangeToFunction(ctxt XPathParserContextPtr, nargs c.Int)

//go:linkname XPtrBuildNodeList C.xmlXPtrBuildNodeList
func XPtrBuildNodeList(obj XPathObjectPtr) NodePtr

//go:linkname XPtrEvalRangePredicate C.xmlXPtrEvalRangePredicate
func XPtrEvalRangePredicate(ctxt XPathParserContextPtr)
