package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const GPIO_ID_PIN0 = 0
const GPIO_FUNC_IN_HIGH = 0x38
const GPIO_FUNC_IN_LOW = 0x3C

type GPIOINTTYPE c.Int

const (
	GPIO_PIN_INTR_DISABLE GPIOINTTYPE = 0
	GPIO_PIN_INTR_POSEDGE GPIOINTTYPE = 1
	GPIO_PIN_INTR_NEGEDGE GPIOINTTYPE = 2
	GPIO_PIN_INTR_ANYEDGE GPIOINTTYPE = 3
	GPIO_PIN_INTR_LOLEVEL GPIOINTTYPE = 4
	GPIO_PIN_INTR_HILEVEL GPIOINTTYPE = 5
)

/**
 * @brief Change GPIO(0-31) pin output by setting, clearing, or disabling pins, GPIO0<->BIT(0).
 *         There is no particular ordering guaranteed; so if the order of writes is significant,
 *         calling code should divide a single call into multiple calls.
 *
 * @param  uint32_t set_mask : the gpios that need high level.
 *
 * @param  uint32_t clear_mask : the gpios that need low level.
 *
 * @param  uint32_t enable_mask : the gpios that need be changed.
 *
 * @param  uint32_t disable_mask : the gpios that need disable output.
 *
 * @return None
 */
//go:linkname GpioOutputSet C.gpio_output_set
func GpioOutputSet(set_mask c.Uint32T, clear_mask c.Uint32T, enable_mask c.Uint32T, disable_mask c.Uint32T)

/**
 * @brief Sample the value of GPIO input pins(0-31) and returns a bitmask.
 *
 * @param None
 *
 * @return uint32_t : bitmask for GPIO input pins, BIT(0) for GPIO0.
 */
//go:linkname GpioInputGet C.gpio_input_get
func GpioInputGet() c.Uint32T

/**
 * @brief Set GPIO to wakeup the ESP32.
 *        Please do not call this function in SDK.
 *
 * @param uint32_t i: gpio number.
 *
 * @param GPIO_INT_TYPE intr_state : only GPIO_PIN_INTR_LOLEVEL\GPIO_PIN_INTR_HILEVEL can be used
 *
 * @return None
 */
//go:linkname GpioPinWakeupEnable C.gpio_pin_wakeup_enable
func GpioPinWakeupEnable(i c.Uint32T, intr_state GPIOINTTYPE)

/**
 * @brief disable GPIOs to wakeup the ESP32.
 *        Please do not call this function in SDK.
 *
 * @param None
 *
 * @return None
 */
//go:linkname GpioPinWakeupDisable C.gpio_pin_wakeup_disable
func GpioPinWakeupDisable()

/**
 * @brief set gpio input to a signal, one gpio can input to several signals.
 *
 * @param uint32_t gpio : gpio number, 0~0x2f
 *                        gpio == 0x3C, input 0 to signal
 *                        gpio == 0x3A, input nothing to signal
 *                        gpio == 0x38, input 1 to signal
 *
 * @param uint32_t signal_idx : signal index.
 *
 * @param bool inv : the signal is inv or not
 *
 * @return None
 */
//go:linkname GpioMatrixIn C.gpio_matrix_in
func GpioMatrixIn(gpio c.Uint32T, signal_idx c.Uint32T, inv bool)

/**
 * @brief set signal output to gpio, one signal can output to several gpios.
 *
 * @param uint32_t gpio : gpio number, 0~0x2f
 *
 * @param uint32_t signal_idx : signal index.
 *                        signal_idx == 0x100, cancel output put to the gpio
 *
 * @param bool out_inv : the signal output is invert or not
 *
 * @param bool oen_inv : the signal output enable is invert or not
 *
 * @return None
 */
//go:linkname GpioMatrixOut C.gpio_matrix_out
func GpioMatrixOut(gpio c.Uint32T, signal_idx c.Uint32T, out_inv bool, oen_inv bool)

/**
 * @brief Select pad as a gpio function from IOMUX.
 *
 * @param uint32_t gpio_num : gpio number, 0~0x2f
 *
 * @return None
 */
//go:linkname GpioPadSelectGpio C.gpio_pad_select_gpio
func GpioPadSelectGpio(gpio_num c.Uint32T)

/**
 * @brief Set pad driver capability.
 *
 * @param uint32_t gpio_num : gpio number, 0~0x2f
 *
 * @param uint32_t drv : 0-3
 *
 * @return None
 */
//go:linkname GpioPadSetDrv C.gpio_pad_set_drv
func GpioPadSetDrv(gpio_num c.Uint32T, drv c.Uint32T)

/**
 * @brief Pull up the pad from gpio number.
 *
 * @param uint32_t gpio_num : gpio number, 0~0x2f
 *
 * @return None
 */
//go:linkname GpioPadPullup C.gpio_pad_pullup
func GpioPadPullup(gpio_num c.Uint32T)

/**
 * @brief Pull down the pad from gpio number.
 *
 * @param uint32_t gpio_num : gpio number, 0~0x2f
 *
 * @return None
 */
//go:linkname GpioPadPulldown C.gpio_pad_pulldown
func GpioPadPulldown(gpio_num c.Uint32T)

/**
 * @brief Unhold the pad from gpio number.
 *
 * @param uint32_t gpio_num : gpio number, 0~0x2f
 *
 * @return None
 */
//go:linkname GpioPadUnhold C.gpio_pad_unhold
func GpioPadUnhold(gpio_num c.Uint32T)

/**
 * @brief Hold the pad from gpio number.
 *
 * @param uint32_t gpio_num : gpio number, 0~0x2f
 *
 * @return None
 */
//go:linkname GpioPadHold C.gpio_pad_hold
func GpioPadHold(gpio_num c.Uint32T)

/**
 * @brief enable gpio pad input.
 *
 * @param uint32_t gpio_num : gpio number, 0~0x2f
 *
 * @return None
 */
//go:linkname GpioPadInputEnable C.gpio_pad_input_enable
func GpioPadInputEnable(gpio_num c.Uint32T)

/**
 * @brief disable gpio pad input.
 *
 * @param uint32_t gpio_num : gpio number, 0~0x2f
 *
 * @return None
 */
//go:linkname GpioPadInputDisable C.gpio_pad_input_disable
func GpioPadInputDisable(gpio_num c.Uint32T)
