package functions

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative number", -1, 1, 0},
		{"zeros", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	if Multiply(3, 4) != 12 {
		t.Error("Multiply(3, 4) should be 12")
	}
	if Multiply(-2, 5) != -10 {
		t.Error("Multiply(-2, 5) should be -10")
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name      string
		a, b      int
		expected  int
		expectErr bool
	}{
		{"normal division", 10, 2, 5, false},
		{"division by zero", 5, 0, 0, true},
		{"floor division", 7, 2, 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)
			if (err != nil) != tt.expectErr {
				t.Errorf("Divide(%d, %d) error = %v, expectErr %v", tt.a, tt.b, err, tt.expectErr)
				return
			}
			if !tt.expectErr && result != tt.expected {
				t.Errorf("Divide(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDivideByZeroError(t *testing.T) {
	_, err := Divide(5, 0)
	if err == nil {
		t.Error("Expected error for division by zero")
	}
	var dzb ErrDivisionByZero
	if !errors.As(err, &dzb) {
		t.Errorf("Expected ErrDivisionByZero, got %T", err)
	}
}
