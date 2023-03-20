package http

import (
	"github.com/labstack/echo/v4"
)

func RegisterRouter(e *echo.Echo) {

	e.GET("/", HelloWorld())
}
