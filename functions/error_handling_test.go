package functions

import (
	"errors"
	"testing"
)

func TestValidateAge(t *testing.T) {
	tests := []struct {
		name     string
		age      int
		expected error
	}{
		{"valid age", 25, nil},
		{"boundary age 0", 0, nil},
		{"boundary age 150", 150, nil},
		{"negative age", -1, errors.New("age cannot be negative")},
		{"too old", 151, errors.New("age is too large")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAge(tt.age)
			if tt.expected == nil && err != nil {
				t.Errorf("ValidateAge(%d) unexpected error: %v", tt.age, err)
			}
			if tt.expected != nil && (err == nil || err.Error() != tt.expected.Error()) {
				t.Errorf("ValidateAge(%d) = %v, want %v", tt.age, err, tt.expected)
			}
		})
	}
}

func TestValidateAgeWithCustom(t *testing.T) {
	err := ValidateAgeWithCustom(-1)
	if err == nil {
		t.Error("Expected error for negative age")
	}

	var invalidErr ErrInvalidAge
	if !errors.As(err, &invalidErr) {
		t.Errorf("Expected ErrInvalidAge, got %T", err)
	}
}

func TestSafeDivide(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"normal", 10, 2, 5},
		{"zero denominator", 5, 0, 0},
		{"negative", -10, 2, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SafeDivide(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("SafeDivide(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
