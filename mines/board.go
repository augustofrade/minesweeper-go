package mines

import (
	gamestate "github.com/augustofrade/minesweeper-go/game"
	"github.com/augustofrade/minesweeper-go/shared"
)

type Board struct {
	Size      shared.Size
	MineGrid  [][]*Mine
	MineList  []*Mine
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
		MineGrid: make([][]*Mine, size.Width),
		Size:     size,
		MineSize: mineSize,
		Offset: shared.Point{
			X: 0,
			Y: 0,
		},
	}

	board.UpdateWindowOffset()

	return board
}

func (board *Board) UpdateWindowOffset() {
	game := gamestate.Instance()
	board.Offset.X = (game.ScreenSize.Width - (board.MineSize * board.Size.Width)) / 2
	board.Offset.Y = (game.ScreenSize.Height - (board.MineSize * board.Size.Height)) / 2
}

func (board *Board) UpdateMinesPositionOnScreen() {
	for col := 0; col < board.Size.Width; col++ {
		for row := 0; row < board.Size.Height; row++ {
			mine := board.MineGrid[col][row]
			mine.Position.X = board.getMineXPosition(col)
			mine.Position.Y = board.getMineYPosition(row)
		}
	}
}

func (board *Board) CreateMines() {
	for col := 0; col < board.Size.Width; col++ {
		board.MineGrid[col] = make([]*Mine, board.Size.Height)
		for row := 0; row < board.Size.Height; row++ {
			mineX := board.getMineXPosition(col)
			mineY := board.getMineYPosition(row)

			mine := NewMine(shared.Point{X: mineX, Y: mineY}, shared.Size{
				Width:  board.MineSize,
				Height: board.MineSize,
			})

			mine.TextureRect = gamestate.Instance().GetDefaultTileTextureRect()
			board.MineGrid[col][row] = mine
			board.MineList = append(board.MineList, mine)
			board.MineCount++
		}
	}
}

func (board *Board) Draw() {

	for col := 0; col < board.Size.Width; col++ {
		for row := 0; row < board.Size.Height; row++ {
			board.MineGrid[col][row].Draw()
		}
	}
}

func (board *Board) getMineXPosition(col int) int {
	return (board.MineSize * col) + board.Offset.X
}

func (board *Board) getMineYPosition(row int) int {
	return (board.MineSize * row) + board.Offset.Y
}
