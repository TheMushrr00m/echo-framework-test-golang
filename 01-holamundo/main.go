package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	// Inicia instancia de "Echo" framework
	e := echo.New()

	// Manejador de rutas
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hola mundo")
	})

	// Inicia el servidor
	e.Start(":8080")
}
