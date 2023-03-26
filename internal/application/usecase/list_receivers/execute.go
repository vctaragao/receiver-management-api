package list_receivers

import (
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
)

type List struct {
	Repo entity.Repository
}

func (l *List) Execute(dto *InputDto) (*OutputDto, error) {
	var receivers []entity.Receiver
	var err error

	if dto.SearchParam == "" {
		receivers, err = l.Repo.FindReceivers(dto.Page)
	} else {
		receivers, err = l.Repo.FindReceiversBy(dto.SearchParam, dto.Page)
	}

	if err != nil {
		return returnError(err)
	}

	return &OutputDto{Receivers: receivers}, nil
}

func returnError(err error) (*OutputDto, error) {
	return &OutputDto{}, err
}
