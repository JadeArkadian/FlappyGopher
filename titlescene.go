package main

import (
	"fmt"

	"github.com/Zyko0/go-sdl3/img"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/ttf"
)

type TitleScene struct {
	backgroundTexture *sdl.Texture
	font              *ttf.Font
	textSurface       *sdl.Surface
	textTexture       *sdl.Texture
}

func NewTitleScene(renderer *sdl.Renderer, backgroundPath string) (*TitleScene, error) {
	texture, err := img.LoadTexture(renderer, backgroundPath)
	if err != nil {
		return nil, fmt.Errorf("Error while loading background image  %v", err)
	}

	font, err := ttf.OpenFont(FlappyTtfFont, TitleFontSize)
	if err != nil {
		return nil, fmt.Errorf("Error while opening flappy.ttf  %v", err)
	}

	textSurface, err := font.RenderTextSolid(GameName, ColorWhite)
	if err != nil {
		return nil, fmt.Errorf("Error while rendering title text  %v", err)
	}

	textTexture, err := renderer.CreateTextureFromSurface(textSurface)
	if err != nil {
		return nil, fmt.Errorf("Error while creating texture from Surface %v", err)
	}

	return &TitleScene{backgroundTexture: texture, font: font, textSurface: textSurface, textTexture: textTexture}, nil
}

func (scene *TitleScene) DrawScene(renderer *sdl.Renderer) {
	renderer.Clear()
	scene.drawBackground(renderer)
	scene.drawTitleText(renderer)
	renderer.Present()
}

func (scene *TitleScene) Destroy() {
	defer scene.backgroundTexture.Destroy()
	defer scene.textSurface.Destroy()
	defer scene.textTexture.Destroy()
}

func (scene *TitleScene) drawBackground(renderer *sdl.Renderer) {
	dst := sdl.FRect{X: 0, Y: 0, W: float32(WindowWidth), H: float32(WindowHeight)}
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.RenderTexture(scene.backgroundTexture, nil, &dst)
}

func (scene *TitleScene) drawTitleText(renderer *sdl.Renderer) {
	// Define destination rect to render at original size (no scaling)
	dst := sdl.FRect{X: 40, Y: 250, W: float32(scene.textSurface.W), H: float32(scene.textSurface.H)}
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.RenderTexture(scene.textTexture, nil, &dst)
}
