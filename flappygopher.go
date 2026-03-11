package main

import (
	"log"
	"os"

	"github.com/Zyko0/go-sdl3/bin/binsdl"
	"github.com/Zyko0/go-sdl3/sdl"
)

const GameName string = "Flappy Gopher"
const WindowWidth int = 800
const WindowHeight int = 600

func main() {
	defer binsdl.Load().Unload()
	defer sdl.Quit()

	var err error = sdl.Init(sdl.INIT_VIDEO | sdl.INIT_AUDIO | sdl.INIT_EVENTS | sdl.INIT_GAMEPAD)

	if err != nil {
		log.Println("Error while loading SDL Library")
		os.Exit(2)
	}

	sdl.CreateWindowAndRenderer(GameName, WindowWidth, WindowHeight, sdl.WINDOW_ALWAYS_ON_TOP)

	log.Println("Hello world!")
}
