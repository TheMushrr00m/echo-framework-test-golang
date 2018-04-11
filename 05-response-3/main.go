package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()

	e.Static("/", "public")

	e.GET("/gopher/:format", func(c echo.Context) error {
		p := c.Param("format")
		if p == "download" {
			return c.Inline("img/gopher.jpeg", "gopher.jpeg")
		} else if p == "jpeg" {
			return c.File("img/gopher.jpeg")
		} else if p == "att" {
			return c.Attachment("img/gopher.jpeg", "gopher-file.jpeg")
		}
		return c.HTML(http.StatusNotFound, "<h1>Content not found</h1>")
	})

	e.Start(":8080")
}
