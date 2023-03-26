package schemas

import "gorm.io/gorm"

type Receiver struct {
	CorporateName string
	CpfCnpj       string `gorm:"unique"`
	Email         string `gorm:"size:250;unique"`
	Status        string
	Pix           []Pix
	gorm.Model
}
