package delete_receiver

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vctaragao/receiver-management-api/internal/test/mocks"
)

type testCase struct {
	description string
	input       *InputDto
	err         error
}

var testCases = []testCase{
	{
		description: "By passing ids delete receivers",
		input: &InputDto{
			ReceiversIds: []uint{1, 2},
		},
		err: nil,
	},
	{
		description: "By passing no ids return an error",
		input: &InputDto{
			ReceiversIds: []uint{},
		},
		err: ErrReceiversIdsAreRequired,
	},
	{
		description: "If cant delete return an error",
		input: &InputDto{
			ReceiversIds: []uint{1},
		},
		err: dReceiverErr,
	},
}

func TestDeleteReceiver(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			repo := setupMock(tc)

			usecase := &Delete{Repo: repo}

			err := usecase.Execute(tc.input)

			if tc.err != nil {
				assert.ErrorAs(t, err, &tc.err)
				assert.ErrorIs(t, err, tc.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func setupMock(tc testCase) *mocks.MockRepo {
	repo := &mocks.MockRepo{}
	setDeleteReceiversExpectation(tc, repo)

	return repo
}

func setDeleteReceiversExpectation(tc testCase, repo *mocks.MockRepo) {

	var err error
	if tc.err != nil && errors.As(tc.err, &dReceiverErr) {
		err = errors.New("error from database")
	}

	repo.On("DeleteReceivers", tc.input.ReceiversIds).Return(err)
}
