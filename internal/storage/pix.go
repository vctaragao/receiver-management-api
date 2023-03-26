package storage

import (
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
	"github.com/vctaragao/receiver-management-api/internal/storage/schemas"
	"gorm.io/gorm"
)

func (postgres *Postgress) AddPix(receiverId uint, p *entity.Pix) (*entity.Pix, error) {
	pix := schemas.Pix{
		Type:       p.Type,
		Key:        p.Key,
		ReceiverId: receiverId,
	}

	if err := postgres.Db.Create(&pix).Error; err != nil {
		return &entity.Pix{}, ErrUnableToInsert
	}

	p.Id = pix.ID

	return p, nil
}

func (postgress *Postgress) UpdatePix(p *entity.Pix) (*entity.Pix, error) {
	pix := &schemas.Pix{
		Type:  p.Type,
		Key:   p.Key,
		Model: gorm.Model{ID: p.Id},
	}

	if err := postgress.Db.Save(pix).Error; err != nil {
		return &entity.Pix{}, ErrUnableToUpdate
	}

	return p, nil
}
