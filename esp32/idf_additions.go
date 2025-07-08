package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Create a new task that is pinned to a particular core
 *
 * This function is similar to xTaskCreate(), but allows the creation of a pinned
 * task. The task's pinned core is specified by the xCoreID argument. If xCoreID
 * is set to tskNO_AFFINITY, then the task is unpinned and can run on any core.
 *
 * @note If ( configNUMBER_OF_CORES == 1 ), setting xCoreID to tskNO_AFFINITY will be
 * be treated as 0.
 *
 * @param pxTaskCode Pointer to the task entry function.
 * @param pcName A descriptive name for the task.
 * @param ulStackDepth The size of the task stack specified as the NUMBER OF
 * BYTES. Note that this differs from vanilla FreeRTOS.
 * @param pvParameters Pointer that will be used as the parameter for the task
 * being created.
 * @param uxPriority The priority at which the task should run.
 * @param pxCreatedTask Used to pass back a handle by which the created task can
 * be referenced.
 * @param xCoreID The core to which the task is pinned to, or tskNO_AFFINITY if
 * the task has no core affinity.
 * @return pdPASS if the task was successfully created and added to a ready
 * list, otherwise an error code defined in the file projdefs.h
 */
//go:linkname XTaskCreatePinnedToCore C.xTaskCreatePinnedToCore
func XTaskCreatePinnedToCore(pxTaskCode c.Int, pcName *c.Char, ulStackDepth c.Uint32T, pvParameters c.Pointer, uxPriority UBaseTypeT, pxCreatedTask *TaskHandleT, xCoreID BaseTypeT) BaseTypeT

/**
 * @brief Create a new static task that is pinned to a particular core
 *
 * This function is similar to xTaskCreateStatic(), but allows the creation of a
 * pinned task. The task's pinned core is specified by the xCoreID argument. If
 * xCoreID is set to tskNO_AFFINITY, then the task is unpinned and can run on any
 * core.
 *
 * @note If ( configNUMBER_OF_CORES == 1 ), setting xCoreID to tskNO_AFFINITY will be
 * be treated as 0.
 *
 * @param pxTaskCode Pointer to the task entry function.
 * @param pcName A descriptive name for the task.
 * @param ulStackDepth The size of the task stack specified as the NUMBER OF
 * BYTES. Note that this differs from vanilla FreeRTOS.
 * @param pvParameters Pointer that will be used as the parameter for the task
 * being created.
 * @param uxPriority The priority at which the task should run.
 * @param puxStackBuffer Must point to a StackType_t array that has at least
 * ulStackDepth indexes
 * @param pxTaskBuffer Must point to a variable of type StaticTask_t, which will
 * then be used to hold the task's data structures,
 * @param xCoreID The core to which the task is pinned to, or tskNO_AFFINITY if
 * the task has no core affinity.
 * @return The task handle if the task was created, NULL otherwise.
 */
//go:linkname XTaskCreateStaticPinnedToCore C.xTaskCreateStaticPinnedToCore
func XTaskCreateStaticPinnedToCore(pxTaskCode c.Int, pcName *c.Char, ulStackDepth c.Uint32T, pvParameters c.Pointer, uxPriority UBaseTypeT, puxStackBuffer *StackTypeT, pxTaskBuffer *StaticTaskT, xCoreID BaseTypeT) TaskHandleT

/**
 * @brief Get the current core ID of a particular task
 *
 * Helper function to get the core ID of a particular task. If the task is
 * pinned to a particular core, the core ID is returned. If the task is not
 * pinned to a particular core, tskNO_AFFINITY is returned.
 *
 * If CONFIG_FREERTOS_UNICORE is enabled, this function simply returns 0.
 *
 * [refactor-todo] See if this needs to be deprecated (IDF-8145)(IDF-8164)
 *
 * @note If CONFIG_FREERTOS_SMP is enabled, please call vTaskCoreAffinityGet()
 * instead.
 * @note In IDF FreerTOS when configNUMBER_OF_CORES == 1, this function will
 * always return 0,
 * @param xTask The task to query
 * @return The task's core ID or tskNO_AFFINITY
 */
//go:linkname XTaskGetCoreID C.xTaskGetCoreID
func XTaskGetCoreID(xTask TaskHandleT) BaseTypeT

/**
 * @brief Get the handle of idle task for the given core.
 *
 * [refactor-todo] See if this needs to be deprecated (IDF-8145)
 *
 * @param xCoreID The core to query
 * @return Handle of the idle task for the queried core
 */
// llgo:link BaseTypeT.XTaskGetIdleTaskHandleForCore C.xTaskGetIdleTaskHandleForCore
func (recv_ BaseTypeT) XTaskGetIdleTaskHandleForCore() TaskHandleT {
	return nil
}

/**
 * @brief Get the handle of the task currently running on a certain core
 *
 * Because of the nature of SMP processing, there is no guarantee that this
 * value will still be valid on return and should only be used for debugging
 * purposes.
 *
 * [refactor-todo] See if this needs to be deprecated (IDF-8145)
 *
 * @param xCoreID The core to query
 * @return Handle of the current task running on the queried core
 */
// llgo:link BaseTypeT.XTaskGetCurrentTaskHandleForCore C.xTaskGetCurrentTaskHandleForCore
func (recv_ BaseTypeT) XTaskGetCurrentTaskHandleForCore() TaskHandleT {
	return nil
}

/**
 * Returns the start of the stack associated with xTask.
 *
 * Returns the lowest stack memory address, regardless of whether the stack
 * grows up or down.
 *
 * [refactor-todo] Change return type to StackType_t (IDF-8158)
 *
 * @param xTask Handle of the task associated with the stack returned.
 * Set xTask to NULL to return the stack of the calling task.
 *
 * @return A pointer to the start of the stack.
 */
//go:linkname PxTaskGetStackStart C.pxTaskGetStackStart
func PxTaskGetStackStart(xTask TaskHandleT) *c.Uint8T

// llgo:type C
type TlsDeleteCallbackFunctionT func(c.Int, c.Pointer)

/**
 * Set local storage pointer and deletion callback.
 *
 * Each task contains an array of pointers that is dimensioned by the
 * configNUM_THREAD_LOCAL_STORAGE_POINTERS setting in FreeRTOSConfig.h. The
 * kernel does not use the pointers itself, so the application writer can use
 * the pointers for any purpose they wish.
 *
 * Local storage pointers set for a task can reference dynamically allocated
 * resources. This function is similar to vTaskSetThreadLocalStoragePointer, but
 * provides a way to release these resources when the task gets deleted. For
 * each pointer, a callback function can be set. This function will be called
 * when task is deleted, with the local storage pointer index and value as
 * arguments.
 *
 * @param xTaskToSet  Task to set thread local storage pointer for
 * @param xIndex The index of the pointer to set, from 0 to
 * configNUM_THREAD_LOCAL_STORAGE_POINTERS - 1.
 * @param pvValue Pointer value to set.
 * @param pvDelCallback  Function to call to dispose of the local storage
 * pointer when the task is deleted.
 */
//go:linkname VTaskSetThreadLocalStoragePointerAndDelCallback C.vTaskSetThreadLocalStoragePointerAndDelCallback
func VTaskSetThreadLocalStoragePointerAndDelCallback(xTaskToSet TaskHandleT, xIndex BaseTypeT, pvValue c.Pointer, pvDelCallback TlsDeleteCallbackFunctionT)

/**
 * @brief Creates a pinned task where its stack has specific memory capabilities
 *
 * This function is similar to xTaskCreatePinnedToCore(), except that it allows
 * the memory allocated for the task's stack to have specific capabilities
 * (e.g., MALLOC_CAP_SPIRAM).
 *
 * However, the specified capabilities will NOT apply to the task's TCB as a TCB
 * must always be in internal RAM.
 *
 * @param pvTaskCode Pointer to the task entry function
 * @param pcName A descriptive name for the task
 * @param usStackDepth The size of the task stack specified as the number of
 * bytes
 * @param pvParameters Pointer that will be used as the parameter for the task
 * being created.
 * @param uxPriority The priority at which the task should run.
 * @param pvCreatedTask Used to pass back a handle by which the created task can
 * be referenced.
 * @param xCoreID Core to which the task is pinned to, or tskNO_AFFINITY if
 * unpinned.
 * @param uxMemoryCaps Memory capabilities of the task stack's memory (see
 * esp_heap_caps.h)
 * @return pdPASS if the task was successfully created and added to a ready
 * list, otherwise an error code defined in the file projdefs.h
 */
//go:linkname XTaskCreatePinnedToCoreWithCaps C.xTaskCreatePinnedToCoreWithCaps
func XTaskCreatePinnedToCoreWithCaps(pvTaskCode c.Int, pcName *c.Char, usStackDepth c.Uint32T, pvParameters c.Pointer, uxPriority UBaseTypeT, pvCreatedTask *TaskHandleT, xCoreID BaseTypeT, uxMemoryCaps UBaseTypeT) BaseTypeT

/**
 * @brief Deletes a task previously created using xTaskCreateWithCaps() or
 * xTaskCreatePinnedToCoreWithCaps()
 *
 * @note It is recommended to use this API to delete tasks from another task's
 * context, rather than self-deletion. When the task is being deleted, it is vital
 * to ensure that it is not running on another core. This API must not be called
 * from an interrupt context.
 *
 * @param xTaskToDelete A handle to the task to be deleted
 */
//go:linkname VTaskDeleteWithCaps C.vTaskDeleteWithCaps
func VTaskDeleteWithCaps(xTaskToDelete TaskHandleT)

/**
 * @brief Creates a queue with specific memory capabilities
 *
 * This function is similar to xQueueCreate(), except that it allows the memory
 * allocated for the queue to have specific capabilities (e.g.,
 * MALLOC_CAP_INTERNAL).
 *
 * @note A queue created using this function must only be deleted using
 * vQueueDeleteWithCaps()
 * @param uxQueueLength The maximum number of items that the queue can contain.
 * @param uxItemSize The number of bytes each item in the queue will require.
 * @param uxMemoryCaps Memory capabilities of the queue's memory (see
 * esp_heap_caps.h)
 * @return Handle to the created queue or NULL on failure.
 */
// llgo:link UBaseTypeT.XQueueCreateWithCaps C.xQueueCreateWithCaps
func (recv_ UBaseTypeT) XQueueCreateWithCaps(uxItemSize UBaseTypeT, uxMemoryCaps UBaseTypeT) QueueHandleT {
	return nil
}

/**
 * @brief Deletes a queue previously created using xQueueCreateWithCaps()
 *
 * @param xQueue A handle to the queue to be deleted.
 */
//go:linkname VQueueDeleteWithCaps C.vQueueDeleteWithCaps
func VQueueDeleteWithCaps(xQueue QueueHandleT)

// llgo:link UBaseTypeT.XSemaphoreCreateGenericWithCaps C.xSemaphoreCreateGenericWithCaps
func (recv_ UBaseTypeT) XSemaphoreCreateGenericWithCaps(uxInitialCount UBaseTypeT, ucQueueType c.Uint8T, uxMemoryCaps UBaseTypeT) SemaphoreHandleT {
	return nil
}

/**
 * @brief Deletes a semaphore previously created using one of the
 * xSemaphoreCreate...WithCaps() functions
 *
 * @param xSemaphore A handle to the semaphore to be deleted.
 */
//go:linkname VSemaphoreDeleteWithCaps C.vSemaphoreDeleteWithCaps
func VSemaphoreDeleteWithCaps(xSemaphore SemaphoreHandleT)

//go:linkname XStreamBufferGenericCreateWithCaps C.xStreamBufferGenericCreateWithCaps
func XStreamBufferGenericCreateWithCaps(xBufferSizeBytes c.SizeT, xTriggerLevelBytes c.SizeT, xIsMessageBuffer BaseTypeT, uxMemoryCaps UBaseTypeT) StreamBufferHandleT

//go:linkname VStreamBufferGenericDeleteWithCaps C.vStreamBufferGenericDeleteWithCaps
func VStreamBufferGenericDeleteWithCaps(xStreamBuffer StreamBufferHandleT, xIsMessageBuffer BaseTypeT)

/**
 * @brief Creates an event group with specific memory capabilities
 *
 * This function is similar to xEventGroupCreate(), except that it allows the
 * memory allocated for the event group to have specific capabilities (e.g.,
 * MALLOC_CAP_INTERNAL).
 *
 * @note An event group created using this function must only be deleted using
 * vEventGroupDeleteWithCaps()
 * @param uxMemoryCaps Memory capabilities of the event group's memory (see
 * esp_heap_caps.h)
 * @return Handle to the created event group or NULL on failure.
 */
// llgo:link UBaseTypeT.XEventGroupCreateWithCaps C.xEventGroupCreateWithCaps
func (recv_ UBaseTypeT) XEventGroupCreateWithCaps() EventGroupHandleT {
	return nil
}

/**
 * @brief Deletes an event group previously created using
 * xEventGroupCreateWithCaps()
 *
 * @param xEventGroup A handle to the event group to be deleted.
 */
//go:linkname VEventGroupDeleteWithCaps C.vEventGroupDeleteWithCaps
func VEventGroupDeleteWithCaps(xEventGroup EventGroupHandleT)
