package counter_test

import (
	"testing"

	"github.com/cassiuspaim/go-concurrency-primitives/counter"
)

func TestRunConcurrent(t *testing.T) {
	const n = 1000
	result := counter.RunConcurrent(n)
	if result != n {
		t.Fatalf("expected %d, got %d", n, result)
	}
}
