// Package multiplex demonstrates select-based channel multiplexing with context.
package multiplex

import (
	"context"
	"fmt"
)

// Job and Result are simple types for the worker example.
type Job struct{ ID int }
type Result struct {
	JobID  int
	Output string
}

// Worker consumes jobs until the jobs channel is closed or ctx is cancelled.
// It demonstrates the correct select pattern for a clean goroutine exit path.
func Worker(ctx context.Context, id int, jobs <-chan Job, results chan<- Result) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				// channel closed: no more jobs, exit cleanly
				return
			}
			results <- Result{
				JobID:  job.ID,
				Output: fmt.Sprintf("worker %d processed job %d", id, job.ID),
			}
		case <-ctx.Done():
			// context cancelled: stop immediately
			fmt.Printf("worker %d stopping: %v\n", id, ctx.Err())
			return
		}
	}
}
