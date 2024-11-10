// Package line implements the Line type and methods.
package line

import "github.com/stvmln86/arikr/arikr/core/cell"

// Line is a single sequence of stored Cells.
type Line struct {
	At    int
	Cells [256]cell.Cell
}

// New returns a new Line from a Cell sequence.
func New(elems ...cell.Cell) *Line {
	line := new(Line)
	for n, elem := range elems {
		line.Cells[n] = elem
	}

	return line
}

// Get increments the Line's pointer and returns the current Cell.
func (l *Line) Get() cell.Cell {
	elem := l.Cells[l.At]
	l.At++
	if l.At > 255 {
		l.At = 255
	}

	return elem
}

// Seek sets the Line's pointer.
func (l *Line) Seek(size cell.Cell) {
	l.At = int(size)
	if l.At > 255 {
		l.At = 255
	}
}

// Set sets the Line's current Cell.
func (l *Line) Set(elem cell.Cell) {
	l.Cells[l.At] = elem
}
