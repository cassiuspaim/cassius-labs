// Package channels demonstrates unbuffered and buffered channel patterns.
package channels

import (
	"fmt"
	"sync"
)

// Pipeline connects a producer and a consumer through typed, directional channels.
// The producer sends squares of integers; the consumer collects results.
func Pipeline(count int) []int {
	jobs := make(chan int, count)
	results := make(chan int, count)
	var wg sync.WaitGroup

	// producer: sends integers and closes the channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range count {
			jobs <- i
		}
		close(jobs) // signals the consumer that no more values will arrive
	}()

	// consumer: reads until jobs is closed and drained
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range jobs {
			results <- v * v
		}
	}()

	// close results once both goroutines are done
	go func() {
		wg.Wait()
		close(results)
	}()

	var out []int
	for r := range results {
		out = append(out, r)
	}
	return out
}

// Transfer demonstrates an unbuffered channel: sender and receiver rendezvous.
func Transfer(payload string) string {
	ch := make(chan string) // unbuffered: blocks until both sides are ready
	go func() {
		ch <- fmt.Sprintf("processed: %s", payload)
	}()
	return <-ch
}
