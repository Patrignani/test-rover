package handlers

import (
	"sync"

	"github.com/Patrignani/test-rover/core/handlers"
	interfaces "github.com/Patrignani/test-rover/core/interfaces/handlers"
	singletonsServices "github.com/Patrignani/test-rover/core/singletons/services"
)

var lock = &sync.Mutex{}

var handler interfaces.FileHandler

func GetFileHandlerInstance() interfaces.FileHandler {
	if handler == nil {
		lock.Lock()
		defer lock.Unlock()
		if handler == nil {

			handler = handlers.NewFileHandler(singletonsServices.GetRoverServiceInstance())
		}
	}
	return handler
}
