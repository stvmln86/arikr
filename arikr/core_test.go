package arikr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCore(t *testing.T) {
	// success
	core := NewCore([]byte{0x01, 0x02, 0x03})
	assert.Equal(t, [8]uint8{0, 0, 0, 0, 0, 0, 0, 0}, core.Array)
	assert.Equal(t, uint8(0), core.Index)
	assert.Equal(t, []byte{0x01, 0x02, 0x03}, core.Store)
}

func TestRead(t *testing.T) {
	// setup
	core := NewCore([]byte{0x01})

	// success
	elem, err := core.Read()
	assert.Equal(t, byte(0x01), elem)
	assert.Equal(t, uint8(1), core.Index)
	assert.NoError(t, err)

	// failure - out of bounds
	elem, err = core.Read()
	assert.Zero(t, elem)
	assert.EqualError(t, err, "core index 1 out of bounds")
}

func TestReadN(t *testing.T) {
	// setup
	core := NewCore([]byte{0x01, 0x02})

	// success
	elems, err := core.ReadN(2)
	assert.Equal(t, []byte{0x01, 0x02}, elems)
	assert.Equal(t, uint8(2), core.Index)
	assert.NoError(t, err)

	// failure - out of bounds
	elems, err = core.ReadN(1)
	assert.Empty(t, elems)
	assert.EqualError(t, err, "core index 2 out of bounds")
}
