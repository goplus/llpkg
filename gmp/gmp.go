package gmp

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const X__GMP_HAVE_HOST_CPU_FAMILY_power = 0
const X__GMP_HAVE_HOST_CPU_FAMILY_powerpc = 0
const GMP_LIMB_BITS = 64
const GMP_NAIL_BITS = 0
const X__GNU_MP__ = 6
const X__GMP_LIBGMP_DLL = 0
const X__GMP_MP_SIZE_T_INT = 0
const X__GMP_INLINE_PROTOTYPES = 1
const X__GMP_CC = "gcc"
const X__GMP_CFLAGS = " -O3"
const X__GNU_MP_VERSION = 6
const X__GNU_MP_VERSION_MINOR = 3
const X__GNU_MP_VERSION_PATCHLEVEL = 0

type MpLimbT c.Ulong
type MpLimbSignedT c.Long
type MpBitcntT c.Ulong

/*
For reference, note that the name __mpz_struct gets into C++ mangled

	function names, which means although the "__" suggests an internal, we
	must leave this name for binary compatibility.
*/
type X__mpzStruct struct {
	X_mpAlloc c.Int
	X_mpSize  c.Int
	X_mpD     *MpLimbT
}
type MPINT X__mpzStruct
type MpzT [1]X__mpzStruct
type MpPtr *MpLimbT
type MpSrcptr *MpLimbT
type MpSizeT c.Long
type MpExpT c.Long

type X__mpqStruct struct {
	X_mpNum X__mpzStruct
	X_mpDen X__mpzStruct
}
type MPRAT X__mpqStruct
type MpqT [1]X__mpqStruct

type X__mpfStruct struct {
	X_mpPrec c.Int
	X_mpSize c.Int
	X_mpExp  MpExpT
	X_mpD    *MpLimbT
}
type MpfT [1]X__mpfStruct
type GmpRandalgT c.Int

const (
	GMPRANDALGDEFAULT GmpRandalgT = 0
	GMPRANDALGLC      GmpRandalgT = 0
)

/* Random state struct.  */

type X__gmpRandstateStruct struct {
	X_mpSeed    MpzT
	X_mpAlg     GmpRandalgT
	X_mpAlgdata struct {
		X_mpLc unsafe.Pointer
	}
}
type GmpRandstateT [1]X__gmpRandstateStruct
type MpzSrcptr *X__mpzStruct
type MpzPtr *X__mpzStruct
type MpfSrcptr *X__mpfStruct
type MpfPtr *X__mpfStruct
type MpqSrcptr *X__mpqStruct
type MpqPtr *X__mpqStruct
type GmpRandstatePtr *X__gmpRandstateStruct
type GmpRandstateSrcptr *X__gmpRandstateStruct

//go:linkname X__gmpSetMemoryFunctions C.__gmp_set_memory_functions
func X__gmpSetMemoryFunctions(func(uintptr) unsafe.Pointer, func(unsafe.Pointer, uintptr, uintptr) unsafe.Pointer, func(unsafe.Pointer, uintptr))

//go:linkname X__gmpGetMemoryFunctions C.__gmp_get_memory_functions
func X__gmpGetMemoryFunctions(func(uintptr) unsafe.Pointer, func(unsafe.Pointer, uintptr, uintptr) unsafe.Pointer, func(unsafe.Pointer, uintptr))

//go:linkname X__gmpRandinit C.__gmp_randinit
func X__gmpRandinit(__llgo_arg_0 GmpRandstatePtr, __llgo_arg_1 GmpRandalgT, __llgo_va_list ...interface{})

//go:linkname X__gmpRandinitDefault C.__gmp_randinit_default
func X__gmpRandinitDefault(GmpRandstatePtr)

//go:linkname X__gmpRandinitLc2exp C.__gmp_randinit_lc_2exp
func X__gmpRandinitLc2exp(GmpRandstatePtr, MpzSrcptr, c.Ulong, MpBitcntT)

//go:linkname X__gmpRandinitLc2expSize C.__gmp_randinit_lc_2exp_size
func X__gmpRandinitLc2expSize(GmpRandstatePtr, MpBitcntT) c.Int

//go:linkname X__gmpRandinitMt C.__gmp_randinit_mt
func X__gmpRandinitMt(GmpRandstatePtr)

//go:linkname X__gmpRandinitSet C.__gmp_randinit_set
func X__gmpRandinitSet(GmpRandstatePtr, GmpRandstateSrcptr)

//go:linkname X__gmpRandseed C.__gmp_randseed
func X__gmpRandseed(GmpRandstatePtr, MpzSrcptr)

//go:linkname X__gmpRandseedUi C.__gmp_randseed_ui
func X__gmpRandseedUi(GmpRandstatePtr, c.Ulong)

//go:linkname X__gmpRandclear C.__gmp_randclear
func X__gmpRandclear(GmpRandstatePtr)

//go:linkname X__gmpUrandombUi C.__gmp_urandomb_ui
func X__gmpUrandombUi(GmpRandstatePtr, c.Ulong) c.Ulong

//go:linkname X__gmpUrandommUi C.__gmp_urandomm_ui
func X__gmpUrandommUi(GmpRandstatePtr, c.Ulong) c.Ulong

//go:linkname X__gmpAsprintf C.__gmp_asprintf
func X__gmpAsprintf(__llgo_arg_0 **int8, __llgo_arg_1 *int8, __llgo_va_list ...interface{}) c.Int

//go:linkname X__gmpPrintf C.__gmp_printf
func X__gmpPrintf(__llgo_arg_0 *int8, __llgo_va_list ...interface{}) c.Int

//go:linkname X__gmpSnprintf C.__gmp_snprintf
func X__gmpSnprintf(__llgo_arg_0 *int8, __llgo_arg_1 uintptr, __llgo_arg_2 *int8, __llgo_va_list ...interface{}) c.Int

//go:linkname X__gmpSprintf C.__gmp_sprintf
func X__gmpSprintf(__llgo_arg_0 *int8, __llgo_arg_1 *int8, __llgo_va_list ...interface{}) c.Int

//go:linkname X__gmpScanf C.__gmp_scanf
func X__gmpScanf(__llgo_arg_0 *int8, __llgo_va_list ...interface{}) c.Int

//go:linkname X__gmpSscanf C.__gmp_sscanf
func X__gmpSscanf(__llgo_arg_0 *int8, __llgo_arg_1 *int8, __llgo_va_list ...interface{}) c.Int

//go:linkname X__gmpzRealloc C.__gmpz_realloc
func X__gmpzRealloc(MpzPtr, MpSizeT) unsafe.Pointer

//go:linkname X__gmpzAbs C.__gmpz_abs
func X__gmpzAbs(MpzPtr, MpzSrcptr)

//go:linkname X__gmpzAdd C.__gmpz_add
func X__gmpzAdd(MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzAddUi C.__gmpz_add_ui
func X__gmpzAddUi(MpzPtr, MpzSrcptr, c.Ulong)

//go:linkname X__gmpzAddmul C.__gmpz_addmul
func X__gmpzAddmul(MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzAddmulUi C.__gmpz_addmul_ui
func X__gmpzAddmulUi(MpzPtr, MpzSrcptr, c.Ulong)

//go:linkname X__gmpzAnd C.__gmpz_and
func X__gmpzAnd(MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzArrayInit C.__gmpz_array_init
func X__gmpzArrayInit(MpzPtr, MpSizeT, MpSizeT)

//go:linkname X__gmpzBinUi C.__gmpz_bin_ui
func X__gmpzBinUi(MpzPtr, MpzSrcptr, c.Ulong)

//go:linkname X__gmpzBinUiui C.__gmpz_bin_uiui
func X__gmpzBinUiui(MpzPtr, c.Ulong, c.Ulong)

//go:linkname X__gmpzCdivQ C.__gmpz_cdiv_q
func X__gmpzCdivQ(MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzCdivQ2exp C.__gmpz_cdiv_q_2exp
func X__gmpzCdivQ2exp(MpzPtr, MpzSrcptr, MpBitcntT)

//go:linkname X__gmpzCdivQUi C.__gmpz_cdiv_q_ui
func X__gmpzCdivQUi(MpzPtr, MpzSrcptr, c.Ulong) c.Ulong

//go:linkname X__gmpzCdivQr C.__gmpz_cdiv_qr
func X__gmpzCdivQr(MpzPtr, MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzCdivQrUi C.__gmpz_cdiv_qr_ui
func X__gmpzCdivQrUi(MpzPtr, MpzPtr, MpzSrcptr, c.Ulong) c.Ulong

//go:linkname X__gmpzCdivR C.__gmpz_cdiv_r
func X__gmpzCdivR(MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzCdivR2exp C.__gmpz_cdiv_r_2exp
func X__gmpzCdivR2exp(MpzPtr, MpzSrcptr, MpBitcntT)

//go:linkname X__gmpzCdivRUi C.__gmpz_cdiv_r_ui
func X__gmpzCdivRUi(MpzPtr, MpzSrcptr, c.Ulong) c.Ulong

//go:linkname X__gmpzCdivUi C.__gmpz_cdiv_ui
func X__gmpzCdivUi(MpzSrcptr, c.Ulong) c.Ulong

//go:linkname X__gmpzClear C.__gmpz_clear
func X__gmpzClear(MpzPtr)

//go:linkname X__gmpzClears C.__gmpz_clears
func X__gmpzClears(__llgo_arg_0 MpzPtr, __llgo_va_list ...interface{})

//go:linkname X__gmpzClrbit C.__gmpz_clrbit
func X__gmpzClrbit(MpzPtr, MpBitcntT)

//go:linkname X__gmpzCmp C.__gmpz_cmp
func X__gmpzCmp(MpzSrcptr, MpzSrcptr) c.Int

//go:linkname X__gmpzCmpD C.__gmpz_cmp_d
func X__gmpzCmpD(MpzSrcptr, float64) c.Int

//go:linkname X__gmpzCmpSi C.__gmpz_cmp_si
func X__gmpzCmpSi(MpzSrcptr, c.Long) c.Int

//go:linkname X__gmpzCmpUi C.__gmpz_cmp_ui
func X__gmpzCmpUi(MpzSrcptr, c.Ulong) c.Int

//go:linkname X__gmpzCmpabs C.__gmpz_cmpabs
func X__gmpzCmpabs(MpzSrcptr, MpzSrcptr) c.Int

//go:linkname X__gmpzCmpabsD C.__gmpz_cmpabs_d
func X__gmpzCmpabsD(MpzSrcptr, float64) c.Int

//go:linkname X__gmpzCmpabsUi C.__gmpz_cmpabs_ui
func X__gmpzCmpabsUi(MpzSrcptr, c.Ulong) c.Int

//go:linkname X__gmpzCom C.__gmpz_com
func X__gmpzCom(MpzPtr, MpzSrcptr)

//go:linkname X__gmpzCombit C.__gmpz_combit
func X__gmpzCombit(MpzPtr, MpBitcntT)

//go:linkname X__gmpzCongruentP C.__gmpz_congruent_p
func X__gmpzCongruentP(MpzSrcptr, MpzSrcptr, MpzSrcptr) c.Int

//go:linkname X__gmpzCongruent2expP C.__gmpz_congruent_2exp_p
func X__gmpzCongruent2expP(MpzSrcptr, MpzSrcptr, MpBitcntT) c.Int

//go:linkname X__gmpzCongruentUiP C.__gmpz_congruent_ui_p
func X__gmpzCongruentUiP(MpzSrcptr, c.Ulong, c.Ulong) c.Int

//go:linkname X__gmpzDivexact C.__gmpz_divexact
func X__gmpzDivexact(MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzDivexactUi C.__gmpz_divexact_ui
func X__gmpzDivexactUi(MpzPtr, MpzSrcptr, c.Ulong)

//go:linkname X__gmpzDivisibleP C.__gmpz_divisible_p
func X__gmpzDivisibleP(MpzSrcptr, MpzSrcptr) c.Int

//go:linkname X__gmpzDivisibleUiP C.__gmpz_divisible_ui_p
func X__gmpzDivisibleUiP(MpzSrcptr, c.Ulong) c.Int

//go:linkname X__gmpzDivisible2expP C.__gmpz_divisible_2exp_p
func X__gmpzDivisible2expP(MpzSrcptr, MpBitcntT) c.Int

//go:linkname X__gmpzDump C.__gmpz_dump
func X__gmpzDump(MpzSrcptr)

//go:linkname X__gmpzExport C.__gmpz_export
func X__gmpzExport(unsafe.Pointer, *uintptr, c.Int, uintptr, c.Int, uintptr, MpzSrcptr) unsafe.Pointer

//go:linkname X__gmpzFacUi C.__gmpz_fac_ui
func X__gmpzFacUi(MpzPtr, c.Ulong)

//go:linkname X__gmpz2facUi C.__gmpz_2fac_ui
func X__gmpz2facUi(MpzPtr, c.Ulong)

//go:linkname X__gmpzMfacUiui C.__gmpz_mfac_uiui
func X__gmpzMfacUiui(MpzPtr, c.Ulong, c.Ulong)

//go:linkname X__gmpzPrimorialUi C.__gmpz_primorial_ui
func X__gmpzPrimorialUi(MpzPtr, c.Ulong)

//go:linkname X__gmpzFdivQ C.__gmpz_fdiv_q
func X__gmpzFdivQ(MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzFdivQ2exp C.__gmpz_fdiv_q_2exp
func X__gmpzFdivQ2exp(MpzPtr, MpzSrcptr, MpBitcntT)

//go:linkname X__gmpzFdivQUi C.__gmpz_fdiv_q_ui
func X__gmpzFdivQUi(MpzPtr, MpzSrcptr, c.Ulong) c.Ulong

//go:linkname X__gmpzFdivQr C.__gmpz_fdiv_qr
func X__gmpzFdivQr(MpzPtr, MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzFdivQrUi C.__gmpz_fdiv_qr_ui
func X__gmpzFdivQrUi(MpzPtr, MpzPtr, MpzSrcptr, c.Ulong) c.Ulong

//go:linkname X__gmpzFdivR C.__gmpz_fdiv_r
func X__gmpzFdivR(MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzFdivR2exp C.__gmpz_fdiv_r_2exp
func X__gmpzFdivR2exp(MpzPtr, MpzSrcptr, MpBitcntT)

//go:linkname X__gmpzFdivRUi C.__gmpz_fdiv_r_ui
func X__gmpzFdivRUi(MpzPtr, MpzSrcptr, c.Ulong) c.Ulong

//go:linkname X__gmpzFdivUi C.__gmpz_fdiv_ui
func X__gmpzFdivUi(MpzSrcptr, c.Ulong) c.Ulong

//go:linkname X__gmpzFibUi C.__gmpz_fib_ui
func X__gmpzFibUi(MpzPtr, c.Ulong)

//go:linkname X__gmpzFib2Ui C.__gmpz_fib2_ui
func X__gmpzFib2Ui(MpzPtr, MpzPtr, c.Ulong)

//go:linkname X__gmpzFitsSintP C.__gmpz_fits_sint_p
func X__gmpzFitsSintP(MpzSrcptr) c.Int

//go:linkname X__gmpzFitsSlongP C.__gmpz_fits_slong_p
func X__gmpzFitsSlongP(MpzSrcptr) c.Int

//go:linkname X__gmpzFitsSshortP C.__gmpz_fits_sshort_p
func X__gmpzFitsSshortP(MpzSrcptr) c.Int

//go:linkname X__gmpzFitsUintP C.__gmpz_fits_uint_p
func X__gmpzFitsUintP(MpzSrcptr) c.Int

//go:linkname X__gmpzFitsUlongP C.__gmpz_fits_ulong_p
func X__gmpzFitsUlongP(MpzSrcptr) c.Int

//go:linkname X__gmpzFitsUshortP C.__gmpz_fits_ushort_p
func X__gmpzFitsUshortP(MpzSrcptr) c.Int

//go:linkname X__gmpzGcd C.__gmpz_gcd
func X__gmpzGcd(MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzGcdUi C.__gmpz_gcd_ui
func X__gmpzGcdUi(MpzPtr, MpzSrcptr, c.Ulong) c.Ulong

//go:linkname X__gmpzGcdext C.__gmpz_gcdext
func X__gmpzGcdext(MpzPtr, MpzPtr, MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzGetD C.__gmpz_get_d
func X__gmpzGetD(MpzSrcptr) float64

//go:linkname X__gmpzGetD2exp C.__gmpz_get_d_2exp
func X__gmpzGetD2exp(*c.Long, MpzSrcptr) float64

/* signed */
//go:linkname X__gmpzGetSi C.__gmpz_get_si
func X__gmpzGetSi(MpzSrcptr) c.Long

//go:linkname X__gmpzGetStr C.__gmpz_get_str
func X__gmpzGetStr(*int8, c.Int, MpzSrcptr) *int8

//go:linkname X__gmpzGetUi C.__gmpz_get_ui
func X__gmpzGetUi(MpzSrcptr) c.Ulong

//go:linkname X__gmpzGetlimbn C.__gmpz_getlimbn
func X__gmpzGetlimbn(MpzSrcptr, MpSizeT) MpLimbT

//go:linkname X__gmpzHamdist C.__gmpz_hamdist
func X__gmpzHamdist(MpzSrcptr, MpzSrcptr) MpBitcntT

//go:linkname X__gmpzImport C.__gmpz_import
func X__gmpzImport(MpzPtr, uintptr, c.Int, uintptr, c.Int, uintptr, unsafe.Pointer)

//go:linkname X__gmpzInit C.__gmpz_init
func X__gmpzInit(MpzPtr)

//go:linkname X__gmpzInit2 C.__gmpz_init2
func X__gmpzInit2(MpzPtr, MpBitcntT)

//go:linkname X__gmpzInits C.__gmpz_inits
func X__gmpzInits(__llgo_arg_0 MpzPtr, __llgo_va_list ...interface{})

//go:linkname X__gmpzInitSet C.__gmpz_init_set
func X__gmpzInitSet(MpzPtr, MpzSrcptr)

//go:linkname X__gmpzInitSetD C.__gmpz_init_set_d
func X__gmpzInitSetD(MpzPtr, float64)

//go:linkname X__gmpzInitSetSi C.__gmpz_init_set_si
func X__gmpzInitSetSi(MpzPtr, c.Long)

//go:linkname X__gmpzInitSetStr C.__gmpz_init_set_str
func X__gmpzInitSetStr(MpzPtr, *int8, c.Int) c.Int

//go:linkname X__gmpzInitSetUi C.__gmpz_init_set_ui
func X__gmpzInitSetUi(MpzPtr, c.Ulong)

//go:linkname X__gmpzInvert C.__gmpz_invert
func X__gmpzInvert(MpzPtr, MpzSrcptr, MpzSrcptr) c.Int

//go:linkname X__gmpzIor C.__gmpz_ior
func X__gmpzIor(MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzJacobi C.__gmpz_jacobi
func X__gmpzJacobi(MpzSrcptr, MpzSrcptr) c.Int

//go:linkname X__gmpzKroneckerSi C.__gmpz_kronecker_si
func X__gmpzKroneckerSi(MpzSrcptr, c.Long) c.Int

//go:linkname X__gmpzKroneckerUi C.__gmpz_kronecker_ui
func X__gmpzKroneckerUi(MpzSrcptr, c.Ulong) c.Int

//go:linkname X__gmpzSiKronecker C.__gmpz_si_kronecker
func X__gmpzSiKronecker(c.Long, MpzSrcptr) c.Int

//go:linkname X__gmpzUiKronecker C.__gmpz_ui_kronecker
func X__gmpzUiKronecker(c.Ulong, MpzSrcptr) c.Int

//go:linkname X__gmpzLcm C.__gmpz_lcm
func X__gmpzLcm(MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzLcmUi C.__gmpz_lcm_ui
func X__gmpzLcmUi(MpzPtr, MpzSrcptr, c.Ulong)

//go:linkname X__gmpzLucnumUi C.__gmpz_lucnum_ui
func X__gmpzLucnumUi(MpzPtr, c.Ulong)

//go:linkname X__gmpzLucnum2Ui C.__gmpz_lucnum2_ui
func X__gmpzLucnum2Ui(MpzPtr, MpzPtr, c.Ulong)

//go:linkname X__gmpzMillerrabin C.__gmpz_millerrabin
func X__gmpzMillerrabin(MpzSrcptr, c.Int) c.Int

//go:linkname X__gmpzMod C.__gmpz_mod
func X__gmpzMod(MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzMul C.__gmpz_mul
func X__gmpzMul(MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzMul2exp C.__gmpz_mul_2exp
func X__gmpzMul2exp(MpzPtr, MpzSrcptr, MpBitcntT)

//go:linkname X__gmpzMulSi C.__gmpz_mul_si
func X__gmpzMulSi(MpzPtr, MpzSrcptr, c.Long)

//go:linkname X__gmpzMulUi C.__gmpz_mul_ui
func X__gmpzMulUi(MpzPtr, MpzSrcptr, c.Ulong)

//go:linkname X__gmpzNeg C.__gmpz_neg
func X__gmpzNeg(MpzPtr, MpzSrcptr)

//go:linkname X__gmpzNextprime C.__gmpz_nextprime
func X__gmpzNextprime(MpzPtr, MpzSrcptr)

//go:linkname X__gmpzPrevprime C.__gmpz_prevprime
func X__gmpzPrevprime(MpzPtr, MpzSrcptr) c.Int

//go:linkname X__gmpzPerfectPowerP C.__gmpz_perfect_power_p
func X__gmpzPerfectPowerP(MpzSrcptr) c.Int

//go:linkname X__gmpzPerfectSquareP C.__gmpz_perfect_square_p
func X__gmpzPerfectSquareP(MpzSrcptr) c.Int

//go:linkname X__gmpzPopcount C.__gmpz_popcount
func X__gmpzPopcount(MpzSrcptr) MpBitcntT

//go:linkname X__gmpzPowUi C.__gmpz_pow_ui
func X__gmpzPowUi(MpzPtr, MpzSrcptr, c.Ulong)

//go:linkname X__gmpzPowm C.__gmpz_powm
func X__gmpzPowm(MpzPtr, MpzSrcptr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzPowmSec C.__gmpz_powm_sec
func X__gmpzPowmSec(MpzPtr, MpzSrcptr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzPowmUi C.__gmpz_powm_ui
func X__gmpzPowmUi(MpzPtr, MpzSrcptr, c.Ulong, MpzSrcptr)

//go:linkname X__gmpzProbabPrimeP C.__gmpz_probab_prime_p
func X__gmpzProbabPrimeP(MpzSrcptr, c.Int) c.Int

//go:linkname X__gmpzRandom C.__gmpz_random
func X__gmpzRandom(MpzPtr, MpSizeT)

//go:linkname X__gmpzRandom2 C.__gmpz_random2
func X__gmpzRandom2(MpzPtr, MpSizeT)

//go:linkname X__gmpzRealloc2 C.__gmpz_realloc2
func X__gmpzRealloc2(MpzPtr, MpBitcntT)

//go:linkname X__gmpzRemove C.__gmpz_remove
func X__gmpzRemove(MpzPtr, MpzSrcptr, MpzSrcptr) MpBitcntT

//go:linkname X__gmpzRoot C.__gmpz_root
func X__gmpzRoot(MpzPtr, MpzSrcptr, c.Ulong) c.Int

//go:linkname X__gmpzRootrem C.__gmpz_rootrem
func X__gmpzRootrem(MpzPtr, MpzPtr, MpzSrcptr, c.Ulong)

//go:linkname X__gmpzRrandomb C.__gmpz_rrandomb
func X__gmpzRrandomb(MpzPtr, GmpRandstatePtr, MpBitcntT)

//go:linkname X__gmpzScan0 C.__gmpz_scan0
func X__gmpzScan0(MpzSrcptr, MpBitcntT) MpBitcntT

//go:linkname X__gmpzScan1 C.__gmpz_scan1
func X__gmpzScan1(MpzSrcptr, MpBitcntT) MpBitcntT

//go:linkname X__gmpzSet C.__gmpz_set
func X__gmpzSet(MpzPtr, MpzSrcptr)

//go:linkname X__gmpzSetD C.__gmpz_set_d
func X__gmpzSetD(MpzPtr, float64)

//go:linkname X__gmpzSetF C.__gmpz_set_f
func X__gmpzSetF(MpzPtr, MpfSrcptr)

//go:linkname X__gmpzSetQ C.__gmpz_set_q
func X__gmpzSetQ(MpzPtr, MpqSrcptr)

//go:linkname X__gmpzSetSi C.__gmpz_set_si
func X__gmpzSetSi(MpzPtr, c.Long)

//go:linkname X__gmpzSetStr C.__gmpz_set_str
func X__gmpzSetStr(MpzPtr, *int8, c.Int) c.Int

//go:linkname X__gmpzSetUi C.__gmpz_set_ui
func X__gmpzSetUi(MpzPtr, c.Ulong)

//go:linkname X__gmpzSetbit C.__gmpz_setbit
func X__gmpzSetbit(MpzPtr, MpBitcntT)

//go:linkname X__gmpzSize C.__gmpz_size
func X__gmpzSize(MpzSrcptr) uintptr

//go:linkname X__gmpzSizeinbase C.__gmpz_sizeinbase
func X__gmpzSizeinbase(MpzSrcptr, c.Int) uintptr

//go:linkname X__gmpzSqrt C.__gmpz_sqrt
func X__gmpzSqrt(MpzPtr, MpzSrcptr)

//go:linkname X__gmpzSqrtrem C.__gmpz_sqrtrem
func X__gmpzSqrtrem(MpzPtr, MpzPtr, MpzSrcptr)

//go:linkname X__gmpzSub C.__gmpz_sub
func X__gmpzSub(MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzSubUi C.__gmpz_sub_ui
func X__gmpzSubUi(MpzPtr, MpzSrcptr, c.Ulong)

//go:linkname X__gmpzUiSub C.__gmpz_ui_sub
func X__gmpzUiSub(MpzPtr, c.Ulong, MpzSrcptr)

//go:linkname X__gmpzSubmul C.__gmpz_submul
func X__gmpzSubmul(MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzSubmulUi C.__gmpz_submul_ui
func X__gmpzSubmulUi(MpzPtr, MpzSrcptr, c.Ulong)

//go:linkname X__gmpzSwap C.__gmpz_swap
func X__gmpzSwap(MpzPtr, MpzPtr)

//go:linkname X__gmpzTdivUi C.__gmpz_tdiv_ui
func X__gmpzTdivUi(MpzSrcptr, c.Ulong) c.Ulong

//go:linkname X__gmpzTdivQ C.__gmpz_tdiv_q
func X__gmpzTdivQ(MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzTdivQ2exp C.__gmpz_tdiv_q_2exp
func X__gmpzTdivQ2exp(MpzPtr, MpzSrcptr, MpBitcntT)

//go:linkname X__gmpzTdivQUi C.__gmpz_tdiv_q_ui
func X__gmpzTdivQUi(MpzPtr, MpzSrcptr, c.Ulong) c.Ulong

//go:linkname X__gmpzTdivQr C.__gmpz_tdiv_qr
func X__gmpzTdivQr(MpzPtr, MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzTdivQrUi C.__gmpz_tdiv_qr_ui
func X__gmpzTdivQrUi(MpzPtr, MpzPtr, MpzSrcptr, c.Ulong) c.Ulong

//go:linkname X__gmpzTdivR C.__gmpz_tdiv_r
func X__gmpzTdivR(MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzTdivR2exp C.__gmpz_tdiv_r_2exp
func X__gmpzTdivR2exp(MpzPtr, MpzSrcptr, MpBitcntT)

//go:linkname X__gmpzTdivRUi C.__gmpz_tdiv_r_ui
func X__gmpzTdivRUi(MpzPtr, MpzSrcptr, c.Ulong) c.Ulong

//go:linkname X__gmpzTstbit C.__gmpz_tstbit
func X__gmpzTstbit(MpzSrcptr, MpBitcntT) c.Int

//go:linkname X__gmpzUiPowUi C.__gmpz_ui_pow_ui
func X__gmpzUiPowUi(MpzPtr, c.Ulong, c.Ulong)

//go:linkname X__gmpzUrandomb C.__gmpz_urandomb
func X__gmpzUrandomb(MpzPtr, GmpRandstatePtr, MpBitcntT)

//go:linkname X__gmpzUrandomm C.__gmpz_urandomm
func X__gmpzUrandomm(MpzPtr, GmpRandstatePtr, MpzSrcptr)

//go:linkname X__gmpzXor C.__gmpz_xor
func X__gmpzXor(MpzPtr, MpzSrcptr, MpzSrcptr)

//go:linkname X__gmpzLimbsRead C.__gmpz_limbs_read
func X__gmpzLimbsRead(MpzSrcptr) MpSrcptr

//go:linkname X__gmpzLimbsWrite C.__gmpz_limbs_write
func X__gmpzLimbsWrite(MpzPtr, MpSizeT) MpPtr

//go:linkname X__gmpzLimbsModify C.__gmpz_limbs_modify
func X__gmpzLimbsModify(MpzPtr, MpSizeT) MpPtr

//go:linkname X__gmpzLimbsFinish C.__gmpz_limbs_finish
func X__gmpzLimbsFinish(MpzPtr, MpSizeT)

//go:linkname X__gmpzRoinitN C.__gmpz_roinit_n
func X__gmpzRoinitN(MpzPtr, MpSrcptr, MpSizeT) MpzSrcptr

//go:linkname X__gmpqAbs C.__gmpq_abs
func X__gmpqAbs(MpqPtr, MpqSrcptr)

//go:linkname X__gmpqAdd C.__gmpq_add
func X__gmpqAdd(MpqPtr, MpqSrcptr, MpqSrcptr)

//go:linkname X__gmpqCanonicalize C.__gmpq_canonicalize
func X__gmpqCanonicalize(MpqPtr)

//go:linkname X__gmpqClear C.__gmpq_clear
func X__gmpqClear(MpqPtr)

//go:linkname X__gmpqClears C.__gmpq_clears
func X__gmpqClears(__llgo_arg_0 MpqPtr, __llgo_va_list ...interface{})

//go:linkname X__gmpqCmp C.__gmpq_cmp
func X__gmpqCmp(MpqSrcptr, MpqSrcptr) c.Int

//go:linkname X__gmpqCmpSi C.__gmpq_cmp_si
func X__gmpqCmpSi(MpqSrcptr, c.Long, c.Ulong) c.Int

//go:linkname X__gmpqCmpUi C.__gmpq_cmp_ui
func X__gmpqCmpUi(MpqSrcptr, c.Ulong, c.Ulong) c.Int

//go:linkname X__gmpqCmpZ C.__gmpq_cmp_z
func X__gmpqCmpZ(MpqSrcptr, MpzSrcptr) c.Int

//go:linkname X__gmpqDiv C.__gmpq_div
func X__gmpqDiv(MpqPtr, MpqSrcptr, MpqSrcptr)

//go:linkname X__gmpqDiv2exp C.__gmpq_div_2exp
func X__gmpqDiv2exp(MpqPtr, MpqSrcptr, MpBitcntT)

//go:linkname X__gmpqEqual C.__gmpq_equal
func X__gmpqEqual(MpqSrcptr, MpqSrcptr) c.Int

//go:linkname X__gmpqGetNum C.__gmpq_get_num
func X__gmpqGetNum(MpzPtr, MpqSrcptr)

//go:linkname X__gmpqGetDen C.__gmpq_get_den
func X__gmpqGetDen(MpzPtr, MpqSrcptr)

//go:linkname X__gmpqGetD C.__gmpq_get_d
func X__gmpqGetD(MpqSrcptr) float64

//go:linkname X__gmpqGetStr C.__gmpq_get_str
func X__gmpqGetStr(*int8, c.Int, MpqSrcptr) *int8

//go:linkname X__gmpqInit C.__gmpq_init
func X__gmpqInit(MpqPtr)

//go:linkname X__gmpqInits C.__gmpq_inits
func X__gmpqInits(__llgo_arg_0 MpqPtr, __llgo_va_list ...interface{})

//go:linkname X__gmpqInv C.__gmpq_inv
func X__gmpqInv(MpqPtr, MpqSrcptr)

//go:linkname X__gmpqMul C.__gmpq_mul
func X__gmpqMul(MpqPtr, MpqSrcptr, MpqSrcptr)

//go:linkname X__gmpqMul2exp C.__gmpq_mul_2exp
func X__gmpqMul2exp(MpqPtr, MpqSrcptr, MpBitcntT)

//go:linkname X__gmpqNeg C.__gmpq_neg
func X__gmpqNeg(MpqPtr, MpqSrcptr)

//go:linkname X__gmpqSet C.__gmpq_set
func X__gmpqSet(MpqPtr, MpqSrcptr)

//go:linkname X__gmpqSetD C.__gmpq_set_d
func X__gmpqSetD(MpqPtr, float64)

//go:linkname X__gmpqSetDen C.__gmpq_set_den
func X__gmpqSetDen(MpqPtr, MpzSrcptr)

//go:linkname X__gmpqSetF C.__gmpq_set_f
func X__gmpqSetF(MpqPtr, MpfSrcptr)

//go:linkname X__gmpqSetNum C.__gmpq_set_num
func X__gmpqSetNum(MpqPtr, MpzSrcptr)

//go:linkname X__gmpqSetSi C.__gmpq_set_si
func X__gmpqSetSi(MpqPtr, c.Long, c.Ulong)

//go:linkname X__gmpqSetStr C.__gmpq_set_str
func X__gmpqSetStr(MpqPtr, *int8, c.Int) c.Int

//go:linkname X__gmpqSetUi C.__gmpq_set_ui
func X__gmpqSetUi(MpqPtr, c.Ulong, c.Ulong)

//go:linkname X__gmpqSetZ C.__gmpq_set_z
func X__gmpqSetZ(MpqPtr, MpzSrcptr)

//go:linkname X__gmpqSub C.__gmpq_sub
func X__gmpqSub(MpqPtr, MpqSrcptr, MpqSrcptr)

//go:linkname X__gmpqSwap C.__gmpq_swap
func X__gmpqSwap(MpqPtr, MpqPtr)

//go:linkname X__gmpfAbs C.__gmpf_abs
func X__gmpfAbs(MpfPtr, MpfSrcptr)

//go:linkname X__gmpfAdd C.__gmpf_add
func X__gmpfAdd(MpfPtr, MpfSrcptr, MpfSrcptr)

//go:linkname X__gmpfAddUi C.__gmpf_add_ui
func X__gmpfAddUi(MpfPtr, MpfSrcptr, c.Ulong)

//go:linkname X__gmpfCeil C.__gmpf_ceil
func X__gmpfCeil(MpfPtr, MpfSrcptr)

//go:linkname X__gmpfClear C.__gmpf_clear
func X__gmpfClear(MpfPtr)

//go:linkname X__gmpfClears C.__gmpf_clears
func X__gmpfClears(__llgo_arg_0 MpfPtr, __llgo_va_list ...interface{})

//go:linkname X__gmpfCmp C.__gmpf_cmp
func X__gmpfCmp(MpfSrcptr, MpfSrcptr) c.Int

//go:linkname X__gmpfCmpZ C.__gmpf_cmp_z
func X__gmpfCmpZ(MpfSrcptr, MpzSrcptr) c.Int

//go:linkname X__gmpfCmpD C.__gmpf_cmp_d
func X__gmpfCmpD(MpfSrcptr, float64) c.Int

//go:linkname X__gmpfCmpSi C.__gmpf_cmp_si
func X__gmpfCmpSi(MpfSrcptr, c.Long) c.Int

//go:linkname X__gmpfCmpUi C.__gmpf_cmp_ui
func X__gmpfCmpUi(MpfSrcptr, c.Ulong) c.Int

//go:linkname X__gmpfDiv C.__gmpf_div
func X__gmpfDiv(MpfPtr, MpfSrcptr, MpfSrcptr)

//go:linkname X__gmpfDiv2exp C.__gmpf_div_2exp
func X__gmpfDiv2exp(MpfPtr, MpfSrcptr, MpBitcntT)

//go:linkname X__gmpfDivUi C.__gmpf_div_ui
func X__gmpfDivUi(MpfPtr, MpfSrcptr, c.Ulong)

//go:linkname X__gmpfDump C.__gmpf_dump
func X__gmpfDump(MpfSrcptr)

//go:linkname X__gmpfEq C.__gmpf_eq
func X__gmpfEq(MpfSrcptr, MpfSrcptr, MpBitcntT) c.Int

//go:linkname X__gmpfFitsSintP C.__gmpf_fits_sint_p
func X__gmpfFitsSintP(MpfSrcptr) c.Int

//go:linkname X__gmpfFitsSlongP C.__gmpf_fits_slong_p
func X__gmpfFitsSlongP(MpfSrcptr) c.Int

//go:linkname X__gmpfFitsSshortP C.__gmpf_fits_sshort_p
func X__gmpfFitsSshortP(MpfSrcptr) c.Int

//go:linkname X__gmpfFitsUintP C.__gmpf_fits_uint_p
func X__gmpfFitsUintP(MpfSrcptr) c.Int

//go:linkname X__gmpfFitsUlongP C.__gmpf_fits_ulong_p
func X__gmpfFitsUlongP(MpfSrcptr) c.Int

//go:linkname X__gmpfFitsUshortP C.__gmpf_fits_ushort_p
func X__gmpfFitsUshortP(MpfSrcptr) c.Int

//go:linkname X__gmpfFloor C.__gmpf_floor
func X__gmpfFloor(MpfPtr, MpfSrcptr)

//go:linkname X__gmpfGetD C.__gmpf_get_d
func X__gmpfGetD(MpfSrcptr) float64

//go:linkname X__gmpfGetD2exp C.__gmpf_get_d_2exp
func X__gmpfGetD2exp(*c.Long, MpfSrcptr) float64

//go:linkname X__gmpfGetDefaultPrec C.__gmpf_get_default_prec
func X__gmpfGetDefaultPrec() MpBitcntT

//go:linkname X__gmpfGetPrec C.__gmpf_get_prec
func X__gmpfGetPrec(MpfSrcptr) MpBitcntT

//go:linkname X__gmpfGetSi C.__gmpf_get_si
func X__gmpfGetSi(MpfSrcptr) c.Long

//go:linkname X__gmpfGetStr C.__gmpf_get_str
func X__gmpfGetStr(*int8, *MpExpT, c.Int, uintptr, MpfSrcptr) *int8

//go:linkname X__gmpfGetUi C.__gmpf_get_ui
func X__gmpfGetUi(MpfSrcptr) c.Ulong

//go:linkname X__gmpfInit C.__gmpf_init
func X__gmpfInit(MpfPtr)

//go:linkname X__gmpfInit2 C.__gmpf_init2
func X__gmpfInit2(MpfPtr, MpBitcntT)

//go:linkname X__gmpfInits C.__gmpf_inits
func X__gmpfInits(__llgo_arg_0 MpfPtr, __llgo_va_list ...interface{})

//go:linkname X__gmpfInitSet C.__gmpf_init_set
func X__gmpfInitSet(MpfPtr, MpfSrcptr)

//go:linkname X__gmpfInitSetD C.__gmpf_init_set_d
func X__gmpfInitSetD(MpfPtr, float64)

//go:linkname X__gmpfInitSetSi C.__gmpf_init_set_si
func X__gmpfInitSetSi(MpfPtr, c.Long)

//go:linkname X__gmpfInitSetStr C.__gmpf_init_set_str
func X__gmpfInitSetStr(MpfPtr, *int8, c.Int) c.Int

//go:linkname X__gmpfInitSetUi C.__gmpf_init_set_ui
func X__gmpfInitSetUi(MpfPtr, c.Ulong)

//go:linkname X__gmpfIntegerP C.__gmpf_integer_p
func X__gmpfIntegerP(MpfSrcptr) c.Int

//go:linkname X__gmpfMul C.__gmpf_mul
func X__gmpfMul(MpfPtr, MpfSrcptr, MpfSrcptr)

//go:linkname X__gmpfMul2exp C.__gmpf_mul_2exp
func X__gmpfMul2exp(MpfPtr, MpfSrcptr, MpBitcntT)

//go:linkname X__gmpfMulUi C.__gmpf_mul_ui
func X__gmpfMulUi(MpfPtr, MpfSrcptr, c.Ulong)

//go:linkname X__gmpfNeg C.__gmpf_neg
func X__gmpfNeg(MpfPtr, MpfSrcptr)

//go:linkname X__gmpfPowUi C.__gmpf_pow_ui
func X__gmpfPowUi(MpfPtr, MpfSrcptr, c.Ulong)

//go:linkname X__gmpfRandom2 C.__gmpf_random2
func X__gmpfRandom2(MpfPtr, MpSizeT, MpExpT)

//go:linkname X__gmpfReldiff C.__gmpf_reldiff
func X__gmpfReldiff(MpfPtr, MpfSrcptr, MpfSrcptr)

//go:linkname X__gmpfSet C.__gmpf_set
func X__gmpfSet(MpfPtr, MpfSrcptr)

//go:linkname X__gmpfSetD C.__gmpf_set_d
func X__gmpfSetD(MpfPtr, float64)

// llgo:link MpBitcntT.X__gmpfSetDefaultPrec C.__gmpf_set_default_prec
func (recv_ MpBitcntT) X__gmpfSetDefaultPrec() {
}

//go:linkname X__gmpfSetPrec C.__gmpf_set_prec
func X__gmpfSetPrec(MpfPtr, MpBitcntT)

//go:linkname X__gmpfSetPrecRaw C.__gmpf_set_prec_raw
func X__gmpfSetPrecRaw(MpfPtr, MpBitcntT)

//go:linkname X__gmpfSetQ C.__gmpf_set_q
func X__gmpfSetQ(MpfPtr, MpqSrcptr)

//go:linkname X__gmpfSetSi C.__gmpf_set_si
func X__gmpfSetSi(MpfPtr, c.Long)

//go:linkname X__gmpfSetStr C.__gmpf_set_str
func X__gmpfSetStr(MpfPtr, *int8, c.Int) c.Int

//go:linkname X__gmpfSetUi C.__gmpf_set_ui
func X__gmpfSetUi(MpfPtr, c.Ulong)

//go:linkname X__gmpfSetZ C.__gmpf_set_z
func X__gmpfSetZ(MpfPtr, MpzSrcptr)

//go:linkname X__gmpfSize C.__gmpf_size
func X__gmpfSize(MpfSrcptr) uintptr

//go:linkname X__gmpfSqrt C.__gmpf_sqrt
func X__gmpfSqrt(MpfPtr, MpfSrcptr)

//go:linkname X__gmpfSqrtUi C.__gmpf_sqrt_ui
func X__gmpfSqrtUi(MpfPtr, c.Ulong)

//go:linkname X__gmpfSub C.__gmpf_sub
func X__gmpfSub(MpfPtr, MpfSrcptr, MpfSrcptr)

//go:linkname X__gmpfSubUi C.__gmpf_sub_ui
func X__gmpfSubUi(MpfPtr, MpfSrcptr, c.Ulong)

//go:linkname X__gmpfSwap C.__gmpf_swap
func X__gmpfSwap(MpfPtr, MpfPtr)

//go:linkname X__gmpfTrunc C.__gmpf_trunc
func X__gmpfTrunc(MpfPtr, MpfSrcptr)

//go:linkname X__gmpfUiDiv C.__gmpf_ui_div
func X__gmpfUiDiv(MpfPtr, c.Ulong, MpfSrcptr)

//go:linkname X__gmpfUiSub C.__gmpf_ui_sub
func X__gmpfUiSub(MpfPtr, c.Ulong, MpfSrcptr)

//go:linkname X__gmpfUrandomb C.__gmpf_urandomb
func X__gmpfUrandomb(MpfPtr, GmpRandstatePtr, MpBitcntT)

//go:linkname X__gmpnAdd C.__gmpn_add
func X__gmpnAdd(MpPtr, MpSrcptr, MpSizeT, MpSrcptr, MpSizeT) MpLimbT

//go:linkname X__gmpnAdd1 C.__gmpn_add_1
func X__gmpnAdd1(MpPtr, MpSrcptr, MpSizeT, MpLimbT) MpLimbT

//go:linkname X__gmpnAddN C.__gmpn_add_n
func X__gmpnAddN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT) MpLimbT

//go:linkname X__gmpnAddmul1 C.__gmpn_addmul_1
func X__gmpnAddmul1(MpPtr, MpSrcptr, MpSizeT, MpLimbT) MpLimbT

//go:linkname X__gmpnCmp C.__gmpn_cmp
func X__gmpnCmp(MpSrcptr, MpSrcptr, MpSizeT) c.Int

//go:linkname X__gmpnZeroP C.__gmpn_zero_p
func X__gmpnZeroP(MpSrcptr, MpSizeT) c.Int

//go:linkname X__gmpnDivexact1 C.__gmpn_divexact_1
func X__gmpnDivexact1(MpPtr, MpSrcptr, MpSizeT, MpLimbT)

//go:linkname X__gmpnDivexactBy3c C.__gmpn_divexact_by3c
func X__gmpnDivexactBy3c(MpPtr, MpSrcptr, MpSizeT, MpLimbT) MpLimbT

//go:linkname X__gmpnDivrem C.__gmpn_divrem
func X__gmpnDivrem(MpPtr, MpSizeT, MpPtr, MpSizeT, MpSrcptr, MpSizeT) MpLimbT

//go:linkname X__gmpnDivrem1 C.__gmpn_divrem_1
func X__gmpnDivrem1(MpPtr, MpSizeT, MpSrcptr, MpSizeT, MpLimbT) MpLimbT

//go:linkname X__gmpnDivrem2 C.__gmpn_divrem_2
func X__gmpnDivrem2(MpPtr, MpSizeT, MpPtr, MpSizeT, MpSrcptr) MpLimbT

//go:linkname X__gmpnDivQr1 C.__gmpn_div_qr_1
func X__gmpnDivQr1(MpPtr, *MpLimbT, MpSrcptr, MpSizeT, MpLimbT) MpLimbT

//go:linkname X__gmpnDivQr2 C.__gmpn_div_qr_2
func X__gmpnDivQr2(MpPtr, MpPtr, MpSrcptr, MpSizeT, MpSrcptr) MpLimbT

//go:linkname X__gmpnGcd C.__gmpn_gcd
func X__gmpnGcd(MpPtr, MpPtr, MpSizeT, MpPtr, MpSizeT) MpSizeT

// llgo:link MpLimbT.X__gmpnGcd11 C.__gmpn_gcd_11
func (recv_ MpLimbT) X__gmpnGcd11(MpLimbT) MpLimbT {
	return 0
}

//go:linkname X__gmpnGcd1 C.__gmpn_gcd_1
func X__gmpnGcd1(MpSrcptr, MpSizeT, MpLimbT) MpLimbT

// llgo:link (*MpLimbSignedT).X__gmpnGcdext1 C.__gmpn_gcdext_1
func (recv_ *MpLimbSignedT) X__gmpnGcdext1(*MpLimbSignedT, MpLimbT, MpLimbT) MpLimbT {
	return 0
}

//go:linkname X__gmpnGcdext C.__gmpn_gcdext
func X__gmpnGcdext(MpPtr, MpPtr, *MpSizeT, MpPtr, MpSizeT, MpPtr, MpSizeT) MpSizeT

//go:linkname X__gmpnGetStr C.__gmpn_get_str
func X__gmpnGetStr(*int8, c.Int, MpPtr, MpSizeT) uintptr

//go:linkname X__gmpnHamdist C.__gmpn_hamdist
func X__gmpnHamdist(MpSrcptr, MpSrcptr, MpSizeT) MpBitcntT

//go:linkname X__gmpnLshift C.__gmpn_lshift
func X__gmpnLshift(MpPtr, MpSrcptr, MpSizeT, c.Uint) MpLimbT

//go:linkname X__gmpnMod1 C.__gmpn_mod_1
func X__gmpnMod1(MpSrcptr, MpSizeT, MpLimbT) MpLimbT

//go:linkname X__gmpnMul C.__gmpn_mul
func X__gmpnMul(MpPtr, MpSrcptr, MpSizeT, MpSrcptr, MpSizeT) MpLimbT

//go:linkname X__gmpnMul1 C.__gmpn_mul_1
func X__gmpnMul1(MpPtr, MpSrcptr, MpSizeT, MpLimbT) MpLimbT

//go:linkname X__gmpnMulN C.__gmpn_mul_n
func X__gmpnMulN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT)

//go:linkname X__gmpnSqr C.__gmpn_sqr
func X__gmpnSqr(MpPtr, MpSrcptr, MpSizeT)

//go:linkname X__gmpnNeg C.__gmpn_neg
func X__gmpnNeg(MpPtr, MpSrcptr, MpSizeT) MpLimbT

//go:linkname X__gmpnCom C.__gmpn_com
func X__gmpnCom(MpPtr, MpSrcptr, MpSizeT)

//go:linkname X__gmpnPerfectSquareP C.__gmpn_perfect_square_p
func X__gmpnPerfectSquareP(MpSrcptr, MpSizeT) c.Int

//go:linkname X__gmpnPerfectPowerP C.__gmpn_perfect_power_p
func X__gmpnPerfectPowerP(MpSrcptr, MpSizeT) c.Int

//go:linkname X__gmpnPopcount C.__gmpn_popcount
func X__gmpnPopcount(MpSrcptr, MpSizeT) MpBitcntT

//go:linkname X__gmpnPow1 C.__gmpn_pow_1
func X__gmpnPow1(MpPtr, MpSrcptr, MpSizeT, MpLimbT, MpPtr) MpSizeT

//go:linkname X__gmpnPreinvMod1 C.__gmpn_preinv_mod_1
func X__gmpnPreinvMod1(MpSrcptr, MpSizeT, MpLimbT, MpLimbT) MpLimbT

//go:linkname X__gmpnRandom C.__gmpn_random
func X__gmpnRandom(MpPtr, MpSizeT)

//go:linkname X__gmpnRandom2 C.__gmpn_random2
func X__gmpnRandom2(MpPtr, MpSizeT)

//go:linkname X__gmpnRshift C.__gmpn_rshift
func X__gmpnRshift(MpPtr, MpSrcptr, MpSizeT, c.Uint) MpLimbT

//go:linkname X__gmpnScan0 C.__gmpn_scan0
func X__gmpnScan0(MpSrcptr, MpBitcntT) MpBitcntT

//go:linkname X__gmpnScan1 C.__gmpn_scan1
func X__gmpnScan1(MpSrcptr, MpBitcntT) MpBitcntT

//go:linkname X__gmpnSetStr C.__gmpn_set_str
func X__gmpnSetStr(MpPtr, *int8, uintptr, c.Int) MpSizeT

//go:linkname X__gmpnSizeinbase C.__gmpn_sizeinbase
func X__gmpnSizeinbase(MpSrcptr, MpSizeT, c.Int) uintptr

//go:linkname X__gmpnSqrtrem C.__gmpn_sqrtrem
func X__gmpnSqrtrem(MpPtr, MpPtr, MpSrcptr, MpSizeT) MpSizeT

//go:linkname X__gmpnSub C.__gmpn_sub
func X__gmpnSub(MpPtr, MpSrcptr, MpSizeT, MpSrcptr, MpSizeT) MpLimbT

//go:linkname X__gmpnSub1 C.__gmpn_sub_1
func X__gmpnSub1(MpPtr, MpSrcptr, MpSizeT, MpLimbT) MpLimbT

//go:linkname X__gmpnSubN C.__gmpn_sub_n
func X__gmpnSubN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT) MpLimbT

//go:linkname X__gmpnSubmul1 C.__gmpn_submul_1
func X__gmpnSubmul1(MpPtr, MpSrcptr, MpSizeT, MpLimbT) MpLimbT

//go:linkname X__gmpnTdivQr C.__gmpn_tdiv_qr
func X__gmpnTdivQr(MpPtr, MpPtr, MpSizeT, MpSrcptr, MpSizeT, MpSrcptr, MpSizeT)

//go:linkname X__gmpnAndN C.__gmpn_and_n
func X__gmpnAndN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT)

//go:linkname X__gmpnAndnN C.__gmpn_andn_n
func X__gmpnAndnN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT)

//go:linkname X__gmpnNandN C.__gmpn_nand_n
func X__gmpnNandN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT)

//go:linkname X__gmpnIorN C.__gmpn_ior_n
func X__gmpnIorN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT)

//go:linkname X__gmpnIornN C.__gmpn_iorn_n
func X__gmpnIornN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT)

//go:linkname X__gmpnNiorN C.__gmpn_nior_n
func X__gmpnNiorN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT)

//go:linkname X__gmpnXorN C.__gmpn_xor_n
func X__gmpnXorN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT)

//go:linkname X__gmpnXnorN C.__gmpn_xnor_n
func X__gmpnXnorN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT)

//go:linkname X__gmpnCopyi C.__gmpn_copyi
func X__gmpnCopyi(MpPtr, MpSrcptr, MpSizeT)

//go:linkname X__gmpnCopyd C.__gmpn_copyd
func X__gmpnCopyd(MpPtr, MpSrcptr, MpSizeT)

//go:linkname X__gmpnZero C.__gmpn_zero
func X__gmpnZero(MpPtr, MpSizeT)

// llgo:link MpLimbT.X__gmpnCndAddN C.__gmpn_cnd_add_n
func (recv_ MpLimbT) X__gmpnCndAddN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT) MpLimbT {
	return 0
}

// llgo:link MpLimbT.X__gmpnCndSubN C.__gmpn_cnd_sub_n
func (recv_ MpLimbT) X__gmpnCndSubN(MpPtr, MpSrcptr, MpSrcptr, MpSizeT) MpLimbT {
	return 0
}

//go:linkname X__gmpnSecAdd1 C.__gmpn_sec_add_1
func X__gmpnSecAdd1(MpPtr, MpSrcptr, MpSizeT, MpLimbT, MpPtr) MpLimbT

// llgo:link MpSizeT.X__gmpnSecAdd1Itch C.__gmpn_sec_add_1_itch
func (recv_ MpSizeT) X__gmpnSecAdd1Itch() MpSizeT {
	return 0
}

//go:linkname X__gmpnSecSub1 C.__gmpn_sec_sub_1
func X__gmpnSecSub1(MpPtr, MpSrcptr, MpSizeT, MpLimbT, MpPtr) MpLimbT

// llgo:link MpSizeT.X__gmpnSecSub1Itch C.__gmpn_sec_sub_1_itch
func (recv_ MpSizeT) X__gmpnSecSub1Itch() MpSizeT {
	return 0
}

// llgo:link MpLimbT.X__gmpnCndSwap C.__gmpn_cnd_swap
func (recv_ MpLimbT) X__gmpnCndSwap(*MpLimbT, *MpLimbT, MpSizeT) {
}

//go:linkname X__gmpnSecMul C.__gmpn_sec_mul
func X__gmpnSecMul(MpPtr, MpSrcptr, MpSizeT, MpSrcptr, MpSizeT, MpPtr)

// llgo:link MpSizeT.X__gmpnSecMulItch C.__gmpn_sec_mul_itch
func (recv_ MpSizeT) X__gmpnSecMulItch(MpSizeT) MpSizeT {
	return 0
}

//go:linkname X__gmpnSecSqr C.__gmpn_sec_sqr
func X__gmpnSecSqr(MpPtr, MpSrcptr, MpSizeT, MpPtr)

// llgo:link MpSizeT.X__gmpnSecSqrItch C.__gmpn_sec_sqr_itch
func (recv_ MpSizeT) X__gmpnSecSqrItch() MpSizeT {
	return 0
}

//go:linkname X__gmpnSecPowm C.__gmpn_sec_powm
func X__gmpnSecPowm(MpPtr, MpSrcptr, MpSizeT, MpSrcptr, MpBitcntT, MpSrcptr, MpSizeT, MpPtr)

// llgo:link MpSizeT.X__gmpnSecPowmItch C.__gmpn_sec_powm_itch
func (recv_ MpSizeT) X__gmpnSecPowmItch(MpBitcntT, MpSizeT) MpSizeT {
	return 0
}

// llgo:link (*MpLimbT).X__gmpnSecTabselect C.__gmpn_sec_tabselect
func (recv_ *MpLimbT) X__gmpnSecTabselect(*MpLimbT, MpSizeT, MpSizeT, MpSizeT) {
}

//go:linkname X__gmpnSecDivQr C.__gmpn_sec_div_qr
func X__gmpnSecDivQr(MpPtr, MpPtr, MpSizeT, MpSrcptr, MpSizeT, MpPtr) MpLimbT

// llgo:link MpSizeT.X__gmpnSecDivQrItch C.__gmpn_sec_div_qr_itch
func (recv_ MpSizeT) X__gmpnSecDivQrItch(MpSizeT) MpSizeT {
	return 0
}

//go:linkname X__gmpnSecDivR C.__gmpn_sec_div_r
func X__gmpnSecDivR(MpPtr, MpSizeT, MpSrcptr, MpSizeT, MpPtr)

// llgo:link MpSizeT.X__gmpnSecDivRItch C.__gmpn_sec_div_r_itch
func (recv_ MpSizeT) X__gmpnSecDivRItch(MpSizeT) MpSizeT {
	return 0
}

//go:linkname X__gmpnSecInvert C.__gmpn_sec_invert
func X__gmpnSecInvert(MpPtr, MpPtr, MpSrcptr, MpSizeT, MpBitcntT, MpPtr) c.Int

// llgo:link MpSizeT.X__gmpnSecInvertItch C.__gmpn_sec_invert_itch
func (recv_ MpSizeT) X__gmpnSecInvertItch() MpSizeT {
	return 0
}

const (
	GMPERRORNONE                c.Int = 0
	GMPERRORUNSUPPORTEDARGUMENT c.Int = 1
	GMPERRORDIVISIONBYZERO      c.Int = 2
	GMPERRORSQRTOFNEGATIVE      c.Int = 4
	GMPERRORINVALIDARGUMENT     c.Int = 8
	GMPERRORMPZOVERFLOW         c.Int = 16
)
