package processor

import (
	"fmt"
	"strconv"
	"strings"
)

// User is intentionally small so the examples focus on allocation behavior,
// not on a complex domain model.
type User struct {
	Name  string
	Score int
}

// FormatUsersSlow uses general formatting and repeated string concatenation.
// It is intentionally allocation-heavy for benchmark comparison.
func FormatUsersSlow(users []User) string {
	out := ""
	for _, user := range users {
		out += fmt.Sprintf("%s:%d\n", user.Name, user.Score)
	}
	return out
}

// FormatUsersBuilder uses strings.Builder and specialized integer conversion.
// Grow reserves capacity in the builder's internal byte storage. The multiplier
// is an approximate capacity hint, not a correctness requirement.
func FormatUsersBuilder(users []User) string {
	var b strings.Builder
	b.Grow(len(users) * 16)

	for _, user := range users {
		b.WriteString(user.Name)
		b.WriteByte(':')
		b.WriteString(strconv.Itoa(user.Score))
		b.WriteByte('\n')
	}
	return b.String()
}

// NamesSlow lets append grow the output slice as needed.
func NamesSlow(users []User) []string {
	var names []string
	for _, user := range users {
		names = append(names, user.Name)
	}
	return names
}

// NamesPreallocated expresses the expected output size up front.
func NamesPreallocated(users []User) []string {
	names := make([]string, 0, len(users))
	for _, user := range users {
		names = append(names, user.Name)
	}
	return names
}

// PrefixView returns a small slice that may retain a large backing array.
func PrefixView(data []byte, n int) []byte {
	return data[:n]
}

// PrefixCopy copies the prefix into a new backing array with the intended lifetime.
func PrefixCopy(data []byte, n int) []byte {
	out := make([]byte, n)
	copy(out, data[:n])
	return out
}
