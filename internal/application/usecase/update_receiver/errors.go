package update_receiver

import (
	"errors"
	"fmt"
)

var sPixErr *SavingPixErr
var cPixErr *CreatingPixErr
var iReceiverErr *invalidReceiverErr
var fReceiverErr *findingReceiverError

type findingReceiverError struct {
	err error
}

func (e *findingReceiverError) Error() string {
	return fmt.Sprintf("finding receiver: %v", e.err)
}

type invalidReceiverErr struct {
	err error
}

func (e *invalidReceiverErr) Error() string {
	return fmt.Sprintf("validating receiver: %v", e.err)
}

type SavingPixErr struct {
	err error
}

func (e *SavingPixErr) Error() string {
	return fmt.Sprintf("saving pix: %v", e.err)
}

type CreatingPixErr struct {
	err error
}

func (e *CreatingPixErr) Error() string {
	return fmt.Sprintf("creating pix: %v", e.err)
}

func IsCreateBussinesLogicError(err error) bool {
	if errors.As(err, &cPixErr) || errors.As(err, &fReceiverErr) {
		return true
	}

	return false
}
