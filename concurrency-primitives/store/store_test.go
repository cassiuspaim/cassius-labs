package store_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/cassiuspaim/go-concurrency-primitives/store"
)

func TestStoreConcurrent(t *testing.T) {
	s := store.New()
	var wg sync.WaitGroup

	// concurrent writers
	for i := range 50 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("val%d", i))
		}(i)
	}

	// concurrent readers (may run before writes complete — that's fine)
	for i := range 50 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Get(fmt.Sprintf("key%d", i))
		}(i)
	}

	wg.Wait()
}
