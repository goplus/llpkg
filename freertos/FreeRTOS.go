package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ConfigUSE_C_RUNTIME_TLS_SUPPORT = 1
const INCLUDE_xQueueGetMutexHolder = 0
const ConfigUSE_DAEMON_TASK_STARTUP_HOOK = 0
const ConfigUSE_APPLICATION_TASK_TAG = 0
const ConfigUSE_ALTERNATIVE_API = 0
const ConfigASSERT_DEFINED = 1
const ConfigPRECONDITION_DEFINED = 0
const ConfigUSE_MINI_LIST_ITEM = 1
const ConfigGENERATE_RUN_TIME_STATS = 0
const ConfigUSE_MALLOC_FAILED_HOOK = 0
const ConfigEXPECTED_IDLE_TIME_BEFORE_SLEEP = 2
const ConfigINCLUDE_APPLICATION_DEFINED_PRIVILEGED_FUNCTIONS = 0
const ConfigUSE_STATS_FORMATTING_FUNCTIONS = 0
const ConfigUSE_TRACE_FACILITY = 0
const ConfigUSE_POSIX_ERRNO = 0
const ConfigUSE_SB_COMPLETED_CALLBACK = 0
const ConfigINITIAL_TICK_COUNT = 0
const ConfigUSE_TASK_FPU_SUPPORT = 1
const ConfigENABLE_MPU = 0
const ConfigENABLE_FPU = 1
const ConfigENABLE_MVE = 0
const ConfigENABLE_TRUSTZONE = 1
const ConfigRUN_FREERTOS_SECURE_ONLY = 0
const ConfigRUN_ADDITIONAL_TESTS = 0

/*
 * In line with software engineering best practice, FreeRTOS implements a strict
 * data hiding policy, so the real structures used by FreeRTOS to maintain the
 * state of tasks, queues, semaphores, etc. are not accessible to the application
 * code.  However, if the application writer wants to statically allocate such
 * an object then the size of the object needs to be known.  Dummy structures
 * that are guaranteed to have the same size and alignment requirements of the
 * real objects are used for this purpose.  The dummy list and list item
 * structures below are used for inclusion in such a dummy structure.
 */
type XSTATICLISTITEM struct {
	XDummy2  TickTypeT
	PvDummy3 [4]c.Pointer
}
type StaticListItemT XSTATICLISTITEM

/* See the comments above the struct xSTATIC_LIST_ITEM definition. */

type XSTATICMINILISTITEM struct {
	XDummy2  TickTypeT
	PvDummy3 [2]c.Pointer
}
type StaticMiniListItemT XSTATICMINILISTITEM

/* See the comments above the struct xSTATIC_LIST_ITEM definition. */

type XSTATICLIST struct {
	UxDummy2 UBaseTypeT
	PvDummy3 c.Pointer
	XDummy4  StaticMiniListItemT
}
type StaticListT XSTATICLIST

/*
 * In line with software engineering best practice, especially when supplying a
 * library that is likely to change in future versions, FreeRTOS implements a
 * strict data hiding policy.  This means the Task structure used internally by
 * FreeRTOS is not accessible to application code.  However, if the application
 * writer wants to statically allocate the memory required to create a task then
 * the size of the task object needs to be known.  The StaticTask_t structure
 * below is provided for this purpose.  Its sizes and alignment requirements are
 * guaranteed to match those of the genuine structure, no matter which
 * architecture is being used, and no matter how the values in FreeRTOSConfig.h
 * are set.  Its contents are somewhat obfuscated in the hope users will
 * recognise that it would be unwise to make direct use of the structure members.
 */
type XSTATICTCB struct {
	PxDummy1     c.Pointer
	XDummy3      [2]StaticListItemT
	UxDummy5     UBaseTypeT
	PxDummy6     c.Pointer
	UcDummy7     [16]c.Uint8T
	XDummyCoreID BaseTypeT
	PxDummy8     c.Pointer
	UxDummy12    [2]UBaseTypeT
	PvDummy15    [2]c.Pointer
	XDummy17     X_reent
	UlDummy18    [1]c.Uint32T
	UcDummy19    [1]c.Uint8T
	UxDummy20    c.Uint8T
	UcDummy21    c.Uint8T
}
type StaticTaskT XSTATICTCB

/*
 * In line with software engineering best practice, especially when supplying a
 * library that is likely to change in future versions, FreeRTOS implements a
 * strict data hiding policy.  This means the Queue structure used internally by
 * FreeRTOS is not accessible to application code.  However, if the application
 * writer wants to statically allocate the memory required to create a queue
 * then the size of the queue object needs to be known.  The StaticQueue_t
 * structure below is provided for this purpose.  Its sizes and alignment
 * requirements are guaranteed to match those of the genuine structure, no
 * matter which architecture is being used, and no matter how the values in
 * FreeRTOSConfig.h are set.  Its contents are somewhat obfuscated in the hope
 * users will recognise that it would be unwise to make direct use of the
 * structure members.
 */
type XSTATICQUEUE struct {
	PvDummy1 [3]c.Pointer
	U        struct {
		PvDummy2 c.Pointer
	}
	XDummy3         [2]StaticListT
	UxDummy4        [3]UBaseTypeT
	UcDummy5        [2]c.Uint8T
	UcDummy6        c.Uint8T
	PvDummy7        c.Pointer
	XDummyQueueLock PortMUXTYPE
}
type StaticQueueT XSTATICQUEUE
type StaticSemaphoreT StaticQueueT

/*
 * In line with software engineering best practice, especially when supplying a
 * library that is likely to change in future versions, FreeRTOS implements a
 * strict data hiding policy.  This means the event group structure used
 * internally by FreeRTOS is not accessible to application code.  However, if
 * the application writer wants to statically allocate the memory required to
 * create an event group then the size of the event group object needs to be
 * know.  The StaticEventGroup_t structure below is provided for this purpose.
 * Its sizes and alignment requirements are guaranteed to match those of the
 * genuine structure, no matter which architecture is being used, and no matter
 * how the values in FreeRTOSConfig.h are set.  Its contents are somewhat
 * obfuscated in the hope users will recognise that it would be unwise to make
 * direct use of the structure members.
 */
type XSTATICEVENTGROUP struct {
	XDummy1              TickTypeT
	XDummy2              StaticListT
	UcDummy4             c.Uint8T
	XDummyEventGroupLock PortMUXTYPE
}
type StaticEventGroupT XSTATICEVENTGROUP

/*
 * In line with software engineering best practice, especially when supplying a
 * library that is likely to change in future versions, FreeRTOS implements a
 * strict data hiding policy.  This means the software timer structure used
 * internally by FreeRTOS is not accessible to application code.  However, if
 * the application writer wants to statically allocate the memory required to
 * create a software timer then the size of the queue object needs to be known.
 * The StaticTimer_t structure below is provided for this purpose.  Its sizes
 * and alignment requirements are guaranteed to match those of the genuine
 * structure, no matter which architecture is being used, and no matter how the
 * values in FreeRTOSConfig.h are set.  Its contents are somewhat obfuscated in
 * the hope users will recognise that it would be unwise to make direct use of
 * the structure members.
 */
type XSTATICTIMER struct {
	PvDummy1 c.Pointer
	XDummy2  StaticListItemT
	XDummy3  TickTypeT
	PvDummy5 c.Pointer
	PvDummy6 TaskFunctionT
	UcDummy8 c.Uint8T
}
type StaticTimerT XSTATICTIMER

/*
 * In line with software engineering best practice, especially when supplying a
 * library that is likely to change in future versions, FreeRTOS implements a
 * strict data hiding policy.  This means the stream buffer structure used
 * internally by FreeRTOS is not accessible to application code.  However, if
 * the application writer wants to statically allocate the memory required to
 * create a stream buffer then the size of the stream buffer object needs to be
 * known.  The StaticStreamBuffer_t structure below is provided for this
 * purpose.  Its size and alignment requirements are guaranteed to match those
 * of the genuine structure, no matter which architecture is being used, and
 * no matter how the values in FreeRTOSConfig.h are set.  Its contents are
 * somewhat obfuscated in the hope users will recognise that it would be unwise
 * to make direct use of the structure members.
 */
type XSTATICSTREAMBUFFER struct {
	UxDummy1               [4]c.SizeT
	PvDummy2               [3]c.Pointer
	UcDummy3               c.Uint8T
	XDummyStreamBufferLock PortMUXTYPE
}
type StaticStreamBufferT XSTATICSTREAMBUFFER
type StaticMessageBufferT StaticStreamBufferT
