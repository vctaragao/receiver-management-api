package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vctaragao/receiver-management-api/internal/application"
)

type ReceiverInputDto struct {
	RazaoSocial string `json:"razao_social"`
	Cpf         string `json:"cpf"`
	Cnpj        string `json:"cnpj"`
	Email       string `json:"email"`
}

type ReceiverOutputDto struct {
	Id uint `json:"recebedor_id"`
}

func CreateReceiver(rm *application.ReceiverManagement) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		dto := new(ReceiverInputDto)

		if err := ctx.Bind(dto); err != nil {
			return err
		}

		resultDto, err := rm.Create(dto.RazaoSocial, dto.Cpf, dto.Cnpj, dto.Email)

		if err != nil {
			return ctx.String(http.StatusBadRequest, err.Error())
		}

		output := &ReceiverOutputDto{Id: resultDto.Id}

		return ctx.JSON(http.StatusOK, output)
	}
}
