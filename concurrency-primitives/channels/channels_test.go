package channels_test

import (
	"testing"

	"github.com/cassiuspaim/go-concurrency-primitives/channels"
)

func TestPipeline(t *testing.T) {
	results := channels.Pipeline(5)
	if len(results) != 5 {
		t.Fatalf("expected 5 results, got %d", len(results))
	}
}

func TestTransfer(t *testing.T) {
	result := channels.Transfer("hello")
	if result != "processed: hello" {
		t.Fatalf("unexpected result: %s", result)
	}
}
