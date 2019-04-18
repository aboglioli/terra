package core

import "github.com/veandco/go-sdl2/sdl"

type Entity interface {
	Update(delta float64) error
	Render(delta float64) error
}

type Square struct {
	X      int
	Y      int
	Width  int
	Height int
}

func (s *Square) Render(d float64) *sdl.Texture {
	return nil
}

func (s *Square) Update(d float64) error {
	return nil
}
