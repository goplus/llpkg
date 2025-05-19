package quickjs

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

//go:linkname LreCompile C.lre_compile
func LreCompile(plen *c.Int, error_msg *int8, error_msg_size c.Int, buf *int8, buf_len uintptr, re_flags c.Int, opaque unsafe.Pointer) *uint8

//go:linkname LreGetCaptureCount C.lre_get_capture_count
func LreGetCaptureCount(bc_buf *uint8) c.Int

//go:linkname LreGetFlags C.lre_get_flags
func LreGetFlags(bc_buf *uint8) c.Int

//go:linkname LreGetGroupnames C.lre_get_groupnames
func LreGetGroupnames(bc_buf *uint8) *int8

//go:linkname LreExec C.lre_exec
func LreExec(capture **uint8, bc_buf *uint8, cbuf *uint8, cindex c.Int, clen c.Int, cbuf_type c.Int, opaque unsafe.Pointer) c.Int

//go:linkname LreParseEscape C.lre_parse_escape
func LreParseEscape(pp **uint8, allow_utf16 c.Int) c.Int

//go:linkname LreIsSpace C.lre_is_space
func LreIsSpace(c c.Int) c.Int

/* must be provided by the user */
//go:linkname LreCheckStackOverflow C.lre_check_stack_overflow
func LreCheckStackOverflow(opaque unsafe.Pointer, alloca_size uintptr) c.Int

//go:linkname LreRealloc C.lre_realloc
func LreRealloc(opaque unsafe.Pointer, ptr unsafe.Pointer, size uintptr) unsafe.Pointer

//go:linkname LreJsIsIdentNext C.lre_js_is_ident_next
func LreJsIsIdentNext(c c.Int) c.Int
