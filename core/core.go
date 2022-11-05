package core

import (
	"github.com/Noreen-py/Bander/interpreter/token"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Core struct {
	ballPosition rl.Vector2
	lexed        []token.Token //change to parsed
}

func (c *Core) Window() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		c.ballPosition = rl.NewVector2(float32(800)/2, float32(450)/2)
		rl.ClearBackground(rl.RayWhite)
		rl.DrawCircleV(c.ballPosition, 50, rl.Maroon)
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
