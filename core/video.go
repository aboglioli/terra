package core

import "github.com/veandco/go-sdl2/sdl"

type Video struct {
	window  *sdl.Window
	Surface *sdl.Surface
}

func InitSDL() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	check(err)
}

func DestroySDL() {
	sdl.Quit()
}

func InitVideo(title string, width, height int32) *Video {
	// Create window
	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	check(err)

	// Get main surface
	surface, err := window.GetSurface()
	check(err)

	return &Video{
		window:  window,
		Surface: surface,
	}
}

func (v *Video) Update() {
	v.window.UpdateSurface()
}

func (v *Video) Destroy() {
	v.window.Destroy()
}
