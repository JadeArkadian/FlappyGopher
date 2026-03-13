package main

import (
	"fmt"
	"math/rand"

	"github.com/Zyko0/go-sdl3/img"
	"github.com/Zyko0/go-sdl3/sdl"
)

// Pipe represents a single pipe obstacle.
type Pipe struct {
	texture  *sdl.Texture
	x        float32
	y        float32
	rotation float64
}

// NewPipe creates a Pipe at horizontal position x.
func NewPipe(renderer *sdl.Renderer, x float32) (*Pipe, error) {
	texture, err := img.LoadTexture(renderer, PipeGreenPath)
	if err != nil {
		return nil, fmt.Errorf("error loading pipe image: %w", err)
	}

	p := &Pipe{texture: texture, x: x}
	p.reset()
	return p, nil
}

// ColliderBounds returns the axis-aligned bounding box of the pipe.
func (pipe *Pipe) ColliderBounds() sdl.FRect {
	return sdl.FRect{X: pipe.x, Y: pipe.y, W: PipeWidth, H: PipeHeight}
}

// reset randomizes the pipe's orientation and vertical position.
func (pipe *Pipe) reset() {
	// Randomly place pipe at top (180°) or bottom (0°)
	pipe.rotation = 180 * float64(rand.Intn(2))

	height := randomPipeOffset()
	if pipe.rotation == 180 {
		// Ceiling pipe: offset from the top
		pipe.y = height*100.0 - 150.0
	} else {
		// Floor pipe: offset from the bottom
		pipe.y = float32(WindowHeight) - FloorHeight - height*100.0 - 150.0
	}
}

// Update moves the pipe leftward and recycles it when off-screen.
func (pipe *Pipe) Update(deltaTime float32) {
	pipe.x -= PipesSpeed * (deltaTime * 60)
	if pipe.x <= -float32(WindowWidth) {
		pipe.x = float32(WindowWidth)
		pipe.reset()
	}
}

// Draw renders the pipe with its current rotation.
func (pipe *Pipe) Draw(renderer *sdl.Renderer) {
	dst := sdl.FRect{X: pipe.x, Y: pipe.y, W: PipeWidth, H: PipeHeight}
	renderer.RenderTextureRotated(pipe.texture, nil, &dst, pipe.rotation, nil, sdl.FLIP_NONE)
}

// Destroy releases the pipe texture.
func (pipe *Pipe) Destroy() {
	pipe.texture.Destroy()
}

// randomPipeOffset returns a random float32 in [0, 2] to vary pipe heights.
func randomPipeOffset() float32 {
	return float32(rand.Intn(3))
}
