package quickjs

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const UTF8_CHAR_LEN_MAX = 6

type BOOL c.Int

const (
	FALSE c.Int = 0
	TRUE  c.Int = 1
)

//go:linkname Pstrcpy C.pstrcpy
func Pstrcpy(buf *int8, buf_size c.Int, str *int8)

//go:linkname Pstrcat C.pstrcat
func Pstrcat(buf *int8, buf_size c.Int, s *int8) *int8

//go:linkname Strstart C.strstart
func Strstart(str *int8, val *int8, ptr **int8) c.Int

//go:linkname HasSuffix C.has_suffix
func HasSuffix(str *int8, suffix *int8) c.Int

type PackedU64 struct {
	V uint64
}

type PackedU32 struct {
	V uint32
}

type PackedU16 struct {
	V uint16
}

// llgo:type C
type DynBufReallocFunc func(unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer

type DynBuf struct {
	Buf           *uint8
	Size          uintptr
	AllocatedSize uintptr
	Error         BOOL
	ReallocFunc   *unsafe.Pointer
	Opaque        unsafe.Pointer
}

// llgo:link (*DynBuf).DbufInit C.dbuf_init
func (recv_ *DynBuf) DbufInit() {
}

// llgo:link (*DynBuf).DbufInit2 C.dbuf_init2
func (recv_ *DynBuf) DbufInit2(opaque unsafe.Pointer, realloc_func DynBufReallocFunc) {
}

// llgo:link (*DynBuf).DbufRealloc C.dbuf_realloc
func (recv_ *DynBuf) DbufRealloc(new_size uintptr) c.Int {
	return 0
}

// llgo:link (*DynBuf).DbufWrite C.dbuf_write
func (recv_ *DynBuf) DbufWrite(offset uintptr, data *uint8, len uintptr) c.Int {
	return 0
}

// llgo:link (*DynBuf).DbufPut C.dbuf_put
func (recv_ *DynBuf) DbufPut(data *uint8, len uintptr) c.Int {
	return 0
}

// llgo:link (*DynBuf).DbufPutSelf C.dbuf_put_self
func (recv_ *DynBuf) DbufPutSelf(offset uintptr, len uintptr) c.Int {
	return 0
}

// llgo:link (*DynBuf).DbufPutc C.dbuf_putc
func (recv_ *DynBuf) DbufPutc(c uint8) c.Int {
	return 0
}

// llgo:link (*DynBuf).DbufPutstr C.dbuf_putstr
func (recv_ *DynBuf) DbufPutstr(str *int8) c.Int {
	return 0
}

// llgo:link (*DynBuf).DbufPrintf C.dbuf_printf
func (recv_ *DynBuf) DbufPrintf(fmt *int8, __llgo_va_list ...interface{}) c.Int {
	return 0
}

// llgo:link (*DynBuf).DbufFree C.dbuf_free
func (recv_ *DynBuf) DbufFree() {
}

//go:linkname UnicodeToUtf8 C.unicode_to_utf8
func UnicodeToUtf8(buf *uint8, c c.Uint) c.Int

//go:linkname UnicodeFromUtf8 C.unicode_from_utf8
func UnicodeFromUtf8(p *uint8, max_len c.Int, pp **uint8) c.Int

//go:linkname Rqsort C.rqsort
func Rqsort(base unsafe.Pointer, nmemb uintptr, size uintptr, cmp func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) c.Int, arg unsafe.Pointer)
