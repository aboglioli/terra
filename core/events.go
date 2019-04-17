package core

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type EventPool struct {
	quitEvents     []func()
	keyboardEvents map[sdl.Keycode][]func()
}

func NewEventPool() *EventPool {
	return &EventPool{
		quitEvents:     make([]func(), 0),
		keyboardEvents: make(map[sdl.Keycode][]func()),
	}
}

func (ep *EventPool) AddQuitEvent(fn func()) {
	ep.quitEvents = append(ep.quitEvents, fn)
}

func (ep *EventPool) AddKeyboardEvent(key sdl.Keycode, fn func()) {
	ke := ep.keyboardEvents

	if ke[key] == nil {
		ke[key] = make([]func(), 0)
	}

	ke[key] = append(ke[key], fn)
}

func (ep *EventPool) HandleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.KeyboardEvent:
			if e.Type == sdl.KEYDOWN {
				key := e.Keysym.Sym

				if eventsByKey, ok := ep.keyboardEvents[key]; ok {
					for _, ke := range eventsByKey {
						ke()
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
