package http

import (
	"github.com/labstack/echo/v4"
	"github.com/vctaragao/receiver-management-api/internal/application"
)

func RegisterRouter(e *echo.Echo, rm *application.ReceiverManagement) {
	e.GET("/list", ListReceiver(rm))
	e.POST("/create", CreateReceiver(rm))
	e.PATCH("/update", UpdateReceiver(rm))
	e.POST("/delete", DeleteReceiver(rm))
}
