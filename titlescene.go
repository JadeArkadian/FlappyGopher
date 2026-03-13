package main

import (
	"fmt"

	"github.com/Zyko0/go-sdl3/img"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/ttf"
)

// TitleScene renders the start screen with background and title text.
type TitleScene struct {
	backgroundTexture *sdl.Texture
	font              *ttf.Font
	textSurface       *sdl.Surface
	textTexture       *sdl.Texture
}

// NewTitleScene creates the title screen scene.
func NewTitleScene(renderer *sdl.Renderer, backgroundPath string) (*TitleScene, error) {
	texture, err := img.LoadTexture(renderer, backgroundPath)
	if err != nil {
		return nil, fmt.Errorf("error loading background image: %w", err)
	}

	font, err := ttf.OpenFont(FlappyTtfFont, TitleFontSize)
	if err != nil {
		return nil, fmt.Errorf("error opening font %q: %w", FlappyTtfFont, err)
	}

	textSurface, err := font.RenderTextSolid(GameName, ColorWhite)
	if err != nil {
		return nil, fmt.Errorf("error rendering title text: %w", err)
	}

	textTexture, err := renderer.CreateTextureFromSurface(textSurface)
	if err != nil {
		return nil, fmt.Errorf("error creating texture from surface: %w", err)
	}

	return &TitleScene{
		backgroundTexture: texture,
		font:              font,
		textSurface:       textSurface,
		textTexture:       textTexture,
	}, nil
}

// DrawScene renders the title screen for the current frame.
func (titleScene *TitleScene) DrawScene(renderer *sdl.Renderer) {
	renderer.Clear()
	titleScene.drawBackground(renderer)
	titleScene.drawTitleTextShadow(renderer)
	titleScene.drawTitleText(renderer)
	renderer.Present()
}

// Destroy releases all resources held by the title scene.
func (titleScene *TitleScene) Destroy() {
	titleScene.backgroundTexture.Destroy()
	titleScene.textTexture.Destroy()
	titleScene.textSurface.Destroy()
	titleScene.font.Close()
}

func (titleScene *TitleScene) drawBackground(renderer *sdl.Renderer) {
	renderer.SetDrawColor(255, 255, 255, 255)
	dst := sdl.FRect{X: 0, Y: 0, W: float32(WindowWidth), H: float32(WindowHeight)}
	renderer.RenderTexture(titleScene.backgroundTexture, nil, &dst)
}

func (titleScene *TitleScene) drawTitleText(renderer *sdl.Renderer) {
	dst := sdl.FRect{
		X: TitleTextX,
		Y: TitleTextY,
		W: float32(titleScene.textSurface.W),
		H: float32(titleScene.textSurface.H),
	}
	titleScene.textTexture.SetColorMod(255, 255, 255)
	renderer.RenderTexture(titleScene.textTexture, nil, &dst)
}

func (titleScene *TitleScene) drawTitleTextShadow(renderer *sdl.Renderer) {
	dst := sdl.FRect{
		X: TitleShadowX,
		Y: TitleShadowY,
		W: float32(titleScene.textSurface.W),
		H: float32(titleScene.textSurface.H),
	}
	titleScene.textTexture.SetColorMod(0, 0, 0)
	renderer.RenderTexture(titleScene.textTexture, nil, &dst)
}
