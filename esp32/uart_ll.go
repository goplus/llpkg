package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type UartIntrT c.Int

const (
	UART_INTR_RXFIFO_FULL      UartIntrT = 1
	UART_INTR_TXFIFO_EMPTY     UartIntrT = 2
	UART_INTR_PARITY_ERR       UartIntrT = 4
	UART_INTR_FRAM_ERR         UartIntrT = 8
	UART_INTR_RXFIFO_OVF       UartIntrT = 16
	UART_INTR_DSR_CHG          UartIntrT = 32
	UART_INTR_CTS_CHG          UartIntrT = 64
	UART_INTR_BRK_DET          UartIntrT = 128
	UART_INTR_RXFIFO_TOUT      UartIntrT = 256
	UART_INTR_SW_XON           UartIntrT = 512
	UART_INTR_SW_XOFF          UartIntrT = 1024
	UART_INTR_GLITCH_DET       UartIntrT = 2048
	UART_INTR_TX_BRK_DONE      UartIntrT = 4096
	UART_INTR_TX_BRK_IDLE      UartIntrT = 8192
	UART_INTR_TX_DONE          UartIntrT = 16384
	UART_INTR_RS485_PARITY_ERR UartIntrT = 32768
	UART_INTR_RS485_FRM_ERR    UartIntrT = 65536
	UART_INTR_RS485_CLASH      UartIntrT = 131072
	UART_INTR_CMD_CHAR_DET     UartIntrT = 262144
	UART_INTR_WAKEUP           UartIntrT = 524288
)
