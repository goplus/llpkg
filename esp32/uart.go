package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const RX_BUFF_SIZE = 0x400
const TX_BUFF_SIZE = 100
const UART_INT_FLAG_MASK = 0x0E
const FRAME_FLAG = 0x7E

type UartIntType c.Int

const (
	UART_LINE_STATUS_INT_FLAG  UartIntType = 6
	UART_RCV_FIFO_INT_FLAG     UartIntType = 4
	UART_RCV_TMOUT_INT_FLAG    UartIntType = 12
	UART_TXBUFF_EMPTY_INT_FLAG UartIntType = 2
)

type UartRcvFifoTrgLvl c.Int

const (
	RCV_ONE_BYTE      UartRcvFifoTrgLvl = 0
	RCV_FOUR_BYTE     UartRcvFifoTrgLvl = 1
	RCV_EIGHT_BYTE    UartRcvFifoTrgLvl = 2
	RCV_FOURTEEN_BYTE UartRcvFifoTrgLvl = 3
)

type UartBitsNum4Char c.Int

const (
	FIVE_BITS  UartBitsNum4Char = 0
	SIX_BITS   UartBitsNum4Char = 1
	SEVEN_BITS UartBitsNum4Char = 2
	EIGHT_BITS UartBitsNum4Char = 3
)

type UartStopBitsNum c.Int

const (
	ONE_STOP_BIT      UartStopBitsNum = 1
	ONE_HALF_STOP_BIT UartStopBitsNum = 2
	TWO_STOP_BIT      UartStopBitsNum = 3
)

type UartParityMode c.Int

const (
	NONE_BITS UartParityMode = 0
	ODD_BITS  UartParityMode = 2
	EVEN_BITS UartParityMode = 3
)

type UartExistParity c.Int

const (
	STICK_PARITY_DIS UartExistParity = 0
	STICK_PARITY_EN  UartExistParity = 2
)

type UartBautRate c.Int

const (
	BIT_RATE_9600   UartBautRate = 9600
	BIT_RATE_19200  UartBautRate = 19200
	BIT_RATE_38400  UartBautRate = 38400
	BIT_RATE_57600  UartBautRate = 57600
	BIT_RATE_115200 UartBautRate = 115200
	BIT_RATE_230400 UartBautRate = 230400
	BIT_RATE_460800 UartBautRate = 460800
	BIT_RATE_921600 UartBautRate = 921600
)

type UartFlowCtrl c.Int

const (
	NONE_CTRL     UartFlowCtrl = 0
	HARDWARE_CTRL UartFlowCtrl = 1
	XON_XOFF_CTRL UartFlowCtrl = 2
)

type RcvMsgBuffState c.Int

const (
	EMPTY       RcvMsgBuffState = 0
	UNDER_WRITE RcvMsgBuffState = 1
	WRITE_OVER  RcvMsgBuffState = 2
)

type RcvMsgBuff struct {
	PRcvMsgBuff *c.Uint8T
	PWritePos   *c.Uint8T
	PReadPos    *c.Uint8T
	TrigLvl     c.Uint8T
	BuffState   RcvMsgBuffState
}

type TrxMsgBuff struct {
	TrxBuffSize c.Uint32T
	PTrxBuff    *c.Uint8T
}
type RcvMsgState c.Int

const (
	BAUD_RATE_DET RcvMsgState = 0
	WAIT_SYNC_FRM RcvMsgState = 1
	SRCH_MSG_HEAD RcvMsgState = 2
	RCV_MSG_BODY  RcvMsgState = 3
	RCV_ESC_CHAR  RcvMsgState = 4
)

type UartDevice struct {
	BautRate    UartBautRate
	DataBits    UartBitsNum4Char
	ExistParity UartExistParity
	Parity      UartParityMode
	StopBits    UartStopBitsNum
	FlowCtrl    UartFlowCtrl
	BuffUartNo  c.Uint8T
	RcvBuff     RcvMsgBuff
	RcvState    RcvMsgState
	Received    c.Int
}

/**
 * @brief Init uart device struct value and reset uart0/uart1 rx.
 *        Please do not call this function in SDK.
 *
 * @param  rxBuffer, must be a pointer to RX_BUFF_SIZE bytes or NULL
 *
 * @return None
 */
//go:linkname UartAttach C.uartAttach
func UartAttach(rxBuffer c.Pointer)

/**
 * @brief Init uart0 or uart1 for UART download booting mode.
 *        Please do not call this function in SDK.
 *
 * @param  uint8_t uart_no : 0 for UART0, else for UART1.
 *
 * @param  uint32_t clock : clock used by uart module, to adjust baudrate.
 *
 * @return None
 */
//go:linkname UartInit C.Uart_Init
func UartInit(uart_no c.Uint8T, clock c.Uint32T)

/**
 * @brief Modify uart baudrate.
 *        This function will reset RX/TX fifo for uart.
 *
 * @param  uint8_t uart_no : 0 for UART0, 1 for UART1.
 *
 * @param  uint32_t DivLatchValue : (clock << 4)/baudrate.
 *
 * @return None
 */
//go:linkname UartDivModify C.uart_div_modify
func UartDivModify(uart_no c.Uint8T, DivLatchValue c.Uint32T)

/**
 * @brief Switch printf channel of uart_tx_one_char.
 *        Please do not call this function when printf.
 *
 * @param  uint8_t uart_no : 0 for UART0, 1 for UART1.
 *
 * @return None
 */
//go:linkname UartTxSwitch C.uart_tx_switch
func UartTxSwitch(uart_no c.Uint8T)

/**
 * @brief Output a char to printf channel, wait until fifo not full.
 *
 * @param  None
 *
 * @return OK.
 */
//go:linkname UartTxOneChar C.uart_tx_one_char
func UartTxOneChar(TxChar c.Uint8T) ETSSTATUS

/**
 * @brief Output a char to message exchange channel, wait until fifo not full.
 *        Please do not call this function in SDK.
 *
 * @param  None
 *
 * @return OK.
 */
//go:linkname UartTxOneChar2 C.uart_tx_one_char2
func UartTxOneChar2(TxChar c.Uint8T) ETSSTATUS

/**
 * @brief Wait until uart tx full empty.
 *
 * @param  uint8_t uart_no : 0 for UART0, 1 for UART1.
 *
 * @return None.
 */
//go:linkname UartTxFlush C.uart_tx_flush
func UartTxFlush(uart_no c.Uint8T)

/**
 * @brief Wait until uart tx full empty and the last char send ok.
 *
 * @param  uart_no : 0 for UART0, 1 for UART1, 2 for UART2
 *
 * The function defined in ROM code has a bug, so we define the correct version
 * here for compatibility.
 */
//go:linkname UartTxWaitIdle C.uart_tx_wait_idle
func UartTxWaitIdle(uart_no c.Uint8T)

/**
 * @brief Get an input char from message channel.
 *        Please do not call this function in SDK.
 *
 * @param  uint8_t *pRxChar : the pointer to store the char.
 *
 * @return OK for successful.
 *         FAIL for failed.
 */
//go:linkname UartRxOneChar C.uart_rx_one_char
func UartRxOneChar(pRxChar *c.Uint8T) ETSSTATUS

/**
 * @brief Get an input char from message channel, wait until successful.
 *        Please do not call this function in SDK.
 *
 * @param  None
 *
 * @return char : input char value.
 */
//go:linkname UartRxOneCharBlock C.uart_rx_one_char_block
func UartRxOneCharBlock() c.Char

/**
 * @brief Get an input string line from message channel.
 *        Please do not call this function in SDK.
 *
 * @param  uint8_t *pString : the pointer to store the string.
 *
 * @param  uint8_t MaxStrlen : the max string length, include '\0'.
 *
 * @return OK.
 */
//go:linkname UartRxString C.UartRxString
func UartRxString(pString *c.Uint8T, MaxStrlen c.Uint8T) ETSSTATUS

/**
 * @brief Get an char from receive buffer.
 *        Please do not call this function in SDK.
 *
 * @param  RcvMsgBuff *pRxBuff : the pointer to the struct that include receive buffer.
 *
 * @param  uint8_t *pRxByte : the pointer to store the char.
 *
 * @return OK for successful.
 *         FAIL for failed.
 */
// llgo:link (*RcvMsgBuff).UartRxReadbuff C.uart_rx_readbuff
func (recv_ *RcvMsgBuff) UartRxReadbuff(pRxByte *c.Uint8T) ETSSTATUS {
	return 0
}

/**
 * @brief Get uart configuration struct.
 *        Please do not call this function in SDK.
 *
 * @param  None
 *
 * @return UartDevice * : uart configuration struct pointer.
 */
//go:linkname GetUartDevice C.GetUartDevice
func GetUartDevice() *UartDevice

/**
 * @brief Send an packet to download tool, with SLIP escaping.
 *        Please do not call this function in SDK.
 *
 * @param  uint8_t *p : the pointer to output string.
 *
 * @param  int len : the string length.
 *
 * @return None.
 */
//go:linkname SendPacket C.send_packet
func SendPacket(p *c.Uint8T, len c.Int)

/**
 * @brief Receive an packet from download tool, with SLIP escaping.
 *        Please do not call this function in SDK.
 *
 * @param  uint8_t *p : the pointer to input string.
 *
 * @param  int len : If string length > len, the string will be truncated.
 *
 * @param  uint8_t is_sync : 0, only one UART module;
 *                           1, two UART modules.
 *
 * @return int : the length of the string.
 */
//go:linkname RecvPacket C.recv_packet
func RecvPacket(p *c.Uint8T, len c.Int, is_sync c.Uint8T) c.Int

/**
 * @brief Initialize the USB ACM UART
 * Needs to be fed a buffer of at least 128 bytes, plus any rx buffer you may want to have.
 *
 * @param cdc_acm_work_mem Pointer to work mem for CDC-ACM code
 * @param cdc_acm_work_mem_len Length of work mem
 */
//go:linkname UartInitUSB C.Uart_Init_USB
func UartInitUSB(cdc_acm_work_mem c.Pointer, cdc_acm_work_mem_len c.Int)
