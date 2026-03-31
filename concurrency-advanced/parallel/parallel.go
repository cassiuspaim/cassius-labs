// Package parallel demonstrates errgroup for concurrent work with error coordination.
package parallel

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

// FetchAll fetches multiple resources concurrently.
// If any fetch fails, the context is cancelled and all remaining fetches stop.
// Returns all results or the first error encountered.
func FetchAll(ctx context.Context, ids []int) ([]string, error) {
	g, ctx := errgroup.WithContext(ctx)
	results := make([]string, len(ids))

	for i, id := range ids {
		i, id := i, id // safe in Go 1.22+ but explicit capture is clearer
		g.Go(func() error {
			data, err := fetch(ctx, id)
			if err != nil {
				return fmt.Errorf("fetch id %d: %w", id, err)
			}
			results[i] = data // safe: each goroutine writes a unique index
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}

func fetch(ctx context.Context, id int) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		if id < 0 {
			return "", fmt.Errorf("invalid id: %d", id)
		}
		return fmt.Sprintf("data-%d", id), nil
	}
}
