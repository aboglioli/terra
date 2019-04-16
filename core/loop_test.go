package core

import (
	"testing"
)

func TestLoop(t *testing.T) {
	frames := 0
	duration := 0.0
	done := make(chan bool)

	loop := NewLoop(5000, func(d float64) {
		frames++
		duration += d

		if frames == 500 {
			done <- true
		}
	})

	go loop.Start()
	<-done
	loop.Stop()

	if duration < 100 || duration > 101 {
		t.Errorf("%d frames processed in %.2fms", frames, duration)
	}
}
