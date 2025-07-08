package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 *
 * Type by which event groups are referenced.  For example, a call to
 * xEventGroupCreate() returns an EventGroupHandle_t variable that can then
 * be used as a parameter to other event group functions.
 *
 * \ingroup EventGroup
 */

type EventGroupDefT struct {
	Unused [8]uint8
}
type EventGroupHandleT *EventGroupDefT
type EventBitsT TickTypeT

/**
 *
 * Create a new event group.
 *
 * Internally, within the FreeRTOS implementation, event groups use a [small]
 * block of memory, in which the event group's structure is stored.  If an event
 * groups is created using xEventGroupCreate() then the required memory is
 * automatically dynamically allocated inside the xEventGroupCreate() function.
 * (see https://www.FreeRTOS.org/a00111.html).  If an event group is created
 * using xEventGroupCreateStatic() then the application writer must instead
 * provide the memory that will get used by the event group.
 * xEventGroupCreateStatic() therefore allows an event group to be created
 * without using any dynamic memory allocation.
 *
 * Although event groups are not related to ticks, for internal implementation
 * reasons the number of bits available for use in an event group is dependent
 * on the configUSE_16_BIT_TICKS setting in FreeRTOSConfig.h.  If
 * configUSE_16_BIT_TICKS is 1 then each event group contains 8 usable bits (bit
 * 0 to bit 7).  If configUSE_16_BIT_TICKS is set to 0 then each event group has
 * 24 usable bits (bit 0 to bit 23).  The EventBits_t type is used to store
 * event bits within an event group.
 *
 * @return If the event group was created then a handle to the event group is
 * returned.  If there was insufficient FreeRTOS heap available to create the
 * event group then NULL is returned.  See https://www.FreeRTOS.org/a00111.html
 *
 * Example usage:
 * @code{c}
 *  // Declare a variable to hold the created event group.
 *  EventGroupHandle_t xCreatedEventGroup;
 *
 *  // Attempt to create the event group.
 *  xCreatedEventGroup = xEventGroupCreate();
 *
 *  // Was the event group created successfully?
 *  if( xCreatedEventGroup == NULL )
 *  {
 *      // The event group was not created because there was insufficient
 *      // FreeRTOS heap available.
 *  }
 *  else
 *  {
 *      // The event group was created.
 *  }
 * @endcode
 * \ingroup EventGroup
 */
//go:linkname XEventGroupCreate C.xEventGroupCreate
func XEventGroupCreate() EventGroupHandleT

/**
 *
 * Create a new event group.
 *
 * Internally, within the FreeRTOS implementation, event groups use a [small]
 * block of memory, in which the event group's structure is stored.  If an event
 * groups is created using xEventGroupCreate() then the required memory is
 * automatically dynamically allocated inside the xEventGroupCreate() function.
 * (see https://www.FreeRTOS.org/a00111.html).  If an event group is created
 * using xEventGroupCreateStatic() then the application writer must instead
 * provide the memory that will get used by the event group.
 * xEventGroupCreateStatic() therefore allows an event group to be created
 * without using any dynamic memory allocation.
 *
 * Although event groups are not related to ticks, for internal implementation
 * reasons the number of bits available for use in an event group is dependent
 * on the configUSE_16_BIT_TICKS setting in FreeRTOSConfig.h.  If
 * configUSE_16_BIT_TICKS is 1 then each event group contains 8 usable bits (bit
 * 0 to bit 7).  If configUSE_16_BIT_TICKS is set to 0 then each event group has
 * 24 usable bits (bit 0 to bit 23).  The EventBits_t type is used to store
 * event bits within an event group.
 *
 * @param pxEventGroupBuffer pxEventGroupBuffer must point to a variable of type
 * StaticEventGroup_t, which will be then be used to hold the event group's data
 * structures, removing the need for the memory to be allocated dynamically.
 *
 * @return If the event group was created then a handle to the event group is
 * returned.  If pxEventGroupBuffer was NULL then NULL is returned.
 *
 * Example usage:
 * @code{c}
 *  // StaticEventGroup_t is a publicly accessible structure that has the same
 *  // size and alignment requirements as the real event group structure.  It is
 *  // provided as a mechanism for applications to know the size of the event
 *  // group (which is dependent on the architecture and configuration file
 *  // settings) without breaking the strict data hiding policy by exposing the
 *  // real event group internals.  This StaticEventGroup_t variable is passed
 *  // into the xSemaphoreCreateEventGroupStatic() function and is used to store
 *  // the event group's data structures
 *  StaticEventGroup_t xEventGroupBuffer;
 *
 *  // Create the event group without dynamically allocating any memory.
 *  xEventGroup = xEventGroupCreateStatic( &xEventGroupBuffer );
 * @endcode
 */
// llgo:link (*StaticEventGroupT).XEventGroupCreateStatic C.xEventGroupCreateStatic
func (recv_ *StaticEventGroupT) XEventGroupCreateStatic() EventGroupHandleT {
	return nil
}

/**
 *
 * [Potentially] block to wait for one or more bits to be set within a
 * previously created event group.
 *
 * This function cannot be called from an interrupt.
 *
 * @param xEventGroup The event group in which the bits are being tested.  The
 * event group must have previously been created using a call to
 * xEventGroupCreate().
 *
 * @param uxBitsToWaitFor A bitwise value that indicates the bit or bits to test
 * inside the event group.  For example, to wait for bit 0 and/or bit 2 set
 * uxBitsToWaitFor to 0x05.  To wait for bits 0 and/or bit 1 and/or bit 2 set
 * uxBitsToWaitFor to 0x07.  Etc.
 *
 * @param xClearOnExit If xClearOnExit is set to pdTRUE then any bits within
 * uxBitsToWaitFor that are set within the event group will be cleared before
 * xEventGroupWaitBits() returns if the wait condition was met (if the function
 * returns for a reason other than a timeout).  If xClearOnExit is set to
 * pdFALSE then the bits set in the event group are not altered when the call to
 * xEventGroupWaitBits() returns.
 *
 * @param xWaitForAllBits If xWaitForAllBits is set to pdTRUE then
 * xEventGroupWaitBits() will return when either all the bits in uxBitsToWaitFor
 * are set or the specified block time expires.  If xWaitForAllBits is set to
 * pdFALSE then xEventGroupWaitBits() will return when any one of the bits set
 * in uxBitsToWaitFor is set or the specified block time expires.  The block
 * time is specified by the xTicksToWait parameter.
 *
 * @param xTicksToWait The maximum amount of time (specified in 'ticks') to wait
 * for one/all (depending on the xWaitForAllBits value) of the bits specified by
 * uxBitsToWaitFor to become set. A value of portMAX_DELAY can be used to block
 * indefinitely (provided INCLUDE_vTaskSuspend is set to 1 in FreeRTOSConfig.h).
 *
 * @return The value of the event group at the time either the bits being waited
 * for became set, or the block time expired.  Test the return value to know
 * which bits were set.  If xEventGroupWaitBits() returned because its timeout
 * expired then not all the bits being waited for will be set.  If
 * xEventGroupWaitBits() returned because the bits it was waiting for were set
 * then the returned value is the event group value before any bits were
 * automatically cleared in the case that xClearOnExit parameter was set to
 * pdTRUE.
 *
 * Example usage:
 * @code{c}
 * #define BIT_0 ( 1 << 0 )
 * #define BIT_4 ( 1 << 4 )
 *
 * void aFunction( EventGroupHandle_t xEventGroup )
 * {
 * EventBits_t uxBits;
 * const TickType_t xTicksToWait = 100 / portTICK_PERIOD_MS;
 *
 *      // Wait a maximum of 100ms for either bit 0 or bit 4 to be set within
 *      // the event group.  Clear the bits before exiting.
 *      uxBits = xEventGroupWaitBits(
 *                  xEventGroup,    // The event group being tested.
 *                  BIT_0 | BIT_4,  // The bits within the event group to wait for.
 *                  pdTRUE,         // BIT_0 and BIT_4 should be cleared before returning.
 *                  pdFALSE,        // Don't wait for both bits, either bit will do.
 *                  xTicksToWait ); // Wait a maximum of 100ms for either bit to be set.
 *
 *      if( ( uxBits & ( BIT_0 | BIT_4 ) ) == ( BIT_0 | BIT_4 ) )
 *      {
 *          // xEventGroupWaitBits() returned because both bits were set.
 *      }
 *      else if( ( uxBits & BIT_0 ) != 0 )
 *      {
 *          // xEventGroupWaitBits() returned because just BIT_0 was set.
 *      }
 *      else if( ( uxBits & BIT_4 ) != 0 )
 *      {
 *          // xEventGroupWaitBits() returned because just BIT_4 was set.
 *      }
 *      else
 *      {
 *          // xEventGroupWaitBits() returned because xTicksToWait ticks passed
 *          // without either BIT_0 or BIT_4 becoming set.
 *      }
 * }
 * @endcode
 * \ingroup EventGroup
 */
//go:linkname XEventGroupWaitBits C.xEventGroupWaitBits
func XEventGroupWaitBits(xEventGroup EventGroupHandleT, uxBitsToWaitFor EventBitsT, xClearOnExit BaseTypeT, xWaitForAllBits BaseTypeT, xTicksToWait TickTypeT) EventBitsT

/**
 *
 * Clear bits within an event group.  This function cannot be called from an
 * interrupt.
 *
 * @param xEventGroup The event group in which the bits are to be cleared.
 *
 * @param uxBitsToClear A bitwise value that indicates the bit or bits to clear
 * in the event group.  For example, to clear bit 3 only, set uxBitsToClear to
 * 0x08.  To clear bit 3 and bit 0 set uxBitsToClear to 0x09.
 *
 * @return The value of the event group before the specified bits were cleared.
 *
 * Example usage:
 * @code{c}
 * #define BIT_0 ( 1 << 0 )
 * #define BIT_4 ( 1 << 4 )
 *
 * void aFunction( EventGroupHandle_t xEventGroup )
 * {
 * EventBits_t uxBits;
 *
 *      // Clear bit 0 and bit 4 in xEventGroup.
 *      uxBits = xEventGroupClearBits(
 *                              xEventGroup,    // The event group being updated.
 *                              BIT_0 | BIT_4 );// The bits being cleared.
 *
 *      if( ( uxBits & ( BIT_0 | BIT_4 ) ) == ( BIT_0 | BIT_4 ) )
 *      {
 *          // Both bit 0 and bit 4 were set before xEventGroupClearBits() was
 *          // called.  Both will now be clear (not set).
 *      }
 *      else if( ( uxBits & BIT_0 ) != 0 )
 *      {
 *          // Bit 0 was set before xEventGroupClearBits() was called.  It will
 *          // now be clear.
 *      }
 *      else if( ( uxBits & BIT_4 ) != 0 )
 *      {
 *          // Bit 4 was set before xEventGroupClearBits() was called.  It will
 *          // now be clear.
 *      }
 *      else
 *      {
 *          // Neither bit 0 nor bit 4 were set in the first place.
 *      }
 * }
 * @endcode
 * \ingroup EventGroup
 */
//go:linkname XEventGroupClearBits C.xEventGroupClearBits
func XEventGroupClearBits(xEventGroup EventGroupHandleT, uxBitsToClear EventBitsT) EventBitsT

/**
 *
 * Set bits within an event group.
 * This function cannot be called from an interrupt.  xEventGroupSetBitsFromISR()
 * is a version that can be called from an interrupt.
 *
 * Setting bits in an event group will automatically unblock tasks that are
 * blocked waiting for the bits.
 *
 * @param xEventGroup The event group in which the bits are to be set.
 *
 * @param uxBitsToSet A bitwise value that indicates the bit or bits to set.
 * For example, to set bit 3 only, set uxBitsToSet to 0x08.  To set bit 3
 * and bit 0 set uxBitsToSet to 0x09.
 *
 * @return The value of the event group at the time the call to
 * xEventGroupSetBits() returns.  There are two reasons why the returned value
 * might have the bits specified by the uxBitsToSet parameter cleared.  First,
 * if setting a bit results in a task that was waiting for the bit leaving the
 * blocked state then it is possible the bit will be cleared automatically
 * (see the xClearBitOnExit parameter of xEventGroupWaitBits()).  Second, any
 * unblocked (or otherwise Ready state) task that has a priority above that of
 * the task that called xEventGroupSetBits() will execute and may change the
 * event group value before the call to xEventGroupSetBits() returns.
 *
 * Example usage:
 * @code{c}
 * #define BIT_0 ( 1 << 0 )
 * #define BIT_4 ( 1 << 4 )
 *
 * void aFunction( EventGroupHandle_t xEventGroup )
 * {
 * EventBits_t uxBits;
 *
 *      // Set bit 0 and bit 4 in xEventGroup.
 *      uxBits = xEventGroupSetBits(
 *                          xEventGroup,    // The event group being updated.
 *                          BIT_0 | BIT_4 );// The bits being set.
 *
 *      if( ( uxBits & ( BIT_0 | BIT_4 ) ) == ( BIT_0 | BIT_4 ) )
 *      {
 *          // Both bit 0 and bit 4 remained set when the function returned.
 *      }
 *      else if( ( uxBits & BIT_0 ) != 0 )
 *      {
 *          // Bit 0 remained set when the function returned, but bit 4 was
 *          // cleared.  It might be that bit 4 was cleared automatically as a
 *          // task that was waiting for bit 4 was removed from the Blocked
 *          // state.
 *      }
 *      else if( ( uxBits & BIT_4 ) != 0 )
 *      {
 *          // Bit 4 remained set when the function returned, but bit 0 was
 *          // cleared.  It might be that bit 0 was cleared automatically as a
 *          // task that was waiting for bit 0 was removed from the Blocked
 *          // state.
 *      }
 *      else
 *      {
 *          // Neither bit 0 nor bit 4 remained set.  It might be that a task
 *          // was waiting for both of the bits to be set, and the bits were
 *          // cleared as the task left the Blocked state.
 *      }
 * }
 * @endcode
 * \ingroup EventGroup
 */
//go:linkname XEventGroupSetBits C.xEventGroupSetBits
func XEventGroupSetBits(xEventGroup EventGroupHandleT, uxBitsToSet EventBitsT) EventBitsT

/**
 *
 * Atomically set bits within an event group, then wait for a combination of
 * bits to be set within the same event group.  This functionality is typically
 * used to synchronise multiple tasks, where each task has to wait for the other
 * tasks to reach a synchronisation point before proceeding.
 *
 * This function cannot be used from an interrupt.
 *
 * The function will return before its block time expires if the bits specified
 * by the uxBitsToWait parameter are set, or become set within that time.  In
 * this case all the bits specified by uxBitsToWait will be automatically
 * cleared before the function returns.
 *
 * @param xEventGroup The event group in which the bits are being tested.  The
 * event group must have previously been created using a call to
 * xEventGroupCreate().
 *
 * @param uxBitsToSet The bits to set in the event group before determining
 * if, and possibly waiting for, all the bits specified by the uxBitsToWait
 * parameter are set.
 *
 * @param uxBitsToWaitFor A bitwise value that indicates the bit or bits to test
 * inside the event group.  For example, to wait for bit 0 and bit 2 set
 * uxBitsToWaitFor to 0x05.  To wait for bits 0 and bit 1 and bit 2 set
 * uxBitsToWaitFor to 0x07.  Etc.
 *
 * @param xTicksToWait The maximum amount of time (specified in 'ticks') to wait
 * for all of the bits specified by uxBitsToWaitFor to become set.
 *
 * @return The value of the event group at the time either the bits being waited
 * for became set, or the block time expired.  Test the return value to know
 * which bits were set.  If xEventGroupSync() returned because its timeout
 * expired then not all the bits being waited for will be set.  If
 * xEventGroupSync() returned because all the bits it was waiting for were
 * set then the returned value is the event group value before any bits were
 * automatically cleared.
 *
 * Example usage:
 * @code{c}
 * // Bits used by the three tasks.
 * #define TASK_0_BIT     ( 1 << 0 )
 * #define TASK_1_BIT     ( 1 << 1 )
 * #define TASK_2_BIT     ( 1 << 2 )
 *
 * #define ALL_SYNC_BITS ( TASK_0_BIT | TASK_1_BIT | TASK_2_BIT )
 *
 * // Use an event group to synchronise three tasks.  It is assumed this event
 * // group has already been created elsewhere.
 * EventGroupHandle_t xEventBits;
 *
 * void vTask0( void *pvParameters )
 * {
 * EventBits_t uxReturn;
 * TickType_t xTicksToWait = 100 / portTICK_PERIOD_MS;
 *
 *   for( ;; )
 *   {
 *      // Perform task functionality here.
 *
 *      // Set bit 0 in the event flag to note this task has reached the
 *      // sync point.  The other two tasks will set the other two bits defined
 *      // by ALL_SYNC_BITS.  All three tasks have reached the synchronisation
 *      // point when all the ALL_SYNC_BITS are set.  Wait a maximum of 100ms
 *      // for this to happen.
 *      uxReturn = xEventGroupSync( xEventBits, TASK_0_BIT, ALL_SYNC_BITS, xTicksToWait );
 *
 *      if( ( uxReturn & ALL_SYNC_BITS ) == ALL_SYNC_BITS )
 *      {
 *          // All three tasks reached the synchronisation point before the call
 *          // to xEventGroupSync() timed out.
 *      }
 *  }
 * }
 *
 * void vTask1( void *pvParameters )
 * {
 *   for( ;; )
 *   {
 *      // Perform task functionality here.
 *
 *      // Set bit 1 in the event flag to note this task has reached the
 *      // synchronisation point.  The other two tasks will set the other two
 *      // bits defined by ALL_SYNC_BITS.  All three tasks have reached the
 *      // synchronisation point when all the ALL_SYNC_BITS are set.  Wait
 *      // indefinitely for this to happen.
 *      xEventGroupSync( xEventBits, TASK_1_BIT, ALL_SYNC_BITS, portMAX_DELAY );
 *
 *      // xEventGroupSync() was called with an indefinite block time, so
 *      // this task will only reach here if the synchronisation was made by all
 *      // three tasks, so there is no need to test the return value.
 *   }
 * }
 *
 * void vTask2( void *pvParameters )
 * {
 *   for( ;; )
 *   {
 *      // Perform task functionality here.
 *
 *      // Set bit 2 in the event flag to note this task has reached the
 *      // synchronisation point.  The other two tasks will set the other two
 *      // bits defined by ALL_SYNC_BITS.  All three tasks have reached the
 *      // synchronisation point when all the ALL_SYNC_BITS are set.  Wait
 *      // indefinitely for this to happen.
 *      xEventGroupSync( xEventBits, TASK_2_BIT, ALL_SYNC_BITS, portMAX_DELAY );
 *
 *      // xEventGroupSync() was called with an indefinite block time, so
 *      // this task will only reach here if the synchronisation was made by all
 *      // three tasks, so there is no need to test the return value.
 *  }
 * }
 *
 * @endcode
 * \ingroup EventGroup
 */
//go:linkname XEventGroupSync C.xEventGroupSync
func XEventGroupSync(xEventGroup EventGroupHandleT, uxBitsToSet EventBitsT, uxBitsToWaitFor EventBitsT, xTicksToWait TickTypeT) EventBitsT

/**
 *
 * A version of xEventGroupGetBits() that can be called from an ISR.
 *
 * @param xEventGroup The event group being queried.
 *
 * @return The event group bits at the time xEventGroupGetBitsFromISR() was called.
 *
 * \ingroup EventGroup
 */
//go:linkname XEventGroupGetBitsFromISR C.xEventGroupGetBitsFromISR
func XEventGroupGetBitsFromISR(xEventGroup EventGroupHandleT) EventBitsT

/**
 *
 * Delete an event group that was previously created by a call to
 * xEventGroupCreate().  Tasks that are blocked on the event group will be
 * unblocked and obtain 0 as the event group's value.
 *
 * @param xEventGroup The event group being deleted.
 */
//go:linkname VEventGroupDelete C.vEventGroupDelete
func VEventGroupDelete(xEventGroup EventGroupHandleT)

/**
 *
 * Retrieve a pointer to a statically created event groups's data structure
 * buffer. It is the same buffer that is supplied at the time of creation.
 *
 * @param xEventGroup The event group for which to retrieve the buffer.
 *
 * @param ppxEventGroupBuffer Used to return a pointer to the event groups's
 * data structure buffer.
 *
 * @return pdTRUE if the buffer was retrieved, pdFALSE otherwise.
 */
//go:linkname XEventGroupGetStaticBuffer C.xEventGroupGetStaticBuffer
func XEventGroupGetStaticBuffer(xEventGroup EventGroupHandleT, ppxEventGroupBuffer **StaticEventGroupT) BaseTypeT

/* For internal use only. */
//go:linkname VEventGroupSetBitsCallback C.vEventGroupSetBitsCallback
func VEventGroupSetBitsCallback(pvEventGroup c.Pointer, ulBitsToSet c.Uint32T)

//go:linkname VEventGroupClearBitsCallback C.vEventGroupClearBitsCallback
func VEventGroupClearBitsCallback(pvEventGroup c.Pointer, ulBitsToClear c.Uint32T)
