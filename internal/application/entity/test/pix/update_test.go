package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
)

type updateTestCase struct {
	description string
	pix         entity.Pix
	t           string
	key         string
	expected    entity.Pix
	err         error
}

func TestUpdatePix(t *testing.T) {
	for _, tc := range getUpdateTestCases() {
		t.Run(tc.description, func(t *testing.T) {

			tc.pix.Update(tc.t, tc.key)
			assert.Equal(t, tc.expected, tc.pix)
		})
	}
}

func getUpdateTestCases() []updateTestCase {
	p := entity.Pix{
		Type: "CPF",
		Key:  "428.639.342-95",
	}

	return []updateTestCase{
		{
			description: "given valid data return update pix",
			pix:         p,
			t:           entity.EMAIL,
			key:         "suporte@esteretiagodocessalgadosltda.com.br",
			expected: entity.Pix{
				Type: "EMAIL",
				Key:  "suporte@esteretiagodocessalgadosltda.com.br",
			},
			err: nil,
		},
	}
}
