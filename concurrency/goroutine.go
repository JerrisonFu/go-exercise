package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func SimpleGoroutine() {
	done := make(chan bool)
	go func() {
		fmt.Println("Running in goroutine")
		time.Sleep(100 * time.Millisecond)
		done <- true
	}()
	<-done
	fmt.Println("Goroutine finished")
}

func MultipleGoroutines() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d started\n", id)
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("Goroutine %d finished\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Println("All goroutines finished")
}
