package gamestate

import (
	"fmt"

	"github.com/augustofrade/minesweeper-go/shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var gameSingleton *GameState

type GameState struct {
	Spritesheet  rl.Texture2D
	textureRects map[string]*rl.Rectangle
	ScreenSize   shared.Size
	SpriteSize   float32
	Score        int
}

func create() *GameState {
	game := GameState{
		Spritesheet:  rl.LoadTexture("assets/spritesheet.png"),
		textureRects: make(map[string]*rl.Rectangle),
		Score:        0,
	}

	maxBombs := 9

	game.SpriteSize = float32(game.Spritesheet.Width / int32(maxBombs))

	for i := 1; i <= maxBombs; i++ {
		game.createTextureRect(fmt.Sprintf("mine_%d", i), shared.Point{X: 1, Y: 0})
	}

	game.createTextureRect("empty", shared.Point{X: 0, Y: 1})
	game.createTextureRect("default", shared.Point{X: 1, Y: 1})
	game.createTextureRect("flag", shared.Point{X: 2, Y: 1})
	game.createTextureRect("bomb", shared.Point{X: 3, Y: 1})

	gameSingleton = &game
	return gameSingleton
}

func Instance() *GameState {
	if gameSingleton == nil {
		return create()
	}
	return gameSingleton
}

func (game *GameState) SetWindowSize(width int, height int) {
	game.ScreenSize = shared.Size{Width: width, Height: height}
}

func (game *GameState) GetTextureRectForMineNumber(number int) *rl.Rectangle {
	textureRect := game.textureRects[fmt.Sprint(number)]
	if textureRect == nil {
		panic("Invalid texture")
	}
	return textureRect
}

func (game *GameState) GetDefaultTileTextureRect() *rl.Rectangle {
	return game.textureRects["default"]
}

func (game *GameState) GetEmptyTileTextureRect() *rl.Rectangle {
	return game.textureRects["empty"]
}
func (game *GameState) GetFlagTileTextureRect() *rl.Rectangle {
	return game.textureRects["flag"]
}
func (game *GameState) GetBombTileTextureRect() *rl.Rectangle {
	return game.textureRects["bomb"]
}

func (game *GameState) createTextureRect(key string, position shared.Point) {
	textureRect := rl.NewRectangle(float32(position.X*int(game.SpriteSize)), float32(position.Y*int(game.SpriteSize)), game.SpriteSize, game.SpriteSize)
	game.textureRects[key] = &textureRect
}
