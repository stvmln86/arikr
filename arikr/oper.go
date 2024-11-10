package arikr

import (
	"fmt"
	"strconv"
)

// Oper is a single CPU opcode.
type Oper byte

// OperFunc is a function that performs operations.
type OperFunc func(*Core) error

// Opers is a map of opcodes to OperFuncs.
var Opers = map[uint8]OperFunc{
	// Nil operators.
	0x00: OperNOOP,

	// Program control operators.
	0x10: OperCOND,
	0x12: OperJUMP,
	0x13: OperLOAD,

	// Input/output operators.
	0x20: OperDUMP,
	0x21: OperECHO,

	// Calculation and logic operators.
	0x30: OperADDI,
	0x31: OperISEQ,
}

// OperNOOP (NOOP) does nothing.
func OperNOOP(c *Core) error {
	return nil
}

// OperCOND (COND R I) jumps to index I if R is not zero.
func OperCOND(c *Core) error {
	elems, err := c.GetN(2)
	if err != nil {
		return err
	}

	if c.Array[elems[0]] > 0 {
		c.Index = elems[1]
	}

	return nil
}

// OperJUMP (COND I) jumps to index I.
func OperJUMP(c *Core) error {
	elem, err := c.Get()
	if err != nil {
		return err
	}

	c.Index = elem
	return nil
}

// OperLOAD (LOAD R I) loads integer I into register R.
func OperLOAD(c *Core) error {
	elems, err := c.GetN(2)
	if err != nil {
		return err
	}

	c.Array[elems[0]] = elems[1]
	return nil
}

// OperDUMP (DUMP R) prints R to STDOUT as a numerical value.
func OperDUMP(c *Core) error {
	elem, err := c.Get()
	if err != nil {
		return err
	}

	fmt.Print(strconv.FormatInt(int64(elem), 10))
	return nil
}

// OperECHO (ECHO R) prints R to STDOUT as an ASCII character.
func OperECHO(c *Core) error {
	elem, err := c.Get()
	if err != nil {
		return err
	}

	fmt.Printf("%c", c.Array[elem])
	return nil
}

// OperADDI (ADDI R R) loads R + R into register 7.
func OperADDI(c *Core) error {
	elems, err := c.GetN(2)
	if err != nil {
		return err
	}

	c.Array[7] = c.Array[elems[0]] + c.Array[elems[1]]
	return nil
}

// OperISEQ (ISEQ R R) loads R == R into register 7.
func OperISEQ(c *Core) error {
	elems, err := c.GetN(2)
	if err != nil {
		return err
	}

	if c.Array[elems[0]] == c.Array[elems[1]] {
		c.Array[7] = 1
	} else {
		c.Array[7] = 0
	}

	return nil
}
