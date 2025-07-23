package inventory

import (
	"context"
	"fmt"

	"github.com/ehharvey/malleus/internal/outcome"
)

// initialize check function arrays here!
var createDomainServiceCheckFunctions = [...]outcome.BusinessValidationFunction[CreateDomainParams, Repository]{
	checkDomainUniqueness,
}

// --

func checkDomainUniqueness(
	context context.Context,
	input CreateDomainParams,
	repository Repository,
) outcome.BusinessValidationTest {
	// Check if check already exists with this name
	check, dbResult := repository.CheckExistsDomainByName(context, input.Name)

	result := outcome.BusinessValidationTest{
		Succeeded: false,
		Field:     "Name",
		DbResult:  dbResult,
	}

	if result.DbResult.Err != nil {
		result.Code = "DbError"
	} else if check {
		result.Code = "NotUnique"
		result.Message = fmt.Sprintf("domain %s already exists", input.Name)
	} else {
		result.Succeeded = true
	}

	return result
}
