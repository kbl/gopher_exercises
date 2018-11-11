package intset

import (
	"testing"
)

func TestLen(t *testing.T) {
	var s IntSet
	len := 0

	if len = s.Len(); len != 0 {
		t.Errorf("s.Len() = %v, wants %v", len, 0)
	}

	s.Add(111)
	if len = s.Len(); len != 1 {
		t.Errorf("s.Len() = %v, wants %v", len, 1)
	}

	s.Add(111)
	s.Add(222)
	if len = s.Len(); len != 2 {
		t.Errorf("s.Len() = %v, wants %v", len, 2)
	}
}

func TestRemove(t *testing.T) {
	var s IntSet
	has := false
	s.Remove(10)

	s.Add(10)
	s.Add(20)
	s.Remove(10)
	if has = s.Has(10); has {
		t.Errorf("s.Has(%v) = %v, wants %v", 10, has, false)
	}
	if has = s.Has(20); !has {
		t.Errorf("s.Has(%v) = %v, wants %v", 20, has, true)
	}
}

func TestClear(t *testing.T) {
	var s IntSet
	len := 0
	s.Add(10)
	s.Add(120)
	if len = s.Len(); len != 2 {
		t.Errorf("s.Len() = %v, wants %v", len, 2)
	}

	s.Clear()
	if len = s.Len(); len != 0 {
		t.Errorf("s.Len() = %v, wants %v", len, 0)
	}
}

func TestCopy(t *testing.T) {
	var has bool
	var orig IntSet
	orig.Add(10)
	orig.Add(120)

	copy := orig.Copy()
	copy.Remove(120)

	if has = orig.Has(10); !has {
		t.Errorf("orig.Has(%v) = %v, wants %v", 10, has, true)
	}
	if has = copy.Has(10); !has {
		t.Errorf("copy.Has(%v) = %v, wants %v", 10, has, true)
	}
	if has = orig.Has(120); !has {
		t.Errorf("orig.Has(%v) = %v, wants %v", 120, has, true)
	}
	if has = copy.Has(120); has {
		t.Errorf("copy.Has(%v) = %v, wants %v", 120, has, false)
	}
}
