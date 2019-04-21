package entity

import (
	"github.com/aboglioli/terra/core"
	"github.com/veandco/go-sdl2/sdl"
)

type Dwarf struct {
	*core.TexturedRect
}

func NewDwarf() *Dwarf {
	texture, _ := core.Texture("dwarves")

	dwarf := &Dwarf{
		&core.TexturedRect{
			&sdl.Rect{
				X: 0,
				Y: 0,
				W: 32,
				H: 32,
			},
			&core.BoundTexture{
				Bound:   &sdl.Rect{32, 0, core.Tile, core.Tile},
				Texture: texture,
			},
		},
	}

	events := core.Events()

	events.AddKeyboardEvent(func(e *sdl.KeyboardEvent) {
		switch e.Keysym.Sym {
		case sdl.K_a:
			dwarf.X -= core.Tile
		case sdl.K_d:
			dwarf.X += core.Tile
		case sdl.K_w:
			dwarf.Y -= core.Tile
		case sdl.K_s:
			dwarf.Y += core.Tile
		}
	})

	return dwarf
}
