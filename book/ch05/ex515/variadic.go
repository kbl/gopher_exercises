package ex515

import (
	"errors"
	"math"
)

func Max(items ...int) int {
	max := math.MinInt32
	for _, i := range items {
		if i > max {
			max = i
		}
	}
	return max
}

func MaxAtLeastOne(items ...int) (int, error) {
	if len(items) == 0 {
		return 0, errors.New("At least one argument is required!")
	}
	return Max(items...), nil
}

func Min(items ...int) int {
	min := math.MaxInt32
	for _, i := range items {
		if i < min {
			min = i
		}
	}
	return min
}

func MinAtLeastOne(items ...int) (int, error) {
	if len(items) == 0 {
		return 0, errors.New("At least one argument is required!")
	}
	return Min(items...), nil
}
