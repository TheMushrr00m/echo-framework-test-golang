package main

import "github.com/labstack/echo"

func main() {
	e := echo.New()

	// Manejadores de rutas
	// Servimos un archivo estatico
	//e.File("/", "public/index.html")
	e.Static("/", "public")


	e.Start(":8080")
}
