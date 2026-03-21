package goroutines_test

import (
	"testing"

	"github.com/cassiuspaim/go-concurrency-primitives/goroutines"
)

func TestRunOrders(t *testing.T) {
	// Run with: go test -race ./goroutines/
	goroutines.RunOrders(10)
}
