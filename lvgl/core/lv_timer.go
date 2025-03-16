package core

import (
	"unsafe"
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

// LV_NO_TIMER_READY defines the value indicating no timer is ready
const LV_NO_TIMER_READY = 0xFFFFFFFF

// LvTimerCbT represents the timer callback function type
//
//llgo:type C
type LvTimerCbT func(*LvTimerT)

// LvTimerHandlerResumeCbT represents the timer handler resume callback function type
//
//llgo:type C
type LvTimerHandlerResumeCbT func(*c.Void)

//go:linkname LvTimerHandler C.lv_timer_handler
func LvTimerHandler() c.Uint32T

//go:linkname LvTimerHandlerRunInPeriod C.lv_timer_handler_run_in_period
func LvTimerHandlerRunInPeriod(period c.Uint32T) c.Uint32T

//go:linkname LvTimerPeriodicHandler C.lv_timer_periodic_handler
func LvTimerPeriodicHandler()

//go:linkname LvTimerHandlerSetResumeCb C.lv_timer_handler_set_resume_cb
func LvTimerHandlerSetResumeCb(cb LvTimerHandlerResumeCbT, data unsafe.Pointer)

//go:linkname LvTimerCreateBasic C.lv_timer_create_basic
func LvTimerCreateBasic() *LvTimerT

//go:linkname LvTimerCreate C.lv_timer_create
func LvTimerCreate(timerXcb LvTimerCbT, period c.Uint32T, userData unsafe.Pointer) *LvTimerT

//go:linkname LvTimerDelete C.lv_timer_delete
func LvTimerDelete(timer *LvTimerT)

//go:linkname LvTimerPause C.lv_timer_pause
func LvTimerPause(timer *LvTimerT)

//go:linkname LvTimerResume C.lv_timer_resume
func LvTimerResume(timer *LvTimerT)

//go:linkname LvTimerSetCb C.lv_timer_set_cb
func LvTimerSetCb(timer *LvTimerT, timerCb LvTimerCbT)

//go:linkname LvTimerSetPeriod C.lv_timer_set_period
func LvTimerSetPeriod(timer *LvTimerT, period c.Uint32T)

//go:linkname LvTimerReady C.lv_timer_ready
func LvTimerReady(timer *LvTimerT)

//go:linkname LvTimerSetRepeatCount C.lv_timer_set_repeat_count
func LvTimerSetRepeatCount(timer *LvTimerT, repeatCount c.Int32T)

//go:linkname LvTimerSetAutoDelete C.lv_timer_set_auto_delete
func LvTimerSetAutoDelete(timer *LvTimerT, autoDelete c.Char)

//go:linkname LvTimerSetUserData C.lv_timer_set_user_data
func LvTimerSetUserData(timer *LvTimerT, userData unsafe.Pointer)

//go:linkname LvTimerReset C.lv_timer_reset
func LvTimerReset(timer *LvTimerT)

//go:linkname LvTimerEnable C.lv_timer_enable
func LvTimerEnable(en c.Char)

//go:linkname LvTimerGetIdle C.lv_timer_get_idle
func LvTimerGetIdle() c.Uint32T

//go:linkname LvTimerGetTimeUntilNext C.lv_timer_get_time_until_next
func LvTimerGetTimeUntilNext() c.Uint32T

//go:linkname LvTimerGetNext C.lv_timer_get_next
func LvTimerGetNext(timer *LvTimerT) *LvTimerT

//go:linkname LvTimerGetUserData C.lv_timer_get_user_data
func LvTimerGetUserData(timer *LvTimerT) unsafe.Pointer

//go:linkname LvTimerGetPaused C.lv_timer_get_paused
func LvTimerGetPaused(timer *LvTimerT) c.Char
