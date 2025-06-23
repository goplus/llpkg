package quickjs

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const LRE_CC_RES_LEN_MAX = 3

type UnicodeNormalizationEnum c.Int

const (
	UNICODE_NFC  UnicodeNormalizationEnum = 0
	UNICODE_NFD  UnicodeNormalizationEnum = 1
	UNICODE_NFKC UnicodeNormalizationEnum = 2
	UNICODE_NFKD UnicodeNormalizationEnum = 3
)

//go:linkname LreCaseConv C.lre_case_conv
func LreCaseConv(res *c.Uint32T, c c.Uint32T, conv_type c.Int) c.Int

//go:linkname LreCanonicalize C.lre_canonicalize
func LreCanonicalize(c c.Uint32T, is_unicode c.Int) c.Int

//go:linkname LreIsCased C.lre_is_cased
func LreIsCased(c c.Uint32T) c.Int

//go:linkname LreIsCaseIgnorable C.lre_is_case_ignorable
func LreIsCaseIgnorable(c c.Uint32T) c.Int

/* char ranges */

type CharRange struct {
	Len         c.Int
	Size        c.Int
	Points      *c.Uint32T
	MemOpaque   c.Pointer
	ReallocFunc c.Pointer
}
type CharRangeOpEnum c.Int

const (
	CR_OP_UNION CharRangeOpEnum = 0
	CR_OP_INTER CharRangeOpEnum = 1
	CR_OP_XOR   CharRangeOpEnum = 2
)

// llgo:link (*CharRange).CrInit C.cr_init
func (recv_ *CharRange) CrInit(mem_opaque c.Pointer, realloc_func func(c.Pointer, c.Pointer, c.SizeT) c.Pointer) {
}

// llgo:link (*CharRange).CrFree C.cr_free
func (recv_ *CharRange) CrFree() {
}

// llgo:link (*CharRange).CrRealloc C.cr_realloc
func (recv_ *CharRange) CrRealloc(size c.Int) c.Int {
	return 0
}

// llgo:link (*CharRange).CrCopy C.cr_copy
func (recv_ *CharRange) CrCopy(cr1 *CharRange) c.Int {
	return 0
}

// llgo:link (*CharRange).CrUnion1 C.cr_union1
func (recv_ *CharRange) CrUnion1(b_pt *c.Uint32T, b_len c.Int) c.Int {
	return 0
}

// llgo:link (*CharRange).CrOp C.cr_op
func (recv_ *CharRange) CrOp(a_pt *c.Uint32T, a_len c.Int, b_pt *c.Uint32T, b_len c.Int, op c.Int) c.Int {
	return 0
}

// llgo:link (*CharRange).CrInvert C.cr_invert
func (recv_ *CharRange) CrInvert() c.Int {
	return 0
}

// llgo:link (*CharRange).CrRegexpCanonicalize C.cr_regexp_canonicalize
func (recv_ *CharRange) CrRegexpCanonicalize(is_unicode c.Int) c.Int {
	return 0
}

//go:linkname LreIsIdStart C.lre_is_id_start
func LreIsIdStart(c c.Uint32T) c.Int

//go:linkname LreIsIdContinue C.lre_is_id_continue
func LreIsIdContinue(c c.Uint32T) c.Int

//go:linkname UnicodeNormalize C.unicode_normalize
func UnicodeNormalize(pdst **c.Uint32T, src *c.Uint32T, src_len c.Int, n_type UnicodeNormalizationEnum, opaque c.Pointer, realloc_func func(c.Pointer, c.Pointer, c.SizeT) c.Pointer) c.Int

/* Unicode character range functions */
// llgo:link (*CharRange).UnicodeScript C.unicode_script
func (recv_ *CharRange) UnicodeScript(script_name *c.Char, is_ext c.Int) c.Int {
	return 0
}

// llgo:link (*CharRange).UnicodeGeneralCategory C.unicode_general_category
func (recv_ *CharRange) UnicodeGeneralCategory(gc_name *c.Char) c.Int {
	return 0
}

// llgo:link (*CharRange).UnicodeProp C.unicode_prop
func (recv_ *CharRange) UnicodeProp(prop_name *c.Char) c.Int {
	return 0
}
