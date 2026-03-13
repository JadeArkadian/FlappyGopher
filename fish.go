package main

import (
	"fmt"
	"log"

	"github.com/Zyko0/go-sdl3/img"
	"github.com/Zyko0/go-sdl3/sdl"
)

type Fish struct {
	fishTexture []*sdl.Texture
	fishX       float32
	fishY       float32
	fishSpeed   float32
	isDead      bool
}

func NewFish(renderer *sdl.Renderer, backgroundPath string) (*Fish, error) {
	var fishTextures []*sdl.Texture
	for i := 0; i < 3; i++ {
		fishTexture, err := img.LoadTexture(renderer, chooseBird()+"/"+fmt.Sprintf("bird-%02d.png", i))
		if err != nil {
			return nil, fmt.Errorf("Error while loading fish image  %v", err)
		}
		fishTextures = append(fishTextures, fishTexture)
	}

	return &Fish{fishTexture: fishTextures, fishX: 100, fishY: 200, isDead: false}, nil
}

func (fish *Fish) ColliderBounds() sdl.FRect {
	return sdl.FRect{X: fish.fishX, Y: fish.fishY, W: BirdWidth, H: BirdHeight}
}

func (fish *Fish) Flap() {
	if fish.isDead {
		return
	}
	log.Println("Flap!")
	fish.fishSpeed = BirdFlapStrength
}

func (fish *Fish) HandleInput(event sdl.Event) {
	if event.Type == sdl.EVENT_MOUSE_BUTTON_DOWN {
		fish.Flap()
	}
}

func (fish *Fish) ResetFish() {
	fish.fishX = 100
	fish.fishY = 200
	fish.fishSpeed = 0
	fish.isDead = false
}

func (fish *Fish) UpdateFish() {
	if fish.isDead {
		return
	}

	fish.fishY += fish.fishSpeed
	fish.fishSpeed += BirdGravity
	if fish.fishSpeed > BirdMaxFallSpeed {
		fish.fishSpeed = BirdMaxFallSpeed
	}

	// Check if the fish hits the ground
	if fish.fishY > float32(WindowHeight)- FloorHeight - BirdHeight {
		fish.fishY = float32(WindowHeight) - FloorHeight - BirdHeight
		fish.isDead = true
	}

	// Check if the fish hits the ceiling
	if fish.fishY < FloorHeight {
		fish.fishY = FloorHeight
		fish.isDead = true
	}
}

func (fish *Fish) DrawFish(renderer *sdl.Renderer) {
	if fish.isDead {
		return
	}
	frame := int(sdl.Ticks()/100) % len(fish.fishTexture)
	dst := sdl.FRect{X: fish.fishX, Y: fish.fishY, W: BirdWidth, H: BirdHeight}
	renderer.RenderTexture(fish.fishTexture[frame], nil, &dst)
}

func (fish *Fish) Destroy() {
	for i := 0; i < len(fish.fishTexture); i++ {
		defer fish.fishTexture[i].Destroy()
	}
}
