package entity

import (
	"testing"

	"github.com/Patrignani/test-rover/core/constant"
	"github.com/stretchr/testify/assert"
)

func Test_When_Moving_Left_While_In_The_North_Should_Stay_In_The_West(t *testing.T) {
	testMoveLeft(t, constant.North, constant.West)
}

func Test_When_Moving_Left_While_In_The_East_Should_Stay_In_The_North(t *testing.T) {
	testMoveLeft(t, constant.East, constant.North)
}

func Test_When_Moving_Left_While_In_The_South_Should_Stay_In_The_East(t *testing.T) {
	testMoveLeft(t, constant.South, constant.East)
}

func Test_When_Moving_Left_While_In_The_West_Should_Stay_In_The_South(t *testing.T) {
	testMoveLeft(t, constant.West, constant.South)
}

func Test_When_Moving_Right_While_In_The_North_Should_Stay_In_The_East(t *testing.T) {
	testMoveRight(t, constant.North, constant.East)
}

func Test_When_Moving_Right_While_In_The_East_Should_Stay_In_The_South(t *testing.T) {
	testMoveRight(t, constant.East, constant.South)
}

func Test_When_Moving_Right_While_In_The_South_Should_Stay_In_The_West(t *testing.T) {
	testMoveRight(t, constant.South, constant.West)
}

func Test_When_Moving_Right_While_In_The_West_Should_Stay_In_The_North(t *testing.T) {
	testMoveRight(t, constant.West, constant.North)
}

func Test_When_Moving_Axis_X_To_West_Should_Subtract_Value(t *testing.T) {
	testMoveAxisX(t, true, 1, 0)
}

func Test_When_Moving_Axis_X_To_East_Should_Add_Value(t *testing.T) {
	testMoveAxisX(t, false, 1, 2)
}

func Test_When_Moving_Axis_X_To_West_Should_Subtract_Valuee_And_Keep_Negative_Value_As_Zero(t *testing.T) {
	testMoveAxisX(t, true, 0, 0)
}

func Test_When_Moving_Axis_Y_To_North_Should_Add_Value(t *testing.T) {
	testMoveAxisY(t, false, 1, 2)
}

func Test_When_Moving_Axis_Y_To_South_Should_Subtract_Value(t *testing.T) {
	testMoveAxisY(t, true, 1, 0)
}

func Test_When_Moving_Axis_Y_To_South_Should_Subtract_Value_And_Keep_Negative_Value_As_Zero(t *testing.T) {
	testMoveAxisY(t, true, 0, 0)
}

func getRoverTest(direction constant.Direction) Rover {
	return *NewRover(0, 0, direction)
}

func testMoveLeft(t *testing.T, entryValue constant.Direction, outputValue constant.Direction) {
	r := getRoverTest(entryValue)
	r.MoveLeft()
	assert.Equal(t, outputValue, r.Direction)
}

func testMoveRight(t *testing.T, entryValue constant.Direction, outputValue constant.Direction) {
	r := getRoverTest(entryValue)
	r.MoveRight()
	assert.Equal(t, outputValue, r.Direction)
}

func testMoveAxisX(t *testing.T, subtract bool, entryValue int, outputValue int) {
	r := Rover{
		Coordinates: Coordinates{
			AxisX: entryValue,
			AxisY: 0,
		},
	}
	r.MoveAxisX(subtract)
	assert.Equal(t, outputValue, r.AxisX)
}

func testMoveAxisY(t *testing.T, subtract bool, entryValue int, outputValue int) {
	r := Rover{
		Coordinates: Coordinates{
			AxisX: 0,
			AxisY: entryValue,
		},
	}
	r.MoveAxisY(subtract)
	assert.Equal(t, outputValue, r.AxisY)
}
