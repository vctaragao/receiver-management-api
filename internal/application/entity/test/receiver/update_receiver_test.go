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
	receiver := entity.Receiver{
		CorporateName: "Clarice Rayssa Tereza Assunção",
		CpfCnpj:       "428.639.342-95",
		Email:         "claricerayssaassuncao@muvacademia.com.br",
		Status:        entity.STATUS_DRAFT,
	}

	validatedReceiver := entity.Receiver{
		CorporateName: "Leonardo e Geraldo Pães e Doces Ltda",
		CpfCnpj:       "32.651.546/0001-07",
		Email:         "representantes@leonardoegeraldopaesedocesltda.com.br",
		Status:        entity.STATUS_VALID,
	}

	return []updateTestCase{
		{
			description:   "given valid data update receiver",
			receiver:      receiver,
			corporateName: "Bryan e Sophie Ferragens Ltda",
			cpfCnpj:       "13.992.684/0001-05",
			email:         "estoque@bryanesophieferragensltda.com.br",
			expected: entity.Receiver{
				CorporateName: "Bryan e Sophie Ferragens Ltda",
				CpfCnpj:       "13.992.684/0001-05",
				Email:         "estoque@bryanesophieferragensltda.com.br",
				Status:        receiver.Status,
			},
			err: nil,
		},
		{
			description:   "given receiver in validated status update email",
			receiver:      validatedReceiver,
			corporateName: "",
			cpfCnpj:       "",
			email:         "marketing@tiagoeflaviamarcenariame.com.br",
			expected: entity.Receiver{
				CorporateName: validatedReceiver.CorporateName,
				CpfCnpj:       validatedReceiver.CpfCnpj,
				Email:         "marketing@tiagoeflaviamarcenariame.com.br",
				Status:        validatedReceiver.Status,
			},
			err: nil,
		},
		{
			description:   "given receiver in validated status dont update when try to update other fields than email",
			receiver:      validatedReceiver,
			corporateName: "Kevin e Isaac Consultoria Financeira Ltda",
			cpfCnpj:       "06.699.980/0001-49",
			email:         "marketing@tiagoeflaviamarcenariame.com.br",
			expected:      validatedReceiver,
			err:           entity.ErrCanOlyUpdateEmailOnValidatedReceiver,
		},
	}
}
