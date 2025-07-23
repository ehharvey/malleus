package outcome

import (
	"context"
	"errors"
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

func (se *BusinessValidationResult) CombineErrors() error {
	var errs []error
	for _, e := range se.Tests {
		errs = append(errs, e)
	}
	return errors.Join(errs...)
}

func (st BusinessValidationTest) Error() string {
	if st.Succeeded {
		return fmt.Sprintf(
			"service validation passed, Code %s, Field %s, Message %s",
			st.Code, st.Field, st.Message,
		)
	} else if st.DbResult.Err == nil {
		return fmt.Sprintf(
			"service validation failed, Code %s, Field %s, Message %s",
			st.Code, st.Field, st.Message,
		)
	} else {

		return fmt.Errorf(
			"service validation failed, Code %s, Field %s, Message %s: %w",
			st.Code, st.Field, st.Message, st.DbResult,
		).Error()
	}
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
