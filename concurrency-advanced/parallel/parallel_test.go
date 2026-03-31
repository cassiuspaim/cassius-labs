package parallel_test

import (
	"context"
	"testing"

	"github.com/cassiuspaim/go-concurrency-advanced/parallel"
)

func TestFetchAll_success(t *testing.T) {
	results, err := parallel.FetchAll(context.Background(), []int{1, 2, 3, 4, 5})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 5 {
		t.Fatalf("expected 5 results, got %d", len(results))
	}
}

func TestFetchAll_error(t *testing.T) {
	_, err := parallel.FetchAll(context.Background(), []int{1, -1, 3})
	if err == nil {
		t.Fatal("expected error for negative id, got nil")
	}
}
