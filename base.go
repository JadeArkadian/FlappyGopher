package main

import (
	"fmt"

	"github.com/Zyko0/go-sdl3/img"
	"github.com/Zyko0/go-sdl3/sdl"
)

// Base represents a scrolling horizontal strip (floor or ceiling).
type Base struct {
	texture  *sdl.Texture
	x        float32
	y        float32
	rotation float64
}

// NewBase creates a Base at position (x, y) with the given rotation in degrees.
func NewBase(renderer *sdl.Renderer, x, y, rotation float32) (*Base, error) {
	texture, err := img.LoadTexture(renderer, BaseImgPath)
	if err != nil {
		return nil, fmt.Errorf("error loading base image: %w", err)
	}
	return &Base{texture: texture, x: x, y: y, rotation: float64(rotation)}, nil
}

// Update scrolls the base strip leftward and wraps it seamlessly.
func (base *Base) Update(deltaTime float32) {
	base.x -= PipesSpeed * (deltaTime * 60)
	if base.x <= -float32(WindowWidth) {
		base.x = 0
	}
}

// Draw renders the base strip (doubled for seamless scrolling).
func (base *Base) Draw(renderer *sdl.Renderer) {
	dst := sdl.FRect{X: base.x, Y: base.y, W: float32(WindowWidth), H: FloorHeight}
	renderer.RenderTextureRotated(base.texture, nil, &dst, base.rotation, nil, sdl.FLIP_NONE)
	dst2 := sdl.FRect{X: base.x + float32(WindowWidth), Y: base.y, W: float32(WindowWidth), H: FloorHeight}
	renderer.RenderTextureRotated(base.texture, nil, &dst2, base.rotation, nil, sdl.FLIP_NONE)
}

// Destroy releases the base texture.
func (base *Base) Destroy() {
	base.texture.Destroy()
}
