package inventory

import (
	"context"
	"log"

	"github.com/ehharvey/malleus/internal/outcome"
)

func (s *Service) CreateDomain(
	context context.Context,
	input CreateDomainParams,
	validationDetailLevel outcome.ValidationDetailLevel,
) outcome.ServiceResult[Domain] {
	result := outcome.ServiceResult[Domain]{
		Model: "Domain",
	}

	// Domain checks
	result.ModelValidationResult = outcome.ValidateModel(
		input,
		validationDetailLevel,
		createDomainModelCheckFunctions[:],
	)

	// Service checks
	if result.ModelValidationResult.Succeeded() {
		result.ServiceValidationResult = outcome.ValidateBusinessRules(
			context,
			input,
			s.Repository,
			validationDetailLevel,
			createDomainServiceCheckFunctions[:],
		)
	}

	// Creation
	if result.ModelValidationResult.Succeeded() && result.ServiceValidationResult.Succeded() {
		createResult, dbResult := s.Repository.CreateDomain(
			context,
			input,
		)

		result.PersistenceResult = dbResult
		result.Result = createResult
	}

	log.Printf("%v", result)

	return result
}
