package freertos

import (
	"github.com/goplus/lib/c/time"
	_ "unsafe"
)

type Event struct {
	Unused [8]uint8
}

//go:linkname EventCreate C.event_create
func EventCreate() *Event

// llgo:link (*Event).EventDelete C.event_delete
func (recv_ *Event) EventDelete() {
}

// llgo:link (*Event).EventWait C.event_wait
func (recv_ *Event) EventWait() bool {
	return false
}

// llgo:link (*Event).EventWaitTimed C.event_wait_timed
func (recv_ *Event) EventWaitTimed(ms time.TimeT) bool {
	return false
}

// llgo:link (*Event).EventSignal C.event_signal
func (recv_ *Event) EventSignal() {
}
