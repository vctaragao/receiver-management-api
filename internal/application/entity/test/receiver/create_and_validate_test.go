package receiver

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
)

type createTestCase struct {
	description   string
	corporateName string
	cpfCnpj       string
	email         string
	status        string
	expected      *entity.Receiver
	err           error
}

func TestCreateAndValidateNewReceiver(t *testing.T) {
	for _, tc := range getCreateAndValidateTestCases() {
		t.Run(tc.description, func(t *testing.T) {

			receiver := entity.NewReceiver(tc.corporateName, tc.cpfCnpj, tc.email, tc.status)
			err := receiver.Validate()

			if tc.err != nil {
				assert.ErrorIs(t, err, tc.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, receiver)
			}
		})
	}
}

func getCreateAndValidateTestCases() []createTestCase {
	r := entity.Receiver{
		CorporateName: "Clarice Rayssa Tereza Assunção",
		CpfCnpj:       "428.639.342-95",
		Email:         "claricerayssaassuncao@muvacademia.com.br",
		Status:        "RASCUNHO",
	}

	return []createTestCase{
		{
			description:   "given valid data return a valid receiver",
			corporateName: r.CorporateName,
			cpfCnpj:       r.CpfCnpj,
			email:         r.Email,
			status:        r.Status,
			expected:      &r,
			err:           nil,
		},
		{
			description:   "given invalid coporate name return an error",
			corporateName: "",
			cpfCnpj:       r.CpfCnpj,
			email:         r.Email,
			status:        r.Status,
			expected:      nil,
			err:           entity.ErrInvalidCorporateName,
		},
		{
			description:   "given invalid coporate name return an error",
			corporateName: "a",
			cpfCnpj:       r.CpfCnpj,
			email:         r.Email,
			status:        r.Status,
			expected:      nil,
			err:           entity.ErrInvalidCorporateName,
		},
		{
			description:   "given invalid cpf return an error",
			corporateName: r.CorporateName,
			cpfCnpj:       "219.334.250-4",
			email:         r.Email,
			status:        r.Status,
			expected:      nil,
			err:           entity.ErrInvalidCpfCnpj,
		},
		{
			description:   "given invalid cnpj return an error",
			corporateName: r.CorporateName,
			cpfCnpj:       "15.663.178/0001-0",
			email:         r.Email,
			status:        r.Status,
			expected:      nil,
			err:           entity.ErrInvalidCpfCnpj,
		},
		{
			description:   "given invalid email return an error",
			corporateName: r.CorporateName,
			cpfCnpj:       r.CpfCnpj,
			email:         "claricerayssaassuncaomuvacademia.com.br",
			status:        r.Status,
			expected:      nil,
			err:           entity.ErrInvalidEmail,
		},
		{
			description:   "given invalid status return an error",
			corporateName: r.CorporateName,
			cpfCnpj:       r.CpfCnpj,
			email:         r.Email,
			status:        "Status inválido",
			expected:      nil,
			err:           entity.ErrInvalidStatus,
		},
	}
}
