package ex515

import (
	"math"
	"testing"
)

func TestMax(t *testing.T) {
	parameters := []int{1, 2, 4, 3}
	result := Max(parameters...)

	if result != 4 {
		t.Errorf("Max(%v) = %v, want %v", parameters, result, 4)
	}
}

func TestMaxWhenEmpty(t *testing.T) {
	var parameters []int
	result := Max(parameters...)

	if result != math.MinInt32 {
		t.Errorf("Max(%v) = %v, want %v", parameters, result, math.MinInt32)
	}
}

func TestMaxWithZeroArguments(t *testing.T) {
	var parameters []int
	_, err := MaxAtLeastOne(parameters...)

	if err == nil {
		t.Error("Error should be returned!")
	}
}
