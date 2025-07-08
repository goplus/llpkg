package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type LdoRegulatorChannelT struct {
	Unused [8]uint8
}
type EspLdoChannelHandleT *LdoRegulatorChannelT

/**
 * @brief LDO channel configurations
 */

type EspLdoChannelConfigT struct {
	ChanId    c.Int
	VoltageMv c.Int
	Flags     LdoExtraFlags
}

type LdoExtraFlags struct {
	Unused [8]uint8
}

/**
 * @brief Acquire an LDO channel with the specified configuration
 *
 * @note This function can't automatically search a LDO channel for you, you must specify a LDO channel ID manually, based on your schematic.
 * @note The same channel can be acquired multiple times in different places of the application code, however,
 *       if the LDO channel is adjustable, you can't acquire it multiple times, in case user A changes the voltage and breaks the voltage setting of user B.
 * @note You should release the channel by `esp_ldo_release_channel` when it's no longer needed.
 *
 * @param[in] config The configuration of the LDO channel
 * @param[out] out_handle The returned LDO channel handle
 * @return
 *      - ESP_OK: Acquire the LDO channel successfully
 *      - ESP_ERR_INVALID_ARG: Acquire the LDO channel failed due to invalid arguments
 *      - ESP_FAIL: Acquire the LDO channel failed due to other reasons
 */
// llgo:link (*EspLdoChannelConfigT).EspLdoAcquireChannel C.esp_ldo_acquire_channel
func (recv_ *EspLdoChannelConfigT) EspLdoAcquireChannel(out_handle *EspLdoChannelHandleT) EspErrT {
	return 0
}

/**
 * @brief Release the LDO channel
 *
 * @param[in] chan The LDO channel handle returned from `esp_ldo_acquire_channel`
 * @return
 *      - ESP_OK: Release the LDO channel successfully
 *      - ESP_ERR_INVALID_ARG: Release the LDO channel failed due to invalid arguments
 *      - ESP_ERR_INVALID_STATE: Release the LDO channel failed due to invalid state, e.g., the channel handle is double released
 *      - ESP_FAIL: Release the LDO channel failed due to other reasons
 */
//go:linkname EspLdoReleaseChannel C.esp_ldo_release_channel
func EspLdoReleaseChannel(chan_ EspLdoChannelHandleT) EspErrT

/**
 * @brief Adjust the voltage of the LDO channel
 *
 * @param[in] chan The LDO channel handle returned from `esp_ldo_acquire_channel`
 * @param[in] voltage_mv The voltage value to be set to the LDO channel, in millivolts
 * @return
 *      - ESP_OK: Adjust the voltage of the LDO channel successfully
 *      - ESP_ERR_INVALID_ARG: Adjust the voltage of the LDO channel failed due to invalid arguments
 *      - ESP_ERR_NOT_SUPPORTED: Adjust the voltage of the LDO channel failed due to the channel is not adjustable
 *      - ESP_FAIL: Adjust the voltage of the LDO channel failed due to other reasons
 */
//go:linkname EspLdoChannelAdjustVoltage C.esp_ldo_channel_adjust_voltage
func EspLdoChannelAdjustVoltage(chan_ EspLdoChannelHandleT, voltage_mv c.Int) EspErrT

/**
 * @brief Dump LDO channel status to the specified stream
 *
 * @param[in] stream IO stream. Can be stdout, stderr, or a file/string stream.
 * @return
 *      - ESP_OK: Dump the LDO channel status successfully
 *      - ESP_FAIL: Dump the LDO channel status failed
 */
//go:linkname EspLdoDump C.esp_ldo_dump
func EspLdoDump(stream *c.FILE) EspErrT
