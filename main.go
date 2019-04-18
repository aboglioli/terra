package main

import (
	"fmt"
	"github.com/aboglioli/terra/core"
	"github.com/aboglioli/terra/entity"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	core.InitEngine("Terra", 1024, 768)
	defer core.DestroyEngine()

	events := core.Events()

	var x, y int32 = 0, 0

	events.AddKeyboardEvent(func(e *sdl.KeyboardEvent) {
		switch e.Keysym.Sym {
		case sdl.K_a:
			x -= 64
		case sdl.K_d:
			x += 64
		case sdl.K_w:
			y -= 64
		case sdl.K_s:
			y += 64
		}
	})

	renderer := core.Renderer()

	var terrain [64][64]*entity.Grass

	fmt.Println("Making terrain")
	for i := 0; i < 64; i++ {
		for j := 0; j < 64; j++ {
			terrain[i][j] = entity.NewGrass()
		}
	}
	fmt.Println("Ready")

	loop := core.NewLoop(60, func(d float64) {
		// Clear screen
		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Clear()

		// Draw terrain
		for i := (0); i < 64; i++ {
			for j := 0; j < 64; j++ {
				grass := terrain[i][j]
				grassTexture := grass.Render(d)
				renderer.Copy(grassTexture, &sdl.Rect{0, 0, 32, 32}, &sdl.Rect{int32(i * 32), int32(j * 32), 32, 32})
			}
		}

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
