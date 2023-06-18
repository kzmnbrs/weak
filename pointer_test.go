package weak

import (
	"crypto/rand"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestPointer_IsNull(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		var ptr Pointer[int]
		assert.True(t, ptr.IsNil())
	})

	t.Run("true:new_pointer", func(t *testing.T) {
		ptr := NewPointer[int](nil)
		assert.True(t, ptr.IsNil())
	})

	t.Run("false", func(t *testing.T) {
		ptr := NewPointer(new(int))
		assert.False(t, ptr.IsNil())

		ptr.Set(nil)
		assert.True(t, ptr.IsNil())
	})
}

func TestPointer_Cast(t *testing.T) {
	x := &struct {
		y int
		z string
	}{}
	xptr := NewPointer(x)
	assert.Equal(t, x, xptr.Cast())
}

func TestPointer_Indirect(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ss := make([]byte, 64)
		rand.Read(ss)
		ssptr := NewPointer[[]byte](&ss)

		runtime.GC()
		runtime.GC()

		ss1 := ssptr.Indirect()
		assert.NotPanics(t, func() {
			ss1[0] = '1'
		})
		assert.Len(t, ss1, 64)
	})

	t.Run("segfault", func(t *testing.T) {
		ss := make([]byte, 64)
		n, err := rand.Read(ss)
		assert.Equal(t, 64, n)
		assert.Nil(t, err)

		ssptr := NewPointer[[]byte](&ss)

		ss = nil
		runtime.GC()
		runtime.GC()

		assert.Panics(t, func() {
			ss1 := ssptr.Indirect()
			ss1[0] = '1'
		})
	})
}
