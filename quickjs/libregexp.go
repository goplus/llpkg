package quickjs

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

//go:linkname LreCompile C.lre_compile
func LreCompile(plen *c.Int, error_msg *c.Char, error_msg_size c.Int, buf *c.Char, buf_len c.SizeT, re_flags c.Int, opaque c.Pointer) *c.Uint8T

//go:linkname LreGetCaptureCount C.lre_get_capture_count
func LreGetCaptureCount(bc_buf *c.Uint8T) c.Int

//go:linkname LreGetFlags C.lre_get_flags
func LreGetFlags(bc_buf *c.Uint8T) c.Int

//go:linkname LreGetGroupnames C.lre_get_groupnames
func LreGetGroupnames(bc_buf *c.Uint8T) *c.Char

//go:linkname LreExec C.lre_exec
func LreExec(capture **c.Uint8T, bc_buf *c.Uint8T, cbuf *c.Uint8T, cindex c.Int, clen c.Int, cbuf_type c.Int, opaque c.Pointer) c.Int

//go:linkname LreParseEscape C.lre_parse_escape
func LreParseEscape(pp **c.Uint8T, allow_utf16 c.Int) c.Int

//go:linkname LreIsSpace C.lre_is_space
func LreIsSpace(c c.Int) c.Int

/* must be provided by the user */
//go:linkname LreCheckStackOverflow C.lre_check_stack_overflow
func LreCheckStackOverflow(opaque c.Pointer, alloca_size c.SizeT) c.Int

//go:linkname LreRealloc C.lre_realloc
func LreRealloc(opaque c.Pointer, ptr c.Pointer, size c.SizeT) c.Pointer
