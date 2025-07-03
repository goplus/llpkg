package main

import (
	"time"

	"github.com/goplus/lib/c"
	"github.com/goplus/llpkg/lvgl"
)

func main() {
	// Initialize LVGL
	lvgl.LvInit()
	disp := lvgl.LvSdlWindowCreate(480, 320)
	lvgl.LvSdlMouseCreate()
	lvgl.LvSdlKeyboardCreate()

	disp.LvSdlWindowSetZoom(1.0)
	disp.LvSdlWindowSetTitle(c.Str("Profile Page"))

	//lv_example_style_1() // 背景
	//lv_example_style_2() // 渐变 -- 有问题
	//lv_example_style_3() // 边框
	//lv_example_style_4() // 轮廓
	//lv_example_style_5() // 阴影
	//lv_example_style_6() // 图片
	lv_example_get_started_2() // 按钮
	//lv_example_get_started_3() // 滑块
	//lv_example_get_started_4() // 标签
	//lv_example_anim_1() // 动画
	// Main event loop
	for {
		lvgl.LvTimerHandler()
		time.Sleep(time.Millisecond * 5)
	}
}
