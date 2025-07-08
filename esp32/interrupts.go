package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type PeriphInterrputT c.Int

const (
	ETS_WIFI_MAC_INTR_SOURCE              PeriphInterrputT = 0
	ETS_WIFI_MAC_NMI_SOURCE               PeriphInterrputT = 1
	ETS_WIFI_PWR_INTR_SOURCE              PeriphInterrputT = 2
	ETS_WIFI_BB_INTR_SOURCE               PeriphInterrputT = 3
	ETS_BT_MAC_INTR_SOURCE                PeriphInterrputT = 4
	ETS_BT_BB_INTR_SOURCE                 PeriphInterrputT = 5
	ETS_BT_BB_NMI_SOURCE                  PeriphInterrputT = 6
	ETS_RWBT_INTR_SOURCE                  PeriphInterrputT = 7
	ETS_RWBLE_INTR_SOURCE                 PeriphInterrputT = 8
	ETS_RWBT_NMI_SOURCE                   PeriphInterrputT = 9
	ETS_RWBLE_NMI_SOURCE                  PeriphInterrputT = 10
	ETS_I2C_MASTER_SOURCE                 PeriphInterrputT = 11
	ETS_SLC0_INTR_SOURCE                  PeriphInterrputT = 12
	ETS_SLC1_INTR_SOURCE                  PeriphInterrputT = 13
	ETS_UHCI0_INTR_SOURCE                 PeriphInterrputT = 14
	ETS_UHCI1_INTR_SOURCE                 PeriphInterrputT = 15
	ETS_GPIO_INTR_SOURCE                  PeriphInterrputT = 16
	ETS_GPIO_NMI_SOURCE                   PeriphInterrputT = 17
	ETS_GPIO_INTR_SOURCE2                 PeriphInterrputT = 18
	ETS_GPIO_NMI_SOURCE2                  PeriphInterrputT = 19
	ETS_SPI1_INTR_SOURCE                  PeriphInterrputT = 20
	ETS_SPI2_INTR_SOURCE                  PeriphInterrputT = 21
	ETS_SPI3_INTR_SOURCE                  PeriphInterrputT = 22
	ETS_LCD_CAM_INTR_SOURCE               PeriphInterrputT = 24
	ETS_I2S0_INTR_SOURCE                  PeriphInterrputT = 25
	ETS_I2S1_INTR_SOURCE                  PeriphInterrputT = 26
	ETS_UART0_INTR_SOURCE                 PeriphInterrputT = 27
	ETS_UART1_INTR_SOURCE                 PeriphInterrputT = 28
	ETS_UART2_INTR_SOURCE                 PeriphInterrputT = 29
	ETS_SDIO_HOST_INTR_SOURCE             PeriphInterrputT = 30
	ETS_PWM0_INTR_SOURCE                  PeriphInterrputT = 31
	ETS_PWM1_INTR_SOURCE                  PeriphInterrputT = 32
	ETS_LEDC_INTR_SOURCE                  PeriphInterrputT = 35
	ETS_EFUSE_INTR_SOURCE                 PeriphInterrputT = 36
	ETS_TWAI_INTR_SOURCE                  PeriphInterrputT = 37
	ETS_USB_INTR_SOURCE                   PeriphInterrputT = 38
	ETS_RTC_CORE_INTR_SOURCE              PeriphInterrputT = 39
	ETS_RMT_INTR_SOURCE                   PeriphInterrputT = 40
	ETS_PCNT_INTR_SOURCE                  PeriphInterrputT = 41
	ETS_I2C_EXT0_INTR_SOURCE              PeriphInterrputT = 42
	ETS_I2C_EXT1_INTR_SOURCE              PeriphInterrputT = 43
	ETS_SPI2_DMA_INTR_SOURCE              PeriphInterrputT = 44
	ETS_SPI3_DMA_INTR_SOURCE              PeriphInterrputT = 45
	ETS_WDT_INTR_SOURCE                   PeriphInterrputT = 47
	ETS_TIMER1_INTR_SOURCE                PeriphInterrputT = 48
	ETS_TIMER2_INTR_SOURCE                PeriphInterrputT = 49
	ETS_TG0_T0_LEVEL_INTR_SOURCE          PeriphInterrputT = 50
	ETS_TG0_T1_LEVEL_INTR_SOURCE          PeriphInterrputT = 51
	ETS_TG0_WDT_LEVEL_INTR_SOURCE         PeriphInterrputT = 52
	ETS_TG1_T0_LEVEL_INTR_SOURCE          PeriphInterrputT = 53
	ETS_TG1_T1_LEVEL_INTR_SOURCE          PeriphInterrputT = 54
	ETS_TG1_WDT_LEVEL_INTR_SOURCE         PeriphInterrputT = 55
	ETS_CACHE_IA_INTR_SOURCE              PeriphInterrputT = 56
	ETS_SYSTIMER_TARGET0_INTR_SOURCE      PeriphInterrputT = 57
	ETS_SYSTIMER_TARGET1_INTR_SOURCE      PeriphInterrputT = 58
	ETS_SYSTIMER_TARGET2_INTR_SOURCE      PeriphInterrputT = 59
	ETS_SYSTIMER_TARGET0_EDGE_INTR_SOURCE PeriphInterrputT = 57
	ETS_SYSTIMER_TARGET1_EDGE_INTR_SOURCE PeriphInterrputT = 58
	ETS_SYSTIMER_TARGET2_EDGE_INTR_SOURCE PeriphInterrputT = 59
	ETS_SPI_MEM_REJECT_CACHE_INTR_SOURCE  PeriphInterrputT = 60
	ETS_DCACHE_PRELOAD0_INTR_SOURCE       PeriphInterrputT = 61
	ETS_ICACHE_PRELOAD0_INTR_SOURCE       PeriphInterrputT = 62
	ETS_DCACHE_SYNC0_INTR_SOURCE          PeriphInterrputT = 63
	ETS_ICACHE_SYNC0_INTR_SOURCE          PeriphInterrputT = 64
	ETS_APB_ADC_INTR_SOURCE               PeriphInterrputT = 65
	ETS_DMA_IN_CH0_INTR_SOURCE            PeriphInterrputT = 66
	ETS_DMA_IN_CH1_INTR_SOURCE            PeriphInterrputT = 67
	ETS_DMA_IN_CH2_INTR_SOURCE            PeriphInterrputT = 68
	ETS_DMA_IN_CH3_INTR_SOURCE            PeriphInterrputT = 69
	ETS_DMA_IN_CH4_INTR_SOURCE            PeriphInterrputT = 70
	ETS_DMA_OUT_CH0_INTR_SOURCE           PeriphInterrputT = 71
	ETS_DMA_OUT_CH1_INTR_SOURCE           PeriphInterrputT = 72
	ETS_DMA_OUT_CH2_INTR_SOURCE           PeriphInterrputT = 73
	ETS_DMA_OUT_CH3_INTR_SOURCE           PeriphInterrputT = 74
	ETS_DMA_OUT_CH4_INTR_SOURCE           PeriphInterrputT = 75
	ETS_RSA_INTR_SOURCE                   PeriphInterrputT = 76
	ETS_AES_INTR_SOURCE                   PeriphInterrputT = 77
	ETS_SHA_INTR_SOURCE                   PeriphInterrputT = 78
	ETS_FROM_CPU_INTR0_SOURCE             PeriphInterrputT = 79
	ETS_FROM_CPU_INTR1_SOURCE             PeriphInterrputT = 80
	ETS_FROM_CPU_INTR2_SOURCE             PeriphInterrputT = 81
	ETS_FROM_CPU_INTR3_SOURCE             PeriphInterrputT = 82
	ETS_ASSIST_DEBUG_INTR_SOURCE          PeriphInterrputT = 83
	ETS_DMA_APBPERI_PMS_INTR_SOURCE       PeriphInterrputT = 84
	ETS_CORE0_IRAM0_PMS_INTR_SOURCE       PeriphInterrputT = 85
	ETS_CORE0_DRAM0_PMS_INTR_SOURCE       PeriphInterrputT = 86
	ETS_CORE0_PIF_PMS_INTR_SOURCE         PeriphInterrputT = 87
	ETS_CORE0_PIF_PMS_SIZE_INTR_SOURCE    PeriphInterrputT = 88
	ETS_CORE1_IRAM0_PMS_INTR_SOURCE       PeriphInterrputT = 89
	ETS_CORE1_DRAM0_PMS_INTR_SOURCE       PeriphInterrputT = 90
	ETS_CORE1_PIF_PMS_INTR_SOURCE         PeriphInterrputT = 91
	ETS_CORE1_PIF_PMS_SIZE_INTR_SOURCE    PeriphInterrputT = 92
	ETS_BACKUP_PMS_VIOLATE_INTR_SOURCE    PeriphInterrputT = 93
	ETS_CACHE_CORE0_ACS_INTR_SOURCE       PeriphInterrputT = 94
	ETS_CACHE_CORE1_ACS_INTR_SOURCE       PeriphInterrputT = 95
	ETS_USB_SERIAL_JTAG_INTR_SOURCE       PeriphInterrputT = 96
	ETS_PERI_BACKUP_INTR_SOURCE           PeriphInterrputT = 97
	ETS_DMA_EXTMEM_REJECT_SOURCE          PeriphInterrputT = 98
	ETS_MAX_INTR_SOURCE                   PeriphInterrputT = 99
)
