package entity

import (
	"github.com/aboglioli/terra/core"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Grass struct {
	*core.Square
	texture *sdl.Texture
}

func NewGrass() *Grass {
	renderer := core.Renderer()

	image, _ := img.Load("./assets/map.png")
	defer image.Free()
	texture, _ := renderer.CreateTextureFromSurface(image)

	return &Grass{
		Square: &core.Square{
			X:      0,
			Y:      0,
			Width:  32,
			Height: 32,
		},
		texture: texture,
	}
}

func (g *Grass) Update(d float64) error {
	return nil
}

func (g *Grass) Render(d float64) *sdl.Texture {
	return g.texture
}
