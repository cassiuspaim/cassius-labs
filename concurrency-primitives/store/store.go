// Package store demonstrates a thread-safe key-value store using sync.RWMutex.
package store

import "sync"

// Store is a concurrent-safe in-memory key-value store.
// RWMutex is used because reads are expected to outnumber writes.
type Store struct {
	mu   sync.RWMutex
	data map[string]string
}

// New returns an initialized Store.
func New() *Store {
	return &Store{data: make(map[string]string)}
}

// Set writes a key-value pair under an exclusive write lock.
func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

// Get reads a value under a shared read lock.
// Multiple goroutines may call Get concurrently without blocking each other.
func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.data[key]
	return v, ok
}

// Delete removes a key under an exclusive write lock.
func (s *Store) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}
