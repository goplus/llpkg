package display

import (
	"unsafe"
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

/**
 * Get the active screen of the default display
 * @return          pointer to the active screen
 */
//lv_obj_t * lv_screen_active(void);
//
//go:linkname LvScreenActive C.lv_screen_active
func LvScreenActive() *core.LvObjT

//go:linkname LvDisplayCreate C.lv_display_create
func LvDisplayCreate(horRes c.Int32T, verRes c.Int32T) *core.LvDisplayT

//go:linkname LvDisplayDelete C.lv_display_delete
func LvDisplayDelete(disp *core.LvDisplayT)

//go:linkname LvDisplaySetDefault C.lv_display_set_default
func LvDisplaySetDefault(disp *core.LvDisplayT)

//go:linkname LvDisplayGetDefault C.lv_display_get_default
func LvDisplayGetDefault() *core.LvDisplayT

//go:linkname LvDisplayGetNext C.lv_display_get_next
func LvDisplayGetNext(disp *core.LvDisplayT) *core.LvDisplayT

//go:linkname LvDisplaySetResolution C.lv_display_set_resolution
func LvDisplaySetResolution(disp *core.LvDisplayT, horRes c.Int32T, verRes c.Int32T)

//go:linkname LvDisplaySetPhysicalResolution C.lv_display_set_physical_resolution
func LvDisplaySetPhysicalResolution(disp *core.LvDisplayT, horRes c.Int32T, verRes c.Int32T)

//go:linkname LvDisplaySetOffset C.lv_display_set_offset
func LvDisplaySetOffset(disp *core.LvDisplayT, x c.Int32T, y c.Int32T)

//go:linkname LvDisplaySetRotation C.lv_display_set_rotation
func LvDisplaySetRotation(disp *core.LvDisplayT, rotation c.Int32T)

//go:linkname LvDisplaySetDpi C.lv_display_set_dpi
func LvDisplaySetDpi(disp *core.LvDisplayT, dpi c.Int32T)

//go:linkname LvDisplayGetHorizontalResolution C.lv_display_get_horizontal_resolution
func LvDisplayGetHorizontalResolution(disp *core.LvDisplayT) c.Int32T

//go:linkname LvDisplayGetVerticalResolution C.lv_display_get_vertical_resolution
func LvDisplayGetVerticalResolution(disp *core.LvDisplayT) c.Int32T

//go:linkname LvDisplayGetPhysicalHorizontalResolution C.lv_display_get_physical_horizontal_resolution
func LvDisplayGetPhysicalHorizontalResolution(disp *core.LvDisplayT) c.Int32T

//go:linkname LvDisplayGetPhysicalVerticalResolution C.lv_display_get_physical_vertical_resolution
func LvDisplayGetPhysicalVerticalResolution(disp *core.LvDisplayT) c.Int32T

//go:linkname LvDisplayGetOffsetX C.lv_display_get_offset_x
func LvDisplayGetOffsetX(disp *core.LvDisplayT) c.Int32T

//go:linkname LvDisplayGetOffsetY C.lv_display_get_offset_y
func LvDisplayGetOffsetY(disp *core.LvDisplayT) c.Int32T

//go:linkname LvDisplayGetRotation C.lv_display_get_rotation
func LvDisplayGetRotation(disp *core.LvDisplayT) c.Int32T

//go:linkname LvDisplayGetDpi C.lv_display_get_dpi
func LvDisplayGetDpi(disp *core.LvDisplayT) c.Int32T

//go:linkname LvDisplaySetBuffers C.lv_display_set_buffers
func LvDisplaySetBuffers(disp *core.LvDisplayT, buf1 unsafe.Pointer, buf2 unsafe.Pointer, bufSize c.Uint32T, renderMode c.Int32T)

//go:linkname LvDisplaySetBuffersWithStride C.lv_display_set_buffers_with_stride
func LvDisplaySetBuffersWithStride(disp *core.LvDisplayT, buf1 unsafe.Pointer, buf2 unsafe.Pointer, bufSize c.Uint32T, stride c.Uint32T, renderMode c.Int32T)

//go:linkname LvDisplaySetColorFormat C.lv_display_set_color_format
func LvDisplaySetColorFormat(disp *core.LvDisplayT, colorFormat c.Int32T)

//go:linkname LvDisplayGetColorFormat C.lv_display_get_color_format
func LvDisplayGetColorFormat(disp *core.LvDisplayT) c.Int32T

//go:linkname LvDisplaySetAntialiasing C.lv_display_set_antialiasing
func LvDisplaySetAntialiasing(disp *core.LvDisplayT, en c.Char)

//go:linkname LvDisplayGetAntialiasing C.lv_display_get_antialiasing
func LvDisplayGetAntialiasing(disp *core.LvDisplayT) c.Char

//go:linkname LvDisplayGetScreenActive C.lv_display_get_screen_active
func LvDisplayGetScreenActive(disp *core.LvDisplayT) *core.LvObjT

//go:linkname LvDisplayGetScreenPrev C.lv_display_get_screen_prev
func LvDisplayGetScreenPrev(disp *core.LvDisplayT) *core.LvObjT

//go:linkname LvDisplayGetLayerTop C.lv_display_get_layer_top
func LvDisplayGetLayerTop(disp *core.LvDisplayT) *core.LvObjT

//go:linkname LvDisplayGetLayerSys C.lv_display_get_layer_sys
func LvDisplayGetLayerSys(disp *core.LvDisplayT) *core.LvObjT

//go:linkname LvDisplayGetLayerBottom C.lv_display_get_layer_bottom
func LvDisplayGetLayerBottom(disp *core.LvDisplayT) *core.LvObjT

//go:linkname LvScreenLoad C.lv_screen_load
func LvScreenLoad(scr *core.LvObjT)

//go:linkname LvScreenLoadAnim C.lv_screen_load_anim
func LvScreenLoadAnim(scr *core.LvObjT, animType c.Int32T, time c.Uint32T, delay c.Uint32T, autoDel c.Char)

//go:linkname LvLayerTop C.lv_layer_top
func LvLayerTop() *core.LvObjT

//go:linkname LvLayerSys C.lv_layer_sys
func LvLayerSys() *core.LvObjT

//go:linkname LvLayerBottom C.lv_layer_bottom
func LvLayerBottom() *core.LvObjT

//go:linkname LvDisplayAddEventCb C.lv_display_add_event_cb
func LvDisplayAddEventCb(disp *core.LvDisplayT, eventCb core.LvEventCbT, filter core.LvEventCodeT, userData unsafe.Pointer)

//go:linkname LvDisplayGetEventCount C.lv_display_get_event_count
func LvDisplayGetEventCount(disp *core.LvDisplayT) c.Uint32T

//go:linkname LvDisplayGetEventDsc C.lv_display_get_event_dsc
func LvDisplayGetEventDsc(disp *core.LvDisplayT, index c.Uint32T) *core.LvEventDscT

//go:linkname LvDisplayDeleteEvent C.lv_display_delete_event
func LvDisplayDeleteEvent(disp *core.LvDisplayT, index c.Uint32T) c.Char

//go:linkname LvDisplayRemoveEventCbWithUserData C.lv_display_remove_event_cb_with_user_data
func LvDisplayRemoveEventCbWithUserData(disp *core.LvDisplayT, eventCb core.LvEventCbT, userData unsafe.Pointer) c.Uint32T

//go:linkname LvDisplaySendEvent C.lv_display_send_event
func LvDisplaySendEvent(disp *core.LvDisplayT, code core.LvEventCodeT, param unsafe.Pointer) core.LvResultT

//go:linkname LvEventGetInvalidatedArea C.lv_event_get_invalidated_area
func LvEventGetInvalidatedArea(e *core.LvEventT) *core.LvAreaT

//go:linkname LvDisplaySetTheme C.lv_display_set_theme
func LvDisplaySetTheme(disp *core.LvDisplayT, th *core.LvThemeT)

//go:linkname LvDisplayGetTheme C.lv_display_get_theme
func LvDisplayGetTheme(disp *core.LvDisplayT) *core.LvThemeT

//go:linkname LvDisplayGetInactiveTime C.lv_display_get_inactive_time
func LvDisplayGetInactiveTime(disp *core.LvDisplayT) c.Uint32T

//go:linkname LvDisplayTriggerActivity C.lv_display_trigger_activity
func LvDisplayTriggerActivity(disp *core.LvDisplayT)

//go:linkname LvDisplayEnableInvalidation C.lv_display_enable_invalidation
func LvDisplayEnableInvalidation(disp *core.LvDisplayT, en c.Char)

//go:linkname LvDisplayIsInvalidationEnabled C.lv_display_is_invalidation_enabled
func LvDisplayIsInvalidationEnabled(disp *core.LvDisplayT) c.Char

//go:linkname LvDisplayGetRefrTimer C.lv_display_get_refr_timer
func LvDisplayGetRefrTimer(disp *core.LvDisplayT) *core.LvTimerT

//go:linkname LvDisplayDeleteRefrTimer C.lv_display_delete_refr_timer
func LvDisplayDeleteRefrTimer(disp *core.LvDisplayT)

//go:linkname LvDisplaySetUserData C.lv_display_set_user_data
func LvDisplaySetUserData(disp *core.LvDisplayT, userData unsafe.Pointer)

//go:linkname LvDisplaySetDriverData C.lv_display_set_driver_data
func LvDisplaySetDriverData(disp *core.LvDisplayT, driverData unsafe.Pointer)

//go:linkname LvDisplayGetUserData C.lv_display_get_user_data
func LvDisplayGetUserData(disp *core.LvDisplayT) unsafe.Pointer

//go:linkname LvDisplayGetDriverData C.lv_display_get_driver_data
func LvDisplayGetDriverData(disp *core.LvDisplayT) unsafe.Pointer

//go:linkname LvDisplayGetBufActive C.lv_display_get_buf_active
func LvDisplayGetBufActive(disp *core.LvDisplayT) *core.LvDrawBufT

//go:linkname LvDisplayRotateArea C.lv_display_rotate_area
func LvDisplayRotateArea(disp *core.LvDisplayT, area *core.LvAreaT)

//go:linkname LvDisplayGetDrawBufSize C.lv_display_get_draw_buf_size
func LvDisplayGetDrawBufSize(disp *core.LvDisplayT) c.Uint32T

//go:linkname LvDisplayGetInvalidatedDrawBufSize C.lv_display_get_invalidated_draw_buf_size
func LvDisplayGetInvalidatedDrawBufSize(disp *core.LvDisplayT, width c.Uint32T, height c.Uint32T) c.Uint32T

//go:linkname LvDpx C.lv_dpx
func LvDpx(n c.Int32T) c.Int32T

//go:linkname LvDisplayDpx C.lv_display_dpx
func LvDisplayDpx(disp *core.LvDisplayT, n c.Int32T) c.Int32T

// 定义 LV_HOR_RES 变量
var LV_HOR_RES int

// 定义 LV_VER_RES 变量
var LV_VER_RES int

// 初始化 LV_HOR_RES 和 LV_VER_RES
func init() {
	LV_HOR_RES = int(LvDisplayGetHorizontalResolution(LvDisplayGetDefault()))
	LV_VER_RES = int(LvDisplayGetVerticalResolution(LvDisplayGetDefault()))
}

/**
 * See `lv_dpx()` and `lv_display_dpx()`.
 * Same as Android's DIP. (Different name is chosen to avoid mistype between LV_DPI and LV_DIP)
 *
 * - 40 dip is 40 px on a 160 DPI screen (distance = 1/4 inch).
 * - 40 dip is 80 px on a 320 DPI screen (distance still = 1/4 inch).
 *
 * @sa https://stackoverflow.com/questions/2025282/what-is-the-difference-between-px-dip-dp-and-sp
 */
// #define LV_DPX_CALC(dpi, n)   ((n) == 0 ? 0 :LV_MAX((( (dpi) * (n) + 80) / 160), 1)) /*+80 for rounding*/
// #define LV_DPX(n)   LV_DPX_CALC(lv_display_get_dpi(NULL), n)

// LV_MAX 计算最大值
func LV_MAX(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 计算 LV_DPX
func LV_DPX_CALC(dpi, n int) int {
	if n == 0 {
		return 0
	}
	return LV_MAX(((dpi*n + 80) / 160), 1) // +80 用于四舍五入
}

// LV_DPX 计算 DIP 到 PX
func LV_DPX(n int) int {
	return LV_DPX_CALC(int(LvDisplayGetDpi(nil)), n)
}
