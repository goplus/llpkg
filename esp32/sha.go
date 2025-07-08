package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SHATYPE c.Int

const (
	SHA1         SHATYPE = 0
	SHA2_224     SHATYPE = 1
	SHA2_256     SHATYPE = 2
	SHA2_384     SHATYPE = 3
	SHA2_512     SHATYPE = 4
	SHA2_512224  SHATYPE = 5
	SHA2_512256  SHATYPE = 6
	SHA2_512T    SHATYPE = 7
	SHA_TYPE_MAX SHATYPE = 8
)

type SHAContext struct {
	Start      bool
	InHardware bool
	Type       SHATYPE
	State      [16]c.Uint32T
	Buffer     [128]c.Char
	TotalBits  [4]c.Uint32T
}
type SHACTX SHAContext

//go:linkname EtsShaEnable C.ets_sha_enable
func EtsShaEnable()

//go:linkname EtsShaDisable C.ets_sha_disable
func EtsShaDisable()

// llgo:link (*SHACTX).EtsShaInit C.ets_sha_init
func (recv_ *SHACTX) EtsShaInit(type_ SHATYPE) EtsStatusT {
	return 0
}

// llgo:link (*SHACTX).EtsShaStarts C.ets_sha_starts
func (recv_ *SHACTX) EtsShaStarts(sha512_t c.Uint16T) EtsStatusT {
	return 0
}

// llgo:link (*SHACTX).EtsShaGetState C.ets_sha_get_state
func (recv_ *SHACTX) EtsShaGetState() {
}

// llgo:link (*SHACTX).EtsShaProcess C.ets_sha_process
func (recv_ *SHACTX) EtsShaProcess(input *c.Char) {
}

// llgo:link (*SHACTX).EtsShaUpdate C.ets_sha_update
func (recv_ *SHACTX) EtsShaUpdate(input *c.Char, input_bytes c.Uint32T, update_ctx bool) {
}

// llgo:link (*SHACTX).EtsShaFinish C.ets_sha_finish
func (recv_ *SHACTX) EtsShaFinish(output *c.Char) EtsStatusT {
	return 0
}
