package concurrency

import (
	"fmt"
	"sync"
)

func RunAll() {
	println("=== Simple Goroutine ===")
	SimpleGoroutine()

	println("\n=== Multiple Goroutines ===")
	MultipleGoroutines()

	println("\n=== Channel Basics ===")
	SimpleChannel()
	BufferedChannel()
	println("Channels work!")

	println("\n=== Select ===")
	msg := ChannelWithSelect()
	fmt.Println("Received:", msg)

	println("\n=== Producer Consumer ===")
	sum := ProducerConsumer()
	fmt.Println("Sum:", sum)

	println("\n=== Sync Mutex ===")
	RaceConditionDemo()
	c := NewCounter()
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				c.Increment()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", c.Value())
}
