package libxslt

import (
	_ "unsafe"

	"github.com/goplus/lib/c"
	"github.com/goplus/llpkg/libxml2"
)

//go:linkname FunctionNodeSet C.xsltFunctionNodeSet
func FunctionNodeSet(ctxt libxml2.XPathParserContextPtr, nargs c.Int)

//go:linkname Debug C.xsltDebug
func Debug(ctxt TransformContextPtr, node libxml2.NodePtr, inst libxml2.NodePtr, comp ElemPreCompPtr)

//go:linkname RegisterExtras C.xsltRegisterExtras
func RegisterExtras(ctxt TransformContextPtr)

//go:linkname RegisterAllExtras C.xsltRegisterAllExtras
func RegisterAllExtras()
