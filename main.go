package main

import (
	"github.com/aboglioli/terra/core"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	title  = "Terra"
	width  = 1024
	height = 768
)

func main() {
	core.InitEngine(title, width, height)
	defer core.DestroyEngine()

	events := core.Events()

	var x, y int32 = 0, 0

	events.AddKeyboardEvent(func(e *sdl.KeyboardEvent) {
		switch e.Keysym.Sym {
		case sdl.K_a:
			x -= 16
		case sdl.K_d:
			x += 16
		case sdl.K_w:
			y -= 16
		case sdl.K_s:
			y += 16
		}
	})

	renderer := core.Renderer()

	media := core.Media()

	loop := core.NewLoop(60, func(d float64) {
		// Clear screen
		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Clear()

		// Draw terrain
		for i := (0); i < 64; i++ {
			for j := 0; j < 64; j++ {
				grass := media["grass1"]
				renderer.Copy(grass.Texture, grass.Bound, &sdl.Rect{int32(i * 16), int32(j * 16), 16, 16})
			}
		}

		totem := media["totem"]
		renderer.Copy(totem.Texture, totem.Bound, &sdl.Rect{128, 128, 16, 16})

		// Draw character
		renderer.SetDrawColor(255, 0, 0, 255)
		renderer.FillRect(&sdl.Rect{x, y, 16, 16})

		// Update
		renderer.Present()

		// Handle events
		events.HandleEvents()
	})

	events.AddKeyboardEvent(func(e *sdl.KeyboardEvent) {
		if e.Keysym.Sym == sdl.K_ESCAPE {
			loop.Stop()
		}
	})

	events.AddQuitEvent(func() {
		loop.Stop()
	})

	loop.Start()
}
