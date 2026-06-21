package processor

import (
	"bytes"
	"testing"
)

func TestEncodeRecord(t *testing.T) {
	record := Record{ID: "user-123", Count: 42, Name: "Cassius"}
	want := []byte("user-123:42:Cassius")

	if got := EncodeRecordPlain(record); !bytes.Equal(got, want) {
		t.Fatalf("EncodeRecordPlain() = %q, want %q", got, want)
	}

	if got := EncodeRecordWithPool(record); !bytes.Equal(got, want) {
		t.Fatalf("EncodeRecordWithPool() = %q, want %q", got, want)
	}
}

func BenchmarkEncodeRecordPlain(b *testing.B) {
	record := Record{ID: "user-123", Count: 42, Name: "Cassius Paim"}

	b.ReportAllocs()
	for b.Loop() {
		_ = EncodeRecordPlain(record)
	}
}

func BenchmarkEncodeRecordWithPool(b *testing.B) {
	record := Record{ID: "user-123", Count: 42, Name: "Cassius Paim"}

	b.ReportAllocs()
	for b.Loop() {
		_ = EncodeRecordWithPool(record)
	}
}
