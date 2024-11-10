package line

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/arikr/arikr/core/cell"
)

func TestNew(t *testing.T) {
	// success
	line := New(0xFF)
	assert.Equal(t, cell.Cell(0xFF), line.Cells[0])
	assert.Zero(t, line.At)
}

func TestGet(t *testing.T) {
	// setup
	line := New(0xFF)

	// success
	elem := line.Get()
	assert.Equal(t, cell.Cell(0xFF), elem)
	assert.Equal(t, 1, line.At)
}

func TestSeek(t *testing.T) {
	// setup
	line := New()

	// success
	line.Seek(0x01)
	assert.Equal(t, 1, line.At)
}

func TestSet(t *testing.T) {
	// setup
	line := New()

	// success
	line.Set(0xFF)
	assert.Equal(t, cell.Cell(0xFF), line.Cells[0])
}
