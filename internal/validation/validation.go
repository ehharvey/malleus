package validation

type ValidationDetailLevel int

const (
	ValidationReturnOnlyFailures = iota
	ValidationReturnAllResults
)
