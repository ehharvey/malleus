package validation

import "fmt"

type ModelValidationFunction[T any] func(input T) ModelValidationCheckResult

type ModelValidationResult struct {
	Model string
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

func (ve *ModelValidationCheckResult) Error() string {
	return fmt.Sprintf(
		"model validation failed for Field %s, and Value %s. Message: %s",
		ve.Field, ve.Value, ve.Message,
	)
}

func ValidateModel[T any](
	input T,
	name string,
	validationDetailLevel ValidationDetailLevel,
	testFunctions []ModelValidationFunction[T],
) ModelValidationResult {
	var checkResults []ModelValidationCheckResult

	for _, tf := range testFunctions {
		test_result := tf(input)
		if validationDetailLevel == ValidationReturnAllResults || !test_result.Succeeded {
			checkResults = append(checkResults, test_result)
		}
	}

	return ModelValidationResult{
		Model: name,
		Tests: checkResults,
	}
}
