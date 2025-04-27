package quickjs

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const LRE_CC_RES_LEN_MAX = 3

type UnicodeNormalizationEnum c.Int

const (
	UNICODENFC  UnicodeNormalizationEnum = 0
	UNICODENFD  UnicodeNormalizationEnum = 1
	UNICODENFKC UnicodeNormalizationEnum = 2
	UNICODENFKD UnicodeNormalizationEnum = 3
)

//go:linkname LreCaseConv C.lre_case_conv
func LreCaseConv(res *uint32, c uint32, conv_type c.Int) c.Int

//go:linkname LreCanonicalize C.lre_canonicalize
func LreCanonicalize(c uint32, is_unicode c.Int) c.Int

//go:linkname LreIsCased C.lre_is_cased
func LreIsCased(c uint32) c.Int

//go:linkname LreIsCaseIgnorable C.lre_is_case_ignorable
func LreIsCaseIgnorable(c uint32) c.Int

/* char ranges */

type CharRange struct {
	Len         c.Int
	Size        c.Int
	Points      *uint32
	MemOpaque   unsafe.Pointer
	ReallocFunc unsafe.Pointer
}
type CharRangeOpEnum c.Int

const (
	CROPUNION CharRangeOpEnum = 0
	CROPINTER CharRangeOpEnum = 1
	CROPXOR   CharRangeOpEnum = 2
)

// llgo:link (*CharRange).CrInit C.cr_init
func (recv_ *CharRange) CrInit(mem_opaque unsafe.Pointer, realloc_func func(unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer) {
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
func (recv_ *CharRange) CrUnion1(b_pt *uint32, b_len c.Int) c.Int {
	return 0
}

// llgo:link (*CharRange).CrOp C.cr_op
func (recv_ *CharRange) CrOp(a_pt *uint32, a_len c.Int, b_pt *uint32, b_len c.Int, op c.Int) c.Int {
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
func LreIsIdStart(c uint32) c.Int

//go:linkname LreIsIdContinue C.lre_is_id_continue
func LreIsIdContinue(c uint32) c.Int

//go:linkname UnicodeNormalize C.unicode_normalize
func UnicodeNormalize(pdst **uint32, src *uint32, src_len c.Int, n_type UnicodeNormalizationEnum, opaque unsafe.Pointer, realloc_func func(unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer) c.Int

/* Unicode character range functions */
// llgo:link (*CharRange).UnicodeScript C.unicode_script
func (recv_ *CharRange) UnicodeScript(script_name *int8, is_ext c.Int) c.Int {
	return 0
}

// llgo:link (*CharRange).UnicodeGeneralCategory C.unicode_general_category
func (recv_ *CharRange) UnicodeGeneralCategory(gc_name *int8) c.Int {
	return 0
}

// llgo:link (*CharRange).UnicodeProp C.unicode_prop
func (recv_ *CharRange) UnicodeProp(prop_name *int8) c.Int {
	return 0
}
