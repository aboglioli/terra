package core

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type engine struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	events   *EventPool
}

var instance *engine

func InitEngine(title string, width, height int32) {
	// Initialize SDL and dependencies
	err := sdl.Init(sdl.INIT_EVERYTHING)
	check(err)

	img.Init(img.INIT_PNG)

	// Create window
	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	check(err)

	// Get main surface
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	check(err)

	// Create event pool
	events := NewEventPool()

	instance = &engine{
		window:   window,
		renderer: renderer,
		events:   events,
	}
}

func DestroyEngine() {
	instance.renderer.Destroy()
	instance.window.Destroy()
	img.Quit()
	sdl.Quit()
}

func Window() *sdl.Window {
	return instance.window
}

func Renderer() *sdl.Renderer {
	return instance.renderer
}

func Events() *EventPool {
	return instance.events
}
