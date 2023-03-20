package storage

import (
	"gorm.io/gorm"
)

type DbAdapter struct {
	Db *gorm.DB
}
