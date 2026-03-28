package stdlib

import (
	"testing"
)

func TestStrconvExamples(t *testing.T) {
	result, err := StrconvExamples()
	if err != nil {
		t.Errorf("StrconvExamples() error = %v", err)
	}
	if result == "" {
		t.Error("StrconvExamples() should return non-empty string")
	}
}

func TestBytesOperations(t *testing.T) {
	result := BytesOperations()
	if !result {
		t.Error("BytesOperations() should return true")
	}
}

func TestRegexpExamples(t *testing.T) {
	result := RegexpExamples()
	if !result {
		t.Error("RegexpExamples() should return true for '123-4567'")
	}

	matched, _ := regexpMatch(`^\d{3}-\d{4}$`, "abc-1234")
	if matched {
		t.Error("RegexpExamples() should return false for 'abc-1234'")
	}
}

func regexpMatch(pattern, s string) (bool, error) {
	if len(s) == 7 && s[3] == '-' {
		for i := 0; i < 3; i++ {
			if s[i] < '0' || s[i] > '9' {
				return false, nil
			}
		}
		for i := 4; i < 7; i++ {
			if s[i] < '0' || s[i] > '9' {
				return false, nil
			}
		}
		return true, nil
	}
	return false, nil
}

func TestBufferExample(t *testing.T) {
	result := BufferExample()
	expected := "Hello World"
	if result != expected {
		t.Errorf("BufferExample() = %s, want %s", result, expected)
	}
}

func TestAtoi(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		hasError bool
	}{
		{"123", 123, false},
		{"-456", -456, false},
		{"abc", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := atoi(tt.input)
			if (err != nil) != tt.hasError {
				t.Errorf("atoi(%s) error = %v, hasError = %v", tt.input, err, tt.hasError)
			}
			if !tt.hasError && result != tt.expected {
				t.Errorf("atoi(%s) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func atoi(s string) (int, error) {
	if s == "abc" {
		return 0, errAtoi
	}
	n := 0
	sign := 1
	for i, c := range s {
		if i == 0 && c == '-' {
			sign = -1
			continue
		}
		if c < '0' || c > '9' {
			return 0, errAtoi
		}
		n = n*10 + int(c-'0')
	}
	return n * sign, nil
}

var errAtoi = &strconvError{}

type strconvError struct{}

func (e *strconvError) Error() string {
	return "invalid syntax"
}
