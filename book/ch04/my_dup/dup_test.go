package my_dup

import (
	"reflect"
	"testing"
)

func TestDup(t *testing.T) {
	s1 := []string{"a", "a", "a", "b", "c", "c", "d"}
	input := copy(make([]string, len(s1)), s1)
	output := Dup(s1)
	expected := []string{"a", "b", "c", "d"}

	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Dup(%v) = %v, want %v", input, output, expected)
	}
}

func TestDupWhitespace(t *testing.T) {
	b1 := []byte("żó   \t\n  łć")
	input := copy(make([]byte, len(b1)), b1)
	output := DupWhitespace(b1)
	expected := []byte("żó łć")

	if !reflect.DeepEqual(output, expected) {
		t.Errorf("DupWhitespace(%v) = %v, want %v", input, output, expected)
	}
}
