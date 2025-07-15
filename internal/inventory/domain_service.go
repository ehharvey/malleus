package inventory

import (
	"context"

	"github.com/ehharvey/malleus/internal/service"
	"github.com/ehharvey/malleus/internal/validation"
)

func (s *Service) CreateDomain(
	context context.Context,
	input *CreateDomainParams,
	validationDetailLevel validation.ValidationDetailLevel,
) service.ServiceResult[Domain] {
	result := service.ServiceResult[Domain]{}

	// Domain checks
	result.ModelValidationResult = validation.ValidateModel(
		input,
		"Domain",
		validationDetailLevel,
		createDomainModelCheckFunctions[:],
	)

	// Service checks
	if result.ModelValidationResult.Succeeded() {
		result.ServiceValidationResult = validation.ValidateService(
			context,
			input,
			"Domain",
			s.Repository,
			validationDetailLevel,
			createDomainServiceCheckFunctions[:],
		)
	}

	// Creation
	if result.ModelValidationResult.Succeeded() && result.ServiceValidationResult.Succeded() {
		create_result, create_err := s.Repository.CreateDomain(
			context,
			input,
		)

		if create_err != nil {
			result.DbError = create_err
		} else {
			result.Result = create_result
		}
	}

	return result
}
