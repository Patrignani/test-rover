package services

import (
	"github.com/Patrignani/test-rover/core/constant"
	"github.com/Patrignani/test-rover/core/entity"
)

type RoverService interface {
	SetNewCoordinates(initialCoordinates entity.Coordinates, rover *entity.Rover, instruction []constant.Instruction)
}
