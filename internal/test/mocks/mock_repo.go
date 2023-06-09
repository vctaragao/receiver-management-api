package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
)

type MockRepo struct {
	mock.Mock
}

func (mr *MockRepo) AddReceiver(r *entity.Receiver) (*entity.Receiver, error) {
	args := mr.Called(r)
	return args.Get(0).(*entity.Receiver), args.Error(1)
}

func (mr *MockRepo) UpdateReceiver(r *entity.Receiver, corporateName, cpfCnpj, email string) (*entity.Receiver, error) {
	args := mr.Called(r, corporateName, cpfCnpj, email)
	return args.Get(0).(*entity.Receiver), args.Error(1)
}

func (mr *MockRepo) GetReceiverWithPix(receiverId uint) (*entity.Receiver, *entity.Pix, error) {
	args := mr.Called(receiverId)
	return args.Get(0).(*entity.Receiver), args.Get(1).(*entity.Pix), args.Error(2)
}

func (mr *MockRepo) AddPix(receiverId uint, p *entity.Pix) (*entity.Pix, error) {
	args := mr.Called(receiverId, p)
	return args.Get(0).(*entity.Pix), args.Error(1)
}

func (mr *MockRepo) UpdatePix(p *entity.Pix) (*entity.Pix, error) {
	args := mr.Called(p)
	return args.Get(0).(*entity.Pix), args.Error(1)
}

func (mr *MockRepo) FindReceiversBy(searchParam string, page int) ([]entity.Receiver, int, error) {
	args := mr.Called(searchParam, page)
	return args.Get(0).([]entity.Receiver), args.Int(1), args.Error(2)
}

func (mr *MockRepo) FindReceivers(page int) ([]entity.Receiver, int, error) {
	args := mr.Called(page)
	return args.Get(0).([]entity.Receiver), args.Int(1), args.Error(2)
}

func (mr *MockRepo) DeleteReceivers(receversIds []uint) error {
	args := mr.Called(receversIds)
	return args.Error(0)
}
