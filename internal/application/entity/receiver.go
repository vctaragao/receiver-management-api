package entity

import (
	"errors"

	"github.com/vctaragao/receiver-management-api/internal/application/entity/helper"
)

const (
	STATUS_DRAFT     = "RASCUNHO"
	STATUS_VALID     = "VALIDADO"
	MAX_EMAIL_LENGTH = 250
)

var ErrInvalidCpf = errors.New("invalid cpf")
var ErrInvalidCnpj = errors.New("invalid cnpj")
var ErrInvalidEmail = errors.New("invalid email")
var ErrInvalidStatus = errors.New("invalid status")
var ErrInvalidCorporateName = errors.New("corporate name must be greater that 2 caracters")

type Receiver struct {
	Id            uint
	CorporateName string
	Cpf           string
	Cnpj          string
	Email         string
	Status        string
}

func NewReceiver(corporateName, cpf, cnpj, email, status string) *Receiver {
	return &Receiver{
		CorporateName: corporateName,
		Cpf:           cpf,
		Cnpj:          cnpj,
		Email:         email,
		Status:        status,
	}
}

func (r *Receiver) Validate() error {
	if len(r.CorporateName) < 2 {
		return ErrInvalidCorporateName
	}

	if r.Cpf != "" && !helper.IsValidCpf(r.Cpf) {
		return ErrInvalidCpf
	}

	if r.Cnpj != "" && !helper.IsValidCnpj(r.Cnpj) {
		return ErrInvalidCnpj
	}

	if r.Email != "" && !r.hasValidEmail() {
		return ErrInvalidEmail
	}

	if r.Status != "" && !r.hasValidStatus() {
		return ErrInvalidStatus
	}

	return nil
}

func (r *Receiver) hasValidEmail() bool {
	if !helper.IsValidEmail(r.Email) || len(r.Email) > MAX_EMAIL_LENGTH {
		return false
	}

	return true
}

func (r *Receiver) hasValidStatus() bool {
	for _, status := range GetValidReciverStatus() {
		if status == r.Status {
			return true
		}
	}

	return false
}

func GetValidReciverStatus() []string {
	return []string{STATUS_DRAFT, STATUS_VALID}
}

func (r *Receiver) Update(corporateName, cpf, cnpj, email string) error {
	if r.canUpdate(corporateName, cpf, cnpj) {
		return errors.New("cantUpdate")
	}

	if corporateName != "" {
		r.CorporateName = corporateName
	}

	if cpf != "" {
		r.Cpf = cpf
	}

	if cnpj != "" {
		r.Cnpj = cnpj
	}

	if email != "" {
		r.Email = email
	}

	return nil
}

func (r *Receiver) canUpdate(corporateName, cpf, cnpj string) bool {
	if r.Status == STATUS_DRAFT {
		return true
	}

	if corporateName != "" || cpf != "" || cnpj != "" {
		return false
	}

	return true
}
