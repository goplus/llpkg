package cjson

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const VERSION_MAJOR = 1
const VERSION_MINOR = 7
const VERSION_PATCH = 18
const IsReference = 256
const StringIsConst = 512
const NESTING_LIMIT = 1000

/* The cJSON structure: */

type CJSON struct {
	Next        *CJSON
	Prev        *CJSON
	Child       *CJSON
	Type        c.Int
	Valuestring *c.Char
	Valueint    c.Int
	Valuedouble c.Double
	String      *c.Char
}

type Hooks struct {
	MallocFn c.Pointer
	FreeFn   c.Pointer
}
type Bool c.Int

/* returns the version of cJSON as a string */
//go:linkname Version C.cJSON_Version
func Version() *c.Char

/* Supply malloc, realloc and free functions to cJSON */
// llgo:link (*Hooks).InitHooks C.cJSON_InitHooks
func (recv_ *Hooks) InitHooks() {
}

/* Memory Management: the caller is always responsible to free the results from all variants of cJSON_Parse (with cJSON_Delete) and cJSON_Print (with stdlib free, cJSON_Hooks.free_fn, or cJSON_free as appropriate). The exception is cJSON_PrintPreallocated, where the caller has full responsibility of the buffer. */
/* Supply a block of JSON, and this returns a cJSON object you can interrogate. */
//go:linkname Parse C.cJSON_Parse
func Parse(value *c.Char) *CJSON

//go:linkname ParseWithLength C.cJSON_ParseWithLength
func ParseWithLength(value *c.Char, buffer_length c.SizeT) *CJSON

/* ParseWithOpts allows you to require (and check) that the JSON is null terminated, and to retrieve the pointer to the final byte parsed. */
/* If you supply a ptr in return_parse_end and parsing fails, then return_parse_end will contain a pointer to the error so will match cJSON_GetErrorPtr(). */
//go:linkname ParseWithOpts C.cJSON_ParseWithOpts
func ParseWithOpts(value *c.Char, return_parse_end **c.Char, require_null_terminated Bool) *CJSON

//go:linkname ParseWithLengthOpts C.cJSON_ParseWithLengthOpts
func ParseWithLengthOpts(value *c.Char, buffer_length c.SizeT, return_parse_end **c.Char, require_null_terminated Bool) *CJSON

/* Render a cJSON entity to text for transfer/storage. */
// llgo:link (*CJSON).Print C.cJSON_Print
func (recv_ *CJSON) Print() *c.Char {
	return nil
}

/* Render a cJSON entity to text for transfer/storage without any formatting. */
// llgo:link (*CJSON).CStr C.cJSON_PrintUnformatted
func (recv_ *CJSON) CStr() *c.Char {
	return nil
}

/* Render a cJSON entity to text using a buffered strategy. prebuffer is a guess at the final size. guessing well reduces reallocation. fmt=0 gives unformatted, =1 gives formatted */
// llgo:link (*CJSON).PrintBuffered C.cJSON_PrintBuffered
func (recv_ *CJSON) PrintBuffered(prebuffer c.Int, fmt Bool) *c.Char {
	return nil
}

/* Render a cJSON entity to text using a buffer already allocated in memory with given length. Returns 1 on success and 0 on failure. */
/* NOTE: cJSON is not always 100% accurate in estimating how much memory it will use, so to be safe allocate 5 bytes more than you actually need */
// llgo:link (*CJSON).PrintPreallocated C.cJSON_PrintPreallocated
func (recv_ *CJSON) PrintPreallocated(buffer *c.Char, length c.Int, format Bool) Bool {
	return 0
}

/* Delete a cJSON entity and all subentities. */
// llgo:link (*CJSON).Delete C.cJSON_Delete
func (recv_ *CJSON) Delete() {
}

/* Returns the number of items in an array (or object). */
// llgo:link (*CJSON).GetArraySize C.cJSON_GetArraySize
func (recv_ *CJSON) GetArraySize() c.Int {
	return 0
}

/* Retrieve item number "index" from array "array". Returns NULL if unsuccessful. */
// llgo:link (*CJSON).GetArrayItem C.cJSON_GetArrayItem
func (recv_ *CJSON) GetArrayItem(index c.Int) *CJSON {
	return nil
}

/* Get item "string" from object. Case insensitive. */
// llgo:link (*CJSON).GetObjectItem C.cJSON_GetObjectItem
func (recv_ *CJSON) GetObjectItem(string *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).GetObjectItemCaseSensitive C.cJSON_GetObjectItemCaseSensitive
func (recv_ *CJSON) GetObjectItemCaseSensitive(string *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).HasObjectItem C.cJSON_HasObjectItem
func (recv_ *CJSON) HasObjectItem(string *c.Char) Bool {
	return 0
}

/* For analysing failed parses. This returns a pointer to the parse error. You'll probably need to look a few chars back to make sense of it. Defined when cJSON_Parse() returns 0. 0 when cJSON_Parse() succeeds. */
//go:linkname GetErrorPtr C.cJSON_GetErrorPtr
func GetErrorPtr() *c.Char

/* Check item type and return its value */
// llgo:link (*CJSON).GetStringValue C.cJSON_GetStringValue
func (recv_ *CJSON) GetStringValue() *c.Char {
	return nil
}

// llgo:link (*CJSON).GetNumberValue C.cJSON_GetNumberValue
func (recv_ *CJSON) GetNumberValue() c.Double {
	return 0
}

/* These functions check the type of an item */
// llgo:link (*CJSON).IsInvalid C.cJSON_IsInvalid
func (recv_ *CJSON) IsInvalid() Bool {
	return 0
}

// llgo:link (*CJSON).IsFalse C.cJSON_IsFalse
func (recv_ *CJSON) IsFalse() Bool {
	return 0
}

// llgo:link (*CJSON).IsTrue C.cJSON_IsTrue
func (recv_ *CJSON) IsTrue() Bool {
	return 0
}

// llgo:link (*CJSON).IsBool C.cJSON_IsBool
func (recv_ *CJSON) IsBool() Bool {
	return 0
}

// llgo:link (*CJSON).IsNull C.cJSON_IsNull
func (recv_ *CJSON) IsNull() Bool {
	return 0
}

// llgo:link (*CJSON).IsNumber C.cJSON_IsNumber
func (recv_ *CJSON) IsNumber() Bool {
	return 0
}

// llgo:link (*CJSON).IsString C.cJSON_IsString
func (recv_ *CJSON) IsString() Bool {
	return 0
}

// llgo:link (*CJSON).IsArray C.cJSON_IsArray
func (recv_ *CJSON) IsArray() Bool {
	return 0
}

// llgo:link (*CJSON).IsObject C.cJSON_IsObject
func (recv_ *CJSON) IsObject() Bool {
	return 0
}

// llgo:link (*CJSON).IsRaw C.cJSON_IsRaw
func (recv_ *CJSON) IsRaw() Bool {
	return 0
}

/* These calls create a cJSON item of the appropriate type. */
//go:linkname CreateNull C.cJSON_CreateNull
func CreateNull() *CJSON

//go:linkname CreateTrue C.cJSON_CreateTrue
func CreateTrue() *CJSON

//go:linkname CreateFalse C.cJSON_CreateFalse
func CreateFalse() *CJSON

// llgo:link Bool.CreateBool C.cJSON_CreateBool
func (recv_ Bool) CreateBool() *CJSON {
	return nil
}

//go:linkname CreateNumber C.cJSON_CreateNumber
func CreateNumber(num c.Double) *CJSON

//go:linkname String C.cJSON_CreateString
func String(string *c.Char) *CJSON

/* raw json */
//go:linkname CreateRaw C.cJSON_CreateRaw
func CreateRaw(raw *c.Char) *CJSON

//go:linkname Array C.cJSON_CreateArray
func Array() *CJSON

//go:linkname Object C.cJSON_CreateObject
func Object() *CJSON

/* Create a string where valuestring references a string so
 * it will not be freed by cJSON_Delete */
//go:linkname CreateStringReference C.cJSON_CreateStringReference
func CreateStringReference(string *c.Char) *CJSON

/* Create an object/array that only references it's elements so
 * they will not be freed by cJSON_Delete */
// llgo:link (*CJSON).CreateObjectReference C.cJSON_CreateObjectReference
func (recv_ *CJSON) CreateObjectReference() *CJSON {
	return nil
}

// llgo:link (*CJSON).CreateArrayReference C.cJSON_CreateArrayReference
func (recv_ *CJSON) CreateArrayReference() *CJSON {
	return nil
}

/* These utilities create an Array of count items.
 * The parameter count cannot be greater than the number of elements in the number array, otherwise array access will be out of bounds.*/
//go:linkname CreateIntArray C.cJSON_CreateIntArray
func CreateIntArray(numbers *c.Int, count c.Int) *CJSON

//go:linkname CreateFloatArray C.cJSON_CreateFloatArray
func CreateFloatArray(numbers *c.Float, count c.Int) *CJSON

//go:linkname CreateDoubleArray C.cJSON_CreateDoubleArray
func CreateDoubleArray(numbers *c.Double, count c.Int) *CJSON

//go:linkname CreateStringArray C.cJSON_CreateStringArray
func CreateStringArray(strings **c.Char, count c.Int) *CJSON

/* Append item to the specified array/object. */
// llgo:link (*CJSON).AddItem C.cJSON_AddItemToArray
func (recv_ *CJSON) AddItem(item *CJSON) Bool {
	return 0
}

// llgo:link (*CJSON).SetItem C.cJSON_AddItemToObject
func (recv_ *CJSON) SetItem(string *c.Char, item *CJSON) Bool {
	return 0
}

/* Use this when string is definitely const (i.e. a literal, or as good as), and will definitely survive the cJSON object.
 * WARNING: When this function was used, make sure to always check that (item->type & cJSON_StringIsConst) is zero before
 * writing to `item->string` */
// llgo:link (*CJSON).AddItemToObjectCS C.cJSON_AddItemToObjectCS
func (recv_ *CJSON) AddItemToObjectCS(string *c.Char, item *CJSON) Bool {
	return 0
}

/* Append reference to item to the specified array/object. Use this when you want to add an existing cJSON to a new cJSON, but don't want to corrupt your existing cJSON. */
// llgo:link (*CJSON).AddItemReferenceToArray C.cJSON_AddItemReferenceToArray
func (recv_ *CJSON) AddItemReferenceToArray(item *CJSON) Bool {
	return 0
}

// llgo:link (*CJSON).AddItemReferenceToObject C.cJSON_AddItemReferenceToObject
func (recv_ *CJSON) AddItemReferenceToObject(string *c.Char, item *CJSON) Bool {
	return 0
}

/* Remove/Detach items from Arrays/Objects. */
// llgo:link (*CJSON).DetachItemViaPointer C.cJSON_DetachItemViaPointer
func (recv_ *CJSON) DetachItemViaPointer(item *CJSON) *CJSON {
	return nil
}

// llgo:link (*CJSON).DetachItemFromArray C.cJSON_DetachItemFromArray
func (recv_ *CJSON) DetachItemFromArray(which c.Int) *CJSON {
	return nil
}

// llgo:link (*CJSON).DeleteItemFromArray C.cJSON_DeleteItemFromArray
func (recv_ *CJSON) DeleteItemFromArray(which c.Int) {
}

// llgo:link (*CJSON).DetachItemFromObject C.cJSON_DetachItemFromObject
func (recv_ *CJSON) DetachItemFromObject(string *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).DetachItemFromObjectCaseSensitive C.cJSON_DetachItemFromObjectCaseSensitive
func (recv_ *CJSON) DetachItemFromObjectCaseSensitive(string *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).DeleteItemFromObject C.cJSON_DeleteItemFromObject
func (recv_ *CJSON) DeleteItemFromObject(string *c.Char) {
}

// llgo:link (*CJSON).DeleteItemFromObjectCaseSensitive C.cJSON_DeleteItemFromObjectCaseSensitive
func (recv_ *CJSON) DeleteItemFromObjectCaseSensitive(string *c.Char) {
}

/* Update array items. */
// llgo:link (*CJSON).InsertItemInArray C.cJSON_InsertItemInArray
func (recv_ *CJSON) InsertItemInArray(which c.Int, newitem *CJSON) Bool {
	return 0
}

// llgo:link (*CJSON).ReplaceItemViaPointer C.cJSON_ReplaceItemViaPointer
func (recv_ *CJSON) ReplaceItemViaPointer(item *CJSON, replacement *CJSON) Bool {
	return 0
}

// llgo:link (*CJSON).ReplaceItemInArray C.cJSON_ReplaceItemInArray
func (recv_ *CJSON) ReplaceItemInArray(which c.Int, newitem *CJSON) Bool {
	return 0
}

// llgo:link (*CJSON).ReplaceItemInObject C.cJSON_ReplaceItemInObject
func (recv_ *CJSON) ReplaceItemInObject(string *c.Char, newitem *CJSON) Bool {
	return 0
}

// llgo:link (*CJSON).ReplaceItemInObjectCaseSensitive C.cJSON_ReplaceItemInObjectCaseSensitive
func (recv_ *CJSON) ReplaceItemInObjectCaseSensitive(string *c.Char, newitem *CJSON) Bool {
	return 0
}

/* Duplicate a cJSON item */
// llgo:link (*CJSON).Duplicate C.cJSON_Duplicate
func (recv_ *CJSON) Duplicate(recurse Bool) *CJSON {
	return nil
}

/* Duplicate will create a new, identical cJSON item to the one you pass, in new memory that will
 * need to be released. With recurse!=0, it will duplicate any children connected to the item.
 * The item->next and ->prev pointers are always zero on return from Duplicate. */
/* Recursively compare two cJSON items for equality. If either a or b is NULL or invalid, they will be considered unequal.
 * case_sensitive determines if object keys are treated case sensitive (1) or case insensitive (0) */
// llgo:link (*CJSON).Compare C.cJSON_Compare
func (recv_ *CJSON) Compare(b *CJSON, case_sensitive Bool) Bool {
	return 0
}

/* Minify a strings, remove blank characters(such as ' ', '\t', '\r', '\n') from strings.
 * The input pointer json cannot point to a read-only address area, such as a string constant,
 * but should point to a readable and writable address area. */
//go:linkname Minify C.cJSON_Minify
func Minify(json *c.Char)

/* Helper functions for creating and adding items to an object at the same time.
 * They return the added item or NULL on failure. */
// llgo:link (*CJSON).AddNullToObject C.cJSON_AddNullToObject
func (recv_ *CJSON) AddNullToObject(name *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).AddTrueToObject C.cJSON_AddTrueToObject
func (recv_ *CJSON) AddTrueToObject(name *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).AddFalseToObject C.cJSON_AddFalseToObject
func (recv_ *CJSON) AddFalseToObject(name *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).AddBoolToObject C.cJSON_AddBoolToObject
func (recv_ *CJSON) AddBoolToObject(name *c.Char, boolean Bool) *CJSON {
	return nil
}

// llgo:link (*CJSON).AddNumberToObject C.cJSON_AddNumberToObject
func (recv_ *CJSON) AddNumberToObject(name *c.Char, number c.Double) *CJSON {
	return nil
}

// llgo:link (*CJSON).AddStringToObject C.cJSON_AddStringToObject
func (recv_ *CJSON) AddStringToObject(name *c.Char, string *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).AddRawToObject C.cJSON_AddRawToObject
func (recv_ *CJSON) AddRawToObject(name *c.Char, raw *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).AddObjectToObject C.cJSON_AddObjectToObject
func (recv_ *CJSON) AddObjectToObject(name *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).AddArrayToObject C.cJSON_AddArrayToObject
func (recv_ *CJSON) AddArrayToObject(name *c.Char) *CJSON {
	return nil
}

/* helper for the cJSON_SetNumberValue macro */
// llgo:link (*CJSON).SetNumberHelper C.cJSON_SetNumberHelper
func (recv_ *CJSON) SetNumberHelper(number c.Double) c.Double {
	return 0
}

/* Change the valuestring of a cJSON_String object, only takes effect when type of object is cJSON_String */
// llgo:link (*CJSON).SetValuestring C.cJSON_SetValuestring
func (recv_ *CJSON) SetValuestring(valuestring *c.Char) *c.Char {
	return nil
}

/* malloc/free objects using the malloc/free functions that have been set with cJSON_InitHooks */
//go:linkname Malloc C.cJSON_malloc
func Malloc(size c.SizeT) c.Pointer

//go:linkname FreeCStr C.cJSON_free
func FreeCStr(object c.Pointer)
