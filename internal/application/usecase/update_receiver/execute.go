package update_receiver

import (
	"errors"

	"github.com/vctaragao/receiver-management-api/internal/application/entity"
)

type Update struct {
	Repo entity.Repository
}

func (u *Update) Execute(receiverId uint, dto *InputDto) (*OutputDto, error) {
	receiver, err := u.Repo.GetReceiverWithPix(receiverId)
	if err != nil {
		return returnError(&findingReceiverError{err: err})
	}

	if err := receiver.Update(dto.CorporateName, dto.Cpf, dto.Cnpj, dto.Email); err != nil {
		return returnError(errors.New("error updating"))
	}

	if err := receiver.Validate(); err != nil {
		return returnError(&invalidReceiverErr{err: err})
	}

	// to do update pix info

	pix := entity.NewPix(dto.PixType, dto.PixKey)

	if err := pix.Validate(); err != nil {
		return returnError(&CreatingPixErr{err: err})
	}

	if err != nil {
		return returnError(&saveReceiverErr{err: err})
	}

	_, err = u.Repo.AddPix(receiver.Id, pix)
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
