package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	RECTSIZE       = 20
	COUNT_FOR_TURN = 40
)

func (c *PlayerControlled) GrowBody(body []Cell) {
	tail := body[0]
	if len(c.Body) != 0 {
		tail = c.Body[len(c.Body)-1]
	}
	newBody := Cell{
		Position:  rl.Vector2{X: tail.Position.X - RECTSIZE, Y: tail.Position.Y},
		Direction: rl.Vector2{X: tail.Direction.X, Y: tail.Direction.Y},
	}
	c.Body = append(c.Body, newBody)
}
