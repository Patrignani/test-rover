package handlers

import (
	"errors"
	"testing"

	mocks "github.com/Patrignani/test-rover/core/mocks/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_When_Empty_File_Should_Return_Error(t *testing.T) {
	roverService := &mocks.RoverServiceImplMock{}
	fileHandler := NewFileHandler(roverService)

	emptyFile := []byte{}
	expectedErr := errors.New("Arquivo vazio")

	rovers, err := fileHandler.ReadFile(emptyFile)

	assert.Nil(t, rovers)
	assert.EqualError(t, err, expectedErr.Error())
}

func Test_When_One_Line_Should_Return_Error(t *testing.T) {
	roverService := &mocks.RoverServiceImplMock{}
	fileHandler := NewFileHandler(roverService)

	file := []byte("5 5")
	expectedErr := errors.New("Arquivo contém número de linhas inválidos")

	rovers, err := fileHandler.ReadFile(file)

	assert.Nil(t, rovers)
	assert.EqualError(t, err, expectedErr.Error())
}

func Test_When_Even_Line_Value_Should_Return_Error(t *testing.T) {
	roverService := &mocks.RoverServiceImplMock{}
	fileHandler := NewFileHandler(roverService)

	file := []byte("5 5\r\n 1 1 N")
	expectedErr := errors.New("Arquivo contém número de linhas inválidos")

	rovers, err := fileHandler.ReadFile(file)

	assert.Nil(t, rovers)
	assert.EqualError(t, err, expectedErr.Error())
}

func Test_When_First_Line_Invalid_Should_Return_Error(t *testing.T) {
	roverService := &mocks.RoverServiceImplMock{}
	fileHandler := NewFileHandler(roverService)

	file := []byte("5\r\n 1 1 N\r\nMLLM")
	expectedErr := errors.New("Primeira linha do arquivo está inválida")

	rovers, err := fileHandler.ReadFile(file)

	assert.Nil(t, rovers)
	assert.EqualError(t, err, expectedErr.Error())
}

func Test_When_First_Line_Invalid_Axis_Y_Should_Return_Error(t *testing.T) {
	roverService := &mocks.RoverServiceImplMock{}
	fileHandler := NewFileHandler(roverService)

	file := []byte("5 N\r\n 1 1 N\r\nMLLM")
	expectedErr := errors.New("Erro ao converter o eixo Y do arquivo TXT da cordenada inicial")

	rovers, err := fileHandler.ReadFile(file)

	assert.Nil(t, rovers)
	assert.EqualError(t, err, expectedErr.Error())
}

func Test_When_First_Line_Invalid_Axis_X_Should_Return_Error(t *testing.T) {
	roverService := &mocks.RoverServiceImplMock{}
	fileHandler := NewFileHandler(roverService)

	file := []byte("N 5\r\n 1 1 N\r\nMLLM")
	expectedErr := errors.New("Erro ao converter o eixo X do arquivo TXT da cordenada inicial")

	rovers, err := fileHandler.ReadFile(file)

	assert.Nil(t, rovers)
	assert.EqualError(t, err, expectedErr.Error())
}

func Test_When_Invalid_command_Rover_Should_Return_Error(t *testing.T) {
	roverService := &mocks.RoverServiceImplMock{}
	fileHandler := NewFileHandler(roverService)

	file := []byte("5 5\r\n1 1\r\nMLLM")
	expectedErr := errors.New("Erro valor do comando do rover inválido da linha: 2")

	rovers, err := fileHandler.ReadFile(file)

	assert.Nil(t, rovers)
	assert.EqualError(t, err, expectedErr.Error())
}

func Test_When_Command_Rover_Invalid_Axis_Y_Should_Return_Error(t *testing.T) {
	roverService := &mocks.RoverServiceImplMock{}
	fileHandler := NewFileHandler(roverService)

	file := []byte("5 5\r\n1 N N\r\nMLLM")
	expectedErr := errors.New("Erro ao converter o eixo Y do arquivo TXT da cordenada do rover da linha: 2")

	rovers, err := fileHandler.ReadFile(file)

	assert.Nil(t, rovers)
	assert.EqualError(t, err, expectedErr.Error())
}

func Test_When_Command_Rover_Invalid_Axis_X_Should_Return_Error(t *testing.T) {
	roverService := &mocks.RoverServiceImplMock{}
	fileHandler := NewFileHandler(roverService)

	file := []byte("5 5\r\nN 1 N\r\nMLLM")
	expectedErr := errors.New("Erro ao converter o eixo X do arquivo TXT da cordenada do rover da linha: 2")

	rovers, err := fileHandler.ReadFile(file)

	assert.Nil(t, rovers)
	assert.EqualError(t, err, expectedErr.Error())
}

func Test_When_Command_Rover_Invalid_Direction_Should_Return_Error(t *testing.T) {
	roverService := &mocks.RoverServiceImplMock{}
	fileHandler := NewFileHandler(roverService)

	file := []byte("5 5\r\n1 1 J\r\nMLLM")
	expectedErr := errors.New("Erro ao converter a cordenada do rover da linha: 2")

	rovers, err := fileHandler.ReadFile(file)

	assert.Nil(t, rovers)
	assert.EqualError(t, err, expectedErr.Error())
}

func Test_When_Instruction_Invalid_Should_Return_Error(t *testing.T) {
	roverService := &mocks.RoverServiceImplMock{}
	fileHandler := NewFileHandler(roverService)

	file := []byte("5 5\r\n1 1 N\r\nJ")
	expectedErr := errors.New("Instrução não reconhecida J")

	rovers, err := fileHandler.ReadFile(file)

	assert.Nil(t, rovers)
	assert.EqualError(t, err, expectedErr.Error())
}

func Test_When_File_Is_Valid_Should_Return_Rovers(t *testing.T) {
	roverService := &mocks.RoverServiceImplMock{}
	fileHandler := NewFileHandler(roverService)

	file := []byte("5 5\r\n1 1 N\r\nMLLM")

	roverService.On("SetNewCoordinates", mock.Anything, mock.Anything, mock.Anything)

	rovers, err := fileHandler.ReadFile(file)

	assert.Nil(t, err)
	assert.NotNil(t, rovers)

}
