package entity

import (
	"github.com/aboglioli/terra/core"
	"github.com/veandco/go-sdl2/sdl"
)

type Grass struct {
	texture *core.BoundTexture
}

func NewGrass() *Grass {
	texture, _ := core.Texture("terrain")

	return &Grass{
		texture: &core.BoundTexture{
			Bound:   &sdl.Rect{0 * 32, 6 * 32, core.Tile, core.Tile},
			Texture: texture,
		},
	}
}

func (g *Grass) Update(d float64) {
}

func (g *Grass) Render(d float64) *core.BoundTexture {
	return g.texture
}
