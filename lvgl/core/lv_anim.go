/**
 * @file lv_anim.h
 *
 */

package core

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

const (
	LV_ANIM_REPEAT_INFINITE   = 0xFFFFFFFF
	LV_ANIM_PLAYTIME_INFINITE = 0xFFFFFFFF
	LV_ANIM_PAUSE_FOREVER     = 0xFFFFFFFF
)

const (
	LV_ANIM_OFF = false
	LV_ANIM_ON  = true
)

// Get the current value during an animation
// typedef int32_t (*lv_anim_path_cb_t)(const lv_anim_t *);
//
//llgo:type C
type LvAnimPathCbT func(*LvAnimT) int32

/** Generic prototype of "animator" functions.
 * First parameter is the variable to animate.
 * Second parameter is the value to set.
 * Compatible with `lv_xxx_set_yyy(obj, value)` functions
 * The `x` in `_xcb_t` means it's not a fully generic prototype because
 * it doesn't receive `lv_anim_t *` as its first argument*/
//typedef int32_t (*lv_anim_path_cb_t)(const lv_anim_t *);
//llgo:type C
type LvAnimExecXcbT func(c.Pointer, int32)

/** Same as `lv_anim_exec_xcb_t` but receives `lv_anim_t *` as the first parameter.
 * It's more consistent but less convenient. Might be used by binding generator functions.*/
//typedef void (*lv_anim_custom_exec_cb_t)(lv_anim_t *, int32_t);
//llgo:type C
type LvAnimCustomExecCbT func(*LvAnimT, int32)

/** Callback to call when the animation is ready*/
//typedef void (*lv_anim_completed_cb_t)(lv_anim_t *);
//llgo:type C
type LvAnimCompletedCbT func(*LvAnimT)

/** Callback to call when the animation really stars (considering `delay`)*/
//typedef void (*lv_anim_start_cb_t)(lv_anim_t *);
//llgo:type C
type LvAnimStartCbT func(*LvAnimT)

/** Callback used when the animation values are relative to get the current value*/
//typedef int32_t (*lv_anim_get_value_cb_t)(lv_anim_t *);
//llgo:type C
type LvAnimGetValueCbT func(*LvAnimT) int32

/** Callback used when the animation is deleted*/
//typedef void (*lv_anim_deleted_cb_t)(lv_anim_t *);
//llgo:type C
type LvAnimDeletedCbT func(*LvAnimT)

/** Parameter used when path is custom_bezier */
type LvAnimBezier3ParaT struct {
	X1 c.Uint16T
	Y1 c.Uint16T
	X2 c.Uint16T
	Y2 c.Uint16T
}

type LvAnimT struct {
	Unused [0]byte
}

//go:linkname LvAnimInit C.lv_anim_init
func LvAnimInit(a *LvAnimT)

//go:linkname LvAnimSetVar C.lv_anim_set_var
func LvAnimSetVar(a *LvAnimT, var_ interface{})

//go:linkname LvAnimSetExecCb C.lv_anim_set_exec_cb
func LvAnimSetExecCb(a *LvAnimT, exec_cb LvAnimExecXcbT)

//go:linkname LvAnimSetDuration C.lv_anim_set_duration
func LvAnimSetDuration(a *LvAnimT, duration uint32)

//go:linkname LvAnimSetDelay C.lv_anim_set_delay
func LvAnimSetDelay(a *LvAnimT, delay uint32)

//go:linkname LvAnimResume C.lv_anim_resume
func LvAnimResume(a *LvAnimT)

//go:linkname LvAnimPause C.lv_anim_pause
func LvAnimPause(a *LvAnimT)

//go:linkname LvAnimPauseFor C.lv_anim_pause_for
func LvAnimPauseFor(a *LvAnimT, ms uint32)

//go:linkname LvAnimIsPaused C.lv_anim_is_paused
func LvAnimIsPaused(a *LvAnimT) bool

//go:linkname LvAnimSetValues C.lv_anim_set_values
func LvAnimSetValues(a *LvAnimT, start, end int32)

//go:linkname LvAnimSetCustomExecCb C.lv_anim_set_custom_exec_cb
func LvAnimSetCustomExecCb(a *LvAnimT, exec_cb LvAnimCustomExecCbT)

//go:linkname LvAnimSetPathCb C.lv_anim_set_path_cb
func LvAnimSetPathCb(a *LvAnimT, path_cb LvAnimPathCbT)

//go:linkname LvAnimSetStartCb C.lv_anim_set_start_cb
func LvAnimSetStartCb(a *LvAnimT, start_cb LvAnimStartCbT)

//go:linkname LvAnimSetGetValueCb C.lv_anim_set_get_value_cb
func LvAnimSetGetValueCb(a *LvAnimT, get_value_cb LvAnimGetValueCbT)

//go:linkname LvAnimSetCompletedCb C.lv_anim_set_completed_cb
func LvAnimSetCompletedCb(a *LvAnimT, completed_cb LvAnimCompletedCbT)

//go:linkname LvAnimSetDeletedCb C.lv_anim_set_deleted_cb
func LvAnimSetDeletedCb(a *LvAnimT, deleted_cb LvAnimDeletedCbT)

//go:linkname LvAnimSetReverseDuration C.lv_anim_set_reverse_duration
func LvAnimSetReverseDuration(a *LvAnimT, duration uint32)

//go:linkname LvAnimSetReverseTime C.lv_anim_set_reverse_time
func LvAnimSetReverseTime(a *LvAnimT, duration uint32)

//go:linkname LvAnimSetReverseDelay C.lv_anim_set_reverse_delay
func LvAnimSetReverseDelay(a *LvAnimT, delay uint32)

//go:linkname LvAnimSetRepeatCount C.lv_anim_set_repeat_count
func LvAnimSetRepeatCount(a *LvAnimT, cnt uint32)

//go:linkname LvAnimSetRepeatDelay C.lv_anim_set_repeat_delay
func LvAnimSetRepeatDelay(a *LvAnimT, delay uint32)

//go:linkname LvAnimSetEarlyApply C.lv_anim_set_early_apply
func LvAnimSetEarlyApply(a *LvAnimT, en bool)

//go:linkname LvAnimSetUserData C.lv_anim_set_user_data
func LvAnimSetUserData(a *LvAnimT, user_data interface{})

//go:linkname LvAnimSetBezier3Param C.lv_anim_set_bezier3_param
func LvAnimSetBezier3Param(a *LvAnimT, x1, y1, x2, y2 int16)

//go:linkname LvAnimStart C.lv_anim_start
func LvAnimStart(a *LvAnimT) *LvAnimT

//go:linkname LvAnimGetDelay C.lv_anim_get_delay
func LvAnimGetDelay(a *LvAnimT) uint32

//go:linkname LvAnimGetPlaytime C.lv_anim_get_playtime
func LvAnimGetPlaytime(a *LvAnimT) uint32

//go:linkname LvAnimGetTime C.lv_anim_get_time
func LvAnimGetTime(a *LvAnimT) uint32

//go:linkname LvAnimGetRepeatCount C.lv_anim_get_repeat_count
func LvAnimGetRepeatCount(a *LvAnimT) uint32

//go:linkname LvAnimGetUserData C.lv_anim_get_user_data
func LvAnimGetUserData(a *LvAnimT) interface{}

//go:linkname LvAnimDelete C.lv_anim_delete
func LvAnimDelete(var_ interface{}, exec_cb LvAnimExecXcbT) bool

//go:linkname LvAnimDeleteAll C.lv_anim_delete_all
func LvAnimDeleteAll()

//go:linkname LvAnimGet C.lv_anim_get
func LvAnimGet(var_ interface{}, exec_cb LvAnimExecXcbT) *LvAnimT

//go:linkname LvAnimGetTimer C.lv_anim_get_timer
func LvAnimGetTimer() *LvTimerT

//go:linkname LvAnimCustomDelete C.lv_anim_custom_delete
func LvAnimCustomDelete(a *LvAnimT, exec_cb LvAnimCustomExecCbT) bool

//go:linkname LvAnimCustomGet C.lv_anim_custom_get
func LvAnimCustomGet(a *LvAnimT, exec_cb LvAnimCustomExecCbT) *LvAnimT

//go:linkname LvAnimCountRunning C.lv_anim_count_running
func LvAnimCountRunning() uint16

//go:linkname LvAnimSpeed C.lv_anim_speed
func LvAnimSpeed(speed uint32) uint32

//go:linkname LvAnimSpeedClamped C.lv_anim_speed_clamped
func LvAnimSpeedClamped(speed, min_time, max_time uint32) uint32

//go:linkname LvAnimResolveSpeed C.lv_anim_resolve_speed
func LvAnimResolveSpeed(speed uint32, start, end int32) uint32

//go:linkname LvAnimSpeedToTime C.lv_anim_speed_to_time
func LvAnimSpeedToTime(speed uint32, start, end int32) uint32

//go:linkname LvAnimRefrNow C.lv_anim_refr_now
func LvAnimRefrNow()

//go:linkname LvAnimPathLinear C.lv_anim_path_linear
func LvAnimPathLinear(a *LvAnimT) int32

//go:linkname LvAnimPathEaseIn C.lv_anim_path_ease_in
func LvAnimPathEaseIn(a *LvAnimT) int32

//go:linkname LvAnimPathEaseOut C.lv_anim_path_ease_out
func LvAnimPathEaseOut(a *LvAnimT) int32

//go:linkname LvAnimPathEaseInOut C.lv_anim_path_ease_in_out
func LvAnimPathEaseInOut(a *LvAnimT) int32

//go:linkname LvAnimPathOvershoot C.lv_anim_path_overshoot
func LvAnimPathOvershoot(a *LvAnimT) int32

//go:linkname LvAnimPathBounce C.lv_anim_path_bounce
func LvAnimPathBounce(a *LvAnimT) int32

//go:linkname LvAnimPathStep C.lv_anim_path_step
func LvAnimPathStep(a *LvAnimT) int32

//go:linkname LvAnimPathCustomBezier3 C.lv_anim_path_custom_bezier3
func LvAnimPathCustomBezier3(a *LvAnimT) int32
