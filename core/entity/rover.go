package entity

import (
	"github.com/Patrignani/test-rover/core/constant"
)

var movementLeft map[constant.Direction]constant.Direction
var movementRight map[constant.Direction]constant.Direction

type Rover struct {
	Coordinates
	Direction constant.Direction
}

func NewRover(axisY int, axisX int, direction constant.Direction) *Rover {
	return &Rover{
		Coordinates: NewCoordinates(axisY, axisX),
		Direction:   direction,
	}
}

func (r *Rover) MoveLeft() {
	r.loadMovementLeft()
	r.Direction = movementLeft[r.Direction]
}

func (r *Rover) MoveRight() {
	r.loadMovementRight()
	r.Direction = movementRight[r.Direction]
}

func (r *Rover) MoveAxisX(subtract bool) {
	if subtract {
		r.AxisX -= 1
	} else {
		r.AxisX += 1
	}

	if r.AxisX < 0 {
		r.AxisX = 0
	}
}

func (r *Rover) MoveAxisY(subtract bool) {
	if subtract {
		r.AxisY -= 1
	} else {
		r.AxisY += 1
	}

	if r.AxisY < 0 {
		r.AxisY = 0
	}
}

func (r *Rover) loadMovementRight() {
	if movementRight == nil {
		movementRight = make(map[constant.Direction]constant.Direction)
		movementRight[constant.North] = constant.East
		movementRight[constant.East] = constant.South
		movementRight[constant.South] = constant.West
		movementRight[constant.West] = constant.North
	}
}

func (r *Rover) loadMovementLeft() {
	if movementLeft == nil {
		movementLeft = make(map[constant.Direction]constant.Direction)
		movementLeft[constant.North] = constant.West
		movementLeft[constant.West] = constant.South
		movementLeft[constant.South] = constant.East
		movementLeft[constant.East] = constant.North
	}
}
