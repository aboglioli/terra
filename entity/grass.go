package entity

import (
	"github.com/aboglioli/terra/core"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

type Grass struct {
	*core.Square
	texture *sdl.Texture
}

func NewGrass() *Grass {
	renderer := core.Renderer()

	texture, _ := renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, 32, 32)

	renderer.SetRenderTarget(texture)
	renderer.SetDrawColor(0, 128, 0, 255)
	renderer.Clear()

	renderer.SetDrawColor(0, 200, 0, 255)

	for i := 0; i < 10; i++ {
		x := rand.Int31n(32)
		y := rand.Int31n(32)
		renderer.FillRect(&sdl.Rect{x, y, 1, 1})
	}

	renderer.SetRenderTarget(nil)

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
