package core

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func loadTexture(renderer *sdl.Renderer, path string) (*sdl.Texture, error) {
	mapImage, err := img.Load(path)
	check(err)
	defer mapImage.Free()

	return renderer.CreateTextureFromSurface(mapImage)
}

func loadTextures(renderer *sdl.Renderer) map[string]*sdl.Texture {
	m := make(map[string]*sdl.Texture)

	texture, _ := loadTexture(renderer, "./assets/terrain.png")
	m["terrain"] = texture

	texture, _ = loadTexture(renderer, "./assets/characters.png")
	m["dwarves"] = texture

	return m
}
