package validation

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestServiceValidationResultSucceededFalse(t *testing.T) {
	arrange := ServiceValidationResult{
		Service: "TestServiceValidationResultSuceededFalse",
		Tests: []ServiceValidationTest{
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

func TestServiceValidationResultSucceededTrue(t *testing.T) {
	arrange := ServiceValidationResult{
		Service: "TestServiceValidationResultSuceededFalse",
		Tests: []ServiceValidationTest{
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
	arrange := ServiceValidationResult{
		Service: "TestServiceValidationResultSuceededFalse",
		Tests:   []ServiceValidationTest{},
	}

	expected := true

	actual := arrange.Succeded()

	if expected != actual {
		t.Errorf("wanted %t got %t", expected, actual)
	}
}

func TestServiceValidationTestError(t *testing.T) {
	arrange := ServiceValidationTest{
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

func stubTestFunctionFailAlways(context context.Context, _ int, _ struct{}) ServiceValidationTest {
	return ServiceValidationTest{
		Succeeded: false,
	}
}

func stubTestFunctionFailNever(context context.Context, _ int, _ struct{}) ServiceValidationTest {
	return ServiceValidationTest{
		Succeeded: true,
	}
}

func TestValidateServiceReturnFailureOnly(t *testing.T) {
	arrangeTestFuncs := [...]ServiceValidationFunction[int, struct{}]{
		stubTestFunctionFailAlways,
		stubTestFunctionFailNever,
		stubTestFunctionFailAlways,
	}

	expected := ServiceValidationResult{
		Service: "TestValidateServiceReturnFailureOnly",
		Tests: []ServiceValidationTest{
			{
				Succeeded: false,
			},
			{
				Succeeded: false,
			},
		},
	}

	actual := ValidateService(
		t.Context(),
		3,
		"TestValidateServiceReturnFailureOnly",
		struct{}{},
		ValidationDetailLevel(ValidationReturnOnlyFailures),
		arrangeTestFuncs[:],
	)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("failed, expected %+v, actual %+v", expected, actual)
	}
}

func TestValidateServiceReturnAll(t *testing.T) {
	arrangeTestFuncs := [...]ServiceValidationFunction[int, struct{}]{
		stubTestFunctionFailAlways,
		stubTestFunctionFailNever,
		stubTestFunctionFailAlways,
	}

	expected := ServiceValidationResult{
		Service: "TestValidateServiceReturnAll",
		Tests: []ServiceValidationTest{
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

	actual := ValidateService(
		t.Context(),
		3,
		"TestValidateServiceReturnAll",
		struct{}{},
		ValidationDetailLevel(ValidationReturnAllResults),
		arrangeTestFuncs[:],
	)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("failed, expected %+v, actual %+v", expected, actual)
	}
}
