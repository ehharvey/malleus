package graphqlglue

import (
	"github.com/ehharvey/malleus/internal/graph/model"
	"github.com/ehharvey/malleus/internal/inventory"
)

func DomainModelToGraphql(domain inventory.Domain) *model.Domain {
	return &model.Domain{
		ID:   domain.ID,
		Name: domain.Name,
	}
}
