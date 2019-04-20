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

	return &Dwarf{
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
}
