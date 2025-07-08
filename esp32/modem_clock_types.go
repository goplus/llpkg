package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type ModemClockDomainT c.Int

const (
	MODEM_CLOCK_DOMAIN_MODEM_APB    ModemClockDomainT = 0
	MODEM_CLOCK_DOMAIN_MODEM_PERIPH ModemClockDomainT = 1
	MODEM_CLOCK_DOMAIN_WIFI         ModemClockDomainT = 2
	MODEM_CLOCK_DOMAIN_BT           ModemClockDomainT = 3
	MODEM_CLOCK_DOMAIN_MODEM_FE     ModemClockDomainT = 4
	MODEM_CLOCK_DOMAIN_IEEE802154   ModemClockDomainT = 5
	MODEM_CLOCK_DOMAIN_LP_APB       ModemClockDomainT = 6
	MODEM_CLOCK_DOMAIN_I2C_MASTER   ModemClockDomainT = 7
	MODEM_CLOCK_DOMAIN_COEX         ModemClockDomainT = 8
	MODEM_CLOCK_DOMAIN_WIFIPWR      ModemClockDomainT = 9
	MODEM_CLOCK_DOMAIN_MAX          ModemClockDomainT = 10
)

type ModemClockLpclkSrcT c.Int

const (
	MODEM_CLOCK_LPCLK_SRC_INVALID   ModemClockLpclkSrcT = -1
	MODEM_CLOCK_LPCLK_SRC_RC_SLOW   ModemClockLpclkSrcT = 0
	MODEM_CLOCK_LPCLK_SRC_RC_FAST   ModemClockLpclkSrcT = 1
	MODEM_CLOCK_LPCLK_SRC_MAIN_XTAL ModemClockLpclkSrcT = 2
	MODEM_CLOCK_LPCLK_SRC_RC32K     ModemClockLpclkSrcT = 3
	MODEM_CLOCK_LPCLK_SRC_XTAL32K   ModemClockLpclkSrcT = 4
	MODEM_CLOCK_LPCLK_SRC_EXT32K    ModemClockLpclkSrcT = 5
	MODEM_CLOCK_LPCLK_SRC_MAX       ModemClockLpclkSrcT = 6
)
