package create_receiver

import "fmt"

type CreatingReceiverErr struct {
	err error
}

func (e *CreatingReceiverErr) Error() string {
	return fmt.Sprintf("creating receiver: %v", e.err)
}

type SaveReceiverErr struct {
	err error
}

func (e *SaveReceiverErr) Error() string {
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
