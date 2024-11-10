package arikr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCore(t *testing.T) {
	// success
	core := NewCore([]byte{0x01, 0x02, 0x03})
	assert.Equal(t, [8]uint8{0, 0, 0, 0, 0, 0, 0, 0}, core.Array)
	assert.Equal(t, uint8(0x00), core.Index)
	assert.Equal(t, []byte{0x01, 0x02, 0x03}, core.Store)
}

func TestExec(t *testing.T) {
	// setup
	core := NewCore([]byte{0x13, 0x00, 0xFF})

	// success
	err := core.Execute()
	assert.Equal(t, [8]uint8{0xFF, 0, 0, 0, 0, 0, 0, 0}, core.Array)
	assert.Equal(t, uint8(0x03), core.Index)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	core := NewCore([]byte{0x01})

	// success
	elem, err := core.Get()
	assert.Equal(t, byte(0x01), elem)
	assert.Equal(t, uint8(0x01), core.Index)
	assert.NoError(t, err)

	// failure - out of bounds
	elem, err = core.Get()
	assert.Zero(t, elem)
	assert.EqualError(t, err, "core index 1 out of bounds")
}

func TestGetN(t *testing.T) {
	// setup
	core := NewCore([]byte{0x01, 0x02})

	// success
	elems, err := core.GetN(2)
	assert.Equal(t, []byte{0x01, 0x02}, elems)
	assert.Equal(t, uint8(0x02), core.Index)
	assert.NoError(t, err)

	// failure - out of bounds
	elems, err = core.GetN(1)
	assert.Empty(t, elems)
	assert.EqualError(t, err, "core index 2 out of bounds")
}

func TestRun(t *testing.T) {
	// setup
	core := NewCore([]byte{0x13, 0x00, 0xFF})

	// success
	err := core.Run()
	assert.Equal(t, [8]uint8{0xFF, 0, 0, 0, 0, 0, 0, 0}, core.Array)
	assert.Equal(t, uint8(0x03), core.Index)
	assert.NoError(t, err)
}
