package outcome

import (
	"errors"
	"fmt"
)

type ModelValidationFunction[T any] func(input T) ModelValidationCheckResult

type ModelValidationResult struct {
	Tests []ModelValidationCheckResult
}

type ModelValidationCheckResult struct {
	Succeeded bool
	Name      string
	Field     string
	Value     string
	Message   string
}

func (ve *ModelValidationResult) Succeeded() bool {
	for _, t := range ve.Tests {
		if !t.Succeeded {
			return false
		}
	}

	return true
}

func (ve ModelValidationResult) CombineErrors() error {
	var errs []error
	for _, e := range ve.Tests {
		errs = append(errs, e)
	}
	return errors.Join(errs...)
}

func (ve ModelValidationCheckResult) Error() string {
	return fmt.Sprintf(
		"model validation failed for Field %s, and Value %s. Message: %s",
		ve.Field, ve.Value, ve.Message,
	)
}

func ValidateModel[T any](
	input T,
	validationDetailLevel ValidationDetailLevel,
	testFunctions []ModelValidationFunction[T],
) ModelValidationResult {
	checkResults := []ModelValidationCheckResult{}

	for _, tf := range testFunctions {
		test_result := tf(input)
		if validationDetailLevel == ValidationReturnAllResults || !test_result.Succeeded {
			checkResults = append(checkResults, test_result)
		}
	}

	return ModelValidationResult{
		Tests: checkResults,
	}
}
