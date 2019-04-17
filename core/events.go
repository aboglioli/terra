package core

import (
	"github.com/veandco/go-sdl2/sdl"
)

type keyboardEvent struct {
	key sdl.Keycode
	fn  func()
}

type EventPool struct {
	quitEvents     []func()
	keyboardEvents []keyboardEvent
}

func NewEventPool() *EventPool {
	return &EventPool{
		quitEvents:     make([]func(), 0),
		keyboardEvents: make([]keyboardEvent, 0),
	}
}

func (ep *EventPool) AddQuitEvent(fn func()) {
	ep.quitEvents = append(ep.quitEvents, fn)
}

func (ep *EventPool) AddKeyboardEvent(key sdl.Keycode, fn func()) {
	e := keyboardEvent{key, fn}
	ep.keyboardEvents = append(ep.keyboardEvents, e)
}

func (ep *EventPool) HandleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.KeyboardEvent:
			if e.Type == sdl.KEYDOWN {
				sym := e.Keysym.Sym
				for _, ke := range ep.keyboardEvents {
					if ke.key == sym {
						ke.fn()
					}
				}
			}
		case *sdl.QuitEvent:
			for _, qe := range ep.quitEvents {
				qe()
			}
		}
	}
}
