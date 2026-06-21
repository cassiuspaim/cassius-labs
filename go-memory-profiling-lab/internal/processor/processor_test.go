package processor

import (
	"fmt"
	"testing"
)

func sampleUsers(n int) []User {
	users := make([]User, 0, n)
	for i := 0; i < n; i++ {
		users = append(users, User{
			Name:  fmt.Sprintf("user-%04d", i),
			Score: i % 100,
		})
	}
	return users
}

func BenchmarkFormatUsersSlow(b *testing.B) {
	users := sampleUsers(100)

	b.ReportAllocs()
	for b.Loop() {
		_ = FormatUsersSlow(users)
	}
}

func BenchmarkFormatUsersBuilder(b *testing.B) {
	users := sampleUsers(100)

	b.ReportAllocs()
	for b.Loop() {
		_ = FormatUsersBuilder(users)
	}
}

func BenchmarkNamesSlow(b *testing.B) {
	users := sampleUsers(1000)

	b.ReportAllocs()
	for b.Loop() {
		_ = NamesSlow(users)
	}
}

func BenchmarkNamesPreallocated(b *testing.B) {
	users := sampleUsers(1000)

	b.ReportAllocs()
	for b.Loop() {
		_ = NamesPreallocated(users)
	}
}

func BenchmarkPrefixView(b *testing.B) {
	data := make([]byte, 8<<20)

	b.ReportAllocs()
	for b.Loop() {
		_ = PrefixView(data, 16)
	}
}

func BenchmarkPrefixCopy(b *testing.B) {
	data := make([]byte, 8<<20)

	b.ReportAllocs()
	for b.Loop() {
		_ = PrefixCopy(data, 16)
	}
}
