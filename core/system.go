package core

import "github.com/veandco/go-sdl2/sdl"

type system struct {
	window  *sdl.Window
	surface *sdl.Surface
}

var instance *system

func InitSystems(title string, width, height int32) {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	check(err)

	// Create window
	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	check(err)

	// Get main surface
	surface, err := window.GetSurface()
	check(err)

	instance = &system{
		window:  window,
		surface: surface,
	}
}

func DestroySystems() {
	instance.window.Destroy()
	sdl.Quit()
}

func Window() *sdl.Window {
	return instance.window
}

func Surface() *sdl.Surface {
	return instance.surface
}

func Update() {
	instance.window.UpdateSurface()
}
