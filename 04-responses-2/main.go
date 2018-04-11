package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type Person struct {
	FirstName string `json:"first_name"`
	LasName string `json:"last_name"`
	Age int `json:"age"`
}

func main() {
	e := echo.New()

	ps := make([]Person, 0)
	p := Person{
		FirstName:	"Gabriel",
		LasName: 	"Cueto",
		Age: 		23,
	}
	ps = append(ps, p)
	ps = append(ps, p)
	e.GET("/json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, ps)
	})

	e.GET("/xml", func(c echo.Context) error {
		return c.XML(http.StatusOK, ps)
	})

	e.Start(":8080")
}
