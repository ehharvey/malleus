package repository

import (
	"context"
	"testing"

	"github.com/ehharvey/malleus/internal/infrastructure/db/pg/inventory/generated"
	"github.com/ehharvey/malleus/internal/inventory"
	"github.com/ehharvey/malleus/internal/outcome"
	"github.com/google/go-cmp/cmp"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type queriesStub struct{}

func (q *queriesStub) InsertOneDomain(
	ctx context.Context,
	name string,
) (generated.Domain, error) {
	return generated.Domain{
		ID: pgtype.UUID{
			Bytes: [16]byte{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		Name: name,
	}, nil
}

func (q *queriesStub) CheckExistsDomainByName(
	ctx context.Context,
	name string,
) (bool, error) {
	return false, nil
}

func (q *queriesStub) WithTx(tx pgx.Tx) *generated.Queries {
	return &generated.Queries{}
}

func TestCreateDomain(t *testing.T) {
	arrangeRepository := InventoryRepository{
		queries: &queriesStub{},
	}

	arrangeInput := inventory.CreateDomainParams{
		Name: "TestCreateDomain",
	}

	expectedDomain := inventory.Domain{
		ID: pgtype.UUID{
			Bytes: [16]byte{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		}.String(),
		Name: "TestCreateDomain",
	}
	expectedErr := outcome.DbResult{
		QueryFunction: "CreateDomain",
		Err:           nil,
	}

	actualDomain, actualErr := arrangeRepository.CreateDomain(
		context.Background(),
		arrangeInput,
	)

	if expectedErr != actualErr {
		t.Errorf("expectedErr != actualErr: \n%s", cmp.Diff(expectedErr, actualErr))
	}

	if !cmp.Equal(expectedDomain, actualDomain) {
		t.Errorf("Mismatch (-want +got):\n%s", cmp.Diff(expectedDomain, actualDomain))
	}
}

func TestCheckExistsDomainByName(t *testing.T) {
	arrangeRepository := InventoryRepository{
		queries: &queriesStub{},
	}

	arrangeName := "TestCheckExistsDomainByName"

	expectedBool := false

	expectedErr := outcome.DbResult{
		QueryFunction: "CheckExistsDomainByName",
		Err:           nil,
	}

	actualBool, actualErr := arrangeRepository.CheckExistsDomainByName(
		context.Background(),
		arrangeName,
	)

	if expectedErr != actualErr {
		t.Errorf("expectedErr != actualErr: \n%s", cmp.Diff(expectedErr, actualErr))
	}

	if expectedBool != actualBool {
		t.Errorf("expectedBool %t != actualBool %t", expectedBool, actualBool)
	}
}
