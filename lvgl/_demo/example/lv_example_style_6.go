package main

import (
	"log"
	"unsafe"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/lvgl/core"
	"github.com/goplus/llpkg/lvgl/display"
	"github.com/goplus/llpkg/lvgl/widgets"

	"github.com/goplus/llpkg/lvgl/_demo/example/assert"
)

func lv_example_style_6() {
	var style core.LvStyleT
	core.LvStyleInit(&style)

	/*Set a background color and a radius*/
	core.LvStyleSetRadius(&style, 5)
	core.LvStyleSetBgOpa(&style, core.LV_OPA_COVER)
	core.LvStyleSetBgColor(&style, core.LvPaletteLighten(core.LV_PALETTE_GREY, 3))
	core.LvStyleSetBorderWidth(&style, 2)
	core.LvStyleSetBorderColor(&style, core.LvPaletteMain(core.LV_PALETTE_BLUE))

	core.LvStyleSetImageRecolor(&style, core.LvPaletteMain(core.LV_PALETTE_BLUE))
	core.LvStyleSetImageRecolorOpa(&style, core.LV_OPA_50)
	core.LvStyleSetTransformRotation(&style, 300)

	/*Create an object with the new style*/
	obj := widgets.LvImageCreate(display.LvScreenActive())
	core.LvObjAddStyle(obj, &style, 0)

	img := assert.ImgCogwheelARGB
	log.Println(img)
	widgets.LvImageSetSrc(obj, (*c.Void)(unsafe.Pointer(&img)))

	core.LvObjCenter(obj)
}
