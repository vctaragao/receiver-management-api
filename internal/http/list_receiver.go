package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vctaragao/receiver-management-api/internal/application"
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
)

type ListInputDto struct {
	Search string `query:"search"`
	Page   int    `query:"page"`
}

type ListOutputDto struct {
	Total       int           `json:"total_recebedores"`
	CurrentPage int           `json:"pagina_atual"`
	Receivers   []ReceiverDto `json:"recebedores"`
}

type ReceiverDto struct {
	ReceiverId    uint   `json:"recebedor_id"`
	CorporateName string `json:"razao_social"`
	CpfCnpj       string `json:"cpf_cnpj"`
	Email         string `json:"email"`
	Status        string `json:"status"`
	PixType       string `json:"pix_type"`
	PixKey        string `json:"pix_key"`
}

func ListReceiver(rm *application.ReceiverManagement) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		dto := new(ListInputDto)

		if err := ctx.Bind(dto); err != nil {
			return err
		}

		if dto.Page <= 0 {
			dto.Page = 1
		}

		resultDto, err := rm.List(dto.Search, dto.Page)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, &ErrorOutputDto{Message: err.Error()})
		}

		outputDto := createListOutPuts(resultDto.Receivers)
		outputDto.Total = resultDto.Total
		outputDto.CurrentPage = dto.Page

		return ctx.JSON(http.StatusOK, outputDto)
	}
}

func createListOutPuts(receivers []entity.Receiver) *ListOutputDto {
	rList := []ReceiverDto{}
	for _, receiver := range receivers {
		rDto := ReceiverDto{
			ReceiverId:    receiver.Id,
			CorporateName: receiver.CorporateName,
			CpfCnpj:       receiver.CpfCnpj,
			Email:         receiver.Email,
			Status:        receiver.Status,
			PixType:       receiver.Pix.Type,
			PixKey:        receiver.Pix.Key,
		}

		rList = append(rList, rDto)
	}

	return &ListOutputDto{Receivers: rList}
}
