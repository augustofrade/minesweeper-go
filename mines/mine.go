package mines

import (
	"fmt"

	gamestate "github.com/augustofrade/minesweeper-go/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Mine struct {
	Uncovered   bool
	TextureRect *rl.Rectangle
	Bounds      *rl.Rectangle
	Size        *int
	HasBomb     bool
	IsFlagged   bool
}

func NewMine(bounds rl.Rectangle, size *int) *Mine {
	fmt.Println(bounds)
	return &Mine{
		Uncovered: false,
		Bounds:    &bounds,
		Size:      size,
		IsFlagged: false,
	}
}

func (mine *Mine) Draw() {
	pos := rl.NewVector2(float32(mine.Bounds.X), float32(mine.Bounds.Y))
	destRect := rl.NewRectangle(
		pos.X,
		pos.Y,
		float32(*mine.Size),
		float32(*mine.Size),
	)

	rl.DrawTexturePro(gamestate.Instance().Spritesheet, *mine.TextureRect, destRect, rl.NewVector2(0, 0), 0, rl.White)
}

func (mine *Mine) Flag() {
	if !mine.IsFlagged {
		mine.IsFlagged = true
		mine.TextureRect = gamestate.Instance().GetFlagTileTextureRect()
	} else {
		mine.IsFlagged = false
		mine.TextureRect = gamestate.Instance().GetDefaultTileTextureRect()
	}
}
