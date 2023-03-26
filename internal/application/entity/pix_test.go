package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type pixTestCase struct {
	description string
	t           string
	key         string
	expected    *Pix
	err         error
}

var p Pix = Pix{
	Type: "CPF",
	Key:  "428.639.342-95",
}

var pixTestCases = []pixTestCase{
	{
		description: "given valid data return a valid pix",
		t:           p.Type,
		key:         p.Key,
		expected:    &p,
		err:         nil,
	},
	{
		description: "given invalid type return an error",
		t:           "tipo inválido",
		key:         p.Key,
		expected:    nil,
		err:         ErrInvalidType,
	},
	{
		description: "given empty type return an error",
		t:           "",
		key:         p.Key,
		expected:    nil,
		err:         ErrInvalidType,
	},
	{
		description: "given invalid key for cpf type return an error",
		t:           "CPF",
		key:         "219.334.250-",
		expected:    nil,
		err:         ErrInvalidKey,
	},
	{
		description: "given invalid key for cnpj type return an error",
		t:           "CNPJ",
		key:         "05.029.616/0001-0",
		expected:    nil,
		err:         ErrInvalidKey,
	},
	{
		description: "given invalid key for email type return an error",
		t:           "EMAIL",
		key:         "@yhaoo.com.br",
		expected:    nil,
		err:         ErrInvalidKey,
	},
	{
		description: "given invalid key for phone type return an error",
		t:           "TELEFONE",
		key:         "12307382917",
		expected:    nil,
		err:         ErrInvalidKey,
	},
	{
		description: "given invalid key for random type return an error",
		t:           "CHAVE_ALEATORIA",
		key:         "8751-8692",
		expected:    nil,
		err:         ErrInvalidKey,
	},
}

func TestCreateAndValidateNewPix(t *testing.T) {
	for _, tc := range pixTestCases {
		t.Run(tc.description, func(t *testing.T) {

			pix := NewPix(tc.t, tc.key)
			err := pix.Validate()

			if tc.err != nil {
				assert.ErrorIs(t, err, tc.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, pix)
			}
		})
	}
}
