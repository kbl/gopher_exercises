package ex704

import (
	"testing"
)

func TestNewReader(t *testing.T) {
	buffer := make([]byte, 5)
	r := NewReader("1234567")

	c, _ := r.Read(buffer)
	if c != 5 {
		t.Errorf("%v != 5", c)
	}

	c, _ = r.Read(buffer)
	if c != 2 {
		t.Errorf("%v != 2", c)
	}
}
