package inventory

import (
	"strings"
	"testing"

	"github.com/ehharvey/malleus/internal/validation"
)

var sucessfulModelValidationCheckResult = validation.ModelValidationCheckResult{
	Succeeded: true,
}

var failedModelValidationCheckResult = validation.ModelValidationCheckResult{
	Succeeded: false,
}

func TestCheckValidDomainNameFormat(t *testing.T) {
	tests := []struct {
		name     string
		arrange  CreateDomainParams
		expected validation.ModelValidationCheckResult
	}{
		{"Example_com", CreateDomainParams{Name: "example.com"}, sucessfulModelValidationCheckResult},
		{"Invalid format", CreateDomainParams{Name: "-asdf.com"}, failedModelValidationCheckResult},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := checkValidDomainNameFormat(&tt.arrange)

			if tt.expected.Succeeded != actual.Succeeded {
				t.Errorf("isValidDomainNameFormat(%s) = %t; want %t", tt.arrange, actual.Succeeded, tt.expected.Succeeded)
			}
		})
	}
}

func TestIsValidDomainNameLength(t *testing.T) {
	tests := []struct {
		name     string
		arrange  CreateDomainParams
		expected validation.ModelValidationCheckResult
	}{
		{"Correct length", CreateDomainParams{Name: "example.com"}, sucessfulModelValidationCheckResult},
		{"Too long", CreateDomainParams{Name: strings.Repeat("example.com", 300)}, failedModelValidationCheckResult},
		{"Zero length", CreateDomainParams{Name: ""}, failedModelValidationCheckResult},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := checkValidDomainNameLength(&tt.arrange)

			if tt.expected.Succeeded != actual.Succeeded {
				t.Errorf("isValidDomainNameLength(%s) = %t; want %t", tt.arrange, actual.Succeeded, tt.expected.Succeeded)
			}
		})
	}
}
