package entity

import (
	"errors"

	"github.com/vctaragao/receiver-management-api/internal/application/entity/helper"
)

const (
	CPF            = "CPF"
	CNPJ           = "CNPJ"
	EMAIL          = "EMAIL"
	PHONE          = "TELEFONE"
	RANDOM_KEY     = "CHAVE_ALEATORIA"
	MAX_LENGTH_KEY = 140
)

type Pix struct {
	Type string
	Key  string
}

func NewPix(t, key string) (*Pix, error) {
	p := &Pix{
		Type: t,
		Key:  key,
	}

	if err := p.isValid(); err != nil {
		return &Pix{}, err
	}

	return p, nil
}

func (p *Pix) isValid() error {
	if !p.isValidType() {
		return errors.New("invalid type")
	}

	if p.Key == "" || len(p.Key) > MAX_LENGTH_KEY || !p.isValidKey() {
		return errors.New("invalid key")
	}

	return nil
}

func (p *Pix) isValidType() bool {
	for _, t := range getValidTypes() {
		if t == p.Type {
			return true
		}
	}

	return false
}

func (p *Pix) isValidKey() bool {
	result := false

	switch p.Type {
	case CPF:
		result = helper.IsValidCnpj(p.Key)
	case CNPJ:
		result = helper.IsValidCnpj(p.Key)
	case EMAIL:
		result = helper.IsValidEmail(p.Key)
	case PHONE:
		result = helper.IsValidPhone(p.Key)
	case RANDOM_KEY:
		result = p.isValidRandomKey()
	}

	return result
}

func (p *Pix) isValidRandomKey() bool {
	return helper.MatchPattern(`/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/i`, p.Key)
}

func getValidTypes() []string {
	return []string{CPF, CNPJ, EMAIL, PHONE, RANDOM_KEY}
}
