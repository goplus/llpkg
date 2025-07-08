package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Input parameters to the ``spi_hal_cal_clock_conf`` to calculate the timing configuration
 */

type SpiHalTimingParamT struct {
	ClkSrcHz     c.Uint32T
	HalfDuplex   c.Uint32T
	NoCompensate c.Uint32T
	ExpectedFreq c.Uint32T
	DutyCycle    c.Uint32T
	InputDelayNs c.Uint32T
	UseGpio      bool
}

/**
 * Timing configuration structure that should be calculated by
 * ``spi_hal_cal_clock_conf`` at initialization and hold. Filled into the
 * ``timing_conf`` member of the context of HAL before setup a device.
 */

type SpiHalTimingConfT struct {
	ClockReg        SpiLlClockValT
	ClockSource     SpiClockSourceT
	SourcePreDiv    c.Uint32T
	RealFreq        c.Int
	TimingDummy     c.Int
	TimingMisoDelay c.Int
	RxSamplePoint   SpiSamplingPointT
}

/**
 * Transaction configuration structure, this should be assigned by driver each time.
 * All these parameters will be updated to the peripheral every transaction.
 */

type SpiHalTransConfigT struct {
	Cmd          c.Uint16T
	CmdBits      c.Int
	AddrBits     c.Int
	DummyBits    c.Int
	TxBitlen     c.Int
	RxBitlen     c.Int
	Addr         c.Uint64T
	SendBuffer   *c.Uint8T
	RcvBuffer    *c.Uint8T
	LineMode     SpiLineModeT
	CsKeepActive c.Int
}

/**
 * Context that should be maintained by both the driver and the HAL.
 */

type SpiHalContextT struct {
	Hw          *SpiDevT
	DmaEnabled  bool
	TransConfig SpiHalTransConfigT
}

/**
 * Device configuration structure, this should be initialised by driver based on different devices respectively.
 * All these parameters will be updated to the peripheral only when ``spi_hal_setup_device``.
 * They may not get updated when ``spi_hal_setup_trans``.
 */

type SpiHalDevConfigT struct {
	Mode       c.Int
	CsSetup    c.Int
	CsHold     c.Int
	CsPinId    c.Int
	TimingConf SpiHalTimingConfT
}

/**
 * SCT mode required configurations, per segment
 */

type SpiHalSegConfigT struct {
	SegEnd    bool
	SegGapLen c.Uint32T
	CsSetup   c.Int
	Cmd       c.Uint16T
	CmdBits   c.Int
	Addr      c.Uint64T
	AddrBits  c.Int
	DummyBits c.Int
	TxBitlen  c.Int
	RxBitlen  c.Int
	CsHold    c.Int
}

/**
 * Init the peripheral and the context.
 *
 * @param hal        Context of the HAL layer.
 * @param host_id    Index of the SPI peripheral. 0 for SPI1, 1 for SPI2 and 2 for SPI3.
 */
// llgo:link (*SpiHalContextT).SpiHalInit C.spi_hal_init
func (recv_ *SpiHalContextT) SpiHalInit(host_id c.Uint32T) {
}

/**
 * Config default output IO level when don't have transaction
 *
 * @param hal Context of the HAL layer.
 * @param level IO level to config
 */
// llgo:link (*SpiHalContextT).SpiHalConfigIoDefaultLevel C.spi_hal_config_io_default_level
func (recv_ *SpiHalContextT) SpiHalConfigIoDefaultLevel(level bool) {
}

/**
 * Deinit the peripheral (and the context if needed).
 *
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SpiHalContextT).SpiHalDeinit C.spi_hal_deinit
func (recv_ *SpiHalContextT) SpiHalDeinit() {
}

/**
 * Setup device-related configurations according to the settings in the context.
 *
 * @param hal       Context of the HAL layer.
 * @param hal_dev   Device configuration
 */
// llgo:link (*SpiHalContextT).SpiHalSetupDevice C.spi_hal_setup_device
func (recv_ *SpiHalContextT) SpiHalSetupDevice(hal_dev *SpiHalDevConfigT) {
}

/**
 * Setup transaction related configurations according to the settings in the context.
 *
 * @param hal       Context of the HAL layer.
 * @param hal_dev   Device configuration
 * @param hal_trans Transaction configuration
 */
// llgo:link (*SpiHalContextT).SpiHalSetupTrans C.spi_hal_setup_trans
func (recv_ *SpiHalContextT) SpiHalSetupTrans(hal_dev *SpiHalDevConfigT, hal_trans *SpiHalTransConfigT) {
}

/**
 * Enable/Disable miso/mosi signals on peripheral side
 *
 * @param hw        Beginning address of the peripheral registers.
 * @param mosi_ena  enable/disable mosi line
 * @param miso_ena  enable/disable miso line
 */
// llgo:link (*SpiDevT).SpiHalEnableDataLine C.spi_hal_enable_data_line
func (recv_ *SpiDevT) SpiHalEnableDataLine(mosi_ena bool, miso_ena bool) {
}

/**
 * Prepare tx hardware for a new DMA trans
 *
 * @param hw Beginning address of the peripheral registers.
 */
// llgo:link (*SpiDevT).SpiHalHwPrepareRx C.spi_hal_hw_prepare_rx
func (recv_ *SpiDevT) SpiHalHwPrepareRx() {
}

/**
 * Prepare tx hardware for a new DMA trans
 *
 * @param hw Beginning address of the peripheral registers.
 */
// llgo:link (*SpiDevT).SpiHalHwPrepareTx C.spi_hal_hw_prepare_tx
func (recv_ *SpiDevT) SpiHalHwPrepareTx() {
}

/**
 * Trigger start a user-defined transaction.
 *
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SpiHalContextT).SpiHalUserStart C.spi_hal_user_start
func (recv_ *SpiHalContextT) SpiHalUserStart() {
}

/**
 * Check whether the transaction is done (trans_done is set).
 *
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SpiHalContextT).SpiHalUsrIsDone C.spi_hal_usr_is_done
func (recv_ *SpiHalContextT) SpiHalUsrIsDone() bool {
	return false
}

/**
 * Setup transaction operations, write tx buffer to HW registers
 *
 * @param hal       Context of the HAL layer.
 * @param hal_trans Transaction configuration.
 */
// llgo:link (*SpiHalContextT).SpiHalPushTxBuffer C.spi_hal_push_tx_buffer
func (recv_ *SpiHalContextT) SpiHalPushTxBuffer(hal_trans *SpiHalTransConfigT) {
}

/**
 * Post transaction operations, mainly fetch data from the buffer.
 *
 * @param hal       Context of the HAL layer.
 */
// llgo:link (*SpiHalContextT).SpiHalFetchResult C.spi_hal_fetch_result
func (recv_ *SpiHalContextT) SpiHalFetchResult() {
}

/*----------------------------------------------------------
 * Utils
 * ---------------------------------------------------------*/
/**
 * Calculate the configuration of clock and timing. The configuration will be used when ``spi_hal_setup_device``.
 *
 * It is highly suggested to do this at initialization, since it takes long time.
 *
 * @param timing_param   Input parameters to calculate timing configuration
 * @param timing_conf    Output of the timing configuration.
 *
 * @return ESP_OK if desired is available, otherwise fail.
 */
// llgo:link (*SpiHalTimingParamT).SpiHalCalClockConf C.spi_hal_cal_clock_conf
func (recv_ *SpiHalTimingParamT) SpiHalCalClockConf(timing_conf *SpiHalTimingConfT) EspErrT {
	return 0
}

/**
 * Get the frequency actual used.
 *
 * @param hal            Context of the HAL layer.
 * @param fapb           APB clock frequency.
 * @param hz             Desired frequencyc.
 * @param duty_cycle     Desired duty cycle.
 */
//go:linkname SpiHalMasterCalClock C.spi_hal_master_cal_clock
func SpiHalMasterCalClock(fapb c.Int, hz c.Int, duty_cycle c.Int) c.Int

/**
 * Get the timing configuration for given parameters.
 *
 * @param source_freq_hz Clock freq of selected clock source for SPI in Hz.
 * @param eff_clk        Actual SPI clock frequency
 * @param gpio_is_used   true if the GPIO matrix is used, otherwise false.
 * @param input_delay_ns Maximum delay between SPI launch clock and the data to
 *                       be valid. This is used to compensate/calculate the maximum frequency
 *                       allowed. Left 0 if not known.
 * @param dummy_n        Dummy cycles required to correctly read the data.
 * @param miso_delay_n   suggested delay on the MISO line, in APB clocks.
 */
//go:linkname SpiHalCalTiming C.spi_hal_cal_timing
func SpiHalCalTiming(source_freq_hz c.Int, eff_clk c.Int, gpio_is_used bool, input_delay_ns c.Int, dummy_n *c.Int, miso_delay_n *c.Int)

/**
 * Get the maximum frequency allowed to read if no compensation is used.
 *
 * @param gpio_is_used   true if the GPIO matrix is used, otherwise false.
 * @param input_delay_ns Maximum delay between SPI launch clock and the data to
 *                       be valid. This is used to compensate/calculate the maximum frequency
 *                       allowed. Left 0 if not known.
 */
//go:linkname SpiHalGetFreqLimit C.spi_hal_get_freq_limit
func SpiHalGetFreqLimit(gpio_is_used bool, input_delay_ns c.Int) c.Int

/*----------------------------------------------------------
 * Segmented-Configure-Transfer (SCT) Mode
 * ---------------------------------------------------------*/
/**
 * Initialise SCT mode required registers and hal states
 *
 * @param hal            Context of the HAL layer.
 */
// llgo:link (*SpiHalContextT).SpiHalSctInit C.spi_hal_sct_init
func (recv_ *SpiHalContextT) SpiHalSctInit() {
}

/**
 * Initialise conf buffer, give it an initial value
 *
 * @param hal            Context of the HAL layer.
 */
// llgo:link (*SpiHalContextT).SpiHalSctInitConfBuffer C.spi_hal_sct_init_conf_buffer
func (recv_ *SpiHalContextT) SpiHalSctInitConfBuffer(conf_buffer *c.Uint32T) {
}

/**
 * Format the conf buffer
 * According to the `spi_hal_seg_config_t`, update the conf buffer
 *
 * @param hal            Context of the HAL layer.
 * @param config         Conf buffer configuration, per segment. See `spi_hal_seg_config_t` to know what can be configured
 * @param conf_buffer    Conf buffer
 */
// llgo:link (*SpiHalContextT).SpiHalSctFormatConfBuffer C.spi_hal_sct_format_conf_buffer
func (recv_ *SpiHalContextT) SpiHalSctFormatConfBuffer(config *SpiHalSegConfigT, dev *SpiHalDevConfigT, conf_buffer *c.Uint32T) {
}

/**
 * Deinit SCT mode related registers and hal states
 */
// llgo:link (*SpiHalContextT).SpiHalSctDeinit C.spi_hal_sct_deinit
func (recv_ *SpiHalContextT) SpiHalSctDeinit() {
}

/**
 * Set conf_bitslen to HW for sct.
 */
// llgo:link (*SpiHalContextT).SpiHalSctSetConfBitsLen C.spi_hal_sct_set_conf_bits_len
func (recv_ *SpiHalContextT) SpiHalSctSetConfBitsLen(conf_len c.Uint32T) {
}

/**
 * Clear SPI interrupt bits by mask
 */
// llgo:link (*SpiHalContextT).SpiHalClearIntrMask C.spi_hal_clear_intr_mask
func (recv_ *SpiHalContextT) SpiHalClearIntrMask(mask c.Uint32T) {
}

/**
 * Get SPI interrupt bits status by mask
 */
// llgo:link (*SpiHalContextT).SpiHalGetIntrMask C.spi_hal_get_intr_mask
func (recv_ *SpiHalContextT) SpiHalGetIntrMask(mask c.Uint32T) bool {
	return false
}
