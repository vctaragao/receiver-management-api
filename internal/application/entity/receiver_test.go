package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type receiverTestCase struct {
	description   string
	corporateName string
	cpfCnpj       string
	email         string
	status        string
	expected      *Receiver
	err           error
}

var r Receiver = Receiver{
	CorporateName: "Clarice Rayssa Tereza Assunção",
	CpfCnpj:       "428.639.342-95",
	Email:         "claricerayssaassuncao@muvacademia.com.br",
	Status:        "RASCUNHO",
}

var receiverTestCases = []receiverTestCase{
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
		err:           ErrInvalidCorporateName,
	},
	{
		description:   "given invalid coporate name return an error",
		corporateName: "a",
		cpfCnpj:       r.CpfCnpj,
		email:         r.Email,
		status:        r.Status,
		expected:      nil,
		err:           ErrInvalidCorporateName,
	},
	{
		description:   "given invalid cpf return an error",
		corporateName: r.CorporateName,
		cpfCnpj:       "219.334.250-4",
		email:         r.Email,
		status:        r.Status,
		expected:      nil,
		err:           ErrInvalidCpfCnpj,
	},
	{
		description:   "given invalid cnpj return an error",
		corporateName: r.CorporateName,
		cpfCnpj:       "15.663.178/0001-0",
		email:         r.Email,
		status:        r.Status,
		expected:      nil,
		err:           ErrInvalidCpfCnpj,
	},
	{
		description:   "given invalid email return an error",
		corporateName: r.CorporateName,
		cpfCnpj:       r.CpfCnpj,
		email:         "claricerayssaassuncaomuvacademia.com.br",
		status:        r.Status,
		expected:      nil,
		err:           ErrInvalidEmail,
	},
	{
		description:   "given invalid status return an error",
		corporateName: r.CorporateName,
		cpfCnpj:       r.CpfCnpj,
		email:         r.Email,
		status:        "Status inválido",
		expected:      nil,
		err:           ErrInvalidStatus,
	},
}

func TestCreateAndValidateNewReceiver(t *testing.T) {
	for _, tc := range receiverTestCases {
		t.Run(tc.description, func(t *testing.T) {

			receiver := NewReceiver(tc.corporateName, tc.cpfCnpj, tc.email, tc.status)
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
