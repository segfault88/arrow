// Code generated by array/numericbuilder.gen.go.tmpl.
// DO NOT EDIT.

package array

import (
	"sync/atomic"

	"github.com/influxdata/arrow"
	"github.com/influxdata/arrow/internal/bitutil"
	"github.com/influxdata/arrow/internal/debug"
	"github.com/influxdata/arrow/memory"
)

type Int64Builder struct {
	builder

	data    *memory.Buffer
	rawData []int64
}

func NewInt64Builder(mem memory.Allocator) *Int64Builder {
	return &Int64Builder{builder: builder{refCount: 1, mem: mem}}
}

// Release decreases the reference count by 1.
// When the reference count goes to zero, the memory is freed.
func (b *Int64Builder) Release() {
	debug.Assert(atomic.LoadInt64(&b.refCount) > 0, "too many releases")

	if atomic.AddInt64(&b.refCount, -1) == 0 {
		if b.nullBitmap != nil {
			b.nullBitmap.Release()
			b.nullBitmap = nil
		}
		if b.data != nil {
			b.data.Release()
			b.data = nil
		}
	}
}

func (b *Int64Builder) Append(v int64) {
	b.Reserve(1)
	b.UnsafeAppend(v)
}

func (b *Int64Builder) AppendNull() {
	b.Reserve(1)
	b.UnsafeAppendBoolToBitmap(false)
}

func (b *Int64Builder) UnsafeAppend(v int64) {
	bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	b.rawData[b.length] = v
	b.length++
}

func (b *Int64Builder) UnsafeAppendBoolToBitmap(isValid bool) {
	if isValid {
		bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	} else {
		b.nullN++
	}
	b.length++
}

// AppendValues will append the values in the v slice. The valid slice determines which values
// in v are valid (not null). The valid slice must either be empty or be equal in length to v. If empty,
// all values in v are appended and considered valid.
func (b *Int64Builder) AppendValues(v []int64, valid []bool) {
	if len(v) != len(valid) && len(valid) != 0 {
		panic("len(v) != len(valid) && len(valid) != 0")
	}

	b.Reserve(len(v))
	if len(v) > 0 {
		arrow.Int64Traits.Copy(b.rawData[b.length:], v)
	}
	b.builder.unsafeAppendBoolsToBitmap(valid, len(v))
}

func (b *Int64Builder) init(capacity int) {
	b.builder.init(capacity)

	b.data = memory.NewResizableBuffer(b.mem)
	bytesN := arrow.Int64Traits.BytesRequired(capacity)
	b.data.Resize(bytesN)
	b.rawData = arrow.Int64Traits.CastFromBytes(b.data.Bytes())
}

// Reserve ensures there is enough space for appending n elements
// by checking the capacity and calling Resize if necessary.
func (b *Int64Builder) Reserve(n int) {
	b.builder.reserve(n, b.Resize)
}

// Resize adjusts the space allocated by b to n elements. If n is greater than b.Cap(),
// additional memory will be allocated. If n is smaller, the allocated memory may reduced.
func (b *Int64Builder) Resize(n int) {
	if n < minBuilderCapacity {
		n = minBuilderCapacity
	}

	if b.capacity == 0 {
		b.init(n)
	} else {
		b.builder.resize(n, b.init)
		b.data.Resize(arrow.Int64Traits.BytesRequired(n))
		b.rawData = arrow.Int64Traits.CastFromBytes(b.data.Bytes())
	}
}

// NewInt64Array creates a Int64 array from the memory buffers used by the builder and resets the Int64Builder
// so it can be used to build a new array.
func (b *Int64Builder) NewInt64Array() (a *Int64) {
	data := b.newData()
	a = NewInt64Data(data)
	data.Release()
	return
}

func (b *Int64Builder) newData() (data *Data) {
	bytesRequired := arrow.Int64Traits.BytesRequired(b.length)
	if bytesRequired > 0 && bytesRequired < b.data.Len() {
		// trim buffers
		b.data.Resize(bytesRequired)
	}
	data = NewData(arrow.PrimitiveTypes.Int64, b.length, []*memory.Buffer{b.nullBitmap, b.data}, b.nullN)
	b.reset()

	b.data.Release()
	b.data = nil
	b.rawData = nil

	return
}

type Uint64Builder struct {
	builder

	data    *memory.Buffer
	rawData []uint64
}

func NewUint64Builder(mem memory.Allocator) *Uint64Builder {
	return &Uint64Builder{builder: builder{refCount: 1, mem: mem}}
}

// Release decreases the reference count by 1.
// When the reference count goes to zero, the memory is freed.
func (b *Uint64Builder) Release() {
	debug.Assert(atomic.LoadInt64(&b.refCount) > 0, "too many releases")

	if atomic.AddInt64(&b.refCount, -1) == 0 {
		if b.nullBitmap != nil {
			b.nullBitmap.Release()
			b.nullBitmap = nil
		}
		if b.data != nil {
			b.data.Release()
			b.data = nil
		}
	}
}

func (b *Uint64Builder) Append(v uint64) {
	b.Reserve(1)
	b.UnsafeAppend(v)
}

func (b *Uint64Builder) AppendNull() {
	b.Reserve(1)
	b.UnsafeAppendBoolToBitmap(false)
}

func (b *Uint64Builder) UnsafeAppend(v uint64) {
	bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	b.rawData[b.length] = v
	b.length++
}

func (b *Uint64Builder) UnsafeAppendBoolToBitmap(isValid bool) {
	if isValid {
		bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	} else {
		b.nullN++
	}
	b.length++
}

// AppendValues will append the values in the v slice. The valid slice determines which values
// in v are valid (not null). The valid slice must either be empty or be equal in length to v. If empty,
// all values in v are appended and considered valid.
func (b *Uint64Builder) AppendValues(v []uint64, valid []bool) {
	if len(v) != len(valid) && len(valid) != 0 {
		panic("len(v) != len(valid) && len(valid) != 0")
	}

	b.Reserve(len(v))
	if len(v) > 0 {
		arrow.Uint64Traits.Copy(b.rawData[b.length:], v)
	}
	b.builder.unsafeAppendBoolsToBitmap(valid, len(v))
}

func (b *Uint64Builder) init(capacity int) {
	b.builder.init(capacity)

	b.data = memory.NewResizableBuffer(b.mem)
	bytesN := arrow.Uint64Traits.BytesRequired(capacity)
	b.data.Resize(bytesN)
	b.rawData = arrow.Uint64Traits.CastFromBytes(b.data.Bytes())
}

// Reserve ensures there is enough space for appending n elements
// by checking the capacity and calling Resize if necessary.
func (b *Uint64Builder) Reserve(n int) {
	b.builder.reserve(n, b.Resize)
}

// Resize adjusts the space allocated by b to n elements. If n is greater than b.Cap(),
// additional memory will be allocated. If n is smaller, the allocated memory may reduced.
func (b *Uint64Builder) Resize(n int) {
	if n < minBuilderCapacity {
		n = minBuilderCapacity
	}

	if b.capacity == 0 {
		b.init(n)
	} else {
		b.builder.resize(n, b.init)
		b.data.Resize(arrow.Uint64Traits.BytesRequired(n))
		b.rawData = arrow.Uint64Traits.CastFromBytes(b.data.Bytes())
	}
}

// NewUint64Array creates a Uint64 array from the memory buffers used by the builder and resets the Uint64Builder
// so it can be used to build a new array.
func (b *Uint64Builder) NewUint64Array() (a *Uint64) {
	data := b.newData()
	a = NewUint64Data(data)
	data.Release()
	return
}

func (b *Uint64Builder) newData() (data *Data) {
	bytesRequired := arrow.Uint64Traits.BytesRequired(b.length)
	if bytesRequired > 0 && bytesRequired < b.data.Len() {
		// trim buffers
		b.data.Resize(bytesRequired)
	}
	data = NewData(arrow.PrimitiveTypes.Uint64, b.length, []*memory.Buffer{b.nullBitmap, b.data}, b.nullN)
	b.reset()

	b.data.Release()
	b.data = nil
	b.rawData = nil

	return
}

type Float64Builder struct {
	builder

	data    *memory.Buffer
	rawData []float64
}

func NewFloat64Builder(mem memory.Allocator) *Float64Builder {
	return &Float64Builder{builder: builder{refCount: 1, mem: mem}}
}

// Release decreases the reference count by 1.
// When the reference count goes to zero, the memory is freed.
func (b *Float64Builder) Release() {
	debug.Assert(atomic.LoadInt64(&b.refCount) > 0, "too many releases")

	if atomic.AddInt64(&b.refCount, -1) == 0 {
		if b.nullBitmap != nil {
			b.nullBitmap.Release()
			b.nullBitmap = nil
		}
		if b.data != nil {
			b.data.Release()
			b.data = nil
		}
	}
}

func (b *Float64Builder) Append(v float64) {
	b.Reserve(1)
	b.UnsafeAppend(v)
}

func (b *Float64Builder) AppendNull() {
	b.Reserve(1)
	b.UnsafeAppendBoolToBitmap(false)
}

func (b *Float64Builder) UnsafeAppend(v float64) {
	bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	b.rawData[b.length] = v
	b.length++
}

func (b *Float64Builder) UnsafeAppendBoolToBitmap(isValid bool) {
	if isValid {
		bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	} else {
		b.nullN++
	}
	b.length++
}

// AppendValues will append the values in the v slice. The valid slice determines which values
// in v are valid (not null). The valid slice must either be empty or be equal in length to v. If empty,
// all values in v are appended and considered valid.
func (b *Float64Builder) AppendValues(v []float64, valid []bool) {
	if len(v) != len(valid) && len(valid) != 0 {
		panic("len(v) != len(valid) && len(valid) != 0")
	}

	b.Reserve(len(v))
	if len(v) > 0 {
		arrow.Float64Traits.Copy(b.rawData[b.length:], v)
	}
	b.builder.unsafeAppendBoolsToBitmap(valid, len(v))
}

func (b *Float64Builder) init(capacity int) {
	b.builder.init(capacity)

	b.data = memory.NewResizableBuffer(b.mem)
	bytesN := arrow.Float64Traits.BytesRequired(capacity)
	b.data.Resize(bytesN)
	b.rawData = arrow.Float64Traits.CastFromBytes(b.data.Bytes())
}

// Reserve ensures there is enough space for appending n elements
// by checking the capacity and calling Resize if necessary.
func (b *Float64Builder) Reserve(n int) {
	b.builder.reserve(n, b.Resize)
}

// Resize adjusts the space allocated by b to n elements. If n is greater than b.Cap(),
// additional memory will be allocated. If n is smaller, the allocated memory may reduced.
func (b *Float64Builder) Resize(n int) {
	if n < minBuilderCapacity {
		n = minBuilderCapacity
	}

	if b.capacity == 0 {
		b.init(n)
	} else {
		b.builder.resize(n, b.init)
		b.data.Resize(arrow.Float64Traits.BytesRequired(n))
		b.rawData = arrow.Float64Traits.CastFromBytes(b.data.Bytes())
	}
}

// NewFloat64Array creates a Float64 array from the memory buffers used by the builder and resets the Float64Builder
// so it can be used to build a new array.
func (b *Float64Builder) NewFloat64Array() (a *Float64) {
	data := b.newData()
	a = NewFloat64Data(data)
	data.Release()
	return
}

func (b *Float64Builder) newData() (data *Data) {
	bytesRequired := arrow.Float64Traits.BytesRequired(b.length)
	if bytesRequired > 0 && bytesRequired < b.data.Len() {
		// trim buffers
		b.data.Resize(bytesRequired)
	}
	data = NewData(arrow.PrimitiveTypes.Float64, b.length, []*memory.Buffer{b.nullBitmap, b.data}, b.nullN)
	b.reset()

	b.data.Release()
	b.data = nil
	b.rawData = nil

	return
}

type Int32Builder struct {
	builder

	data    *memory.Buffer
	rawData []int32
}

func NewInt32Builder(mem memory.Allocator) *Int32Builder {
	return &Int32Builder{builder: builder{refCount: 1, mem: mem}}
}

// Release decreases the reference count by 1.
// When the reference count goes to zero, the memory is freed.
func (b *Int32Builder) Release() {
	debug.Assert(atomic.LoadInt64(&b.refCount) > 0, "too many releases")

	if atomic.AddInt64(&b.refCount, -1) == 0 {
		if b.nullBitmap != nil {
			b.nullBitmap.Release()
			b.nullBitmap = nil
		}
		if b.data != nil {
			b.data.Release()
			b.data = nil
		}
	}
}

func (b *Int32Builder) Append(v int32) {
	b.Reserve(1)
	b.UnsafeAppend(v)
}

func (b *Int32Builder) AppendNull() {
	b.Reserve(1)
	b.UnsafeAppendBoolToBitmap(false)
}

func (b *Int32Builder) UnsafeAppend(v int32) {
	bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	b.rawData[b.length] = v
	b.length++
}

func (b *Int32Builder) UnsafeAppendBoolToBitmap(isValid bool) {
	if isValid {
		bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	} else {
		b.nullN++
	}
	b.length++
}

// AppendValues will append the values in the v slice. The valid slice determines which values
// in v are valid (not null). The valid slice must either be empty or be equal in length to v. If empty,
// all values in v are appended and considered valid.
func (b *Int32Builder) AppendValues(v []int32, valid []bool) {
	if len(v) != len(valid) && len(valid) != 0 {
		panic("len(v) != len(valid) && len(valid) != 0")
	}

	b.Reserve(len(v))
	if len(v) > 0 {
		arrow.Int32Traits.Copy(b.rawData[b.length:], v)
	}
	b.builder.unsafeAppendBoolsToBitmap(valid, len(v))
}

func (b *Int32Builder) init(capacity int) {
	b.builder.init(capacity)

	b.data = memory.NewResizableBuffer(b.mem)
	bytesN := arrow.Int32Traits.BytesRequired(capacity)
	b.data.Resize(bytesN)
	b.rawData = arrow.Int32Traits.CastFromBytes(b.data.Bytes())
}

// Reserve ensures there is enough space for appending n elements
// by checking the capacity and calling Resize if necessary.
func (b *Int32Builder) Reserve(n int) {
	b.builder.reserve(n, b.Resize)
}

// Resize adjusts the space allocated by b to n elements. If n is greater than b.Cap(),
// additional memory will be allocated. If n is smaller, the allocated memory may reduced.
func (b *Int32Builder) Resize(n int) {
	if n < minBuilderCapacity {
		n = minBuilderCapacity
	}

	if b.capacity == 0 {
		b.init(n)
	} else {
		b.builder.resize(n, b.init)
		b.data.Resize(arrow.Int32Traits.BytesRequired(n))
		b.rawData = arrow.Int32Traits.CastFromBytes(b.data.Bytes())
	}
}

// NewInt32Array creates a Int32 array from the memory buffers used by the builder and resets the Int32Builder
// so it can be used to build a new array.
func (b *Int32Builder) NewInt32Array() (a *Int32) {
	data := b.newData()
	a = NewInt32Data(data)
	data.Release()
	return
}

func (b *Int32Builder) newData() (data *Data) {
	bytesRequired := arrow.Int32Traits.BytesRequired(b.length)
	if bytesRequired > 0 && bytesRequired < b.data.Len() {
		// trim buffers
		b.data.Resize(bytesRequired)
	}
	data = NewData(arrow.PrimitiveTypes.Int32, b.length, []*memory.Buffer{b.nullBitmap, b.data}, b.nullN)
	b.reset()

	b.data.Release()
	b.data = nil
	b.rawData = nil

	return
}

type Uint32Builder struct {
	builder

	data    *memory.Buffer
	rawData []uint32
}

func NewUint32Builder(mem memory.Allocator) *Uint32Builder {
	return &Uint32Builder{builder: builder{refCount: 1, mem: mem}}
}

// Release decreases the reference count by 1.
// When the reference count goes to zero, the memory is freed.
func (b *Uint32Builder) Release() {
	debug.Assert(atomic.LoadInt64(&b.refCount) > 0, "too many releases")

	if atomic.AddInt64(&b.refCount, -1) == 0 {
		if b.nullBitmap != nil {
			b.nullBitmap.Release()
			b.nullBitmap = nil
		}
		if b.data != nil {
			b.data.Release()
			b.data = nil
		}
	}
}

func (b *Uint32Builder) Append(v uint32) {
	b.Reserve(1)
	b.UnsafeAppend(v)
}

func (b *Uint32Builder) AppendNull() {
	b.Reserve(1)
	b.UnsafeAppendBoolToBitmap(false)
}

func (b *Uint32Builder) UnsafeAppend(v uint32) {
	bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	b.rawData[b.length] = v
	b.length++
}

func (b *Uint32Builder) UnsafeAppendBoolToBitmap(isValid bool) {
	if isValid {
		bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	} else {
		b.nullN++
	}
	b.length++
}

// AppendValues will append the values in the v slice. The valid slice determines which values
// in v are valid (not null). The valid slice must either be empty or be equal in length to v. If empty,
// all values in v are appended and considered valid.
func (b *Uint32Builder) AppendValues(v []uint32, valid []bool) {
	if len(v) != len(valid) && len(valid) != 0 {
		panic("len(v) != len(valid) && len(valid) != 0")
	}

	b.Reserve(len(v))
	if len(v) > 0 {
		arrow.Uint32Traits.Copy(b.rawData[b.length:], v)
	}
	b.builder.unsafeAppendBoolsToBitmap(valid, len(v))
}

func (b *Uint32Builder) init(capacity int) {
	b.builder.init(capacity)

	b.data = memory.NewResizableBuffer(b.mem)
	bytesN := arrow.Uint32Traits.BytesRequired(capacity)
	b.data.Resize(bytesN)
	b.rawData = arrow.Uint32Traits.CastFromBytes(b.data.Bytes())
}

// Reserve ensures there is enough space for appending n elements
// by checking the capacity and calling Resize if necessary.
func (b *Uint32Builder) Reserve(n int) {
	b.builder.reserve(n, b.Resize)
}

// Resize adjusts the space allocated by b to n elements. If n is greater than b.Cap(),
// additional memory will be allocated. If n is smaller, the allocated memory may reduced.
func (b *Uint32Builder) Resize(n int) {
	if n < minBuilderCapacity {
		n = minBuilderCapacity
	}

	if b.capacity == 0 {
		b.init(n)
	} else {
		b.builder.resize(n, b.init)
		b.data.Resize(arrow.Uint32Traits.BytesRequired(n))
		b.rawData = arrow.Uint32Traits.CastFromBytes(b.data.Bytes())
	}
}

// NewUint32Array creates a Uint32 array from the memory buffers used by the builder and resets the Uint32Builder
// so it can be used to build a new array.
func (b *Uint32Builder) NewUint32Array() (a *Uint32) {
	data := b.newData()
	a = NewUint32Data(data)
	data.Release()
	return
}

func (b *Uint32Builder) newData() (data *Data) {
	bytesRequired := arrow.Uint32Traits.BytesRequired(b.length)
	if bytesRequired > 0 && bytesRequired < b.data.Len() {
		// trim buffers
		b.data.Resize(bytesRequired)
	}
	data = NewData(arrow.PrimitiveTypes.Uint32, b.length, []*memory.Buffer{b.nullBitmap, b.data}, b.nullN)
	b.reset()

	b.data.Release()
	b.data = nil
	b.rawData = nil

	return
}

type Float32Builder struct {
	builder

	data    *memory.Buffer
	rawData []float32
}

func NewFloat32Builder(mem memory.Allocator) *Float32Builder {
	return &Float32Builder{builder: builder{refCount: 1, mem: mem}}
}

// Release decreases the reference count by 1.
// When the reference count goes to zero, the memory is freed.
func (b *Float32Builder) Release() {
	debug.Assert(atomic.LoadInt64(&b.refCount) > 0, "too many releases")

	if atomic.AddInt64(&b.refCount, -1) == 0 {
		if b.nullBitmap != nil {
			b.nullBitmap.Release()
			b.nullBitmap = nil
		}
		if b.data != nil {
			b.data.Release()
			b.data = nil
		}
	}
}

func (b *Float32Builder) Append(v float32) {
	b.Reserve(1)
	b.UnsafeAppend(v)
}

func (b *Float32Builder) AppendNull() {
	b.Reserve(1)
	b.UnsafeAppendBoolToBitmap(false)
}

func (b *Float32Builder) UnsafeAppend(v float32) {
	bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	b.rawData[b.length] = v
	b.length++
}

func (b *Float32Builder) UnsafeAppendBoolToBitmap(isValid bool) {
	if isValid {
		bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	} else {
		b.nullN++
	}
	b.length++
}

// AppendValues will append the values in the v slice. The valid slice determines which values
// in v are valid (not null). The valid slice must either be empty or be equal in length to v. If empty,
// all values in v are appended and considered valid.
func (b *Float32Builder) AppendValues(v []float32, valid []bool) {
	if len(v) != len(valid) && len(valid) != 0 {
		panic("len(v) != len(valid) && len(valid) != 0")
	}

	b.Reserve(len(v))
	if len(v) > 0 {
		arrow.Float32Traits.Copy(b.rawData[b.length:], v)
	}
	b.builder.unsafeAppendBoolsToBitmap(valid, len(v))
}

func (b *Float32Builder) init(capacity int) {
	b.builder.init(capacity)

	b.data = memory.NewResizableBuffer(b.mem)
	bytesN := arrow.Float32Traits.BytesRequired(capacity)
	b.data.Resize(bytesN)
	b.rawData = arrow.Float32Traits.CastFromBytes(b.data.Bytes())
}

// Reserve ensures there is enough space for appending n elements
// by checking the capacity and calling Resize if necessary.
func (b *Float32Builder) Reserve(n int) {
	b.builder.reserve(n, b.Resize)
}

// Resize adjusts the space allocated by b to n elements. If n is greater than b.Cap(),
// additional memory will be allocated. If n is smaller, the allocated memory may reduced.
func (b *Float32Builder) Resize(n int) {
	if n < minBuilderCapacity {
		n = minBuilderCapacity
	}

	if b.capacity == 0 {
		b.init(n)
	} else {
		b.builder.resize(n, b.init)
		b.data.Resize(arrow.Float32Traits.BytesRequired(n))
		b.rawData = arrow.Float32Traits.CastFromBytes(b.data.Bytes())
	}
}

// NewFloat32Array creates a Float32 array from the memory buffers used by the builder and resets the Float32Builder
// so it can be used to build a new array.
func (b *Float32Builder) NewFloat32Array() (a *Float32) {
	data := b.newData()
	a = NewFloat32Data(data)
	data.Release()
	return
}

func (b *Float32Builder) newData() (data *Data) {
	bytesRequired := arrow.Float32Traits.BytesRequired(b.length)
	if bytesRequired > 0 && bytesRequired < b.data.Len() {
		// trim buffers
		b.data.Resize(bytesRequired)
	}
	data = NewData(arrow.PrimitiveTypes.Float32, b.length, []*memory.Buffer{b.nullBitmap, b.data}, b.nullN)
	b.reset()

	b.data.Release()
	b.data = nil
	b.rawData = nil

	return
}

type Int16Builder struct {
	builder

	data    *memory.Buffer
	rawData []int16
}

func NewInt16Builder(mem memory.Allocator) *Int16Builder {
	return &Int16Builder{builder: builder{refCount: 1, mem: mem}}
}

// Release decreases the reference count by 1.
// When the reference count goes to zero, the memory is freed.
func (b *Int16Builder) Release() {
	debug.Assert(atomic.LoadInt64(&b.refCount) > 0, "too many releases")

	if atomic.AddInt64(&b.refCount, -1) == 0 {
		if b.nullBitmap != nil {
			b.nullBitmap.Release()
			b.nullBitmap = nil
		}
		if b.data != nil {
			b.data.Release()
			b.data = nil
		}
	}
}

func (b *Int16Builder) Append(v int16) {
	b.Reserve(1)
	b.UnsafeAppend(v)
}

func (b *Int16Builder) AppendNull() {
	b.Reserve(1)
	b.UnsafeAppendBoolToBitmap(false)
}

func (b *Int16Builder) UnsafeAppend(v int16) {
	bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	b.rawData[b.length] = v
	b.length++
}

func (b *Int16Builder) UnsafeAppendBoolToBitmap(isValid bool) {
	if isValid {
		bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	} else {
		b.nullN++
	}
	b.length++
}

// AppendValues will append the values in the v slice. The valid slice determines which values
// in v are valid (not null). The valid slice must either be empty or be equal in length to v. If empty,
// all values in v are appended and considered valid.
func (b *Int16Builder) AppendValues(v []int16, valid []bool) {
	if len(v) != len(valid) && len(valid) != 0 {
		panic("len(v) != len(valid) && len(valid) != 0")
	}

	b.Reserve(len(v))
	if len(v) > 0 {
		arrow.Int16Traits.Copy(b.rawData[b.length:], v)
	}
	b.builder.unsafeAppendBoolsToBitmap(valid, len(v))
}

func (b *Int16Builder) init(capacity int) {
	b.builder.init(capacity)

	b.data = memory.NewResizableBuffer(b.mem)
	bytesN := arrow.Int16Traits.BytesRequired(capacity)
	b.data.Resize(bytesN)
	b.rawData = arrow.Int16Traits.CastFromBytes(b.data.Bytes())
}

// Reserve ensures there is enough space for appending n elements
// by checking the capacity and calling Resize if necessary.
func (b *Int16Builder) Reserve(n int) {
	b.builder.reserve(n, b.Resize)
}

// Resize adjusts the space allocated by b to n elements. If n is greater than b.Cap(),
// additional memory will be allocated. If n is smaller, the allocated memory may reduced.
func (b *Int16Builder) Resize(n int) {
	if n < minBuilderCapacity {
		n = minBuilderCapacity
	}

	if b.capacity == 0 {
		b.init(n)
	} else {
		b.builder.resize(n, b.init)
		b.data.Resize(arrow.Int16Traits.BytesRequired(n))
		b.rawData = arrow.Int16Traits.CastFromBytes(b.data.Bytes())
	}
}

// NewInt16Array creates a Int16 array from the memory buffers used by the builder and resets the Int16Builder
// so it can be used to build a new array.
func (b *Int16Builder) NewInt16Array() (a *Int16) {
	data := b.newData()
	a = NewInt16Data(data)
	data.Release()
	return
}

func (b *Int16Builder) newData() (data *Data) {
	bytesRequired := arrow.Int16Traits.BytesRequired(b.length)
	if bytesRequired > 0 && bytesRequired < b.data.Len() {
		// trim buffers
		b.data.Resize(bytesRequired)
	}
	data = NewData(arrow.PrimitiveTypes.Int16, b.length, []*memory.Buffer{b.nullBitmap, b.data}, b.nullN)
	b.reset()

	b.data.Release()
	b.data = nil
	b.rawData = nil

	return
}

type Uint16Builder struct {
	builder

	data    *memory.Buffer
	rawData []uint16
}

func NewUint16Builder(mem memory.Allocator) *Uint16Builder {
	return &Uint16Builder{builder: builder{refCount: 1, mem: mem}}
}

// Release decreases the reference count by 1.
// When the reference count goes to zero, the memory is freed.
func (b *Uint16Builder) Release() {
	debug.Assert(atomic.LoadInt64(&b.refCount) > 0, "too many releases")

	if atomic.AddInt64(&b.refCount, -1) == 0 {
		if b.nullBitmap != nil {
			b.nullBitmap.Release()
			b.nullBitmap = nil
		}
		if b.data != nil {
			b.data.Release()
			b.data = nil
		}
	}
}

func (b *Uint16Builder) Append(v uint16) {
	b.Reserve(1)
	b.UnsafeAppend(v)
}

func (b *Uint16Builder) AppendNull() {
	b.Reserve(1)
	b.UnsafeAppendBoolToBitmap(false)
}

func (b *Uint16Builder) UnsafeAppend(v uint16) {
	bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	b.rawData[b.length] = v
	b.length++
}

func (b *Uint16Builder) UnsafeAppendBoolToBitmap(isValid bool) {
	if isValid {
		bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	} else {
		b.nullN++
	}
	b.length++
}

// AppendValues will append the values in the v slice. The valid slice determines which values
// in v are valid (not null). The valid slice must either be empty or be equal in length to v. If empty,
// all values in v are appended and considered valid.
func (b *Uint16Builder) AppendValues(v []uint16, valid []bool) {
	if len(v) != len(valid) && len(valid) != 0 {
		panic("len(v) != len(valid) && len(valid) != 0")
	}

	b.Reserve(len(v))
	if len(v) > 0 {
		arrow.Uint16Traits.Copy(b.rawData[b.length:], v)
	}
	b.builder.unsafeAppendBoolsToBitmap(valid, len(v))
}

func (b *Uint16Builder) init(capacity int) {
	b.builder.init(capacity)

	b.data = memory.NewResizableBuffer(b.mem)
	bytesN := arrow.Uint16Traits.BytesRequired(capacity)
	b.data.Resize(bytesN)
	b.rawData = arrow.Uint16Traits.CastFromBytes(b.data.Bytes())
}

// Reserve ensures there is enough space for appending n elements
// by checking the capacity and calling Resize if necessary.
func (b *Uint16Builder) Reserve(n int) {
	b.builder.reserve(n, b.Resize)
}

// Resize adjusts the space allocated by b to n elements. If n is greater than b.Cap(),
// additional memory will be allocated. If n is smaller, the allocated memory may reduced.
func (b *Uint16Builder) Resize(n int) {
	if n < minBuilderCapacity {
		n = minBuilderCapacity
	}

	if b.capacity == 0 {
		b.init(n)
	} else {
		b.builder.resize(n, b.init)
		b.data.Resize(arrow.Uint16Traits.BytesRequired(n))
		b.rawData = arrow.Uint16Traits.CastFromBytes(b.data.Bytes())
	}
}

// NewUint16Array creates a Uint16 array from the memory buffers used by the builder and resets the Uint16Builder
// so it can be used to build a new array.
func (b *Uint16Builder) NewUint16Array() (a *Uint16) {
	data := b.newData()
	a = NewUint16Data(data)
	data.Release()
	return
}

func (b *Uint16Builder) newData() (data *Data) {
	bytesRequired := arrow.Uint16Traits.BytesRequired(b.length)
	if bytesRequired > 0 && bytesRequired < b.data.Len() {
		// trim buffers
		b.data.Resize(bytesRequired)
	}
	data = NewData(arrow.PrimitiveTypes.Uint16, b.length, []*memory.Buffer{b.nullBitmap, b.data}, b.nullN)
	b.reset()

	b.data.Release()
	b.data = nil
	b.rawData = nil

	return
}

type Int8Builder struct {
	builder

	data    *memory.Buffer
	rawData []int8
}

func NewInt8Builder(mem memory.Allocator) *Int8Builder {
	return &Int8Builder{builder: builder{refCount: 1, mem: mem}}
}

// Release decreases the reference count by 1.
// When the reference count goes to zero, the memory is freed.
func (b *Int8Builder) Release() {
	debug.Assert(atomic.LoadInt64(&b.refCount) > 0, "too many releases")

	if atomic.AddInt64(&b.refCount, -1) == 0 {
		if b.nullBitmap != nil {
			b.nullBitmap.Release()
			b.nullBitmap = nil
		}
		if b.data != nil {
			b.data.Release()
			b.data = nil
		}
	}
}

func (b *Int8Builder) Append(v int8) {
	b.Reserve(1)
	b.UnsafeAppend(v)
}

func (b *Int8Builder) AppendNull() {
	b.Reserve(1)
	b.UnsafeAppendBoolToBitmap(false)
}

func (b *Int8Builder) UnsafeAppend(v int8) {
	bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	b.rawData[b.length] = v
	b.length++
}

func (b *Int8Builder) UnsafeAppendBoolToBitmap(isValid bool) {
	if isValid {
		bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	} else {
		b.nullN++
	}
	b.length++
}

// AppendValues will append the values in the v slice. The valid slice determines which values
// in v are valid (not null). The valid slice must either be empty or be equal in length to v. If empty,
// all values in v are appended and considered valid.
func (b *Int8Builder) AppendValues(v []int8, valid []bool) {
	if len(v) != len(valid) && len(valid) != 0 {
		panic("len(v) != len(valid) && len(valid) != 0")
	}

	b.Reserve(len(v))
	if len(v) > 0 {
		arrow.Int8Traits.Copy(b.rawData[b.length:], v)
	}
	b.builder.unsafeAppendBoolsToBitmap(valid, len(v))
}

func (b *Int8Builder) init(capacity int) {
	b.builder.init(capacity)

	b.data = memory.NewResizableBuffer(b.mem)
	bytesN := arrow.Int8Traits.BytesRequired(capacity)
	b.data.Resize(bytesN)
	b.rawData = arrow.Int8Traits.CastFromBytes(b.data.Bytes())
}

// Reserve ensures there is enough space for appending n elements
// by checking the capacity and calling Resize if necessary.
func (b *Int8Builder) Reserve(n int) {
	b.builder.reserve(n, b.Resize)
}

// Resize adjusts the space allocated by b to n elements. If n is greater than b.Cap(),
// additional memory will be allocated. If n is smaller, the allocated memory may reduced.
func (b *Int8Builder) Resize(n int) {
	if n < minBuilderCapacity {
		n = minBuilderCapacity
	}

	if b.capacity == 0 {
		b.init(n)
	} else {
		b.builder.resize(n, b.init)
		b.data.Resize(arrow.Int8Traits.BytesRequired(n))
		b.rawData = arrow.Int8Traits.CastFromBytes(b.data.Bytes())
	}
}

// NewInt8Array creates a Int8 array from the memory buffers used by the builder and resets the Int8Builder
// so it can be used to build a new array.
func (b *Int8Builder) NewInt8Array() (a *Int8) {
	data := b.newData()
	a = NewInt8Data(data)
	data.Release()
	return
}

func (b *Int8Builder) newData() (data *Data) {
	bytesRequired := arrow.Int8Traits.BytesRequired(b.length)
	if bytesRequired > 0 && bytesRequired < b.data.Len() {
		// trim buffers
		b.data.Resize(bytesRequired)
	}
	data = NewData(arrow.PrimitiveTypes.Int8, b.length, []*memory.Buffer{b.nullBitmap, b.data}, b.nullN)
	b.reset()

	b.data.Release()
	b.data = nil
	b.rawData = nil

	return
}

type Uint8Builder struct {
	builder

	data    *memory.Buffer
	rawData []uint8
}

func NewUint8Builder(mem memory.Allocator) *Uint8Builder {
	return &Uint8Builder{builder: builder{refCount: 1, mem: mem}}
}

// Release decreases the reference count by 1.
// When the reference count goes to zero, the memory is freed.
func (b *Uint8Builder) Release() {
	debug.Assert(atomic.LoadInt64(&b.refCount) > 0, "too many releases")

	if atomic.AddInt64(&b.refCount, -1) == 0 {
		if b.nullBitmap != nil {
			b.nullBitmap.Release()
			b.nullBitmap = nil
		}
		if b.data != nil {
			b.data.Release()
			b.data = nil
		}
	}
}

func (b *Uint8Builder) Append(v uint8) {
	b.Reserve(1)
	b.UnsafeAppend(v)
}

func (b *Uint8Builder) AppendNull() {
	b.Reserve(1)
	b.UnsafeAppendBoolToBitmap(false)
}

func (b *Uint8Builder) UnsafeAppend(v uint8) {
	bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	b.rawData[b.length] = v
	b.length++
}

func (b *Uint8Builder) UnsafeAppendBoolToBitmap(isValid bool) {
	if isValid {
		bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	} else {
		b.nullN++
	}
	b.length++
}

// AppendValues will append the values in the v slice. The valid slice determines which values
// in v are valid (not null). The valid slice must either be empty or be equal in length to v. If empty,
// all values in v are appended and considered valid.
func (b *Uint8Builder) AppendValues(v []uint8, valid []bool) {
	if len(v) != len(valid) && len(valid) != 0 {
		panic("len(v) != len(valid) && len(valid) != 0")
	}

	b.Reserve(len(v))
	if len(v) > 0 {
		arrow.Uint8Traits.Copy(b.rawData[b.length:], v)
	}
	b.builder.unsafeAppendBoolsToBitmap(valid, len(v))
}

func (b *Uint8Builder) init(capacity int) {
	b.builder.init(capacity)

	b.data = memory.NewResizableBuffer(b.mem)
	bytesN := arrow.Uint8Traits.BytesRequired(capacity)
	b.data.Resize(bytesN)
	b.rawData = arrow.Uint8Traits.CastFromBytes(b.data.Bytes())
}

// Reserve ensures there is enough space for appending n elements
// by checking the capacity and calling Resize if necessary.
func (b *Uint8Builder) Reserve(n int) {
	b.builder.reserve(n, b.Resize)
}

// Resize adjusts the space allocated by b to n elements. If n is greater than b.Cap(),
// additional memory will be allocated. If n is smaller, the allocated memory may reduced.
func (b *Uint8Builder) Resize(n int) {
	if n < minBuilderCapacity {
		n = minBuilderCapacity
	}

	if b.capacity == 0 {
		b.init(n)
	} else {
		b.builder.resize(n, b.init)
		b.data.Resize(arrow.Uint8Traits.BytesRequired(n))
		b.rawData = arrow.Uint8Traits.CastFromBytes(b.data.Bytes())
	}
}

// NewUint8Array creates a Uint8 array from the memory buffers used by the builder and resets the Uint8Builder
// so it can be used to build a new array.
func (b *Uint8Builder) NewUint8Array() (a *Uint8) {
	data := b.newData()
	a = NewUint8Data(data)
	data.Release()
	return
}

func (b *Uint8Builder) newData() (data *Data) {
	bytesRequired := arrow.Uint8Traits.BytesRequired(b.length)
	if bytesRequired > 0 && bytesRequired < b.data.Len() {
		// trim buffers
		b.data.Resize(bytesRequired)
	}
	data = NewData(arrow.PrimitiveTypes.Uint8, b.length, []*memory.Buffer{b.nullBitmap, b.data}, b.nullN)
	b.reset()

	b.data.Release()
	b.data = nil
	b.rawData = nil

	return
}

type TimestampBuilder struct {
	builder

	typE    *arrow.TimestampType
	data    *memory.Buffer
	rawData []arrow.Timestamp
}

func NewTimestampBuilder(mem memory.Allocator, typE *arrow.TimestampType) *TimestampBuilder {
	return &TimestampBuilder{builder: builder{refCount: 1, mem: mem}, typE: typE}
}

// Release decreases the reference count by 1.
// When the reference count goes to zero, the memory is freed.
func (b *TimestampBuilder) Release() {
	debug.Assert(atomic.LoadInt64(&b.refCount) > 0, "too many releases")

	if atomic.AddInt64(&b.refCount, -1) == 0 {
		if b.nullBitmap != nil {
			b.nullBitmap.Release()
			b.nullBitmap = nil
		}
		if b.data != nil {
			b.data.Release()
			b.data = nil
		}
	}
}

func (b *TimestampBuilder) Append(v arrow.Timestamp) {
	b.Reserve(1)
	b.UnsafeAppend(v)
}

func (b *TimestampBuilder) AppendNull() {
	b.Reserve(1)
	b.UnsafeAppendBoolToBitmap(false)
}

func (b *TimestampBuilder) UnsafeAppend(v arrow.Timestamp) {
	bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	b.rawData[b.length] = v
	b.length++
}

func (b *TimestampBuilder) UnsafeAppendBoolToBitmap(isValid bool) {
	if isValid {
		bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	} else {
		b.nullN++
	}
	b.length++
}

// AppendValues will append the values in the v slice. The valid slice determines which values
// in v are valid (not null). The valid slice must either be empty or be equal in length to v. If empty,
// all values in v are appended and considered valid.
func (b *TimestampBuilder) AppendValues(v []arrow.Timestamp, valid []bool) {
	if len(v) != len(valid) && len(valid) != 0 {
		panic("len(v) != len(valid) && len(valid) != 0")
	}

	b.Reserve(len(v))
	if len(v) > 0 {
		arrow.TimestampTraits.Copy(b.rawData[b.length:], v)
	}
	b.builder.unsafeAppendBoolsToBitmap(valid, len(v))
}

func (b *TimestampBuilder) init(capacity int) {
	b.builder.init(capacity)

	b.data = memory.NewResizableBuffer(b.mem)
	bytesN := arrow.TimestampTraits.BytesRequired(capacity)
	b.data.Resize(bytesN)
	b.rawData = arrow.TimestampTraits.CastFromBytes(b.data.Bytes())
}

// Reserve ensures there is enough space for appending n elements
// by checking the capacity and calling Resize if necessary.
func (b *TimestampBuilder) Reserve(n int) {
	b.builder.reserve(n, b.Resize)
}

// Resize adjusts the space allocated by b to n elements. If n is greater than b.Cap(),
// additional memory will be allocated. If n is smaller, the allocated memory may reduced.
func (b *TimestampBuilder) Resize(n int) {
	if n < minBuilderCapacity {
		n = minBuilderCapacity
	}

	if b.capacity == 0 {
		b.init(n)
	} else {
		b.builder.resize(n, b.init)
		b.data.Resize(arrow.TimestampTraits.BytesRequired(n))
		b.rawData = arrow.TimestampTraits.CastFromBytes(b.data.Bytes())
	}
}

// NewTimestampArray creates a Timestamp array from the memory buffers used by the builder and resets the TimestampBuilder
// so it can be used to build a new array.
func (b *TimestampBuilder) NewTimestampArray() (a *Timestamp) {
	data := b.newData()
	a = NewTimestampData(data)
	data.Release()
	return
}

func (b *TimestampBuilder) newData() (data *Data) {
	bytesRequired := arrow.TimestampTraits.BytesRequired(b.length)
	if bytesRequired > 0 && bytesRequired < b.data.Len() {
		// trim buffers
		b.data.Resize(bytesRequired)
	}
	data = NewData(b.typE, b.length, []*memory.Buffer{b.nullBitmap, b.data}, b.nullN)
	b.reset()

	b.data.Release()
	b.data = nil
	b.rawData = nil

	return
}
