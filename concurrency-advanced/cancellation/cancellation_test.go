package cancellation_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/cassiuspaim/go-concurrency-advanced/cancellation"
)

func TestFetchWithTimeout_success(t *testing.T) {
	result, err := cancellation.FetchWithTimeout(1, 500*time.Millisecond, 50*time.Millisecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "result-1" {
		t.Fatalf("unexpected result: %s", result)
	}
}

func TestFetchWithTimeout_timeout(t *testing.T) {
	_, err := cancellation.FetchWithTimeout(2, 50*time.Millisecond, 500*time.Millisecond)
	if err == nil {
		t.Fatal("expected timeout error, got nil")
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected DeadlineExceeded, got: %v", err)
	}
}
