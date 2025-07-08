package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspInterfaceT c.Int

const (
	ESP_IF_WIFI_STA EspInterfaceT = 0
	ESP_IF_WIFI_AP  EspInterfaceT = 1
	ESP_IF_WIFI_NAN EspInterfaceT = 2
	ESP_IF_ETH      EspInterfaceT = 3
	ESP_IF_MAX      EspInterfaceT = 4
)
