package concurrency

import (
	"sync"
	"testing"
	"time"
)

func TestSimpleGoroutine(t *testing.T) {
	done := make(chan bool, 1)
	go func() {
		time.Sleep(10 * time.Millisecond)
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Error("Goroutine did not complete in time")
	}
}

func TestMultipleGoroutines(t *testing.T) {
	var wg sync.WaitGroup
	counter := 0

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	if counter != 5 {
		t.Errorf("Expected counter=5, got %d", counter)
	}
}
