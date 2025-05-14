package cjson

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/* Implement RFC6901 (https://tools.ietf.org/html/rfc6901) JSON Pointer spec. */
// llgo:link (*CJSON).GetPointer C.cJSONUtils_GetPointer
func (recv_ *CJSON) GetPointer(pointer *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).GetPointerCaseSensitive C.cJSONUtils_GetPointerCaseSensitive
func (recv_ *CJSON) GetPointerCaseSensitive(pointer *c.Char) *CJSON {
	return nil
}

/* Implement RFC6902 (https://tools.ietf.org/html/rfc6902) JSON Patch spec. */
/* NOTE: This modifies objects in 'from' and 'to' by sorting the elements by their key */
// llgo:link (*CJSON).GeneratePatches C.cJSONUtils_GeneratePatches
func (recv_ *CJSON) GeneratePatches(to *CJSON) *CJSON {
	return nil
}

// llgo:link (*CJSON).GeneratePatchesCaseSensitive C.cJSONUtils_GeneratePatchesCaseSensitive
func (recv_ *CJSON) GeneratePatchesCaseSensitive(to *CJSON) *CJSON {
	return nil
}

/* Utility for generating patch array entries. */
// llgo:link (*CJSON).AddPatchToArray C.cJSONUtils_AddPatchToArray
func (recv_ *CJSON) AddPatchToArray(operation *c.Char, path *c.Char, value *CJSON) {
}

/* Returns 0 for success. */
// llgo:link (*CJSON).ApplyPatches C.cJSONUtils_ApplyPatches
func (recv_ *CJSON) ApplyPatches(patches *CJSON) c.Int {
	return 0
}

// llgo:link (*CJSON).ApplyPatchesCaseSensitive C.cJSONUtils_ApplyPatchesCaseSensitive
func (recv_ *CJSON) ApplyPatchesCaseSensitive(patches *CJSON) c.Int {
	return 0
}

/* Implement RFC7386 (https://tools.ietf.org/html/rfc7396) JSON Merge Patch spec. */
/* target will be modified by patch. return value is new ptr for target. */
// llgo:link (*CJSON).MergePatch C.cJSONUtils_MergePatch
func (recv_ *CJSON) MergePatch(patch *CJSON) *CJSON {
	return nil
}

// llgo:link (*CJSON).MergePatchCaseSensitive C.cJSONUtils_MergePatchCaseSensitive
func (recv_ *CJSON) MergePatchCaseSensitive(patch *CJSON) *CJSON {
	return nil
}

/* generates a patch to move from -> to */
/* NOTE: This modifies objects in 'from' and 'to' by sorting the elements by their key */
// llgo:link (*CJSON).GenerateMergePatch C.cJSONUtils_GenerateMergePatch
func (recv_ *CJSON) GenerateMergePatch(to *CJSON) *CJSON {
	return nil
}

// llgo:link (*CJSON).GenerateMergePatchCaseSensitive C.cJSONUtils_GenerateMergePatchCaseSensitive
func (recv_ *CJSON) GenerateMergePatchCaseSensitive(to *CJSON) *CJSON {
	return nil
}

/* Given a root object and a target object, construct a pointer from one to the other. */
// llgo:link (*CJSON).FindPointerFromObjectTo C.cJSONUtils_FindPointerFromObjectTo
func (recv_ *CJSON) FindPointerFromObjectTo(target *CJSON) *c.Char {
	return nil
}

/* Sorts the members of the object into alphabetical order. */
// llgo:link (*CJSON).SortObject C.cJSONUtils_SortObject
func (recv_ *CJSON) SortObject() {
}

// llgo:link (*CJSON).SortObjectCaseSensitive C.cJSONUtils_SortObjectCaseSensitive
func (recv_ *CJSON) SortObjectCaseSensitive() {
}
