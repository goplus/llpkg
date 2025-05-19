package main

import (
	"github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/quickjs"
)

func main() {
	rt := quickjs.JSNewRuntime()
	if rt == nil {
		panic("none")
	}
	rt.JsStdInitHandlers()
	ctx := rt.JSNewContext()
	if ctx == nil {
		panic("none")
	}
	ctx.JsStdAddHelpers(0, nil)
	jsCode := c.Str("console.log('Hello from QuickJS! 1 + 2 =', 1 + 2);")
	result := ctx.JSEval(jsCode, c.Strlen(jsCode), c.Str("<input>"), 0)
	ctx.X__JSFreeValue(result)
	ctx.JSFreeContext()
	rt.JSFreeRuntime()
}
