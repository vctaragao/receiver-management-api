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

type Receiver struct {
	Id            uint
	CorporateName string
	Cpf           string
	Cnpj          string
	Email         string
	Status        string
	Pix           Pix
}

func NewReceiver(corporateName, cpf, cnpj, email, status string) (*Receiver, error) {
	r := &Receiver{
		CorporateName: corporateName,
		Cpf:           cpf,
		Cnpj:          cnpj,
		Email:         email,
		Status:        status,
	}

	if err := r.isValid(); err != nil {
		return &Receiver{}, err
	}

	return r, nil
}

func (r *Receiver) isValid() error {
	if r.CorporateName == "" {
		return errors.New("invalid corporate name")
	}

	if r.Cpf != "" && r.Cnpj != "" {
		return errors.New("receiver can only be PF or PJ")
	}

	if r.Cpf != "" && !helper.IsValidCpf(r.Cpf) {
		return errors.New("invalid cpf")
	}

	if r.Cnpj != "" && !helper.IsValidCnpj(r.Cnpj) {
		return errors.New("invalid cnpj")
	}

	if r.Email != "" && !r.isValidEmail() {
		return errors.New("invalid email")
	}

	if r.Status != "" && !r.isValidStatus() {
		return errors.New("invalid status")
	}

	return nil
}

func (r *Receiver) isValidEmail() bool {
	if !helper.IsValidEmail(r.Email) || len(r.Email) > MAX_EMAIL_LENGTH {
		return false
	}

	return true
}

func (r *Receiver) isValidStatus() bool {
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
