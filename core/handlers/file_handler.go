package handlers

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Patrignani/test-rover/core/constant"
	"github.com/Patrignani/test-rover/core/entity"
	interfacesHandlers "github.com/Patrignani/test-rover/core/interfaces/handlers"
	interfacesServices "github.com/Patrignani/test-rover/core/interfaces/services"
)

type FileHandlerImp struct {
	roverService interfacesServices.RoverService
}

func NewFileHandler(roverService interfacesServices.RoverService) interfacesHandlers.FileHandler {
	return &FileHandlerImp{roverService: roverService}
}

func (f *FileHandlerImp) ReadFile(file []byte) ([]entity.Rover, error) {
	if len(file) == 0 {
		return nil, errors.New("Arquivo vazio")
	}

	txtFile := string(file)
	fileLines := strings.Split(txtFile, "\r\n")
	var initialCoordinates entity.Coordinates
	var rovers []entity.Rover
	var currentRover *entity.Rover

	if len(fileLines) < 3 || len(fileLines)%2 == 0 {
		return nil, errors.New("Arquivo contém número de linhas inválidos")
	}

	for i, line := range fileLines {
		if i == 0 {
			initlaValue, err := f.getInitialCoordinates(line)

			if err != nil {
				return nil, err
			}

			initialCoordinates = *initlaValue
		} else if i%2 != 0 {
			roverResult, err := f.getRover(line, i)

			if err != nil {
				return nil, err
			}

			currentRover = roverResult
		} else {
			var instructions []constant.Instruction

			for _, instruction := range line {
				newInstruction, err := constant.ToInstruction(string(instruction))

				if err != nil {
					return nil, err
				}

				instructions = append(instructions, newInstruction)
			}

			f.roverService.SetNewCoordinates(initialCoordinates, currentRover, instructions)

			rovers = append(rovers, *currentRover)
		}
	}

	return rovers, nil
}

func (f *FileHandlerImp) getInitialCoordinates(line string) (*entity.Coordinates, error) {
	coordinates := strings.Split(line, " ")

	if len(coordinates) < 2 {
		return nil, errors.New("Primeira linha do arquivo está inválida")
	}

	axisX, err := strconv.Atoi(coordinates[0])

	if err != nil {
		return nil, errors.New("Erro ao converter o eixo X do arquivo TXT da cordenada inicial")
	}

	axisY, err := strconv.Atoi(coordinates[1])

	if err != nil {
		return nil, errors.New("Erro ao converter o eixo Y do arquivo TXT da cordenada inicial")
	}

	initialCoordinates := entity.NewCoordinates(axisY, axisX)

	return &initialCoordinates, nil
}

func (f *FileHandlerImp) getRover(line string, index int) (*entity.Rover, error) {
	roverInfo := strings.Split(line, " ")

	if len(roverInfo) < 3 {
		return nil, errors.New("Erro valor do comando do rover inválido da linha: " + strconv.Itoa(index+1))
	}

	axisX, err := strconv.Atoi(roverInfo[0])

	if err != nil {
		return nil, errors.New("Erro ao converter o eixo X do arquivo TXT da cordenada do rover da linha: " + strconv.Itoa(index+1))
	}

	axisY, err := strconv.Atoi(roverInfo[1])

	if err != nil {
		return nil, errors.New("Erro ao converter o eixo Y do arquivo TXT da cordenada do rover da linha: " + strconv.Itoa(index+1))
	}

	direction, err := constant.ToDirection(roverInfo[2])

	if err != nil {
		return nil, errors.New("Erro ao converter a cordenada do rover da linha: " + strconv.Itoa(index+1))
	}

	return entity.NewRover(axisY, axisX, direction), nil
}
