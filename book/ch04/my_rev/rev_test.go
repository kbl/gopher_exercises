package my_rev

import (
	"reflect"
	"testing"
)

func TestReverseEmptySlice(t *testing.T) {
	input := []byte{}
	before := []byte{}
	expected := []byte{}
	Reverse(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Reverse(%v) = %v, want %v", before, input, expected)
	}
}

func TestReverseNilSlice(t *testing.T) {
	var input []byte
	var before []byte
	var expected []byte
	Reverse(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Reverse(%v) = %v, want %v", before, input, expected)
	}
}

func TestReverseOddLengthSlice(t *testing.T) {
	input := []byte{1, 2, 3}
	before := []byte{1, 2, 3}
	expected := []byte{3, 2, 1}
	Reverse(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Reverse(%v) = %v, want %v", before, input, expected)
	}
}

func TestReverseEvenLengthSlice(t *testing.T) {
	input := []byte{1, 2, 3, 4}
	before := []byte{1, 2, 3, 4}
	expected := []byte{4, 3, 2, 1}
	Reverse(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Reverse(%v) = %v, want %v", before, input, expected)
	}
}

func TestReverseArrayEvenLengthSlice(t *testing.T) {
	input := [...]int{1, 2, 3, 4}
	before := [...]int{1, 2, 3, 4}
	expected := [...]int{4, 3, 2, 1}
	ReverseArray(&input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("ReverseArray(%v) = %v, want %v", before, input, expected)
	}
}

func TestReverseUtf8(t *testing.T) {
	input := []byte("zażółć gęślą jaźń")
	before := []byte("zażółć gęślą jaźń")
	expected := []byte("ńźaj ąlśęg ćłóżaz")
	ReverseUtf8(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("ReverseUtf8(%q) = %q, want %q", before, input, expected)
	}
}
