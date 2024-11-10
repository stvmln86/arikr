package arikr

import "fmt"

// Core is a single emulated CPU core.
type Core struct {
	Array [8]byte
	Index uint8
	Opers map[byte]OperFunc
	Store []byte
}

// NewCore returns a new Core containing a program and default operators.
func NewCore(elems []byte) *Core {
	return &Core{[8]byte{}, 0, Opers, elems}
}

// Execute executes the next instruction in the Core.
func (c *Core) Execute() error {
	oper, err := c.Get()
	if err != nil {
		return err
	}

	ofun, ok := c.Opers[oper]
	if !ok {
		return fmt.Errorf("operator %X does not exist", oper)
	}

	return ofun(c)
}

// Get returns the next byte from the Core's memory and increments the Index.
func (c *Core) Get() (uint8, error) {
	if c.Index < uint8(len(c.Store)) {
		c.Index++
		return c.Store[c.Index-1], nil
	}

	return 0, fmt.Errorf("core index %d out of bounds", c.Index)
}

// GetN returns the next N bytes from the Core's memory and increments the Index.
func (c *Core) GetN(size int) ([]uint8, error) {
	var elems []uint8
	for range size {
		elem, err := c.Get()
		if err != nil {
			return nil, err
		}
		elems = append(elems, elem)
	}

	return elems, nil
}

// Run executes all instructions in the Core.
func (c *Core) Run() error {
	for int(c.Index) < len(c.Store) {
		if err := c.Execute(); err != nil {
			return err
		}
	}

	return nil
}
