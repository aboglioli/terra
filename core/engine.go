package core

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type engine struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	events   *EventPool
	textures map[string]*sdl.Texture
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

	// Load media
	textures := loadTextures(renderer)

	instance = &engine{
		window:   window,
		renderer: renderer,
		events:   events,
		textures: textures,
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

func Texture(key string) (*sdl.Texture, bool) {
	texture, ok := instance.textures[key]
	return texture, ok
}
