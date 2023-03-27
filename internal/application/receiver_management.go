package application

import (
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
	"github.com/vctaragao/receiver-management-api/internal/application/usecase/create_receiver"
	"github.com/vctaragao/receiver-management-api/internal/application/usecase/delete_receiver"
	"github.com/vctaragao/receiver-management-api/internal/application/usecase/list_receivers"
	"github.com/vctaragao/receiver-management-api/internal/application/usecase/update_receiver"
)

type ReceiverManagement struct {
	Repo *entity.Repository
}

func NewReceiverManagement(repo entity.Repository) *ReceiverManagement {
	return &ReceiverManagement{
		Repo: &repo,
	}
}

func (rm *ReceiverManagement) Create(corporateName, cpfCnpj, email, pixType, pixKey string) (*create_receiver.OutputDto, error) {
	dto := &create_receiver.InputDto{
		CorporateName: corporateName,
		CpfCnpj:       cpfCnpj,
		Email:         email,
		PixType:       pixType,
		PixKey:        pixKey,
	}

	usecase := &create_receiver.Create{Repo: *rm.Repo}

	return usecase.Execute(dto)
}

func (rm *ReceiverManagement) IsCreateBussinesLogicError(err error) bool {
	return create_receiver.IsBussinesLogicError(err)
}

func (rm *ReceiverManagement) Update(receiverId uint, corporateName, cpfCnpj, email, pixType, pixKey string) (*update_receiver.OutputDto, error) {
	dto := &update_receiver.InputDto{
		ReceiverId:    receiverId,
		CorporateName: corporateName,
		CpfCnpj:       cpfCnpj,
		Email:         email,
		PixType:       pixType,
		PixKey:        pixKey,
	}

	usecase := &update_receiver.Update{Repo: *rm.Repo}

	return usecase.Execute(dto)
}

func (rm *ReceiverManagement) IsUpdateBussinesLogicError(err error) bool {
	return update_receiver.IsBussinesLogicError(err)
}

func (rm *ReceiverManagement) List(searchParam string, page int) (*list_receivers.OutputDto, error) {
	dto := &list_receivers.InputDto{
		SearchParam: searchParam,
		Page:        page,
	}

	usecase := &list_receivers.List{Repo: *rm.Repo}

	return usecase.Execute(dto)
}

func (rm *ReceiverManagement) Delete(receiversIds []uint) error {
	dto := &delete_receiver.InputDto{
		ReceiversIds: receiversIds,
	}

	usecase := &delete_receiver.Delete{Repo: *rm.Repo}

	return usecase.Execute(dto)
}
func (rm *ReceiverManagement) IsDeleteBussinesLogicError(err error) bool {
	return delete_receiver.IsBusinessLogicError(err)
}
