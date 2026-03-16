package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/Zyko0/go-sdl3/bin/binimg"
	"github.com/Zyko0/go-sdl3/bin/binsdl"
	"github.com/Zyko0/go-sdl3/bin/binttf"
	"github.com/Zyko0/go-sdl3/img"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/ttf"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

// GameState represents the current phase of the game.
type GameState int

const (
	StartScreen GameState = iota
	Playing
	GameOver
)

// Game holds all mutable state that was previously in global variables.
type Game struct {
	state     GameState
	deltaTime float32
	lastTime  uint64
}

// ColorWhite is a convenience color constant for white.
var ColorWhite = sdl.Color{R: 255, G: 255, B: 255, A: 255}

func chooseBackground() string {
	backgrounds := []string{
		BackgroundsDir + "/background-day.png",
		BackgroundsDir + "/background-night.png",
	}
	return backgrounds[rand.Intn(len(backgrounds))]
}

func chooseBird() string {
	birds := []string{RedBird, BlueBird, YellowBird}
	return birds[rand.Intn(len(birds))]
}

func (game *Game) run() error {
	defer binsdl.Load().Unload()
	defer binttf.Load().Unload()
	defer binimg.Load().Unload()
	defer sdl.Quit()
	defer ttf.Quit()
	defer sdl.CloseLibrary()
	defer ttf.CloseLibrary()
	defer img.CloseLibrary()

	// Initialize SDL
	if err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_EVENTS | sdl.INIT_GAMEPAD); err != nil {
		return fmt.Errorf("error initializing SDL: %w", err)
	}

	// Initialize speaker
	speaker.Init(beep.SampleRate(44100), 1024)

	// Create window and renderer
	window, renderer, err := sdl.CreateWindowAndRenderer(GameName, WindowWidth, WindowHeight, sdl.WINDOW_ALWAYS_ON_TOP)
	if err != nil {
		return fmt.Errorf("error creating window and renderer: %w", err)
	}
	defer renderer.Destroy()
	defer window.Destroy()

	// Initialize TTF module
	if err := ttf.Init(); err != nil {
		return fmt.Errorf("error initializing TTF module: %w", err)
	}

	// Choose random background and bird for this session
	background := chooseBackground()
	bird := chooseBird()

	// Initialize scenes
	titleScene, err := NewTitleScene(renderer, background)
	if err != nil {
		return fmt.Errorf("error creating title scene: %w", err)
	}
	defer titleScene.Destroy()

	scene, err := NewScene(renderer, background, bird)
	if err != nil {
		return fmt.Errorf("error creating game scene: %w", err)
	}
	defer scene.Destroy()

	game.lastTime = sdl.Ticks()

	sdl.RunLoop(func() error {
		currentTime := sdl.Ticks()
		if game.lastTime != 0 {
			game.deltaTime = float32(currentTime-game.lastTime) / 1000.0
		}
		if game.deltaTime == 0 {
			game.deltaTime = 1.0 / 60.0
		}
		game.lastTime = currentTime

		var event sdl.Event
		for sdl.PollEvent(&event) {
			switch event.Type {
			case sdl.EVENT_KEY_DOWN, sdl.EVENT_MOUSE_BUTTON_DOWN:
				if game.state == Playing {
					scene.fish.HandleInput(event)
					break
				}
				log.Println("Starting game")
				game.state = Playing
			case sdl.EVENT_QUIT:
				return sdl.EndLoop
			}
		}

		switch game.state {
		case StartScreen:
			titleScene.DrawScene(renderer)
		case Playing:
			scene.DrawScene(renderer)
			scene.UpdateScene(game.deltaTime)
		case GameOver:
		}

		return nil
	})

	return nil
}

func main() {
	game := &Game{}
	if err := game.run(); err != nil {
		log.Printf("Fatal error: %v", err)
		os.Exit(2)
	}
}
