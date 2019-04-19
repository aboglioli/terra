package core

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type BoundTexture struct {
	Bound   *sdl.Rect
	Texture *sdl.Texture
}

func LoadMedia(renderer *sdl.Renderer) map[string]*BoundTexture {
	m := make(map[string]*BoundTexture)

	mapImage, err := img.Load("./assets/map.png")
	check(err)
	defer mapImage.Free()

	texture, _ := renderer.CreateTextureFromSurface(mapImage)

	m["grass1"] = &BoundTexture{
		Bound:   &sdl.Rect{0, 0, 16, 16},
		Texture: texture,
	}

	m["totem"] = &BoundTexture{
		Bound:   &sdl.Rect{96, 0, 16, 16},
		Texture: texture,
	}

	return m
}
