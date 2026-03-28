package stdlib

import (
	"fmt"
	"time"
)

func TimeOperations() string {
	now := time.Now()
	unix := now.Unix()

	t := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	formatted := t.Format("2006-01-02 15:04:05")

	parsed, _ := time.Parse("2006-01-02", "2024-06-20")
	year, month, day := parsed.Date()

	return fmt.Sprintf("Now: %v, Unix: %d, Formatted: %s, Parsed: %d-%02d-%02d",
		now, unix, formatted, year, month, day)
}

func DurationOperations() string {
	d1 := 5 * time.Minute
	d2 := 120 * time.Second

	total := d1 + d2
	half := total / 2

	return fmt.Sprintf("5min + 120sec = %v, half = %v", total, half)
}

func TimerAndTicker() int {
	counter := 0

	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	done := time.After(50 * time.Millisecond)

loop:
	for {
		select {
		case <-ticker.C:
			counter++
		case <-done:
			break loop
		}
	}

	return counter
}

func SleepExample() int {
	start := time.Now()
	time.Sleep(10 * time.Millisecond)
	elapsed := time.Since(start)

	if elapsed > 0 {
		return 1
	}
	return 0
}
