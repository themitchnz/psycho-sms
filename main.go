package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(256, 224, "Psycho SMS")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Wow!, Nothing works yet!", 20, 20, 15, rl.LightGray)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
