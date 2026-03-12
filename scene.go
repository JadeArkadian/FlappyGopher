package main

import (
	"fmt"

	"github.com/Zyko0/go-sdl3/img"
	"github.com/Zyko0/go-sdl3/sdl"
)

type Scene struct {
	backgroundTexture *sdl.Texture
	fishTexture[] *sdl.Texture
	fishX float32
	fishY float32
}

func NewScene(renderer *sdl.Renderer, backgroundPath string) (*Scene, error) {
	bgTexture, err := img.LoadTexture(renderer, backgroundPath)
	if err != nil {
		return nil, fmt.Errorf("Error while loading background image  %v", err)
	}

	var fishTextures[] *sdl.Texture
	for i := 0; i < 3; i++ {
		fishTexture, err := img.LoadTexture(renderer, chooseBird() + "/" + fmt.Sprintf("bird-%02d.png", i))
		if err != nil {
			return nil, fmt.Errorf("Error while loading fish image  %v", err)
		}
		fishTextures = append(fishTextures, fishTexture)
	}

	return &Scene{backgroundTexture: bgTexture, fishTexture: fishTextures, fishX: 100, fishY: 100}, nil
}

func (scene *Scene) DrawScene(renderer *sdl.Renderer) {
	renderer.Clear()
	scene.DrawBackground(renderer)
	scene.DrawFish(renderer, scene.fishX, scene.fishY)
	renderer.Present()
}

func (scene *Scene) DrawFish(renderer *sdl.Renderer, x, y float32) {
	frame := int(sdl.Ticks() / 100) % len(scene.fishTexture)

	dst := sdl.FRect{X: x, Y: y, W: BirdWidth, H: BirdHeight}
	renderer.RenderTexture(scene.fishTexture[frame], nil, &dst)
}

func (scene *Scene) DrawBackground(renderer *sdl.Renderer) {
	dst := sdl.FRect{X: 0, Y: 0, W: float32(WindowWidth), H: float32(WindowHeight)}
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.RenderTexture(scene.backgroundTexture, nil, &dst)
}

func (scene *Scene) Destroy() {
	defer scene.backgroundTexture.Destroy()
	for i := 0; i < len(scene.fishTexture); i++ {
		defer scene.fishTexture[i].Destroy()
	}
}
