package main

import (
	"fmt"
	"os"

	"github.com/Patrignani/test-rover/core/singletons/handlers"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Erro: informe o caminho do arquivo como argumento")
		return
	}
	filePath := os.Args[1]
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	rovers, err := handlers.GetFileHandlerInstance().ReadFile(fileBytes)

	if err != nil {
		panic(err)
	}

	println("-------------- RESULTADO --------------")
	for _, rover := range rovers {
		println(rover.AxisX, rover.AxisY, rover.Direction)
	}

}
