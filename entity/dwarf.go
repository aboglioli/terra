package entity

import (
	"github.com/aboglioli/terra/core"
	"github.com/veandco/go-sdl2/sdl"
)

type Dwarf struct {
	texture *core.BoundTexture
}

func NewDwarf() *Dwarf {
	texture, _ := core.Texture("dwarves")

	return &Dwarf{
		texture: &core.BoundTexture{
			Bound:   &sdl.Rect{64, 96, core.Tile, core.Tile},
			Texture: texture,
		},
	}
}

func (g *Dwarf) Update(d float64) {
}

func (g *Dwarf) Render(d float64) *core.BoundTexture {
	return g.texture
}
