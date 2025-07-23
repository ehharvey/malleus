package graphqlglue

import (
	"testing"

	"github.com/ehharvey/malleus/internal/outcome"
	"github.com/google/go-cmp/cmp"
)

type testStruct struct {
	Foo string
}

func testStructMapper(t testStruct) testStruct {
	return t
}

func TestServiceResultToGraphql(t *testing.T) {
	arrange := outcome.ServiceResult[testStruct]{
		Result: testStruct{
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
		Result: testStruct{},
	}

	expected := outcome.ServiceResult[testStruct]{
		Result: testStruct{},
	}

	_, err := ServiceResultToGraphql(arrange, testStructMapper)

	if !cmp.Equal(expected, err) {
		t.Errorf("Mismatch (-want +got):\n%s", cmp.Diff(expected, err))
	}
}
