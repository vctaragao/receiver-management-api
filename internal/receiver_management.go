package internal

import (
	"github.com/vctaragao/receiver-management-api/internal/application"
)

type ReceiverManagement struct {
	Repo application.Repository
}

func NewReceiverManagement(repo application.Repository) *ReceiverManagement {
	return &ReceiverManagement{
		Repo: repo,
	}
}
