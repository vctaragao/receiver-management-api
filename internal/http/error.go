package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type bussinesLogicErrorChecker func(error) bool

type ErrorOutputDto struct {
	Message string `json:"message"`
}

func returnError(ctx echo.Context, checker bussinesLogicErrorChecker, err error) {
	if checker(err) {
		ctx.JSON(http.StatusBadRequest, &ErrorOutputDto{Message: err.Error()})
	} else {
		fmt.Println("Erro inesperado: ", err)
		ctx.JSON(http.StatusInternalServerError, &ErrorOutputDto{Message: "erro inesperado"})
	}
}
