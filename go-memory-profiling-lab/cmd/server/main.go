package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"

	"github.com/cassiuspaim/go-memory-profiling-lab/internal/processor"
)

func main() {
	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		n := 1000
		if raw := r.URL.Query().Get("n"); raw != "" {
			parsed, err := strconv.Atoi(raw)
			if err == nil && parsed > 0 {
				n = parsed
			}
		}

		users := make([]processor.User, 0, n)
		for i := 0; i < n; i++ {
			users = append(users, processor.User{
				Name:  fmt.Sprintf("user-%04d", i),
				Score: i % 100,
			})
		}

		result := processor.FormatUsersBuilder(users)
		_, _ = w.Write([]byte(result))
	})

	log.Println("listening on http://localhost:8080")
	log.Println("pprof endpoints are available under /debug/pprof/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
