package multiplex_test

import (
	"context"
	"sync"
	"testing"

	"github.com/cassiuspaim/go-concurrency-advanced/multiplex"
)

func TestWorker_processesJobs(t *testing.T) {
	ctx := context.Background()
	jobs := make(chan multiplex.Job, 5)
	results := make(chan multiplex.Result, 5)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		multiplex.Worker(ctx, 1, jobs, results)
	}()

	for i := range 5 {
		jobs <- multiplex.Job{ID: i}
	}
	close(jobs)

	wg.Wait()
	close(results)

	var count int
	for range results {
		count++
	}
	if count != 5 {
		t.Fatalf("expected 5 results, got %d", count)
	}
}

func TestWorker_respectsCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	jobs := make(chan multiplex.Job) // unbuffered: worker will block on select
	results := make(chan multiplex.Result, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		multiplex.Worker(ctx, 1, jobs, results)
	}()

	cancel()  // cancel before sending any jobs
	wg.Wait() // worker should exit cleanly
}
