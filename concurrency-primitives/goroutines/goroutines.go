// Package goroutines demonstrates basic goroutine lifecycle with WaitGroup.
package goroutines

import (
	"fmt"
	"sync"
	"time"
)

// ProcessOrder simulates processing a single order concurrently.
// It receives a WaitGroup pointer and calls Done when finished.
func ProcessOrder(id int, wg *sync.WaitGroup) {
	defer wg.Done() // guaranteed to run even if the function panics
	fmt.Printf("processing order %d\n", id)
	time.Sleep(50 * time.Millisecond)
	fmt.Printf("order %d done\n", id)
}

// RunOrders launches one goroutine per order and waits for all to finish.
func RunOrders(count int) {
	var wg sync.WaitGroup
	for i := range count { // Go 1.22+: i is a new variable per iteration
		wg.Add(1)
		go ProcessOrder(i, &wg)
	}
	wg.Wait()
}
