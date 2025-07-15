package inventory

import (
	"context"
	"fmt"

	"github.com/ehharvey/malleus/internal/validation"
)

// initialize check function arrays here!
var createDomainServiceCheckFunctions = [...]validation.ServiceValidationFunction[*CreateDomainParams, Repository]{
	checkDomainUniqueness,
}

// --

func checkDomainUniqueness(context context.Context, input *CreateDomainParams, repository Repository) validation.ServiceValidationTest {
	// Check if domain already exists with this name
	check_domain, checK_err := repository.GetDomainByName(context, input.Name)

	result := validation.ServiceValidationTest{
		Name:      "checkDomainUniqueness",
		Succeeded: false,
		Field:     "Name",
	}

	if checK_err != nil {
		result.Code = "DbError"
		result.DbError = checK_err
	} else if check_domain != nil {
		result.Code = "NotUnique"
		result.Message = fmt.Sprintf("domain %s already exists", input.Name)
	} else {
		result.Succeeded = true
	}

	return result
}
