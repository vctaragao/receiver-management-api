package list_receivers

import "fmt"

var fReceiverErr *findingReceiverErr

type findingReceiverErr struct {
	err error
}

func (e *findingReceiverErr) Error() string {
	return fmt.Sprintf("creating receiver: %v", e.err)
}
