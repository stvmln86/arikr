package arikr

import "fmt"

// Core is a single emulated CPU core.
type Core struct {
	Array [8]uint8
	Index uint8
	Opers map[uint8]OperFunc
	Store []uint8
}

// NewCore returns a new Core containing a program and default operators.
func NewCore(elems []byte) *Core {
	return &Core{[8]uint8{}, 0, Opers, elems}
}

// Read returns the next byte from the Core's memory and increments the Index.
func (c *Core) Read() (uint8, error) {
	if c.Index < uint8(len(c.Store)) {
		c.Index++
		return c.Store[c.Index-1], nil
	}

	return 0, fmt.Errorf("core index %d out of bounds", c.Index)
}

// ReadN returns the next N bytes from the Core's memory and increments the Index.
func (c *Core) ReadN(size int) ([]uint8, error) {
	var elems []uint8
	for range size {
		elem, err := c.Read()
		if err != nil {
			return nil, err
		}
		elems = append(elems, elem)
	}

	return elems, nil
}
