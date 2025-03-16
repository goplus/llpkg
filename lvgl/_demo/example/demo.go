package main

import (
	"log"
	"time"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/lvgl"
	"github.com/goplus/llpkg/lvgl/drivers/sdl"
)

func main() {
	// Initialize LVGL
	lvgl.LvglInit()

	// Create a display
	disp := sdl.LvSdlWindowCreate(480, 320)
	sdl.LvSdlWindowSetZoom(disp, 1.0)
	sdl.LvSdlWindowSetTitle(disp, c.Str("Profile Page"))
	log.Println("disp", disp)

	//lv_example_style_1()
	//lv_example_get_started_2()
	lv_example_get_started_3()
	// Main event loop
	for {
		lvgl.LvTimerHandler()
		time.Sleep(time.Millisecond * 5)
	}
}
