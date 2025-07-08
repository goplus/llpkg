package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Set RTC IO direction.
 *
 * Configure RTC IO direction, such as output only, input only,
 * output and input.
 *
 * @param rtcio_num The index of rtcio. 0 ~ SOC_RTCIO_PIN_COUNT.
 * @param mode IO direction.
 */
//go:linkname RtcioHalSetDirection C.rtcio_hal_set_direction
func RtcioHalSetDirection(rtcio_num c.Int, mode RtcGpioModeT)

/**
 * Set RTC IO direction in deep sleep or disable sleep status.
 *
 * NOTE: ESP32 supports INPUT_ONLY mode.
 *       The rest targets support INPUT_ONLY, OUTPUT_ONLY, INPUT_OUTPUT mode.
 *
 * @param rtcio_num The index of rtcio. 0 ~ SOC_RTCIO_PIN_COUNT.
 * @param mode IO direction.
 */
//go:linkname RtcioHalSetDirectionInSleep C.rtcio_hal_set_direction_in_sleep
func RtcioHalSetDirectionInSleep(rtcio_num c.Int, mode RtcGpioModeT)

/**
 * Helper function to disconnect internal circuits from an RTC IO
 * This function disables input, output, pullup, pulldown, and enables
 * hold feature for an RTC IO.
 * Use this function if an RTC IO needs to be disconnected from internal
 * circuits in deep sleep, to minimize leakage current.
 *
 * In particular, for ESP32-WROVER module, call
 * rtc_gpio_isolate(GPIO_NUM_12) before entering deep sleep, to reduce
 * deep sleep current.
 *
 * @param rtcio_num The index of rtcio. 0 ~ SOC_RTCIO_PIN_COUNT.
 */
//go:linkname RtcioHalIsolate C.rtcio_hal_isolate
func RtcioHalIsolate(rtc_num c.Int)
