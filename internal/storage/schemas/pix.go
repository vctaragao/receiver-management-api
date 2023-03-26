package schemas

import "gorm.io/gorm"

type Pix struct {
	Type       string
	Key        string `gorm:"size:140"`
	ReceiverId uint
	gorm.Model
}
