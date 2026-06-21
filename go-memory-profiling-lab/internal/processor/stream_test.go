package processor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestFilterActiveUsers(t *testing.T) {
	input := []byte(`[
		{"id":"u1","active":true,"name":"Ana"},
		{"id":"u2","active":false,"name":"Bob"},
		{"id":"u3","active":true,"name":"Carla"}
	]`)

	want := []OutputUser{{ID: "u1"}, {ID: "u3"}}

	for name, fn := range map[string]func(*bytes.Reader, *bytes.Buffer) error{
		"in-memory": func(r *bytes.Reader, w *bytes.Buffer) error {
			return FilterActiveUsersInMemory(r, w)
		},
		"streaming": func(r *bytes.Reader, w *bytes.Buffer) error {
			return FilterActiveUsersStreaming(r, w)
		},
	} {
		t.Run(name, func(t *testing.T) {
			var out bytes.Buffer
			if err := fn(bytes.NewReader(input), &out); err != nil {
				t.Fatalf("filter returned error: %v", err)
			}

			var got []OutputUser
			if err := json.Unmarshal(out.Bytes(), &got); err != nil {
				t.Fatalf("invalid JSON output %q: %v", out.String(), err)
			}

			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %+v, want %+v", got, want)
			}
		})
	}
}

func BenchmarkFilterActiveUsersInMemory(b *testing.B) {
	input := generateUsersJSON(10_000)

	b.ReportAllocs()
	for b.Loop() {
		var out bytes.Buffer
		if err := FilterActiveUsersInMemory(bytes.NewReader(input), &out); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFilterActiveUsersStreaming(b *testing.B) {
	input := generateUsersJSON(10_000)

	b.ReportAllocs()
	for b.Loop() {
		var out bytes.Buffer
		if err := FilterActiveUsersStreaming(bytes.NewReader(input), &out); err != nil {
			b.Fatal(err)
		}
	}
}

func generateUsersJSON(n int) []byte {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i := 0; i < n; i++ {
		if i > 0 {
			buf.WriteByte(',')
		}
		active := i%3 != 0
		fmt.Fprintf(&buf, `{"id":"u-%06d","active":%t,"name":"User %06d"}`, i, active, i)
	}
	buf.WriteByte(']')
	return buf.Bytes()
}
