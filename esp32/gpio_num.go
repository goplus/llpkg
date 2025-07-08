package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GpioNumT c.Int

const (
	GPIO_NUM_NC  GpioNumT = -1
	GPIO_NUM_0   GpioNumT = 0
	GPIO_NUM_1   GpioNumT = 1
	GPIO_NUM_2   GpioNumT = 2
	GPIO_NUM_3   GpioNumT = 3
	GPIO_NUM_4   GpioNumT = 4
	GPIO_NUM_5   GpioNumT = 5
	GPIO_NUM_6   GpioNumT = 6
	GPIO_NUM_7   GpioNumT = 7
	GPIO_NUM_8   GpioNumT = 8
	GPIO_NUM_9   GpioNumT = 9
	GPIO_NUM_10  GpioNumT = 10
	GPIO_NUM_11  GpioNumT = 11
	GPIO_NUM_12  GpioNumT = 12
	GPIO_NUM_13  GpioNumT = 13
	GPIO_NUM_14  GpioNumT = 14
	GPIO_NUM_15  GpioNumT = 15
	GPIO_NUM_16  GpioNumT = 16
	GPIO_NUM_17  GpioNumT = 17
	GPIO_NUM_18  GpioNumT = 18
	GPIO_NUM_19  GpioNumT = 19
	GPIO_NUM_20  GpioNumT = 20
	GPIO_NUM_21  GpioNumT = 21
	GPIO_NUM_26  GpioNumT = 26
	GPIO_NUM_27  GpioNumT = 27
	GPIO_NUM_28  GpioNumT = 28
	GPIO_NUM_29  GpioNumT = 29
	GPIO_NUM_30  GpioNumT = 30
	GPIO_NUM_31  GpioNumT = 31
	GPIO_NUM_32  GpioNumT = 32
	GPIO_NUM_33  GpioNumT = 33
	GPIO_NUM_34  GpioNumT = 34
	GPIO_NUM_35  GpioNumT = 35
	GPIO_NUM_36  GpioNumT = 36
	GPIO_NUM_37  GpioNumT = 37
	GPIO_NUM_38  GpioNumT = 38
	GPIO_NUM_39  GpioNumT = 39
	GPIO_NUM_40  GpioNumT = 40
	GPIO_NUM_41  GpioNumT = 41
	GPIO_NUM_42  GpioNumT = 42
	GPIO_NUM_43  GpioNumT = 43
	GPIO_NUM_44  GpioNumT = 44
	GPIO_NUM_45  GpioNumT = 45
	GPIO_NUM_46  GpioNumT = 46
	GPIO_NUM_47  GpioNumT = 47
	GPIO_NUM_48  GpioNumT = 48
	GPIO_NUM_MAX GpioNumT = 49
)
