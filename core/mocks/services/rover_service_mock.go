package services

import (
	"github.com/Patrignani/test-rover/core/constant"
	"github.com/Patrignani/test-rover/core/entity"
	"github.com/stretchr/testify/mock"
)

type RoverServiceImplMock struct {
	mock.Mock
}

func (r *RoverServiceImplMock) SetNewCoordinates(initialCoordinates entity.Coordinates, rover *entity.Rover, instruction []constant.Instruction) {
	r.Called(initialCoordinates, rover, instruction)
}
