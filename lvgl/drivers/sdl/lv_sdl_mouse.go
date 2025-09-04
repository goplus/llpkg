package sdl

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

// lv_indev_t * lv_sdl_mouse_create(void);
//
//go:linkname LvSdlMouseCreate C.lv_sdl_mouse_create
func LvSdlMouseCreate() *c.Void

// lv_indev_t * lv_sdl_mousewheel_create(void);
//
//go:linkname LvSdlMouseWheelCreate C.lv_sdl_mousewheel_create
func LvSdlMouseWheelCreate() *c.Void

// lv_indev_t * lv_sdl_keyboard_create(void);
//
//go:linkname LvSdlKeyboardCreate C.lv_sdl_keyboard_create
func LvSdlKeyboardCreate() *c.Void
