package main

import (
	"fmt"

	"github.com/Zyko0/go-sdl3/img"
	"github.com/Zyko0/go-sdl3/sdl"
)

type Pipe struct {
	pipeTexture *sdl.Texture
	pipeX       float32
	pipeY       float32
	rotation    float64
}

func NewPipe(renderer *sdl.Renderer, rotation float64) (*Pipe, error) {
	var pipeTextures *sdl.Texture

	pipeTexture, err := img.LoadTexture(renderer, PipeGreenPath)
	if err != nil {
		return nil, fmt.Errorf("Error while loading pipe image  %v", err)
	}
	pipeTextures = pipeTexture

	return &Pipe{pipeTexture: pipeTextures, pipeX: 100, pipeY: 200, rotation: rotation}, nil
}

func (pipe *Pipe) UpdatePipe() {
	pipe.pipeX -= PipesSpeed
	if pipe.pipeX <= -float32(WindowWidth) {
		pipe.pipeX = float32(WindowWidth)
	}
}

func (pipe *Pipe) DrawPipe(renderer *sdl.Renderer) {
	dst := sdl.FRect{X: pipe.pipeX, Y: pipe.pipeY, W: float32(WindowWidth), H: FloorHeight}
	renderer.RenderTextureRotated(pipe.pipeTexture, nil, &dst, pipe.rotation, nil, sdl.FLIP_NONE)
}

func (pipe *Pipe) Destroy() {

}
