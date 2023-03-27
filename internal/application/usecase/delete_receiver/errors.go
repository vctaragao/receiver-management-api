package delete_receiver

import (
	"errors"
	"fmt"
)

var dReceiverErr *deletingReceiverErr

var ErrReceiversIdsAreRequired = errors.New("at leat one receiver_id is required")

type deletingReceiverErr struct {
	err error
}

func (e *deletingReceiverErr) Error() string {
	return fmt.Sprintf("deleting receiver: %v", e.err)
}

func IsBusinessLogicError(err error) bool {
	return errors.Is(err, ErrReceiversIdsAreRequired)
}
