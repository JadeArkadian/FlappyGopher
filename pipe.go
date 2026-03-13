package main

import (
	"fmt"
	"math/rand"

	"github.com/Zyko0/go-sdl3/img"
	"github.com/Zyko0/go-sdl3/sdl"
)

type Pipe struct {
	pipeTexture *sdl.Texture
	pipeX       float32
	pipeY       float32
	rotation    float64
}

func NewPipe(renderer *sdl.Renderer, x float32) (*Pipe, error) {
	var pipeTextures *sdl.Texture

	pipeTexture, err := img.LoadTexture(renderer, PipeGreenPath)
	if err != nil {
		return nil, fmt.Errorf("Error while loading pipe image  %v", err)
	}
	pipeTextures = pipeTexture

	pipe := &Pipe{pipeTexture: pipeTextures, pipeX: x}
	pipe.ResetPipe()

	return pipe, nil
}

func (pipe *Pipe) ColliderBounds() sdl.FRect {
	return sdl.FRect{X: pipe.pipeX, Y: pipe.pipeY, W: PipeWidth, H: PipeHeight}
}

func (pipe *Pipe) ResetPipe() error {
	// Randomly choose between 180 and 0 degrees for the pipe rotation
	pipe.rotation = 180 * float64(rand.Intn(2))

	// Randomly choose a height for the pipe between 150 and 350 pixels from the top if it's a ceiling pipe, or from the bottom if it's a floor pipe
	if pipe.rotation == 180 {
		// Ceiling pipe
		pipe.pipeY = choosePipeHeight()*100.0 - 150.0
	} else {
		pipe.pipeY = float32(WindowHeight) - FloorHeight - choosePipeHeight()*100.0 - 150.0
	}
	return nil
}

func (pipe *Pipe) UpdatePipe() {
	pipe.pipeX -= PipesSpeed * (deltaTime * 60)
	if pipe.pipeX <= -float32(WindowWidth) {
		pipe.pipeX = float32(WindowWidth)
		pipe.ResetPipe()
	}
}

func (pipe *Pipe) DrawPipe(renderer *sdl.Renderer) {
	dst := sdl.FRect{X: pipe.pipeX, Y: pipe.pipeY, W: float32(PipeWidth), H: float32(PipeHeight)}
	renderer.RenderTextureRotated(pipe.pipeTexture, nil, &dst, pipe.rotation, nil, sdl.FLIP_NONE)
}

func (pipe *Pipe) Destroy() {
	pipe.pipeTexture.Destroy()
}

func choosePipeHeight() float32 {
	return float32(rand.Intn(3))
}