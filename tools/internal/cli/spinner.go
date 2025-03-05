package cli

import (
	"fmt"
	"time"
)

type loadingSpinner struct {
	stopChan chan struct{}
	message  string
}

func NewLoadingSpinner(message string) *loadingSpinner {
	return &loadingSpinner{
		stopChan: make(chan struct{}),
		message:  message,
	}
}

func (s *loadingSpinner) Start() {
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	go func() {
		for i := 0; ; i++ {
			select {
			case <-s.stopChan:
				fmt.Printf("\r%s... Done!    \n", s.message)
				return
			default:
				frame := frames[i%len(frames)]
				fmt.Printf("\r%s %s ", frame, s.message)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
}

func (s *loadingSpinner) Stop() {
	s.stopChan <- struct{}{}
}
