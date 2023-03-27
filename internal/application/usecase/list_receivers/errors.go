package list_receivers

import (
	"errors"
	"fmt"
)

var iPixErr *invalidPixErr
var uPixErr *UpdatingPixErr
var iReceiverErr *invalidReceiverErr
var uReceiverErr *UpdatingReceiverErr
var fReceiverErr *findingReceiverError

var invalidStatusErr = errors.New("invalid status")
var invalidCorporateNameErr = errors.New("invalid corporate name")
var invalidPixTypeErr = errors.New("invalid pix_type")
var invalidPixKeyErr = errors.New("invalid pix_key")

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

type invalidPixErr struct {
	err error
}

func (e *invalidPixErr) Error() string {
	return fmt.Sprintf("validating pix: %v", e.err)
}

type UpdatingPixErr struct {
	err error
}

func (e *UpdatingPixErr) Error() string {
	return fmt.Sprintf("updating pix: %v", e.err)
}

type UpdatingReceiverErr struct {
	err error
}

func (e *UpdatingReceiverErr) Error() string {
	return fmt.Sprintf("updating receiver: %v", e.err)
}

func IsBussinesLogicError(err error) bool {
	return errors.As(err, &iReceiverErr) || errors.As(err, &fReceiverErr) || errors.As(err, &iPixErr) || errors.As(err, &uReceiverErr)
}
