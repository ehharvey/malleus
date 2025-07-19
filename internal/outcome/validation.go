package outcome

type ValidationDetailLevel int

const (
	ValidationReturnOnlyFailures = iota
	ValidationReturnAllResults
)
