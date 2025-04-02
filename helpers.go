package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/constraints"
	"math/rand"
)

var VICOLOR = rl.Color{252, 163, 17, 255}

type Number interface {
	constraints.Integer | constraints.Float
}

func CandyGenerator() map[ComponentID]any {
	c := make(map[ComponentID]any)
	c[candyID] = Candy{}
	x := float32(rand.Intn(500) + 40)
	y := float32(rand.Intn(500) + 40)
	c[positionID] = Position{X: x, Y: y}
	c[spriteID] = Sprite{Width: 20, Height: 20, Color: rl.Blue}
	c[collidesID] = Collides{X: x, Y: y, Width: 20, Height: 20}

	return c
}

func GetMaskFromComponents(componentsID ...ComponentID) ComponentID {
	var mask ComponentID
	for i := range len(componentsID) {
		mask |= componentsID[i]
	}
	return mask
}

// HACK: Add components here as needed.
func GetArrayComponentsFromID(id ComponentID) any {
	switch id {
	case positionID:
		return make([]Position, 0)
	case spriteID:
		return make([]Sprite, 0)
	case movementID:
		return make([]Movement, 0)
	case healthID:
		return make([]Health, 0)
	case aliveID:
		return make([]Alive, 0)
	case animationID:
		return make([]Animation, 0)
	case playerControlledID:
		return make([]PlayerControlled, 0)
	case IAControlledID:
		return make([]IAControlled, 0)
	case collidesID:
		return make([]Collides, 0)
	case enemyID:
		return make([]Enemy, 0)
	case candyID:
		return make([]Candy, 0)

	default:
		return nil
	}
}

func GetComponentsFromMask(mask ComponentID) []ComponentID {
	components := make([]ComponentID, 32)
	count := 0
	bitValue := 1
	for mask != 0 {
		bit := mask & 1
		if bit != 0 {
			components[count] = ComponentID(bitValue)
			count++
		}
		bitValue = bitValue << 1
		mask = mask >> 1
	}
	return components[:count]
}

func hasComponent(mask, componentsMask ComponentID) bool {
	newMask := (uint32(mask) & uint32(componentsMask))
	result := newMask == uint32(mask)
	return result
}

func GetInput(c Movement, dt float32) Movement {
	CurrentDirection := c.Direction

	if rl.IsKeyDown(rl.KeyUp) {
		CurrentDirection = DIRECTIONS[0]
	} else if rl.IsKeyDown(rl.KeyRight) {
		CurrentDirection = DIRECTIONS[1]
	} else if rl.IsKeyDown(rl.KeyDown) {
		CurrentDirection = DIRECTIONS[2]
	} else if rl.IsKeyDown(rl.KeyLeft) {
		CurrentDirection = DIRECTIONS[3]
	}

	c.Direction = CurrentDirection

	return c
}

type collisionType uint8

const (
	noC collisionType = iota
	topC
	bottomC
	rightC
	leftC
	overlapC
)

func CheckRectCollision(aPos Position, aSize Collides, bPos Position, bSize Collides) collisionType {
	aTop := aPos.Y
	aBottom := aPos.Y + aSize.Height
	aRight := aPos.X + aSize.Width
	aLeft := aPos.X

	bTop := bPos.Y
	bBottom := bPos.Y + bSize.Height
	bRight := bPos.X + bSize.Width
	bLeft := bPos.X

	centerA := rl.Vector2{X: aPos.X + aSize.Width/2, Y: aPos.Y + aSize.Height/2}
	centerB := rl.Vector2{X: bPos.X + bSize.Width/2, Y: bPos.Y + bSize.Height/2}

	// No collision
	if aRight < bLeft || aLeft > bRight ||
		aBottom < bTop || aTop > bBottom {
		return noC
	}

	relativePosition := rl.Vector2{X: centerA.X - centerB.X, Y: centerA.Y - centerB.Y}
	if aBottom >= bTop && aTop <= bTop &&
		aRight > bLeft && aLeft < bRight {
		if abs(relativePosition.X)/bSize.Width > abs(relativePosition.Y)/bSize.Height {
			if relativePosition.X > 0 {
				return rightC
			}
			return leftC
		}
		return topC
	}

	if aTop <= bBottom && aBottom >= bBottom &&
		aRight > bLeft && aLeft < bRight {
		if abs(relativePosition.X)/bSize.Width > abs(relativePosition.Y)/bSize.Height {
			if relativePosition.X > 0 {
				return rightC
			}
			return leftC
		}
		return bottomC
	}

	if aRight >= bLeft && aLeft <= bLeft &&
		aBottom > bTop && aTop < bBottom {
		if abs(relativePosition.X)/bSize.Width > abs(relativePosition.Y)/bSize.Height {
			if relativePosition.X > 0 {
				return rightC
			}
			return leftC
		}
		return bottomC
	}

	if aLeft <= bRight && aRight >= bRight &&
		aBottom > bTop && aTop < bBottom {

		if abs(relativePosition.X)/bSize.Width > abs(relativePosition.Y)/bSize.Height {
			if relativePosition.X > 0 {
				return rightC
			}
			return leftC
		}
		return bottomC
	}
	// If we get here, we got full overlap
	return overlapC
}

func convertToRectangle(v Collides) rl.Rectangle {
	return rl.Rectangle{
		X:      v.X,
		Y:      v.Y,
		Width:  v.Width,
		Height: v.Height,
	}

}

func abs[T Number](a T) T {
	var v T
	if a < v {
		return -a
	}
	return a
}
func min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}
