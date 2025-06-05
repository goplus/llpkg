package quickjs

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const UTF8_CHAR_LEN_MAX = 6

type BOOL c.Int

const (
	FALSE c.Int = 0
	TRUE  c.Int = 1
)

//go:linkname Pstrcpy C.pstrcpy
func Pstrcpy(buf *c.Char, buf_size c.Int, str *c.Char)

//go:linkname Pstrcat C.pstrcat
func Pstrcat(buf *c.Char, buf_size c.Int, s *c.Char) *c.Char

//go:linkname Strstart C.strstart
func Strstart(str *c.Char, val *c.Char, ptr **c.Char) c.Int

//go:linkname HasSuffix C.has_suffix
func HasSuffix(str *c.Char, suffix *c.Char) c.Int

type PackedU64 struct {
	V c.Uint64T
}

type PackedU32 struct {
	V c.Uint32T
}

type PackedU16 struct {
	V c.Uint16T
}

// llgo:type C
type DynBufReallocFunc func(c.Pointer, c.Pointer, c.SizeT) c.Pointer

type DynBuf struct {
	Buf           *c.Uint8T
	Size          c.SizeT
	AllocatedSize c.SizeT
	Error         BOOL
	ReallocFunc   *c.Pointer
	Opaque        c.Pointer
}

// llgo:link (*DynBuf).DbufInit C.dbuf_init
func (recv_ *DynBuf) DbufInit() {
}

// llgo:link (*DynBuf).DbufInit2 C.dbuf_init2
func (recv_ *DynBuf) DbufInit2(opaque c.Pointer, realloc_func DynBufReallocFunc) {
}

// llgo:link (*DynBuf).DbufRealloc C.dbuf_realloc
func (recv_ *DynBuf) DbufRealloc(new_size c.SizeT) c.Int {
	return 0
}

// llgo:link (*DynBuf).DbufWrite C.dbuf_write
func (recv_ *DynBuf) DbufWrite(offset c.SizeT, data *c.Uint8T, len c.SizeT) c.Int {
	return 0
}

// llgo:link (*DynBuf).DbufPut C.dbuf_put
func (recv_ *DynBuf) DbufPut(data *c.Uint8T, len c.SizeT) c.Int {
	return 0
}

// llgo:link (*DynBuf).DbufPutSelf C.dbuf_put_self
func (recv_ *DynBuf) DbufPutSelf(offset c.SizeT, len c.SizeT) c.Int {
	return 0
}

// llgo:link (*DynBuf).DbufPutc C.dbuf_putc
func (recv_ *DynBuf) DbufPutc(c c.Uint8T) c.Int {
	return 0
}

// llgo:link (*DynBuf).DbufPutstr C.dbuf_putstr
func (recv_ *DynBuf) DbufPutstr(str *c.Char) c.Int {
	return 0
}

// llgo:link (*DynBuf).DbufPrintf C.dbuf_printf
func (recv_ *DynBuf) DbufPrintf(fmt *c.Char, __llgo_va_list ...interface{}) c.Int {
	return 0
}

// llgo:link (*DynBuf).DbufFree C.dbuf_free
func (recv_ *DynBuf) DbufFree() {
}

//go:linkname UnicodeToUtf8 C.unicode_to_utf8
func UnicodeToUtf8(buf *c.Uint8T, c c.Uint) c.Int

//go:linkname UnicodeFromUtf8 C.unicode_from_utf8
func UnicodeFromUtf8(p *c.Uint8T, max_len c.Int, pp **c.Uint8T) c.Int

//go:linkname Rqsort C.rqsort
func Rqsort(base c.Pointer, nmemb c.SizeT, size c.SizeT, cmp func(c.Pointer, c.Pointer, c.Pointer) c.Int, arg c.Pointer)
