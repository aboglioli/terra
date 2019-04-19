package core

import "github.com/veandco/go-sdl2/sdl"

type BoundTexture struct {
	Bound   *sdl.Rect
	Texture *sdl.Texture
}

type Entity interface {
	Update(delta float64)
	Render(delta float64) *BoundTexture
}

type Square sdl.Rect

func (s *Square) Render(d float64) *BoundTexture {
	return nil
}

func (s *Square) Update(d float64) {
}
