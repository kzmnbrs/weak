package weak

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointer_Set(t *testing.T) {
	x, y := 10, 20

	wp := NewPointer(&x)
	assert.Equal(t, &x, wp.Cast())

	wp.Set(&y)
	assert.Equal(t, &y, wp.Cast())
}

func TestPointer_Cast(t *testing.T) {
	x := 10
	wp := NewPointer(&x)

	assert.Equal(t, &x, wp.Cast())
}

func TestPointer_Indirect(t *testing.T) {
	x := 10
	wp := NewPointer(&x)

	assert.Equal(t, x, wp.Indirect())
}

func TestPointer_IsNil(t *testing.T) {
	t.Run("true:default", func(t *testing.T) {
		var wp Pointer[int]
		assert.True(t, wp.IsNil())
	})

	t.Run("true:new_nullptr", func(t *testing.T) {
		wp := NewPointer[int](nil)
		assert.True(t, wp.IsNil())
	})

	t.Run("false:new_ptr", func(t *testing.T) {
		wp := NewPointer(new(int))
		assert.False(t, wp.IsNil())
	})
}

//go:noinline
func stackWeakPtr(i int) Pointer[int] {
	return NewPointer(&i)
}

func TestPointer_FaultyIndirect_StackReuse(t *testing.T) {
	wp := stackWeakPtr(10)
	assert.Equal(t, 10, wp.Indirect())

	_ = stackWeakPtr(11)
	assert.Equal(t, 11, wp.Indirect())
}
