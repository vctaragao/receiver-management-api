package application

import (
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
	"github.com/vctaragao/receiver-management-api/internal/application/usecase"
)

type ReceiverManagement struct {
	Repo entity.Repository
}

func NewReceiverManagement(repo entity.Repository) *ReceiverManagement {
	return &ReceiverManagement{
		Repo: repo,
	}
}

func (rm *ReceiverManagement) Create(corporateName, cpf, cnpj, email string) (*usecase.CreateOutputDto, error) {
	dto := &usecase.CreateInputDto{
		CorporateName: corporateName,
		Cpf:           cpf,
		Cnpj:          cnpj,
		Email:         email,
	}

	return usecase.Create(dto, rm.Repo)
}
