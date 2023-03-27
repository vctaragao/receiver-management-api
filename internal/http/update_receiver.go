package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vctaragao/receiver-management-api/internal/application"
)

type UpdateInputDto struct {
	ReceiverId    uint   `json:"recebedor_id"`
	CorporateName string `json:"razao_social"`
	CpfCnpj       string `json:"cpf_cnpj"`
	Email         string `json:"email"`
	PixType       string `json:"pix_type"`
	PixKey        string `json:"pix_key"`
}

type UpdateOutputDto struct {
	ReceiverId    uint   `json:"recebedor_id"`
	CorporateName string `json:"razao_social"`
	CpfCnpj       string `json:"cpf_cnpj"`
	Email         string `json:"email"`
	Status        string `json:"status"`
	PixType       string `json:"pix_type"`
	PixKey        string `json:"pix_key"`
}

func UpdateReceiver(rm *application.ReceiverManagement) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		dto := new(UpdateInputDto)

		if err := ctx.Bind(dto); err != nil {
			return err
		}

		resultDto, err := rm.Update(dto.ReceiverId, dto.CorporateName, dto.CpfCnpj, dto.Email, dto.PixType, dto.PixKey)

		if err != nil {
			return ctx.JSON(http.StatusBadRequest, &ErrorOutputDto{Message: err.Error()})
		}

		return ctx.JSON(http.StatusOK, &UpdateOutputDto{
			ReceiverId:    resultDto.ReceiverId,
			CorporateName: resultDto.CorporateName,
			CpfCnpj:       resultDto.CpfCnpj,
			Email:         resultDto.Email,
			Status:        resultDto.Status,
			PixType:       resultDto.PixType,
			PixKey:        resultDto.PixKey,
		})
	}
}
