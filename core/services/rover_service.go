package services

import (
	"github.com/Patrignani/test-rover/core/constant"
	"github.com/Patrignani/test-rover/core/entity"
	interfaces "github.com/Patrignani/test-rover/core/interfaces/services"
)

type RoverServiceImpl struct {
}

func NewRoverService() interfaces.RoverService {
	return &RoverServiceImpl{}
}

func (r *RoverServiceImpl) SetNewCoordinates(initialCoordinates entity.Coordinates, rover *entity.Rover, instruction []constant.Instruction) {

	for _, i := range instruction {

		if i == constant.Left {
			rover.MoveLeft()
		} else if i == constant.Right {
			rover.MoveRight()
		} else {
			switch rover.Direction {
			case constant.North:
				rover.MoveAxisY(false)
			case constant.East:
				rover.MoveAxisX(false)
			case constant.South:
				rover.MoveAxisY(true)
			case constant.West:
				rover.MoveAxisX(true)
			}
		}

		if rover.AxisX > initialCoordinates.AxisX {
			rover.AxisX = initialCoordinates.AxisX
		}

		if rover.AxisY > initialCoordinates.AxisY {
			rover.AxisY = initialCoordinates.AxisY
		}

	}

}
