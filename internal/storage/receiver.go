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
	receiver := schemas.Receiver{
		Model: gorm.Model{ID: receiverId},
	}

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

func (postgress *Postgress) UpdateReceiver(r *entity.Receiver, corporateName, cpfCnpj, email string) (*entity.Receiver, error) {
	receiver := &schemas.Receiver{
		CorporateName: corporateName,
		CpfCnpj:       cpfCnpj,
		Email:         email,
	}

	err := postgress.Db.Model(&schemas.Receiver{}).Where("id = ?", r.Id).Updates(receiver).Error

	if err != nil {
		return &entity.Receiver{}, ErrUnableToUpdate
	}

	return r, nil
}

func (postgress *Postgress) FindReceivers(page int) ([]entity.Receiver, int, error) {
	var records []schemas.Receiver
	err := postgress.Db.
		Preload("Pix").
		Limit(PAGE_SIZE).
		Offset(PAGE_SIZE * (page - 1)).
		Find(&records).Error

	if err != nil {
		return []entity.Receiver{}, 0, ErrUnableToFind
	}

	var quantity int64
	err = postgress.Db.
		Model(&schemas.Receiver{}).
		Preload("Pix").
		Count(&quantity).Error

	if err != nil {
		return []entity.Receiver{}, 0, ErrUnableToFind
	}

	return createReceiversSlice(records), int(quantity), nil
}

func (postgress *Postgress) FindReceiversBy(searchParam string, page int) ([]entity.Receiver, int, error) {
	var records []schemas.Receiver
	err := postgress.Db.
		Preload("Pix").
		Where("corporate_name = ? OR status = ? OR EXISTS (SELECT 1 FROM pixes WHERE receiver_id = receivers.id AND (type = ? OR key = ?))", searchParam, searchParam, searchParam, searchParam).
		Limit(PAGE_SIZE).
		Offset(PAGE_SIZE * (page - 1)).
		Find(&records).Error

	if err != nil {
		return []entity.Receiver{}, 0, ErrUnableToFind
	}

	var quantity int64
	err = postgress.Db.
		Model(&schemas.Receiver{}).
		Preload("Pix").
		Where("corporate_name = ? OR status = ? OR EXISTS (SELECT 1 FROM pixes WHERE receiver_id = receivers.id AND (type = ? OR key = ?))", searchParam, searchParam, searchParam, searchParam).
		Count(&quantity).Error

	if err != nil {
		return []entity.Receiver{}, 0, ErrUnableToFind
	}

	return createReceiversSlice(records), int(quantity), nil
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

func (postgress *Postgress) DeleteReceivers(receiversIds []uint) error {
	var receiversToDelete []schemas.Receiver
	for _, id := range receiversIds {
		receiver := schemas.Receiver{Model: gorm.Model{ID: id}}
		receiversToDelete = append(receiversToDelete, receiver)
	}

	if err := postgress.Db.Select("Pix").Delete(&receiversToDelete).Error; err != nil {
		return ErrUnableToDelete
	}

	return nil
}
