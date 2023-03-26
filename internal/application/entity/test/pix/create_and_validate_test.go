package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
)

type createTestCase struct {
	description string
	t           string
	key         string
	expected    *entity.Pix
	err         error
}

func TestCreateAndValidateNewPix(t *testing.T) {
	for _, tc := range getCreateAndValidateTestCases() {
		t.Run(tc.description, func(t *testing.T) {

			pix := entity.NewPix(tc.t, tc.key)
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

func getCreateAndValidateTestCases() []createTestCase {

	p := entity.Pix{
		Type: "CPF",
		Key:  "428.639.342-95",
	}

	return []createTestCase{
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
			err:         entity.ErrInvalidType,
		},
		{
			description: "given empty type return an error",
			t:           "",
			key:         p.Key,
			expected:    nil,
			err:         entity.ErrInvalidType,
		},
		{
			description: "given invalid key for cpf type return an error",
			t:           "CPF",
			key:         "219.334.250-",
			expected:    nil,
			err:         entity.ErrInvalidKey,
		},
		{
			description: "given invalid key for cnpj type return an error",
			t:           "CNPJ",
			key:         "05.029.616/0001-0",
			expected:    nil,
			err:         entity.ErrInvalidKey,
		},
		{
			description: "given invalid key for email type return an error",
			t:           "EMAIL",
			key:         "@yhaoo.com.br",
			expected:    nil,
			err:         entity.ErrInvalidKey,
		},
		{
			description: "given invalid key for phone type return an error",
			t:           "TELEFONE",
			key:         "12307382917",
			expected:    nil,
			err:         entity.ErrInvalidKey,
		},
		{
			description: "given invalid key for random type return an error",
			t:           "CHAVE_ALEATORIA",
			key:         "8751-8692",
			expected:    nil,
			err:         entity.ErrInvalidKey,
		},
	}
}
