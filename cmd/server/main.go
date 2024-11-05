package main

import (
	"log"

	"github.com/FreitasGabriel/live-streaming-server/live-streaming-server/config/db"
	"github.com/FreitasGabriel/live-streaming-server/live-streaming-server/internal/handler"
	"github.com/FreitasGabriel/live-streaming-server/live-streaming-server/internal/repository"
	"github.com/FreitasGabriel/live-streaming-server/live-streaming-server/internal/service"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := db.OpenConn()
	if err != nil {
		log.Fatalln("error to connect database", err)
	}
	defer db.Close()

	repo := repository.NewKeysRepository(db)
	service := service.NewKeysService(repo)
	handler := handler.NewKeysHandler(service)
	log.Default().Println("Routing...")

	e.POST("/auth", handler.AuthStreamingKey)
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	e.Logger.Fatal(e.Start(":8000").Error())
}
