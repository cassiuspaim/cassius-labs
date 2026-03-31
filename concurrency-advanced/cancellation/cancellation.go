// Package cancellation demonstrates context.Context propagation patterns.
package cancellation

import (
	"context"
	"fmt"
	"time"
)

// FetchRemote simulates a slow external call that respects context cancellation.
// Returns an error if the context is cancelled before the work completes.
func FetchRemote(ctx context.Context, id int, latency time.Duration) (string, error) {
	select {
	case <-time.After(latency):
		return fmt.Sprintf("result-%d", id), nil
	case <-ctx.Done():
		// ctx.Err() is context.Canceled or context.DeadlineExceeded
		return "", fmt.Errorf("fetch %d: %w", id, ctx.Err())
	}
}

// FetchWithTimeout wraps FetchRemote with a per-call timeout.
// Always defers cancel to prevent context leak.
func FetchWithTimeout(id int, timeout time.Duration, latency time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel() // prevents context leak even on early return
	return FetchRemote(ctx, id, latency)
}
