package outcome

import (
	"context"
	"fmt"
)

type BusinessValidationFunction[T any, R any] func(
	context context.Context,
	input T,
	repository R,
) BusinessValidationTest

type BusinessValidationResult struct {
	Tests []BusinessValidationTest
}

type BusinessValidationTest struct {
	Succeeded bool
	Code      string
	Field     string
	Message   string
	DbResult  DbResult
}

func (se *BusinessValidationResult) Succeded() bool {
	for _, te := range se.Tests {
		if !te.Succeeded {
			return false
		}
	}

	return true
}

func (se *BusinessValidationTest) Error() string {
	return fmt.Sprintf(
		"service validation failed, Code %s, Field  %s, Message %s",
		se.Code, se.Field, se.Message,
	)
}

func ValidateBusinessRules[T any, R any](
	context context.Context,
	input T,
	repository R,
	validationDetailLevel ValidationDetailLevel,
	testFunctions []BusinessValidationFunction[T, R],
) BusinessValidationResult {
	results := []BusinessValidationTest{}

	for _, tf := range testFunctions {
		test_result := tf(context, input, repository)

		if validationDetailLevel == ValidationReturnAllResults ||
			!test_result.Succeeded {
			results = append(results, test_result)
		}
	}

	return BusinessValidationResult{
		Tests: results,
	}
}
