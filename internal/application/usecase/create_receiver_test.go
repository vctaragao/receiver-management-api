package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
)

type testCase struct {
	description string
	input       CreateInputDto
	output      CreateOutputDto
	err         error
}

type MockRepo struct {
	mock.Mock
}

func (mr *MockRepo) AddReceiver(r *entity.Receiver) error {
	args := mr.Called(r)
	r.Id = 1
	return args.Error(0)
}

var testCases = []testCase{
	{
		description: "with valid data insert add receiver",
		input: CreateInputDto{
			CorporateName: "Lara Natália Ana Almeida",
			Cpf:           "009.016.853-48",
			Cnpj:          "",
			Email:         "laranataliaalmeida@chavao.com.br",
		},
		output: CreateOutputDto{
			Id: 1,
			CreateInputDto: CreateInputDto{
				CorporateName: "Lara Natália Ana Almeida",
				Cpf:           "009.016.853-48",
				Cnpj:          "",
				Email:         "laranataliaalmeida@chavao.com.br",
			},
		},
		err: nil,
	},
	{
		description: "if an error accour on save return error",
		input: CreateInputDto{
			CorporateName: "Lara Natália Ana Almeida",
			Cpf:           "009.016.853-48",
			Cnpj:          "",
			Email:         "laranataliaalmeida@chavao.com.br",
		},
		output: CreateOutputDto{},
		err:    errors.New("unable to add receiver"),
	},
}

func TestCreateReceiver(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			repo := setupMock(tc)

			outputDto, err := Create(&tc.input, repo)

			if tc.err != nil {
				assert.Error(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.output, *outputDto)
		})
	}
}

func setupMock(tc testCase) *MockRepo {
	repo := &MockRepo{}
	repo.On("AddReceiver", mock.Anything).Return(tc.err)
	return repo
}
