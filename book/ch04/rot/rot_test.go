package rot

import (
	"reflect"
	"testing"
)

func TestRotEmptySlice(t *testing.T) {
	input := []int{}
	before := []int{}
	expected := []int{}
	Rot(input, 2)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Rot(%v) = %v, want %v", before, input, expected)
	}
}

func TestRotNilSlice(t *testing.T) {
	var input []int
	var before []int
	var expected []int
	Rot(input, 2)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Rot(%v) = %v, want %v", before, input, expected)
	}
}

func TestRotOddLengthSlice(t *testing.T) {
	input := []int{1, 2, 3}
	before := []int{1, 2, 3}
	expected := []int{3, 1, 2}
	Rot(input, 2)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Rot(%v) = %v, want %v", before, input, expected)
	}
}

func TestRotEvenLengthSlice(t *testing.T) {
	input := []int{1, 2, 3, 4}
	before := []int{1, 2, 3, 4}
	expected := []int{3, 4, 1, 2}
	Rot(input, 2)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Rot(%v) = %v, want %v", before, input, expected)
	}
}
