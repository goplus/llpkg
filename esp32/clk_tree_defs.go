package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SOC_CLK_RC_FAST_FREQ_APPROX = 17500000
const SOC_CLK_RC_SLOW_FREQ_APPROX = 136000
const SOC_CLK_XTAL32K_FREQ_APPROX = 32768

type SocRootClkT c.Int

const (
	SOC_ROOT_CLK_INT_RC_FAST SocRootClkT = 0
	SOC_ROOT_CLK_INT_RC_SLOW SocRootClkT = 1
	SOC_ROOT_CLK_EXT_XTAL    SocRootClkT = 2
	SOC_ROOT_CLK_EXT_XTAL32K SocRootClkT = 3
)

type SocCpuClkSrcT c.Int

const (
	SOC_CPU_CLK_SRC_XTAL    SocCpuClkSrcT = 0
	SOC_CPU_CLK_SRC_PLL     SocCpuClkSrcT = 1
	SOC_CPU_CLK_SRC_RC_FAST SocCpuClkSrcT = 2
	SOC_CPU_CLK_SRC_INVALID SocCpuClkSrcT = 3
)

type SocRtcSlowClkSrcT c.Int

const (
	SOC_RTC_SLOW_CLK_SRC_RC_SLOW      SocRtcSlowClkSrcT = 0
	SOC_RTC_SLOW_CLK_SRC_XTAL32K      SocRtcSlowClkSrcT = 1
	SOC_RTC_SLOW_CLK_SRC_RC_FAST_D256 SocRtcSlowClkSrcT = 2
	SOC_RTC_SLOW_CLK_SRC_INVALID      SocRtcSlowClkSrcT = 3
)

type SocRtcFastClkSrcT c.Int

const (
	SOC_RTC_FAST_CLK_SRC_XTAL_D2  SocRtcFastClkSrcT = 0
	SOC_RTC_FAST_CLK_SRC_XTAL_DIV SocRtcFastClkSrcT = 0
	SOC_RTC_FAST_CLK_SRC_RC_FAST  SocRtcFastClkSrcT = 1
	SOC_RTC_FAST_CLK_SRC_INVALID  SocRtcFastClkSrcT = 2
)

type SocXtalFreqT c.Int

const (
	SOC_XTAL_FREQ_32M SocXtalFreqT = 32
	SOC_XTAL_FREQ_40M SocXtalFreqT = 40
)

type SocModuleClkT c.Int

const (
	SOC_MOD_CLK_CPU          SocModuleClkT = 1
	SOC_MOD_CLK_RTC_FAST     SocModuleClkT = 2
	SOC_MOD_CLK_RTC_SLOW     SocModuleClkT = 3
	SOC_MOD_CLK_APB          SocModuleClkT = 4
	SOC_MOD_CLK_PLL_F80M     SocModuleClkT = 5
	SOC_MOD_CLK_PLL_F160M    SocModuleClkT = 6
	SOC_MOD_CLK_PLL_D2       SocModuleClkT = 7
	SOC_MOD_CLK_XTAL32K      SocModuleClkT = 8
	SOC_MOD_CLK_RC_FAST      SocModuleClkT = 9
	SOC_MOD_CLK_RC_FAST_D256 SocModuleClkT = 10
	SOC_MOD_CLK_XTAL         SocModuleClkT = 11
	SOC_MOD_CLK_TEMP_SENSOR  SocModuleClkT = 12
	SOC_MOD_CLK_INVALID      SocModuleClkT = 13
)

type SocPeriphSystimerClkSrcT c.Int

const (
	SYSTIMER_CLK_SRC_XTAL    SocPeriphSystimerClkSrcT = 11
	SYSTIMER_CLK_SRC_DEFAULT SocPeriphSystimerClkSrcT = 11
)

type SocPeriphGptimerClkSrcT c.Int

const (
	GPTIMER_CLK_SRC_APB     SocPeriphGptimerClkSrcT = 4
	GPTIMER_CLK_SRC_XTAL    SocPeriphGptimerClkSrcT = 11
	GPTIMER_CLK_SRC_DEFAULT SocPeriphGptimerClkSrcT = 4
)

type SocPeriphTgClkSrcLegacyT c.Int

const (
	TIMER_SRC_CLK_APB     SocPeriphTgClkSrcLegacyT = 4
	TIMER_SRC_CLK_XTAL    SocPeriphTgClkSrcLegacyT = 11
	TIMER_SRC_CLK_DEFAULT SocPeriphTgClkSrcLegacyT = 4
)

type SocPeriphLcdClkSrcT c.Int

const (
	LCD_CLK_SRC_PLL160M SocPeriphLcdClkSrcT = 6
	LCD_CLK_SRC_PLL240M SocPeriphLcdClkSrcT = 7
	LCD_CLK_SRC_XTAL    SocPeriphLcdClkSrcT = 11
	LCD_CLK_SRC_DEFAULT SocPeriphLcdClkSrcT = 6
)

type SocPeriphRmtClkSrcT c.Int

const (
	RMT_CLK_SRC_APB     SocPeriphRmtClkSrcT = 4
	RMT_CLK_SRC_RC_FAST SocPeriphRmtClkSrcT = 9
	RMT_CLK_SRC_XTAL    SocPeriphRmtClkSrcT = 11
	RMT_CLK_SRC_DEFAULT SocPeriphRmtClkSrcT = 4
)

type SocPeriphRmtClkSrcLegacyT c.Int

const (
	RMT_BASECLK_APB     SocPeriphRmtClkSrcLegacyT = 4
	RMT_BASECLK_XTAL    SocPeriphRmtClkSrcLegacyT = 11
	RMT_BASECLK_DEFAULT SocPeriphRmtClkSrcLegacyT = 4
)

type SocPeriphTemperatureSensorClkSrcT c.Int

const (
	TEMPERATURE_SENSOR_CLK_SRC_RC_FAST SocPeriphTemperatureSensorClkSrcT = 12
	TEMPERATURE_SENSOR_CLK_SRC_DEFAULT SocPeriphTemperatureSensorClkSrcT = 12
)

type SocPeriphUartClkSrcLegacyT c.Int

const (
	UART_SCLK_APB     SocPeriphUartClkSrcLegacyT = 4
	UART_SCLK_RTC     SocPeriphUartClkSrcLegacyT = 9
	UART_SCLK_XTAL    SocPeriphUartClkSrcLegacyT = 11
	UART_SCLK_DEFAULT SocPeriphUartClkSrcLegacyT = 4
)

type SocPeriphMcpwmTimerClkSrcT c.Int

const (
	MCPWM_TIMER_CLK_SRC_PLL160M SocPeriphMcpwmTimerClkSrcT = 6
	MCPWM_TIMER_CLK_SRC_DEFAULT SocPeriphMcpwmTimerClkSrcT = 6
)

type SocPeriphMcpwmCaptureClkSrcT c.Int

const (
	MCPWM_CAPTURE_CLK_SRC_APB     SocPeriphMcpwmCaptureClkSrcT = 4
	MCPWM_CAPTURE_CLK_SRC_DEFAULT SocPeriphMcpwmCaptureClkSrcT = 4
)

type SocPeriphMcpwmCarrierClkSrcT c.Int

const (
	MCPWM_CARRIER_CLK_SRC_PLL160M SocPeriphMcpwmCarrierClkSrcT = 6
	MCPWM_CARRIER_CLK_SRC_DEFAULT SocPeriphMcpwmCarrierClkSrcT = 6
)

type SocPeriphI2sClkSrcT c.Int

const (
	I2S_CLK_SRC_DEFAULT  SocPeriphI2sClkSrcT = 6
	I2S_CLK_SRC_PLL_160M SocPeriphI2sClkSrcT = 6
	I2S_CLK_SRC_XTAL     SocPeriphI2sClkSrcT = 11
	I2S_CLK_SRC_EXTERNAL SocPeriphI2sClkSrcT = -1
)

type SocPeriphI2cClkSrcT c.Int

const (
	I2C_CLK_SRC_XTAL    SocPeriphI2cClkSrcT = 11
	I2C_CLK_SRC_RC_FAST SocPeriphI2cClkSrcT = 9
	I2C_CLK_SRC_DEFAULT SocPeriphI2cClkSrcT = 11
)

type SocPeriphSpiClkSrcT c.Int

const (
	SPI_CLK_SRC_DEFAULT SocPeriphSpiClkSrcT = 4
	SPI_CLK_SRC_APB     SocPeriphSpiClkSrcT = 4
	SPI_CLK_SRC_XTAL    SocPeriphSpiClkSrcT = 11
)

type SocPeriphSdmClkSrcT c.Int

const (
	SDM_CLK_SRC_APB     SocPeriphSdmClkSrcT = 4
	SDM_CLK_SRC_DEFAULT SocPeriphSdmClkSrcT = 4
)

type SocPeriphGlitchFilterClkSrcT c.Int

const (
	GLITCH_FILTER_CLK_SRC_APB     SocPeriphGlitchFilterClkSrcT = 4
	GLITCH_FILTER_CLK_SRC_DEFAULT SocPeriphGlitchFilterClkSrcT = 4
)

type SocPeriphTwaiClkSrcT c.Int

const (
	TWAI_CLK_SRC_APB     SocPeriphTwaiClkSrcT = 4
	TWAI_CLK_SRC_DEFAULT SocPeriphTwaiClkSrcT = 4
)

type SocPeriphAdcDigiClkSrcT c.Int

const (
	ADC_DIGI_CLK_SRC_APB       SocPeriphAdcDigiClkSrcT = 4
	ADC_DIGI_CLK_SRC_PLL_F240M SocPeriphAdcDigiClkSrcT = 7
	ADC_DIGI_CLK_SRC_DEFAULT   SocPeriphAdcDigiClkSrcT = 4
)

type SocPeriphAdcRtcClkSrcT c.Int

const (
	ADC_RTC_CLK_SRC_RC_FAST SocPeriphAdcRtcClkSrcT = 9
	ADC_RTC_CLK_SRC_DEFAULT SocPeriphAdcRtcClkSrcT = 9
)

type SocPeriphMwdtClkSrcT c.Int

const (
	MWDT_CLK_SRC_APB     SocPeriphMwdtClkSrcT = 4
	MWDT_CLK_SRC_DEFAULT SocPeriphMwdtClkSrcT = 4
)

type SocPeriphLedcClkSrcLegacyT c.Int

const (
	LEDC_AUTO_CLK        SocPeriphLedcClkSrcLegacyT = 0
	LEDC_USE_APB_CLK     SocPeriphLedcClkSrcLegacyT = 4
	LEDC_USE_RC_FAST_CLK SocPeriphLedcClkSrcLegacyT = 9
	LEDC_USE_XTAL_CLK    SocPeriphLedcClkSrcLegacyT = 11
	LEDC_USE_RTC8M_CLK   SocPeriphLedcClkSrcLegacyT = 9
)

type SocPeriphSdmmcClkSrcT c.Int

const (
	SDMMC_CLK_SRC_DEFAULT SocPeriphSdmmcClkSrcT = 6
	SDMMC_CLK_SRC_PLL160M SocPeriphSdmmcClkSrcT = 6
	SDMMC_CLK_SRC_XTAL    SocPeriphSdmmcClkSrcT = 11
)

type SocClkoutSigIdT c.Int

const (
	CLKOUT_SIG_PLL      SocClkoutSigIdT = 1
	CLKOUT_SIG_RC_SLOW  SocClkoutSigIdT = 4
	CLKOUT_SIG_XTAL     SocClkoutSigIdT = 5
	CLKOUT_SIG_PLL_F80M SocClkoutSigIdT = 13
	CLKOUT_SIG_RC_FAST  SocClkoutSigIdT = 14
	CLKOUT_SIG_INVALID  SocClkoutSigIdT = 255
)
