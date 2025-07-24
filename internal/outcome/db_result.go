package outcome

import "fmt"

type DbResult struct {
	QueryFunction string
	Err           error // nil when no error
}

func (dr DbResult) Succeeded() bool {
	return dr.Err == nil
}

func (dr DbResult) Error() string {
	return fmt.Sprintf("DbResult error in QueryFunction %s: %v", dr.QueryFunction, dr.Err)
}

func (dr DbResult) AsError() error {
	if dr.Err == nil {
		return dr
	} else {
		return nil
	}
}

// Allow unwrapping of the underlying error
func (dr DbResult) Unwrap() error {
	return dr.Err
}
