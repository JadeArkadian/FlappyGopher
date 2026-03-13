package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/Zyko0/go-sdl3/bin/binimg"
	"github.com/Zyko0/go-sdl3/bin/binsdl"
	"github.com/Zyko0/go-sdl3/bin/binttf"
	"github.com/Zyko0/go-sdl3/img"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/ttf"
)

type GameState int

const (
	StartScreen GameState = iota
	Playing
	GameOver
)

const GameName string = "Flappy Bird"
const WindowWidth int = 576
const WindowHeight int = 1024
const TitleFontSize float32 = 90.0
const BirdWidth float32 = 68.0
const BirdHeight float32 = 48.0
const BirdGravity float32 = 0.25
const BirdMaxFallSpeed float32 = 14.0
const BirdFlapStrength float32 = -6.0
const FloorHeight float32 = 112.0
const PipesSpeed float32 = 2.0
const ParallaxSpeed float32 = 0.5

const RessourcesDir string = "res"
const FontsDir string = RessourcesDir + "/" + "fonts"
const FlappyTtfFont = FontsDir + "/" + "flappy.ttf"
const ImgDir string = RessourcesDir + "/" + "imgs"
const PipesDir string = ImgDir + "/" + "pipes"
const PipeGreenPath string = PipesDir + "/" + "pipegreen.png"
const PipeRedPath string = PipesDir + "/" + "pipered.png"
const BackgroundsDir string = ImgDir + "/" + "backgrounds"
const BaseImgPath string = BackgroundsDir + "/" + "base.png"
const RedBird string = ImgDir + "/redbird"
const BlueBird string = ImgDir + "/bluebird"
const YellowBird string = ImgDir + "/yellowbird"


var ColorWhite sdl.Color = sdl.Color{R: 255, G: 255, B: 255, A: 255}

var (
	chosenBackground string    = chooseBackground()
	chosenBird       string    = chooseBird()
	gameState        GameState = StartScreen
)

func initialize() error {
	defer binsdl.Load().Unload()
	defer binttf.Load().Unload()
	defer binimg.Load().Unload()
	defer sdl.Quit()
	defer ttf.Quit()
	defer sdl.CloseLibrary()
	defer ttf.CloseLibrary()
	defer img.CloseLibrary()

	// Init SDL
	var err error = sdl.Init(sdl.INIT_VIDEO | sdl.INIT_AUDIO | sdl.INIT_EVENTS | sdl.INIT_GAMEPAD)
	if err != nil {
		return fmt.Errorf("Error while loading SDL Library %v", err)
	}

	// Init Window
	window, renderer, err := sdl.CreateWindowAndRenderer(GameName, WindowWidth, WindowHeight, sdl.WINDOW_ALWAYS_ON_TOP)
	if err != nil {
		return fmt.Errorf("Error while intializing  %v", err)
	}

	defer renderer.Destroy()
	defer window.Destroy()

	// Init TTF Module
	err = ttf.Init()
	if err != nil {
		return fmt.Errorf("Error while intializing TTF module  %v", err)
	}

	// Init scenes
	titleScene, err := NewTitleScene(renderer, chosenBackground)
	if err != nil {
		return fmt.Errorf("Error while creating title scene  %v", err)
	}
	defer titleScene.Destroy()

	scene, err := NewScene(renderer, chosenBackground)
	if err != nil {
		return fmt.Errorf("Error while creating scene  %v", err)
	}
	defer scene.Destroy()

	sdl.RunLoop(func() error {
		var event sdl.Event

		for sdl.PollEvent(&event) {
			switch event.Type {
			// press anykey to start the game
			case sdl.EVENT_KEY_DOWN, sdl.EVENT_MOUSE_BUTTON_DOWN:
				if gameState == Playing {
					scene.fish.HandleInput(event)
					break
				}

				log.Println("Start the game!")
				gameState = Playing
			// Quit the game when the user clicks the close button
			case sdl.EVENT_QUIT:
				return sdl.EndLoop
			}
		}

		// Draw logic
		switch gameState {
		case StartScreen:
			titleScene.DrawScene(renderer)
		case Playing:
			scene.DrawScene(renderer)
		case GameOver:
		}

		// Update logic
		switch gameState {
		case StartScreen:
		case Playing:
			scene.UpdateScene()
		case GameOver:
		}

		return nil
	})

	return nil
}

func chooseBackground() string {
	random := int(math.Floor(float64(time.Now().Unix()))) % 2

	switch random {
	case 0:
		return BackgroundsDir + "/" + "background-day.png"
	case 1:
		return BackgroundsDir + "/" + "background-night.png"
	default:
		return BackgroundsDir + "/" + "background-day.png"
	}
}

func chooseBird() string {
	random := int(math.Floor(float64(time.Now().Unix()))) % 3

	switch random {
	case 0:
		return RedBird
	case 1:
		return BlueBird
	case 2:
		return YellowBird
	default:
		return RedBird
	}
}

func main() {
	var err error = initialize()
	if err != nil {
		log.Printf("Error: %v", err)
		os.Exit(2)
	}

	log.Println("Hello world!")
}
