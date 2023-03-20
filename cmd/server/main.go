package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vctaragao/receiver-management-api/api/http"
)

func main() {
	e := echo.New()

	http.RegisterRouter(e)

	e.Logger.Fatal(e.Start(":1323"))
}
