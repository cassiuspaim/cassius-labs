# go-concurrency-advanced

Example project for the Medium article **"Advanced Go Concurrency Patterns in Production"** by Cassius Paim.

## Requirements

- Go 1.22 or later
- `golang.org/x/sync` v0.7.0 (see `go.mod`)

## Install dependencies

```bash
go mod tidy
```

## Packages

```
concurrency-advanced/
├── workerpool/    — bounded worker pool with typed jobs and results
├── cancellation/  — context.WithTimeout and cancellation propagation
├── parallel/      — errgroup for concurrent work with coordinated error handling
└── multiplex/     — select-based worker with context cancellation exit path
```

## Running tests with race detector

```bash
go test -race ./...
```

All packages pass cleanly under `-race`.
