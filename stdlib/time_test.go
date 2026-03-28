package stdlib

import (
	"strings"
	"testing"
	"time"
)

func TestTimeOperations(t *testing.T) {
	result := TimeOperations()
	if result == "" {
		t.Error("TimeOperations() should return non-empty string")
	}
}

func TestDurationOperations(t *testing.T) {
	result := DurationOperations()
	if result == "" {
		t.Error("DurationOperations() should return non-empty string")
	}
}

func TestTimerAndTicker(t *testing.T) {
	counter := TimerAndTicker()
	if counter == 0 {
		t.Error("TimerAndTicker() should return non-zero counter")
	}
}

func TestSleepExample(t *testing.T) {
	result := SleepExample()
	if result != 1 {
		t.Error("SleepExample() should return 1")
	}
}

func TestTimeNow(t *testing.T) {
	now := time.Now()
	if now.IsZero() {
		t.Error("time.Now() should not be zero")
	}
}

func TestTimeFormat(t *testing.T) {
	t1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	formatted := t1.Format("2006-01-02")
	if !strings.Contains(formatted, "2024-01-01") {
		t.Errorf("Time format incorrect: %s", formatted)
	}
}
