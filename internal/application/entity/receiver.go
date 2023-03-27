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

var ErrInvalidEmail = errors.New("invalid email")
var ErrInvalidStatus = errors.New("invalid status")
var ErrInvalidCpfCnpj = errors.New("invalid cpf or cnpj")
var ErrInvalidCorporateName = errors.New("corporate name must be greater that 2 caracters")
var ErrCanOlyUpdateEmailOnValidatedReceiver = errors.New("can only update email on validated receiver")

type Receiver struct {
	Id            uint
	CorporateName string
	CpfCnpj       string
	Email         string
	Status        string
	Pix           *Pix
}

func NewReceiver(corporateName, cpfCnpj, email, status string) *Receiver {
	return &Receiver{
		CorporateName: corporateName,
		CpfCnpj:       cpfCnpj,
		Email:         email,
		Status:        status,
	}
}

func (r *Receiver) Validate() error {
	if len(r.CorporateName) < 2 {
		return ErrInvalidCorporateName
	}

	if r.CpfCnpj != "" && !helper.IsValidCpf(r.CpfCnpj) && !helper.IsValidCnpj(r.CpfCnpj) {
		return ErrInvalidCpfCnpj
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

func IsValidStatus(s string) bool {
	for _, status := range GetValidReciverStatus() {
		if status == s {
			return true
		}
	}

	return false
}

func (r *Receiver) Update(corporateName, cpfCnpj, email string) error {
	if !r.canUpdate(corporateName, cpfCnpj) {
		return ErrCanOlyUpdateEmailOnValidatedReceiver
	}

	if corporateName != "" {
		r.CorporateName = corporateName
	}

	if cpfCnpj != "" {
		r.CpfCnpj = cpfCnpj
	}

	if email != "" {
		r.Email = email
	}

	return nil
}

func (r *Receiver) canUpdate(corporateName, cpfCnpj string) bool {
	if r.Status == STATUS_DRAFT || (corporateName == "" && cpfCnpj == "") {
		return true
	}

	return false
}

func (r *Receiver) IsInDraft() bool {
	return r.Status == STATUS_DRAFT
}

func (r *Receiver) SetPix(p *Pix) {
	r.Pix = p
}
