package array

import (
	"github.com/influxdata/arrow/internal/bitutil"
	"github.com/influxdata/arrow/memory"
)

// A bufferBuilder provides common functionality for populating memory with a sequence of type-specific values.
// Specialized implementations provide type-safe APIs for appending and accessing the memory.
type bufferBuilder struct {
	mem      memory.Allocator
	buffer   *memory.ResizableBuffer
	length   int
	capacity int

	bytes []byte
}

// Len returns the length of the memory buffer in bytes.
func (b *bufferBuilder) Len() int { return b.length }

// Cap returns the total number of bytes that can be stored without allocating additional memory.
func (b *bufferBuilder) Cap() int { return b.capacity }

// Bytes returns a slice of length b.Len().
// The slice is only valid for use until the next buffer modification. That is, until the next call
// to Advance, Reset, Finish or any Append function. The slice aliases the buffer content at least until the next
// buffer modification.
func (b *bufferBuilder) Bytes() []byte { return b.bytes[:b.length] }

func (b *bufferBuilder) resize(elements int) {
	if b.buffer == nil {
		b.buffer = memory.NewResizableBuffer(b.mem)
	}

	b.buffer.Resize(elements)
	oldCapacity := b.capacity
	b.capacity = b.buffer.Cap()
	b.bytes = b.buffer.Buf()

	if b.capacity > oldCapacity {
		memory.Set(b.bytes[oldCapacity:], 0)
	}
}

// Advance increases the buffer by length and initializes the skipped bytes to zero.
func (b *bufferBuilder) Advance(length int) {
	if b.capacity < b.length+length {
		newCapacity := bitutil.NextPowerOf2(b.length + length)
		b.resize(newCapacity)
	}
	b.length += length
}

// Append appends the contents of v to the buffer, resizing it if necessary.
func (b *bufferBuilder) Append(v []byte) {
	if b.capacity < b.length+len(v) {
		newCapacity := bitutil.NextPowerOf2(b.length + len(v))
		b.resize(newCapacity)
	}
	b.unsafeAppend(v)
}

// Reset returns the buffer to an empty state. Reset releases the memory and sets the length and capacity to zero.
func (b *bufferBuilder) Reset() {
	b.buffer, b.bytes = nil, nil
	b.capacity, b.length = 0, 0
}

// Finish TODO(sgc)
func (b *bufferBuilder) Finish() *memory.Buffer {
	if b.length > 0 {
		b.buffer.ResizeNoShrink(b.length)
	}
	res := &b.buffer.Buffer
	b.Reset()
	return res
}

func (b *bufferBuilder) unsafeAppend(data []byte) {
	copy(b.bytes[b.length:], data)
	b.length += len(data)
}
