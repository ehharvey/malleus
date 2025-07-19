package inventory

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewService(t *testing.T) {
	arrangeRepo := repoMockNoUnique{}

	expected := &Service{
		Repository: repoMockNoUnique{},
	}

	actual := NewService(arrangeRepo)

	if !cmp.Equal(expected, actual) {
		t.Errorf("Mismatch (-want +got):\n%s", cmp.Diff(expected, actual))
	}
}
