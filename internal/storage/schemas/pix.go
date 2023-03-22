package schemas

import "gorm.io/gorm"

type Pix struct {
	Type       string
	Key        string
	ReceiverId uint
	gorm.Model
}
