package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/FreitasGabriel/live-streaming-server/live-streaming-server/internal/model"
	"github.com/FreitasGabriel/live-streaming-server/live-streaming-server/internal/service"
	"github.com/labstack/echo/v4"
)

type IKeysHandler interface {
	AuthStreamingKey(ctx echo.Context) error
}

type keysHandler struct {
	keysService service.IKeysService
}

func NewKeysHandler(keysService service.IKeysService) IKeysHandler {
	return &keysHandler{keysService}
}

func (kh *keysHandler) AuthStreamingKey(c echo.Context) error {
	log.Default().Println("Running Auth")
	body := c.Request().Body
	defer body.Close()

	fields, _ := io.ReadAll(body)
	streamingKeys := getStreamKey(fields)

	keys, err := kh.keysService.AuthStramingKey(streamingKeys.Name, streamingKeys.Key)
	if err != nil {
		return c.String(http.StatusBadRequest, "problem with streaming key")
	}

	if keys.Key == "" {
		log.Default().Println("Forbidden User.")
		return c.String(http.StatusForbidden, "Forbidden")
	}

	log.Default().Println("User authenticated")

	newStreamURL := fmt.Sprintf("rtmp://127.0.0.1:1935/hls-live/%s", keys.Name)
	log.Default().Println("Redirecting to: ", newStreamURL)
	return c.Redirect(http.StatusFound, newStreamURL)
}

func getStreamKey(s []byte) model.Keys {
	var authValues model.Keys

	pairs := strings.Split(string(s), "&")

	for _, pair := range pairs {
		sliptPair := strings.Split(pair, "=")
		key := sliptPair[0]
		value := sliptPair[1]

		if key == "name" {
			allPassedValues := strings.Split(value, "_")
			authValues.Name = allPassedValues[0]
			authValues.Key = allPassedValues[1]
		}

	}

	return authValues
}
