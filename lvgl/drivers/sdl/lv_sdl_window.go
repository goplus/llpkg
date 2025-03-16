package sdl

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/lvgl/core"
)

// lv_display_t * lv_sdl_window_create(int32_t hor_res, int32_t ver_res);
//
//go:linkname LvSdlWindowCreate C.lv_sdl_window_create
func LvSdlWindowCreate(hor_res int32, ver_res int32) *core.LvDisplayT

// void lv_sdl_window_set_resizeable(lv_display_t * disp, bool value);
//
//go:linkname LvSdlWindowSetResizeable C.lv_sdl_window_set_resizeable
func LvSdlWindowSetResizeable(disp *c.Void, value bool)

// void lv_sdl_window_set_zoom(lv_display_t * disp, float zoom);
//
//go:linkname LvSdlWindowSetZoom C.lv_sdl_window_set_zoom
func LvSdlWindowSetZoom(disp *core.LvDisplayT, zoom c.Float)

// float lv_sdl_window_get_zoom(lv_display_t * disp);
//
//go:linkname LvSdlWindowGetZoom C.lv_sdl_window_get_zoom
func LvSdlWindowGetZoom(disp *core.LvDisplayT) c.Float

// void lv_sdl_window_set_title(lv_display_t * disp, const char * title);
//
//go:linkname LvSdlWindowSetTitle C.lv_sdl_window_set_title
func LvSdlWindowSetTitle(disp *core.LvDisplayT, title *c.Char)

// void * lv_sdl_window_get_renderer(lv_display_t * disp);
//
//go:linkname LvSdlWindowGetRenderer C.lv_sdl_window_get_renderer
func LvSdlWindowGetRenderer(disp *core.LvDisplayT) *c.Char

// void lv_sdl_quit(void);
//
//go:linkname LvSdlQuit C.lv_sdl_quit
func LvSdlQuit() {}
