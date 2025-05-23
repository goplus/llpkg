package widgets

import (
	"unsafe"
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

//go:linkname LvCanvasCreate C.lv_canvas_create
func LvCanvasCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvCanvasSetBuffer C.lv_canvas_set_buffer
func LvCanvasSetBuffer(obj *core.LvObjT, buf unsafe.Pointer, w c.Int32T, h c.Int32T, cf core.LvColorFormatT)

//go:linkname LvCanvasSetDrawBuf C.lv_canvas_set_draw_buf
func LvCanvasSetDrawBuf(obj *core.LvObjT, draw_buf *core.LvDrawBufT)

//go:linkname LvCanvasSetPx C.lv_canvas_set_px
func LvCanvasSetPx(obj *core.LvObjT, x c.Int32T, y c.Int32T, color core.LvColorT, opa core.LvOpaT)

//go:linkname LvCanvasSetPalette C.lv_canvas_set_palette
func LvCanvasSetPalette(obj *core.LvObjT, index c.Uint8T, color core.LvColor32T)

//go:linkname LvCanvasGetDrawBuf C.lv_canvas_get_draw_buf
func LvCanvasGetDrawBuf(obj *core.LvObjT) *core.LvDrawBufT

//go:linkname LvCanvasGetPx C.lv_canvas_get_px
func LvCanvasGetPx(obj *core.LvObjT, x c.Int32T, y c.Int32T) core.LvColor32T

//go:linkname LvCanvasGetImage C.lv_canvas_get_image
func LvCanvasGetImage(canvas *core.LvObjT) *core.LvImageDscT

//go:linkname LvCanvasGetBuf C.lv_canvas_get_buf
func LvCanvasGetBuf(canvas *core.LvObjT) unsafe.Pointer

//go:linkname LvCanvasCopyBuf C.lv_canvas_copy_buf
func LvCanvasCopyBuf(obj *core.LvObjT, canvas_area *core.LvAreaT, dest_buf *core.LvDrawBufT, dest_area *core.LvAreaT)

//go:linkname LvCanvasFillBg C.lv_canvas_fill_bg
func LvCanvasFillBg(obj *core.LvObjT, color core.LvColorT, opa core.LvOpaT)

//go:linkname LvCanvasInitLayer C.lv_canvas_init_layer
func LvCanvasInitLayer(canvas *core.LvObjT, layer *core.LvLayerT)

//go:linkname LvCanvasFinishLayer C.lv_canvas_finish_layer
func LvCanvasFinishLayer(canvas *core.LvObjT, layer *core.LvLayerT)

//go:linkname LvCanvasBufSize C.lv_canvas_buf_size
func LvCanvasBufSize(w c.Int32T, h c.Int32T, bpp c.Uint8T, stride c.Uint8T) c.Uint32T
