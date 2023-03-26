package create_receiver

import (
	"errors"
	"fmt"
)

var sPixErr *SavingPixErr
var cPixErr *CreatingPixErr
var sReceiverErr *saveReceiverErr
var cReceiverErr *CreatingReceiverErr

type CreatingReceiverErr struct {
	err error
}

func (e *CreatingReceiverErr) Error() string {
	return fmt.Sprintf("creating receiver: %v", e.err)
}

type saveReceiverErr struct {
	err error
}

func (e *saveReceiverErr) Error() string {
	return fmt.Sprintf("saving receiver: %v", e.err)
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

func IsBussinesLogicError(err error) bool {
	if errors.As(err, &cPixErr) || errors.As(err, &cReceiverErr) {
		return true
	}

	return false
}
