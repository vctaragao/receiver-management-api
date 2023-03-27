package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vctaragao/receiver-management-api/internal/application"
)

type DeleteInputDto struct {
	ReceiversIds []uint `json:"recebedores_id"`
}

func DeleteReceiver(rm *application.ReceiverManagement) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		dto := new(DeleteInputDto)

		if err := ctx.Bind(dto); err != nil {
			return err
		}

		if err := rm.Delete(dto.ReceiversIds); err != nil {
			returnError(ctx, rm.IsDeleteBussinesLogicError, err)
			return nil
		}

		return ctx.JSON(http.StatusNoContent, nil)
	}
}
