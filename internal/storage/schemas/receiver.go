package schemas

import "gorm.io/gorm"

type Receiver struct {
	RazaoSocial string
	Cpf         string
	Cnpj        string
	Email       string
	Status      string
	Pix         []Pix
	gorm.Model
}
