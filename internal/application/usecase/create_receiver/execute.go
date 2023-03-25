package create_receiver

import (
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
)

type Create struct {
	Repo entity.Repository
}

func (c *Create) Execute(dto *InputDto) (*OutputDto, error) {
	receiver := entity.NewReceiver(dto.CorporateName, dto.Cpf, dto.Cnpj, dto.Email, entity.STATUS_DRAFT)

	if err := receiver.Validate(); err != nil {
		return returnError(&CreatingReceiverErr{err: err})
	}

	pix := entity.NewPix(dto.Pix_type, dto.Pix_key)

	if err := pix.Validate(); err != nil {
		return returnError(&CreatingPixErr{err: err})
	}

	receiver, err := c.Repo.AddReceiver(receiver)
	if err != nil {
		return returnError(&SaveReceiverErr{err: err})
	}

	_, err = c.Repo.AddPix(receiver.Id, pix)
	if err != nil {
		return returnError(&SavingPixErr{err: err})
	}

	out := &OutputDto{
		Id:       receiver.Id,
		InputDto: *dto,
	}

	return out, nil
}

func returnError(err error) (*OutputDto, error) {
	return &OutputDto{}, err
}