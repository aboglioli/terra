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

	var x, y int32 = 0, 0

	events.AddKeyboardEvent(func(e *sdl.KeyboardEvent) {
		switch e.Keysym.Sym {
		case sdl.K_a:
			x -= core.Tile
		case sdl.K_d:
			x += core.Tile
		case sdl.K_w:
			y -= core.Tile
		case sdl.K_s:
			y += core.Tile
		}
	})

	renderer := core.Renderer()

	var grasses [60][60]core.Entity
	for i := 0; i < 60; i++ {
		for j := 0; j < 60; j++ {
			grasses[i][j] = entity.NewGrass()
		}
	}

	dwarf := entity.NewDwarf()

	loop := core.NewLoop(60, func(d float64) {
		// Clear screen
		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Clear()

		// Draw terrain
		for i := 0; i < 60; i++ {
			for j := 0; j < 60; j++ {
				grass := grasses[i][j]
				texture := grass.Render(d)
				renderer.Copy(texture.Texture, texture.Bound, &sdl.Rect{int32(i * core.Tile), int32(j * core.Tile), core.Tile, core.Tile})
			}
		}

		// Draw character
		texture := dwarf.Render(d)
		renderer.Copy(texture.Texture, texture.Bound, &sdl.Rect{x, y, core.Tile, core.Tile})

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
