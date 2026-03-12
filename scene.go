package main

import (
	"fmt"

	"github.com/Zyko0/go-sdl3/img"
	"github.com/Zyko0/go-sdl3/sdl"
)

type Scene struct {
	background *sdl.Texture
}

func NewScene(renderer *sdl.Renderer, backgroundPath string) (*Scene, error) {
	texture, err := img.LoadTexture(renderer, backgroundPath)
	if err != nil {
		return nil, fmt.Errorf("Error while loading background image  %v", err)
	}

	return &Scene{background: texture}, nil
}

func (scene *Scene) DrawScene(renderer *sdl.Renderer) {
	renderer.Clear()
	dst := sdl.FRect{X: 0, Y: 0, W: float32(WindowWidth), H: float32(WindowHeight)}
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.RenderTexture(scene.background, nil, &dst)
	renderer.Present()
}

func (scene *Scene) Destroy() {
	defer scene.background.Destroy()
}
