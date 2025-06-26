package widgets

import (
	"unsafe"
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

//go:linkname LvLottieCreate C.lv_lottie_create
func LvLottieCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvLottieSetBuffer C.lv_lottie_set_buffer
func LvLottieSetBuffer(obj *core.LvObjT, w c.Int32T, h c.Int32T, buf *c.Void)

//go:linkname LvLottieSetDrawBuf C.lv_lottie_set_draw_buf
func LvLottieSetDrawBuf(obj *core.LvObjT, draw_buf *core.LvDrawBufT)

//go:linkname LvLottieSetSrcData C.lv_lottie_set_src_data
func LvLottieSetSrcData(obj *core.LvObjT, src unsafe.Pointer, src_size c.SizeT)

//go:linkname LvLottieSetSrcFile C.lv_lottie_set_src_file
func LvLottieSetSrcFile(obj *core.LvObjT, src *c.Char)

//go:linkname LvLottieGetAnim C.lv_lottie_get_anim
func LvLottieGetAnim(obj *core.LvObjT) *core.LvAnimT
