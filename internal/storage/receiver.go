package storage

import (
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
	"github.com/vctaragao/receiver-management-api/internal/storage/schemas"
	"gorm.io/gorm"
)

func (p *Postgress) AddReceiver(r *entity.Receiver) (*entity.Receiver, error) {
	receiver := schemas.Receiver{
		CorporateName: r.CorporateName,
		CpfCnpj:       r.CpfCnpj,
		Email:         r.Email,
		Status:        r.Status,
	}

	result := p.Db.Create(&receiver)

	if result.Error != nil {
		return &entity.Receiver{}, ErrUnableToInsert
	}

	r.Id = receiver.ID

	return r, nil
}

func (postgress *Postgress) GetReceiverWithPix(receiverId uint) (*entity.Receiver, *entity.Pix, error) {
	var receiver schemas.Receiver
	err := postgress.Db.Model(&schemas.Receiver{}).Preload("Pix").First(&receiver).Error
	if err != nil {
		return &entity.Receiver{}, &entity.Pix{}, ErrUnableToFetch
	}

	r := &entity.Receiver{
		Id:            receiverId,
		CorporateName: receiver.CorporateName,
		CpfCnpj:       receiver.CpfCnpj,
		Email:         receiver.Email,
		Status:        receiver.Status,
	}

	p := &entity.Pix{
		Id:   receiver.Pix[0].ID,
		Type: receiver.Pix[0].Type,
		Key:  receiver.Pix[0].Key,
	}

	return r, p, nil
}

func (postgress *Postgress) UpdateReceiver(r *entity.Receiver) (*entity.Receiver, error) {
	receiver := &schemas.Receiver{
		CorporateName: r.CorporateName,
		CpfCnpj:       r.CpfCnpj,
		Email:         r.Email,
		Status:        r.Status,
		Model:         gorm.Model{ID: r.Id},
	}

	if err := postgress.Db.Save(receiver).Error; err != nil {
		return &entity.Receiver{}, ErrUnableToUpdate
	}

	return r, nil
}
