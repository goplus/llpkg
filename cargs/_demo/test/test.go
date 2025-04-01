package main

import (
	"fmt"
	"os"

	C "github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/cargs"
)

func main() {
	// Define command-line options
	options := []cargs.CagOption{
		{
			Identifier:    'h',
			AccessLetters: C.Str("h"),
			AccessName:    C.Str("help"),
			ValueName:     nil,
			Description:   C.Str("Show help information"),
		},
		{
			Identifier:    'v',
			AccessLetters: C.Str("v"),
			AccessName:    C.Str("version"),
			ValueName:     nil,
			Description:   C.Str("Show version information"),
		},
	}

	args := os.Args

	// Convert Go string array to C-style argv
	argv := make([]*int8, len(args))
	for i, arg := range args {
		argv[i] = C.AllocaCStr(arg)
	}

	// Initialize option context
	var context cargs.CagOptionContext
	context.CagOptionInit(&options[0], uintptr(len(options)), C.Int(len(args)), &argv[0])

	// Process all options
	identifierFound := false
	for context.CagOptionFetch() {
		identifierFound = true
		identifier := context.CagOptionGetIdentifier()
		switch identifier {
		case 'h':
			fmt.Println("Help: This is a simple command-line parser demo")
		case 'v':
			fmt.Println("Version: 1.0.0")
		}
	}

	// Default output if no identifier is found
	if !identifierFound {
		fmt.Println("Demo Command-line Tool\nIdentifier:\n\t-h: Help\n\t-v: Version")
	}
}
