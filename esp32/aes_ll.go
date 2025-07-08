package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspAesStateT c.Int

const (
	ESP_AES_STATE_IDLE EspAesStateT = 0
	ESP_AES_STATE_BUSY EspAesStateT = 1
	ESP_AES_STATE_DONE EspAesStateT = 2
)
