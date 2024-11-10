package cell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// success
	cell := New(0xFF)
	assert.Equal(t, Cell(0xFF), cell)
}

func TestRune(t *testing.T) {
	// success
	elem := New(0x41).Rune()
	assert.Equal(t, "A", elem)
}

func TestString(t *testing.T) {
	// success
	elem := New(0xFF).String()
	assert.Equal(t, "FF", elem)
}
