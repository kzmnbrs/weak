package weak

import (
	"crypto/rand"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice_IsNil(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		s := Slice[int]{}
		assert.True(t, s.IsNil())
		assert.Equal(t, 0, s.Len())
		assert.Equal(t, 0, s.Cap())
	})

	t.Run("true", func(t *testing.T) {
		s := NewSlice([]int(nil))
		assert.True(t, s.IsNil())
		assert.Equal(t, 0, s.Len())
		assert.Equal(t, 0, s.Cap())
	})

	t.Run("false", func(t *testing.T) {
		s0 := make([]int, 0, 3)
		s0 = append(s0, 1, 2, 3)
		s := NewSlice(s0)
		assert.False(t, s.IsNil())
		assert.Equal(t, 3, s.Len())
		assert.Equal(t, 3, s.Cap())
	})
}

func TestSlice_Indirect(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ss := make([]byte, 64)
		n, err := rand.Read(ss)
		assert.Equal(t, 64, n)
		assert.Nil(t, err)

		w := NewSlice[byte](ss)
		assert.Equal(t, ss, w.Indirect())
	})
}
