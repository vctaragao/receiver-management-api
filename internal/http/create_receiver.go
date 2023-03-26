package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vctaragao/receiver-management-api/internal/application"
)

type ReceiverInputDto struct {
	CorporateName string `json:"razao_social"`
	Cpf           string `json:"cpf"`
	Cnpj          string `json:"cnpj"`
	Email         string `json:"email"`
	PixType       string `json:"pix_type"`
	PixKey        string `json:"pix_key"`
}

type ReceiverOutputDto struct {
	Id uint `json:"recebedor_id"`
	ReceiverInputDto
}

type ErrorOutputDto struct {
	Message string `json:"message"`
}

func CreateReceiver(rm *application.ReceiverManagement) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		dto := new(ReceiverInputDto)

		if err := ctx.Bind(dto); err != nil {
			return err
		}

		resultDto, err := rm.Create(dto.CorporateName, dto.Cpf, dto.Cnpj, dto.Email, dto.PixType, dto.PixKey)

		if err != nil {
			returnError(ctx, rm, err)
			return nil
		}

		return ctx.JSON(http.StatusOK, &ReceiverOutputDto{Id: resultDto.Id, ReceiverInputDto: *dto})
	}
}

func returnError(ctx echo.Context, rm *application.ReceiverManagement, err error) {
	if rm.IsCreateBussinesLogicError(err) {
		ctx.JSON(http.StatusBadRequest, &ErrorOutputDto{Message: err.Error()})
	} else {
		ctx.JSON(http.StatusInternalServerError, &ErrorOutputDto{Message: "erro inesperado"})
	}
}
