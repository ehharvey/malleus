package repository

import (
	"context"

	"github.com/ehharvey/malleus/internal/infrastructure/db/pg/inventory/generated"
	"github.com/ehharvey/malleus/internal/inventory"
	"github.com/ehharvey/malleus/internal/outcome"
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

func (repo InventoryRepository) CreateDomain(
	context context.Context,
	input *inventory.CreateDomainParams,
) (*inventory.Domain, outcome.DbResult) {
	dbResult := outcome.DbResult{
		QueryFunction: "CreateDomain",
		Succeded:      false,
	}
	insert, err := repo.queries.InsertOneDomain(context, input.Name)

	if err != nil {
		return nil, dbResult
	} else {
		dbResult.Succeded = true
		return &inventory.Domain{
			ID:   insert.ID.String(),
			Name: input.Name,
		}, dbResult
	}
}

func (repo InventoryRepository) GetDomainByName(
	context context.Context,
	input string,
) (*inventory.Domain, outcome.DbResult) {
	dbResult := outcome.DbResult{
		QueryFunction: "GetDomainByName",
		Succeded:      false,
	}
	get, err := repo.queries.SelectOneDomainByName(context, input)

	if err != nil {
		return nil, dbResult
	} else {
		dbResult.Succeded = true
		return &inventory.Domain{
			ID:   get.ID.String(),
			Name: get.Name,
		}, dbResult
	}
}
