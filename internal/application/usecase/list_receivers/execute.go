package list_receivers

import (
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
)

type List struct {
	Repo entity.Repository
}

func (l *List) Execute(dto *InputDto) (*OutputDto, error) {
	var receivers []entity.Receiver
	var total int
	var err error

	if dto.SearchParam == "" {
		receivers, total, err = l.Repo.FindReceivers(dto.Page)
	} else {
		receivers, total, err = l.Repo.FindReceiversBy(dto.SearchParam, dto.Page)
	}

	if err != nil {
		return &OutputDto{}, &findingReceiverErr{err: err}
	}

	return &OutputDto{Receivers: receivers, Total: total}, nil
}
