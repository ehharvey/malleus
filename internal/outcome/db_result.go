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
	if dr.Err == nil {
		return fmt.Sprintf("no error from QueryFunction: %s", dr.QueryFunction)
	}
	return fmt.Sprintf("DbResult error in QueryFunction %s: %v", dr.QueryFunction, dr.Err)
}

// Allow unwrapping of the underlying error
func (dr DbResult) Unwrap() error {
	return dr.Err
}
