package mines

import (
	"math/rand"
	"time"

	gamestate "github.com/augustofrade/minesweeper-go/game"
	"github.com/augustofrade/minesweeper-go/shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const MaxRectWidth int = 1000

type Board struct {
	RectWidth int
	Size      shared.Size
	MineGrid  [][]*Mine
	MineList  []*Mine
	MineCount int
	MineSize  *int
	BombCount int
	Offset    shared.Point
}

func NewEasyBoard() Board {
	return NewEmptyBoard(shared.Size{Width: 9, Height: 9}, 10)
}

func NewMediumBoard() Board {
	return NewEmptyBoard(shared.Size{Width: 16, Height: 16}, 40)
}
func NewHardBoard() Board {
	return NewEmptyBoard(shared.Size{Width: 30, Height: 16}, 99)
}

func NewEmptyBoard(size shared.Size, bombCount int) Board {

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
		BombCount: bombCount,
		Offset: shared.Point{
			X: 0,
			Y: 0,
		},
	}

	board.UpdateWindowOffset()

	return board
}

func (board *Board) HandleMouseClicks(mousePosition *rl.Vector2) {
	for _, mine := range board.MineList {
		if rl.IsMouseButtonReleased(rl.MouseLeftButton) && rl.CheckCollisionPointRec(*mousePosition, *mine.Bounds) {
			board.handleMineClick(mine)
			return
		}

		if rl.IsMouseButtonReleased(rl.MouseRightButton) && rl.CheckCollisionPointRec(*mousePosition, *mine.Bounds) {
			board.handleMineFlagClick(mine)
			return
		}
	}
}

func (board *Board) handleMineClick(mine *Mine) {
	board.revealMineAndNeighbors(mine)
}

func (board *Board) revealMineAndNeighbors(mine *Mine) {
	if !mine.CanBeInteracted() {
		return
	}

	if mine.HasBomb {
		mine.Reveal(0)
		panic("a")
	}

	bombAmount := board.getSurroundingBombAmount(mine)
	mine.Reveal(bombAmount)

	if bombAmount == 0 {
		board.forEachSurroundingMine(board.revealMineAndNeighbors, mine)
	}
}

func (board *Board) handleMineFlagClick(mine *Mine) {
	mine.Flag()
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
	mineSize := board.RectWidth / board.Size.Width
	if mineSize*board.Size.Height < gamestate.Instance().ScreenSize.Height {
		*board.MineSize = mineSize
	}
}

func (board *Board) UpdateMinesPositionOnScreen() {
	mineSize := float32(*board.MineSize)

	for col := 0; col < board.Size.Width; col++ {
		for row := 0; row < board.Size.Height; row++ {
			mine := board.MineGrid[col][row]

			mineX := board.getMineXPosition(col)
			mineY := board.getMineYPosition(row)
			mine.Bounds = &rl.Rectangle{
				X:      float32(mineX),
				Y:      float32(mineY),
				Width:  mineSize,
				Height: mineSize,
			}

		}
	}
}

func (board *Board) CreateMines() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for col := 0; col < board.Size.Width; col++ {
		board.MineGrid[col] = make([]*Mine, board.Size.Height)
		for row := 0; row < board.Size.Height; row++ {
			mineX := board.getMineXPosition(col)
			mineY := board.getMineYPosition(row)

			mine := NewMine(rl.Rectangle{
				X:      float32(mineX),
				Y:      float32(mineY),
				Width:  float32(*board.MineSize),
				Height: float32(*board.MineSize),
			},
				&shared.Point{
					X: col,
					Y: row,
				},
				board.MineSize,
				false)

			mine.TextureRect = gamestate.Instance().GetDefaultTileTextureRect()
			board.MineGrid[col][row] = mine
			board.MineList = append(board.MineList, mine)
			board.MineCount++
		}
	}

	rng.Shuffle(len(board.MineList), func(i, j int) {
		board.MineList[i], board.MineList[j] = board.MineList[j], board.MineList[i]
	})
	for i := 0; i < board.BombCount && i < len(board.MineList); i++ {
		board.MineList[i].HasBomb = true
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

func (board *Board) getSurroundingBombAmount(mine *Mine) int {
	amount := 0
	board.forEachSurroundingMine(func(foundMine *Mine) {
		if foundMine.HasBomb {
			amount++
		}
	}, mine)
	return amount
}

func (board *Board) forEachSurroundingMine(callback func(foundMine *Mine), mine *Mine) {
	grid := board.MineGrid
	x := mine.GridCoords.X
	y := mine.GridCoords.Y
	max := board.Size

	// left
	if x > 0 {
		callback(grid[x-1][y])
	}
	// top-left
	if x > 0 && y > 0 {
		callback(grid[x-1][y-1])
	}
	// bottom-left
	if x > 0 && y < max.Height-1 {
		callback(grid[x-1][y+1])
	}
	// top
	if y > 0 {
		callback(grid[x][y-1])
	}
	// bottom
	if y < max.Height-1 {
		callback(grid[x][y+1])
	}
	// top-right
	if x < max.Width-1 && y > 0 {
		callback(grid[x+1][y-1])
	}
	// right
	if x < max.Width-1 {
		callback(grid[x+1][y])
	}
	// bottom-right
	if x < max.Width-1 && y < max.Height-1 {
		callback(grid[x+1][y+1])
	}
}
