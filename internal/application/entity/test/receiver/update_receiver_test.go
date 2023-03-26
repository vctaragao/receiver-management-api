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
	cpfCnpj       string
	email         string
	expected      entity.Receiver
	err           error
}

func TestUpdateReceiver(t *testing.T) {
	for _, tc := range getUpdateTestCasesSuccess() {
		t.Run(tc.description, func(t *testing.T) {
			err := tc.receiver.Update(tc.corporateName, tc.cpfCnpj, tc.email)

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
		CpfCnpj:       "428.639.342-95",
		Email:         "claricerayssaassuncao@muvacademia.com.br",
		Status:        "RASCUNHO",
	}

	return []updateTestCase{
		{
			description:   "given valid data update receiver",
			receiver:      r,
			corporateName: "Valid name",
			cpfCnpj:       r.CpfCnpj,
			email:         r.Email,
			expected: entity.Receiver{
				CorporateName: "Valid name",
				CpfCnpj:       r.CpfCnpj,
				Email:         r.Email,
				Status:        r.Status,
			},
			err: nil,
		},
	}
}
