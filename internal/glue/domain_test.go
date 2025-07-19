package glue

import (
	"testing"

	"github.com/ehharvey/malleus/internal/graph/model"
	"github.com/ehharvey/malleus/internal/inventory"
	"github.com/ehharvey/malleus/internal/outcome"
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

	actual := DomainModelToGraphql(&arrange)

	if !cmp.Equal(expected, actual) {
		t.Errorf("Mismatch (-want +got):\n%s", cmp.Diff(expected, actual))
	}
}

type testStruct struct {
	Foo string
}

func testStructMapper(t *testStruct) *testStruct {
	return t
}

func TestServiceResultToGraphql(t *testing.T) {
	arrange := outcome.ServiceResult[testStruct]{
		Result: &testStruct{
			Foo: "bar",
		},
	}

	expected := &testStruct{
		Foo: "bar",
	}

	actual, err := ServiceResultToGraphql(arrange, testStructMapper)

	if err != nil {
		t.Errorf("wanted err to be nil")
	} else if !cmp.Equal(expected, actual) {
		t.Errorf("Mismatch (-want +got):\n%s", cmp.Diff(expected, actual))
	}
}

func TestServiceResultToGraphqlError(t *testing.T) {
	arrange := outcome.ServiceResult[testStruct]{
		Result: nil,
	}

	expected := outcome.ServiceResult[testStruct]{
		Result: nil,
	}

	actual, err := ServiceResultToGraphql(arrange, testStructMapper)

	if actual != nil {
		t.Errorf("wanted actual to be nil")
	} else if !cmp.Equal(expected, err) {
		t.Errorf("Mismatch (-want +got):\n%s", cmp.Diff(expected, err))
	}
}
