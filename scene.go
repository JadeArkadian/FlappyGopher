package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Zyko0/go-sdl3/img"
	"github.com/Zyko0/go-sdl3/sdl"
)

// Scene holds all game objects for the main gameplay phase.
type Scene struct {
	mu                sync.Mutex
	pipes             []*Pipe
	backgroundTexture *sdl.Texture
	floor             *Base
	ceiling           *Base
	fish              *Fish
	isGameOver        bool
}

// NewScene creates the gameplay scene with background, fish, floor, ceiling and pipes.
func NewScene(renderer *sdl.Renderer, backgroundPath, birdPath string) (*Scene, error) {
	bgTexture, err := img.LoadTexture(renderer, backgroundPath)
	if err != nil {
		return nil, fmt.Errorf("error loading background image: %w", err)
	}

	fish, err := NewFish(renderer, birdPath)
	if err != nil {
		return nil, fmt.Errorf("error creating fish: %w", err)
	}

	floor, err := NewBase(renderer, 0, float32(WindowHeight)-FloorHeight, 0)
	if err != nil {
		return nil, fmt.Errorf("error creating floor: %w", err)
	}

	ceiling, err := NewBase(renderer, 0, 0, 180)
	if err != nil {
		return nil, fmt.Errorf("error creating ceiling: %w", err)
	}

	var pipes []*Pipe
	for i := 0; i < 3; i++ {
		pipe, err := NewPipe(renderer, float32(WindowWidth)+float32(i*400))
		if err != nil {
			return nil, fmt.Errorf("error creating pipe %d: %w", i, err)
		}
		pipes = append(pipes, pipe)
	}

	return &Scene{
		backgroundTexture: bgTexture,
		fish:              fish,
		floor:             floor,
		ceiling:           ceiling,
		pipes:             pipes,
	}, nil
}

// UpdateScene advances the game simulation by one frame.
func (scene *Scene) UpdateScene(deltaTime float32) {
	scene.mu.Lock()
	gameOver := scene.isGameOver
	scene.mu.Unlock()

	if scene.fish.isDead && !gameOver {
		scene.mu.Lock()
		scene.isGameOver = true
		scene.mu.Unlock()

		log.Println("Game Over!")
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("Resetting scene")
			scene.reset()
			scene.mu.Lock()
			scene.isGameOver = false
			scene.mu.Unlock()
		}()
		return
	}

	if gameOver {
		return
	}

	scene.fish.Update(deltaTime)
	scene.floor.Update(deltaTime)
	scene.ceiling.Update(deltaTime)
	for _, pipe := range scene.pipes {
		pipe.Update(deltaTime)
		if scene.checkCollision(scene.fish.ColliderBounds(), pipe.ColliderBounds()) {
			log.Println("Collision detected!")
			scene.fish.isDead = true
		}
	}
}

// checkCollision returns true if two rectangles overlap (AABB).
func (scene *Scene) checkCollision(a, b sdl.FRect) bool {
	return a.X < b.X+b.W &&
		a.X+a.W > b.X &&
		a.Y < b.Y+b.H &&
		a.Y+a.H > b.Y
}

// reset restores the scene to its initial state.
func (scene *Scene) reset() {
	scene.fish.Reset()
}

// DrawScene renders the complete scene for the current frame.
func (scene *Scene) DrawScene(renderer *sdl.Renderer) {
	renderer.Clear()
	renderer.SetDrawColor(255, 255, 255, 255)

	dst := sdl.FRect{X: 0, Y: 0, W: float32(WindowWidth), H: float32(WindowHeight)}
	renderer.RenderTexture(scene.backgroundTexture, nil, &dst)

	for _, pipe := range scene.pipes {
		pipe.Draw(renderer)
	}
	scene.ceiling.Draw(renderer)
	scene.floor.Draw(renderer)
	scene.fish.Draw(renderer)

	renderer.Present()
}

// Destroy releases all resources held by the scene.
func (scene *Scene) Destroy() {
	scene.backgroundTexture.Destroy()
	scene.fish.Destroy()
	scene.floor.Destroy()
	scene.ceiling.Destroy()
	for _, pipe := range scene.pipes {
		pipe.Destroy()
	}
}
