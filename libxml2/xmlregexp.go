package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type X_xmlRegexp struct {
	Unused [8]uint8
}
type Regexp X_xmlRegexp
type RegexpPtr *Regexp

type X_xmlRegExecCtxt struct {
	Unused [8]uint8
}
type RegExecCtxt X_xmlRegExecCtxt
type RegExecCtxtPtr *RegExecCtxt

/*
 * The POSIX like API
 */
// llgo:link (*Char).RegexpCompile C.xmlRegexpCompile
func (recv_ *Char) RegexpCompile() RegexpPtr {
	return nil
}

//go:linkname RegFreeRegexp C.xmlRegFreeRegexp
func RegFreeRegexp(regexp RegexpPtr)

//go:linkname RegexpExec C.xmlRegexpExec
func RegexpExec(comp RegexpPtr, value *Char) c.Int

//go:linkname RegexpPrint C.xmlRegexpPrint
func RegexpPrint(output *c.FILE, regexp RegexpPtr)

//go:linkname RegexpIsDeterminist C.xmlRegexpIsDeterminist
func RegexpIsDeterminist(comp RegexpPtr) c.Int

// llgo:type C
type RegExecCallbacks func(RegExecCtxtPtr, *Char, unsafe.Pointer, unsafe.Pointer)

/*
 * The progressive API
 */
//go:linkname RegNewExecCtxt C.xmlRegNewExecCtxt
func RegNewExecCtxt(comp RegexpPtr, callback RegExecCallbacks, data unsafe.Pointer) RegExecCtxtPtr

//go:linkname RegFreeExecCtxt C.xmlRegFreeExecCtxt
func RegFreeExecCtxt(exec RegExecCtxtPtr)

//go:linkname RegExecPushString C.xmlRegExecPushString
func RegExecPushString(exec RegExecCtxtPtr, value *Char, data unsafe.Pointer) c.Int

//go:linkname RegExecPushString2 C.xmlRegExecPushString2
func RegExecPushString2(exec RegExecCtxtPtr, value *Char, value2 *Char, data unsafe.Pointer) c.Int

//go:linkname RegExecNextValues C.xmlRegExecNextValues
func RegExecNextValues(exec RegExecCtxtPtr, nbval *c.Int, nbneg *c.Int, values **Char, terminal *c.Int) c.Int

//go:linkname RegExecErrInfo C.xmlRegExecErrInfo
func RegExecErrInfo(exec RegExecCtxtPtr, string **Char, nbval *c.Int, nbneg *c.Int, values **Char, terminal *c.Int) c.Int

type X_xmlExpCtxt struct {
	Unused [8]uint8
}
type ExpCtxt X_xmlExpCtxt
type ExpCtxtPtr *ExpCtxt

//go:linkname ExpFreeCtxt C.xmlExpFreeCtxt
func ExpFreeCtxt(ctxt ExpCtxtPtr)

//go:linkname ExpNewCtxt C.xmlExpNewCtxt
func ExpNewCtxt(maxNodes c.Int, dict DictPtr) ExpCtxtPtr

//go:linkname ExpCtxtNbNodes C.xmlExpCtxtNbNodes
func ExpCtxtNbNodes(ctxt ExpCtxtPtr) c.Int

//go:linkname ExpCtxtNbCons C.xmlExpCtxtNbCons
func ExpCtxtNbCons(ctxt ExpCtxtPtr) c.Int

type X_xmlExpNode struct {
	Unused [8]uint8
}
type ExpNode X_xmlExpNode
type ExpNodePtr *ExpNode
type ExpNodeType c.Int

const (
	EXPEMPTY  ExpNodeType = 0
	EXPFORBID ExpNodeType = 1
	EXPATOM   ExpNodeType = 2
	EXPSEQ    ExpNodeType = 3
	EXPOR     ExpNodeType = 4
	EXPCOUNT  ExpNodeType = 5
)

/*
 * Expressions are reference counted internally
 */
//go:linkname ExpFree C.xmlExpFree
func ExpFree(ctxt ExpCtxtPtr, expr ExpNodePtr)

//go:linkname ExpRef C.xmlExpRef
func ExpRef(expr ExpNodePtr)

/*
 * constructors can be either manual or from a string
 */
//go:linkname ExpParse C.xmlExpParse
func ExpParse(ctxt ExpCtxtPtr, expr *int8) ExpNodePtr

//go:linkname ExpNewAtom C.xmlExpNewAtom
func ExpNewAtom(ctxt ExpCtxtPtr, name *Char, len c.Int) ExpNodePtr

//go:linkname ExpNewOr C.xmlExpNewOr
func ExpNewOr(ctxt ExpCtxtPtr, left ExpNodePtr, right ExpNodePtr) ExpNodePtr

//go:linkname ExpNewSeq C.xmlExpNewSeq
func ExpNewSeq(ctxt ExpCtxtPtr, left ExpNodePtr, right ExpNodePtr) ExpNodePtr

//go:linkname ExpNewRange C.xmlExpNewRange
func ExpNewRange(ctxt ExpCtxtPtr, subset ExpNodePtr, min c.Int, max c.Int) ExpNodePtr

/*
 * The really interesting APIs
 */
//go:linkname ExpIsNillable C.xmlExpIsNillable
func ExpIsNillable(expr ExpNodePtr) c.Int

//go:linkname ExpMaxToken C.xmlExpMaxToken
func ExpMaxToken(expr ExpNodePtr) c.Int

//go:linkname ExpGetLanguage C.xmlExpGetLanguage
func ExpGetLanguage(ctxt ExpCtxtPtr, expr ExpNodePtr, langList **Char, len c.Int) c.Int

//go:linkname ExpGetStart C.xmlExpGetStart
func ExpGetStart(ctxt ExpCtxtPtr, expr ExpNodePtr, tokList **Char, len c.Int) c.Int

//go:linkname ExpStringDerive C.xmlExpStringDerive
func ExpStringDerive(ctxt ExpCtxtPtr, expr ExpNodePtr, str *Char, len c.Int) ExpNodePtr

//go:linkname ExpExpDerive C.xmlExpExpDerive
func ExpExpDerive(ctxt ExpCtxtPtr, expr ExpNodePtr, sub ExpNodePtr) ExpNodePtr

//go:linkname ExpSubsume C.xmlExpSubsume
func ExpSubsume(ctxt ExpCtxtPtr, expr ExpNodePtr, sub ExpNodePtr) c.Int

//go:linkname ExpDump C.xmlExpDump
func ExpDump(buf BufferPtr, expr ExpNodePtr)
