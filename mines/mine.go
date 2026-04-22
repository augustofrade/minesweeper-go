package mines

import (
	gamestate "github.com/augustofrade/minesweeper-go/game"
	"github.com/augustofrade/minesweeper-go/shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Mine struct {
	Uncovered   bool
	TextureRect *rl.Rectangle
	Position    shared.Point
	Size        shared.Size
	HasBomb     bool
}

func NewMine(position shared.Point, size shared.Size) *Mine {
	return &Mine{
		Uncovered: false,
		Position:  position,
		Size:      size,
	}
}

func (mine *Mine) Draw() {
	pos := rl.NewVector2(float32(mine.Position.X), float32(mine.Position.Y))
	destRect := rl.NewRectangle(
		pos.X,
		pos.Y,
		float32(mine.Size.Width),
		float32(mine.Size.Height),
	)

	rl.DrawTexturePro(gamestate.Instance().Spritesheet, *mine.TextureRect, destRect, rl.NewVector2(0, 0), 0, rl.White)
}
