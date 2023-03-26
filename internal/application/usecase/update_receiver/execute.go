package update_receiver

import (
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
)

type Update struct {
	Repo entity.Repository
}

func (u *Update) Execute(dto *InputDto) (*OutputDto, error) {
	receiver, pix, err := u.Repo.GetReceiverWithPix(dto.ReceiverId)
	if err != nil {
		return returnError(&findingReceiverError{err: err})
	}
	if err := u.updateReciver(receiver, dto); err != nil {
		return returnError(err)
	}

	if receiver.IsInDraft() {
		if err := u.updatePix(pix, dto); err != nil {
			return returnError(err)
		}
	}

	out := &OutputDto{
		CorporateName: receiver.CorporateName,
		CpfCnpj:       receiver.CpfCnpj,
		Email:         receiver.Email,
		ReceiverId:    receiver.Id,
		PixType:       pix.Type,
		PixKey:        pix.Key,
	}

	return out, nil
}

func (u *Update) updateReciver(receiver *entity.Receiver, dto *InputDto) error {
	if err := receiver.Update(dto.CorporateName, dto.CpfCnpj, dto.Email); err != nil {
		return &UpdatingReceiverErr{err: err}
	}

	if err := receiver.Validate(); err != nil {
		return &invalidReceiverErr{err: err}
	}

	if _, err := u.Repo.UpdateReceiver(receiver); err != nil {
		return &UpdatingReceiverErr{err: err}
	}

	return nil
}

func (u *Update) updatePix(pix *entity.Pix, dto *InputDto) error {
	if dto.PixType == "" || dto.PixKey == "" {
		return nil
	}

	pix.Update(dto.PixType, dto.PixKey)

	if err := pix.Validate(); err != nil {
		return &invalidPixErr{err: err}
	}

	if _, err := u.Repo.UpdatePix(pix); err != nil {
		return &UpdatingPixErr{err: err}
	}

	return nil
}

func returnError(err error) (*OutputDto, error) {
	return &OutputDto{}, err
}
