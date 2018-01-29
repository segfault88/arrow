// Code generated by arraybuilder_numeric.gen.go.tmpl.
// DO NOT EDIT.

package arrow

import (
	"github.com/influxdata/arrow/internal/bitutil"
	"github.com/influxdata/arrow/memory"
)

type Int32ArrayBuilder struct {
	arrayBuilder

	data    *memory.PoolBuffer
	rawData []int32
}

func NewInt32ArrayBuilder(pool memory.Allocator) *Int32ArrayBuilder {
	return &Int32ArrayBuilder{arrayBuilder: arrayBuilder{pool: pool}}
}

func (b *Int32ArrayBuilder) Append(v int32) {
	b.Reserve(1)
	b.UnsafeAppend(v)
}

func (b *Int32ArrayBuilder) AppendNull() {
	b.Reserve(1)
	b.UnsafeAppendBoolToBitmap(false)
}

func (b *Int32ArrayBuilder) UnsafeAppend(v int32) {
	bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	b.rawData[b.length] = v
	b.length++
}

func (b *Int32ArrayBuilder) UnsafeAppendBoolToBitmap(isValid bool) {
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
func (b *Int32ArrayBuilder) AppendValues(v []int32, valid []bool) {
	if len(v) != len(valid) && len(valid) != 0 {
		panic("len(v) != len(valid) && len(valid) != 0")
	}

	b.Reserve(len(v))
	if len(v) > 0 {
		Int32Traits{}.Copy(b.rawData[b.length:], v)
	}
	b.arrayBuilder.unsafeAppendBoolsToBitmap(valid, len(v))
}

func (b *Int32ArrayBuilder) init(capacity int) {
	b.arrayBuilder.init(capacity)

	b.data = memory.NewPoolBuffer(b.pool)
	bytesN := Int32Traits{}.BytesRequired(capacity)
	b.data.Resize(bytesN)
	b.rawData = Int32Traits{}.CastFromBytes(b.data.Bytes())
}

// Reserve ensures there is enough space for appending n elements
// by checking the capacity and calling Resize if necessary.
func (b *Int32ArrayBuilder) Reserve(n int) {
	b.arrayBuilder.reserve(n, b.Resize)
}

// Resize adjusts the space allocated by b to n elements. If n is greater than b.Cap(),
// additional memory will be allocated. If n is smaller, the allocated memory may reduced.
func (b *Int32ArrayBuilder) Resize(n int) {
	if n < minBuilderCapacity {
		n = minBuilderCapacity
	}

	if b.capacity == 0 {
		b.init(n)
	} else {
		b.arrayBuilder.resize(n, b.init)
		b.data.Resize(Int32Traits{}.BytesRequired(n))
		b.rawData = Int32Traits{}.CastFromBytes(b.data.Bytes())
	}
}

func (b *Int32ArrayBuilder) Finish() *Int32Array {
	data := b.finishInternal()
	return NewInt32Array(data)
}

func (b *Int32ArrayBuilder) finishInternal() *ArrayData {
	bytesRequired := Int32Traits{}.BytesRequired(b.length)
	if bytesRequired > 0 && bytesRequired < b.data.Len() {
		// trim buffers
		b.data.Resize(bytesRequired)
	}
	res := NewArrayData(PrimitiveTypes.Int32, b.length, []*memory.Buffer{&b.nullBitmap.Buffer, &b.data.Buffer}, b.nullN)
	b.reset()

	return res
}

type Int64ArrayBuilder struct {
	arrayBuilder

	data    *memory.PoolBuffer
	rawData []int64
}

func NewInt64ArrayBuilder(pool memory.Allocator) *Int64ArrayBuilder {
	return &Int64ArrayBuilder{arrayBuilder: arrayBuilder{pool: pool}}
}

func (b *Int64ArrayBuilder) Append(v int64) {
	b.Reserve(1)
	b.UnsafeAppend(v)
}

func (b *Int64ArrayBuilder) AppendNull() {
	b.Reserve(1)
	b.UnsafeAppendBoolToBitmap(false)
}

func (b *Int64ArrayBuilder) UnsafeAppend(v int64) {
	bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	b.rawData[b.length] = v
	b.length++
}

func (b *Int64ArrayBuilder) UnsafeAppendBoolToBitmap(isValid bool) {
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
func (b *Int64ArrayBuilder) AppendValues(v []int64, valid []bool) {
	if len(v) != len(valid) && len(valid) != 0 {
		panic("len(v) != len(valid) && len(valid) != 0")
	}

	b.Reserve(len(v))
	if len(v) > 0 {
		Int64Traits{}.Copy(b.rawData[b.length:], v)
	}
	b.arrayBuilder.unsafeAppendBoolsToBitmap(valid, len(v))
}

func (b *Int64ArrayBuilder) init(capacity int) {
	b.arrayBuilder.init(capacity)

	b.data = memory.NewPoolBuffer(b.pool)
	bytesN := Int64Traits{}.BytesRequired(capacity)
	b.data.Resize(bytesN)
	b.rawData = Int64Traits{}.CastFromBytes(b.data.Bytes())
}

// Reserve ensures there is enough space for appending n elements
// by checking the capacity and calling Resize if necessary.
func (b *Int64ArrayBuilder) Reserve(n int) {
	b.arrayBuilder.reserve(n, b.Resize)
}

// Resize adjusts the space allocated by b to n elements. If n is greater than b.Cap(),
// additional memory will be allocated. If n is smaller, the allocated memory may reduced.
func (b *Int64ArrayBuilder) Resize(n int) {
	if n < minBuilderCapacity {
		n = minBuilderCapacity
	}

	if b.capacity == 0 {
		b.init(n)
	} else {
		b.arrayBuilder.resize(n, b.init)
		b.data.Resize(Int64Traits{}.BytesRequired(n))
		b.rawData = Int64Traits{}.CastFromBytes(b.data.Bytes())
	}
}

func (b *Int64ArrayBuilder) Finish() *Int64Array {
	data := b.finishInternal()
	return NewInt64Array(data)
}

func (b *Int64ArrayBuilder) finishInternal() *ArrayData {
	bytesRequired := Int64Traits{}.BytesRequired(b.length)
	if bytesRequired > 0 && bytesRequired < b.data.Len() {
		// trim buffers
		b.data.Resize(bytesRequired)
	}
	res := NewArrayData(PrimitiveTypes.Int64, b.length, []*memory.Buffer{&b.nullBitmap.Buffer, &b.data.Buffer}, b.nullN)
	b.reset()

	return res
}

type Uint64ArrayBuilder struct {
	arrayBuilder

	data    *memory.PoolBuffer
	rawData []uint64
}

func NewUint64ArrayBuilder(pool memory.Allocator) *Uint64ArrayBuilder {
	return &Uint64ArrayBuilder{arrayBuilder: arrayBuilder{pool: pool}}
}

func (b *Uint64ArrayBuilder) Append(v uint64) {
	b.Reserve(1)
	b.UnsafeAppend(v)
}

func (b *Uint64ArrayBuilder) AppendNull() {
	b.Reserve(1)
	b.UnsafeAppendBoolToBitmap(false)
}

func (b *Uint64ArrayBuilder) UnsafeAppend(v uint64) {
	bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	b.rawData[b.length] = v
	b.length++
}

func (b *Uint64ArrayBuilder) UnsafeAppendBoolToBitmap(isValid bool) {
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
func (b *Uint64ArrayBuilder) AppendValues(v []uint64, valid []bool) {
	if len(v) != len(valid) && len(valid) != 0 {
		panic("len(v) != len(valid) && len(valid) != 0")
	}

	b.Reserve(len(v))
	if len(v) > 0 {
		Uint64Traits{}.Copy(b.rawData[b.length:], v)
	}
	b.arrayBuilder.unsafeAppendBoolsToBitmap(valid, len(v))
}

func (b *Uint64ArrayBuilder) init(capacity int) {
	b.arrayBuilder.init(capacity)

	b.data = memory.NewPoolBuffer(b.pool)
	bytesN := Uint64Traits{}.BytesRequired(capacity)
	b.data.Resize(bytesN)
	b.rawData = Uint64Traits{}.CastFromBytes(b.data.Bytes())
}

// Reserve ensures there is enough space for appending n elements
// by checking the capacity and calling Resize if necessary.
func (b *Uint64ArrayBuilder) Reserve(n int) {
	b.arrayBuilder.reserve(n, b.Resize)
}

// Resize adjusts the space allocated by b to n elements. If n is greater than b.Cap(),
// additional memory will be allocated. If n is smaller, the allocated memory may reduced.
func (b *Uint64ArrayBuilder) Resize(n int) {
	if n < minBuilderCapacity {
		n = minBuilderCapacity
	}

	if b.capacity == 0 {
		b.init(n)
	} else {
		b.arrayBuilder.resize(n, b.init)
		b.data.Resize(Uint64Traits{}.BytesRequired(n))
		b.rawData = Uint64Traits{}.CastFromBytes(b.data.Bytes())
	}
}

func (b *Uint64ArrayBuilder) Finish() *Uint64Array {
	data := b.finishInternal()
	return NewUint64Array(data)
}

func (b *Uint64ArrayBuilder) finishInternal() *ArrayData {
	bytesRequired := Uint64Traits{}.BytesRequired(b.length)
	if bytesRequired > 0 && bytesRequired < b.data.Len() {
		// trim buffers
		b.data.Resize(bytesRequired)
	}
	res := NewArrayData(PrimitiveTypes.Uint64, b.length, []*memory.Buffer{&b.nullBitmap.Buffer, &b.data.Buffer}, b.nullN)
	b.reset()

	return res
}

type Float64ArrayBuilder struct {
	arrayBuilder

	data    *memory.PoolBuffer
	rawData []float64
}

func NewFloat64ArrayBuilder(pool memory.Allocator) *Float64ArrayBuilder {
	return &Float64ArrayBuilder{arrayBuilder: arrayBuilder{pool: pool}}
}

func (b *Float64ArrayBuilder) Append(v float64) {
	b.Reserve(1)
	b.UnsafeAppend(v)
}

func (b *Float64ArrayBuilder) AppendNull() {
	b.Reserve(1)
	b.UnsafeAppendBoolToBitmap(false)
}

func (b *Float64ArrayBuilder) UnsafeAppend(v float64) {
	bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	b.rawData[b.length] = v
	b.length++
}

func (b *Float64ArrayBuilder) UnsafeAppendBoolToBitmap(isValid bool) {
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
func (b *Float64ArrayBuilder) AppendValues(v []float64, valid []bool) {
	if len(v) != len(valid) && len(valid) != 0 {
		panic("len(v) != len(valid) && len(valid) != 0")
	}

	b.Reserve(len(v))
	if len(v) > 0 {
		Float64Traits{}.Copy(b.rawData[b.length:], v)
	}
	b.arrayBuilder.unsafeAppendBoolsToBitmap(valid, len(v))
}

func (b *Float64ArrayBuilder) init(capacity int) {
	b.arrayBuilder.init(capacity)

	b.data = memory.NewPoolBuffer(b.pool)
	bytesN := Float64Traits{}.BytesRequired(capacity)
	b.data.Resize(bytesN)
	b.rawData = Float64Traits{}.CastFromBytes(b.data.Bytes())
}

// Reserve ensures there is enough space for appending n elements
// by checking the capacity and calling Resize if necessary.
func (b *Float64ArrayBuilder) Reserve(n int) {
	b.arrayBuilder.reserve(n, b.Resize)
}

// Resize adjusts the space allocated by b to n elements. If n is greater than b.Cap(),
// additional memory will be allocated. If n is smaller, the allocated memory may reduced.
func (b *Float64ArrayBuilder) Resize(n int) {
	if n < minBuilderCapacity {
		n = minBuilderCapacity
	}

	if b.capacity == 0 {
		b.init(n)
	} else {
		b.arrayBuilder.resize(n, b.init)
		b.data.Resize(Float64Traits{}.BytesRequired(n))
		b.rawData = Float64Traits{}.CastFromBytes(b.data.Bytes())
	}
}

func (b *Float64ArrayBuilder) Finish() *Float64Array {
	data := b.finishInternal()
	return NewFloat64Array(data)
}

func (b *Float64ArrayBuilder) finishInternal() *ArrayData {
	bytesRequired := Float64Traits{}.BytesRequired(b.length)
	if bytesRequired > 0 && bytesRequired < b.data.Len() {
		// trim buffers
		b.data.Resize(bytesRequired)
	}
	res := NewArrayData(PrimitiveTypes.Float64, b.length, []*memory.Buffer{&b.nullBitmap.Buffer, &b.data.Buffer}, b.nullN)
	b.reset()

	return res
}

type TimestampArrayBuilder struct {
	arrayBuilder

	typE    *TimestampType
	data    *memory.PoolBuffer
	rawData []Timestamp
}

func NewTimestampArrayBuilder(pool memory.Allocator, typE *TimestampType) *TimestampArrayBuilder {
	return &TimestampArrayBuilder{arrayBuilder: arrayBuilder{pool: pool}, typE: typE}
}

func (b *TimestampArrayBuilder) Append(v Timestamp) {
	b.Reserve(1)
	b.UnsafeAppend(v)
}

func (b *TimestampArrayBuilder) AppendNull() {
	b.Reserve(1)
	b.UnsafeAppendBoolToBitmap(false)
}

func (b *TimestampArrayBuilder) UnsafeAppend(v Timestamp) {
	bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	b.rawData[b.length] = v
	b.length++
}

func (b *TimestampArrayBuilder) UnsafeAppendBoolToBitmap(isValid bool) {
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
func (b *TimestampArrayBuilder) AppendValues(v []Timestamp, valid []bool) {
	if len(v) != len(valid) && len(valid) != 0 {
		panic("len(v) != len(valid) && len(valid) != 0")
	}

	b.Reserve(len(v))
	if len(v) > 0 {
		TimestampTraits{}.Copy(b.rawData[b.length:], v)
	}
	b.arrayBuilder.unsafeAppendBoolsToBitmap(valid, len(v))
}

func (b *TimestampArrayBuilder) init(capacity int) {
	b.arrayBuilder.init(capacity)

	b.data = memory.NewPoolBuffer(b.pool)
	bytesN := TimestampTraits{}.BytesRequired(capacity)
	b.data.Resize(bytesN)
	b.rawData = TimestampTraits{}.CastFromBytes(b.data.Bytes())
}

// Reserve ensures there is enough space for appending n elements
// by checking the capacity and calling Resize if necessary.
func (b *TimestampArrayBuilder) Reserve(n int) {
	b.arrayBuilder.reserve(n, b.Resize)
}

// Resize adjusts the space allocated by b to n elements. If n is greater than b.Cap(),
// additional memory will be allocated. If n is smaller, the allocated memory may reduced.
func (b *TimestampArrayBuilder) Resize(n int) {
	if n < minBuilderCapacity {
		n = minBuilderCapacity
	}

	if b.capacity == 0 {
		b.init(n)
	} else {
		b.arrayBuilder.resize(n, b.init)
		b.data.Resize(TimestampTraits{}.BytesRequired(n))
		b.rawData = TimestampTraits{}.CastFromBytes(b.data.Bytes())
	}
}

func (b *TimestampArrayBuilder) Finish() *TimestampArray {
	data := b.finishInternal()
	return NewTimestampArray(data)
}

func (b *TimestampArrayBuilder) finishInternal() *ArrayData {
	bytesRequired := TimestampTraits{}.BytesRequired(b.length)
	if bytesRequired > 0 && bytesRequired < b.data.Len() {
		// trim buffers
		b.data.Resize(bytesRequired)
	}
	res := NewArrayData(b.typE, b.length, []*memory.Buffer{&b.nullBitmap.Buffer, &b.data.Buffer}, b.nullN)
	b.reset()

	return res
}
