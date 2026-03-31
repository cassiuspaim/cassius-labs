package workerpool_test

import (
	"testing"

	"github.com/cassiuspaim/go-concurrency-advanced/workerpool"
)

func TestRun(t *testing.T) {
	jobs := make([]workerpool.Job, 20)
	for i := range jobs {
		jobs[i] = workerpool.Job{ID: i, Input: "item"}
	}

	results := workerpool.Run(4, jobs)
	var count int
	for r := range results {
		if r.Err != nil {
			t.Fatalf("unexpected error for job %d: %v", r.JobID, r.Err)
		}
		count++
	}

	if count != len(jobs) {
		t.Fatalf("expected %d results, got %d", len(jobs), count)
	}
}
