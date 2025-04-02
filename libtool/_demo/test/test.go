package main

import (
	"fmt"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/libtool"
)

func main() {
	fmt.Println("Simple libtool demonstration")

	// Initialize libtool
	ret := libtool.Init()
	if ret != 0 {
		fmt.Println("Failed to initialize libtool:", c.GoString(libtool.Error()))
		return
	}
	fmt.Println("Successfully initialized libtool")

	// Try to load a common library (libc)
	libName := "libc.so.6" // Linux style
	handle := libtool.Open(c.Str(libName))
	if handle == nil {
		libName = "libc.dylib" // macOS style
		handle = libtool.Open(c.Str(libName))
	}
	if handle == nil {
		libName = "c" // Generic style
		handle = libtool.Open(c.Str(libName))
	}

	if handle != nil {
		fmt.Printf("Successfully opened %s\n", libName)

		// Try to find a common function (printf)
		symPtr := libtool.Sym(handle, c.Str("printf"))
		if symPtr != nil {
			fmt.Println("Found 'printf' function")
		} else {
			fmt.Println("Symbol 'printf' not found:", c.GoString(libtool.Error()))
		}

		// Close the library
		libtool.Close(handle)
		fmt.Println("Closed library")
	} else {
		fmt.Println("Could not open any standard library:", c.GoString(libtool.Error()))
	}

	// Clean up libtool
	libtool.Exit()
	fmt.Println("Successfully cleaned up libtool")
}
