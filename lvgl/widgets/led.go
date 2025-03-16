package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

// LED brightness constants
const (
	LV_LED_BRIGHT_MIN = 80  // Brightness when the LED is OFF
	LV_LED_BRIGHT_MAX = 255 // Brightness when the LED is ON
)

//go:linkname LvLedCreate C.lv_led_create
func LvLedCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvLedSetColor C.lv_led_set_color
func LvLedSetColor(led *core.LvObjT, color core.LvColorT)

//go:linkname LvLedSetBrightness C.lv_led_set_brightness
func LvLedSetBrightness(led *core.LvObjT, bright c.Uint8T)

//go:linkname LvLedOn C.lv_led_on
func LvLedOn(led *core.LvObjT)

//go:linkname LvLedOff C.lv_led_off
func LvLedOff(led *core.LvObjT)

//go:linkname LvLedToggle C.lv_led_toggle
func LvLedToggle(led *core.LvObjT)

//go:linkname LvLedGetBrightness C.lv_led_get_brightness
func LvLedGetBrightness(obj *core.LvObjT) c.Uint8T
