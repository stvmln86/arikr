// Package cell implements the Cell type and functions.
package cell

import (
	"fmt"
)

// Cell is a single byte in memory.
type Cell byte

// New returns a new Cell from an integer.
func New(cell byte) Cell {
	return Cell(cell)
}

// Rune returns the Cell as an ASCII character rune.
func (c Cell) Rune() string {
	return fmt.Sprintf("%c", byte(c))
}

// String returns the Cell as a hexidecimal string.
func (c Cell) String() string {
	return fmt.Sprintf("%X", byte(c))
}
