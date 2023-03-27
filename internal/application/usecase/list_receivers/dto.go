package list_receivers

import "github.com/vctaragao/receiver-management-api/internal/application/entity"

type InputDto struct {
	SearchParam string
	Page        int
}

type OutputDto struct {
	Receivers []entity.Receiver
	Total     int
}
