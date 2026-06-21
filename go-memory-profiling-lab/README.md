# go-memory-profiling-lab

A companion project for the Medium article **Hands-on Memory Profiling and Optimization in Go**.

The project targets Go 1.26 and keeps all article examples in one module:

- Benchmark allocation output with `go test -bench=. -benchmem ./...`.
- Memory profile generation for a single benchmark package with `go test -bench=. -benchmem -memprofile=mem.out ./internal/processor`.
- Local pprof exploration with `top`, `list`, interactive mode, graph views, and the HTTP UI.
- `strings.Builder` and `strconv.Itoa` for explicit string construction.
- Slice preallocation with `make([]T, 0, len(input))`.
- Intentional prefix copying to avoid retaining a large backing array.
- Streaming JSON with `encoding/json.Decoder` and `encoding/json.Encoder`.
- Temporary buffer reuse with `sync.Pool`.
- Runtime heap and goroutine profiling with `net/http/pprof`.

## Project layout

```text
.
├── Makefile
├── README.md
├── go.mod
├── cmd/
│   └── server/
│       └── main.go
└── internal/
    └── processor/
        ├── pool.go
        ├── pool_test.go
        ├── processor.go
        ├── processor_test.go
        ├── stream.go
        └── stream_test.go
```

## Run tests

```bash
go test ./...
```

Or with Make:

```bash
make test
```

## Run benchmarks across the module

```bash
go test -bench=. -benchmem ./...
```

Or with Make:

```bash
make bench
```

`-bench=.` runs all benchmarks matched by the regular expression `.`. `-benchmem` adds allocation columns such as `B/op` and `allocs/op` to the benchmark output.

The benchmark set includes:

- `BenchmarkFormatUsersSlow`
- `BenchmarkFormatUsersBuilder`
- `BenchmarkNamesSlow`
- `BenchmarkNamesPreallocated`
- `BenchmarkPrefixView`
- `BenchmarkPrefixCopy`
- `BenchmarkEncodeRecordPlain`
- `BenchmarkEncodeRecordWithPool`
- `BenchmarkFilterActiveUsersInMemory`
- `BenchmarkFilterActiveUsersStreaming`

## Generate a memory profile

`-memprofile` writes one profile file, so use it against one package instead of `./...`:

```bash
go test -bench=. -benchmem -memprofile=mem.out ./internal/processor
```

Or with Make:

```bash
make memprofile
```

The benchmark package can be overridden:

```bash
make memprofile BENCH_PKG=./internal/processor
```

## Read the memory profile with pprof

Start with the textual top view:

```bash
go tool pprof -top mem.out
```

Or with Make:

```bash
make pprof-top
```

The most important columns in the `top` output are:

- `flat`: memory attributed directly to that function.
- `flat%`: direct contribution as a percentage of the selected profile.
- `sum%`: running total of `flat%` up to the current row.
- `cum`: cumulative memory attributed to the function plus the functions it calls.
- `cum%`: cumulative contribution as a percentage of the selected profile.

For source-level inspection, open the interactive shell:

```bash
go tool pprof mem.out
```

Then run commands such as:

```text
(pprof) top
(pprof) list FormatUsersSlow
(pprof) list FormatUsersBuilder
(pprof) list EncodeRecordWithPool
(pprof) list FilterActiveUsersStreaming
(pprof) help
(pprof) quit
```

`list` connects profile data back to source lines, which is useful when a function contains multiple possible allocation sites.

## Graph views and Graphviz

The interactive command:

```text
(pprof) web
```

renders a call graph. It requires Graphviz because `pprof` uses the `dot` executable to render graph-based views.

On macOS:

```bash
brew install graphviz
```

On Debian or Ubuntu:

```bash
sudo apt-get update
sudo apt-get install graphviz
```

Verify:

```bash
dot -V
```

A browser UI can also be started with:

```bash
go tool pprof -http=:8080 mem.out
```

Or with Make:

```bash
make pprof-http
```

Some graph views in the browser UI may still require Graphviz.

## What to observe

### String construction

`FormatUsersSlow` uses repeated string concatenation and `fmt.Sprintf`. `FormatUsersBuilder` uses `strings.Builder`, `Grow`, and `strconv.Itoa` to build the same text more explicitly.

### Slice preallocation

`NamesSlow` lets `append` grow the slice. `NamesPreallocated` allocates the expected output capacity up front with `make([]string, 0, len(users))`.

### Slice retention

`PrefixView` returns a small view into a potentially large backing array. `PrefixCopy` intentionally allocates a small new backing array so the large original array can be released when nothing else references it.

### sync.Pool

`EncodeRecordPlain` creates a temporary buffer on each call. `EncodeRecordWithPool` reuses a temporary `bytes.Buffer`, resets it before use, copies the returned result to give the caller independent ownership, and discards oversized buffers instead of returning them to the pool.

### Streaming JSON

`FilterActiveUsersInMemory` decodes the full JSON array and builds a full output slice before writing. `FilterActiveUsersStreaming` processes one JSON element at a time and writes matching output progressively.

## Run the example server

```bash
go run ./cmd/server
```

In another terminal:

```bash
curl "http://localhost:8080/work?n=1000"
go tool pprof "http://localhost:8080/debug/pprof/heap"
go tool pprof "http://localhost:8080/debug/pprof/goroutine"
```

Do not expose pprof endpoints publicly in production. Bind them to trusted interfaces or protect them with appropriate access controls.
