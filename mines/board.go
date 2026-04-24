package mines

import (
	gamestate "github.com/augustofrade/minesweeper-go/game"
	"github.com/augustofrade/minesweeper-go/shared"
)

var MaxRectWidth int = 1000

type Board struct {
	RectWidth int
	Size      shared.Size
	MineGrid  [][]*Mine
	MineList  []*Mine
	MineCount int
	MineSize  *int
	Offset    shared.Point
}

func NewEmptyBoard(size shared.Size) Board {

	game := gamestate.Instance()

	rectWidth := min(game.ScreenSize.Width, MaxRectWidth)

	mineSize := rectWidth / size.Width

	if size.Height == 0 {
		size.Height = game.ScreenSize.Height / mineSize
	}

	board := Board{
		RectWidth: rectWidth,
		MineGrid:  make([][]*Mine, size.Width),
		Size:      size,
		MineSize:  &mineSize,
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
	board.Offset.X = (game.ScreenSize.Width - (*board.MineSize * board.Size.Width)) / 2
	board.Offset.Y = (game.ScreenSize.Height - (*board.MineSize * board.Size.Height)) / 2
}

func (board *Board) UpdateRectWidth() {
	game := gamestate.Instance()
	board.RectWidth = min(game.ScreenSize.Width, MaxRectWidth)
}

func (board *Board) UpdateMineSize() {
	// TODO: fiz this to use height as well, otherwise mines get pushed outside the top and bottom of the screen
	mineSize := board.RectWidth / board.Size.Width
	*board.MineSize = mineSize
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

			mine := NewMine(shared.Point{X: mineX, Y: mineY}, board.MineSize)

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
	return (*board.MineSize * col) + board.Offset.X
}

func (board *Board) getMineYPosition(row int) int {
	return (*board.MineSize * row) + board.Offset.Y
}
