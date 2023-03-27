package update_receiver

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
	"github.com/vctaragao/receiver-management-api/internal/test/mocks"
)

type testCase struct {
	description string
	receiver    entity.Receiver
	pix         entity.Pix
	input       *InputDto
	output      *OutputDto
	err         error
}

var receiverDraft entity.Receiver = entity.Receiver{
	Id:            1,
	CorporateName: "Lara Natália Ana Almeida",
	CpfCnpj:       "009.016.853-48",
	Email:         "laranataliaalmeida@chavao.com.br",
	Status:        entity.STATUS_DRAFT,
}

var pixDraft entity.Pix = entity.Pix{
	Id:   1,
	Type: entity.CPF,
	Key:  "009.016.853-48",
}

var receiverValid entity.Receiver = entity.Receiver{
	Id:            2,
	CorporateName: "Nicolas Levi Osvaldo da Rosa",
	CpfCnpj:       "727.340.197-87",
	Email:         "nicolas.levi.darosa@novaface.com.br",
	Status:        entity.STATUS_VALID,
}

var pixValid entity.Pix = entity.Pix{
	Id:   2,
	Type: entity.CPF,
	Key:  "727.340.197-87",
}

func TestUpdateReceiver(t *testing.T) {
	for _, tc := range getTestCases() {
		t.Run(tc.description, func(t *testing.T) {

			repo := setupMock(tc, tc.receiver, tc.pix)

			usecase := &Update{Repo: repo}

			outputDto, err := usecase.Execute(tc.input)

			if tc.err != nil {
				fmt.Println(err)
				assert.ErrorAs(t, err, &tc.err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.output, outputDto)
		})
	}
}

func setupMock(tc testCase, receiver_return entity.Receiver, pix_return entity.Pix) *mocks.MockRepo {
	repo := &mocks.MockRepo{}
	receiver_found, pix_found := setGetReceiverWithPixExpectation(receiver_return, pix_return, tc, repo)
	setUpdateReceiverExpectation(receiver_found, tc, repo)
	setUpdatePixExpectation(pix_found, tc, repo)

	return repo
}

func setGetReceiverWithPixExpectation(receiver_return entity.Receiver, pix_return entity.Pix, tc testCase, repo *mocks.MockRepo) (entity.Receiver, entity.Pix) {
	var err error

	if tc.err != nil && errors.As(tc.err, &fReceiverErr) {
		receiver_return = entity.Receiver{}
		pix_return = entity.Pix{}
		err = errors.New("unexpected error")
	}

	repo.On("GetReceiverWithPix", tc.input.ReceiverId).Return(&receiver_return, &pix_return, err)

	return receiver_return, pix_return
}

func setUpdateReceiverExpectation(receiver_found entity.Receiver, tc testCase, repo *mocks.MockRepo) {
	receiver_found.Update(tc.input.CorporateName, tc.input.CpfCnpj, tc.input.Email)

	receiver_return := receiver_found
	var err error

	if tc.err != nil && errors.As(tc.err, &uReceiverErr) {
		receiver_return = entity.Receiver{}
		err = errors.New("unexpected error")
	}

	repo.On("UpdateReceiver", &receiver_found, tc.input.CorporateName, tc.input.CpfCnpj, tc.input.Email).Return(&receiver_return, err)
}

func setUpdatePixExpectation(pix_found entity.Pix, tc testCase, repo *mocks.MockRepo) {
	pix_found.Update(tc.input.PixType, tc.input.PixKey)

	pix_return := pix_found
	var err error

	if tc.err != nil && errors.As(tc.err, &uPixErr) {
		pix_return = entity.Pix{}
		err = errors.New("unexpected error")
	}

	repo.On("UpdatePix", &pix_found).Return(&pix_return, err)
}

func getTestCases() []testCase {
	return []testCase{
		{
			description: "with a draft receiver and valid data update receiver and pix",
			receiver:    receiverDraft,
			pix:         pixDraft,
			input: &InputDto{
				ReceiverId:    1,
				CorporateName: "Mariana e Emilly Adega ME",
				CpfCnpj:       "76.560.155/0001-44",
				Email:         "qualidade@marianaeemillyadegame.com.br",
				PixType:       "CNPJ",
				PixKey:        "76.560.155/0001-44",
			},
			output: &OutputDto{
				ReceiverId:    1,
				CorporateName: "Mariana e Emilly Adega ME",
				CpfCnpj:       "76.560.155/0001-44",
				Email:         "qualidade@marianaeemillyadegame.com.br",
				PixType:       "CNPJ",
				Status:        "RASCUNHO",
				PixKey:        "76.560.155/0001-44",
			},
			err: nil,
		},
		{
			description: "with a validated receiver and email than update receiver email only",
			receiver:    receiverValid,
			pix:         pixValid,
			input: &InputDto{
				ReceiverId:    1,
				CorporateName: "",
				CpfCnpj:       "",
				Email:         "faleconosco@benjamineruancontabilme.com.br",
				PixType:       "TELEFONE",
				PixKey:        "(11) 2764-3535",
			},
			output: &OutputDto{
				ReceiverId:    receiverValid.Id,
				CorporateName: receiverValid.CorporateName,
				CpfCnpj:       receiverValid.CpfCnpj,
				Email:         "faleconosco@benjamineruancontabilme.com.br",
				PixType:       pixValid.Type,
				Status:        "VALIDADO",
				PixKey:        pixValid.Key,
			},
			err: nil,
		},
		{
			description: "with a draft receiver and only pix info sended than update pix",
			receiver:    receiverDraft,
			pix:         pixDraft,
			input: &InputDto{
				ReceiverId:    1,
				CorporateName: "",
				CpfCnpj:       "",
				Email:         "",
				PixType:       "EMAIL",
				PixKey:        "faleconosco@benjamineruancontabilme.com.br",
			},
			output: &OutputDto{
				ReceiverId:    receiverDraft.Id,
				CorporateName: receiverDraft.CorporateName,
				CpfCnpj:       receiverDraft.CpfCnpj,
				Email:         receiverDraft.Email,
				Status:        "RASCUNHO",
				PixType:       "EMAIL",
				PixKey:        "faleconosco@benjamineruancontabilme.com.br",
			},
			err: nil,
		},
		{
			description: "with an invalid receiver id return not found error",
			receiver:    receiverDraft,
			pix:         pixDraft,
			input: &InputDto{
				ReceiverId:    10,
				CorporateName: "",
				CpfCnpj:       "",
				Email:         "",
				PixType:       "",
				PixKey:        "",
			},
			output: &OutputDto{},
			err:    fReceiverErr,
		},
		{
			description: "with receiver is status valid and try to update corporate name return update receiver error",
			receiver:    receiverValid,
			pix:         pixValid,
			input: &InputDto{
				ReceiverId:    1,
				CorporateName: "Malu Catarina Milena",
				CpfCnpj:       "",
				Email:         "",
				PixType:       "",
				PixKey:        "",
			},
			output: &OutputDto{},
			err:    uReceiverErr,
		},
		{
			description: "with receiver is status draft and try to update with invalid data return invalid receiver error",
			receiver:    receiverDraft,
			pix:         pixDraft,
			input: &InputDto{
				ReceiverId:    1,
				CorporateName: "Malu Catarina Milena",
				CpfCnpj:       "5484651864",
				Email:         "",
				PixType:       "",
				PixKey:        "",
			},
			output: &OutputDto{},
			err:    iReceiverErr,
		},
		{
			description: "with receiver is status draft and try to update with valid data but repo returns an error than return error",
			receiver:    receiverDraft,
			pix:         pixDraft,
			input: &InputDto{
				ReceiverId:    1,
				CorporateName: "Malu Catarina Milena",
				CpfCnpj:       "",
				Email:         "",
				PixType:       "",
				PixKey:        "",
			},
			output: &OutputDto{},
			err:    uReceiverErr,
		},
		{
			description: "with receiver is status draft and try to update with invalid pix data return error",
			receiver:    receiverDraft,
			pix:         pixDraft,
			input: &InputDto{
				ReceiverId:    1,
				CorporateName: "Malu Catarina Milena",
				CpfCnpj:       "",
				Email:         "",
				PixType:       "TELEFONE",
				PixKey:        "8945468",
			},
			output: &OutputDto{},
			err:    iPixErr,
		},
		{
			description: "with receiver is status draft and try to update with valid pix data but repo returns and error than return error",
			receiver:    receiverDraft,
			pix:         pixDraft,
			input: &InputDto{
				ReceiverId:    1,
				CorporateName: "Malu Catarina Milena",
				CpfCnpj:       "",
				Email:         "",
				PixType:       "EMAIL",
				PixKey:        "faleconosco@benjamineruancontabilme.com.br",
			},
			output: &OutputDto{},
			err:    uPixErr,
		},
	}
}
