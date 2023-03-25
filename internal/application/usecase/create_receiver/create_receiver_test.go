package create_receiver

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
	"github.com/vctaragao/receiver-management-api/internal/test/mocks"
)

var savingPixErr *SavingPixErr
var creatingPixErr *CreatingPixErr
var saveReceiverErr *SaveReceiverErr
var creatingReceiverErr *CreatingReceiverErr

type testCase struct {
	description string
	input       *InputDto
	output      *OutputDto
	err         error
}

var inputDto = InputDto{
	CorporateName: "Lara Natália Ana Almeida",
	Cpf:           "009.016.853-48",
	Cnpj:          "",
	Email:         "laranataliaalmeida@chavao.com.br",
	Pix_type:      "CPF",
	Pix_key:       "009.016.853-48",
}

var testCases = []testCase{
	{
		description: "with valid data add receiver",
		input:       &inputDto,
		output: &OutputDto{
			Id:       1,
			InputDto: inputDto,
		},
		err: nil,
	},
	{
		description: "if an error accour on validating the receiver return error",
		input: &InputDto{
			CorporateName: inputDto.CorporateName,
			Cpf:           inputDto.Cpf,
			Cnpj:          inputDto.Cnpj,
			Email:         "laranataliaalmeidachavao.com.br",
			Pix_type:      inputDto.Pix_type,
			Pix_key:       inputDto.Pix_key,
		},
		output: &OutputDto{},
		err:    creatingReceiverErr,
	},
	{
		description: "if an error accour on validating the pix return error",
		input: &InputDto{
			CorporateName: inputDto.CorporateName,
			Cpf:           inputDto.Cpf,
			Cnpj:          inputDto.Cnpj,
			Email:         inputDto.Email,
			Pix_type:      "invalidType",
			Pix_key:       inputDto.Pix_key,
		},
		output: &OutputDto{},
		err:    creatingPixErr,
	},
	{
		description: "if an error accour on saving the receiver return error",
		input:       &inputDto,
		output:      &OutputDto{},
		err:         saveReceiverErr,
	},
	{
		description: "if an error accour on saving the pix return error",
		input:       &inputDto,
		output:      &OutputDto{},
		err:         savingPixErr,
	},
}

func TestCreateReceiver(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			repo := setupMock(tc)

			usecase := &Create{Repo: repo}

			outputDto, err := usecase.Execute(tc.input)

			if tc.err != nil {
				assert.ErrorAs(t, err, &tc.err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.output, outputDto)
		})
	}
}

func setupMock(tc testCase) *mocks.MockRepo {
	repo := &mocks.MockRepo{}
	receiverId := setAddReceiverExpectation(tc, repo)
	setAddPixExpectation(receiverId, tc, repo)

	return repo
}

func setAddReceiverExpectation(tc testCase, repo *mocks.MockRepo) uint {
	receiver_param := entity.NewReceiver(tc.input.CorporateName, tc.input.Cpf, tc.input.Cnpj, tc.input.Email, entity.STATUS_DRAFT)
	receiver_return := *receiver_param
	receiver_return.Id = 1

	err := errors.New("error")

	if tc.err == nil || !errors.As(tc.err, &saveReceiverErr) {
		err = nil
	}

	repo.On("AddReceiver", receiver_param).Return(&receiver_return, err)

	return receiver_return.Id
}

func setAddPixExpectation(receiverId uint, tc testCase, repo *mocks.MockRepo) {
	pix_param := entity.NewPix(tc.input.Pix_type, tc.input.Pix_key)
	pix_return := *pix_param
	pix_return.Id = 1

	err := errors.New("error")

	if tc.err == nil || !errors.As(tc.err, &savingPixErr) {
		err = nil
	}

	repo.On("AddPix", receiverId, pix_param).Return(&pix_return, err)
}
