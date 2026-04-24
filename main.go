package main

import (
	"fmt"

	gamestate "github.com/augustofrade/minesweeper-go/game"
	"github.com/augustofrade/minesweeper-go/mines"
	"github.com/augustofrade/minesweeper-go/shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	windowWidth := 600
	windowHeight := 800
	rl.InitWindow(int32(windowWidth), int32(windowHeight), "Minesweeper")
	defer rl.CloseWindow()

	game := gamestate.Instance()
	game.SetWindowSize(windowWidth, windowHeight)
	board := mines.NewEmptyBoard(shared.Size{Width: 10})
	board.CreateMines()

	var mousePosition rl.Vector2

	rl.SetTargetFPS(60)
	rl.SetWindowState(rl.FlagWindowResizable)
	bg := rl.NewColor(236, 228, 215, 1)

	for !rl.WindowShouldClose() {
		mousePosition = rl.GetMousePosition()

		if rl.IsWindowResized() {
			game.SetWindowSize(rl.GetScreenWidth(), rl.GetScreenHeight())
			board.UpdateRectWidth()
			board.UpdateWindowOffset()
			board.UpdateMineSize()
			board.UpdateMinesPositionOnScreen()
		}

		board.HandleMouseClicks(&mousePosition)

		rl.BeginDrawing()

		rl.ClearBackground(bg)
		board.Draw()
		rl.DrawText(fmt.Sprintf("ScreenSize  = %d, %d", game.ScreenSize.Width, game.ScreenSize.Height), 20, 10, 30, rl.Black)
		rl.DrawText(fmt.Sprintf("Board Width = %d", board.RectWidth), 20, 40, 30, rl.Black)
		rl.DrawText(fmt.Sprintf("Board Size  = %d, %d", board.Size.Width, board.Size.Height), 20, 70, 30, rl.Black)
		rl.DrawText(fmt.Sprintf("Mine Amount = %d", board.MineCount), 20, 100, 30, rl.Black)
		rl.DrawText(fmt.Sprintf("Mine Size   = %d", *board.MineSize), 20, 130, 30, rl.Black)
		rl.DrawText(fmt.Sprintf("Offset 		 = %d, %d", board.Offset.X, board.Offset.Y), 20, 160, 30, rl.Black)
		rl.DrawText(fmt.Sprintf("Mouse 		   = %.2f, %.2f", mousePosition.X, mousePosition.Y), 20, 190, 30, rl.Black)

		rl.EndDrawing()
	}
}
