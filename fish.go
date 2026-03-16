package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Zyko0/go-sdl3/img"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

// Fish represents the player-controlled bird character.
type Fish struct {
	textures  []*sdl.Texture
	x         float32
	y         float32
	speed     float32
	isDead    bool
	flapSound beep.StreamSeeker
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

	f, err := os.Open(FlapSfxPath)
	if err != nil {
		return nil, fmt.Errorf("error opening sound file: %w", err)
	}
	defer f.Close()

	s, format, err := wav.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("error decoding sound: %w", err)
	}
	streamer := beep.Streamer(s)

	// Resample if necessary
	if format.SampleRate != beep.SampleRate(44100) {
		streamer = beep.Resample(4, format.SampleRate, beep.SampleRate(44100), streamer)
	}

	// Create a buffer to make it seekable
	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)

	return &Fish{
		textures:  textures,
		x:         BirdInitialX,
		y:         BirdInitialY,
		flapSound: buffer.Streamer(0, buffer.Len()),
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

	// Play flap sound
	fish.flapSound.Seek(0)
	speaker.Play(fish.flapSound)
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
