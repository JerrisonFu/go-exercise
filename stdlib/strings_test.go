package stdlib

import (
	"slices"
	"testing"
)

func TestStringOperations(t *testing.T) {
	result := StringOperations()
	if result == "" {
		t.Error("StringOperations should return non-empty string")
	}
}

func TestSortOperations(t *testing.T) {
	nums := SortOperations()
	expected := []int{1, 2, 5, 8, 9}
	if !slices.Equal(nums, expected) {
		t.Errorf("SortOperations() = %v, want %v", nums, expected)
	}
}

func TestSliceOperations(t *testing.T) {
	result := SliceOperations()
	if !result {
		t.Error("SliceOperations should return true")
	}
}

func TestStringsContains(t *testing.T) {
	if !contains("hello world", "world") {
		t.Error("Should contain 'world'")
	}
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
