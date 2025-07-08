package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Set parameter of touch sensor filter and detection algorithm.
 * For more details on the detection algorithm, please refer to the application documentation.
 *
 * @param filter_info select filter type and threshold of detection algorithm
 */
// llgo:link (*TouchFilterConfigT).TouchHalFilterSetConfig C.touch_hal_filter_set_config
func (recv_ *TouchFilterConfigT) TouchHalFilterSetConfig() {
}

/**
 * Get parameter of touch sensor filter and detection algorithm.
 * For more details on the detection algorithm, please refer to the application documentation.
 *
 * @param filter_info select filter type and threshold of detection algorithm
 */
// llgo:link (*TouchFilterConfigT).TouchHalFilterGetConfig C.touch_hal_filter_get_config
func (recv_ *TouchFilterConfigT) TouchHalFilterGetConfig() {
}

/**
 * set parameter of denoise pad (TOUCH_PAD_NUM0).
 *        T0 is an internal channel that does not have a corresponding external GPIO.
 *        T0 will work simultaneously with the measured channel Tn. Finally, the actual
 *        measured value of Tn is the value after subtracting lower bits of T0.
 *        This denoise function filters out interference introduced on all channels,
 *        such as noise introduced by the power supply and external EMI.
 *
 * @param denoise parameter of denoise
 */
// llgo:link (*TouchPadDenoiseT).TouchHalDenoiseSetConfig C.touch_hal_denoise_set_config
func (recv_ *TouchPadDenoiseT) TouchHalDenoiseSetConfig() {
}

/**
 * @brief get parameter of denoise pad (TOUCH_PAD_NUM0).
 *
 * @param denoise Pointer to parameter of denoise
 */
// llgo:link (*TouchPadDenoiseT).TouchHalDenoiseGetConfig C.touch_hal_denoise_get_config
func (recv_ *TouchPadDenoiseT) TouchHalDenoiseGetConfig() {
}

/**
 * Enable denoise function.
 * T0 is an internal channel that does not have a corresponding external GPIO.
 * T0 will work simultaneously with the measured channel Tn. Finally, the actual
 * measured value of Tn is the value after subtracting lower bits of T0.
 * This denoise function filters out interference introduced on all channels,
 * such as noise introduced by the power supply and external EMI.
 */
//go:linkname TouchHalDenoiseEnable C.touch_hal_denoise_enable
func TouchHalDenoiseEnable()

/**
 * Set parameter of waterproof function.
 *
 * The waterproof function includes a shielded channel (TOUCH_PAD_NUM14) and a guard channel.
 * Guard pad is used to detect the large area of water covering the touch panel.
 * Shield pad is used to shield the influence of water droplets covering the touch panel.
 * It is generally designed as a grid and is placed around the touch buttons.
 *
 * @param waterproof parameter of waterproof
 */
// llgo:link (*TouchPadWaterproofT).TouchHalWaterproofSetConfig C.touch_hal_waterproof_set_config
func (recv_ *TouchPadWaterproofT) TouchHalWaterproofSetConfig() {
}

/**
 * Get parameter of waterproof function.
 *
 * @param waterproof parameter of waterproof.
 */
// llgo:link (*TouchPadWaterproofT).TouchHalWaterproofGetConfig C.touch_hal_waterproof_get_config
func (recv_ *TouchPadWaterproofT) TouchHalWaterproofGetConfig() {
}

/**
 * Enable parameter of waterproof function.
 * Should be called after function ``touch_hal_waterproof_set_config``.
 */
//go:linkname TouchHalWaterproofEnable C.touch_hal_waterproof_enable
func TouchHalWaterproofEnable()

/**
 * Enable/disable proximity function of touch channels.
 * The proximity sensor measurement is the accumulation of touch channel measurements.
 *
 * @note Supports up to three touch channels configured as proximity sensors.
 * @param touch_num touch pad index
 * @param enabled true: enable the proximity function; false:  disable the proximity function
 * @return
 *     - true: Configured correctly.
 *     - false: Configured error.
 */
// llgo:link TouchPadT.TouchHalEnableProximity C.touch_hal_enable_proximity
func (recv_ TouchPadT) TouchHalEnableProximity(enabled bool) bool {
	return false
}

/**
 * Get parameter of touch sensor sleep channel.
 * The touch sensor can works in sleep mode to wake up sleep.
 * After the sleep channel is configured, users should query the channel reading using a specific function.
 *
 * @param slp_config Point to touch sleep pad config.
 */
// llgo:link (*TouchPadSleepChannelT).TouchHalSleepChannelGetConfig C.touch_hal_sleep_channel_get_config
func (recv_ *TouchPadSleepChannelT) TouchHalSleepChannelGetConfig() {
}

/**
 * Set parameter of touch sensor sleep channel.
 * The touch sensor can works in sleep mode to wake up sleep.
 * After the sleep channel is configured, users should query the channel reading using a specific function.
 *
 * @note ESP32S2 only support one channel to be set sleep channel.
 *
 * @param pad_num touch sleep pad number.
 * @param enable Enable/disable sleep pad function.
 */
// llgo:link TouchPadT.TouchHalSleepChannelEnable C.touch_hal_sleep_channel_enable
func (recv_ TouchPadT) TouchHalSleepChannelEnable(enable bool) {
}

/**
 * Change the operating frequency of touch pad in deep sleep state. Reducing the operating frequency can effectively reduce power consumption.
 * If this function is not called, the working frequency of touch in the deep sleep state is the same as that in the wake-up state.
 *
 * @param sleep_cycle The touch sensor will sleep after each measurement.
 *                    sleep_cycle decide the interval between each measurement.
 *                    t_sleep = sleep_cycle / (RTC_SLOW_CLK frequency).
 *                    The approximate frequency value of RTC_SLOW_CLK can be obtained using rtc_clk_slow_freq_get_hz function.
 * @param meas_times The times of charge and discharge in each measure process of touch channels.
 *                  The timer frequency is 8Mhz. Range: 0 ~ 0xffff.
 *                  Recommended typical value: Modify this value to make the measurement time around 1ms.
 */
//go:linkname TouchHalSleepChannelSetWorkTime C.touch_hal_sleep_channel_set_work_time
func TouchHalSleepChannelSetWorkTime(sleep_cycle c.Uint16T, meas_times c.Uint16T)

/**
 * Get the operating frequency of touch pad in deep sleep state. Reducing the operating frequency can effectively reduce power consumption.
 *
 * @param sleep_cycle The touch sensor will sleep after each measurement.
 *                    sleep_cycle decide the interval between each measurement.
 *                    t_sleep = sleep_cycle / (RTC_SLOW_CLK frequency).
 *                    The approximate frequency value of RTC_SLOW_CLK can be obtained using rtc_clk_slow_freq_get_hz function.
 * @param meas_times The times of charge and discharge in each measure process of touch channels.
 *                  The timer frequency is 8Mhz. Range: 0 ~ 0xffff.
 *                  Recommended typical value: Modify this value to make the measurement time around 1ms.
 */
//go:linkname TouchHalSleepChannelGetWorkTime C.touch_hal_sleep_channel_get_work_time
func TouchHalSleepChannelGetWorkTime(sleep_cycle *c.Uint16T, meas_times *c.Uint16T)
