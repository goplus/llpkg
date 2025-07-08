package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

//go:linkname EtsBigintEnable C.ets_bigint_enable
func EtsBigintEnable()

//go:linkname EtsBigintDisable C.ets_bigint_disable
func EtsBigintDisable()

//go:linkname EtsBigintMultiply C.ets_bigint_multiply
func EtsBigintMultiply(x *c.Uint32T, y *c.Uint32T, len_words c.Uint32T) c.Int

//go:linkname EtsBigintModmult C.ets_bigint_modmult
func EtsBigintModmult(x *c.Uint32T, y *c.Uint32T, m *c.Uint32T, m_dash c.Uint32T, rb *c.Uint32T, len_words c.Uint32T) c.Int

//go:linkname EtsBigintModexp C.ets_bigint_modexp
func EtsBigintModexp(x *c.Uint32T, y *c.Uint32T, m *c.Uint32T, m_dash c.Uint32T, rb *c.Uint32T, constant_time bool, len_words c.Uint32T) c.Int

//go:linkname EtsBigintWaitFinish C.ets_bigint_wait_finish
func EtsBigintWaitFinish()

//go:linkname EtsBigintGetz C.ets_bigint_getz
func EtsBigintGetz(z *c.Uint32T, len_words c.Uint32T) c.Int
