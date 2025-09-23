package pcre

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const REG_ICASE = 0x0001
const REG_NEWLINE = 0x0002
const REG_NOTBOL = 0x0004
const REG_NOTEOL = 0x0008
const REG_DOTALL = 0x0010
const REG_NOSUB = 0x0020
const REG_UTF8 = 0x0040
const REG_STARTEND = 0x0080
const REG_NOTEMPTY = 0x0100
const REG_UNGREEDY = 0x0200
const REG_UCP = 0x0400
const REG_EXTENDED = 0
const (
	REGASSERT   c.Int = 1
	REGBADBR    c.Int = 2
	REGBADPAT   c.Int = 3
	REGBADRPT   c.Int = 4
	REGEBRACE   c.Int = 5
	REGEBRACK   c.Int = 6
	REGECOLLATE c.Int = 7
	REGECTYPE   c.Int = 8
	REGEESCAPE  c.Int = 9
	REGEMPTY    c.Int = 10
	REGEPAREN   c.Int = 11
	REGERANGE   c.Int = 12
	REGESIZE    c.Int = 13
	REGESPACE   c.Int = 14
	REGESUBREG  c.Int = 15
	REGINVARG   c.Int = 16
	REGNOMATCH  c.Int = 17
)

/* The structure representing a compiled regular expression. */

type RegexT struct {
	RePcre      unsafe.Pointer
	ReNsub      uintptr
	ReErroffset uintptr
}
type RegoffT c.Int

type RegmatchT struct {
	RmSo RegoffT
	RmEo RegoffT
}

/* The functions */
// llgo:link (*RegexT).Regcomp C.regcomp
func (recv_ *RegexT) Regcomp(*int8, c.Int) c.Int {
	return 0
}

// llgo:link (*RegexT).Regexec C.regexec
func (recv_ *RegexT) Regexec(*int8, uintptr, *RegmatchT, c.Int) c.Int {
	return 0
}

//go:linkname Regerror C.regerror
func Regerror(c.Int, *RegexT, *int8, uintptr) uintptr

// llgo:link (*RegexT).Regfree C.regfree
func (recv_ *RegexT) Regfree() {
}
