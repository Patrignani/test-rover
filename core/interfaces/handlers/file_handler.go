package handlers

import "github.com/Patrignani/test-rover/core/entity"

type FileHandler interface {
	ReadFile(file []byte) ([]entity.Rover, error)
}
