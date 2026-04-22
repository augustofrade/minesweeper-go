package main

import (
	"fmt"

	gamestate "github.com/augustofrade/minesweeper-go/game"
	"github.com/augustofrade/minesweeper-go/mines"
	"github.com/augustofrade/minesweeper-go/shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	windowWidth := 500
	windowHeight := 800
	rl.InitWindow(int32(windowWidth), int32(windowHeight), "Minesweeper")
	defer rl.CloseWindow()

	game := gamestate.Instance()
	game.SetWindowSize(windowWidth, windowHeight)
	board := mines.NewEmptyBoard(shared.Size{Width: 10})
	board.CreateMines()

	rl.SetTargetFPS(60)
	rl.SetWindowState(rl.FlagWindowResizable)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		board.Draw()
		rl.DrawText(fmt.Sprintf("ScreenSize  = %d, %d", windowWidth, windowHeight), 20, 10, 30, rl.Black)
		rl.DrawText(fmt.Sprintf("Board Size  = %d, %d", board.Size.Width, board.Size.Height), 20, 40, 30, rl.Black)
		rl.DrawText(fmt.Sprintf("Mine Amount = %d", board.MineCount), 20, 70, 30, rl.Black)
		rl.DrawText(fmt.Sprintf("Mine Size   = %d", board.MineSize), 20, 100, 30, rl.Black)
		rl.DrawText(fmt.Sprintf("Offset 		 = %d, %d", board.Offset.X, board.Offset.Y), 20, 130, 30, rl.Black)

		if rl.IsWindowResized() {
			game.SetWindowSize(rl.GetScreenWidth(), rl.GetScreenHeight())
			// TODO: implement board rect size for better responsiveness
			board.UpdateWindowOffset()
			board.UpdateMinesPositionOnScreen()
		}

		rl.EndDrawing()
	}
}
