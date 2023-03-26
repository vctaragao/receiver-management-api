package receiver

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
)

type updateTestCase struct {
	description   string
	receiver      entity.Receiver
	corporateName string
	cpf           string
	cnpj          string
	email         string
	expected      entity.Receiver
	err           error
}

func TestUpdateReceiver(t *testing.T) {
	for _, tc := range getUpdateTestCasesSuccess() {
		t.Run(tc.description, func(t *testing.T) {
			err := tc.receiver.Update(tc.corporateName, tc.cpf, tc.cnpj, tc.email)

			if tc.err != nil {
				assert.ErrorIs(t, err, tc.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, tc.receiver)
			}
		})
	}
}

func getUpdateTestCasesSuccess() []updateTestCase {
	r := entity.Receiver{
		CorporateName: "Clarice Rayssa Tereza Assunção",
		Cpf:           "428.639.342-95",
		Cnpj:          "85.980.829/0001-50",
		Email:         "claricerayssaassuncao@muvacademia.com.br",
		Status:        "RASCUNHO",
	}

	return []updateTestCase{
		{
			description:   "given valid data update receiver",
			receiver:      r,
			corporateName: "Valid name",
			cpf:           r.Cpf,
			cnpj:          r.Cnpj,
			email:         r.Email,
			expected: entity.Receiver{
				CorporateName: "Valid name",
				Cpf:           r.Cpf,
				Cnpj:          r.Cnpj,
				Email:         r.Email,
				Status:        r.Status,
			},
			err: nil,
		},
	}
}
