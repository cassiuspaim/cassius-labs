# go-concurrency-primitives

Example project for the Medium article **"Go Concurrency Primitives in Production"** by Cassius Paim.

## Requirements

- Go 1.22 or later

## Packages

```
concurrency-primitives/
├── goroutines/   — basic goroutine lifecycle with WaitGroup
├── channels/     — unbuffered and buffered channel patterns (pipeline, transfer)
├── store/        — thread-safe key-value store with sync.RWMutex
└── counter/      — lock-free counter with sync/atomic typed API (Go 1.19+)
```

## Running tests with race detector

```bash
go test -race ./...
```

All packages are designed to pass cleanly under `-race`.
