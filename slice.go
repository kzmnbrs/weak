package weak

import (
	"reflect"
	"unsafe"
)

type Slice[T any] struct {
	data     uintptr
	len, cap int
}

func NewSlice[T any](x []T) Slice[T] {
	sh := *(*reflect.SliceHeader)(unsafe.Pointer(&x))
	return Slice[T]{
		data: sh.Data,
		len:  sh.Len,
		cap:  sh.Cap,
	}
}

func (s Slice[T]) Indirect() []T {
	return *(*[]T)(unsafe.Pointer(&reflect.SliceHeader{
		Data: s.data,
		Len:  s.len,
		Cap:  s.cap,
	}))
}

func (s Slice[T]) Len() int {
	return s.len
}

func (s Slice[T]) Cap() int {
	return s.cap
}

func (s Slice[T]) IsNil() bool {
	return s.data == 0
}
