package delete_receiver

import (
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
)

type Delete struct {
	Repo entity.Repository
}

func (d *Delete) Execute(dto *InputDto) error {
	if len(dto.ReceiversIds) == 0 {
		return ErrReceiversIdsAreRequired
	}

	if err := d.Repo.DeleteReceivers(dto.ReceiversIds); err != nil {
		return &deletingReceiverErr{err: err}
	}

	return nil
}
