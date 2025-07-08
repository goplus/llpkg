package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Initialise SAR related peripheral register settings
 * Should only be used when running into app stage
 */
//go:linkname SarPeriphCtrlInit C.sar_periph_ctrl_init
func SarPeriphCtrlInit()

/*------------------------------------------------------------------------------
* ADC Power
*----------------------------------------------------------------------------*/
/**
 * @brief Acquire the ADC oneshot mode power
 */
//go:linkname SarPeriphCtrlAdcOneshotPowerAcquire C.sar_periph_ctrl_adc_oneshot_power_acquire
func SarPeriphCtrlAdcOneshotPowerAcquire()

/**
 * @brief Release the ADC oneshot mode power
 */
//go:linkname SarPeriphCtrlAdcOneshotPowerRelease C.sar_periph_ctrl_adc_oneshot_power_release
func SarPeriphCtrlAdcOneshotPowerRelease()

/**
 * @brief Acquire the ADC continuous mode power
 */
//go:linkname SarPeriphCtrlAdcContinuousPowerAcquire C.sar_periph_ctrl_adc_continuous_power_acquire
func SarPeriphCtrlAdcContinuousPowerAcquire()

/**
 * @brief Release the ADC ADC continuous mode power
 */
//go:linkname SarPeriphCtrlAdcContinuousPowerRelease C.sar_periph_ctrl_adc_continuous_power_release
func SarPeriphCtrlAdcContinuousPowerRelease()

/*------------------------------------------------------------------------------
* PWDET Power
*----------------------------------------------------------------------------*/
/**
 * @brief Acquire the PWDET Power
 */
//go:linkname SarPeriphCtrlPwdetPowerAcquire C.sar_periph_ctrl_pwdet_power_acquire
func SarPeriphCtrlPwdetPowerAcquire()

/**
 * @brief Release the PWDET Power
 */
//go:linkname SarPeriphCtrlPwdetPowerRelease C.sar_periph_ctrl_pwdet_power_release
func SarPeriphCtrlPwdetPowerRelease()

/**
 * @brief Acquire the temperature sensor power
 */
//go:linkname TemperatureSensorPowerAcquire C.temperature_sensor_power_acquire
func TemperatureSensorPowerAcquire()

/**
 * @brief Release the temperature sensor power
 */
//go:linkname TemperatureSensorPowerRelease C.temperature_sensor_power_release
func TemperatureSensorPowerRelease()

/**
 * @brief Get the temperature value and choose the temperature sensor range. Will be both used in phy and peripheral.
 *
 * @param range_changed Pointer to whether range has been changed here. If you don't need this param, you can
 *        set NULL directly.
 *
 * @return temperature sensor value.
 */
//go:linkname TempSensorGetRawValue C.temp_sensor_get_raw_value
func TempSensorGetRawValue(range_changed *bool) c.Int16T

/**
 * @brief Synchronize the tsens_idx between sar_periph and driver
 *
 * @param tsens_idx index value of temperature sensor attribute
 */
//go:linkname TempSensorSyncTsensIdx C.temp_sensor_sync_tsens_idx
func TempSensorSyncTsensIdx(tsens_idx c.Int)

/**
 * @brief Enable SAR power when system wakes up
 */
//go:linkname SarPeriphCtrlPowerEnable C.sar_periph_ctrl_power_enable
func SarPeriphCtrlPowerEnable()

/**
 * @brief Disable SAR power when system goes to sleep
 */
//go:linkname SarPeriphCtrlPowerDisable C.sar_periph_ctrl_power_disable
func SarPeriphCtrlPowerDisable()
