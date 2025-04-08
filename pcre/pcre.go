package pcre

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const PCRE_MAJOR = 8
const PCRE_MINOR = 45
const PCRE_CASELESS = 0x00000001
const PCRE_MULTILINE = 0x00000002
const PCRE_DOTALL = 0x00000004
const PCRE_EXTENDED = 0x00000008
const PCRE_ANCHORED = 0x00000010
const PCRE_DOLLAR_ENDONLY = 0x00000020
const PCRE_EXTRA = 0x00000040
const PCRE_NOTBOL = 0x00000080
const PCRE_NOTEOL = 0x00000100
const PCRE_UNGREEDY = 0x00000200
const PCRE_NOTEMPTY = 0x00000400
const PCRE_UTF8 = 0x00000800
const PCRE_UTF16 = 0x00000800
const PCRE_UTF32 = 0x00000800
const PCRE_NO_AUTO_CAPTURE = 0x00001000
const PCRE_NO_UTF8_CHECK = 0x00002000
const PCRE_NO_UTF16_CHECK = 0x00002000
const PCRE_NO_UTF32_CHECK = 0x00002000
const PCRE_AUTO_CALLOUT = 0x00004000
const PCRE_PARTIAL_SOFT = 0x00008000
const PCRE_PARTIAL = 0x00008000
const PCRE_NEVER_UTF = 0x00010000
const PCRE_DFA_SHORTEST = 0x00010000
const PCRE_NO_AUTO_POSSESS = 0x00020000
const PCRE_DFA_RESTART = 0x00020000
const PCRE_FIRSTLINE = 0x00040000
const PCRE_DUPNAMES = 0x00080000
const PCRE_NEWLINE_CR = 0x00100000
const PCRE_NEWLINE_LF = 0x00200000
const PCRE_NEWLINE_CRLF = 0x00300000
const PCRE_NEWLINE_ANY = 0x00400000
const PCRE_NEWLINE_ANYCRLF = 0x00500000
const PCRE_BSR_ANYCRLF = 0x00800000
const PCRE_BSR_UNICODE = 0x01000000
const PCRE_JAVASCRIPT_COMPAT = 0x02000000
const PCRE_NO_START_OPTIMIZE = 0x04000000
const PCRE_NO_START_OPTIMISE = 0x04000000
const PCRE_PARTIAL_HARD = 0x08000000
const PCRE_NOTEMPTY_ATSTART = 0x10000000
const PCRE_UCP = 0x20000000
const PCRE_UTF8_ERR0 = 0
const PCRE_UTF8_ERR1 = 1
const PCRE_UTF8_ERR2 = 2
const PCRE_UTF8_ERR3 = 3
const PCRE_UTF8_ERR4 = 4
const PCRE_UTF8_ERR5 = 5
const PCRE_UTF8_ERR6 = 6
const PCRE_UTF8_ERR7 = 7
const PCRE_UTF8_ERR8 = 8
const PCRE_UTF8_ERR9 = 9
const PCRE_UTF8_ERR10 = 10
const PCRE_UTF8_ERR11 = 11
const PCRE_UTF8_ERR12 = 12
const PCRE_UTF8_ERR13 = 13
const PCRE_UTF8_ERR14 = 14
const PCRE_UTF8_ERR15 = 15
const PCRE_UTF8_ERR16 = 16
const PCRE_UTF8_ERR17 = 17
const PCRE_UTF8_ERR18 = 18
const PCRE_UTF8_ERR19 = 19
const PCRE_UTF8_ERR20 = 20
const PCRE_UTF8_ERR21 = 21
const PCRE_UTF8_ERR22 = 22
const PCRE_UTF16_ERR0 = 0
const PCRE_UTF16_ERR1 = 1
const PCRE_UTF16_ERR2 = 2
const PCRE_UTF16_ERR3 = 3
const PCRE_UTF16_ERR4 = 4
const PCRE_UTF32_ERR0 = 0
const PCRE_UTF32_ERR1 = 1
const PCRE_UTF32_ERR2 = 2
const PCRE_UTF32_ERR3 = 3
const PCRE_INFO_OPTIONS = 0
const PCRE_INFO_SIZE = 1
const PCRE_INFO_CAPTURECOUNT = 2
const PCRE_INFO_BACKREFMAX = 3
const PCRE_INFO_FIRSTBYTE = 4
const PCRE_INFO_FIRSTCHAR = 4
const PCRE_INFO_FIRSTTABLE = 5
const PCRE_INFO_LASTLITERAL = 6
const PCRE_INFO_NAMEENTRYSIZE = 7
const PCRE_INFO_NAMECOUNT = 8
const PCRE_INFO_NAMETABLE = 9
const PCRE_INFO_STUDYSIZE = 10
const PCRE_INFO_DEFAULT_TABLES = 11
const PCRE_INFO_OKPARTIAL = 12
const PCRE_INFO_JCHANGED = 13
const PCRE_INFO_HASCRORLF = 14
const PCRE_INFO_MINLENGTH = 15
const PCRE_INFO_JIT = 16
const PCRE_INFO_JITSIZE = 17
const PCRE_INFO_MAXLOOKBEHIND = 18
const PCRE_INFO_FIRSTCHARACTER = 19
const PCRE_INFO_FIRSTCHARACTERFLAGS = 20
const PCRE_INFO_REQUIREDCHAR = 21
const PCRE_INFO_REQUIREDCHARFLAGS = 22
const PCRE_INFO_MATCHLIMIT = 23
const PCRE_INFO_RECURSIONLIMIT = 24
const PCRE_INFO_MATCH_EMPTY = 25
const PCRE_CONFIG_UTF8 = 0
const PCRE_CONFIG_NEWLINE = 1
const PCRE_CONFIG_LINK_SIZE = 2
const PCRE_CONFIG_POSIX_MALLOC_THRESHOLD = 3
const PCRE_CONFIG_MATCH_LIMIT = 4
const PCRE_CONFIG_STACKRECURSE = 5
const PCRE_CONFIG_UNICODE_PROPERTIES = 6
const PCRE_CONFIG_MATCH_LIMIT_RECURSION = 7
const PCRE_CONFIG_BSR = 8
const PCRE_CONFIG_JIT = 9
const PCRE_CONFIG_UTF16 = 10
const PCRE_CONFIG_JITTARGET = 11
const PCRE_CONFIG_UTF32 = 12
const PCRE_CONFIG_PARENS_LIMIT = 13
const PCRE_STUDY_JIT_COMPILE = 0x0001
const PCRE_STUDY_JIT_PARTIAL_SOFT_COMPILE = 0x0002
const PCRE_STUDY_JIT_PARTIAL_HARD_COMPILE = 0x0004
const PCRE_STUDY_EXTRA_NEEDED = 0x0008
const PCRE_EXTRA_STUDY_DATA = 0x0001
const PCRE_EXTRA_MATCH_LIMIT = 0x0002
const PCRE_EXTRA_CALLOUT_DATA = 0x0004
const PCRE_EXTRA_TABLES = 0x0008
const PCRE_EXTRA_MATCH_LIMIT_RECURSION = 0x0010
const PCRE_EXTRA_MARK = 0x0020
const PCRE_EXTRA_EXECUTABLE_JIT = 0x0040

/* Types */

type RealPcre8Or16 struct {
	Unused [8]uint8
}
type Pcre RealPcre8Or16
type Pcre16 RealPcre8Or16

type RealPcre32 struct {
	Unused [8]uint8
}
type Pcre32 RealPcre32

type RealPcreJitStack struct {
	Unused [8]uint8
}
type JitStack RealPcreJitStack

type RealPcre16JitStack struct {
	Unused [8]uint8
}
type JitStack__1 RealPcre16JitStack

type RealPcre32JitStack struct {
	Unused [8]uint8
}
type JitStack__2 RealPcre32JitStack

/*
	The structure for passing additional data to pcre_exec(). This is defined in

such as way as to be extensible. Always add new fields at the end, in order to
remain compatible.
*/
type Extra struct {
	Flags               c.Ulong
	StudyData           unsafe.Pointer
	MatchLimit          c.Ulong
	CalloutData         unsafe.Pointer
	Tables              *int8
	MatchLimitRecursion c.Ulong
	Mark                **int8
	ExecutableJit       unsafe.Pointer
}

/* Same structure as above, but with 16 bit char pointers. */

type Extra__1 struct {
	Flags               c.Ulong
	StudyData           unsafe.Pointer
	MatchLimit          c.Ulong
	CalloutData         unsafe.Pointer
	Tables              *int8
	MatchLimitRecursion c.Ulong
	Mark                **uint16
	ExecutableJit       unsafe.Pointer
}

/* Same structure as above, but with 32 bit char pointers. */

type Extra__2 struct {
	Flags               c.Ulong
	StudyData           unsafe.Pointer
	MatchLimit          c.Ulong
	CalloutData         unsafe.Pointer
	Tables              *int8
	MatchLimitRecursion c.Ulong
	Mark                **c.Uint
	ExecutableJit       unsafe.Pointer
}

/*
	The structure for passing out data via the pcre_callout_function. We use a

structure so that new fields can be added on the end in future versions,
without changing the API of the function, thereby allowing old clients to work
without modification.
*/
type CalloutBlock struct {
	Version         c.Int
	CalloutNumber   c.Int
	OffsetVector    *c.Int
	Subject         *int8
	SubjectLength   c.Int
	StartMatch      c.Int
	CurrentPosition c.Int
	CaptureTop      c.Int
	CaptureLast     c.Int
	CalloutData     unsafe.Pointer
	PatternPosition c.Int
	NextItemLength  c.Int
	Mark            *int8
}

/* Same structure as above, but with 16 bit char pointers. */

type CalloutBlock__1 struct {
	Version         c.Int
	CalloutNumber   c.Int
	OffsetVector    *c.Int
	Subject         *uint16
	SubjectLength   c.Int
	StartMatch      c.Int
	CurrentPosition c.Int
	CaptureTop      c.Int
	CaptureLast     c.Int
	CalloutData     unsafe.Pointer
	PatternPosition c.Int
	NextItemLength  c.Int
	Mark            *uint16
}

/* Same structure as above, but with 32 bit char pointers. */

type CalloutBlock__2 struct {
	Version         c.Int
	CalloutNumber   c.Int
	OffsetVector    *c.Int
	Subject         *c.Uint
	SubjectLength   c.Int
	StartMatch      c.Int
	CurrentPosition c.Int
	CaptureTop      c.Int
	CaptureLast     c.Int
	CalloutData     unsafe.Pointer
	PatternPosition c.Int
	NextItemLength  c.Int
	Mark            *c.Uint
}

// llgo:type C
type JitCallback func(unsafe.Pointer) *JitStack

// llgo:type C
type JitCallback__1 func(unsafe.Pointer) *JitStack__1

// llgo:type C
type JitCallback__2 func(unsafe.Pointer) *JitStack__2

/* Exported PCRE functions */
//go:linkname Compile C.pcre_compile
func Compile(*int8, c.Int, **int8, *c.Int, *int8) *Pcre

//go:linkname Compile__1 C.pcre16_compile
func Compile__1(*uint16, c.Int, **int8, *c.Int, *int8) *Pcre16

//go:linkname Compile__2 C.pcre32_compile
func Compile__2(*c.Uint, c.Int, **int8, *c.Int, *int8) *Pcre32

//go:linkname Compile2 C.pcre_compile2
func Compile2(*int8, c.Int, *c.Int, **int8, *c.Int, *int8) *Pcre

//go:linkname Compile2__1 C.pcre16_compile2
func Compile2__1(*uint16, c.Int, *c.Int, **int8, *c.Int, *int8) *Pcre16

//go:linkname Compile2__2 C.pcre32_compile2
func Compile2__2(*c.Uint, c.Int, *c.Int, **int8, *c.Int, *int8) *Pcre32

//go:linkname Config C.pcre_config
func Config(c.Int, unsafe.Pointer) c.Int

//go:linkname Config__1 C.pcre16_config
func Config__1(c.Int, unsafe.Pointer) c.Int

//go:linkname Config__2 C.pcre32_config
func Config__2(c.Int, unsafe.Pointer) c.Int

// llgo:link (*Pcre).CopyNamedSubstring C.pcre_copy_named_substring
func (recv_ *Pcre) CopyNamedSubstring(*int8, *c.Int, c.Int, *int8, *int8, c.Int) c.Int {
	return 0
}

// llgo:link (*Pcre16).CopyNamedSubstring C.pcre16_copy_named_substring
func (recv_ *Pcre16) CopyNamedSubstring(*uint16, *c.Int, c.Int, *uint16, *uint16, c.Int) c.Int {
	return 0
}

// llgo:link (*Pcre32).CopyNamedSubstring C.pcre32_copy_named_substring
func (recv_ *Pcre32) CopyNamedSubstring(*c.Uint, *c.Int, c.Int, *c.Uint, *c.Uint, c.Int) c.Int {
	return 0
}

//go:linkname CopySubstring C.pcre_copy_substring
func CopySubstring(*int8, *c.Int, c.Int, c.Int, *int8, c.Int) c.Int

//go:linkname CopySubstring__1 C.pcre16_copy_substring
func CopySubstring__1(*uint16, *c.Int, c.Int, c.Int, *uint16, c.Int) c.Int

//go:linkname CopySubstring__2 C.pcre32_copy_substring
func CopySubstring__2(*c.Uint, *c.Int, c.Int, c.Int, *c.Uint, c.Int) c.Int

// llgo:link (*Pcre).DfaExec C.pcre_dfa_exec
func (recv_ *Pcre) DfaExec(*Extra, *int8, c.Int, c.Int, c.Int, *c.Int, c.Int, *c.Int, c.Int) c.Int {
	return 0
}

// llgo:link (*Pcre16).DfaExec C.pcre16_dfa_exec
func (recv_ *Pcre16) DfaExec(*Extra__1, *uint16, c.Int, c.Int, c.Int, *c.Int, c.Int, *c.Int, c.Int) c.Int {
	return 0
}

// llgo:link (*Pcre32).DfaExec C.pcre32_dfa_exec
func (recv_ *Pcre32) DfaExec(*Extra__2, *c.Uint, c.Int, c.Int, c.Int, *c.Int, c.Int, *c.Int, c.Int) c.Int {
	return 0
}

// llgo:link (*Pcre).Exec C.pcre_exec
func (recv_ *Pcre) Exec(*Extra, *int8, c.Int, c.Int, c.Int, *c.Int, c.Int) c.Int {
	return 0
}

// llgo:link (*Pcre16).Exec C.pcre16_exec
func (recv_ *Pcre16) Exec(*Extra__1, *uint16, c.Int, c.Int, c.Int, *c.Int, c.Int) c.Int {
	return 0
}

// llgo:link (*Pcre32).Exec C.pcre32_exec
func (recv_ *Pcre32) Exec(*Extra__2, *c.Uint, c.Int, c.Int, c.Int, *c.Int, c.Int) c.Int {
	return 0
}

//go:linkname FreeSubstring C.pcre_free_substring
func FreeSubstring(*int8)

//go:linkname FreeSubstring__1 C.pcre16_free_substring
func FreeSubstring__1(*uint16)

//go:linkname FreeSubstring__2 C.pcre32_free_substring
func FreeSubstring__2(*c.Uint)

//go:linkname FreeSubstringList C.pcre_free_substring_list
func FreeSubstringList(**int8)

//go:linkname FreeSubstringList__1 C.pcre16_free_substring_list
func FreeSubstringList__1(**uint16)

//go:linkname FreeSubstringList__2 C.pcre32_free_substring_list
func FreeSubstringList__2(**c.Uint)

// llgo:link (*Pcre).Fullinfo C.pcre_fullinfo
func (recv_ *Pcre) Fullinfo(*Extra, c.Int, unsafe.Pointer) c.Int {
	return 0
}

// llgo:link (*Pcre16).Fullinfo C.pcre16_fullinfo
func (recv_ *Pcre16) Fullinfo(*Extra__1, c.Int, unsafe.Pointer) c.Int {
	return 0
}

// llgo:link (*Pcre32).Fullinfo C.pcre32_fullinfo
func (recv_ *Pcre32) Fullinfo(*Extra__2, c.Int, unsafe.Pointer) c.Int {
	return 0
}

// llgo:link (*Pcre).GetNamedSubstring C.pcre_get_named_substring
func (recv_ *Pcre) GetNamedSubstring(*int8, *c.Int, c.Int, *int8, **int8) c.Int {
	return 0
}

// llgo:link (*Pcre16).GetNamedSubstring C.pcre16_get_named_substring
func (recv_ *Pcre16) GetNamedSubstring(*uint16, *c.Int, c.Int, *uint16, **uint16) c.Int {
	return 0
}

// llgo:link (*Pcre32).GetNamedSubstring C.pcre32_get_named_substring
func (recv_ *Pcre32) GetNamedSubstring(*c.Uint, *c.Int, c.Int, *c.Uint, **c.Uint) c.Int {
	return 0
}

// llgo:link (*Pcre).GetStringnumber C.pcre_get_stringnumber
func (recv_ *Pcre) GetStringnumber(*int8) c.Int {
	return 0
}

// llgo:link (*Pcre16).GetStringnumber C.pcre16_get_stringnumber
func (recv_ *Pcre16) GetStringnumber(*uint16) c.Int {
	return 0
}

// llgo:link (*Pcre32).GetStringnumber C.pcre32_get_stringnumber
func (recv_ *Pcre32) GetStringnumber(*c.Uint) c.Int {
	return 0
}

// llgo:link (*Pcre).GetStringtableEntries C.pcre_get_stringtable_entries
func (recv_ *Pcre) GetStringtableEntries(*int8, **int8, **int8) c.Int {
	return 0
}

// llgo:link (*Pcre16).GetStringtableEntries C.pcre16_get_stringtable_entries
func (recv_ *Pcre16) GetStringtableEntries(*uint16, **uint16, **uint16) c.Int {
	return 0
}

// llgo:link (*Pcre32).GetStringtableEntries C.pcre32_get_stringtable_entries
func (recv_ *Pcre32) GetStringtableEntries(*c.Uint, **c.Uint, **c.Uint) c.Int {
	return 0
}

//go:linkname GetSubstring C.pcre_get_substring
func GetSubstring(*int8, *c.Int, c.Int, c.Int, **int8) c.Int

//go:linkname GetSubstring__1 C.pcre16_get_substring
func GetSubstring__1(*uint16, *c.Int, c.Int, c.Int, **uint16) c.Int

//go:linkname GetSubstring__2 C.pcre32_get_substring
func GetSubstring__2(*c.Uint, *c.Int, c.Int, c.Int, **c.Uint) c.Int

//go:linkname GetSubstringList C.pcre_get_substring_list
func GetSubstringList(*int8, *c.Int, c.Int, ***int8) c.Int

//go:linkname GetSubstringList__1 C.pcre16_get_substring_list
func GetSubstringList__1(*uint16, *c.Int, c.Int, ***uint16) c.Int

//go:linkname GetSubstringList__2 C.pcre32_get_substring_list
func GetSubstringList__2(*c.Uint, *c.Int, c.Int, ***c.Uint) c.Int

//go:linkname Maketables C.pcre_maketables
func Maketables() *int8

//go:linkname Maketables__1 C.pcre16_maketables
func Maketables__1() *int8

//go:linkname Maketables__2 C.pcre32_maketables
func Maketables__2() *int8

// llgo:link (*Pcre).Refcount C.pcre_refcount
func (recv_ *Pcre) Refcount(c.Int) c.Int {
	return 0
}

// llgo:link (*Pcre16).Refcount C.pcre16_refcount
func (recv_ *Pcre16) Refcount(c.Int) c.Int {
	return 0
}

// llgo:link (*Pcre32).Refcount C.pcre32_refcount
func (recv_ *Pcre32) Refcount(c.Int) c.Int {
	return 0
}

// llgo:link (*Pcre).Study C.pcre_study
func (recv_ *Pcre) Study(c.Int, **int8) *Extra {
	return nil
}

// llgo:link (*Pcre16).Study C.pcre16_study
func (recv_ *Pcre16) Study(c.Int, **int8) *Extra__1 {
	return nil
}

// llgo:link (*Pcre32).Study C.pcre32_study
func (recv_ *Pcre32) Study(c.Int, **int8) *Extra__2 {
	return nil
}

// llgo:link (*Extra).FreeStudy C.pcre_free_study
func (recv_ *Extra) FreeStudy() {
}

// llgo:link (*Extra__1).FreeStudy__1 C.pcre16_free_study
func (recv_ *Extra__1) FreeStudy__1() {
}

// llgo:link (*Extra__2).FreeStudy__2 C.pcre32_free_study
func (recv_ *Extra__2) FreeStudy__2() {
}

//go:linkname Version C.pcre_version
func Version() *int8

//go:linkname Version__1 C.pcre16_version
func Version__1() *int8

//go:linkname Version__2 C.pcre32_version
func Version__2() *int8

/* Utility functions for byte order swaps. */
// llgo:link (*Pcre).PatternToHostByteOrder C.pcre_pattern_to_host_byte_order
func (recv_ *Pcre) PatternToHostByteOrder(*Extra, *int8) c.Int {
	return 0
}

// llgo:link (*Pcre16).PatternToHostByteOrder C.pcre16_pattern_to_host_byte_order
func (recv_ *Pcre16) PatternToHostByteOrder(*Extra__1, *int8) c.Int {
	return 0
}

// llgo:link (*Pcre32).PatternToHostByteOrder C.pcre32_pattern_to_host_byte_order
func (recv_ *Pcre32) PatternToHostByteOrder(*Extra__2, *int8) c.Int {
	return 0
}

//go:linkname Utf16ToHostByteOrder C.pcre16_utf16_to_host_byte_order
func Utf16ToHostByteOrder(*uint16, *uint16, c.Int, *c.Int, c.Int) c.Int

//go:linkname Utf32ToHostByteOrder C.pcre32_utf32_to_host_byte_order
func Utf32ToHostByteOrder(*c.Uint, *c.Uint, c.Int, *c.Int, c.Int) c.Int

/* JIT compiler related functions. */
//go:linkname JitStackAlloc C.pcre_jit_stack_alloc
func JitStackAlloc(c.Int, c.Int) *JitStack

//go:linkname JitStackAlloc__1 C.pcre16_jit_stack_alloc
func JitStackAlloc__1(c.Int, c.Int) *JitStack__1

//go:linkname JitStackAlloc__2 C.pcre32_jit_stack_alloc
func JitStackAlloc__2(c.Int, c.Int) *JitStack__2

// llgo:link (*JitStack).JitStackFree C.pcre_jit_stack_free
func (recv_ *JitStack) JitStackFree() {
}

// llgo:link (*JitStack__1).JitStackFree__1 C.pcre16_jit_stack_free
func (recv_ *JitStack__1) JitStackFree__1() {
}

// llgo:link (*JitStack__2).JitStackFree__2 C.pcre32_jit_stack_free
func (recv_ *JitStack__2) JitStackFree__2() {
}

// llgo:link (*Extra).AssignJitStack C.pcre_assign_jit_stack
func (recv_ *Extra) AssignJitStack(JitCallback, unsafe.Pointer) {
}

// llgo:link (*Extra__1).AssignJitStack__1 C.pcre16_assign_jit_stack
func (recv_ *Extra__1) AssignJitStack__1(JitCallback__1, unsafe.Pointer) {
}

// llgo:link (*Extra__2).AssignJitStack__2 C.pcre32_assign_jit_stack
func (recv_ *Extra__2) AssignJitStack__2(JitCallback__2, unsafe.Pointer) {
}

//go:linkname JitFreeUnusedMemory C.pcre_jit_free_unused_memory
func JitFreeUnusedMemory()

//go:linkname JitFreeUnusedMemory__1 C.pcre16_jit_free_unused_memory
func JitFreeUnusedMemory__1()

//go:linkname JitFreeUnusedMemory__2 C.pcre32_jit_free_unused_memory
func JitFreeUnusedMemory__2()
