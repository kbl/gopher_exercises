package intset

import (
	"testing"
)

func TestIntersectWith(t *testing.T) {
	var s1, s2 IntSet
	s1.Add(1)
	s1.Add(10)
	s1.Add(100)
	s1.Add(1000)

	s2.Add(1)
	s2.Add(100)
	s2.Add(1000)

	var repr string
	expected := "x{1 100 1000}"

	s1.IntersectWith(&s2)
	if repr = s1.String(); repr != expected {
		t.Errorf("s1.String() = %v, wants %v", repr, expected)
	}
}

func TestDifferrenceWith(t *testing.T) {
	var s1, s2 IntSet
	s1.Add(1)
	s1.Add(10)
	s1.Add(100)
	s1.Add(1000)

	s2.Add(1)
	s2.Add(100)
	s2.Add(1000)

	var repr string
	expected := "xxx{10}"

	s1.DifferenceWith(&s2)
	if repr = s1.String(); repr != expected {
		t.Errorf("s1.String() = %v, wants %v", repr, expected)
	}
}
