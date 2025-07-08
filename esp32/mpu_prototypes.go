package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/* MPU versions of task.h API functions. */
//go:linkname MPUXTaskCreate C.MPU_xTaskCreate
func MPUXTaskCreate(pxTaskCode c.Int, pcName *c.Char, usStackDepth c.Uint16T, pvParameters c.Pointer, uxPriority UBaseTypeT, pxCreatedTask *c.Int) BaseTypeT

//go:linkname MPUXTaskCreateStatic C.MPU_xTaskCreateStatic
func MPUXTaskCreateStatic(pxTaskCode c.Int, pcName *c.Char, ulStackDepth c.Uint32T, pvParameters c.Pointer, uxPriority UBaseTypeT, puxStackBuffer *StackTypeT, pxTaskBuffer *StaticTaskT) c.Int

//go:linkname MPUVTaskDelete C.MPU_vTaskDelete
func MPUVTaskDelete(xTaskToDelete c.Int)

// llgo:link TickTypeT.MPUVTaskDelay C.MPU_vTaskDelay
func (recv_ TickTypeT) MPUVTaskDelay() {
}

// llgo:link (*TickTypeT).MPUXTaskDelayUntil C.MPU_xTaskDelayUntil
func (recv_ *TickTypeT) MPUXTaskDelayUntil(xTimeIncrement TickTypeT) BaseTypeT {
	return 0
}

//go:linkname MPUXTaskAbortDelay C.MPU_xTaskAbortDelay
func MPUXTaskAbortDelay(xTask c.Int) BaseTypeT

//go:linkname MPUUxTaskPriorityGet C.MPU_uxTaskPriorityGet
func MPUUxTaskPriorityGet(xTask c.Int) UBaseTypeT

//go:linkname MPUETaskGetState C.MPU_eTaskGetState
func MPUETaskGetState(xTask c.Int) c.Int

//go:linkname MPUVTaskGetInfo C.MPU_vTaskGetInfo
func MPUVTaskGetInfo(xTask c.Int, pxTaskStatus *c.Int, xGetFreeStackSpace BaseTypeT, eState c.Int)

//go:linkname MPUVTaskPrioritySet C.MPU_vTaskPrioritySet
func MPUVTaskPrioritySet(xTask c.Int, uxNewPriority UBaseTypeT)

//go:linkname MPUVTaskSuspend C.MPU_vTaskSuspend
func MPUVTaskSuspend(xTaskToSuspend c.Int)

//go:linkname MPUVTaskResume C.MPU_vTaskResume
func MPUVTaskResume(xTaskToResume c.Int)

//go:linkname MPUVTaskStartScheduler C.MPU_vTaskStartScheduler
func MPUVTaskStartScheduler()

//go:linkname MPUVTaskSuspendAll C.MPU_vTaskSuspendAll
func MPUVTaskSuspendAll()

//go:linkname MPUXTaskResumeAll C.MPU_xTaskResumeAll
func MPUXTaskResumeAll() BaseTypeT

//go:linkname MPUXTaskGetTickCount C.MPU_xTaskGetTickCount
func MPUXTaskGetTickCount() TickTypeT

//go:linkname MPUUxTaskGetNumberOfTasks C.MPU_uxTaskGetNumberOfTasks
func MPUUxTaskGetNumberOfTasks() UBaseTypeT

//go:linkname MPUPcTaskGetName C.MPU_pcTaskGetName
func MPUPcTaskGetName(xTaskToQuery c.Int) *c.Char

//go:linkname MPUXTaskGetHandle C.MPU_xTaskGetHandle
func MPUXTaskGetHandle(pcNameToQuery *c.Char) c.Int

//go:linkname MPUUxTaskGetStackHighWaterMark C.MPU_uxTaskGetStackHighWaterMark
func MPUUxTaskGetStackHighWaterMark(xTask c.Int) UBaseTypeT

//go:linkname MPUUxTaskGetStackHighWaterMark2 C.MPU_uxTaskGetStackHighWaterMark2
func MPUUxTaskGetStackHighWaterMark2(xTask c.Int) c.Uint32T

//go:linkname MPUVTaskSetApplicationTaskTag C.MPU_vTaskSetApplicationTaskTag
func MPUVTaskSetApplicationTaskTag(xTask c.Int, pxHookFunction c.Int)

//go:linkname MPUXTaskGetApplicationTaskTag C.MPU_xTaskGetApplicationTaskTag
func MPUXTaskGetApplicationTaskTag(xTask c.Int) c.Int

//go:linkname MPUVTaskSetThreadLocalStoragePointer C.MPU_vTaskSetThreadLocalStoragePointer
func MPUVTaskSetThreadLocalStoragePointer(xTaskToSet c.Int, xIndex BaseTypeT, pvValue c.Pointer)

//go:linkname MPUPvTaskGetThreadLocalStoragePointer C.MPU_pvTaskGetThreadLocalStoragePointer
func MPUPvTaskGetThreadLocalStoragePointer(xTaskToQuery c.Int, xIndex BaseTypeT) c.Pointer

//go:linkname MPUXTaskCallApplicationTaskHook C.MPU_xTaskCallApplicationTaskHook
func MPUXTaskCallApplicationTaskHook(xTask c.Int, pvParameter c.Pointer) BaseTypeT

//go:linkname MPUXTaskGetIdleTaskHandle C.MPU_xTaskGetIdleTaskHandle
func MPUXTaskGetIdleTaskHandle() c.Int

//go:linkname MPUUxTaskGetSystemState C.MPU_uxTaskGetSystemState
func MPUUxTaskGetSystemState(pxTaskStatusArray *c.Int, uxArraySize UBaseTypeT, pulTotalRunTime *c.Uint32T) UBaseTypeT

//go:linkname MPUUlTaskGetIdleRunTimeCounter C.MPU_ulTaskGetIdleRunTimeCounter
func MPUUlTaskGetIdleRunTimeCounter() c.Uint32T

//go:linkname MPUUlTaskGetIdleRunTimePercent C.MPU_ulTaskGetIdleRunTimePercent
func MPUUlTaskGetIdleRunTimePercent() c.Uint32T

//go:linkname MPUVTaskList C.MPU_vTaskList
func MPUVTaskList(pcWriteBuffer *c.Char)

//go:linkname MPUVTaskGetRunTimeStats C.MPU_vTaskGetRunTimeStats
func MPUVTaskGetRunTimeStats(pcWriteBuffer *c.Char)

//go:linkname MPUXTaskGenericNotify C.MPU_xTaskGenericNotify
func MPUXTaskGenericNotify(xTaskToNotify c.Int, uxIndexToNotify UBaseTypeT, ulValue c.Uint32T, eAction c.Int, pulPreviousNotificationValue *c.Uint32T) BaseTypeT

// llgo:link UBaseTypeT.MPUXTaskGenericNotifyWait C.MPU_xTaskGenericNotifyWait
func (recv_ UBaseTypeT) MPUXTaskGenericNotifyWait(ulBitsToClearOnEntry c.Uint32T, ulBitsToClearOnExit c.Uint32T, pulNotificationValue *c.Uint32T, xTicksToWait TickTypeT) BaseTypeT {
	return 0
}

// llgo:link UBaseTypeT.MPUUlTaskGenericNotifyTake C.MPU_ulTaskGenericNotifyTake
func (recv_ UBaseTypeT) MPUUlTaskGenericNotifyTake(xClearCountOnExit BaseTypeT, xTicksToWait TickTypeT) c.Uint32T {
	return 0
}

//go:linkname MPUXTaskGenericNotifyStateClear C.MPU_xTaskGenericNotifyStateClear
func MPUXTaskGenericNotifyStateClear(xTask c.Int, uxIndexToClear UBaseTypeT) BaseTypeT

//go:linkname MPUUlTaskGenericNotifyValueClear C.MPU_ulTaskGenericNotifyValueClear
func MPUUlTaskGenericNotifyValueClear(xTask c.Int, uxIndexToClear UBaseTypeT, ulBitsToClear c.Uint32T) c.Uint32T

//go:linkname MPUXTaskIncrementTick C.MPU_xTaskIncrementTick
func MPUXTaskIncrementTick() BaseTypeT

//go:linkname MPUXTaskGetCurrentTaskHandle C.MPU_xTaskGetCurrentTaskHandle
func MPUXTaskGetCurrentTaskHandle() c.Int

//go:linkname MPUVTaskSetTimeOutState C.MPU_vTaskSetTimeOutState
func MPUVTaskSetTimeOutState(pxTimeOut *c.Int)

//go:linkname MPUXTaskCheckForTimeOut C.MPU_xTaskCheckForTimeOut
func MPUXTaskCheckForTimeOut(pxTimeOut *c.Int, pxTicksToWait *TickTypeT) BaseTypeT

//go:linkname MPUVTaskMissedYield C.MPU_vTaskMissedYield
func MPUVTaskMissedYield()

//go:linkname MPUXTaskGetSchedulerState C.MPU_xTaskGetSchedulerState
func MPUXTaskGetSchedulerState() BaseTypeT

// llgo:link TickTypeT.MPUXTaskCatchUpTicks C.MPU_xTaskCatchUpTicks
func (recv_ TickTypeT) MPUXTaskCatchUpTicks() BaseTypeT {
	return 0
}

/* MPU versions of queue.h API functions. */
//go:linkname MPUXQueueGenericSend C.MPU_xQueueGenericSend
func MPUXQueueGenericSend(xQueue c.Int, pvItemToQueue c.Pointer, xTicksToWait TickTypeT, xCopyPosition BaseTypeT) BaseTypeT

//go:linkname MPUXQueueReceive C.MPU_xQueueReceive
func MPUXQueueReceive(xQueue c.Int, pvBuffer c.Pointer, xTicksToWait TickTypeT) BaseTypeT

//go:linkname MPUXQueuePeek C.MPU_xQueuePeek
func MPUXQueuePeek(xQueue c.Int, pvBuffer c.Pointer, xTicksToWait TickTypeT) BaseTypeT

//go:linkname MPUXQueueSemaphoreTake C.MPU_xQueueSemaphoreTake
func MPUXQueueSemaphoreTake(xQueue c.Int, xTicksToWait TickTypeT) BaseTypeT

//go:linkname MPUUxQueueMessagesWaiting C.MPU_uxQueueMessagesWaiting
func MPUUxQueueMessagesWaiting(xQueue c.Int) UBaseTypeT

//go:linkname MPUUxQueueSpacesAvailable C.MPU_uxQueueSpacesAvailable
func MPUUxQueueSpacesAvailable(xQueue c.Int) UBaseTypeT

//go:linkname MPUVQueueDelete C.MPU_vQueueDelete
func MPUVQueueDelete(xQueue c.Int)

//go:linkname MPUXQueueCreateMutex C.MPU_xQueueCreateMutex
func MPUXQueueCreateMutex(ucQueueType c.Uint8T) c.Int

//go:linkname MPUXQueueCreateMutexStatic C.MPU_xQueueCreateMutexStatic
func MPUXQueueCreateMutexStatic(ucQueueType c.Uint8T, pxStaticQueue *StaticQueueT) c.Int

// llgo:link UBaseTypeT.MPUXQueueCreateCountingSemaphore C.MPU_xQueueCreateCountingSemaphore
func (recv_ UBaseTypeT) MPUXQueueCreateCountingSemaphore(uxInitialCount UBaseTypeT) c.Int {
	return 0
}

// llgo:link UBaseTypeT.MPUXQueueCreateCountingSemaphoreStatic C.MPU_xQueueCreateCountingSemaphoreStatic
func (recv_ UBaseTypeT) MPUXQueueCreateCountingSemaphoreStatic(uxInitialCount UBaseTypeT, pxStaticQueue *StaticQueueT) c.Int {
	return 0
}

//go:linkname MPUXQueueGetMutexHolder C.MPU_xQueueGetMutexHolder
func MPUXQueueGetMutexHolder(xSemaphore c.Int) c.Int

//go:linkname MPUXQueueTakeMutexRecursive C.MPU_xQueueTakeMutexRecursive
func MPUXQueueTakeMutexRecursive(xMutex c.Int, xTicksToWait TickTypeT) BaseTypeT

//go:linkname MPUXQueueGiveMutexRecursive C.MPU_xQueueGiveMutexRecursive
func MPUXQueueGiveMutexRecursive(pxMutex c.Int) BaseTypeT

//go:linkname MPUVQueueAddToRegistry C.MPU_vQueueAddToRegistry
func MPUVQueueAddToRegistry(xQueue c.Int, pcName *c.Char)

//go:linkname MPUVQueueUnregisterQueue C.MPU_vQueueUnregisterQueue
func MPUVQueueUnregisterQueue(xQueue c.Int)

//go:linkname MPUPcQueueGetName C.MPU_pcQueueGetName
func MPUPcQueueGetName(xQueue c.Int) *c.Char

// llgo:link UBaseTypeT.MPUXQueueGenericCreate C.MPU_xQueueGenericCreate
func (recv_ UBaseTypeT) MPUXQueueGenericCreate(uxItemSize UBaseTypeT, ucQueueType c.Uint8T) c.Int {
	return 0
}

// llgo:link UBaseTypeT.MPUXQueueGenericCreateStatic C.MPU_xQueueGenericCreateStatic
func (recv_ UBaseTypeT) MPUXQueueGenericCreateStatic(uxItemSize UBaseTypeT, pucQueueStorage *c.Uint8T, pxStaticQueue *StaticQueueT, ucQueueType c.Uint8T) c.Int {
	return 0
}

// llgo:link UBaseTypeT.MPUXQueueCreateSet C.MPU_xQueueCreateSet
func (recv_ UBaseTypeT) MPUXQueueCreateSet() c.Int {
	return 0
}

//go:linkname MPUXQueueAddToSet C.MPU_xQueueAddToSet
func MPUXQueueAddToSet(xQueueOrSemaphore c.Int, xQueueSet c.Int) BaseTypeT

//go:linkname MPUXQueueRemoveFromSet C.MPU_xQueueRemoveFromSet
func MPUXQueueRemoveFromSet(xQueueOrSemaphore c.Int, xQueueSet c.Int) BaseTypeT

//go:linkname MPUXQueueSelectFromSet C.MPU_xQueueSelectFromSet
func MPUXQueueSelectFromSet(xQueueSet c.Int, xTicksToWait TickTypeT) c.Int

//go:linkname MPUXQueueGenericReset C.MPU_xQueueGenericReset
func MPUXQueueGenericReset(xQueue c.Int, xNewQueue BaseTypeT) BaseTypeT

//go:linkname MPUVQueueSetQueueNumber C.MPU_vQueueSetQueueNumber
func MPUVQueueSetQueueNumber(xQueue c.Int, uxQueueNumber UBaseTypeT)

//go:linkname MPUUxQueueGetQueueNumber C.MPU_uxQueueGetQueueNumber
func MPUUxQueueGetQueueNumber(xQueue c.Int) UBaseTypeT

//go:linkname MPUUcQueueGetQueueType C.MPU_ucQueueGetQueueType
func MPUUcQueueGetQueueType(xQueue c.Int) c.Uint8T

/* MPU versions of timers.h API functions. */
//go:linkname MPUXTimerCreate C.MPU_xTimerCreate
func MPUXTimerCreate(pcTimerName *c.Char, xTimerPeriodInTicks TickTypeT, uxAutoReload UBaseTypeT, pvTimerID c.Pointer, pxCallbackFunction c.Int) c.Int

//go:linkname MPUXTimerCreateStatic C.MPU_xTimerCreateStatic
func MPUXTimerCreateStatic(pcTimerName *c.Char, xTimerPeriodInTicks TickTypeT, uxAutoReload UBaseTypeT, pvTimerID c.Pointer, pxCallbackFunction c.Int, pxTimerBuffer *StaticTimerT) c.Int

//go:linkname MPUPvTimerGetTimerID C.MPU_pvTimerGetTimerID
func MPUPvTimerGetTimerID(xTimer c.Int) c.Pointer

//go:linkname MPUVTimerSetTimerID C.MPU_vTimerSetTimerID
func MPUVTimerSetTimerID(xTimer c.Int, pvNewID c.Pointer)

//go:linkname MPUXTimerIsTimerActive C.MPU_xTimerIsTimerActive
func MPUXTimerIsTimerActive(xTimer c.Int) BaseTypeT

//go:linkname MPUXTimerGetTimerDaemonTaskHandle C.MPU_xTimerGetTimerDaemonTaskHandle
func MPUXTimerGetTimerDaemonTaskHandle() c.Int

//go:linkname MPUXTimerPendFunctionCall C.MPU_xTimerPendFunctionCall
func MPUXTimerPendFunctionCall(xFunctionToPend c.Int, pvParameter1 c.Pointer, ulParameter2 c.Uint32T, xTicksToWait TickTypeT) BaseTypeT

//go:linkname MPUPcTimerGetName C.MPU_pcTimerGetName
func MPUPcTimerGetName(xTimer c.Int) *c.Char

//go:linkname MPUVTimerSetReloadMode C.MPU_vTimerSetReloadMode
func MPUVTimerSetReloadMode(xTimer c.Int, uxAutoReload UBaseTypeT)

//go:linkname MPUUxTimerGetReloadMode C.MPU_uxTimerGetReloadMode
func MPUUxTimerGetReloadMode(xTimer c.Int) UBaseTypeT

//go:linkname MPUXTimerGetPeriod C.MPU_xTimerGetPeriod
func MPUXTimerGetPeriod(xTimer c.Int) TickTypeT

//go:linkname MPUXTimerGetExpiryTime C.MPU_xTimerGetExpiryTime
func MPUXTimerGetExpiryTime(xTimer c.Int) TickTypeT

//go:linkname MPUXTimerCreateTimerTask C.MPU_xTimerCreateTimerTask
func MPUXTimerCreateTimerTask() BaseTypeT

//go:linkname MPUXTimerGenericCommand C.MPU_xTimerGenericCommand
func MPUXTimerGenericCommand(xTimer c.Int, xCommandID BaseTypeT, xOptionalValue TickTypeT, pxHigherPriorityTaskWoken *BaseTypeT, xTicksToWait TickTypeT) BaseTypeT

/* MPU versions of event_group.h API functions. */
//go:linkname MPUXEventGroupCreate C.MPU_xEventGroupCreate
func MPUXEventGroupCreate() c.Int

// llgo:link (*StaticEventGroupT).MPUXEventGroupCreateStatic C.MPU_xEventGroupCreateStatic
func (recv_ *StaticEventGroupT) MPUXEventGroupCreateStatic() c.Int {
	return 0
}

//go:linkname MPUXEventGroupWaitBits C.MPU_xEventGroupWaitBits
func MPUXEventGroupWaitBits(xEventGroup c.Int, uxBitsToWaitFor c.Int, xClearOnExit BaseTypeT, xWaitForAllBits BaseTypeT, xTicksToWait TickTypeT) c.Int

//go:linkname MPUXEventGroupClearBits C.MPU_xEventGroupClearBits
func MPUXEventGroupClearBits(xEventGroup c.Int, uxBitsToClear c.Int) c.Int

//go:linkname MPUXEventGroupSetBits C.MPU_xEventGroupSetBits
func MPUXEventGroupSetBits(xEventGroup c.Int, uxBitsToSet c.Int) c.Int

//go:linkname MPUXEventGroupSync C.MPU_xEventGroupSync
func MPUXEventGroupSync(xEventGroup c.Int, uxBitsToSet c.Int, uxBitsToWaitFor c.Int, xTicksToWait TickTypeT) c.Int

//go:linkname MPUVEventGroupDelete C.MPU_vEventGroupDelete
func MPUVEventGroupDelete(xEventGroup c.Int)

//go:linkname MPUUxEventGroupGetNumber C.MPU_uxEventGroupGetNumber
func MPUUxEventGroupGetNumber(xEventGroup c.Pointer) UBaseTypeT

/* MPU versions of message/stream_buffer.h API functions. */
//go:linkname MPUXStreamBufferSend C.MPU_xStreamBufferSend
func MPUXStreamBufferSend(xStreamBuffer c.Int, pvTxData c.Pointer, xDataLengthBytes c.SizeT, xTicksToWait TickTypeT) c.SizeT

//go:linkname MPUXStreamBufferReceive C.MPU_xStreamBufferReceive
func MPUXStreamBufferReceive(xStreamBuffer c.Int, pvRxData c.Pointer, xBufferLengthBytes c.SizeT, xTicksToWait TickTypeT) c.SizeT

//go:linkname MPUXStreamBufferNextMessageLengthBytes C.MPU_xStreamBufferNextMessageLengthBytes
func MPUXStreamBufferNextMessageLengthBytes(xStreamBuffer c.Int) c.SizeT

//go:linkname MPUVStreamBufferDelete C.MPU_vStreamBufferDelete
func MPUVStreamBufferDelete(xStreamBuffer c.Int)

//go:linkname MPUXStreamBufferIsFull C.MPU_xStreamBufferIsFull
func MPUXStreamBufferIsFull(xStreamBuffer c.Int) BaseTypeT

//go:linkname MPUXStreamBufferIsEmpty C.MPU_xStreamBufferIsEmpty
func MPUXStreamBufferIsEmpty(xStreamBuffer c.Int) BaseTypeT

//go:linkname MPUXStreamBufferReset C.MPU_xStreamBufferReset
func MPUXStreamBufferReset(xStreamBuffer c.Int) BaseTypeT

//go:linkname MPUXStreamBufferSpacesAvailable C.MPU_xStreamBufferSpacesAvailable
func MPUXStreamBufferSpacesAvailable(xStreamBuffer c.Int) c.SizeT

//go:linkname MPUXStreamBufferBytesAvailable C.MPU_xStreamBufferBytesAvailable
func MPUXStreamBufferBytesAvailable(xStreamBuffer c.Int) c.SizeT

//go:linkname MPUXStreamBufferSetTriggerLevel C.MPU_xStreamBufferSetTriggerLevel
func MPUXStreamBufferSetTriggerLevel(xStreamBuffer c.Int, xTriggerLevel c.SizeT) BaseTypeT

//go:linkname MPUXStreamBufferGenericCreate C.MPU_xStreamBufferGenericCreate
func MPUXStreamBufferGenericCreate(xBufferSizeBytes c.SizeT, xTriggerLevelBytes c.SizeT, xIsMessageBuffer BaseTypeT, pxSendCompletedCallback c.Int, pxReceiveCompletedCallback c.Int) c.Int

//go:linkname MPUXStreamBufferGenericCreateStatic C.MPU_xStreamBufferGenericCreateStatic
func MPUXStreamBufferGenericCreateStatic(xBufferSizeBytes c.SizeT, xTriggerLevelBytes c.SizeT, xIsMessageBuffer BaseTypeT, pucStreamBufferStorageArea *c.Uint8T, pxStaticStreamBuffer *StaticStreamBufferT, pxSendCompletedCallback c.Int, pxReceiveCompletedCallback c.Int) c.Int
