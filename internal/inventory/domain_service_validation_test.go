package inventory

import (
	"context"
	"errors"
	"testing"

	"github.com/ehharvey/malleus/internal/validation"
)

type repoMockFindUnique struct{}
type repoMockNoUnique struct{}
type repoMockDbError struct{}

func stubCreateDomain(input *CreateDomainParams) (*Domain, error) {
	return &Domain{
		ID:   "123",
		Name: input.Name,
	}, nil
}

func (repo repoMockNoUnique) GetDomainByName(context context.Context, input string) (*Domain, error) {
	return nil, nil
}

func (repo repoMockNoUnique) CreateDomain(context context.Context, input *CreateDomainParams) (*Domain, error) {
	return stubCreateDomain(input)
}

func (repo repoMockFindUnique) GetDomainByName(context context.Context, input string) (*Domain, error) {
	return &Domain{
		ID:   "123",
		Name: input,
	}, nil
}

func (repo repoMockFindUnique) CreateDomain(context context.Context, input *CreateDomainParams) (*Domain, error) {
	return stubCreateDomain(input)
}

func (repo repoMockDbError) GetDomainByName(context context.Context, input string) (*Domain, error) {
	return nil, errors.New("db mock error")
}

func (repo repoMockDbError) CreateDomain(context context.Context, input *CreateDomainParams) (*Domain, error) {
	return nil, errors.New("db mock error")
}

func TestCheckDomainUniqueness(t *testing.T) {
	tests := []struct {
		name     string
		arrange  CreateDomainParams
		repoMock Repository
		expected validation.ServiceValidationTest
	}{
		{"Unique", CreateDomainParams{Name: "unique"}, repoMockNoUnique{}, validation.ServiceValidationTest{
			Succeeded: true,
		}},
		{"Unique", CreateDomainParams{Name: "not unique"}, repoMockFindUnique{}, validation.ServiceValidationTest{
			Succeeded: false,
		}},
		{"DB error", CreateDomainParams{Name: "db error"}, repoMockDbError{}, validation.ServiceValidationTest{
			Succeeded: false,
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := checkDomainUniqueness(context.Background(), &tt.arrange, tt.repoMock)

			if tt.expected.Succeeded != actual.Succeeded {
				t.Errorf("checkDomainUniqueness(%s) = %t, want %t", tt.arrange, actual.Succeeded, tt.expected.Succeeded)
			}
		})
	}
}
