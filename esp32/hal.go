package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const XTHAL_RELEASE_MAJOR = 12000
const XTHAL_RELEASE_MINOR = 9
const XTHAL_RELEASE_NAME = "12.0.9"
const XTHAL_REL_12 = 1
const XTHAL_REL_12_0 = 1
const XTHAL_REL_12_0_9 = 1
const XTHAL_MAX_CPS = 8
const XTHAL_LITTLEENDIAN = 0
const XTHAL_BIGENDIAN = 1
const XTHAL_PREFETCH_DISABLE = 0xFFFF0000
const XTHAL_DCACHE_PREFETCH_L1_OFF = 0x90000000
const XTHAL_DCACHE_PREFETCH_L1 = 0x90001000
const XTHAL_ICACHE_PREFETCH_L1_OFF = 0xA0000000
const XTHAL_ICACHE_PREFETCH_L1 = 0xA0002000
const XTHAL_DISASM_BUFSIZE = 80
const XTHAL_DISASM_OPT_ADDR = 0x0001
const XTHAL_DISASM_OPT_OPHEX = 0x0002
const XTHAL_DISASM_OPT_OPCODE = 0x0004
const XTHAL_DISASM_OPT_PARMS = 0x0008
const XTHAL_DISASM_OPT_ALL = 0x0FFF
const XTHAL_MAX_INTERRUPTS = 32
const XTHAL_MAX_INTLEVELS = 16
const XTHAL_MAX_TIMERS = 4
const XTHAL_INTTYPE_UNCONFIGURED = 0
const XTHAL_INTTYPE_SOFTWARE = 1
const XTHAL_INTTYPE_EXTERN_EDGE = 2
const XTHAL_INTTYPE_EXTERN_LEVEL = 3
const XTHAL_INTTYPE_TIMER = 4
const XTHAL_INTTYPE_NMI = 5
const XTHAL_INTTYPE_WRITE_ERROR = 6
const XTHAL_INTTYPE_PROFILING = 7
const XTHAL_INTTYPE_IDMA_DONE = 8
const XTHAL_INTTYPE_IDMA_ERR = 9
const XTHAL_INTTYPE_GS_ERR = 10
const XTHAL_INTTYPE_SG_ERR = 10
const XTHAL_MAX_INTTYPES = 11
const XTHAL_MEMEP_PARITY = 1
const XTHAL_MEMEP_ECC = 2
const XTHAL_MEMEP_F_LOCAL = 0
const XTHAL_MEMEP_F_DCACHE_DATA = 4
const XTHAL_MEMEP_F_DCACHE_TAG = 5
const XTHAL_MEMEP_F_ICACHE_DATA = 6
const XTHAL_MEMEP_F_ICACHE_TAG = 7
const XTHAL_MEMEP_F_CORRECTABLE = 16
const XTHAL_AMB_EXCEPTION = 0
const XTHAL_AMB_HITCACHE = 1
const XTHAL_AMB_ALLOCATE = 2
const XTHAL_AMB_WRITETHRU = 3
const XTHAL_AMB_ISOLATE = 4
const XTHAL_AMB_GUARD = 5
const XTHAL_AMB_COHERENT = 6
const XTHAL_FAM_EXCEPTION = 0x001
const XTHAL_FAM_BYPASS = 0x000
const XTHAL_FAM_CACHED = 0x006
const XTHAL_LAM_EXCEPTION = 0x001
const XTHAL_LAM_ISOLATE = 0x012
const XTHAL_LAM_BYPASS = 0x000
const XTHAL_LAM_BYPASSG = 0x020
const XTHAL_LAM_CACHED_NOALLOC = 0x002
const XTHAL_LAM_NACACHEDG = 0x022
const XTHAL_LAM_CACHED = 0x006
const XTHAL_LAM_COHCACHED = 0x046
const XTHAL_SAM_EXCEPTION = 0x001
const XTHAL_SAM_ISOLATE = 0x032
const XTHAL_SAM_BYPASS = 0x028
const XTHAL_SAM_WRITETHRU = 0x02A
const XTHAL_SAM_WRITEBACK = 0x026
const XTHAL_SAM_WRITEBACK_NOALLOC = 0x022
const XTHAL_SAM_COHWRITEBACK = 0x066
const XTHAL_PAM_BYPASS = 0x000
const XTHAL_PAM_BYPASS_BUF = 0x010
const XTHAL_PAM_CACHED_NOALLOC = 0x030
const XTHAL_PAM_WRITETHRU = 0x0B0
const XTHAL_PAM_WRITEBACK_NOALLOC = 0x0F0
const XTHAL_PAM_WRITEBACK = 0x1F0
const XTHAL_CAFLAG_EXPAND = 0x000100
const XTHAL_CAFLAG_EXACT = 0x000200
const XTHAL_CAFLAG_NO_PARTIAL = 0x000400
const XTHAL_CAFLAG_NO_AUTO_WB = 0x000800
const XTHAL_CAFLAG_NO_AUTO_INV = 0x001000
const XTHAL_SUCCESS = 0
const XTHAL_AR_NONE = 0
const XTHAL_AR_R = 4
const XTHAL_AR_RX = 5
const XTHAL_AR_RW = 6
const XTHAL_AR_RWX = 7
const XTHAL_AR_Ww = 8
const XTHAL_AR_RWrwx = 9
const XTHAL_AR_RWr = 10
const XTHAL_AR_RWXrx = 11
const XTHAL_AR_Rr = 12
const XTHAL_AR_RXrx = 13
const XTHAL_AR_RWrw = 14
const XTHAL_AR_RWXrwx = 15
const XTHAL_AR_WIDTH = 4
const XTHAL_MPU_USE_EXISTING_ACCESS_RIGHTS = 0x00002000
const XTHAL_MPU_USE_EXISTING_MEMORY_TYPE = 0x00004000
const XTHAL_MEM_DEVICE = 0x00008000
const XTHAL_MEM_NON_CACHEABLE = 0x00090000
const XTHAL_MEM_WRITETHRU_NOALLOC = 0x00080000
const XTHAL_MEM_WRITETHRU = 0x00040000
const XTHAL_MEM_WRITETHRU_WRITEALLOC = 0x00060000
const XTHAL_MEM_WRITEBACK_NOALLOC = 0x00050000
const XTHAL_MEM_WRITEBACK = 0x00070000
const XTHAL_MEM_INTERRUPTIBLE = 0x08000000
const XTHAL_MEM_BUFFERABLE = 0x01000000
const XTHAL_MEM_NON_SHAREABLE = 0x00000000
const XTHAL_MEM_INNER_SHAREABLE = 0x02000000
const XTHAL_MEM_OUTER_SHAREABLE = 0x04000000
const XTHAL_MEM_SYSTEM_SHAREABLE = 0x06000000
const X_XTHAL_SYSTEM_CACHE_BITS = 0x000f0000
const X_XTHAL_LOCAL_CACHE_BITS = 0x00f00000
const X_XTHAL_MEM_SYSTEM_RWC_MASK = 0x00070000
const X_XTHAL_MEM_LOCAL_RWC_MASK = 0x00700000
const X_XTHAL_SHIFT_RWC = 16
const XTHAL_MEM_SW_SHAREABLE = 0

/* save & restore the extra processor state */
//go:linkname XthalSaveExtra C.xthal_save_extra
func XthalSaveExtra(base c.Pointer)

//go:linkname XthalRestoreExtra C.xthal_restore_extra
func XthalRestoreExtra(base c.Pointer)

//go:linkname XthalSaveCpregs C.xthal_save_cpregs
func XthalSaveCpregs(base c.Pointer, __llgo_arg_1 c.Int)

//go:linkname XthalRestoreCpregs C.xthal_restore_cpregs
func XthalRestoreCpregs(base c.Pointer, __llgo_arg_1 c.Int)

/* versions specific to each coprocessor id */
//go:linkname XthalSaveCp0 C.xthal_save_cp0
func XthalSaveCp0(base c.Pointer)

//go:linkname XthalSaveCp1 C.xthal_save_cp1
func XthalSaveCp1(base c.Pointer)

//go:linkname XthalSaveCp2 C.xthal_save_cp2
func XthalSaveCp2(base c.Pointer)

//go:linkname XthalSaveCp3 C.xthal_save_cp3
func XthalSaveCp3(base c.Pointer)

//go:linkname XthalSaveCp4 C.xthal_save_cp4
func XthalSaveCp4(base c.Pointer)

//go:linkname XthalSaveCp5 C.xthal_save_cp5
func XthalSaveCp5(base c.Pointer)

//go:linkname XthalSaveCp6 C.xthal_save_cp6
func XthalSaveCp6(base c.Pointer)

//go:linkname XthalSaveCp7 C.xthal_save_cp7
func XthalSaveCp7(base c.Pointer)

//go:linkname XthalRestoreCp0 C.xthal_restore_cp0
func XthalRestoreCp0(base c.Pointer)

//go:linkname XthalRestoreCp1 C.xthal_restore_cp1
func XthalRestoreCp1(base c.Pointer)

//go:linkname XthalRestoreCp2 C.xthal_restore_cp2
func XthalRestoreCp2(base c.Pointer)

//go:linkname XthalRestoreCp3 C.xthal_restore_cp3
func XthalRestoreCp3(base c.Pointer)

//go:linkname XthalRestoreCp4 C.xthal_restore_cp4
func XthalRestoreCp4(base c.Pointer)

//go:linkname XthalRestoreCp5 C.xthal_restore_cp5
func XthalRestoreCp5(base c.Pointer)

//go:linkname XthalRestoreCp6 C.xthal_restore_cp6
func XthalRestoreCp6(base c.Pointer)

//go:linkname XthalRestoreCp7 C.xthal_restore_cp7
func XthalRestoreCp7(base c.Pointer)

/* initialize the extra processor */
//go:linkname XthalInitMemExtra C.xthal_init_mem_extra
func XthalInitMemExtra(c.Pointer)

/* initialize the TIE coprocessor */
//go:linkname XthalInitMemCp C.xthal_init_mem_cp
func XthalInitMemCp(c.Pointer, c.Int)

/* cache region operations*/
//go:linkname XthalIcacheRegionInvalidate C.xthal_icache_region_invalidate
func XthalIcacheRegionInvalidate(addr c.Pointer, size c.Uint)

//go:linkname XthalDcacheRegionInvalidate C.xthal_dcache_region_invalidate
func XthalDcacheRegionInvalidate(addr c.Pointer, size c.Uint)

//go:linkname XthalDcacheRegionWriteback C.xthal_dcache_region_writeback
func XthalDcacheRegionWriteback(addr c.Pointer, size c.Uint)

//go:linkname XthalDcacheRegionWritebackInv C.xthal_dcache_region_writeback_inv
func XthalDcacheRegionWritebackInv(addr c.Pointer, size c.Uint)

/* cache line operations*/
//go:linkname XthalIcacheLineInvalidate C.xthal_icache_line_invalidate
func XthalIcacheLineInvalidate(addr c.Pointer)

//go:linkname XthalDcacheLineInvalidate C.xthal_dcache_line_invalidate
func XthalDcacheLineInvalidate(addr c.Pointer)

//go:linkname XthalDcacheLineWriteback C.xthal_dcache_line_writeback
func XthalDcacheLineWriteback(addr c.Pointer)

//go:linkname XthalDcacheLineWritebackInv C.xthal_dcache_line_writeback_inv
func XthalDcacheLineWritebackInv(addr c.Pointer)

/* sync icache and memory */
//go:linkname XthalIcacheSync C.xthal_icache_sync
func XthalIcacheSync()

/* sync dcache and memory */
//go:linkname XthalDcacheSync C.xthal_dcache_sync
func XthalDcacheSync()

/* get/set number of icache ways enabled */
//go:linkname XthalIcacheGetWays C.xthal_icache_get_ways
func XthalIcacheGetWays() c.Uint

//go:linkname XthalIcacheSetWays C.xthal_icache_set_ways
func XthalIcacheSetWays(ways c.Uint)

/* get/set number of dcache ways enabled */
//go:linkname XthalDcacheGetWays C.xthal_dcache_get_ways
func XthalDcacheGetWays() c.Uint

//go:linkname XthalDcacheSetWays C.xthal_dcache_set_ways
func XthalDcacheSetWays(ways c.Uint)

/* coherency (low-level -- not normally called directly) */
//go:linkname XthalCacheCoherenceOn C.xthal_cache_coherence_on
func XthalCacheCoherenceOn()

//go:linkname XthalCacheCoherenceOff C.xthal_cache_coherence_off
func XthalCacheCoherenceOff()

/* coherency (high-level) */
//go:linkname XthalCacheCoherenceOptin C.xthal_cache_coherence_optin
func XthalCacheCoherenceOptin()

//go:linkname XthalCacheCoherenceOptout C.xthal_cache_coherence_optout
func XthalCacheCoherenceOptout()

//go:linkname XthalGetCachePrefetch C.xthal_get_cache_prefetch
func XthalGetCachePrefetch() c.Int

//go:linkname XthalSetCachePrefetch C.xthal_set_cache_prefetch
func XthalSetCachePrefetch(c.Int) c.Int

//go:linkname XthalSetCachePrefetchLong C.xthal_set_cache_prefetch_long
func XthalSetCachePrefetchLong(c.UlongLong) c.Int

/*  Set (plant) and remove software breakpoint, both synchronizing cache:  */
//go:linkname XthalSetSoftBreak C.xthal_set_soft_break
func XthalSetSoftBreak(addr c.Pointer) c.Uint

//go:linkname XthalRemoveSoftBreak C.xthal_remove_soft_break
func XthalRemoveSoftBreak(addr c.Pointer, __llgo_arg_1 c.Uint)

/* routine to get a string for the disassembled instruction */
//go:linkname XthalDisassemble C.xthal_disassemble
func XthalDisassemble(instr_buf *c.Char, tgt_addr c.Pointer, buffer *c.Char, buflen c.Uint, options c.Uint) c.Int

/* routine to get the size of the next instruction. Returns 0 for
   illegal instruction */
//go:linkname XthalDisassembleSize C.xthal_disassemble_size
func XthalDisassembleSize(instr_buf *c.Char) c.Int

/*----------------------------------------------------------------------
			Instruction/Data RAM/ROM Access
  ----------------------------------------------------------------------*/
//go:linkname XthalMemcpy C.xthal_memcpy
func XthalMemcpy(dst c.Pointer, src c.Pointer, len c.Uint) c.Pointer

//go:linkname XthalBcopy C.xthal_bcopy
func XthalBcopy(src c.Pointer, dst c.Pointer, len c.Uint) c.Pointer

/*----------------------------------------------------------------------
                         MP Synchronization
----------------------------------------------------------------------*/
//go:linkname XthalCompareAndSet C.xthal_compare_and_set
func XthalCompareAndSet(addr *c.Int, test_val c.Int, compare_val c.Int) c.Int

/*  Clear any remnant code-dependent state (i.e. clear loop count regs).  */
//go:linkname XthalClearRegcachedCode C.xthal_clear_regcached_code
func XthalClearRegcachedCode()

/*  This spill any live register windows (other than the caller's):
 *  (NOTE:  current implementation require privileged code, but
 *   a user-callable implementation is possible.)  */
//go:linkname XthalWindowSpill C.xthal_window_spill
func XthalWindowSpill()

/* validate & invalidate the TIE register file */
//go:linkname XthalValidateCp C.xthal_validate_cp
func XthalValidateCp(c.Int)

//go:linkname XthalInvalidateCp C.xthal_invalidate_cp
func XthalInvalidateCp(c.Int)

/* read and write cpenable register */
//go:linkname XthalSetCpenable C.xthal_set_cpenable
func XthalSetCpenable(c.Uint)

//go:linkname XthalGetCpenable C.xthal_get_cpenable
func XthalGetCpenable() c.Uint

/*  INTENABLE,INTERRUPT,INTSET,INTCLEAR register access functions:  */
//go:linkname XthalGetIntenable C.xthal_get_intenable
func XthalGetIntenable() c.Uint

//go:linkname XthalSetIntenable C.xthal_set_intenable
func XthalSetIntenable(c.Uint)

//go:linkname XthalGetInterrupt C.xthal_get_interrupt
func XthalGetInterrupt() c.Uint

/*  These two functions are deprecated. Use the newer functions
    xthal_interrupt_trigger and xthal_interrupt_clear instead.  */
//go:linkname XthalSetIntset C.xthal_set_intset
func XthalSetIntset(c.Uint)

//go:linkname XthalSetIntclear C.xthal_set_intclear
func XthalSetIntclear(c.Uint)

/* get CCOUNT register (if not present return 0) */
//go:linkname XthalGetCcount C.xthal_get_ccount
func XthalGetCcount() c.Uint

/* set and get CCOMPAREn registers (if not present, get returns 0) */
//go:linkname XthalSetCcompare C.xthal_set_ccompare
func XthalSetCcompare(c.Int, c.Uint)

//go:linkname XthalGetCcompare C.xthal_get_ccompare
func XthalGetCcompare(c.Int) c.Uint

//go:linkname XthalGetPrid C.xthal_get_prid
func XthalGetPrid() c.Uint

/*  Convert between interrupt levels (as per PS.INTLEVEL) and virtual interrupt priorities:  */
//go:linkname XthalVpriToIntlevel C.xthal_vpri_to_intlevel
func XthalVpriToIntlevel(vpri c.Uint) c.Uint

//go:linkname XthalIntlevelToVpri C.xthal_intlevel_to_vpri
func XthalIntlevelToVpri(intlevel c.Uint) c.Uint

/*  Enables/disables given set (mask) of interrupts; returns previous enabled-mask of all ints:  */
/*  These functions are deprecated. Use xthal_interrupt_enable and xthal_interrupt_disable instead.  */
//go:linkname XthalIntEnable C.xthal_int_enable
func XthalIntEnable(c.Uint) c.Uint

//go:linkname XthalIntDisable C.xthal_int_disable
func XthalIntDisable(c.Uint) c.Uint

/*  Set/get virtual priority of an interrupt:  */
//go:linkname XthalSetIntVpri C.xthal_set_int_vpri
func XthalSetIntVpri(intnum c.Int, vpri c.Int) c.Int

//go:linkname XthalGetIntVpri C.xthal_get_int_vpri
func XthalGetIntVpri(intnum c.Int) c.Int

/*  Set/get interrupt lockout level for exclusive access to virtual priority data structures:  */
//go:linkname XthalSetVpriLocklevel C.xthal_set_vpri_locklevel
func XthalSetVpriLocklevel(intlevel c.Uint)

//go:linkname XthalGetVpriLocklevel C.xthal_get_vpri_locklevel
func XthalGetVpriLocklevel() c.Uint

/*  Set/get current virtual interrupt priority:  */
//go:linkname XthalSetVpri C.xthal_set_vpri
func XthalSetVpri(vpri c.Uint) c.Uint

//go:linkname XthalGetVpri C.xthal_get_vpri
func XthalGetVpri() c.Uint

//go:linkname XthalSetVpriIntlevel C.xthal_set_vpri_intlevel
func XthalSetVpriIntlevel(intlevel c.Uint) c.Uint

//go:linkname XthalSetVpriLock C.xthal_set_vpri_lock
func XthalSetVpriLock() c.Uint

// llgo:type C
type XtHalVoidFunc func()

/*  Trampoline support functions:  */
//go:linkname XthalTramPendingToService C.xthal_tram_pending_to_service
func XthalTramPendingToService() c.Uint

//go:linkname XthalTramDone C.xthal_tram_done
func XthalTramDone(serviced_mask c.Uint)

//go:linkname XthalTramSetSync C.xthal_tram_set_sync
func XthalTramSetSync(intnum c.Int, sync c.Int) c.Int

// llgo:link (*XtHalVoidFunc).XthalSetTramTriggerFunc C.xthal_set_tram_trigger_func
func (recv_ *XtHalVoidFunc) XthalSetTramTriggerFunc() XtHalVoidFunc {
	return nil
}

/* cache attribute register control (used by other HAL routines) */
//go:linkname XthalGetCacheattr C.xthal_get_cacheattr
func XthalGetCacheattr() c.Uint

//go:linkname XthalGetIcacheattr C.xthal_get_icacheattr
func XthalGetIcacheattr() c.Uint

//go:linkname XthalGetDcacheattr C.xthal_get_dcacheattr
func XthalGetDcacheattr() c.Uint

//go:linkname XthalSetCacheattr C.xthal_set_cacheattr
func XthalSetCacheattr(c.Uint)

//go:linkname XthalSetIcacheattr C.xthal_set_icacheattr
func XthalSetIcacheattr(c.Uint)

//go:linkname XthalSetDcacheattr C.xthal_set_dcacheattr
func XthalSetDcacheattr(c.Uint)

/* set cache attribute (access modes) for a range of memory */
//go:linkname XthalSetRegionAttribute C.xthal_set_region_attribute
func XthalSetRegionAttribute(addr c.Pointer, size c.Uint, cattr c.Uint, flags c.Uint) c.Int

/* enable caches */
//go:linkname XthalIcacheEnable C.xthal_icache_enable
func XthalIcacheEnable()

//go:linkname XthalDcacheEnable C.xthal_dcache_enable
func XthalDcacheEnable()

/* disable caches */
//go:linkname XthalIcacheDisable C.xthal_icache_disable
func XthalIcacheDisable()

//go:linkname XthalDcacheDisable C.xthal_dcache_disable
func XthalDcacheDisable()

/* whole cache operations (privileged) */
//go:linkname XthalIcacheAllInvalidate C.xthal_icache_all_invalidate
func XthalIcacheAllInvalidate()

//go:linkname XthalDcacheAllInvalidate C.xthal_dcache_all_invalidate
func XthalDcacheAllInvalidate()

//go:linkname XthalDcacheAllWriteback C.xthal_dcache_all_writeback
func XthalDcacheAllWriteback()

//go:linkname XthalDcacheAllWritebackInv C.xthal_dcache_all_writeback_inv
func XthalDcacheAllWritebackInv()

//go:linkname XthalIcacheAllUnlock C.xthal_icache_all_unlock
func XthalIcacheAllUnlock()

//go:linkname XthalDcacheAllUnlock C.xthal_dcache_all_unlock
func XthalDcacheAllUnlock()

/* address-range cache operations (privileged) */
/* prefetch and lock specified memory range into cache */
//go:linkname XthalIcacheRegionLock C.xthal_icache_region_lock
func XthalIcacheRegionLock(addr c.Pointer, size c.Uint)

//go:linkname XthalDcacheRegionLock C.xthal_dcache_region_lock
func XthalDcacheRegionLock(addr c.Pointer, size c.Uint)

/* unlock from cache */
//go:linkname XthalIcacheRegionUnlock C.xthal_icache_region_unlock
func XthalIcacheRegionUnlock(addr c.Pointer, size c.Uint)

//go:linkname XthalDcacheRegionUnlock C.xthal_dcache_region_unlock
func XthalDcacheRegionUnlock(addr c.Pointer, size c.Uint)

/* huge-range cache operations (privileged) (EXPERIMENTAL) */
//go:linkname XthalIcacheHugerangeInvalidate C.xthal_icache_hugerange_invalidate
func XthalIcacheHugerangeInvalidate(addr c.Pointer, size c.Uint)

//go:linkname XthalIcacheHugerangeUnlock C.xthal_icache_hugerange_unlock
func XthalIcacheHugerangeUnlock(addr c.Pointer, size c.Uint)

//go:linkname XthalDcacheHugerangeInvalidate C.xthal_dcache_hugerange_invalidate
func XthalDcacheHugerangeInvalidate(addr c.Pointer, size c.Uint)

//go:linkname XthalDcacheHugerangeUnlock C.xthal_dcache_hugerange_unlock
func XthalDcacheHugerangeUnlock(addr c.Pointer, size c.Uint)

//go:linkname XthalDcacheHugerangeWriteback C.xthal_dcache_hugerange_writeback
func XthalDcacheHugerangeWriteback(addr c.Pointer, size c.Uint)

//go:linkname XthalDcacheHugerangeWritebackInv C.xthal_dcache_hugerange_writeback_inv
func XthalDcacheHugerangeWritebackInv(addr c.Pointer, size c.Uint)

/* cache line operations (privileged) */
//go:linkname XthalIcacheLineLock C.xthal_icache_line_lock
func XthalIcacheLineLock(addr c.Pointer)

//go:linkname XthalDcacheLineLock C.xthal_dcache_line_lock
func XthalDcacheLineLock(addr c.Pointer)

//go:linkname XthalIcacheLineUnlock C.xthal_icache_line_unlock
func XthalIcacheLineUnlock(addr c.Pointer)

//go:linkname XthalDcacheLineUnlock C.xthal_dcache_line_unlock
func XthalDcacheLineUnlock(addr c.Pointer)

/*  Inject memory errors; flags is bit combination of XTHAL_MEMEP_F_xxx:  */
//go:linkname XthalMemepInjectError C.xthal_memep_inject_error
func XthalMemepInjectError(addr c.Pointer, size c.Int, flags c.Int)

/*  Convert between virtual and physical addresses (through static maps only)
 *  WARNING: these two functions may go away in a future release;
 *  don't depend on them!
 */
//go:linkname XthalStaticV2p C.xthal_static_v2p
func XthalStaticV2p(vaddr c.Uint, paddrp *c.Uint) c.Int

//go:linkname XthalStaticP2v C.xthal_static_p2v
func XthalStaticP2v(paddr c.Uint, vaddrp *c.Uint, cached c.Uint) c.Int

//go:linkname XthalSetRegionTranslation C.xthal_set_region_translation
func XthalSetRegionTranslation(vaddr c.Pointer, paddr c.Pointer, size c.Uint, cache_atr c.Uint, flags c.Uint) c.Int

//go:linkname XthalV2p C.xthal_v2p
func XthalV2p(c.Pointer, *c.Pointer, *c.Uint, *c.Uint) c.Int

//go:linkname XthalInvalidateRegion C.xthal_invalidate_region
func XthalInvalidateRegion(addr c.Pointer) c.Int

//go:linkname XthalSetRegionTranslationRaw C.xthal_set_region_translation_raw
func XthalSetRegionTranslationRaw(vaddr c.Pointer, paddr c.Pointer, cattr c.Uint) c.Int

/*
 * This structure is used to represent each MPU entry (both foreground and
 * background). The internal representation of the structure is subject to
 * change, so it should only be accessed by the XTHAL_MPU_ENTRY_... macros
 * below.
 */

type XthalMPUEntry struct {
	As c.Uint32T
	At c.Uint32T
}

/*
 * These functions accept encoded access rights, and return 1 if the
 * supplied memory type has the property specified by the function name,
 * otherwise they return 0.
 */
//go:linkname XthalIsKernelReadable C.xthal_is_kernel_readable
func XthalIsKernelReadable(accessRights c.Uint32T) c.Int32T

//go:linkname XthalIsKernelWriteable C.xthal_is_kernel_writeable
func XthalIsKernelWriteable(accessRights c.Uint32T) c.Int32T

//go:linkname XthalIsKernelExecutable C.xthal_is_kernel_executable
func XthalIsKernelExecutable(accessRights c.Uint32T) c.Int32T

//go:linkname XthalIsUserReadable C.xthal_is_user_readable
func XthalIsUserReadable(accessRights c.Uint32T) c.Int32T

//go:linkname XthalIsUserWriteable C.xthal_is_user_writeable
func XthalIsUserWriteable(accessRights c.Uint32T) c.Int32T

//go:linkname XthalIsUserExecutable C.xthal_is_user_executable
func XthalIsUserExecutable(accessRights c.Uint32T) c.Int32T

/*
 * This function converts a bit-wise combination of the XTHAL_MEM_.. constants
 * to the corresponding MPU memory type (9-bits).
 *
 * If none of the XTHAL_MEM_.. bits are present in the argument, then
 * bits 4-12 (9-bits) are returned ... this supports using an already encoded
 * memoryType (perhaps obtained from an xthal_MPU_entry structure) as input
 * to xthal_set_region_attribute().
 *
 * This function first checks that the supplied constants are a valid and
 * supported combination.  If not, it returns XTHAL_BAD_MEMORY_TYPE.
 */
//go:linkname XthalEncodeMemoryType C.xthal_encode_memory_type
func XthalEncodeMemoryType(x c.Uint32T) c.Int

/*
 * This function accepts a 9-bit memory type value (such as returned by
 * XTHAL_MEM_ENTRY_GET_MEMORY_TYPE() or xthal_encode_memory_type(). They
 * return 1 if the memoryType has the property specified in the function
 * name and 0 otherwise.
 */
//go:linkname XthalIsCacheable C.xthal_is_cacheable
func XthalIsCacheable(memoryType c.Uint32T) c.Int32T

//go:linkname XthalIsWriteback C.xthal_is_writeback
func XthalIsWriteback(memoryType c.Uint32T) c.Int32T

//go:linkname XthalIsDevice C.xthal_is_device
func XthalIsDevice(memoryType c.Uint32T) c.Int32T

/*
 * Copies the current MPU entry list into 'entries' which
 * must point to available memory of at least
 * sizeof(struct xthal_MPU_entry) * XCHAL_MPU_ENTRIES.
 *
 * This function returns XTHAL_SUCCESS.
 * XTHAL_INVALID, or
 * XTHAL_UNSUPPORTED.
 */
// llgo:link (*XthalMPUEntry).XthalReadMap C.xthal_read_map
func (recv_ *XthalMPUEntry) XthalReadMap() c.Int32T {
	return 0
}

/*
 * Writes the map pointed to by 'entries' to the MPU. Before updating
 * the map, it commits any uncommitted
 * cache writes, and invalidates the cache if necessary.
 *
 * This function does not check for the correctness of the map.  Generally
 * xthal_check_map() should be called first to check the map.
 *
 * If n == 0 then the existing map is cleared, and no new map is written
 * (useful for returning to reset state)
 *
 * If (n > 0 && n < XCHAL_MPU_ENTRIES) then a new map is written with
 * (XCHAL_MPU_ENTRIES-n) padding entries added to ensure a properly ordered
 * map.  The resulting foreground map will be equivalent to the map vector
 * fg, but the position of the padding entries should not be relied upon.
 *
 * If n == XCHAL_MPU_ENTRIES then the complete map as specified by fg is
 * written.
 *
 * The CACHEADRDIS register will be set to enable caching any 512MB region
 * that is overlapped by an MPU region with a cacheable memory type.
 * Caching will be disabled if none of the 512 MB region is cacheable.
 *
 * xthal_write_map() disables the MPU foreground map during the MPU
 * update and relies on the background map.
 *
 * As a result any interrupt that does not meet the following conditions
 * must be disabled before calling xthal_write_map():
 *    1) All code and data needed for the interrupt must be
 *       mapped by the background map with sufficient access rights.
 *    2) The interrupt code must not access the MPU.
 *
 */
// llgo:link (*XthalMPUEntry).XthalWriteMap C.xthal_write_map
func (recv_ *XthalMPUEntry) XthalWriteMap(n c.Uint32T) {
}

/*
 * Checks if entry vector 'entries' of length 'n' is a valid MPU access map.
 * Returns:
 *    XTHAL_SUCCESS if valid,
 *    XTHAL_OUT_OF_ENTRIES
 *    XTHAL_MAP_NOT_ALIGNED,
 *    XTHAL_BAD_ACCESS_RIGHTS,
 *    XTHAL_OUT_OF_ORDER_MAP, or
 *    XTHAL_UNSUPPORTED if config doesn't have an MPU.
 */
// llgo:link (*XthalMPUEntry).XthalCheckMap C.xthal_check_map
func (recv_ *XthalMPUEntry) XthalCheckMap(n c.Uint32T) c.Int {
	return 0
}

/*
 * Returns the MPU entry that maps 'vaddr'. If 'infgmap' is non-NULL then
 * *infgmap is set to 1 if 'vaddr' is mapped by the foreground map, and
 * *infgmap is set to 0 if 'vaddr' is mapped by the background map.
 */
//go:linkname XthalGetEntryForAddress C.xthal_get_entry_for_address
func XthalGetEntryForAddress(vaddr c.Pointer, infgmap *c.Int32T) XthalMPUEntry

/*
 * Scans the supplied MPU map and returns a value suitable for writing to
 * the CACHEADRDIS register:
 * Bits 0-7    -> 1 if there are no cacheable areas in the corresponding 512MB
 * region and 0 otherwise.
 * Bits 8-31   -> undefined.
 * This function can accept a partial memory map in the same manner
 * xthal_write_map() does, */
// llgo:link (*XthalMPUEntry).XthalCalcCacheadrdis C.xthal_calc_cacheadrdis
func (recv_ *XthalMPUEntry) XthalCalcCacheadrdis(n c.Uint32T) c.Uint32T {
	return 0
}

/*
 * This function is intended as an MPU specific version of
 * xthal_set_region_attributes(). xthal_set_region_attributes() calls
 * this function for MPU configurations.
 *
 * This function sets the attributes for the region [vaddr, vaddr+size)
 * in the MPU.
 *
 * Depending on the state of the MPU this function will require from
 * 0 to 3 unused MPU entries.
 *
 * This function typically will move, add, and subtract entries from
 * the MPU map during execution, so that the resulting map may
 * be quite different than when the function was called.
 *
 * This function does make the following guarantees:
 *    1) The MPU access map remains in a valid state at all times
 *       during its execution.
 *    2) At all points during (and after) completion the memoryType
 *       and accessRights remain the same for all addresses
 *       that are not in the range [vaddr, vaddr+size).
 *    3) If XTHAL_SUCCESS is returned, then the range
 *       [vaddr, vaddr+size) will have the accessRights and memoryType
 *       specified.
 *    4) The CACHEADRDIS register will be set to enable caching any 512MB region
 *       that is overlapped by an MPU region with a cacheable memory type.
 *       Caching will be disabled if none of the 512 MB region is cacheable.
 *
 * The accessRights parameter should be either a 4-bit value corresponding
 * to an MPU access mode (as defined by the XTHAL_AR_.. constants), or
 * XTHAL_MPU_USE_EXISTING_ACCESS_RIGHTS.
 *
 * The memoryType parameter should be either a bit-wise or-ing of XTHAL_MEM_..
 * constants that represent a valid MPU memoryType, a 9-bit MPU memoryType
 * value, or XTHAL_MPU_USE_EXISTING_MEMORY_TYPE.
 *
 * In addition to the error codes that xthal_set_region_attribute()
 * returns, this function can also return: XTHAL_BAD_ACCESS_RIGHTS
 * (if the access rights bits map to an unsupported combination), or
 * XTHAL_OUT_OF_MAP_ENTRIES (if there are not enough unused MPU entries)
 *
 * If this function is called with an invalid MPU map, then this function
 * will return one of the codes that is returned by xthal_check_map().
 *
 * The flag, XTHAL_CAFLAG_EXPAND, is not supported
 *
 */
//go:linkname XthalMpuSetRegionAttribute C.xthal_mpu_set_region_attribute
func XthalMpuSetRegionAttribute(vaddr c.Pointer, size c.SizeT, accessRights c.Int32T, memoryType c.Int32T, flags c.Uint32T) c.Int

// llgo:link (*XthalMPUEntry).XthalReadBackgroundMap C.xthal_read_background_map
func (recv_ *XthalMPUEntry) XthalReadBackgroundMap() c.Int32T {
	return 0
}
