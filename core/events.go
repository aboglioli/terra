package core

import (
	"github.com/veandco/go-sdl2/sdl"
)

type EventPool struct {
	quitEvents        []func()
	keyboardEvents    []func(*sdl.KeyboardEvent)
	mouseMotionEvents []func(*sdl.MouseMotionEvent)
	mouseButtonEvents []func(*sdl.MouseButtonEvent)
}

func NewEventPool() *EventPool {
	return &EventPool{
		quitEvents:        make([]func(), 0),
		keyboardEvents:    make([]func(*sdl.KeyboardEvent), 0),
		mouseMotionEvents: make([]func(*sdl.MouseMotionEvent), 0),
		mouseButtonEvents: make([]func(*sdl.MouseButtonEvent), 0),
	}
}

func (ep *EventPool) AddQuitEvent(fn func()) {
	ep.quitEvents = append(ep.quitEvents, fn)
}

func (ep *EventPool) AddKeyboardEvent(fn func(*sdl.KeyboardEvent)) {
	ep.keyboardEvents = append(ep.keyboardEvents, fn)
}

func (ep *EventPool) AddMouseMotionEvent(fn func(*sdl.MouseMotionEvent)) {
	ep.mouseMotionEvents = append(ep.mouseMotionEvents, fn)
}

func (ep *EventPool) AddMouseButtonEvent(fn func(*sdl.MouseButtonEvent)) {
	ep.mouseButtonEvents = append(ep.mouseButtonEvents, fn)
}

func (ep *EventPool) HandleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.KeyboardEvent:
			if e.Type == sdl.KEYDOWN {
				for _, ke := range ep.keyboardEvents {
					ke(e)
				}
			}
		case *sdl.MouseMotionEvent:
			for _, mme := range ep.mouseMotionEvents {
				mme(e)
			}
		case *sdl.MouseButtonEvent:
			if e.Type == sdl.MOUSEBUTTONDOWN {
				for _, mbe := range ep.mouseButtonEvents {
					mbe(e)
				}
			}
		case *sdl.QuitEvent:
			for _, qe := range ep.quitEvents {
				qe()
			}
		}
	}
}
