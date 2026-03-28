package testing

import (
	"errors"
	"testing"
)

func TestSubTest(t *testing.T) {
	tests := []struct {
		name    string
		a, b    int
		want    int
		wantErr bool
	}{
		{"divide positive", 6, 2, 3, false},
		{"divide zero", 5, 1, 5, false},
		{"divide negative", -10, 2, -5, false},
		{"divide by zero", 10, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := divide(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("divide(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func TestTableDriven(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{5}, 5},
		{"multiple elements", []int{1, 2, 3, 4, 5}, 15},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := sum(c.input)
			if got != c.expected {
				t.Errorf("sum(%v) = %d, want %d", c.input, got, c.expected)
			}
		})
	}
}

func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func TestParallel(t *testing.T) {
	t.Run("parallel1", func(t *testing.T) {
		t.Parallel()
		for i := 0; i < 100; i++ {
			_ = i * 2
		}
	})

	t.Run("parallel2", func(t *testing.T) {
		t.Parallel()
		for i := 0; i < 100; i++ {
			_ = i * 3
		}
	})
}

func TestWithSetupTeardown(t *testing.T) {
	setup := func() func() {
		t.Log("Setup")
		return func() {
			t.Log("Teardown")
		}
	}
	teardown := setup()
	defer teardown()

	t.Log("Test running")
}

type MockDB struct {
	data map[string]int
}

func NewMockDB() *MockDB {
	return &MockDB{data: make(map[string]int)}
}

func (m *MockDB) Get(key string) int {
	return m.data[key]
}

func (m *MockDB) Set(key string, value int) {
	m.data[key] = value
}

func TestMockDB(t *testing.T) {
	db := NewMockDB()
	db.Set("count", 10)

	if got := db.Get("count"); got != 10 {
		t.Errorf("Get('count') = %d, want 10", got)
	}
}

func BenchmarkSum(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum(nums)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
		quickSort(data)
	}
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	var left, mid, right []int

	for _, v := range arr {
		switch {
		case v < pivot:
			left = append(left, v)
		case v == pivot:
			mid = append(mid, v)
		default:
			right = append(right, v)
		}
	}

	result := append(quickSort(left), mid...)
	result = append(result, quickSort(right)...)
	return result
}
