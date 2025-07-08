package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*
 * Definition of the only type of object that a list can contain.
 */

type XLIST struct {
	UxNumberOfItems UBaseTypeT
	PxIndex         *ListItemT
	XListEnd        MiniListItemT
}

type XLISTITEM struct {
	XItemValue  TickTypeT
	PxNext      *XLISTITEM
	PxPrevious  *XLISTITEM
	PvOwner     c.Pointer
	PxContainer *XLIST
}
type ListItemT XLISTITEM

type XMINILISTITEM struct {
	XItemValue TickTypeT
	PxNext     *XLISTITEM
	PxPrevious *XLISTITEM
}
type MiniListItemT XMINILISTITEM
type ListT XLIST

/*
 * Must be called before a list is used!  This initialises all the members
 * of the list structure and inserts the xListEnd item into the list as a
 * marker to the back of the list.
 *
 * @param pxList Pointer to the list being initialised.
 *
 * \page vListInitialise vListInitialise
 * \ingroup LinkedList
 */
// llgo:link (*ListT).VListInitialise C.vListInitialise
func (recv_ *ListT) VListInitialise() {
}

/*
 * Must be called before a list item is used.  This sets the list container to
 * null so the item does not think that it is already contained in a list.
 *
 * @param pxItem Pointer to the list item being initialised.
 *
 * \page vListInitialiseItem vListInitialiseItem
 * \ingroup LinkedList
 */
// llgo:link (*ListItemT).VListInitialiseItem C.vListInitialiseItem
func (recv_ *ListItemT) VListInitialiseItem() {
}

/*
 * Insert a list item into a list.  The item will be inserted into the list in
 * a position determined by its item value (ascending item value order).
 *
 * @param pxList The list into which the item is to be inserted.
 *
 * @param pxNewListItem The item that is to be placed in the list.
 *
 * \page vListInsert vListInsert
 * \ingroup LinkedList
 */
// llgo:link (*ListT).VListInsert C.vListInsert
func (recv_ *ListT) VListInsert(pxNewListItem *ListItemT) {
}

/*
 * Insert a list item into a list.  The item will be inserted in a position
 * such that it will be the last item within the list returned by multiple
 * calls to listGET_OWNER_OF_NEXT_ENTRY.
 *
 * The list member pxIndex is used to walk through a list.  Calling
 * listGET_OWNER_OF_NEXT_ENTRY increments pxIndex to the next item in the list.
 * Placing an item in a list using vListInsertEnd effectively places the item
 * in the list position pointed to by pxIndex.  This means that every other
 * item within the list will be returned by listGET_OWNER_OF_NEXT_ENTRY before
 * the pxIndex parameter again points to the item being inserted.
 *
 * @param pxList The list into which the item is to be inserted.
 *
 * @param pxNewListItem The list item to be inserted into the list.
 *
 * \page vListInsertEnd vListInsertEnd
 * \ingroup LinkedList
 */
// llgo:link (*ListT).VListInsertEnd C.vListInsertEnd
func (recv_ *ListT) VListInsertEnd(pxNewListItem *ListItemT) {
}

/*
 * Remove an item from a list.  The list item has a pointer to the list that
 * it is in, so only the list item need be passed into the function.
 *
 * @param uxListRemove The item to be removed.  The item will remove itself from
 * the list pointed to by it's pxContainer parameter.
 *
 * @return The number of items that remain in the list after the list item has
 * been removed.
 *
 * \page uxListRemove uxListRemove
 * \ingroup LinkedList
 */
// llgo:link (*ListItemT).UxListRemove C.uxListRemove
func (recv_ *ListItemT) UxListRemove() UBaseTypeT {
	return 0
}
