package application

import (
	"errors"

	"github.com/vctaragao/receiver-management-api/internal/application/entity"
	"github.com/vctaragao/receiver-management-api/internal/application/usecase/create_receiver"
)

type ReceiverManagement struct {
	Repo *entity.Repository
}

func NewReceiverManagement(repo *entity.Repository) *ReceiverManagement {
	return &ReceiverManagement{
		Repo: repo,
	}
}

func (rm *ReceiverManagement) Create(corporateName, cpf, cnpj, email, pixType, pixKey string) (*create_receiver.OutputDto, error) {
	dto := &create_receiver.InputDto{
		CorporateName: corporateName,
		Cpf:           cpf,
		Cnpj:          cnpj,
		Email:         email,
		PixType:       pixType,
		PixKey:        pixKey,
	}

	usecase := &create_receiver.Create{Repo: *rm.Repo}

	return usecase.Execute(dto)
}

func (rm *ReceiverManagement) IsCreateBussinesLogicError(err error) bool {
	for _, t := range create_receiver.GetBussinesLogicErrors() {
		if errors.As(err, &t) {
			return true
		}
	}
	return false
}
