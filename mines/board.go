package mines

import (
	gamestate "github.com/augustofrade/minesweeper-go/game"
	"github.com/augustofrade/minesweeper-go/shared"
)

type Board struct {
	Size      shared.Size
	Mines     [][]Mine
	MineCount int
	MineSize  int
	Offset    shared.Point
}

func NewEmptyBoard(size shared.Size) Board {

	game := gamestate.Instance()

	mineSize := game.ScreenSize.Width / size.Width

	if size.Height == 0 {
		size.Height = game.ScreenSize.Height / mineSize
	}

	board := Board{
		Mines:    make([][]Mine, size.Width),
		Size:     size,
		MineSize: mineSize,
		Offset: shared.Point{
			X: (game.ScreenSize.Width - (mineSize * size.Width)) / 2,
			Y: (game.ScreenSize.Height - (mineSize * size.Height)) / 2,
		},
	}

	return board
}

func (board *Board) CreateMines() {
	for col := 0; col < board.Size.Width; col++ {
		board.Mines[col] = make([]Mine, board.Size.Height)
		for row := 0; row < board.Size.Height; row++ {
			mineX := (board.MineSize * col) + board.Offset.X
			mineY := (board.MineSize * row) + board.Offset.Y

			mine := NewMine(shared.Point{X: mineX, Y: mineY}, shared.Size{
				Width:  board.MineSize,
				Height: board.MineSize,
			})

			mine.TextureRect = gamestate.Instance().GetDefaultTileTextureRect()
			board.Mines[col][row] = *mine
			board.MineCount++
		}
	}
}

func (board *Board) Draw() {

	for col := 0; col < board.Size.Width; col++ {
		for row := 0; row < board.Size.Height; row++ {
			board.Mines[col][row].Draw()
		}
	}
}
