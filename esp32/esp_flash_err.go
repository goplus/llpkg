package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const (
	ESP_ERR_FLASH_SIZE_NOT_MATCH c.Int = 260
	ESP_ERR_FLASH_NO_RESPONSE    c.Int = 264
)
