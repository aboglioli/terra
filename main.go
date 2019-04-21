package main

import (
	"github.com/aboglioli/terra/core"
	"github.com/aboglioli/terra/entity"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	core.InitEngine(core.Title, core.Width, core.Height)
	defer core.DestroyEngine()

	events := core.Events()

	dwarf := entity.NewDwarf()

	renderer := core.Renderer()

	var grasses [60][60]core.Entity
	for i := 0; i < 60; i++ {
		for j := 0; j < 60; j++ {
			grasses[i][j] = entity.NewGrass(i*core.Tile, j*core.Tile)
		}
	}

	loop := core.NewLoop(60, func(d float64) {
		// Clear screen
		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Clear()

		// Draw terrain
		for i := 0; i < 60; i++ {
			for j := 0; j < 60; j++ {
				grass := grasses[i][j]
				grass.Render(renderer)
			}
		}

		// Draw character
		dwarf.Render(renderer)

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
