package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type CacheTypeT c.Int

const (
	CACHE_TYPE_DATA        CacheTypeT = 0
	CACHE_TYPE_INSTRUCTION CacheTypeT = 1
	CACHE_TYPE_ALL         CacheTypeT = 2
)

type CacheBusMaskT c.Int

const (
	CACHE_BUS_IBUS0 CacheBusMaskT = 1
	CACHE_BUS_IBUS1 CacheBusMaskT = 2
	CACHE_BUS_IBUS2 CacheBusMaskT = 4
	CACHE_BUS_DBUS0 CacheBusMaskT = 8
	CACHE_BUS_DBUS1 CacheBusMaskT = 16
	CACHE_BUS_DBUS2 CacheBusMaskT = 32
)
