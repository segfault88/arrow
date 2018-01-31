// Code generated by type_amd64.go.tmpl.
// DO NOT EDIT.

// +build !noasm

package math

import (
	"unsafe"

	"github.com/influxdata/arrow/array"
)

//go:noescape
func _sum_int64_sse4(buf, len, res unsafe.Pointer)

func sum_int64_sse4(a *array.Int64) int64 {
	buf := a.Int64Values()
	var (
		p1  = unsafe.Pointer(&buf[0])
		p2  = unsafe.Pointer(uintptr(len(buf)))
		res int64
	)
	_sum_int64_sse4(p1, p2, unsafe.Pointer(&res))
	return res
}
