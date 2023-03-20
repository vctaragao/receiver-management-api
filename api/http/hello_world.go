package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func HelloWorld() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		data := &HelloResponse{
			Message: "Hello World",
		}

		return ctx.JSON(http.StatusOK, data)
	}
}
