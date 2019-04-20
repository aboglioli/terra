package entity

import (
	"github.com/aboglioli/terra/core"
	"github.com/veandco/go-sdl2/sdl"
)

type Grass struct {
	*core.TexturedRect
}

func NewGrass(x, y int) *Grass {
	texture, _ := core.Texture("terrain")

	return &Grass{
		&core.TexturedRect{
			&sdl.Rect{
				X: int32(x),
				Y: int32(y),
				W: core.Tile,
				H: core.Tile,
			},
			&core.BoundTexture{
				Bound:   &sdl.Rect{0 * 32, 5 * 32, core.Tile, core.Tile},
				Texture: texture,
			},
		},
	}
}
