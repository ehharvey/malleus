package repository

import (
	"context"

	"github.com/ehharvey/malleus/internal/infrastructure/db/pg/inventory/generated"
	"github.com/ehharvey/malleus/internal/infrastructure/db/pg/inventory/glue"
	"github.com/ehharvey/malleus/internal/inventory"
	"github.com/ehharvey/malleus/internal/outcome"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Queries interface {
	InsertOneDomain(ctx context.Context, name string) (generated.Domain, error)
	CheckExistsDomainByName(ctx context.Context, name string) (bool, error)
	WithTx(tx pgx.Tx) *generated.Queries
}

type InventoryRepository struct {
	queries Queries
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
	input inventory.CreateDomainParams,
) (inventory.Domain, outcome.DbResult) {
	insert, err := repo.queries.InsertOneDomain(context, input.Name)
	dbResult := outcome.DbResult{
		QueryFunction: "CreateDomain",
		Err:           err,
	}
	return glue.ProcessDbDomain(insert, err), dbResult
}

func (repo InventoryRepository) CheckExistsDomainByName(
	context context.Context,
	name string,
) (bool, outcome.DbResult) {
	check, err := repo.queries.CheckExistsDomainByName(context, name)

	dbResult := outcome.DbResult{
		QueryFunction: "CheckExistsDomainByName",
		Err:           err,
	}

	return check, dbResult
}
