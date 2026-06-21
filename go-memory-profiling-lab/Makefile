BENCH_PKG ?= ./internal/processor
PROFILE ?= mem.out
HTTP_ADDR ?= :8080

.PHONY: test bench memprofile pprof-top pprof-interactive pprof-http run clean

test:
	go test ./...

bench:
	go test -bench=. -benchmem ./...

memprofile:
	go test -bench=. -benchmem -memprofile=$(PROFILE) $(BENCH_PKG)

pprof-top:
	go tool pprof -top $(PROFILE)

pprof-interactive:
	go tool pprof $(PROFILE)

pprof-http:
	go tool pprof -http=$(HTTP_ADDR) $(PROFILE)

run:
	go run ./cmd/server

clean:
	rm -f *.out *.prof
