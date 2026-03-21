// Package counter demonstrates sync/atomic typed API (Go 1.19+).
package counter

import (
	"sync"
	"sync/atomic"
)

// SafeCounter uses atomic.Int64 for lock-free concurrent increments.
type SafeCounter struct {
	val atomic.Int64
}

// Inc increments the counter by 1.
func (c *SafeCounter) Inc() {
	c.val.Add(1)
}

// Value returns the current counter value.
func (c *SafeCounter) Value() int64 {
	return c.val.Load()
}

// RunConcurrent launches n goroutines each incrementing the counter once.
// Returns the final value (always n).
func RunConcurrent(n int) int64 {
	var c SafeCounter
	var wg sync.WaitGroup
	for range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	return c.Value()
}
