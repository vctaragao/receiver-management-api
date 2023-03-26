package storage

import (
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
	"github.com/vctaragao/receiver-management-api/internal/storage/schemas"
	"gorm.io/gorm"
)

const PAGE_SIZE = 10

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

func (postgress *Postgress) FindReceivers(page int) ([]entity.Receiver, error) {
	var records []schemas.Receiver
	err := postgress.Db.
		Preload("Pix").
		Limit(PAGE_SIZE).
		Offset(PAGE_SIZE * (page - 1)).
		Find(&records).Error

	if err != nil {
		return []entity.Receiver{}, ErrUnableToFind
	}

	return createReceiversSlice(records), nil
}

func (postgress *Postgress) FindReceiversBy(searchParam string, page int) ([]entity.Receiver, error) {
	var records []schemas.Receiver
	err := postgress.Db.
		Preload("Pix").
		Where("corporate_name = ? OR status = ? OR EXISTS (SELECT 1 FROM pixes WHERE receiver_id = receivers.id AND (type = ? OR key = ?))", searchParam, searchParam, searchParam, searchParam).
		Limit(PAGE_SIZE).
		Offset(PAGE_SIZE * (page - 1)).
		Find(&records).Error

	if err != nil {
		return []entity.Receiver{}, ErrUnableToFind
	}

	return createReceiversSlice(records), nil
}

func createReceiversSlice(records []schemas.Receiver) []entity.Receiver {
	var receivers []entity.Receiver
	for _, record := range records {
		receiver := *entity.NewReceiver(record.CorporateName, record.CpfCnpj, record.Email, record.Status)
		receiver.Id = record.ID

		receiver.SetPix(&entity.Pix{
			Id:   record.Pix[0].ID,
			Type: record.Pix[0].Type,
			Key:  record.Pix[0].Key,
		})

		receivers = append(receivers, receiver)
	}

	return receivers
}
