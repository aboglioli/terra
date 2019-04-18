package main

import (
	"github.com/aboglioli/terra/core"
	"github.com/aboglioli/terra/entity"
	"github.com/veandco/go-sdl2/img"
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
			x -= 32
		case sdl.K_d:
			x += 32
		case sdl.K_w:
			y -= 32
		case sdl.K_s:
			y += 32
		}
	})

	renderer := core.Renderer()

	var terrain [16][16]*entity.Grass

	for i := 0; i < len(terrain); i++ {
		for j := 0; j < len(terrain[0]); j++ {
			terrain[i][j] = entity.NewGrass()
		}
	}

	loop := core.NewLoop(60, func(d float64) {
		// Clear screen
		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Clear()

		// Draw terrain
		for i := (0); i < len(terrain); i++ {
			for j := 0; j < len(terrain[0]); j++ {
				grass := terrain[i][j]
				grassTexture := grass.Render(d)
				renderer.Copy(grassTexture, &sdl.Rect{0, 0, 32, 32}, &sdl.Rect{int32(i * 32), int32(j * 32), 32, 32})
			}
		}

		image, _ := img.Load("./assets/map.png")
		defer image.Free()
		texture, _ := renderer.CreateTextureFromSurface(image)
		renderer.Copy(texture, &sdl.Rect{96, 0, 32, 32}, &sdl.Rect{128, 128, 32, 32})

		// Draw character
		renderer.SetDrawColor(255, 0, 0, 255)
		renderer.FillRect(&sdl.Rect{x, y, 32, 32})

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
