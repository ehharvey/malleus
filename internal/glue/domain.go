package glue

import (
	"github.com/ehharvey/malleus/internal/graph/model"
	"github.com/ehharvey/malleus/internal/inventory"
	"github.com/ehharvey/malleus/internal/outcome"
)

func DomainModelToGraphql(domain *inventory.Domain) *model.Domain {
	return &model.Domain{
		ID:   domain.ID,
		Name: domain.Name,
	}
}

func ServiceResultToGraphql[T any, U any](
	serviceResult outcome.ServiceResult[T],
	mapper func(*T) *U,
) (*U, error) {
	if serviceResult.Result == nil {
		return nil, serviceResult
	} else {
		return mapper(serviceResult.Result), nil
	}
}
