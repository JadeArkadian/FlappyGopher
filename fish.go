package main

import (
	"fmt"
	"log"

	"github.com/Zyko0/go-sdl3/img"
	"github.com/Zyko0/go-sdl3/sdl"
)

// Fish represents the player-controlled bird character.
type Fish struct {
	textures []*sdl.Texture
	x        float32
	y        float32
	speed    float32
	isDead   bool
}

// NewFish creates a Fish loading sprites from birdPath.
func NewFish(renderer *sdl.Renderer, birdPath string) (*Fish, error) {
	var textures []*sdl.Texture
	for i := 0; i < 3; i++ {
		texture, err := img.LoadTexture(renderer, fmt.Sprintf("%s/bird-%02d.png", birdPath, i))
		if err != nil {
			return nil, fmt.Errorf("error loading bird sprite %d: %w", i, err)
		}
		textures = append(textures, texture)
	}

	return &Fish{
		textures: textures,
		x:        BirdInitialX,
		y:        BirdInitialY,
	}, nil
}

// ColliderBounds returns the axis-aligned bounding box of the fish.
func (fish *Fish) ColliderBounds() sdl.FRect {
	return sdl.FRect{X: fish.x, Y: fish.y, W: BirdWidth, H: BirdHeight}
}

// Flap makes the fish jump upward.
func (fish *Fish) Flap() {
	if fish.isDead {
		return
	}
	log.Println("Flap!")
	fish.speed = BirdFlapStrength
}

// HandleInput processes SDL input events for the fish.
func (fish *Fish) HandleInput(event sdl.Event) {
	switch event.Type {
	case sdl.EVENT_MOUSE_BUTTON_DOWN:
		fish.Flap()
	case sdl.EVENT_KEY_DOWN:
		key := event.KeyboardEvent()
		if key.Scancode == sdl.SCANCODE_SPACE {
			fish.Flap()
		}
	}
}

// Reset restores the fish to its initial state.
func (fish *Fish) Reset() {
	fish.x = BirdInitialX
	fish.y = BirdInitialY
	fish.speed = 0
	fish.isDead = false
}

// Update applies gravity and movement to the fish.
func (fish *Fish) Update(deltaTime float32) {
	if fish.isDead {
		return
	}

	fish.y += fish.speed * (deltaTime * 60)
	fish.speed += BirdGravity * (deltaTime * 60)
	if fish.speed > BirdMaxFallSpeed {
		fish.speed = BirdMaxFallSpeed
	}

	// Floor collision
	if fish.y > float32(WindowHeight)-FloorHeight-BirdHeight {
		fish.y = float32(WindowHeight) - FloorHeight - BirdHeight
		fish.isDead = true
	}

	// Ceiling collision
	if fish.y < FloorHeight {
		fish.y = FloorHeight
		fish.isDead = true
	}
}

// Draw renders the current animation frame of the fish.
func (fish *Fish) Draw(renderer *sdl.Renderer) {
	if fish.isDead {
		return
	}
	frame := int(sdl.Ticks()/100) % len(fish.textures)
	dst := sdl.FRect{X: fish.x, Y: fish.y, W: BirdWidth, H: BirdHeight}
	renderer.RenderTexture(fish.textures[frame], nil, &dst)
}

// Destroy releases all loaded textures.
func (fish *Fish) Destroy() {
	for _, t := range fish.textures {
		t.Destroy()
	}
}
