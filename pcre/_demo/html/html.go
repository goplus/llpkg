package main

import (
	"fmt"
	"unsafe"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/pcre"
)

func main() {
	pattern := "Hello (\\w+)"
	subject := "Hello World, test."

	patternC := c.AllocaCStr(pattern)

	subjectC := c.AllocaCStr(subject)

	var (
		errPtr    *int8
		erroffset c.Int
		errorcode c.Int
	)

	re := pcre.Compile2(
		(*int8)(unsafe.Pointer(patternC)),
		0,
		&errorcode,
		&errPtr,
		&erroffset,
		nil,
	)

	if re == nil {
		fmt.Printf("Compile fail (code=%d) at offset %d: %s\n",
			errorcode,
			erroffset,
			c.GoString((*c.Char)(unsafe.Pointer(errPtr))),
		)
		return
	}

	var ovector [30]c.Int

	rc := re.Exec(
		nil,
		(*int8)(unsafe.Pointer(subjectC)),
		c.Int(len(subject)),
		0,
		0,
		&ovector[0],
		c.Int(len(ovector)),
	)

	if rc < 0 {
		fmt.Printf("Match fail: %d\n", rc)
		return
	}

	for i := 0; i < int(rc); i++ {
		var substrPtr *int8

		ret := pcre.GetSubstring(
			(*int8)(unsafe.Pointer(subjectC)),
			&ovector[0],
			c.Int(rc),
			c.Int(i),
			&substrPtr,
		)
		if ret < 0 {
			fmt.Printf("Submatch fail: %d\n", ret)
			continue
		}
		matchedStr := c.GoString((*c.Char)(unsafe.Pointer(substrPtr)))
		fmt.Printf("Capture group #%d = %s\n", i, matchedStr)

		pcre.FreeSubstring(substrPtr)
	}
}
