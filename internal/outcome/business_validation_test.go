package outcome

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestBusinessValidationResultSucceededFalse(t *testing.T) {
	arrange := BusinessValidationResult{
		Tests: []BusinessValidationTest{
			{
				Succeeded: false,
			},
		},
	}

	expected := false

	actual := arrange.Succeded()

	if expected != actual {
		t.Errorf("wanted %t got %t", expected, actual)
	}
}

func TestBusinessValidationResultSucceededTrue(t *testing.T) {
	arrange := BusinessValidationResult{
		Tests: []BusinessValidationTest{
			{
				Succeeded: true,
			},
		},
	}

	expected := true

	actual := arrange.Succeded()

	if expected != actual {
		t.Errorf("wanted %t got %t", expected, actual)
	}
}

func TestServiceValidationResultSucceededTrueEmpty(t *testing.T) {
	arrange := BusinessValidationResult{
		Tests: []BusinessValidationTest{},
	}

	expected := true

	actual := arrange.Succeded()

	if expected != actual {
		t.Errorf("wanted %t got %t", expected, actual)
	}
}

func TestBusinessValidationTestError(t *testing.T) {
	arrange := BusinessValidationTest{
		Code:    "foo",
		Field:   "bar",
		Message: "baz",
	}

	expected := fmt.Sprintf(
		"service validation failed, Code %s, Field  %s, Message %s",
		arrange.Code, arrange.Field, arrange.Message,
	)

	actual := arrange.Error()

	if expected != actual {
		t.Errorf("wanted %s got %s", expected, actual)
	}
}

func stubTestFunctionFailAlways(context context.Context, _ int, _ struct{}) BusinessValidationTest {
	return BusinessValidationTest{
		Succeeded: false,
	}
}

func stubTestFunctionFailNever(context context.Context, _ int, _ struct{}) BusinessValidationTest {
	return BusinessValidationTest{
		Succeeded: true,
	}
}

func TestValidateServiceReturnFailureOnly(t *testing.T) {
	arrangeTestFuncs := [...]BusinessValidationFunction[int, struct{}]{
		stubTestFunctionFailAlways,
		stubTestFunctionFailNever,
		stubTestFunctionFailAlways,
	}

	expected := BusinessValidationResult{
		Tests: []BusinessValidationTest{
			{
				Succeeded: false,
			},
			{
				Succeeded: false,
			},
		},
	}

	actual := ValidateBusinessRules(
		t.Context(),
		3,
		struct{}{},
		ValidationDetailLevel(ValidationReturnOnlyFailures),
		arrangeTestFuncs[:],
	)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("failed, expected %+v, actual %+v", expected, actual)
	}
}

func TestValidateBusinessReturnAll(t *testing.T) {
	arrangeTestFuncs := [...]BusinessValidationFunction[int, struct{}]{
		stubTestFunctionFailAlways,
		stubTestFunctionFailNever,
		stubTestFunctionFailAlways,
	}

	expected := BusinessValidationResult{
		Tests: []BusinessValidationTest{
			{
				Succeeded: false,
			},
			{
				Succeeded: true,
			},
			{
				Succeeded: false,
			},
		},
	}

	actual := ValidateBusinessRules(
		t.Context(),
		3,
		struct{}{},
		ValidationDetailLevel(ValidationReturnAllResults),
		arrangeTestFuncs[:],
	)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("failed, expected %+v, actual %+v", expected, actual)
	}
}
