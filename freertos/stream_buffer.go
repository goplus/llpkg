package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Type by which stream buffers are referenced.  For example, a call to
 * xStreamBufferCreate() returns an StreamBufferHandle_t variable that can
 * then be used as a parameter to xStreamBufferSend(), xStreamBufferReceive(),
 * etc.
 */

type StreamBufferDefT struct {
	Unused [8]uint8
}
type StreamBufferHandleT *StreamBufferDefT

// llgo:type C
type StreamBufferCallbackFunctionT func(StreamBufferHandleT, BaseTypeT, *BaseTypeT)

/**
 *
 * Retrieve pointers to a statically created stream buffer's data structure
 * buffer and storage area buffer. These are the same buffers that are supplied
 * at the time of creation.
 *
 * @param xStreamBuffer The stream buffer for which to retrieve the buffers.
 *
 * @param ppucStreamBufferStorageArea Used to return a pointer to the stream
 * buffer's storage area buffer.
 *
 * @param ppxStaticStreamBuffer Used to return a pointer to the stream
 * buffer's data structure buffer.
 *
 * @return pdTRUE if buffers were retrieved, pdFALSE otherwise.
 *
 * \ingroup StreamBufferManagement
 */
//go:linkname XStreamBufferGetStaticBuffers C.xStreamBufferGetStaticBuffers
func XStreamBufferGetStaticBuffers(xStreamBuffer StreamBufferHandleT, ppucStreamBufferStorageArea **c.Uint8T, ppxStaticStreamBuffer **StaticStreamBufferT) BaseTypeT

/**
 *
 * Sends bytes to a stream buffer.  The bytes are copied into the stream buffer.
 *
 * ***NOTE***:  Uniquely among FreeRTOS objects, the stream buffer
 * implementation (so also the message buffer implementation, as message buffers
 * are built on top of stream buffers) assumes there is only one task or
 * interrupt that will write to the buffer (the writer), and only one task or
 * interrupt that will read from the buffer (the reader).  It is safe for the
 * writer and reader to be different tasks or interrupts, but, unlike other
 * FreeRTOS objects, it is not safe to have multiple different writers or
 * multiple different readers.  If there are to be multiple different writers
 * then the application writer must place each call to a writing API function
 * (such as xStreamBufferSend()) inside a critical section and set the send
 * block time to 0.  Likewise, if there are to be multiple different readers
 * then the application writer must place each call to a reading API function
 * (such as xStreamBufferReceive()) inside a critical section and set the receive
 * block time to 0.
 *
 * Use xStreamBufferSend() to write to a stream buffer from a task.  Use
 * xStreamBufferSendFromISR() to write to a stream buffer from an interrupt
 * service routine (ISR).
 *
 * @param xStreamBuffer The handle of the stream buffer to which a stream is
 * being sent.
 *
 * @param pvTxData A pointer to the buffer that holds the bytes to be copied
 * into the stream buffer.
 *
 * @param xDataLengthBytes   The maximum number of bytes to copy from pvTxData
 * into the stream buffer.
 *
 * @param xTicksToWait The maximum amount of time the task should remain in the
 * Blocked state to wait for enough space to become available in the stream
 * buffer, should the stream buffer contain too little space to hold the
 * another xDataLengthBytes bytes.  The block time is specified in tick periods,
 * so the absolute time it represents is dependent on the tick frequency.  The
 * macro pdMS_TO_TICKS() can be used to convert a time specified in milliseconds
 * into a time specified in ticks.  Setting xTicksToWait to portMAX_DELAY will
 * cause the task to wait indefinitely (without timing out), provided
 * INCLUDE_vTaskSuspend is set to 1 in FreeRTOSConfig.h.  If a task times out
 * before it can write all xDataLengthBytes into the buffer it will still write
 * as many bytes as possible.  A task does not use any CPU time when it is in
 * the blocked state.
 *
 * @return The number of bytes written to the stream buffer.  If a task times
 * out before it can write all xDataLengthBytes into the buffer it will still
 * write as many bytes as possible.
 *
 * Example use:
 * @code{c}
 * void vAFunction( StreamBufferHandle_t xStreamBuffer )
 * {
 * size_t xBytesSent;
 * uint8_t ucArrayToSend[] = { 0, 1, 2, 3 };
 * char *pcStringToSend = "String to send";
 * const TickType_t x100ms = pdMS_TO_TICKS( 100 );
 *
 *  // Send an array to the stream buffer, blocking for a maximum of 100ms to
 *  // wait for enough space to be available in the stream buffer.
 *  xBytesSent = xStreamBufferSend( xStreamBuffer, ( void * ) ucArrayToSend, sizeof( ucArrayToSend ), x100ms );
 *
 *  if( xBytesSent != sizeof( ucArrayToSend ) )
 *  {
 *      // The call to xStreamBufferSend() times out before there was enough
 *      // space in the buffer for the data to be written, but it did
 *      // successfully write xBytesSent bytes.
 *  }
 *
 *  // Send the string to the stream buffer.  Return immediately if there is not
 *  // enough space in the buffer.
 *  xBytesSent = xStreamBufferSend( xStreamBuffer, ( void * ) pcStringToSend, strlen( pcStringToSend ), 0 );
 *
 *  if( xBytesSent != strlen( pcStringToSend ) )
 *  {
 *      // The entire string could not be added to the stream buffer because
 *      // there was not enough free space in the buffer, but xBytesSent bytes
 *      // were sent.  Could try again to send the remaining bytes.
 *  }
 * }
 * @endcode
 * \ingroup StreamBufferManagement
 */
//go:linkname XStreamBufferSend C.xStreamBufferSend
func XStreamBufferSend(xStreamBuffer StreamBufferHandleT, pvTxData c.Pointer, xDataLengthBytes c.SizeT, xTicksToWait TickTypeT) c.SizeT

/**
 *
 * Interrupt safe version of the API function that sends a stream of bytes to
 * the stream buffer.
 *
 * ***NOTE***:  Uniquely among FreeRTOS objects, the stream buffer
 * implementation (so also the message buffer implementation, as message buffers
 * are built on top of stream buffers) assumes there is only one task or
 * interrupt that will write to the buffer (the writer), and only one task or
 * interrupt that will read from the buffer (the reader).  It is safe for the
 * writer and reader to be different tasks or interrupts, but, unlike other
 * FreeRTOS objects, it is not safe to have multiple different writers or
 * multiple different readers.  If there are to be multiple different writers
 * then the application writer must place each call to a writing API function
 * (such as xStreamBufferSend()) inside a critical section and set the send
 * block time to 0.  Likewise, if there are to be multiple different readers
 * then the application writer must place each call to a reading API function
 * (such as xStreamBufferReceive()) inside a critical section and set the receive
 * block time to 0.
 *
 * Use xStreamBufferSend() to write to a stream buffer from a task.  Use
 * xStreamBufferSendFromISR() to write to a stream buffer from an interrupt
 * service routine (ISR).
 *
 * @param xStreamBuffer The handle of the stream buffer to which a stream is
 * being sent.
 *
 * @param pvTxData A pointer to the data that is to be copied into the stream
 * buffer.
 *
 * @param xDataLengthBytes The maximum number of bytes to copy from pvTxData
 * into the stream buffer.
 *
 * @param pxHigherPriorityTaskWoken  It is possible that a stream buffer will
 * have a task blocked on it waiting for data.  Calling
 * xStreamBufferSendFromISR() can make data available, and so cause a task that
 * was waiting for data to leave the Blocked state.  If calling
 * xStreamBufferSendFromISR() causes a task to leave the Blocked state, and the
 * unblocked task has a priority higher than the currently executing task (the
 * task that was interrupted), then, internally, xStreamBufferSendFromISR()
 * will set *pxHigherPriorityTaskWoken to pdTRUE.  If
 * xStreamBufferSendFromISR() sets this value to pdTRUE, then normally a
 * context switch should be performed before the interrupt is exited.  This will
 * ensure that the interrupt returns directly to the highest priority Ready
 * state task.  *pxHigherPriorityTaskWoken should be set to pdFALSE before it
 * is passed into the function.  See the example code below for an example.
 *
 * @return The number of bytes actually written to the stream buffer, which will
 * be less than xDataLengthBytes if the stream buffer didn't have enough free
 * space for all the bytes to be written.
 *
 * Example use:
 * @code{c}
 * // A stream buffer that has already been created.
 * StreamBufferHandle_t xStreamBuffer;
 *
 * void vAnInterruptServiceRoutine( void )
 * {
 * size_t xBytesSent;
 * char *pcStringToSend = "String to send";
 * BaseType_t xHigherPriorityTaskWoken = pdFALSE; // Initialised to pdFALSE.
 *
 *  // Attempt to send the string to the stream buffer.
 *  xBytesSent = xStreamBufferSendFromISR( xStreamBuffer,
 *                                         ( void * ) pcStringToSend,
 *                                         strlen( pcStringToSend ),
 *                                         &xHigherPriorityTaskWoken );
 *
 *  if( xBytesSent != strlen( pcStringToSend ) )
 *  {
 *      // There was not enough free space in the stream buffer for the entire
 *      // string to be written, ut xBytesSent bytes were written.
 *  }
 *
 *  // If xHigherPriorityTaskWoken was set to pdTRUE inside
 *  // xStreamBufferSendFromISR() then a task that has a priority above the
 *  // priority of the currently executing task was unblocked and a context
 *  // switch should be performed to ensure the ISR returns to the unblocked
 *  // task.  In most FreeRTOS ports this is done by simply passing
 *  // xHigherPriorityTaskWoken into portYIELD_FROM_ISR(), which will test the
 *  // variables value, and perform the context switch if necessary.  Check the
 *  // documentation for the port in use for port specific instructions.
 *  portYIELD_FROM_ISR( xHigherPriorityTaskWoken );
 * }
 * @endcode
 * \ingroup StreamBufferManagement
 */
//go:linkname XStreamBufferSendFromISR C.xStreamBufferSendFromISR
func XStreamBufferSendFromISR(xStreamBuffer StreamBufferHandleT, pvTxData c.Pointer, xDataLengthBytes c.SizeT, pxHigherPriorityTaskWoken *BaseTypeT) c.SizeT

/**
 *
 * Receives bytes from a stream buffer.
 *
 * ***NOTE***:  Uniquely among FreeRTOS objects, the stream buffer
 * implementation (so also the message buffer implementation, as message buffers
 * are built on top of stream buffers) assumes there is only one task or
 * interrupt that will write to the buffer (the writer), and only one task or
 * interrupt that will read from the buffer (the reader).  It is safe for the
 * writer and reader to be different tasks or interrupts, but, unlike other
 * FreeRTOS objects, it is not safe to have multiple different writers or
 * multiple different readers.  If there are to be multiple different writers
 * then the application writer must place each call to a writing API function
 * (such as xStreamBufferSend()) inside a critical section and set the send
 * block time to 0.  Likewise, if there are to be multiple different readers
 * then the application writer must place each call to a reading API function
 * (such as xStreamBufferReceive()) inside a critical section and set the receive
 * block time to 0.
 *
 * Use xStreamBufferReceive() to read from a stream buffer from a task.  Use
 * xStreamBufferReceiveFromISR() to read from a stream buffer from an
 * interrupt service routine (ISR).
 *
 * @param xStreamBuffer The handle of the stream buffer from which bytes are to
 * be received.
 *
 * @param pvRxData A pointer to the buffer into which the received bytes will be
 * copied.
 *
 * @param xBufferLengthBytes The length of the buffer pointed to by the
 * pvRxData parameter.  This sets the maximum number of bytes to receive in one
 * call.  xStreamBufferReceive will return as many bytes as possible up to a
 * maximum set by xBufferLengthBytes.
 *
 * @param xTicksToWait The maximum amount of time the task should remain in the
 * Blocked state to wait for data to become available if the stream buffer is
 * empty.  xStreamBufferReceive() will return immediately if xTicksToWait is
 * zero.  The block time is specified in tick periods, so the absolute time it
 * represents is dependent on the tick frequency.  The macro pdMS_TO_TICKS() can
 * be used to convert a time specified in milliseconds into a time specified in
 * ticks.  Setting xTicksToWait to portMAX_DELAY will cause the task to wait
 * indefinitely (without timing out), provided INCLUDE_vTaskSuspend is set to 1
 * in FreeRTOSConfig.h.  A task does not use any CPU time when it is in the
 * Blocked state.
 *
 * @return The number of bytes actually read from the stream buffer, which will
 * be less than xBufferLengthBytes if the call to xStreamBufferReceive() timed
 * out before xBufferLengthBytes were available.
 *
 * Example use:
 * @code{c}
 * void vAFunction( StreamBuffer_t xStreamBuffer )
 * {
 * uint8_t ucRxData[ 20 ];
 * size_t xReceivedBytes;
 * const TickType_t xBlockTime = pdMS_TO_TICKS( 20 );
 *
 *  // Receive up to another sizeof( ucRxData ) bytes from the stream buffer.
 *  // Wait in the Blocked state (so not using any CPU processing time) for a
 *  // maximum of 100ms for the full sizeof( ucRxData ) number of bytes to be
 *  // available.
 *  xReceivedBytes = xStreamBufferReceive( xStreamBuffer,
 *                                         ( void * ) ucRxData,
 *                                         sizeof( ucRxData ),
 *                                         xBlockTime );
 *
 *  if( xReceivedBytes > 0 )
 *  {
 *      // A ucRxData contains another xReceivedBytes bytes of data, which can
 *      // be processed here....
 *  }
 * }
 * @endcode
 * \ingroup StreamBufferManagement
 */
//go:linkname XStreamBufferReceive C.xStreamBufferReceive
func XStreamBufferReceive(xStreamBuffer StreamBufferHandleT, pvRxData c.Pointer, xBufferLengthBytes c.SizeT, xTicksToWait TickTypeT) c.SizeT

/**
 *
 * An interrupt safe version of the API function that receives bytes from a
 * stream buffer.
 *
 * Use xStreamBufferReceive() to read bytes from a stream buffer from a task.
 * Use xStreamBufferReceiveFromISR() to read bytes from a stream buffer from an
 * interrupt service routine (ISR).
 *
 * @param xStreamBuffer The handle of the stream buffer from which a stream
 * is being received.
 *
 * @param pvRxData A pointer to the buffer into which the received bytes are
 * copied.
 *
 * @param xBufferLengthBytes The length of the buffer pointed to by the
 * pvRxData parameter.  This sets the maximum number of bytes to receive in one
 * call.  xStreamBufferReceive will return as many bytes as possible up to a
 * maximum set by xBufferLengthBytes.
 *
 * @param pxHigherPriorityTaskWoken  It is possible that a stream buffer will
 * have a task blocked on it waiting for space to become available.  Calling
 * xStreamBufferReceiveFromISR() can make space available, and so cause a task
 * that is waiting for space to leave the Blocked state.  If calling
 * xStreamBufferReceiveFromISR() causes a task to leave the Blocked state, and
 * the unblocked task has a priority higher than the currently executing task
 * (the task that was interrupted), then, internally,
 * xStreamBufferReceiveFromISR() will set *pxHigherPriorityTaskWoken to pdTRUE.
 * If xStreamBufferReceiveFromISR() sets this value to pdTRUE, then normally a
 * context switch should be performed before the interrupt is exited.  That will
 * ensure the interrupt returns directly to the highest priority Ready state
 * task.  *pxHigherPriorityTaskWoken should be set to pdFALSE before it is
 * passed into the function.  See the code example below for an example.
 *
 * @return The number of bytes read from the stream buffer, if any.
 *
 * Example use:
 * @code{c}
 * // A stream buffer that has already been created.
 * StreamBuffer_t xStreamBuffer;
 *
 * void vAnInterruptServiceRoutine( void )
 * {
 * uint8_t ucRxData[ 20 ];
 * size_t xReceivedBytes;
 * BaseType_t xHigherPriorityTaskWoken = pdFALSE;  // Initialised to pdFALSE.
 *
 *  // Receive the next stream from the stream buffer.
 *  xReceivedBytes = xStreamBufferReceiveFromISR( xStreamBuffer,
 *                                                ( void * ) ucRxData,
 *                                                sizeof( ucRxData ),
 *                                                &xHigherPriorityTaskWoken );
 *
 *  if( xReceivedBytes > 0 )
 *  {
 *      // ucRxData contains xReceivedBytes read from the stream buffer.
 *      // Process the stream here....
 *  }
 *
 *  // If xHigherPriorityTaskWoken was set to pdTRUE inside
 *  // xStreamBufferReceiveFromISR() then a task that has a priority above the
 *  // priority of the currently executing task was unblocked and a context
 *  // switch should be performed to ensure the ISR returns to the unblocked
 *  // task.  In most FreeRTOS ports this is done by simply passing
 *  // xHigherPriorityTaskWoken into portYIELD_FROM_ISR(), which will test the
 *  // variables value, and perform the context switch if necessary.  Check the
 *  // documentation for the port in use for port specific instructions.
 *  portYIELD_FROM_ISR( xHigherPriorityTaskWoken );
 * }
 * @endcode
 * \ingroup StreamBufferManagement
 */
//go:linkname XStreamBufferReceiveFromISR C.xStreamBufferReceiveFromISR
func XStreamBufferReceiveFromISR(xStreamBuffer StreamBufferHandleT, pvRxData c.Pointer, xBufferLengthBytes c.SizeT, pxHigherPriorityTaskWoken *BaseTypeT) c.SizeT

/**
 *
 * Deletes a stream buffer that was previously created using a call to
 * xStreamBufferCreate() or xStreamBufferCreateStatic().  If the stream
 * buffer was created using dynamic memory (that is, by xStreamBufferCreate()),
 * then the allocated memory is freed.
 *
 * A stream buffer handle must not be used after the stream buffer has been
 * deleted.
 *
 * @param xStreamBuffer The handle of the stream buffer to be deleted.
 *
 * \ingroup StreamBufferManagement
 */
//go:linkname VStreamBufferDelete C.vStreamBufferDelete
func VStreamBufferDelete(xStreamBuffer StreamBufferHandleT)

/**
 *
 * Queries a stream buffer to see if it is full.  A stream buffer is full if it
 * does not have any free space, and therefore cannot accept any more data.
 *
 * @param xStreamBuffer The handle of the stream buffer being queried.
 *
 * @return If the stream buffer is full then pdTRUE is returned.  Otherwise
 * pdFALSE is returned.
 *
 * \ingroup StreamBufferManagement
 */
//go:linkname XStreamBufferIsFull C.xStreamBufferIsFull
func XStreamBufferIsFull(xStreamBuffer StreamBufferHandleT) BaseTypeT

/**
 *
 * Queries a stream buffer to see if it is empty.  A stream buffer is empty if
 * it does not contain any data.
 *
 * @param xStreamBuffer The handle of the stream buffer being queried.
 *
 * @return If the stream buffer is empty then pdTRUE is returned.  Otherwise
 * pdFALSE is returned.
 *
 * \ingroup StreamBufferManagement
 */
//go:linkname XStreamBufferIsEmpty C.xStreamBufferIsEmpty
func XStreamBufferIsEmpty(xStreamBuffer StreamBufferHandleT) BaseTypeT

/**
 *
 * Resets a stream buffer to its initial, empty, state.  Any data that was in
 * the stream buffer is discarded.  A stream buffer can only be reset if there
 * are no tasks blocked waiting to either send to or receive from the stream
 * buffer.
 *
 * @param xStreamBuffer The handle of the stream buffer being reset.
 *
 * @return If the stream buffer is reset then pdPASS is returned.  If there was
 * a task blocked waiting to send to or read from the stream buffer then the
 * stream buffer is not reset and pdFAIL is returned.
 *
 * \ingroup StreamBufferManagement
 */
//go:linkname XStreamBufferReset C.xStreamBufferReset
func XStreamBufferReset(xStreamBuffer StreamBufferHandleT) BaseTypeT

/**
 *
 * Queries a stream buffer to see how much free space it contains, which is
 * equal to the amount of data that can be sent to the stream buffer before it
 * is full.
 *
 * @param xStreamBuffer The handle of the stream buffer being queried.
 *
 * @return The number of bytes that can be written to the stream buffer before
 * the stream buffer would be full.
 *
 * \ingroup StreamBufferManagement
 */
//go:linkname XStreamBufferSpacesAvailable C.xStreamBufferSpacesAvailable
func XStreamBufferSpacesAvailable(xStreamBuffer StreamBufferHandleT) c.SizeT

/**
 *
 * Queries a stream buffer to see how much data it contains, which is equal to
 * the number of bytes that can be read from the stream buffer before the stream
 * buffer would be empty.
 *
 * @param xStreamBuffer The handle of the stream buffer being queried.
 *
 * @return The number of bytes that can be read from the stream buffer before
 * the stream buffer would be empty.
 *
 * \ingroup StreamBufferManagement
 */
//go:linkname XStreamBufferBytesAvailable C.xStreamBufferBytesAvailable
func XStreamBufferBytesAvailable(xStreamBuffer StreamBufferHandleT) c.SizeT

/**
 *
 * A stream buffer's trigger level is the number of bytes that must be in the
 * stream buffer before a task that is blocked on the stream buffer to
 * wait for data is moved out of the blocked state.  For example, if a task is
 * blocked on a read of an empty stream buffer that has a trigger level of 1
 * then the task will be unblocked when a single byte is written to the buffer
 * or the task's block time expires.  As another example, if a task is blocked
 * on a read of an empty stream buffer that has a trigger level of 10 then the
 * task will not be unblocked until the stream buffer contains at least 10 bytes
 * or the task's block time expires.  If a reading task's block time expires
 * before the trigger level is reached then the task will still receive however
 * many bytes are actually available.  Setting a trigger level of 0 will result
 * in a trigger level of 1 being used.  It is not valid to specify a trigger
 * level that is greater than the buffer size.
 *
 * A trigger level is set when the stream buffer is created, and can be modified
 * using xStreamBufferSetTriggerLevel().
 *
 * @param xStreamBuffer The handle of the stream buffer being updated.
 *
 * @param xTriggerLevel The new trigger level for the stream buffer.
 *
 * @return If xTriggerLevel was less than or equal to the stream buffer's length
 * then the trigger level will be updated and pdTRUE is returned.  Otherwise
 * pdFALSE is returned.
 *
 * \ingroup StreamBufferManagement
 */
//go:linkname XStreamBufferSetTriggerLevel C.xStreamBufferSetTriggerLevel
func XStreamBufferSetTriggerLevel(xStreamBuffer StreamBufferHandleT, xTriggerLevel c.SizeT) BaseTypeT

/**
 *
 * For advanced users only.
 *
 * The sbSEND_COMPLETED() macro is called from within the FreeRTOS APIs when
 * data is sent to a message buffer or stream buffer.  If there was a task that
 * was blocked on the message or stream buffer waiting for data to arrive then
 * the sbSEND_COMPLETED() macro sends a notification to the task to remove it
 * from the Blocked state.  xStreamBufferSendCompletedFromISR() does the same
 * thing.  It is provided to enable application writers to implement their own
 * version of sbSEND_COMPLETED(), and MUST NOT BE USED AT ANY OTHER TIME.
 *
 * See the example implemented in FreeRTOS/Demo/Minimal/MessageBufferAMP.c for
 * additional information.
 *
 * @param xStreamBuffer The handle of the stream buffer to which data was
 * written.
 *
 * @param pxHigherPriorityTaskWoken *pxHigherPriorityTaskWoken should be
 * initialised to pdFALSE before it is passed into
 * xStreamBufferSendCompletedFromISR().  If calling
 * xStreamBufferSendCompletedFromISR() removes a task from the Blocked state,
 * and the task has a priority above the priority of the currently running task,
 * then *pxHigherPriorityTaskWoken will get set to pdTRUE indicating that a
 * context switch should be performed before exiting the ISR.
 *
 * @return If a task was removed from the Blocked state then pdTRUE is returned.
 * Otherwise pdFALSE is returned.
 *
 * \ingroup StreamBufferManagement
 */
//go:linkname XStreamBufferSendCompletedFromISR C.xStreamBufferSendCompletedFromISR
func XStreamBufferSendCompletedFromISR(xStreamBuffer StreamBufferHandleT, pxHigherPriorityTaskWoken *BaseTypeT) BaseTypeT

/**
 *
 * For advanced users only.
 *
 * The sbRECEIVE_COMPLETED() macro is called from within the FreeRTOS APIs when
 * data is read out of a message buffer or stream buffer.  If there was a task
 * that was blocked on the message or stream buffer waiting for data to arrive
 * then the sbRECEIVE_COMPLETED() macro sends a notification to the task to
 * remove it from the Blocked state.  xStreamBufferReceiveCompletedFromISR()
 * does the same thing.  It is provided to enable application writers to
 * implement their own version of sbRECEIVE_COMPLETED(), and MUST NOT BE USED AT
 * ANY OTHER TIME.
 *
 * See the example implemented in FreeRTOS/Demo/Minimal/MessageBufferAMP.c for
 * additional information.
 *
 * @param xStreamBuffer The handle of the stream buffer from which data was
 * read.
 *
 * @param pxHigherPriorityTaskWoken *pxHigherPriorityTaskWoken should be
 * initialised to pdFALSE before it is passed into
 * xStreamBufferReceiveCompletedFromISR().  If calling
 * xStreamBufferReceiveCompletedFromISR() removes a task from the Blocked state,
 * and the task has a priority above the priority of the currently running task,
 * then *pxHigherPriorityTaskWoken will get set to pdTRUE indicating that a
 * context switch should be performed before exiting the ISR.
 *
 * @return If a task was removed from the Blocked state then pdTRUE is returned.
 * Otherwise pdFALSE is returned.
 *
 * \ingroup StreamBufferManagement
 */
//go:linkname XStreamBufferReceiveCompletedFromISR C.xStreamBufferReceiveCompletedFromISR
func XStreamBufferReceiveCompletedFromISR(xStreamBuffer StreamBufferHandleT, pxHigherPriorityTaskWoken *BaseTypeT) BaseTypeT

/* Functions below here are not part of the public API. */
//go:linkname XStreamBufferGenericCreate C.xStreamBufferGenericCreate
func XStreamBufferGenericCreate(xBufferSizeBytes c.SizeT, xTriggerLevelBytes c.SizeT, xIsMessageBuffer BaseTypeT, pxSendCompletedCallback StreamBufferCallbackFunctionT, pxReceiveCompletedCallback StreamBufferCallbackFunctionT) StreamBufferHandleT

//go:linkname XStreamBufferGenericCreateStatic C.xStreamBufferGenericCreateStatic
func XStreamBufferGenericCreateStatic(xBufferSizeBytes c.SizeT, xTriggerLevelBytes c.SizeT, xIsMessageBuffer BaseTypeT, pucStreamBufferStorageArea *c.Uint8T, pxStaticStreamBuffer *StaticStreamBufferT, pxSendCompletedCallback StreamBufferCallbackFunctionT, pxReceiveCompletedCallback StreamBufferCallbackFunctionT) StreamBufferHandleT

//go:linkname XStreamBufferNextMessageLengthBytes C.xStreamBufferNextMessageLengthBytes
func XStreamBufferNextMessageLengthBytes(xStreamBuffer StreamBufferHandleT) c.SizeT
