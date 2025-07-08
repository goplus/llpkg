package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Task Snapshot structure
 *
 * - Used with the uxTaskGetSnapshotAll() function to save memory snapshot of each task in the system.
 * - We need this structure because TCB_t is defined (hidden) in tasks.c.
 */

type XTASKSNAPSHOT struct {
	PxTCB        c.Pointer
	PxTopOfStack *StackTypeT
	PxEndOfStack *StackTypeT
}
type TaskSnapshotT XTASKSNAPSHOT

/**
 * @brief Task Snapshot iterator
 *
 * Used in xTaskGetNext(). Must be zero/null initialized on the first call.
 */

type TaskIterator struct {
	UxCurrentListIndex UBaseTypeT
	PxNextListItem     *ListItemT
	PxTaskHandle       TaskHandleT
}
type TaskIteratorT TaskIterator

/**
 * @brief Get the next task using the task iterator.
 *
 * This function retrieves the next task in the traversal sequence.
 *
 * @param xIterator Pointer to the task iterator structure.
 *
 * @return Index of the current task list. Returns -1 if all tasks have been traversed.
 *
 * @note The task iterator keeps track of the current state during task traversal,
 *       including the index of the current task list and the pointer of the next task list item.
 *       When all tasks have been traversed, this function returns -1.
 *       If a broken or corrupted task is encountered, the task handle is set to NULL.
 *
 */
// llgo:link (*TaskIteratorT).XTaskGetNext C.xTaskGetNext
func (recv_ *TaskIteratorT) XTaskGetNext() c.Int {
	return 0
}

/**
 * @brief Fill a TaskSnapshot_t structure for specified task.
 *
 * - This function is used by the panic handler to get the snapshot of a particular task.
 *
 * @note This function should only be called when FreeRTOS is no longer running (e.g., during a panic) as this function
 *       does not acquire any locks.
 * @param[in] pxTask Task's handle
 * @param[out] pxTaskSnapshot Snapshot of the task
 * @return pdTRUE if operation was successful else pdFALSE
 */
//go:linkname VTaskGetSnapshot C.vTaskGetSnapshot
func VTaskGetSnapshot(pxTask TaskHandleT, pxTaskSnapshot *TaskSnapshotT) BaseTypeT

/**
 * @brief Fill an array of TaskSnapshot_t structures for every task in the system
 *
 * - This function is used by the panic handler to get a snapshot of all tasks in the system
 *
 * @note This function should only be called when FreeRTOS is no longer running (e.g., during a panic) as this function
 *        does not acquire any locks.
 * @param[out] pxTaskSnapshotArray Array of TaskSnapshot_t structures filled by this function
 * @param[in] uxArrayLength Length of the provided array
 * @param[out] pxTCBSize Size of the a task's TCB structure (can be set to NULL)
 * @return UBaseType_t
 */
// llgo:link (*TaskSnapshotT).UxTaskGetSnapshotAll C.uxTaskGetSnapshotAll
func (recv_ *TaskSnapshotT) UxTaskGetSnapshotAll(uxArrayLength UBaseTypeT, pxTCBSize *UBaseTypeT) UBaseTypeT {
	return 0
}

/**
 * @brief Get a void pointer to the current TCB of a particular core
 *
 * @note This function provides no guarantee that the return TCB will still be the current task (or that the task still
 * exists) when it returns. It is the caller's responsibility to ensure that the task does not get scheduled or deleted.
 * @param xCoreID The core to query
 * @return Void pointer to current TCB
 */
// llgo:link BaseTypeT.PvTaskGetCurrentTCBForCore C.pvTaskGetCurrentTCBForCore
func (recv_ BaseTypeT) PvTaskGetCurrentTCBForCore() c.Pointer {
	return nil
}
