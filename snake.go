package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	RECTSIZE = 20
)

func (c *PlayerControlled) GrowBody(body []rl.Vector2) {
	tail := body[0]
	if len(c.Body) != 0 {
		tail = c.Body[len(c.Body)-1]
	}

	c.Body = append(c.Body, tail)
}
