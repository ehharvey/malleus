package inventory

import (
	"context"
	"testing"

	"github.com/ehharvey/malleus/internal/outcome"
	"github.com/google/go-cmp/cmp"
)

func TestCreateDomain(t *testing.T) {
	arrange := Service{
		Repository: repoMockNoUnique{},
	}

	arrangeInput := CreateDomainParams{
		Name: "example.com",
	}

	expectedResult := Domain{
		ID:   "123",
		Name: "example.com",
	}

	expected := outcome.ServiceResult[Domain]{
		Result: expectedResult,
		Model:  "Domain",
		ModelValidationResult: outcome.ModelValidationResult{
			Tests: []outcome.ModelValidationCheckResult{},
		},
		ServiceValidationResult: outcome.BusinessValidationResult{
			Tests: []outcome.BusinessValidationTest{},
		},
		PersistenceResult: outcome.DbResult{
			QueryFunction: "CreateDomain",
			Err:           nil,
		},
	}

	actual := arrange.CreateDomain(
		context.Background(),
		arrangeInput,
		outcome.ValidationDetailLevel(outcome.ValidationReturnOnlyFailures),
	)

	if !cmp.Equal(expected, actual) {
		t.Errorf("Mismatch (-want +got):\n%s", cmp.Diff(expected, actual))
	}
}
