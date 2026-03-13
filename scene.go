package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Zyko0/go-sdl3/img"
	"github.com/Zyko0/go-sdl3/sdl"
)

type Scene struct {
	pipes []*Pipe
	backgroundTexture *sdl.Texture
	floor *Base
	ceiling *Base
	fish *Fish
	isGameOver bool
}

func NewScene(renderer *sdl.Renderer, backgroundPath string) (*Scene, error) {
	bgTexture, err := img.LoadTexture(renderer, backgroundPath)
	if err != nil {
		return nil, fmt.Errorf("Error while loading background image  %v", err)
	}

	fish, err := NewFish(renderer, "")
	if err != nil {
		return nil, fmt.Errorf("Error while creating fish  %v", err)
	}

	floor, err := NewBase(renderer, 0, float32(WindowHeight)-FloorHeight, 0)
	if err != nil {
		return nil, fmt.Errorf("Error while creating floor  %v", err)
	}

	pipes := []*Pipe{}
	for i := 0; i < 3; i++ {
		pipe, err := NewPipe(renderer, float32(WindowWidth)+float32(i*400))
		if err != nil {
			return nil, fmt.Errorf("Error while creating pipe  %v", err)
		}
		pipes = append(pipes, pipe)
	}

	ceiling, err := NewBase(renderer, 0, 0, 180)
	if err != nil {
		return nil, fmt.Errorf("Error while creating ceiling  %v", err)
	}

	return &Scene{backgroundTexture: bgTexture, fish: fish, floor: floor, ceiling: ceiling, pipes: pipes}, nil
}

func (scene *Scene) UpdateScene() {
	if scene.fish.isDead && !scene.isGameOver {
		scene.isGameOver = true	
		log.Println("Game Over!")

		go func() {
			time.Sleep(3 * time.Second)
			log.Println("Reset scene")
			scene.ResetScene()
			scene.isGameOver = false
		}()
	} else if !scene.isGameOver{
		scene.fish.UpdateFish()
		scene.floor.UpdateBase()
		scene.ceiling.UpdateBase()
		for _, pipe := range scene.pipes {
			pipe.UpdatePipe()
			if scene.CheckCollisions(scene.fish.ColliderBounds(), pipe.ColliderBounds()) {
				log.Println("Collision detected!")
				scene.fish.isDead = true
			}
		}
	}
}

func (scene *Scene) CheckCollisions(fishCollider sdl.FRect, pipeCollider sdl.FRect) bool {
	    return fishCollider.X < pipeCollider.X+pipeCollider.W &&
        fishCollider.X+fishCollider.W > pipeCollider.X &&
        fishCollider.Y < pipeCollider.Y+pipeCollider.H &&
        fishCollider.Y+fishCollider.H > pipeCollider.Y
}

func (scene *Scene) ResetScene() {
	scene.fish.ResetFish()
}

func (scene *Scene) DrawScene(renderer *sdl.Renderer) {
	renderer.Clear()
	scene.DrawBackground(renderer)
	for _, pipe := range scene.pipes {
		pipe.DrawPipe(renderer)
	}
	scene.DrawCeiling(renderer)
	scene.DrawFloor(renderer)
	scene.fish.DrawFish(renderer)
	renderer.Present()
}

func (scene *Scene) DrawBackground(renderer *sdl.Renderer) {
	dst := sdl.FRect{X: 0, Y: 0, W: float32(WindowWidth), H: float32(WindowHeight)}
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.RenderTexture(scene.backgroundTexture, nil, &dst)
}

func (scene *Scene) DrawFloor(renderer *sdl.Renderer) {
	scene.floor.DrawBase(renderer)
}

func (scene *Scene) DrawCeiling(renderer *sdl.Renderer) {
	scene.ceiling.DrawBase(renderer)
}

func (scene *Scene) Destroy() {
	scene.backgroundTexture.Destroy()
	scene.fish.Destroy()
	scene.floor.Destroy()
	scene.ceiling.Destroy()
}
