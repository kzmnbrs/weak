package weak

import "unsafe"

type Pointer[T any] uintptr

func NewPointer[T any](x *T) Pointer[T] {
	return Pointer[T](unsafe.Pointer(x))
}

func (p *Pointer[T]) Set(x *T) {
	*p = Pointer[T](unsafe.Pointer(x))
}

func (p *Pointer[T]) Cast() *T {
	//nolint
	//goland:noinspection GoVetUnsafePointer
	return (*T)(unsafe.Pointer(*p))
}

func (p *Pointer[T]) Indirect() T {
	return *p.Cast()
}

func (p *Pointer[T]) IsNil() bool {
	return *p == 0
}
