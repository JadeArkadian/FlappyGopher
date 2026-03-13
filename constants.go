package main

// Window settings
const (
	GameName     string = "Flappy Bird"
	WindowWidth  int    = 576
	WindowHeight int    = 1024
)

// Bird settings
const (
	BirdWidth        float32 = 68.0
	BirdHeight       float32 = 48.0
	BirdGravity      float32 = 0.25
	BirdMaxFallSpeed float32 = 14.0
	BirdFlapStrength float32 = -6.0
	BirdInitialX     float32 = 100.0
	BirdInitialY     float32 = 200.0
)

// Environment settings
const (
	FloorHeight   float32 = 112.0
	PipeWidth     float32 = 104.0
	PipeHeight    float32 = 640.0
	PipesSpeed    float32 = 3.0
	ParallaxSpeed float32 = 0.5
)

// UI settings
const (
	TitleFontSize float32 = 90.0
	TitleTextX    float32 = 40.0
	TitleTextY    float32 = 250.0
	TitleShadowX  float32 = 45.0
	TitleShadowY  float32 = 255.0
)

// Resource paths
const (
	ResourcesDir   string = "res"
	FontsDir       string = ResourcesDir + "/fonts"
	FlappyTtfFont  string = FontsDir + "/flappy.ttf"
	ImgDir         string = ResourcesDir + "/imgs"
	PipesDir       string = ImgDir + "/pipes"
	PipeGreenPath  string = PipesDir + "/pipegreen.png"
	PipeRedPath    string = PipesDir + "/pipered.png"
	BackgroundsDir string = ImgDir + "/backgrounds"
	BaseImgPath    string = BackgroundsDir + "/base.png"
	RedBird        string = ImgDir + "/redbird"
	BlueBird       string = ImgDir + "/bluebird"
	YellowBird     string = ImgDir + "/yellowbird"
)
