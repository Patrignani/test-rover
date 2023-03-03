package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Patrignani/test-rover/core/singletons/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/upload", uploadTxt)

	e.Start(":8080")
}

func uploadTxt(c echo.Context) error {

	fileBytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, "Falha ao ler o arquivo")
	}

	rovers, err := handlers.GetFileHandlerInstance().ReadFile(fileBytes)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	var result string

	for _, rover := range rovers {
		result += fmt.Sprintf("%d %d %s \n", rover.AxisX, rover.AxisY, rover.Direction)
	}

	return c.String(http.StatusOK, result)
}
