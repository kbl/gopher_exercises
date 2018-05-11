package rev

import (
	"reflect"
	"testing"
)

func TestRevEmptySlice(t *testing.T) {
	input := []int{}
	before := []int{}
	expected := []int{}
	Rev(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Rev(%v) = %v, want %v", before, input, expected)
	}
}

func TestRevNilSlice(t *testing.T) {
	var input []int
	var before []int
	var expected []int
	Rev(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Rev(%v) = %v, want %v", before, input, expected)
	}
}

func TestRevOddLengthSlice(t *testing.T) {
	input := []int{1, 2, 3}
	before := []int{1, 2, 3}
	expected := []int{3, 2, 1}
	Rev(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Rev(%v) = %v, want %v", before, input, expected)
	}
}

func TestRevEvenLengthSlice(t *testing.T) {
	input := []int{1, 2, 3, 4}
	before := []int{1, 2, 3, 4}
	expected := []int{4, 3, 2, 1}
	Rev(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Rev(%v) = %v, want %v", before, input, expected)
	}
}
