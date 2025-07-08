package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Type by which software timers are referenced.  For example, a call to
 * xTimerCreate() returns an TimerHandle_t variable that can then be used to
 * reference the subject timer in calls to other software timer API functions
 * (for example, xTimerStart(), xTimerReset(), etc.).
 */

type TmrTimerControl struct {
	Unused [8]uint8
}
type TimerHandleT *TmrTimerControl

// llgo:type C
type TimerCallbackFunctionT func(TimerHandleT)

// llgo:type C
type PendedFunctionT func(c.Pointer, c.Uint32T)

/**
 *
 * Creates a new software timer instance, and returns a handle by which the
 * created software timer can be referenced.
 *
 * Internally, within the FreeRTOS implementation, software timers use a block
 * of memory, in which the timer data structure is stored.  If a software timer
 * is created using xTimerCreate() then the required memory is automatically
 * dynamically allocated inside the xTimerCreate() function.  (see
 * https://www.FreeRTOS.org/a00111.html).  If a software timer is created using
 * xTimerCreateStatic() then the application writer must provide the memory that
 * will get used by the software timer.  xTimerCreateStatic() therefore allows a
 * software timer to be created without using any dynamic memory allocation.
 *
 * Timers are created in the dormant state.  The xTimerStart(), xTimerReset(),
 * xTimerStartFromISR(), xTimerResetFromISR(), xTimerChangePeriod() and
 * xTimerChangePeriodFromISR() API functions can all be used to transition a
 * timer into the active state.
 *
 * @param pcTimerName A text name that is assigned to the timer.  This is done
 * purely to assist debugging.  The kernel itself only ever references a timer
 * by its handle, and never by its name.
 *
 * @param xTimerPeriodInTicks The timer period.  The time is defined in tick
 * periods so the constant portTICK_PERIOD_MS can be used to convert a time that
 * has been specified in milliseconds.  For example, if the timer must expire
 * after 100 ticks, then xTimerPeriodInTicks should be set to 100.
 * Alternatively, if the timer must expire after 500ms, then xPeriod can be set
 * to ( 500 / portTICK_PERIOD_MS ) provided configTICK_RATE_HZ is less than or
 * equal to 1000.  Time timer period must be greater than 0.
 *
 * @param xAutoReload If xAutoReload is set to pdTRUE then the timer will
 * expire repeatedly with a frequency set by the xTimerPeriodInTicks parameter.
 * If xAutoReload is set to pdFALSE then the timer will be a one-shot timer and
 * enter the dormant state after it expires.
 *
 * @param pvTimerID An identifier that is assigned to the timer being created.
 * Typically this would be used in the timer callback function to identify which
 * timer expired when the same callback function is assigned to more than one
 * timer.
 *
 * @param pxCallbackFunction The function to call when the timer expires.
 * Callback functions must have the prototype defined by TimerCallbackFunction_t,
 * which is "void vCallbackFunction( TimerHandle_t xTimer );".
 *
 * @return If the timer is successfully created then a handle to the newly
 * created timer is returned.  If the timer cannot be created because there is
 * insufficient FreeRTOS heap remaining to allocate the timer
 * structures then NULL is returned.
 *
 * Example usage:
 * @verbatim
 * #define NUM_TIMERS 5
 *
 * // An array to hold handles to the created timers.
 * TimerHandle_t xTimers[ NUM_TIMERS ];
 *
 * // An array to hold a count of the number of times each timer expires.
 * int32_t lExpireCounters[ NUM_TIMERS ] = { 0 };
 *
 * // Define a callback function that will be used by multiple timer instances.
 * // The callback function does nothing but count the number of times the
 * // associated timer expires, and stop the timer once the timer has expired
 * // 10 times.
 * void vTimerCallback( TimerHandle_t pxTimer )
 * {
 * int32_t lArrayIndex;
 * const int32_t xMaxExpiryCountBeforeStopping = 10;
 *
 *     // Optionally do something if the pxTimer parameter is NULL.
 *     configASSERT( pxTimer );
 *
 *     // Which timer expired?
 *     lArrayIndex = ( int32_t ) pvTimerGetTimerID( pxTimer );
 *
 *     // Increment the number of times that pxTimer has expired.
 *     lExpireCounters[ lArrayIndex ] += 1;
 *
 *     // If the timer has expired 10 times then stop it from running.
 *     if( lExpireCounters[ lArrayIndex ] == xMaxExpiryCountBeforeStopping )
 *     {
 *         // Do not use a block time if calling a timer API function from a
 *         // timer callback function, as doing so could cause a deadlock!
 *         xTimerStop( pxTimer, 0 );
 *     }
 * }
 *
 * void main( void )
 * {
 * int32_t x;
 *
 *     // Create then start some timers.  Starting the timers before the scheduler
 *     // has been started means the timers will start running immediately that
 *     // the scheduler starts.
 *     for( x = 0; x < NUM_TIMERS; x++ )
 *     {
 *         xTimers[ x ] = xTimerCreate(    "Timer",             // Just a text name, not used by the kernel.
 *                                         ( 100 * ( x + 1 ) ), // The timer period in ticks.
 *                                         pdTRUE,              // The timers will auto-reload themselves when they expire.
 *                                         ( void * ) x,        // Assign each timer a unique id equal to its array index.
 *                                         vTimerCallback       // Each timer calls the same callback when it expires.
 *                                     );
 *
 *         if( xTimers[ x ] == NULL )
 *         {
 *             // The timer was not created.
 *         }
 *         else
 *         {
 *             // Start the timer.  No block time is specified, and even if one was
 *             // it would be ignored because the scheduler has not yet been
 *             // started.
 *             if( xTimerStart( xTimers[ x ], 0 ) != pdPASS )
 *             {
 *                 // The timer could not be set into the Active state.
 *             }
 *         }
 *     }
 *
 *     // ...
 *     // Create tasks here.
 *     // ...
 *
 *     // Starting the scheduler will start the timers running as they have already
 *     // been set into the active state.
 *     vTaskStartScheduler();
 *
 *     // Should not reach here.
 *     for( ;; );
 * }
 * @endverbatim
 */
//go:linkname XTimerCreate C.xTimerCreate
func XTimerCreate(pcTimerName *c.Char, xTimerPeriodInTicks TickTypeT, xAutoReload BaseTypeT, pvTimerID c.Pointer, pxCallbackFunction TimerCallbackFunctionT) TimerHandleT

/**
 *
 * Creates a new software timer instance, and returns a handle by which the
 * created software timer can be referenced.
 *
 * Internally, within the FreeRTOS implementation, software timers use a block
 * of memory, in which the timer data structure is stored.  If a software timer
 * is created using xTimerCreate() then the required memory is automatically
 * dynamically allocated inside the xTimerCreate() function.  (see
 * https://www.FreeRTOS.org/a00111.html).  If a software timer is created using
 * xTimerCreateStatic() then the application writer must provide the memory that
 * will get used by the software timer.  xTimerCreateStatic() therefore allows a
 * software timer to be created without using any dynamic memory allocation.
 *
 * Timers are created in the dormant state.  The xTimerStart(), xTimerReset(),
 * xTimerStartFromISR(), xTimerResetFromISR(), xTimerChangePeriod() and
 * xTimerChangePeriodFromISR() API functions can all be used to transition a
 * timer into the active state.
 *
 * @param pcTimerName A text name that is assigned to the timer.  This is done
 * purely to assist debugging.  The kernel itself only ever references a timer
 * by its handle, and never by its name.
 *
 * @param xTimerPeriodInTicks The timer period.  The time is defined in tick
 * periods so the constant portTICK_PERIOD_MS can be used to convert a time that
 * has been specified in milliseconds.  For example, if the timer must expire
 * after 100 ticks, then xTimerPeriodInTicks should be set to 100.
 * Alternatively, if the timer must expire after 500ms, then xPeriod can be set
 * to ( 500 / portTICK_PERIOD_MS ) provided configTICK_RATE_HZ is less than or
 * equal to 1000.  The timer period must be greater than 0.
 *
 * @param xAutoReload If xAutoReload is set to pdTRUE then the timer will
 * expire repeatedly with a frequency set by the xTimerPeriodInTicks parameter.
 * If xAutoReload is set to pdFALSE then the timer will be a one-shot timer and
 * enter the dormant state after it expires.
 *
 * @param pvTimerID An identifier that is assigned to the timer being created.
 * Typically this would be used in the timer callback function to identify which
 * timer expired when the same callback function is assigned to more than one
 * timer.
 *
 * @param pxCallbackFunction The function to call when the timer expires.
 * Callback functions must have the prototype defined by TimerCallbackFunction_t,
 * which is "void vCallbackFunction( TimerHandle_t xTimer );".
 *
 * @param pxTimerBuffer Must point to a variable of type StaticTimer_t, which
 * will be then be used to hold the software timer's data structures, removing
 * the need for the memory to be allocated dynamically.
 *
 * @return If the timer is created then a handle to the created timer is
 * returned.  If pxTimerBuffer was NULL then NULL is returned.
 *
 * Example usage:
 * @verbatim
 *
 * // The buffer used to hold the software timer's data structure.
 * static StaticTimer_t xTimerBuffer;
 *
 * // A variable that will be incremented by the software timer's callback
 * // function.
 * UBaseType_t uxVariableToIncrement = 0;
 *
 * // A software timer callback function that increments a variable passed to
 * // it when the software timer was created.  After the 5th increment the
 * // callback function stops the software timer.
 * static void prvTimerCallback( TimerHandle_t xExpiredTimer )
 * {
 * UBaseType_t *puxVariableToIncrement;
 * BaseType_t xReturned;
 *
 *     // Obtain the address of the variable to increment from the timer ID.
 *     puxVariableToIncrement = ( UBaseType_t * ) pvTimerGetTimerID( xExpiredTimer );
 *
 *     // Increment the variable to show the timer callback has executed.
 *     ( *puxVariableToIncrement )++;
 *
 *     // If this callback has executed the required number of times, stop the
 *     // timer.
 *     if( *puxVariableToIncrement == 5 )
 *     {
 *         // This is called from a timer callback so must not block.
 *         xTimerStop( xExpiredTimer, staticDONT_BLOCK );
 *     }
 * }
 *
 *
 * void main( void )
 * {
 *     // Create the software time.  xTimerCreateStatic() has an extra parameter
 *     // than the normal xTimerCreate() API function.  The parameter is a pointer
 *     // to the StaticTimer_t structure that will hold the software timer
 *     // structure.  If the parameter is passed as NULL then the structure will be
 *     // allocated dynamically, just as if xTimerCreate() had been called.
 *     xTimer = xTimerCreateStatic( "T1",             // Text name for the task.  Helps debugging only.  Not used by FreeRTOS.
 *                                  xTimerPeriod,     // The period of the timer in ticks.
 *                                  pdTRUE,           // This is an auto-reload timer.
 *                                  ( void * ) &uxVariableToIncrement,    // A variable incremented by the software timer's callback function
 *                                  prvTimerCallback, // The function to execute when the timer expires.
 *                                  &xTimerBuffer );  // The buffer that will hold the software timer structure.
 *
 *     // The scheduler has not started yet so a block time is not used.
 *     xReturned = xTimerStart( xTimer, 0 );
 *
 *     // ...
 *     // Create tasks here.
 *     // ...
 *
 *     // Starting the scheduler will start the timers running as they have already
 *     // been set into the active state.
 *     vTaskStartScheduler();
 *
 *     // Should not reach here.
 *     for( ;; );
 * }
 * @endverbatim
 */
//go:linkname XTimerCreateStatic C.xTimerCreateStatic
func XTimerCreateStatic(pcTimerName *c.Char, xTimerPeriodInTicks TickTypeT, xAutoReload BaseTypeT, pvTimerID c.Pointer, pxCallbackFunction TimerCallbackFunctionT, pxTimerBuffer *StaticTimerT) TimerHandleT

/**
 *
 * Returns the ID assigned to the timer.
 *
 * IDs are assigned to timers using the pvTimerID parameter of the call to
 * xTimerCreated() that was used to create the timer, and by calling the
 * vTimerSetTimerID() API function.
 *
 * If the same callback function is assigned to multiple timers then the timer
 * ID can be used as time specific (timer local) storage.
 *
 * @param xTimer The timer being queried.
 *
 * @return The ID assigned to the timer being queried.
 *
 * Example usage:
 *
 * See the xTimerCreate() API function example usage scenario.
 */
//go:linkname PvTimerGetTimerID C.pvTimerGetTimerID
func PvTimerGetTimerID(xTimer TimerHandleT) c.Pointer

/**
 *
 * Sets the ID assigned to the timer.
 *
 * IDs are assigned to timers using the pvTimerID parameter of the call to
 * xTimerCreated() that was used to create the timer.
 *
 * If the same callback function is assigned to multiple timers then the timer
 * ID can be used as time specific (timer local) storage.
 *
 * @param xTimer The timer being updated.
 *
 * @param pvNewID The ID to assign to the timer.
 *
 * Example usage:
 *
 * See the xTimerCreate() API function example usage scenario.
 */
//go:linkname VTimerSetTimerID C.vTimerSetTimerID
func VTimerSetTimerID(xTimer TimerHandleT, pvNewID c.Pointer)

/**
 *
 * Queries a timer to see if it is active or dormant.
 *
 * A timer will be dormant if:
 *     1) It has been created but not started, or
 *     2) It is an expired one-shot timer that has not been restarted.
 *
 * Timers are created in the dormant state.  The xTimerStart(), xTimerReset(),
 * xTimerStartFromISR(), xTimerResetFromISR(), xTimerChangePeriod() and
 * xTimerChangePeriodFromISR() API functions can all be used to transition a timer into the
 * active state.
 *
 * @param xTimer The timer being queried.
 *
 * @return pdFALSE will be returned if the timer is dormant.  A value other than
 * pdFALSE will be returned if the timer is active.
 *
 * Example usage:
 * @verbatim
 * // This function assumes xTimer has already been created.
 * void vAFunction( TimerHandle_t xTimer )
 * {
 *     if( xTimerIsTimerActive( xTimer ) != pdFALSE ) // or more simply and equivalently "if( xTimerIsTimerActive( xTimer ) )"
 *     {
 *         // xTimer is active, do something.
 *     }
 *     else
 *     {
 *         // xTimer is not active, do something else.
 *     }
 * }
 * @endverbatim
 */
//go:linkname XTimerIsTimerActive C.xTimerIsTimerActive
func XTimerIsTimerActive(xTimer TimerHandleT) BaseTypeT

/**
 *
 * Simply returns the handle of the timer service/daemon task.  It it not valid
 * to call xTimerGetTimerDaemonTaskHandle() before the scheduler has been started.
 */
//go:linkname XTimerGetTimerDaemonTaskHandle C.xTimerGetTimerDaemonTaskHandle
func XTimerGetTimerDaemonTaskHandle() TaskHandleT

/**
 *
 *
 * Used from application interrupt service routines to defer the execution of a
 * function to the RTOS daemon task (the timer service task, hence this function
 * is implemented in timers.c and is prefixed with 'Timer').
 *
 * Ideally an interrupt service routine (ISR) is kept as short as possible, but
 * sometimes an ISR either has a lot of processing to do, or needs to perform
 * processing that is not deterministic.  In these cases
 * xTimerPendFunctionCallFromISR() can be used to defer processing of a function
 * to the RTOS daemon task.
 *
 * A mechanism is provided that allows the interrupt to return directly to the
 * task that will subsequently execute the pended callback function.  This
 * allows the callback function to execute contiguously in time with the
 * interrupt - just as if the callback had executed in the interrupt itself.
 *
 * @param xFunctionToPend The function to execute from the timer service/
 * daemon task.  The function must conform to the PendedFunction_t
 * prototype.
 *
 * @param pvParameter1 The value of the callback function's first parameter.
 * The parameter has a void * type to allow it to be used to pass any type.
 * For example, unsigned longs can be cast to a void *, or the void * can be
 * used to point to a structure.
 *
 * @param ulParameter2 The value of the callback function's second parameter.
 *
 * @param pxHigherPriorityTaskWoken As mentioned above, calling this function
 * will result in a message being sent to the timer daemon task.  If the
 * priority of the timer daemon task (which is set using
 * configTIMER_TASK_PRIORITY in FreeRTOSConfig.h) is higher than the priority of
 * the currently running task (the task the interrupt interrupted) then
 * *pxHigherPriorityTaskWoken will be set to pdTRUE within
 * xTimerPendFunctionCallFromISR(), indicating that a context switch should be
 * requested before the interrupt exits.  For that reason
 * *pxHigherPriorityTaskWoken must be initialised to pdFALSE.  See the
 * example code below.
 *
 * @return pdPASS is returned if the message was successfully sent to the
 * timer daemon task, otherwise pdFALSE is returned.
 *
 * Example usage:
 * @verbatim
 *
 *  // The callback function that will execute in the context of the daemon task.
 *  // Note callback functions must all use this same prototype.
 *  void vProcessInterface( void *pvParameter1, uint32_t ulParameter2 )
 *  {
 *      BaseType_t xInterfaceToService;
 *
 *      // The interface that requires servicing is passed in the second
 *      // parameter.  The first parameter is not used in this case.
 *      xInterfaceToService = ( BaseType_t ) ulParameter2;
 *
 *      // ...Perform the processing here...
 *  }
 *
 *  // An ISR that receives data packets from multiple interfaces
 *  void vAnISR( void )
 *  {
 *      BaseType_t xInterfaceToService, xHigherPriorityTaskWoken;
 *
 *      // Query the hardware to determine which interface needs processing.
 *      xInterfaceToService = prvCheckInterfaces();
 *
 *      // The actual processing is to be deferred to a task.  Request the
 *      // vProcessInterface() callback function is executed, passing in the
 *      // number of the interface that needs processing.  The interface to
 *      // service is passed in the second parameter.  The first parameter is
 *      // not used in this case.
 *      xHigherPriorityTaskWoken = pdFALSE;
 *      xTimerPendFunctionCallFromISR( vProcessInterface, NULL, ( uint32_t ) xInterfaceToService, &xHigherPriorityTaskWoken );
 *
 *      // If xHigherPriorityTaskWoken is now set to pdTRUE then a context
 *      // switch should be requested.  The macro used is port specific and will
 *      // be either portYIELD_FROM_ISR() or portEND_SWITCHING_ISR() - refer to
 *      // the documentation page for the port being used.
 *      portYIELD_FROM_ISR( xHigherPriorityTaskWoken );
 *
 *  }
 * @endverbatim
 */
//go:linkname XTimerPendFunctionCallFromISR C.xTimerPendFunctionCallFromISR
func XTimerPendFunctionCallFromISR(xFunctionToPend PendedFunctionT, pvParameter1 c.Pointer, ulParameter2 c.Uint32T, pxHigherPriorityTaskWoken *BaseTypeT) BaseTypeT

/**
 *
 *
 * Used to defer the execution of a function to the RTOS daemon task (the timer
 * service task, hence this function is implemented in timers.c and is prefixed
 * with 'Timer').
 *
 * @param xFunctionToPend The function to execute from the timer service/
 * daemon task.  The function must conform to the PendedFunction_t
 * prototype.
 *
 * @param pvParameter1 The value of the callback function's first parameter.
 * The parameter has a void * type to allow it to be used to pass any type.
 * For example, unsigned longs can be cast to a void *, or the void * can be
 * used to point to a structure.
 *
 * @param ulParameter2 The value of the callback function's second parameter.
 *
 * @param xTicksToWait Calling this function will result in a message being
 * sent to the timer daemon task on a queue.  xTicksToWait is the amount of
 * time the calling task should remain in the Blocked state (so not using any
 * processing time) for space to become available on the timer queue if the
 * queue is found to be full.
 *
 * @return pdPASS is returned if the message was successfully sent to the
 * timer daemon task, otherwise pdFALSE is returned.
 *
 */
//go:linkname XTimerPendFunctionCall C.xTimerPendFunctionCall
func XTimerPendFunctionCall(xFunctionToPend PendedFunctionT, pvParameter1 c.Pointer, ulParameter2 c.Uint32T, xTicksToWait TickTypeT) BaseTypeT

/**
 *
 * Returns the name that was assigned to a timer when the timer was created.
 *
 * @param xTimer The handle of the timer being queried.
 *
 * @return The name assigned to the timer specified by the xTimer parameter.
 */
//go:linkname PcTimerGetName C.pcTimerGetName
func PcTimerGetName(xTimer TimerHandleT) *c.Char

/**
 *
 * Updates a timer to be either an auto-reload timer, in which case the timer
 * automatically resets itself each time it expires, or a one-shot timer, in
 * which case the timer will only expire once unless it is manually restarted.
 *
 * @param xTimer The handle of the timer being updated.
 *
 * @param xAutoReload If xAutoReload is set to pdTRUE then the timer will
 * expire repeatedly with a frequency set by the timer's period (see the
 * xTimerPeriodInTicks parameter of the xTimerCreate() API function).  If
 * xAutoReload is set to pdFALSE then the timer will be a one-shot timer and
 * enter the dormant state after it expires.
 */
//go:linkname VTimerSetReloadMode C.vTimerSetReloadMode
func VTimerSetReloadMode(xTimer TimerHandleT, xAutoReload BaseTypeT)

/**
 *
 * Queries a timer to determine if it is an auto-reload timer, in which case the timer
 * automatically resets itself each time it expires, or a one-shot timer, in
 * which case the timer will only expire once unless it is manually restarted.
 *
 * @param xTimer The handle of the timer being queried.
 *
 * @return If the timer is an auto-reload timer then pdTRUE is returned, otherwise
 * pdFALSE is returned.
 */
//go:linkname XTimerGetReloadMode C.xTimerGetReloadMode
func XTimerGetReloadMode(xTimer TimerHandleT) BaseTypeT

/**
 *
 * Queries a timer to determine if it is an auto-reload timer, in which case the timer
 * automatically resets itself each time it expires, or a one-shot timer, in
 * which case the timer will only expire once unless it is manually restarted.
 *
 * @param xTimer The handle of the timer being queried.
 *
 * @return If the timer is an auto-reload timer then pdTRUE is returned, otherwise
 * pdFALSE is returned.
 */
//go:linkname UxTimerGetReloadMode C.uxTimerGetReloadMode
func UxTimerGetReloadMode(xTimer TimerHandleT) UBaseTypeT

/**
 *
 * Returns the period of a timer.
 *
 * @param xTimer The handle of the timer being queried.
 *
 * @return The period of the timer in ticks.
 */
//go:linkname XTimerGetPeriod C.xTimerGetPeriod
func XTimerGetPeriod(xTimer TimerHandleT) TickTypeT

/**
 *
 * Returns the time in ticks at which the timer will expire.  If this is less
 * than the current tick count then the expiry time has overflowed from the
 * current time.
 *
 * @param xTimer The handle of the timer being queried.
 *
 * @return If the timer is running then the time in ticks at which the timer
 * will next expire is returned.  If the timer is not running then the return
 * value is undefined.
 */
//go:linkname XTimerGetExpiryTime C.xTimerGetExpiryTime
func XTimerGetExpiryTime(xTimer TimerHandleT) TickTypeT

/**
 *
 * Retrieve pointer to a statically created timer's data structure
 * buffer. This is the same buffer that is supplied at the time of
 * creation.
 *
 * @param xTimer The timer for which to retrieve the buffer.
 *
 * @param ppxTimerBuffer Used to return a pointer to the timers's data
 * structure buffer.
 *
 * @return pdTRUE if the buffer was retrieved, pdFALSE otherwise.
 */
//go:linkname XTimerGetStaticBuffer C.xTimerGetStaticBuffer
func XTimerGetStaticBuffer(xTimer TimerHandleT, ppxTimerBuffer **StaticTimerT) BaseTypeT

/*
 * Functions beyond this part are not part of the public API and are intended
 * for use by the kernel only.
 */
//go:linkname XTimerCreateTimerTask C.xTimerCreateTimerTask
func XTimerCreateTimerTask() BaseTypeT

//go:linkname XTimerGenericCommand C.xTimerGenericCommand
func XTimerGenericCommand(xTimer TimerHandleT, xCommandID BaseTypeT, xOptionalValue TickTypeT, pxHigherPriorityTaskWoken *BaseTypeT, xTicksToWait TickTypeT) BaseTypeT

/**
 *
 * This function is used to provide a statically allocated block of memory to FreeRTOS to hold the Timer Task TCB.  This function is required when
 * configSUPPORT_STATIC_ALLOCATION is set.  For more information see this URI: https://www.FreeRTOS.org/a00110.html#configSUPPORT_STATIC_ALLOCATION
 *
 * @param ppxTimerTaskTCBBuffer   A handle to a statically allocated TCB buffer
 * @param ppxTimerTaskStackBuffer A handle to a statically allocated Stack buffer for the idle task
 * @param pulTimerTaskStackSize   A pointer to the number of elements that will fit in the allocated stack buffer
 */
//go:linkname VApplicationGetTimerTaskMemory C.vApplicationGetTimerTaskMemory
func VApplicationGetTimerTaskMemory(ppxTimerTaskTCBBuffer **StaticTaskT, ppxTimerTaskStackBuffer **StackTypeT, pulTimerTaskStackSize *c.Uint32T)
