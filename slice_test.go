package weak

import (
	"crypto/rand"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice_IsNil(t *testing.T) {
	t.Run("true:default", func(t *testing.T) {
		ws := Slice[int]{}
		assert.True(t, ws.IsNil())
		assert.Equal(t, 0, ws.Len())
		assert.Equal(t, 0, ws.Cap())
	})

	t.Run("true:new_nullptr", func(t *testing.T) {
		ws := NewSlice([]int(nil))
		assert.True(t, ws.IsNil())
		assert.Equal(t, 0, ws.Len())
		assert.Equal(t, 0, ws.Cap())
	})

	t.Run("false:new_ptr", func(t *testing.T) {
		ws := NewSlice([]int{1, 2, 3})
		assert.False(t, ws.IsNil())
		assert.Equal(t, 3, ws.Len())
		assert.Equal(t, 3, ws.Cap())
	})
}

func TestSlice_Indirect(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		s := make([]byte, 64)
		n, err := rand.Read(s)
		assert.Equal(t, 64, n)
		assert.Nil(t, err)

		w := NewSlice[byte](s)
		assert.Equal(t, s, w.Indirect())
	})
}
