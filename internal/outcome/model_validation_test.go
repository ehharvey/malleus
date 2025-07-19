package outcome

import (
	"fmt"
	"reflect"
	"testing"
)

func stubValidateFunctionFail(input int) ModelValidationCheckResult {
	return ModelValidationCheckResult{
		Succeeded: false,
	}
}

func stubValidateFunctionPass(input int) ModelValidationCheckResult {
	return ModelValidationCheckResult{
		Succeeded: true,
	}
}

func TestValidateModelReportOnlyFailures(t *testing.T) {
	validateFunctions := []ModelValidationFunction[int]{
		stubValidateFunctionFail,
		stubValidateFunctionPass,
		stubValidateFunctionFail,
	}

	expectedTests := []ModelValidationCheckResult{
		{
			Succeeded: false,
		},
		{
			Succeeded: false,
		},
	}

	expected := ModelValidationResult{
		Tests: expectedTests,
	}

	actual := ValidateModel(5,
		ValidationReturnOnlyFailures,
		validateFunctions,
	)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("failed, expected %+v, actual %+v", expected, actual)
	}
}

func TestValidateModelReportAll(t *testing.T) {
	validateFunctions := []ModelValidationFunction[int]{
		stubValidateFunctionFail,
		stubValidateFunctionPass,
		stubValidateFunctionFail,
	}

	expectedTests := []ModelValidationCheckResult{
		{
			Succeeded: false,
		},
		{
			Succeeded: true,
		},
		{
			Succeeded: false,
		},
	}

	expected := ModelValidationResult{
		Tests: expectedTests,
	}

	actual := ValidateModel[int](5,
		ValidationReturnAllResults,
		validateFunctions,
	)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("failed, expected %+v, actual %+v", expected, actual)
	}
}

func TestModelValidationResultSucceededFalse(t *testing.T) {
	arrangeCheckResults := []ModelValidationCheckResult{
		{
			Succeeded: false,
		},
		{
			Succeeded: true,
		},
		{
			Succeeded: false,
		},
	}

	arrange := ModelValidationResult{
		Tests: arrangeCheckResults,
	}

	expected := false

	actual := arrange.Succeeded()

	if expected != actual {
		t.Errorf("wanted %t got %t", expected, actual)
	}
}

func TestModelValidationResultSucceededTrue(t *testing.T) {
	arrangeCheckResults := []ModelValidationCheckResult{
		{
			Succeeded: true,
		},
		{
			Succeeded: true,
		},
		{
			Succeeded: true,
		},
	}

	arrange := ModelValidationResult{
		Tests: arrangeCheckResults,
	}

	expected := true

	actual := arrange.Succeeded()

	if expected != actual {
		t.Errorf("wanted %t got %t", expected, actual)
	}
}

func TestModelValidationCheckResultError(t *testing.T) {
	arrange := ModelValidationCheckResult{
		Field:   "field",
		Value:   "value",
		Message: "message",
	}

	expected := fmt.Sprintf(
		"model validation failed for Field %s, and Value %s. Message: %s",
		arrange.Field, arrange.Value, arrange.Message,
	)

	actual := arrange.Error()

	if expected != actual {
		t.Errorf("wanted %s, got %s", expected, actual)
	}
}
