package testing

import (
	"strings"
	"sync"
	"testing"
	"time"
)

func BenchmarkSimpleAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = 1 + 2
	}
}

func BenchmarkStringConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ""
		for j := 0; j < 100; j++ {
			s += "a"
		}
		_ = s
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		for j := 0; j < 100; j++ {
			sb.WriteString("a")
		}
		_ = sb.String()
	}
}

func BenchmarkMapAccess(b *testing.B) {
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m[500]
	}
}

func BenchmarkSliceAccess(b *testing.B) {
	s := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		s[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s[500]
	}
}

func BenchmarkMutex(b *testing.B) {
	var mu sync.Mutex
	counter := 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

func BenchmarkConcurrentgoroutine(b *testing.B) {
	var wg sync.WaitGroup
	b.ResetTimer()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				_ = j * 2
			}
		}()
	}
	wg.Wait()
}

func BenchmarkSleep(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Nanosecond)
	}
}

func BenchmarkChannel(b *testing.B) {
	ch := make(chan int, 100)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				ch <- j
			}
		}()
	}

	b.ResetTimer()
	go func() {
		wg.Wait()
		close(ch)
	}()

	for range ch {
	}
}
