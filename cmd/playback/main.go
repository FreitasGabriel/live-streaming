package main

import (
	"log"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func serveStream() echo.HandlerFunc {
	return func(c echo.Context) error {
		streamName := c.Param("live")
		filePath := c.Param("*")

		if filePath == "" {
			filePath = "index.m3u8"
		}

		fileStreamPth := filepath.Join("hls/live/", streamName, filePath)
		log.Default().Println("stream file requested", fileStreamPth)

		return c.File(fileStreamPth)
	}
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	e.GET("/live/:live/*", serveStream())
	e.Logger.Fatal(e.Start(":8001").Error())
}
