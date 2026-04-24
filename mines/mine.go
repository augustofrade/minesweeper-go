package mines

import (
	"fmt"

	gamestate "github.com/augustofrade/minesweeper-go/game"
	"github.com/augustofrade/minesweeper-go/shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Mine struct {
	Uncovered   bool
	TextureRect *rl.Rectangle
	Bounds      *rl.Rectangle
	GridCoords  *shared.Point
	Size        *int
	HasBomb     bool
	IsFlagged   bool
	IsRevealed  bool
}

func NewMine(bounds rl.Rectangle, gridCoords *shared.Point, size *int) *Mine {
	fmt.Println(bounds)
	return &Mine{
		Uncovered:  false,
		Bounds:     &bounds,
		GridCoords: gridCoords,
		Size:       size,
		IsFlagged:  false,
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
	if mine.IsRevealed {
		return
	}

	if !mine.IsFlagged {
		mine.IsFlagged = true
		mine.TextureRect = gamestate.Instance().GetFlagTileTextureRect()
	} else {
		mine.IsFlagged = false
		mine.TextureRect = gamestate.Instance().GetDefaultTileTextureRect()
	}
}

func (mine *Mine) Reveal(bombAmount int) {
	if mine.IsRevealed {
		return
	}

	mine.IsRevealed = true

	if bombAmount == 0 {
		mine.TextureRect = gamestate.Instance().GetEmptyTileTextureRect()
	} else {
		mine.TextureRect = gamestate.Instance().GetTextureRectForMineNumber(bombAmount)
	}
}
