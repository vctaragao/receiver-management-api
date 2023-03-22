package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vctaragao/receiver-management-api/internal/application"
	"github.com/vctaragao/receiver-management-api/internal/http"
	"github.com/vctaragao/receiver-management-api/internal/storage"
)

func main() {
	repo := storage.NewPostgress()

	rm := application.NewReceiverManagement(repo)

	e := echo.New()
	http.RegisterRouter(e, rm)

	e.Logger.Fatal(e.Start(":1323"))
}
