package main

import (
	"fmt"

	"github.com/Zyko0/go-sdl3/img"
	"github.com/Zyko0/go-sdl3/sdl"
)

type Base struct {
	baseTexture *sdl.Texture
	x           float32
	y           float32
	rotation    float64
}

func NewBase(renderer *sdl.Renderer, x, y, rotation float32) (*Base, error) {
	baseTexture, err := img.LoadTexture(renderer, BaseImgPath)

	if err != nil {
		return nil, fmt.Errorf("Error while loading base floor image  %v", err)
	}

	return &Base{baseTexture: baseTexture, x: x, y: y, rotation: float64(rotation)}, nil
}

func (base *Base) UpdateBase() {
	base.x -= PipesSpeed * (deltaTime * 60)
	if base.x <= -float32(WindowWidth) {
		base.x = 0
	}
}

func (base *Base) DrawBase(renderer *sdl.Renderer) {
	dst := sdl.FRect{X: base.x, Y: base.y, W: float32(WindowWidth), H: FloorHeight}
	renderer.RenderTextureRotated(base.baseTexture, nil, &dst, base.rotation, nil, sdl.FLIP_NONE)
	dst2 := sdl.FRect{X: base.x + float32(WindowWidth), Y: base.y, W: float32(WindowWidth), H: FloorHeight}
	renderer.RenderTextureRotated(base.baseTexture, nil, &dst2, base.rotation, nil, sdl.FLIP_NONE)
}

func (base *Base) Destroy() {
	base.baseTexture.Destroy()
}
