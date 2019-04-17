package core

import "time"

type Loop struct {
	fps      int
	onUpdate func(float64)
	done     chan bool
}

func NewLoop(fps int, onUpdate func(float64)) *Loop {
	return &Loop{
		fps:      fps,
		onUpdate: onUpdate,
		done:     make(chan bool),
	}
}

func (l *Loop) Start() {
	t := 1e9 / float64(l.fps)
	ticker := time.NewTicker(time.Duration(t) * time.Nanosecond)

	timeStart := time.Now().UnixNano()

	for {
		select {
		case <-ticker.C:
			now := time.Now().UnixNano()
			delta := float64(now-timeStart) / float64(time.Millisecond)
			timeStart = now

			l.onUpdate(delta)
		case <-l.done:
			ticker.Stop()
			return
		}
	}
}

func (l *Loop) Stop() {
	l.done <- true
}
