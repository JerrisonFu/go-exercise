package concurrency

import (
	"sync"
	"testing"
)

func TestNewCounter(t *testing.T) {
	c := NewCounter()
	if c.Value() != 0 {
		t.Error("New counter should start at 0")
	}
}

func TestCounterIncrement(t *testing.T) {
	c := NewCounter()
	c.Increment()
	if c.Value() != 1 {
		t.Error("After increment, counter should be 1")
	}
}

func TestCounterConcurrent(t *testing.T) {
	c := NewCounter()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go c.IncrementWithWG(&wg)
	}
	wg.Wait()

	if c.Value() != 10000 {
		t.Errorf("Expected 10000, got %d", c.Value())
	}
}

func TestAtomicCounter(t *testing.T) {
	c := NewAtomicCounter()
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				c.Increment()
			}
		}()
	}
	wg.Wait()

	if c.Value() != 5000 {
		t.Errorf("Expected 5000, got %d", c.Value())
	}
}

func TestRaceConditionDemo(t *testing.T) {
	RaceConditionDemo()
}
