package freertos

import _ "unsafe"

/*---------------------------------------------------------------
            ADC Hardware Calibration
---------------------------------------------------------------*/
/**
 * @brief Calculate the ADC HW calibration code. (Based on the pre-stored efuse or actual calibration)
 *
 * @param adc_n ADC unit to calibrate
 * @param atten Attenuation to use
 */
// llgo:link AdcUnitT.AdcCalcHwCalibrationCode C.adc_calc_hw_calibration_code
func (recv_ AdcUnitT) AdcCalcHwCalibrationCode(atten AdcAttenT) {
}

/**
 * @brief Set the ADC HW calibration code.
 *
 * @param adc_n ADC unit to calibrate
 * @param atten Attenuation to use
 */
// llgo:link AdcUnitT.AdcSetHwCalibrationCode C.adc_set_hw_calibration_code
func (recv_ AdcUnitT) AdcSetHwCalibrationCode(atten AdcAttenT) {
}

/*---------------------------------------------------------------
            ADC Cross Peripheral Locks
---------------------------------------------------------------*/
/**
 * @brief Acquire ADC lock by unit
 *
 * The lock acquiring sequence will be: ADC1, ADC2, ...
 *
 * @note If any of the locks are taken, this API will wait until the lock is successfully acquired.
 *
 * @param[in] adc_unit    ADC unit ID
 *
 * @return
 *        - ESP_OK: On success
 */
// llgo:link AdcUnitT.AdcLockAcquire C.adc_lock_acquire
func (recv_ AdcUnitT) AdcLockAcquire() EspErrT {
	return 0
}

/**
 * @brief Release ADC lock by unit
 *
 * The lock releasing sequence will be: ..., ADC2, ADC1
 *
 * @param[in] adc_unit    ADC unit ID
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_STATE: The lock(s) isn't acquired yet
 */
// llgo:link AdcUnitT.AdcLockRelease C.adc_lock_release
func (recv_ AdcUnitT) AdcLockRelease() EspErrT {
	return 0
}

/**
 * @brief Try to acquire ADC lock by unit
 *
 * The lock acquiring sequence will be: ADC1, ADC2, ...
 *
 * @note If any of the locks are taken, this API will return immediately with an error `ESP_ERR_TIMEOUT`
 *
 * @param[in] adc_unit    ADC unit ID
 *
 * @return
 *        - ESP_OK:          On success
 *        - ESP_ERR_TIMEOUT: Lock(s) is taken already
 */
// llgo:link AdcUnitT.AdcLockTryAcquire C.adc_lock_try_acquire
func (recv_ AdcUnitT) AdcLockTryAcquire() EspErrT {
	return 0
}

/**
 * @brief For WIFI module to claim the usage of ADC2.
 *
 * Other tasks will be forbidden to use ADC2 between ``adc2_wifi_acquire`` and ``adc2_wifi_release``.
 * The WIFI module may have to wait for a short time for the current conversion (if exist) to finish.
 *
 * @return
 *        - ESP_OK success
 *        - ESP_ERR_TIMEOUT reserved for future use. Currently the function will wait until success.
 */
//go:linkname Adc2WifiAcquire C.adc2_wifi_acquire
func Adc2WifiAcquire() EspErrT

/**
 * @brief For WIFI module to let other tasks use the ADC2 when WIFI is not work.
 *
 * Other tasks will be forbidden to use ADC2 between ``adc2_wifi_acquire`` and ``adc2_wifi_release``.
 * Call this function to release the occupation of ADC2 by WIFI.
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_STATE: The lock(s) isn't acquired yet
 */
//go:linkname Adc2WifiRelease C.adc2_wifi_release
func Adc2WifiRelease() EspErrT

/*------------------------------------------------------------------------------
* For those who use APB_SARADC periph
*----------------------------------------------------------------------------*/
/**
 * @brief Claim the usage of the APB_SARADC periph
 *
 * Reference count inside
 */
//go:linkname AdcApbPeriphClaim C.adc_apb_periph_claim
func AdcApbPeriphClaim()

/**
 * @brief Free the usage of the APB_SARADC periph
 *
 * Reference count inside
 */
//go:linkname AdcApbPeriphFree C.adc_apb_periph_free
func AdcApbPeriphFree()
