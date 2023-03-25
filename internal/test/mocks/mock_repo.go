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

func (mr *MockRepo) AddPix(receiverId uint, p *entity.Pix) (*entity.Pix, error) {
	args := mr.Called(receiverId, p)
	return args.Get(0).(*entity.Pix), args.Error(1)
}