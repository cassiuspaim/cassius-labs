// Package workerpool implements a bounded worker pool pattern.
// A fixed number of goroutines consume from a shared job channel,
// preventing unbounded goroutine creation under high load.
package workerpool

import (
	"fmt"
	"sync"
)

// Job is a unit of work to be processed.
type Job struct {
	ID    int
	Input string
}

// Result carries the output of a processed Job.
type Result struct {
	JobID  int
	Output string
	Err    error
}

// Run launches numWorkers goroutines that consume from jobs.
// It returns a channel of Results that is closed when all workers finish.
// The caller must drain the result channel to avoid blocking workers.
func Run(numWorkers int, jobs []Job) <-chan Result {
	jobCh := make(chan Job, len(jobs))
	resultCh := make(chan Result, len(jobs))
	var wg sync.WaitGroup

	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobCh {
				resultCh <- process(job)
			}
		}()
	}

	for _, job := range jobs {
		jobCh <- job
	}
	close(jobCh) // signals workers: no more jobs

	go func() {
		wg.Wait()
		close(resultCh) // signals caller: no more results
	}()

	return resultCh
}

func process(job Job) Result {
	return Result{
		JobID:  job.ID,
		Output: fmt.Sprintf("processed: %s", job.Input),
	}
}
