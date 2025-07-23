package glue

import (
	"github.com/ehharvey/malleus/internal/infrastructure/db/pg/inventory/generated"
	"github.com/ehharvey/malleus/internal/inventory"
)

func ProcessDbDomain(domain generated.Domain, err error) inventory.Domain {
	return inventory.Domain{
		ID:   domain.ID.String(),
		Name: domain.Name,
	}
}
