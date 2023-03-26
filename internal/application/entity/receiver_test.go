package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type receiverTestCase struct {
	description   string
	corporateName string
	cpf           string
	cnpj          string
	email         string
	status        string
	expected      *Receiver
	err           error
}

var r Receiver = Receiver{
	CorporateName: "Clarice Rayssa Tereza Assunção",
	Cpf:           "428.639.342-95",
	Cnpj:          "85.980.829/0001-50",
	Email:         "claricerayssaassuncao@muvacademia.com.br",
	Status:        "RASCUNHO",
}

var receiverTestCases = []receiverTestCase{
	{
		description:   "given valid data return a valid receiver",
		corporateName: r.CorporateName,
		cpf:           r.Cpf,
		cnpj:          r.Cnpj,
		email:         r.Email,
		status:        r.Status,
		expected:      &r,
		err:           nil,
	},
	{
		description:   "given invalid coporate name return an error",
		corporateName: "",
		cpf:           r.Cpf,
		cnpj:          r.Cnpj,
		email:         r.Email,
		status:        r.Status,
		expected:      nil,
		err:           ErrInvalidCorporateName,
	},
	{
		description:   "given invalid coporate name return an error",
		corporateName: "a",
		cpf:           r.Cpf,
		cnpj:          r.Cnpj,
		email:         r.Email,
		status:        r.Status,
		expected:      nil,
		err:           ErrInvalidCorporateName,
	},
	{
		description:   "given invalid cpf return an error",
		corporateName: r.CorporateName,
		cpf:           "219.334.250-4",
		cnpj:          r.Cnpj,
		email:         r.Email,
		status:        r.Status,
		expected:      nil,
		err:           ErrInvalidCpf,
	},
	{
		description:   "given invalid cnpj return an error",
		corporateName: r.CorporateName,
		cpf:           r.Cpf,
		cnpj:          "15.663.178/0001-0",
		email:         r.Email,
		status:        r.Status,
		expected:      nil,
		err:           ErrInvalidCnpj,
	},
	{
		description:   "given invalid email return an error",
		corporateName: r.CorporateName,
		cpf:           r.Cpf,
		cnpj:          r.Cnpj,
		email:         "claricerayssaassuncaomuvacademia.com.br",
		status:        r.Status,
		expected:      nil,
		err:           ErrInvalidEmail,
	},
	{
		description:   "given invalid status return an error",
		corporateName: r.CorporateName,
		cpf:           r.Cpf,
		cnpj:          r.Cnpj,
		email:         r.Email,
		status:        "Status inválido",
		expected:      nil,
		err:           ErrInvalidStatus,
	},
}

func TestCreateAndValidateNewReceiver(t *testing.T) {
	for _, tc := range receiverTestCases {
		t.Run(tc.description, func(t *testing.T) {

			receiver := NewReceiver(tc.corporateName, tc.cpf, tc.cnpj, tc.email, tc.status)
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
