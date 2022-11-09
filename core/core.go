package core

import (
	"fmt"

	"github.com/Noreen-py/Bander/interpreter/token"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Core struct {
	lexed []token.Token //change to parsed
}

func (c *Core) Window() {
	fmt.Println(c.lexed)
	screenWidth := 512
	screenHeight := 400
	rl.InitWindow(int32(screenWidth), int32(screenHeight), "Bandel")
	rl.SetTargetFPS(60)
	ballPosition := rl.NewVector2(float32(800)/2, float32(450)/2)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawCircleV(ballPosition, 50, rl.Maroon)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

/*
func (c *Core) Eval() {
	switch c.lexed[0].Type {
	case token.EOF:
		fmt.Println(324234)
		c.ballPosition.X += 300
	}
}
*/
