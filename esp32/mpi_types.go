package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type MpiOpT c.Int

const (
	MPI_MULT    MpiOpT = 0
	MPI_MODMULT MpiOpT = 1
	MPI_MODEXP  MpiOpT = 2
)

type MpiParamT c.Int

const (
	MPI_PARAM_X MpiParamT = 0
	MPI_PARAM_Y MpiParamT = 1
	MPI_PARAM_Z MpiParamT = 2
	MPI_PARAM_M MpiParamT = 3
)
