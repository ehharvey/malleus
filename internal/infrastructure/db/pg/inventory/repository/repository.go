package repository

import (
	"context"

	"github.com/ehharvey/malleus/internal/infrastructure/db/pg/inventory/generated"
	"github.com/ehharvey/malleus/internal/inventory"
	"github.com/jackc/pgx/v5/pgxpool"
)

type InventoryRepository struct {
	queries *generated.Queries
}

func NewInventoryRepository(queries *generated.Queries) InventoryRepository {
	return InventoryRepository{
		queries: queries,
	}
}

func NewInventoryQueries(pool *pgxpool.Pool) *generated.Queries {
	return generated.New(pool)
}

func (repo InventoryRepository) CreateDomain(context context.Context, input *inventory.CreateDomainParams) (*inventory.Domain, error) {
	insert, err := repo.queries.InsertOneDomain(context, input.Name)

	if err != nil {
		return nil, err
	} else {
		return &inventory.Domain{
			ID:   insert.ID.String(),
			Name: input.Name,
		}, err
	}
}

func (repo InventoryRepository) GetDomainByName(context context.Context, input string) (*inventory.Domain, error) {
	get, err := repo.queries.SelectOneDomainByName(context, input)

	if err != nil {
		return nil, err
	} else {
		return &inventory.Domain{
			ID:   get.ID.String(),
			Name: get.Name,
		}, err
	}
}
