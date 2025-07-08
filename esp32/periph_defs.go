package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type PeriphModuleT c.Int

const (
	PERIPH_LEDC_MODULE           PeriphModuleT = 0
	PERIPH_UART0_MODULE          PeriphModuleT = 1
	PERIPH_UART1_MODULE          PeriphModuleT = 2
	PERIPH_UART2_MODULE          PeriphModuleT = 3
	PERIPH_USB_MODULE            PeriphModuleT = 4
	PERIPH_I2C0_MODULE           PeriphModuleT = 5
	PERIPH_I2C1_MODULE           PeriphModuleT = 6
	PERIPH_I2S0_MODULE           PeriphModuleT = 7
	PERIPH_I2S1_MODULE           PeriphModuleT = 8
	PERIPH_LCD_CAM_MODULE        PeriphModuleT = 9
	PERIPH_TIMG0_MODULE          PeriphModuleT = 10
	PERIPH_TIMG1_MODULE          PeriphModuleT = 11
	PERIPH_PWM0_MODULE           PeriphModuleT = 12
	PERIPH_PWM1_MODULE           PeriphModuleT = 13
	PERIPH_PWM2_MODULE           PeriphModuleT = 14
	PERIPH_PWM3_MODULE           PeriphModuleT = 15
	PERIPH_UHCI0_MODULE          PeriphModuleT = 16
	PERIPH_UHCI1_MODULE          PeriphModuleT = 17
	PERIPH_RMT_MODULE            PeriphModuleT = 18
	PERIPH_PCNT_MODULE           PeriphModuleT = 19
	PERIPH_SPI_MODULE            PeriphModuleT = 20
	PERIPH_SPI2_MODULE           PeriphModuleT = 21
	PERIPH_SPI3_MODULE           PeriphModuleT = 22
	PERIPH_SDMMC_MODULE          PeriphModuleT = 23
	PERIPH_TWAI_MODULE           PeriphModuleT = 24
	PERIPH_RNG_MODULE            PeriphModuleT = 25
	PERIPH_WIFI_MODULE           PeriphModuleT = 26
	PERIPH_BT_MODULE             PeriphModuleT = 27
	PERIPH_WIFI_BT_COMMON_MODULE PeriphModuleT = 28
	PERIPH_BT_BASEBAND_MODULE    PeriphModuleT = 29
	PERIPH_BT_LC_MODULE          PeriphModuleT = 30
	PERIPH_AES_MODULE            PeriphModuleT = 31
	PERIPH_SHA_MODULE            PeriphModuleT = 32
	PERIPH_HMAC_MODULE           PeriphModuleT = 33
	PERIPH_DS_MODULE             PeriphModuleT = 34
	PERIPH_RSA_MODULE            PeriphModuleT = 35
	PERIPH_SYSTIMER_MODULE       PeriphModuleT = 36
	PERIPH_GDMA_MODULE           PeriphModuleT = 37
	PERIPH_DEDIC_GPIO_MODULE     PeriphModuleT = 38
	PERIPH_SARADC_MODULE         PeriphModuleT = 39
	PERIPH_TEMPSENSOR_MODULE     PeriphModuleT = 40
	PERIPH_MODULE_MAX            PeriphModuleT = 41
)
