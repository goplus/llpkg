package main

import (
	"time"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/lvgl"
	"github.com/goplus/llpkg/lvgl/drivers/sdl"
)

func main() {
	lvgl.LvglInit()

	disp := sdl.LvSdlWindowCreate(480, 320)
	sdl.LvSdlWindowSetZoom(disp, 1.0)
	sdl.LvSdlWindowSetTitle(disp, c.Str("Music"))

	lv_demo_music()

	for {
		lvgl.LvTimerHandler()
		time.Sleep(time.Millisecond * 5)
	}

}
