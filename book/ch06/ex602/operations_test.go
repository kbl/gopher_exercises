package intset

import (
	"testing"
)

func TestAddAll(t *testing.T) {
	var s IntSet
	var repr string
	expected := "{1 10 100 1000}"

	s.AddAll(1, 10, 100, 1000)
	if repr = s.String(); repr != expected {
		t.Errorf("s.String() = %v, wants %v", repr, expected)
	}
}
