package inwx

import "fmt"

// ErrUnexpectedStatus indicates that a Call was responded to
// but the Status doesn't match any of the expected Statuses of the Call.
type ErrUnexpectedStatus struct {
	Expected []Status
	Got      Status
}

func (e *ErrUnexpectedStatus) Error() string {
	const format = "unexpected status code: expected %v, got %v"
	return fmt.Sprintf(format, e.Expected, e.Got)
}
