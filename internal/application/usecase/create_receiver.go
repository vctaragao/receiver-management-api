package usecase

import (
	"errors"

	"github.com/vctaragao/receiver-management-api/internal/application/entity"
)

type CreateInputDto struct {
	CorporateName string
	Cpf           string
	Cnpj          string
	Email         string
}

type CreateOutputDto struct {
	Id uint
	CreateInputDto
}

func Create(dto *CreateInputDto, repo entity.Repository) (*CreateOutputDto, error) {
	receiver := entity.NewReceiver(dto.CorporateName, dto.Cpf, dto.Cnpj, dto.Email, entity.STATUS_DRAFT)
	err := receiver.Validate()
	if err != nil {
		return &CreateOutputDto{}, err
	}

	receiver, err = repo.AddReceiver(receiver)

	if err != nil {
		return &CreateOutputDto{}, errors.New("unable to create receiver")
	}

	out := &CreateOutputDto{
		Id:             receiver.Id,
		CreateInputDto: *dto,
	}

	return out, nil
}
