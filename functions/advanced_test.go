package functions

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"multiple numbers", []int{1, 2, 3, 4, 5}, 15},
		{"single number", []int{10}, 10},
		{"empty", []int{}, 0},
		{"with zero", []int{1, 0, 1}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sum(tt.nums...)
			if result != tt.expected {
				t.Errorf("Sum(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestMakeMultiplier(t *testing.T) {
	double := MakeMultiplier(2)
	triple := MakeMultiplier(3)

	tests := []struct {
		name     string
		fn       func(int) int
		input    int
		expected int
	}{
		{"double 5", double, 5, 10},
		{"double 0", double, 0, 0},
		{"triple 4", triple, 4, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fn(tt.input)
			if result != tt.expected {
				t.Errorf("MakeMultiplier()(%d) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestApply(t *testing.T) {
	nums := []int{1, 2, 3}
	square := func(n int) int { return n * n }

	result := Apply(nums, square)
	expected := []int{1, 4, 9}

	if !slices.Equal(result, expected) {
		t.Errorf("Apply(%v, square) = %v, want %v", nums, result, expected)
	}
}
