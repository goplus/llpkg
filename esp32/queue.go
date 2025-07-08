package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Type by which queues are referenced.  For example, a call to xQueueCreate()
 * returns an QueueHandle_t variable that can then be used as a parameter to
 * xQueueSend(), xQueueReceive(), etc.
 */

type QueueDefinition struct {
	Unused [8]uint8
}
type QueueHandleT *QueueDefinition
type QueueSetHandleT *QueueDefinition
type QueueSetMemberHandleT *QueueDefinition

/**
 *
 * It is preferred that the macros xQueueSend(), xQueueSendToFront() and
 * xQueueSendToBack() are used in place of calling this function directly.
 *
 * Post an item on a queue.  The item is queued by copy, not by reference.
 * This function must not be called from an interrupt service routine.
 * See xQueueSendFromISR () for an alternative which may be used in an ISR.
 *
 * @param xQueue The handle to the queue on which the item is to be posted.
 *
 * @param pvItemToQueue A pointer to the item that is to be placed on the
 * queue.  The size of the items the queue will hold was defined when the
 * queue was created, so this many bytes will be copied from pvItemToQueue
 * into the queue storage area.
 *
 * @param xTicksToWait The maximum amount of time the task should block
 * waiting for space to become available on the queue, should it already
 * be full.  The call will return immediately if this is set to 0 and the
 * queue is full.  The time is defined in tick periods so the constant
 * portTICK_PERIOD_MS should be used to convert to real time if this is required.
 *
 * @param xCopyPosition Can take the value queueSEND_TO_BACK to place the
 * item at the back of the queue, or queueSEND_TO_FRONT to place the item
 * at the front of the queue (for high priority messages).
 *
 * @return pdTRUE if the item was successfully posted, otherwise errQUEUE_FULL.
 *
 * Example usage:
 * @code{c}
 * struct AMessage
 * {
 *  char ucMessageID;
 *  char ucData[ 20 ];
 * } xMessage;
 *
 * uint32_t ulVar = 10UL;
 *
 * void vATask( void *pvParameters )
 * {
 * QueueHandle_t xQueue1, xQueue2;
 * struct AMessage *pxMessage;
 *
 *  // Create a queue capable of containing 10 uint32_t values.
 *  xQueue1 = xQueueCreate( 10, sizeof( uint32_t ) );
 *
 *  // Create a queue capable of containing 10 pointers to AMessage structures.
 *  // These should be passed by pointer as they contain a lot of data.
 *  xQueue2 = xQueueCreate( 10, sizeof( struct AMessage * ) );
 *
 *  // ...
 *
 *  if( xQueue1 != 0 )
 *  {
 *      // Send an uint32_t.  Wait for 10 ticks for space to become
 *      // available if necessary.
 *      if( xQueueGenericSend( xQueue1, ( void * ) &ulVar, ( TickType_t ) 10, queueSEND_TO_BACK ) != pdPASS )
 *      {
 *          // Failed to post the message, even after 10 ticks.
 *      }
 *  }
 *
 *  if( xQueue2 != 0 )
 *  {
 *      // Send a pointer to a struct AMessage object.  Don't block if the
 *      // queue is already full.
 *      pxMessage = & xMessage;
 *      xQueueGenericSend( xQueue2, ( void * ) &pxMessage, ( TickType_t ) 0, queueSEND_TO_BACK );
 *  }
 *
 *  // ... Rest of task code.
 * }
 * @endcode
 * \ingroup QueueManagement
 */
//go:linkname XQueueGenericSend C.xQueueGenericSend
func XQueueGenericSend(xQueue QueueHandleT, pvItemToQueue c.Pointer, xTicksToWait TickTypeT, xCopyPosition BaseTypeT) BaseTypeT

/**
 *
 * Receive an item from a queue without removing the item from the queue.
 * The item is received by copy so a buffer of adequate size must be
 * provided.  The number of bytes copied into the buffer was defined when
 * the queue was created.
 *
 * Successfully received items remain on the queue so will be returned again
 * by the next call, or a call to xQueueReceive().
 *
 * This macro must not be used in an interrupt service routine.  See
 * xQueuePeekFromISR() for an alternative that can be called from an interrupt
 * service routine.
 *
 * @param xQueue The handle to the queue from which the item is to be
 * received.
 *
 * @param pvBuffer Pointer to the buffer into which the received item will
 * be copied.
 *
 * @param xTicksToWait The maximum amount of time the task should block
 * waiting for an item to receive should the queue be empty at the time
 * of the call. The time is defined in tick periods so the constant
 * portTICK_PERIOD_MS should be used to convert to real time if this is required.
 * xQueuePeek() will return immediately if xTicksToWait is 0 and the queue
 * is empty.
 *
 * @return pdTRUE if an item was successfully received from the queue,
 * otherwise pdFALSE.
 *
 * Example usage:
 * @code{c}
 * struct AMessage
 * {
 *  char ucMessageID;
 *  char ucData[ 20 ];
 * } xMessage;
 *
 * QueueHandle_t xQueue;
 *
 * // Task to create a queue and post a value.
 * void vATask( void *pvParameters )
 * {
 * struct AMessage *pxMessage;
 *
 *  // Create a queue capable of containing 10 pointers to AMessage structures.
 *  // These should be passed by pointer as they contain a lot of data.
 *  xQueue = xQueueCreate( 10, sizeof( struct AMessage * ) );
 *  if( xQueue == 0 )
 *  {
 *      // Failed to create the queue.
 *  }
 *
 *  // ...
 *
 *  // Send a pointer to a struct AMessage object.  Don't block if the
 *  // queue is already full.
 *  pxMessage = & xMessage;
 *  xQueueSend( xQueue, ( void * ) &pxMessage, ( TickType_t ) 0 );
 *
 *  // ... Rest of task code.
 * }
 *
 * // Task to peek the data from the queue.
 * void vADifferentTask( void *pvParameters )
 * {
 * struct AMessage *pxRxedMessage;
 *
 *  if( xQueue != 0 )
 *  {
 *      // Peek a message on the created queue.  Block for 10 ticks if a
 *      // message is not immediately available.
 *      if( xQueuePeek( xQueue, &( pxRxedMessage ), ( TickType_t ) 10 ) )
 *      {
 *          // pcRxedMessage now points to the struct AMessage variable posted
 *          // by vATask, but the item still remains on the queue.
 *      }
 *  }
 *
 *  // ... Rest of task code.
 * }
 * @endcode
 * \ingroup QueueManagement
 */
//go:linkname XQueuePeek C.xQueuePeek
func XQueuePeek(xQueue QueueHandleT, pvBuffer c.Pointer, xTicksToWait TickTypeT) BaseTypeT

/**
 *
 * A version of xQueuePeek() that can be called from an interrupt service
 * routine (ISR).
 *
 * Receive an item from a queue without removing the item from the queue.
 * The item is received by copy so a buffer of adequate size must be
 * provided.  The number of bytes copied into the buffer was defined when
 * the queue was created.
 *
 * Successfully received items remain on the queue so will be returned again
 * by the next call, or a call to xQueueReceive().
 *
 * @param xQueue The handle to the queue from which the item is to be
 * received.
 *
 * @param pvBuffer Pointer to the buffer into which the received item will
 * be copied.
 *
 * @return pdTRUE if an item was successfully received from the queue,
 * otherwise pdFALSE.
 *
 * \ingroup QueueManagement
 */
//go:linkname XQueuePeekFromISR C.xQueuePeekFromISR
func XQueuePeekFromISR(xQueue QueueHandleT, pvBuffer c.Pointer) BaseTypeT

/**
 *
 * Receive an item from a queue.  The item is received by copy so a buffer of
 * adequate size must be provided.  The number of bytes copied into the buffer
 * was defined when the queue was created.
 *
 * Successfully received items are removed from the queue.
 *
 * This function must not be used in an interrupt service routine.  See
 * xQueueReceiveFromISR for an alternative that can.
 *
 * @param xQueue The handle to the queue from which the item is to be
 * received.
 *
 * @param pvBuffer Pointer to the buffer into which the received item will
 * be copied.
 *
 * @param xTicksToWait The maximum amount of time the task should block
 * waiting for an item to receive should the queue be empty at the time
 * of the call. xQueueReceive() will return immediately if xTicksToWait
 * is zero and the queue is empty.  The time is defined in tick periods so the
 * constant portTICK_PERIOD_MS should be used to convert to real time if this is
 * required.
 *
 * @return pdTRUE if an item was successfully received from the queue,
 * otherwise pdFALSE.
 *
 * Example usage:
 * @code{c}
 * struct AMessage
 * {
 *  char ucMessageID;
 *  char ucData[ 20 ];
 * } xMessage;
 *
 * QueueHandle_t xQueue;
 *
 * // Task to create a queue and post a value.
 * void vATask( void *pvParameters )
 * {
 * struct AMessage *pxMessage;
 *
 *  // Create a queue capable of containing 10 pointers to AMessage structures.
 *  // These should be passed by pointer as they contain a lot of data.
 *  xQueue = xQueueCreate( 10, sizeof( struct AMessage * ) );
 *  if( xQueue == 0 )
 *  {
 *      // Failed to create the queue.
 *  }
 *
 *  // ...
 *
 *  // Send a pointer to a struct AMessage object.  Don't block if the
 *  // queue is already full.
 *  pxMessage = & xMessage;
 *  xQueueSend( xQueue, ( void * ) &pxMessage, ( TickType_t ) 0 );
 *
 *  // ... Rest of task code.
 * }
 *
 * // Task to receive from the queue.
 * void vADifferentTask( void *pvParameters )
 * {
 * struct AMessage *pxRxedMessage;
 *
 *  if( xQueue != 0 )
 *  {
 *      // Receive a message on the created queue.  Block for 10 ticks if a
 *      // message is not immediately available.
 *      if( xQueueReceive( xQueue, &( pxRxedMessage ), ( TickType_t ) 10 ) )
 *      {
 *          // pcRxedMessage now points to the struct AMessage variable posted
 *          // by vATask.
 *      }
 *  }
 *
 *  // ... Rest of task code.
 * }
 * @endcode
 * \ingroup QueueManagement
 */
//go:linkname XQueueReceive C.xQueueReceive
func XQueueReceive(xQueue QueueHandleT, pvBuffer c.Pointer, xTicksToWait TickTypeT) BaseTypeT

/**
 *
 * Return the number of messages stored in a queue.
 *
 * @param xQueue A handle to the queue being queried.
 *
 * @return The number of messages available in the queue.
 *
 * \ingroup QueueManagement
 */
//go:linkname UxQueueMessagesWaiting C.uxQueueMessagesWaiting
func UxQueueMessagesWaiting(xQueue QueueHandleT) UBaseTypeT

/**
 *
 * Return the number of free spaces available in a queue.  This is equal to the
 * number of items that can be sent to the queue before the queue becomes full
 * if no items are removed.
 *
 * @param xQueue A handle to the queue being queried.
 *
 * @return The number of spaces available in the queue.
 *
 * \ingroup QueueManagement
 */
//go:linkname UxQueueSpacesAvailable C.uxQueueSpacesAvailable
func UxQueueSpacesAvailable(xQueue QueueHandleT) UBaseTypeT

/**
 *
 * Delete a queue - freeing all the memory allocated for storing of items
 * placed on the queue.
 *
 * @param xQueue A handle to the queue to be deleted.
 *
 * \ingroup QueueManagement
 */
//go:linkname VQueueDelete C.vQueueDelete
func VQueueDelete(xQueue QueueHandleT)

/**
 *
 * It is preferred that the macros xQueueSendFromISR(),
 * xQueueSendToFrontFromISR() and xQueueSendToBackFromISR() be used in place
 * of calling this function directly.  xQueueGiveFromISR() is an
 * equivalent for use by semaphores that don't actually copy any data.
 *
 * Post an item on a queue.  It is safe to use this function from within an
 * interrupt service routine.
 *
 * Items are queued by copy not reference so it is preferable to only
 * queue small items, especially when called from an ISR.  In most cases
 * it would be preferable to store a pointer to the item being queued.
 *
 * @param xQueue The handle to the queue on which the item is to be posted.
 *
 * @param pvItemToQueue A pointer to the item that is to be placed on the
 * queue.  The size of the items the queue will hold was defined when the
 * queue was created, so this many bytes will be copied from pvItemToQueue
 * into the queue storage area.
 *
 * @param pxHigherPriorityTaskWoken xQueueGenericSendFromISR() will set
 * *pxHigherPriorityTaskWoken to pdTRUE if sending to the queue caused a task
 * to unblock, and the unblocked task has a priority higher than the currently
 * running task.  If xQueueGenericSendFromISR() sets this value to pdTRUE then
 * a context switch should be requested before the interrupt is exited.
 *
 * @param xCopyPosition Can take the value queueSEND_TO_BACK to place the
 * item at the back of the queue, or queueSEND_TO_FRONT to place the item
 * at the front of the queue (for high priority messages).
 *
 * @return pdTRUE if the data was successfully sent to the queue, otherwise
 * errQUEUE_FULL.
 *
 * Example usage for buffered IO (where the ISR can obtain more than one value
 * per call):
 * @code{c}
 * void vBufferISR( void )
 * {
 * char cIn;
 * BaseType_t xHigherPriorityTaskWokenByPost;
 *
 *  // We have not woken a task at the start of the ISR.
 *  xHigherPriorityTaskWokenByPost = pdFALSE;
 *
 *  // Loop until the buffer is empty.
 *  do
 *  {
 *      // Obtain a byte from the buffer.
 *      cIn = portINPUT_BYTE( RX_REGISTER_ADDRESS );
 *
 *      // Post each byte.
 *      xQueueGenericSendFromISR( xRxQueue, &cIn, &xHigherPriorityTaskWokenByPost, queueSEND_TO_BACK );
 *
 *  } while( portINPUT_BYTE( BUFFER_COUNT ) );
 *
 *  // Now the buffer is empty we can switch context if necessary.  Note that the
 *  // name of the yield function required is port specific.
 *  if( xHigherPriorityTaskWokenByPost )
 *  {
 *      portYIELD_FROM_ISR();
 *  }
 * }
 * @endcode
 *
 * \ingroup QueueManagement
 */
//go:linkname XQueueGenericSendFromISR C.xQueueGenericSendFromISR
func XQueueGenericSendFromISR(xQueue QueueHandleT, pvItemToQueue c.Pointer, pxHigherPriorityTaskWoken *BaseTypeT, xCopyPosition BaseTypeT) BaseTypeT

//go:linkname XQueueGiveFromISR C.xQueueGiveFromISR
func XQueueGiveFromISR(xQueue QueueHandleT, pxHigherPriorityTaskWoken *BaseTypeT) BaseTypeT

/**
 *
 * Receive an item from a queue.  It is safe to use this function from within an
 * interrupt service routine.
 *
 * @param xQueue The handle to the queue from which the item is to be
 * received.
 *
 * @param pvBuffer Pointer to the buffer into which the received item will
 * be copied.
 *
 * @param pxHigherPriorityTaskWoken A task may be blocked waiting for space to
 * become available on the queue.  If xQueueReceiveFromISR causes such a task to
 * unblock *pxTaskWoken will get set to pdTRUE, otherwise *pxTaskWoken will
 * remain unchanged.
 *
 * @return pdTRUE if an item was successfully received from the queue,
 * otherwise pdFALSE.
 *
 * Example usage:
 * @code{c}
 *
 * QueueHandle_t xQueue;
 *
 * // Function to create a queue and post some values.
 * void vAFunction( void *pvParameters )
 * {
 * char cValueToPost;
 * const TickType_t xTicksToWait = ( TickType_t )0xff;
 *
 *  // Create a queue capable of containing 10 characters.
 *  xQueue = xQueueCreate( 10, sizeof( char ) );
 *  if( xQueue == 0 )
 *  {
 *      // Failed to create the queue.
 *  }
 *
 *  // ...
 *
 *  // Post some characters that will be used within an ISR.  If the queue
 *  // is full then this task will block for xTicksToWait ticks.
 *  cValueToPost = 'a';
 *  xQueueSend( xQueue, ( void * ) &cValueToPost, xTicksToWait );
 *  cValueToPost = 'b';
 *  xQueueSend( xQueue, ( void * ) &cValueToPost, xTicksToWait );
 *
 *  // ... keep posting characters ... this task may block when the queue
 *  // becomes full.
 *
 *  cValueToPost = 'c';
 *  xQueueSend( xQueue, ( void * ) &cValueToPost, xTicksToWait );
 * }
 *
 * // ISR that outputs all the characters received on the queue.
 * void vISR_Routine( void )
 * {
 * BaseType_t xTaskWokenByReceive = pdFALSE;
 * char cRxedChar;
 *
 *  while( xQueueReceiveFromISR( xQueue, ( void * ) &cRxedChar, &xTaskWokenByReceive) )
 *  {
 *      // A character was received.  Output the character now.
 *      vOutputCharacter( cRxedChar );
 *
 *      // If removing the character from the queue woke the task that was
 *      // posting onto the queue xTaskWokenByReceive will have been set to
 *      // pdTRUE.  No matter how many times this loop iterates only one
 *      // task will be woken.
 *  }
 *
 *  if( xTaskWokenByReceive != ( char ) pdFALSE;
 *  {
 *      taskYIELD ();
 *  }
 * }
 * @endcode
 * \ingroup QueueManagement
 */
//go:linkname XQueueReceiveFromISR C.xQueueReceiveFromISR
func XQueueReceiveFromISR(xQueue QueueHandleT, pvBuffer c.Pointer, pxHigherPriorityTaskWoken *BaseTypeT) BaseTypeT

/**
 * Queries a queue to determine if the queue is empty. This function should only be used in an ISR.
 *
 * @param xQueue The handle of the queue being queried
 * @return pdFALSE if the queue is not empty, or pdTRUE if the queue is empty.
 */
//go:linkname XQueueIsQueueEmptyFromISR C.xQueueIsQueueEmptyFromISR
func XQueueIsQueueEmptyFromISR(xQueue QueueHandleT) BaseTypeT

/**
 * Queries a queue to determine if the queue is full. This function should only be used in an ISR.
 *
 * @param xQueue The handle of the queue being queried
 * @return pdFALSE if the queue is not full, or pdTRUE if the queue is full.
 */
//go:linkname XQueueIsQueueFullFromISR C.xQueueIsQueueFullFromISR
func XQueueIsQueueFullFromISR(xQueue QueueHandleT) BaseTypeT

/**
 * A version of uxQueueMessagesWaiting() that can be called from an ISR. Return the number of messages stored in a queue.
 *
 * @param xQueue A handle to the queue being queried.
 * @return The number of messages available in the queue.
 */
//go:linkname UxQueueMessagesWaitingFromISR C.uxQueueMessagesWaitingFromISR
func UxQueueMessagesWaitingFromISR(xQueue QueueHandleT) UBaseTypeT

/*
 * The functions defined above are for passing data to and from tasks.  The
 * functions below are the equivalents for passing data to and from
 * co-routines.
 *
 * These functions are called from the co-routine macro implementation and
 * should not be called directly from application code.  Instead use the macro
 * wrappers defined within croutine.h.
 */
//go:linkname XQueueCRSendFromISR C.xQueueCRSendFromISR
func XQueueCRSendFromISR(xQueue QueueHandleT, pvItemToQueue c.Pointer, xCoRoutinePreviouslyWoken BaseTypeT) BaseTypeT

//go:linkname XQueueCRReceiveFromISR C.xQueueCRReceiveFromISR
func XQueueCRReceiveFromISR(xQueue QueueHandleT, pvBuffer c.Pointer, pxTaskWoken *BaseTypeT) BaseTypeT

//go:linkname XQueueCRSend C.xQueueCRSend
func XQueueCRSend(xQueue QueueHandleT, pvItemToQueue c.Pointer, xTicksToWait TickTypeT) BaseTypeT

//go:linkname XQueueCRReceive C.xQueueCRReceive
func XQueueCRReceive(xQueue QueueHandleT, pvBuffer c.Pointer, xTicksToWait TickTypeT) BaseTypeT

/*
 * For internal use only.  Use xSemaphoreCreateMutex(),
 * xSemaphoreCreateCounting() or xSemaphoreGetMutexHolder() instead of calling
 * these functions directly.
 */
//go:linkname XQueueCreateMutex C.xQueueCreateMutex
func XQueueCreateMutex(ucQueueType c.Uint8T) QueueHandleT

//go:linkname XQueueCreateMutexStatic C.xQueueCreateMutexStatic
func XQueueCreateMutexStatic(ucQueueType c.Uint8T, pxStaticQueue *StaticQueueT) QueueHandleT

// llgo:link UBaseTypeT.XQueueCreateCountingSemaphore C.xQueueCreateCountingSemaphore
func (recv_ UBaseTypeT) XQueueCreateCountingSemaphore(uxInitialCount UBaseTypeT) QueueHandleT {
	return nil
}

// llgo:link UBaseTypeT.XQueueCreateCountingSemaphoreStatic C.xQueueCreateCountingSemaphoreStatic
func (recv_ UBaseTypeT) XQueueCreateCountingSemaphoreStatic(uxInitialCount UBaseTypeT, pxStaticQueue *StaticQueueT) QueueHandleT {
	return nil
}

//go:linkname XQueueSemaphoreTake C.xQueueSemaphoreTake
func XQueueSemaphoreTake(xQueue QueueHandleT, xTicksToWait TickTypeT) BaseTypeT

//go:linkname XQueueGetMutexHolder C.xQueueGetMutexHolder
func XQueueGetMutexHolder(xSemaphore QueueHandleT) TaskHandleT

//go:linkname XQueueGetMutexHolderFromISR C.xQueueGetMutexHolderFromISR
func XQueueGetMutexHolderFromISR(xSemaphore QueueHandleT) TaskHandleT

/*
 * For internal use only.  Use xSemaphoreTakeMutexRecursive() or
 * xSemaphoreGiveMutexRecursive() instead of calling these functions directly.
 */
//go:linkname XQueueTakeMutexRecursive C.xQueueTakeMutexRecursive
func XQueueTakeMutexRecursive(xMutex QueueHandleT, xTicksToWait TickTypeT) BaseTypeT

//go:linkname XQueueGiveMutexRecursive C.xQueueGiveMutexRecursive
func XQueueGiveMutexRecursive(xMutex QueueHandleT) BaseTypeT

/*
 * Generic version of the function used to create a queue using dynamic memory
 * allocation.  This is called by other functions and macros that create other
 * RTOS objects that use the queue structure as their base.
 */
// llgo:link UBaseTypeT.XQueueGenericCreate C.xQueueGenericCreate
func (recv_ UBaseTypeT) XQueueGenericCreate(uxItemSize UBaseTypeT, ucQueueType c.Uint8T) QueueHandleT {
	return nil
}

/*
 * Generic version of the function used to create a queue using dynamic memory
 * allocation.  This is called by other functions and macros that create other
 * RTOS objects that use the queue structure as their base.
 */
// llgo:link UBaseTypeT.XQueueGenericCreateStatic C.xQueueGenericCreateStatic
func (recv_ UBaseTypeT) XQueueGenericCreateStatic(uxItemSize UBaseTypeT, pucQueueStorage *c.Uint8T, pxStaticQueue *StaticQueueT, ucQueueType c.Uint8T) QueueHandleT {
	return nil
}

/*
 * Generic version of the function used to retrieve the buffers of statically
 * created queues. This is called by other functions and macros that retrieve
 * the buffers of other statically created RTOS objects that use the queue
 * structure as their base.
 */
//go:linkname XQueueGenericGetStaticBuffers C.xQueueGenericGetStaticBuffers
func XQueueGenericGetStaticBuffers(xQueue QueueHandleT, ppucQueueStorage **c.Uint8T, ppxStaticQueue **StaticQueueT) BaseTypeT

/**
 * Queue sets provide a mechanism to allow a task to block (pend) on a read
 * operation from multiple queues or semaphores simultaneously.
 *
 * See FreeRTOS/Source/Demo/Common/Minimal/QueueSet.c for an example using this
 * function.
 *
 * A queue set must be explicitly created using a call to xQueueCreateSet()
 * before it can be used.  Once created, standard FreeRTOS queues and semaphores
 * can be added to the set using calls to xQueueAddToSet().
 * xQueueSelectFromSet() is then used to determine which, if any, of the queues
 * or semaphores contained in the set is in a state where a queue read or
 * semaphore take operation would be successful.
 *
 * Note 1:  See the documentation on https://www.FreeRTOS.org/RTOS-queue-sets.html
 * for reasons why queue sets are very rarely needed in practice as there are
 * simpler methods of blocking on multiple objects.
 *
 * Note 2:  Blocking on a queue set that contains a mutex will not cause the
 * mutex holder to inherit the priority of the blocked task.
 *
 * Note 3:  An additional 4 bytes of RAM is required for each space in a every
 * queue added to a queue set.  Therefore counting semaphores that have a high
 * maximum count value should not be added to a queue set.
 *
 * Note 4:  A receive (in the case of a queue) or take (in the case of a
 * semaphore) operation must not be performed on a member of a queue set unless
 * a call to xQueueSelectFromSet() has first returned a handle to that set member.
 *
 * @param uxEventQueueLength Queue sets store events that occur on
 * the queues and semaphores contained in the set.  uxEventQueueLength specifies
 * the maximum number of events that can be queued at once.  To be absolutely
 * certain that events are not lost uxEventQueueLength should be set to the
 * total sum of the length of the queues added to the set, where binary
 * semaphores and mutexes have a length of 1, and counting semaphores have a
 * length set by their maximum count value.  Examples:
 *  + If a queue set is to hold a queue of length 5, another queue of length 12,
 *    and a binary semaphore, then uxEventQueueLength should be set to
 *    (5 + 12 + 1), or 18.
 *  + If a queue set is to hold three binary semaphores then uxEventQueueLength
 *    should be set to (1 + 1 + 1 ), or 3.
 *  + If a queue set is to hold a counting semaphore that has a maximum count of
 *    5, and a counting semaphore that has a maximum count of 3, then
 *    uxEventQueueLength should be set to (5 + 3), or 8.
 *
 * @return If the queue set is created successfully then a handle to the created
 * queue set is returned.  Otherwise NULL is returned.
 */
// llgo:link UBaseTypeT.XQueueCreateSet C.xQueueCreateSet
func (recv_ UBaseTypeT) XQueueCreateSet() QueueSetHandleT {
	return nil
}

/**
 * Adds a queue or semaphore to a queue set that was previously created by a
 * call to xQueueCreateSet().
 *
 * See FreeRTOS/Source/Demo/Common/Minimal/QueueSet.c for an example using this
 * function.
 *
 * Note 1:  A receive (in the case of a queue) or take (in the case of a
 * semaphore) operation must not be performed on a member of a queue set unless
 * a call to xQueueSelectFromSet() has first returned a handle to that set member.
 *
 * @param xQueueOrSemaphore The handle of the queue or semaphore being added to
 * the queue set (cast to an QueueSetMemberHandle_t type).
 *
 * @param xQueueSet The handle of the queue set to which the queue or semaphore
 * is being added.
 *
 * @return If the queue or semaphore was successfully added to the queue set
 * then pdPASS is returned.  If the queue could not be successfully added to the
 * queue set because it is already a member of a different queue set then pdFAIL
 * is returned.
 */
//go:linkname XQueueAddToSet C.xQueueAddToSet
func XQueueAddToSet(xQueueOrSemaphore QueueSetMemberHandleT, xQueueSet QueueSetHandleT) BaseTypeT

/**
 * Removes a queue or semaphore from a queue set.  A queue or semaphore can only
 * be removed from a set if the queue or semaphore is empty.
 *
 * See FreeRTOS/Source/Demo/Common/Minimal/QueueSet.c for an example using this
 * function.
 *
 * @param xQueueOrSemaphore The handle of the queue or semaphore being removed
 * from the queue set (cast to an QueueSetMemberHandle_t type).
 *
 * @param xQueueSet The handle of the queue set in which the queue or semaphore
 * is included.
 *
 * @return If the queue or semaphore was successfully removed from the queue set
 * then pdPASS is returned.  If the queue was not in the queue set, or the
 * queue (or semaphore) was not empty, then pdFAIL is returned.
 */
//go:linkname XQueueRemoveFromSet C.xQueueRemoveFromSet
func XQueueRemoveFromSet(xQueueOrSemaphore QueueSetMemberHandleT, xQueueSet QueueSetHandleT) BaseTypeT

/**
 * xQueueSelectFromSet() selects from the members of a queue set a queue or
 * semaphore that either contains data (in the case of a queue) or is available
 * to take (in the case of a semaphore).  xQueueSelectFromSet() effectively
 * allows a task to block (pend) on a read operation on all the queues and
 * semaphores in a queue set simultaneously.
 *
 * See FreeRTOS/Source/Demo/Common/Minimal/QueueSet.c for an example using this
 * function.
 *
 * Note 1:  See the documentation on https://www.FreeRTOS.org/RTOS-queue-sets.html
 * for reasons why queue sets are very rarely needed in practice as there are
 * simpler methods of blocking on multiple objects.
 *
 * Note 2:  Blocking on a queue set that contains a mutex will not cause the
 * mutex holder to inherit the priority of the blocked task.
 *
 * Note 3:  A receive (in the case of a queue) or take (in the case of a
 * semaphore) operation must not be performed on a member of a queue set unless
 * a call to xQueueSelectFromSet() has first returned a handle to that set member.
 *
 * @param xQueueSet The queue set on which the task will (potentially) block.
 *
 * @param xTicksToWait The maximum time, in ticks, that the calling task will
 * remain in the Blocked state (with other tasks executing) to wait for a member
 * of the queue set to be ready for a successful queue read or semaphore take
 * operation.
 *
 * @return xQueueSelectFromSet() will return the handle of a queue (cast to
 * a QueueSetMemberHandle_t type) contained in the queue set that contains data,
 * or the handle of a semaphore (cast to a QueueSetMemberHandle_t type) contained
 * in the queue set that is available, or NULL if no such queue or semaphore
 * exists before before the specified block time expires.
 */
//go:linkname XQueueSelectFromSet C.xQueueSelectFromSet
func XQueueSelectFromSet(xQueueSet QueueSetHandleT, xTicksToWait TickTypeT) QueueSetMemberHandleT

/**
 * A version of xQueueSelectFromSet() that can be used from an ISR.
 */
//go:linkname XQueueSelectFromSetFromISR C.xQueueSelectFromSetFromISR
func XQueueSelectFromSetFromISR(xQueueSet QueueSetHandleT) QueueSetMemberHandleT

/* Not public API functions. */
//go:linkname VQueueWaitForMessageRestricted C.vQueueWaitForMessageRestricted
func VQueueWaitForMessageRestricted(xQueue QueueHandleT, xTicksToWait TickTypeT, xWaitIndefinitely BaseTypeT)

//go:linkname XQueueGenericReset C.xQueueGenericReset
func XQueueGenericReset(xQueue QueueHandleT, xNewQueue BaseTypeT) BaseTypeT

//go:linkname VQueueSetQueueNumber C.vQueueSetQueueNumber
func VQueueSetQueueNumber(xQueue QueueHandleT, uxQueueNumber UBaseTypeT)

//go:linkname UxQueueGetQueueNumber C.uxQueueGetQueueNumber
func UxQueueGetQueueNumber(xQueue QueueHandleT) UBaseTypeT

//go:linkname UcQueueGetQueueType C.ucQueueGetQueueType
func UcQueueGetQueueType(xQueue QueueHandleT) c.Uint8T
