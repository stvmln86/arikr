package arikr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperNOOP(t *testing.T) {
	// setup
	core := NewCore(nil)

	// success
	err := OperNOOP(core)
	assert.NoError(t, err)
}

func TestOperCOND(t *testing.T) {
	// setup
	core := NewCore([]byte{0x00, 0xFF})
	core.Array[0] = 1

	// success - true
	err := OperCOND(core)
	assert.Equal(t, uint8(0xFF), core.Index)
	assert.NoError(t, err)

	// setup
	core = NewCore([]byte{0x00, 0xFF})

	// success - false
	err = OperCOND(core)
	assert.Equal(t, uint8(0x02), core.Index)
	assert.NoError(t, err)
}

func TestOperJUMP(t *testing.T) {
	// setup
	core := NewCore([]byte{0xFF})

	// success
	err := OperJUMP(core)
	assert.Equal(t, uint8(0xFF), core.Index)
	assert.NoError(t, err)
}

func TestOperLOAD(t *testing.T) {
	// setup
	core := NewCore([]byte{0x00, 0xFF})

	// success
	err := OperLOAD(core)
	assert.Equal(t, uint8(0xFF), core.Array[0])
	assert.NoError(t, err)
}

func TestOperDUMP(t *testing.T) {
	// can't test I/O yet
}

func TestOperECHO(t *testing.T) {
	// can't test I/O yet
}

func TestOperADDI(t *testing.T) {
	// setup
	core := NewCore([]byte{0x00, 0x01})
	core.Array[0] = 0x10
	core.Array[1] = 0x20

	// success
	err := OperADDI(core)
	assert.Equal(t, uint8(0x30), core.Array[7])
	assert.NoError(t, err)
}

func TestOperISEQ(t *testing.T) {
	// setup
	core := NewCore([]byte{0x00, 0x01})
	core.Array[0] = 0x10
	core.Array[1] = 0x10

	// success - true
	err := OperISEQ(core)
	assert.Equal(t, uint8(0x01), core.Array[7])
	assert.NoError(t, err)

	// setup
	core.Array[1] = 0x20
	core.Index = 0

	// success - false
	err = OperISEQ(core)
	assert.Equal(t, uint8(0x00), core.Array[7])
	assert.NoError(t, err)
}
