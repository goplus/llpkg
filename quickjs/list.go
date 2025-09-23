package quickjs

import _ "unsafe"

type ListHead struct {
	Prev *ListHead
	Next *ListHead
}
