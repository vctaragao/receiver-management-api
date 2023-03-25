package schemas

import "gorm.io/gorm"

type Receiver struct {
	CorporateName string
	Cpf           string
	Cnpj          string
	Email         string
	Status        string
	Pix           []Pix
	gorm.Model
}
