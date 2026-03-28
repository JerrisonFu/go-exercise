package concurrency

func SimpleChannel() {
	ch := make(chan string)

	go func() {
		ch <- "Hello from channel"
	}()

	msg := <-ch
	_ = msg
}

func BufferedChannel() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
}

func ChannelWithSelect() string {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "from ch1"
	}()
	go func() {
		ch2 <- "from ch2"
	}()

	select {
	case msg := <-ch1:
		return msg
	case msg := <-ch2:
		return msg
	}
}

func ProducerConsumer() int {
	ch := make(chan int, 5)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}
