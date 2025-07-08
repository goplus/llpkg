package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type UartPortT c.Int

const (
	UART_NUM_0   UartPortT = 0
	UART_NUM_1   UartPortT = 1
	UART_NUM_2   UartPortT = 2
	UART_NUM_MAX UartPortT = 3
)

type UartModeT c.Int

const (
	UART_MODE_UART                   UartModeT = 0
	UART_MODE_RS485_HALF_DUPLEX      UartModeT = 1
	UART_MODE_IRDA                   UartModeT = 2
	UART_MODE_RS485_COLLISION_DETECT UartModeT = 3
	UART_MODE_RS485_APP_CTRL         UartModeT = 4
)

type UartWordLengthT c.Int

const (
	UART_DATA_5_BITS   UartWordLengthT = 0
	UART_DATA_6_BITS   UartWordLengthT = 1
	UART_DATA_7_BITS   UartWordLengthT = 2
	UART_DATA_8_BITS   UartWordLengthT = 3
	UART_DATA_BITS_MAX UartWordLengthT = 4
)

type UartStopBitsT c.Int

const (
	UART_STOP_BITS_1   UartStopBitsT = 1
	UART_STOP_BITS_1_5 UartStopBitsT = 2
	UART_STOP_BITS_2   UartStopBitsT = 3
	UART_STOP_BITS_MAX UartStopBitsT = 4
)

type UartParityT c.Int

const (
	UART_PARITY_DISABLE UartParityT = 0
	UART_PARITY_EVEN    UartParityT = 2
	UART_PARITY_ODD     UartParityT = 3
)

type UartHwFlowcontrolT c.Int

const (
	UART_HW_FLOWCTRL_DISABLE UartHwFlowcontrolT = 0
	UART_HW_FLOWCTRL_RTS     UartHwFlowcontrolT = 1
	UART_HW_FLOWCTRL_CTS     UartHwFlowcontrolT = 2
	UART_HW_FLOWCTRL_CTS_RTS UartHwFlowcontrolT = 3
	UART_HW_FLOWCTRL_MAX     UartHwFlowcontrolT = 4
)

type UartSignalInvT c.Int

const (
	UART_SIGNAL_INV_DISABLE UartSignalInvT = 0
	UART_SIGNAL_IRDA_TX_INV UartSignalInvT = 1
	UART_SIGNAL_IRDA_RX_INV UartSignalInvT = 2
	UART_SIGNAL_RXD_INV     UartSignalInvT = 4
	UART_SIGNAL_CTS_INV     UartSignalInvT = 8
	UART_SIGNAL_DSR_INV     UartSignalInvT = 16
	UART_SIGNAL_TXD_INV     UartSignalInvT = 32
	UART_SIGNAL_RTS_INV     UartSignalInvT = 64
	UART_SIGNAL_DTR_INV     UartSignalInvT = 128
)

type UartSclkT SocPeriphUartClkSrcLegacyT

/**
 * @brief UART AT cmd char configuration parameters
 *        Note that this function may different on different chip. Please refer to the TRM at confirguration.
 */

type UartAtCmdT struct {
	CmdChar  c.Uint8T
	CharNum  c.Uint8T
	GapTout  c.Uint32T
	PreIdle  c.Uint32T
	PostIdle c.Uint32T
}

/**
 * @brief UART software flow control configuration parameters
 */

type UartSwFlowctrlT struct {
	XonChar  c.Uint8T
	XoffChar c.Uint8T
	XonThrd  c.Uint8T
	XoffThrd c.Uint8T
}
