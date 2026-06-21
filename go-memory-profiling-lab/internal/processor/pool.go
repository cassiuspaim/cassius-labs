package processor

import (
	"bytes"
	"strconv"
	"sync"
)

const maxPooledBufferCapacity = 64 * 1024

// Record is used by the sync.Pool examples.
type Record struct {
	ID    string
	Count int
	Name  string
}

var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

// EncodeRecordPlain creates a fresh temporary buffer for every call.
func EncodeRecordPlain(record Record) []byte {
	var buf bytes.Buffer

	buf.WriteString(record.ID)
	buf.WriteByte(':')
	buf.WriteString(strconv.Itoa(record.Count))
	buf.WriteByte(':')
	buf.WriteString(record.Name)

	out := make([]byte, buf.Len())
	copy(out, buf.Bytes())
	return out
}

// EncodeRecordWithPool reuses a temporary buffer through sync.Pool.
// The returned byte slice is copied so the caller does not observe a pooled buffer
// that may be reused by another operation later.
func EncodeRecordWithPool(record Record) []byte {
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()

	defer func() {
		if buf.Cap() <= maxPooledBufferCapacity {
			bufferPool.Put(buf)
		}
	}()

	buf.WriteString(record.ID)
	buf.WriteByte(':')
	buf.WriteString(strconv.Itoa(record.Count))
	buf.WriteByte(':')
	buf.WriteString(record.Name)

	out := make([]byte, buf.Len())
	copy(out, buf.Bytes())
	return out
}
