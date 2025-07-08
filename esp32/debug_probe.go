package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type DebugProbeUnitT struct {
	Unused [8]uint8
}
type DebugProbeUnitHandleT *DebugProbeUnitT

type DebugProbeChannelT struct {
	Unused [8]uint8
}
type DebugProbeChannelHandleT *DebugProbeChannelT

/**
 * @brief Configuration for a debug probe unit
 */

type DebugProbeUnitConfigT struct {
	ProbeOutGpioNums [16]GpioNumT
}

/**
 * @brief Create a new debug probe unit
 *
 * @param[in] config Configuration for the debug probe unit
 * @param[out] out_handle Handle of the created debug probe unit
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if the configuration is invalid
 *      - ESP_ERR_NOT_FOUND if there is no free unit slot
 *      - ESP_ERR_NO_MEM if memory allocation failed
 *      - ESP_FAIL if an internal error occurred
 */
// llgo:link (*DebugProbeUnitConfigT).DebugProbeNewUnit C.debug_probe_new_unit
func (recv_ *DebugProbeUnitConfigT) DebugProbeNewUnit(out_handle *DebugProbeUnitHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete a debug probe unit
 *
 * @param[in] unit Handle of the debug probe unit
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if the handle is invalid
 *      - ESP_FAIL if an internal error occurred
 */
//go:linkname DebugProbeDelUnit C.debug_probe_del_unit
func DebugProbeDelUnit(unit DebugProbeUnitHandleT) EspErrT

/**
 * @brief Configuration for a debug probe channel
 */

type DebugProbeChannelConfigT struct {
	TargetModule DebugProbeTargetT
}

/**
 * @brief Create a new debug probe channel in a unit
 *
 * @param[in] unit Handle of the debug probe unit
 * @param[in] config Configuration for the debug probe channel
 * @param[out] out_handle Handle of the created debug probe channel
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if the handle or configuration is invalid
 *      - ESP_ERR_NOT_FOUND if there is no free channel slot
 *      - ESP_ERR_NO_MEM if memory allocation failed
 *      - ESP_FAIL if an internal error occurred
 */
//go:linkname DebugProbeNewChannel C.debug_probe_new_channel
func DebugProbeNewChannel(unit DebugProbeUnitHandleT, config *DebugProbeChannelConfigT, out_handle *DebugProbeChannelHandleT) EspErrT

/**
 * @brief Delete a debug probe channel
 *
 * @param[in] chan Handle of the debug probe channel
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if the handle is invalid
 *      - ESP_FAIL if an internal error occurred
 */
//go:linkname DebugProbeDelChannel C.debug_probe_del_channel
func DebugProbeDelChannel(chan_ DebugProbeChannelHandleT) EspErrT

/**
 * @brief Add signals to a debug probe channel
 *
 * @note The n-th byte of channel output comes from the n-th byte in the sig_group.
 * @note The signals to be added must aligned to the byte boundary.
 *       byte_idx = 0: signal 0-7 in the group
 *       byte_idx = 1: signal 8-15 in the group
 *       ...
 * @note If you add the signals from different groups but with the same byte_idx, only the last added signal will be effective.
 * @note You can save up to 32 signals in a channel, but in the end, only the part of them (e.g. upper or lower 16 signals) can be output to the GPIO pads.
 *
 * @param[in] chan Handle of the debug probe channel
 * @param[in] byte_idx Byte index of the signals, ranges from 0 to 3
 * @param[in] sig_group Signal group of the signal, ranges from 0 to 15
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if the parameters are invalid
 *      - ESP_FAIL if an internal error occurred
 */
//go:linkname DebugProbeChanAddSignalByByte C.debug_probe_chan_add_signal_by_byte
func DebugProbeChanAddSignalByByte(chan_ DebugProbeChannelHandleT, byte_idx c.Uint8T, sig_group c.Uint8T) EspErrT

/**
 * @brief Merge the part of the channel output to the debug probe unit output
 *
 *           +----upper16---+
 * chan_0 ---+              |
 *           +----lower16---+
 *                          |
 *                          +-------unit_output[31:0]
 *                          |
 *           +----upper16---+
 * chan_1 ---+              |
 *           +----lower16---+
 *
 * @param[in] unit Handle of the debug probe unit
 * @param[in] chan0 Handle of the debug probe channel 0, whose output will be merged to the lower 16 signals of the unit output
 * @param[in] split_of_chan0 Part of the channel 0 output to be merged
 * @param[in] chan1 Handle of the debug probe channel 1, whose output will be merged to the upper 16 signals of the unit output
 * @param[in] split_of_chan1 Part of the channel 1 output to be merged
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if the parameters are invalid
 *      - ESP_FAIL if an internal error occurred
 */
//go:linkname DebugProbeUnitMerge16 C.debug_probe_unit_merge16
func DebugProbeUnitMerge16(unit DebugProbeUnitHandleT, chan0 DebugProbeChannelHandleT, split_of_chan0 DebugProbeSplitU16T, chan1 DebugProbeChannelHandleT, split_of_chan1 DebugProbeSplitU16T) EspErrT

/**
 * @brief Read the value of the debug probe unit
 *
 * @note Only the lower 16 signals of the probe unit can be routed to the GPIO pads.
 *
 * @param[in] unit Handle of the debug probe unit
 * @param[out] value Current value of the debug probe unit output
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if the handle or value is invalid
 *      - ESP_FAIL if an internal error occurred
 */
//go:linkname DebugProbeUnitRead C.debug_probe_unit_read
func DebugProbeUnitRead(unit DebugProbeUnitHandleT, value *c.Uint32T) EspErrT
