package entity

import (
	"github.com/aboglioli/terra/core"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

type Grass struct {
	texture *core.BoundTexture
}

func NewGrass() *Grass {
	texture, _ := core.Texture("terrain")

	x := rand.Int31n(3) + 9
	y := rand.Int31n(3) + 0

	return &Grass{
		texture: &core.BoundTexture{
			Bound:   &sdl.Rect{x * core.Tile, y * core.Tile, core.Tile, core.Tile},
			Texture: texture,
		},
	}
}

func (g *Grass) Update(d float64) {
}

func (g *Grass) Render(d float64) *core.BoundTexture {
	return g.texture
}
