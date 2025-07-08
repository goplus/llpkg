package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type DsiHostDevT struct {
	Unused [8]uint8
}
type MipiDsiHostSocHandleT *DsiHostDevT

type DsiBrgDevT struct {
	Unused [8]uint8
}
type MipiDsiBridgeSocHandleT *DsiBrgDevT

/**
 * @brief MIPI DSI HAL driver context
 */

type MipiDsiHalContextT struct {
	Host            MipiDsiHostSocHandleT
	Bridge          MipiDsiBridgeSocHandleT
	LaneBitRateMbps c.Uint32T
	DpiClockFreqMhz c.Uint32T
}

/**
 * @brief MIPI DSI HAL driver configuration
 */

type MipiDsiHalConfigT struct {
	BusId           c.Int
	LaneBitRateMbps c.Uint32T
	NumDataLanes    c.Uint8T
}

/**
 * @brief Initialize MIPI DSI Hal driver context
 *
 * @note Caller should malloc the memory for the hal context
 *
 * @param hal Pointer to the HAL driver context
 * @param config Pointer to the HAL driver configuration
 */
// llgo:link (*MipiDsiHalContextT).MipiDsiHalInit C.mipi_dsi_hal_init
func (recv_ *MipiDsiHalContextT) MipiDsiHalInit(config *MipiDsiHalConfigT) {
}

/**
 * @brief Deinitialize MIPI DSI Hal driver context
 *
 * @param hal Pointer to the HAL driver context
 */
// llgo:link (*MipiDsiHalContextT).MipiDsiHalDeinit C.mipi_dsi_hal_deinit
func (recv_ *MipiDsiHalContextT) MipiDsiHalDeinit() {
}

/**
 * @brief Configure the PHY PLL
 *
 * @param hal Pointer to the HAL driver context
 * @param phy_clk_src_freq_hz PHY clock source frequency in Hz
 * @param lane_bit_rate_mbps Lane bit rate in Mbps
 */
// llgo:link (*MipiDsiHalContextT).MipiDsiHalConfigurePhyPll C.mipi_dsi_hal_configure_phy_pll
func (recv_ *MipiDsiHalContextT) MipiDsiHalConfigurePhyPll(phy_clk_src_freq_hz c.Uint32T, lane_bit_rate_mbps c.Uint32T) {
}

/**
 * @brief Write a value to a PHY register via internal bus (so-called test interface)
 *
 * @param hal Pointer to the HAL driver context
 * @param reg_addr Address of the PHY register
 * @param reg_val Value to be written to the PHY register
 */
// llgo:link (*MipiDsiHalContextT).MipiDsiHalPhyWriteRegister C.mipi_dsi_hal_phy_write_register
func (recv_ *MipiDsiHalContextT) MipiDsiHalPhyWriteRegister(reg_addr c.Uint8T, reg_val c.Uint8T) {
}

/**
 * @brief Send a DCS command with parameters via the generic interface
 *
 * @note The command itself and the parameters are sent in one packet. For simplicity, we use the same
 *       DCS_LONG_WRITE type to send all different commands.
 *
 * @param hal Pointer to the HAL driver context
 * @param vc Virtual channel number
 * @param command DCS command
 * @param command_bytes Number of bytes of the command
 * @param param Pointer to the parameters
 * @param param_size Number of bytes of the parameters
 */
// llgo:link (*MipiDsiHalContextT).MipiDsiHalHostGenWriteDcsCommand C.mipi_dsi_hal_host_gen_write_dcs_command
func (recv_ *MipiDsiHalContextT) MipiDsiHalHostGenWriteDcsCommand(vc c.Uint8T, command c.Uint32T, command_bytes c.Uint32T, param c.Pointer, param_size c.Uint16T) {
}

/**
 * @brief Send a DCS command and return the associated parameters via the generic interface
 *
 * @param hal Pointer to the HAL driver context
 * @param vc Virtual channel number
 * @param command DCS command
 * @param command_bytes Number of bytes of the command
 * @param ret_param Pointer to the buffer to store the returned parameters
 * @param param_buf_size Size of the buffer to store the returned parameters
 */
// llgo:link (*MipiDsiHalContextT).MipiDsiHalHostGenReadDcsCommand C.mipi_dsi_hal_host_gen_read_dcs_command
func (recv_ *MipiDsiHalContextT) MipiDsiHalHostGenReadDcsCommand(vc c.Uint8T, command c.Uint32T, command_bytes c.Uint32T, ret_param c.Pointer, param_buf_size c.Uint16T) {
}

/**
 * @brief Send a long packet via the generic interface
 *
 * @param hal Pointer to the HAL driver context
 * @param vc Virtual channel number
 * @param dt Data type
 * @param buffer Pointer to the buffer
 * @param buffer_size Number of bytes to be sent
 */
// llgo:link (*MipiDsiHalContextT).MipiDsiHalHostGenWriteLongPacket C.mipi_dsi_hal_host_gen_write_long_packet
func (recv_ *MipiDsiHalContextT) MipiDsiHalHostGenWriteLongPacket(vc c.Uint8T, dt MipiDsiDataTypeT, buffer c.Pointer, buffer_size c.Uint16T) {
}

/**
 * @brief Send a short packet via the generic interface
 *
 * @param hal Pointer to the HAL driver context
 * @param vc Virtual channel number
 * @param dt Data type
 * @param header_data Data to be sent, filled into the DSI packet header
 */
// llgo:link (*MipiDsiHalContextT).MipiDsiHalHostGenWriteShortPacket C.mipi_dsi_hal_host_gen_write_short_packet
func (recv_ *MipiDsiHalContextT) MipiDsiHalHostGenWriteShortPacket(vc c.Uint8T, dt MipiDsiDataTypeT, header_data c.Uint16T) {
}

/**
 * @brief Send a short packet via the generic interface and return the associated data
 *
 * @param hal Pointer to the HAL driver context
 * @param vc Virtual channel number
 * @param dt Data type
 * @param header_data Data to be filled into the DSI packet header
 * @param ret_buffer Pointer to the buffer to store the returned data
 * @param buffer_size Size of the buffer to store the returned data
 */
// llgo:link (*MipiDsiHalContextT).MipiDsiHalHostGenReadShortPacket C.mipi_dsi_hal_host_gen_read_short_packet
func (recv_ *MipiDsiHalContextT) MipiDsiHalHostGenReadShortPacket(vc c.Uint8T, dt MipiDsiDataTypeT, header_data c.Uint16T, ret_buffer c.Pointer, buffer_size c.Uint16T) {
}

/**
 * @brief Set DPI color coding
 *
 * @param hal Pointer to the HAL driver context
 * @param color_coding Color coding
 * @param sub_config Sub configuration
 */
// llgo:link (*MipiDsiHalContextT).MipiDsiHalHostDpiSetColorCoding C.mipi_dsi_hal_host_dpi_set_color_coding
func (recv_ *MipiDsiHalContextT) MipiDsiHalHostDpiSetColorCoding(color_coding LcdColorFormatT, sub_config c.Uint32T) {
}

/**
 * @brief Set horizontal timing parameters for DPI
 *
 * @param hal Pointer to the HAL driver context
 * @param hsw Horizontal Synchronization Width
 * @param hbp Horizontal Back Porch
 * @param active_width Active Width
 * @param hfp Horizontal Front Porch
 */
// llgo:link (*MipiDsiHalContextT).MipiDsiHalHostDpiSetHorizontalTiming C.mipi_dsi_hal_host_dpi_set_horizontal_timing
func (recv_ *MipiDsiHalContextT) MipiDsiHalHostDpiSetHorizontalTiming(hsw c.Uint32T, hbp c.Uint32T, active_width c.Uint32T, hfp c.Uint32T) {
}

/**
 * @brief Set vertical timing parameters for DPI
 *
 * @param hal Pointer to the HAL driver context
 * @param vsw Vertical Synchronization Width
 * @param vbp Vertical Back Porch
 * @param active_height Active Height
 * @param vfp Vertical Front Porch
 */
// llgo:link (*MipiDsiHalContextT).MipiDsiHalHostDpiSetVerticalTiming C.mipi_dsi_hal_host_dpi_set_vertical_timing
func (recv_ *MipiDsiHalContextT) MipiDsiHalHostDpiSetVerticalTiming(vsw c.Uint32T, vbp c.Uint32T, active_height c.Uint32T, vfp c.Uint32T) {
}

/**
 * @brief Calculate the divider for DPI clock
 *
 * @note This function will also update the real DPI clock frequency in the HAL context
 *
 * @param hal Pointer to the HAL driver context
 * @param clk_src_mhz Clock source frequency in MHz
 * @param expect_dpi_clk_mhz Expected DPI clock frequency in MHz
 * @return Divider value
 */
// llgo:link (*MipiDsiHalContextT).MipiDsiHalHostDpiCalculateDivider C.mipi_dsi_hal_host_dpi_calculate_divider
func (recv_ *MipiDsiHalContextT) MipiDsiHalHostDpiCalculateDivider(clk_src_mhz c.Uint32T, expect_dpi_clk_mhz c.Uint32T) c.Uint32T {
	return 0
}
