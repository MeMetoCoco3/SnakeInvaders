package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
)

const (
	SCREENWIDTH = 600
	SCREENHEIGHT
)
const (
	PLAYER_MOVEMENT_SPEED = 90
)

func main() {
	rl.InitWindow(SCREENWIDTH, SCREENHEIGHT, "Snake")
	rl.InitAudioDevice()
	defer rl.CloseWindow()
	defer rl.CloseAudioDevice()

	world := NewWorld()
	world.gameState.maxCandies = 5
	world.gameState.currentCandies = 0
	renderSys := *NewSystem(world, &DrawSystem{})
	movementSys := *NewSystem(world, &MovementSystem{})
	collisionSys := *NewSystem(world, &CollisionSystem{})
	audioSys := *NewSystem(world, &AudioSystem{})
	player := make(map[ComponentID]any)

	// player[positionID] = Position{
	// 	X: 200,
	// 	Y: 200,
	// }

	player[movementID] = Movement{Direction: rl.Vector2{0, 0}, Speed: 500, Timer: COUNT_FOR_TURN}
	player[playerControlledID] = PlayerControlled{
		Body: []Cell{
			{Position: rl.Vector2{X: 200, Y: 200}, Direction: rl.Vector2{1, 0}},
		},
	}

	player[collidesID] = Collides{
		X:     player[playerControlledID].(PlayerControlled).Body[0].Position.X,
		Y:     player[playerControlledID].(PlayerControlled).Body[0].Position.Y,
		Width: 20, Height: 20,
	}

	player[spriteID] = Sprite{Width: RECTSIZE, Height: RECTSIZE, Color: rl.Lime}

	border1 := make(map[ComponentID]any)
	border2 := make(map[ComponentID]any)
	border3 := make(map[ComponentID]any)
	border4 := make(map[ComponentID]any)
	border1[positionID] = Position{X: 0, Y: 0}
	border1[spriteID] = Sprite{Width: SCREENWIDTH, Height: 20, Color: rl.Red}
	border1[collidesID] = Collides{X: 0, Y: 0, Width: SCREENWIDTH, Height: 20}
	border2[positionID] = Position{X: SCREENWIDTH - 20, Y: 0}
	border2[spriteID] = Sprite{Width: 100, Height: SCREENHEIGHT, Color: rl.Red}
	border2[collidesID] = Collides{X: SCREENWIDTH - 20, Y: 0, Width: 100, Height: SCREENHEIGHT}
	border3[positionID] = Position{X: 0, Y: 0}
	border3[spriteID] = Sprite{Width: 20, Height: SCREENHEIGHT, Color: rl.Red}
	border3[collidesID] = Collides{X: 0, Y: 0, Width: 20, Height: SCREENHEIGHT}
	border4[positionID] = Position{X: 0, Y: SCREENHEIGHT - 20}
	border4[spriteID] = Sprite{Width: SCREENWIDTH, Height: 100, Color: rl.Red}
	border4[collidesID] = Collides{X: 0, Y: SCREENHEIGHT - 20, Width: SCREENWIDTH, Height: 100}
	world.CreateEntity(player)
	world.CreateEntity(border1)
	world.CreateEntity(border2)
	world.CreateEntity(border3)
	world.CreateEntity(border4)

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		audioSys.Update(dt)
		movementSys.Update(dt)
		collisionSys.Update(dt)
		if world.gameState.currentCandies < world.gameState.maxCandies {
			world.gameState.currentCandies += 1
			log.Println("CANDY GENERATED")
			c := CandyGenerator()
			world.CreateEntity(c)
		}
		log.Println(len(player[playerControlledID].(PlayerControlled).Body))
		rl.BeginDrawing()
		rl.ClearBackground(VICOLOR)
		renderSys.Update(dt)
		rl.EndDrawing()
	}
}
