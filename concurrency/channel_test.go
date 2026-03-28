package concurrency

import (
	"testing"
	"time"
)

func TestSimpleChannel(t *testing.T) {
	ch := make(chan string, 1)
	go func() {
		ch <- "test"
	}()

	select {
	case msg := <-ch:
		if msg != "test" {
			t.Errorf("Expected 'test', got '%s'", msg)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Channel receive timed out")
	}
}

func TestBufferedChannel(t *testing.T) {
	ch := make(chan int, 3)
	for i := 1; i <= 3; i++ {
		ch <- i
	}

	for i := 1; i <= 3; i++ {
		select {
		case v := <-ch:
			if v != i {
				t.Errorf("Expected %d, got %d", i, v)
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("Buffered channel empty too early")
		}
	}
}

func TestChannelWithSelect(t *testing.T) {
	result := ChannelWithSelect()
	if result != "from ch1" && result != "from ch2" {
		t.Errorf("Unexpected result: %s", result)
	}
}

func TestProducerConsumer(t *testing.T) {
	sum := ProducerConsumer()
	expected := 1 + 2 + 3 + 4 + 5
	if sum != expected {
		t.Errorf("ProducerConsumer() = %d, want %d", sum, expected)
	}
}

func TestChannelClose(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	count := 0
	for range ch {
		count++
	}
	if count != 2 {
		t.Errorf("Expected 2 values, got %d", count)
	}
}
