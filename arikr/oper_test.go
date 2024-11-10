package arikr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperNOOP(t *testing.T) {
	// setup
	core := NewCore(nil)

	// success
	err := OperNOOP(core, nil)
	assert.NoError(t, err)
}

func TestOperCOND(t *testing.T) {
	// setup
	core := NewCore(nil)
	core.Array[0] = 1

	// success - true
	err := OperCOND(core, []byte{0x00, 0x01})
	assert.Equal(t, uint8(0x01), core.Index)
	assert.NoError(t, err)

	// success - false
	err = OperCOND(core, []byte{0x01, 0x02})
	assert.Equal(t, uint8(0x01), core.Index)
	assert.NoError(t, err)
}

func TestOperJUMP(t *testing.T) {
	// setup
	core := NewCore(nil)

	// success
	err := OperJUMP(core, []byte{0x01})
	assert.Equal(t, uint8(0x01), core.Index)
	assert.NoError(t, err)
}

func TestOperLOAD(t *testing.T) {
	// setup
	core := NewCore(nil)

	// success
	err := OperLOAD(core, []byte{0x00, 0x01})
	assert.Equal(t, uint8(0x01), core.Array[0])
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
	core := NewCore(nil)
	core.Array[0] = 0x01
	core.Array[1] = 0x02

	// success
	err := OperADDI(core, []byte{0x00, 0x01})
	assert.Equal(t, uint8(0x03), core.Array[7])
	assert.NoError(t, err)
}

func TestOperISEQ(t *testing.T) {
	// setup
	core := NewCore(nil)
	core.Array[0] = 0x01
	core.Array[1] = 0x01

	// success - true
	err := OperISEQ(core, []byte{0x00, 0x01})
	assert.Equal(t, uint8(0x01), core.Array[7])
	assert.NoError(t, err)

	// setup
	core.Array[1] = 0x02

	// success - false
	err = OperISEQ(core, []byte{0x00, 0x01})
	assert.Equal(t, uint8(0x00), core.Array[7])
	assert.NoError(t, err)
}
