package list_receivers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
	"github.com/vctaragao/receiver-management-api/internal/test/mocks"
)

type testCase struct {
	description    string
	input          *InputDto
	output         *OutputDto
	receiversFound []entity.Receiver
	total          int
	err            error
}

var receivers = []entity.Receiver{
	{
		Id:            1,
		CorporateName: "Olivia Daiane Tânia Rezende",
		CpfCnpj:       "106.762.957-20",
		Email:         "bryan_barbosa@prcondominios.com.br",
		Status:        entity.STATUS_DRAFT,
		Pix: &entity.Pix{
			Type: "CPF",
			Key:  "106.762.957-20",
		},
	},
	{
		Id:            2,
		CorporateName: "Alana Sara Silveira",
		CpfCnpj:       "366.101.352-15",
		Email:         "alana_silveira@land.com.br",
		Status:        entity.STATUS_DRAFT,
		Pix: &entity.Pix{
			Type: "CPF",
			Key:  "366.101.352-15",
		},
	},
}

var testCases = []testCase{
	{
		description: "with empty search param return all receivers paginated",
		input: &InputDto{
			SearchParam: "",
			Page:        1,
		},
		output:         &OutputDto{Total: len(receivers), Receivers: receivers},
		receiversFound: receivers,
		total:          len(receivers),
		err:            nil,
	},
	{
		description: "with search param return paginated all receivers that match",
		input: &InputDto{
			SearchParam: "RASCUNHO",
			Page:        1,
		},
		output:         &OutputDto{Total: len(receivers), Receivers: receivers},
		receiversFound: receivers,
		total:          len(receivers),
		err:            nil,
	},
}

func TestListReceiver(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			repo := setupMock(tc)

			usecase := &List{Repo: repo}

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
	setFindReceiversExpectation(tc, repo)
	setFindReceiversByExpectation(tc, repo)

	return repo
}

func setFindReceiversExpectation(tc testCase, repo *mocks.MockRepo) {

	var err error
	if tc.err != nil && errors.As(tc.err, &fReceiverErr) {
		err = errors.New("error")
	}

	repo.On("FindReceivers", tc.input.Page).Return(tc.receiversFound, tc.total, err)
}

func setFindReceiversByExpectation(tc testCase, repo *mocks.MockRepo) {

	var err error
	if tc.err != nil && errors.As(tc.err, &fReceiverErr) {
		err = errors.New("error")
	}

	repo.On("FindReceiversBy", tc.input.SearchParam, tc.input.Page).Return(tc.receiversFound, tc.total, err)
}
