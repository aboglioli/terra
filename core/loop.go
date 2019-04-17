package core

import "time"

type Loop struct {
	fps      int
	onUpdate func(float64)
	done     bool
}

func NewLoop(fps int, onUpdate func(float64)) *Loop {
	return &Loop{
		fps:      fps,
		onUpdate: onUpdate,
		done:     false,
	}
}

func (l *Loop) Start() {
	t := 1e9 / float64(l.fps)
	ticker := time.NewTicker(time.Duration(t) * time.Nanosecond)

	timeStart := time.Now().UnixNano()

	for range ticker.C {
		if l.done {
			return
		}

		now := time.Now().UnixNano()
		delta := float64(now-timeStart) / float64(time.Millisecond)
		timeStart = now

		l.onUpdate(delta)
	}
}

func (l *Loop) Stop() {
	l.done = true
}
