package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Zyko0/go-sdl3/bin/binsdl"
	"github.com/Zyko0/go-sdl3/bin/binttf"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/ttf"
)

const GameName string = "Flappy Gopher"
const WindowWidth int = 800
const WindowHeight int = 600
const TitleFontSize float32 = 100.0
const FlappyTtfFont = "res/fonts/flappy.ttf"

var ColorWhite sdl.Color = sdl.Color{R: 255, G: 255, B: 255, A: 255}


func initialize() error {
	defer binsdl.Load().Unload()
	defer binttf.Load().Unload()
	defer sdl.Quit()
	defer ttf.Quit()
	
	// Init SDL
	var err error = sdl.Init(sdl.INIT_VIDEO | sdl.INIT_AUDIO | sdl.INIT_EVENTS | sdl.INIT_GAMEPAD)
	if err != nil {
		return fmt.Errorf("Error while loading SDL Library %v", err )
	}

	// Init Window
	window, renderer, err := sdl.CreateWindowAndRenderer(GameName, WindowWidth, WindowHeight, sdl.WINDOW_ALWAYS_ON_TOP)
	if err != nil {
		return fmt.Errorf("Error while intializing  %v", err )
	}

	defer renderer.Destroy()
	defer window.Destroy()

	// Init TTF Module
	err = ttf.Init()
	if err != nil {
		return fmt.Errorf("Error while intializing TTF module  %v", err )
	}

	return drawTitle(renderer)
}


func drawTitle(renderer *sdl.Renderer) error {

	font, err := ttf.OpenFont(FlappyTtfFont,TitleFontSize)
	if err != nil {
		return fmt.Errorf("Error while opening flappy.ttf  %v", err )
	}
	defer font.Close()

	surface, err := font.RenderTextSolid(GameName, ColorWhite)
	if err != nil {
		return fmt.Errorf("Error while rendering title text  %v", err )
	}
	defer surface.Destroy()

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return fmt.Errorf("Error while creating texture from Surface %v", err )
	}
	defer texture.Destroy()

	// Define destination rect to render at original size (no scaling)
	dst := sdl.FRect{X: 50, Y: 250, W: float32(surface.W), H: float32(surface.H)}
	renderer.SetDrawColor(255, 255, 255, 255)
	
	sdl.RunLoop(func() error {
		var event sdl.Event

		for sdl.PollEvent(&event) {
			if event.Type == sdl.EVENT_QUIT {
				return sdl.EndLoop
			}
		}

		//renderer.Clear()
		renderer.RenderTexture(texture, nil, &dst)
		renderer.Present()

		return nil
	})

	return nil
}

func main() {

	var err error = initialize()
	if err != nil {
		os.Exit(2)
	}

	log.Println("Hello world!")
}
