package concurrency

import (
	"sync"
)

type Counter struct {
	mu sync.Mutex
	n  int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

func (c *Counter) IncrementWithWG(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		c.Increment()
	}
}

func RaceConditionDemo() {
	c := NewCounter()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go c.IncrementWithWG(&wg)
	}
	wg.Wait()
}

type AtomicCounter struct {
	n int64
}

func NewAtomicCounter() *AtomicCounter {
	return &AtomicCounter{}
}

func (a *AtomicCounter) Increment() {
	a.n++
}

func (a *AtomicCounter) Value() int64 {
	return a.n
}
