package services

import (
	"sync"

	interfaces "github.com/Patrignani/test-rover/core/interfaces/services"
	"github.com/Patrignani/test-rover/core/services"
)

var lock = &sync.Mutex{}

var roverService interfaces.RoverService

func GetRoverServiceInstance() interfaces.RoverService {
	if roverService == nil {
		lock.Lock()
		defer lock.Unlock()
		if roverService == nil {

			roverService = services.NewRoverService()
		}
	}
	return roverService
}
