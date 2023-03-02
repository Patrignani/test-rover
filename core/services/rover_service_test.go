package services

import (
	"testing"

	"github.com/Patrignani/test-rover/core/constant"
	"github.com/Patrignani/test-rover/core/entity"
	"github.com/stretchr/testify/assert"
)

func Test_When_Initial_Coordinates_Is_Smaller_Should_Keep_The_Value_Of_The_Initial_Coordinate(t *testing.T) {
	service := NewRoverService()
	inital := entity.NewCoordinates(0, 0)
	rover := entity.NewRover(2, 2, constant.North)
	instructions := []constant.Instruction{constant.Left}
	service.SetNewCoordinates(inital, rover, instructions)

	assert.Equal(t, rover.AxisX, 0)
	assert.Equal(t, rover.AxisY, 0)
	assert.Equal(t, rover.Direction, constant.West)
}

func Test_When_Set_New_Coordinate_Should_Expected_Output_Direction_North(t *testing.T) {
	service := NewRoverService()

	inital := entity.NewCoordinates(5, 5)
	rover := entity.NewRover(2, 1, constant.North)

	instructions := []constant.Instruction{constant.Left, constant.Move, constant.Left, constant.Move, constant.Left, constant.Move, constant.Left, constant.Move, constant.Move}

	service.SetNewCoordinates(inital, rover, instructions)

	assert.Equal(t, rover.AxisX, 1)
	assert.Equal(t, rover.AxisY, 3)
	assert.Equal(t, rover.Direction, constant.North)
}

func Test_When_Set_New_Coordinate_Should_Expected_Output_Direction_East(t *testing.T) {
	service := NewRoverService()

	inital := entity.NewCoordinates(5, 5)
	rover := entity.NewRover(3, 3, constant.East)

	instructions := []constant.Instruction{constant.Move, constant.Move, constant.Right, constant.Move, constant.Move, constant.Right, constant.Move, constant.Right, constant.Right, constant.Move}

	service.SetNewCoordinates(inital, rover, instructions)

	assert.Equal(t, rover.AxisX, 5)
	assert.Equal(t, rover.AxisY, 1)
	assert.Equal(t, rover.Direction, constant.East)
}
