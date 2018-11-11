package intset

import (
	"testing"
)

func TestElems(t *testing.T) {
	var s IntSet

	expected := []int{1, 10, 30, 90, 100}
	for _, e := range expected {
		s.Add(e)
	}
	actual := s.Elems()
	equal := len(actual) == len(expected)
	for i, e := range actual {
		equal = equal && e == expected[i]
		if !equal {
			break
		}
	}

	if !equal {
		t.Errorf("s.Elems() = %v, wants %v", actual, expected)
	}
}
