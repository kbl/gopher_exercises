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
	expected := "{1 100 1000}"

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
	expected := "{10}"

	s1.DifferenceWith(&s2)
	if repr = s1.String(); repr != expected {
		t.Errorf("s1.String() = %v, wants %v", repr, expected)
	}
}

func TestSymmetricDifference1(t *testing.T) {
	var s1, s2 IntSet
	s1.Add(1)
	s1.Add(10)
	s1.Add(30)
	s1.Add(90)
	s1.Add(100)

	s2.Add(1)
	s2.Add(10)
	s2.Add(20)
	s2.Add(100)
	s2.Add(110)

	var repr string
	expected := "{20 30 90 110}"

	s1.SymetricDifferenceWith(&s2)
	if repr = s1.String(); repr != expected {
		t.Errorf("s1.String() = %v, wants %v", repr, expected)
	}
}

func TestSymmetricDifference2(t *testing.T) {
	var s1, s2 IntSet
	s1.Add(1)
	s1.Add(10)

	s2.Add(10)
	s2.Add(100)
	s2.Add(110)

	var repr string
	expected := "{1 100 110}"

	s1.SymetricDifferenceWith(&s2)
	if repr = s1.String(); repr != expected {
		t.Errorf("s1.String() = %v, wants %v", repr, expected)
	}
}

func TestSymmetricDifference3(t *testing.T) {
	var s1, s2 IntSet
	s1.Add(1)
	s1.Add(10)
	s1.Add(100)
	s1.Add(110)
	s1.Add(210)

	s2.Add(110)

	var repr string
	expected := "{1 10 100 210}"

	s1.SymetricDifferenceWith(&s2)
	if repr = s1.String(); repr != expected {
		t.Errorf("s1.String() = %v, wants %v", repr, expected)
	}
}
