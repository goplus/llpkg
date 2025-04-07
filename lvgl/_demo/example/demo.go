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
	sdl.LvSdlMouseCreate()
	sdl.LvSdlKeyboardCreate()

	sdl.LvSdlWindowSetZoom(disp, 1.0)
	sdl.LvSdlWindowSetTitle(disp, c.Str("Profile Page"))
	log.Println("disp", disp)

	//lv_example_style_1() // 背景
	//lv_example_style_2() // 渐变 -- 有问题
	//lv_example_style_3() // 边框
	//lv_example_style_4() // 轮廓
	lv_example_style_5() // 阴影
	//lv_example_style_6() // 图片
	//lv_example_get_started_2() // 按钮
	//lv_example_get_started_3() // 滑块
	//lv_example_get_started_4() // 标签
	// Main event loop
	for {
		lvgl.LvTimerHandler()
		time.Sleep(time.Millisecond * 5)
	}
}
