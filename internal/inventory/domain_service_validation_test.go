package inventory

import (
	"context"
	"errors"
	"testing"

	"github.com/ehharvey/malleus/internal/outcome"
)

type repoMockFindUnique struct{}
type repoMockNoUnique struct{}
type repoMockDbError struct{}

func stubCreateDomain(input CreateDomainParams) (Domain, outcome.DbResult) {
	dbResult := outcome.DbResult{
		QueryFunction: "CreateDomain",
		Err:           nil,
	}
	return Domain{
		ID:   "123",
		Name: input.Name,
	}, dbResult
}

func (repo repoMockNoUnique) CheckExistsDomainByName(
	context context.Context,
	input string,
) (bool, outcome.DbResult) {
	dbResult := outcome.DbResult{
		Err:           nil,
		QueryFunction: "CheckExistsDomainByName",
	}
	return false, dbResult
}

func (repo repoMockNoUnique) CreateDomain(
	context context.Context,
	input CreateDomainParams,
) (Domain, outcome.DbResult) {
	return stubCreateDomain(input)
}

func (repo repoMockFindUnique) CheckExistsDomainByName(
	context context.Context,
	input string,
) (bool, outcome.DbResult) {
	dbResult := outcome.DbResult{
		QueryFunction: "CheckExistsDomainByName",
		Err:           nil,
	}
	return true, dbResult
}

func (repo repoMockFindUnique) CreateDomain(
	context context.Context,
	input CreateDomainParams,
) (Domain, outcome.DbResult) {
	return stubCreateDomain(input)
}

func (repo repoMockDbError) CheckExistsDomainByName(
	context context.Context,
	input string,
) (bool, outcome.DbResult) {
	dbResult := outcome.DbResult{
		QueryFunction: "CheckExistsDomainByName",
		Err:           errors.New("CheckExistsDomainByNameFailed"),
	}
	return false, dbResult
}

func (repo repoMockDbError) CreateDomain(
	context context.Context,
	input CreateDomainParams,
) (Domain, outcome.DbResult) {
	dbResult := outcome.DbResult{
		QueryFunction: "CreateDomain",
		Err:           errors.New("CreateDomainByNameFailed"),
	}
	return Domain{}, dbResult
}

func TestCheckDomainUniqueness(t *testing.T) {
	tests := []struct {
		name     string
		arrange  CreateDomainParams
		repoMock Repository
		expected outcome.BusinessValidationTest
	}{
		{"Unique", CreateDomainParams{Name: "unique"}, repoMockNoUnique{}, outcome.BusinessValidationTest{
			Succeeded: true,
		}},
		{"Unique", CreateDomainParams{Name: "not unique"}, repoMockFindUnique{}, outcome.BusinessValidationTest{
			Succeeded: false,
		}},
		{"DB error", CreateDomainParams{Name: "db error"}, repoMockDbError{}, outcome.BusinessValidationTest{
			Succeeded: false,
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := checkDomainUniqueness(context.Background(), tt.arrange, tt.repoMock)

			if tt.expected.Succeeded != actual.Succeeded {
				t.Errorf("checkDomainUniqueness(%s) = %t, want %t", tt.arrange, actual.Succeeded, tt.expected.Succeeded)
			}
		})
	}
}
