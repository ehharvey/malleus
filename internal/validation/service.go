package validation

import (
	"context"
	"fmt"
)

type ServiceValidationFunction[T any, R any] func(context context.Context, input T, repository R) ServiceValidationTest

type ServiceValidationResult struct {
	Service string
	Tests   []ServiceValidationTest
}

type ServiceValidationTest struct {
	Succeeded bool
	Name      string
	Code      string
	Field     string
	Message   string
	DbError   error
}

func (se *ServiceValidationResult) Succeded() bool {
	for _, te := range se.Tests {
		if !te.Succeeded {
			return false
		}
	}

	return true
}

func (se *ServiceValidationTest) Error() string {
	return fmt.Sprintf(
		"service validation failed, Code %s, Field  %s, Message %s",
		se.Code, se.Field, se.Message,
	)
}

func ValidateService[T any, R any](
	context context.Context,
	input T,
	name string,
	repository R,
	validationDetailLevel ValidationDetailLevel,
	testFunctions []ServiceValidationFunction[T, R],
) ServiceValidationResult {
	var results []ServiceValidationTest

	for _, tf := range testFunctions {
		test_result := tf(context, input, repository)

		if validationDetailLevel == ValidationReturnAllResults ||
			!test_result.Succeeded {
			results = append(results, test_result)
		}
	}

	return ServiceValidationResult{
		Service: name,
		Tests:   results,
	}
}
