package quickjs

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:link (*JSContext).JsInitModuleStd C.js_init_module_std
func (recv_ *JSContext) JsInitModuleStd(module_name *c.Char) *JSModuleDef {
	return nil
}

// llgo:link (*JSContext).JsInitModuleOs C.js_init_module_os
func (recv_ *JSContext) JsInitModuleOs(module_name *c.Char) *JSModuleDef {
	return nil
}

// llgo:link (*JSContext).JsStdAddHelpers C.js_std_add_helpers
func (recv_ *JSContext) JsStdAddHelpers(argc c.Int, argv **c.Char) {
}

// llgo:link (*JSContext).JsStdLoop C.js_std_loop
func (recv_ *JSContext) JsStdLoop() {
}

// llgo:link (*JSRuntime).JsStdInitHandlers C.js_std_init_handlers
func (recv_ *JSRuntime) JsStdInitHandlers() {
}

// llgo:link (*JSRuntime).JsStdFreeHandlers C.js_std_free_handlers
func (recv_ *JSRuntime) JsStdFreeHandlers() {
}

// llgo:link (*JSContext).JsStdDumpError C.js_std_dump_error
func (recv_ *JSContext) JsStdDumpError() {
}

// llgo:link (*JSContext).JsLoadFile C.js_load_file
func (recv_ *JSContext) JsLoadFile(pbuf_len *c.SizeT, filename *c.Char) *c.Uint8T {
	return nil
}

// llgo:link (*JSContext).JsModuleSetImportMeta C.js_module_set_import_meta
func (recv_ *JSContext) JsModuleSetImportMeta(func_val JSValue, use_realpath c.Int, is_main c.Int) c.Int {
	return 0
}

// llgo:link (*JSContext).JsModuleLoader C.js_module_loader
func (recv_ *JSContext) JsModuleLoader(module_name *c.Char, opaque c.Pointer) *JSModuleDef {
	return nil
}

// llgo:link (*JSContext).JsStdEvalBinary C.js_std_eval_binary
func (recv_ *JSContext) JsStdEvalBinary(buf *c.Uint8T, buf_len c.SizeT, flags c.Int) {
}

// llgo:link (*JSContext).JsStdPromiseRejectionTracker C.js_std_promise_rejection_tracker
func (recv_ *JSContext) JsStdPromiseRejectionTracker(promise JSValue, reason JSValue, is_handled c.Int, opaque c.Pointer) {
}

//go:linkname JsStdSetWorkerNewContextFunc C.js_std_set_worker_new_context_func
func JsStdSetWorkerNewContextFunc(func_ func(*JSRuntime) *JSContext)
