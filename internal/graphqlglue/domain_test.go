package graphqlglue

import (
	"testing"

	"github.com/ehharvey/malleus/internal/graph/model"
	"github.com/ehharvey/malleus/internal/inventory"
	"github.com/google/go-cmp/cmp"
)

func TestDomainModelToGraphql(t *testing.T) {
	arrange := inventory.Domain{
		ID:   "abc",
		Name: "TestDomainModelToGraphql",
	}

	expected := &model.Domain{
		ID:   "abc",
		Name: "TestDomainModelToGraphql",
	}

	actual := DomainModelToGraphql(arrange)

	if !cmp.Equal(expected, actual) {
		t.Errorf("Mismatch (-want +got):\n%s", cmp.Diff(expected, actual))
	}
}
