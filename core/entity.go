package core

import "github.com/veandco/go-sdl2/sdl"

type Entity interface {
	Update(delta float64)
	Render(renderer *sdl.Renderer)
}

type BoundTexture struct {
	Bound   *sdl.Rect
	Texture *sdl.Texture
}

type TexturedRect struct {
	*sdl.Rect
	*BoundTexture
}

func (r *TexturedRect) Render(renderer *sdl.Renderer) {
	renderer.Copy(r.Texture, r.Bound, r.Rect)
}

func (r *TexturedRect) Update(d float64) {
}
