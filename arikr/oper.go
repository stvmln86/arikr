package arikr

import (
	"fmt"
	"strconv"
)

// Oper is a single CPU opcode.
type Oper byte

// OperFunc is a function that performs operations.
type OperFunc func(*Core, []byte) error

// Constants for defined Opers.
const (
	// Nil operators.
	NOOP Oper = 0x00

	// Program control operators.
	COND Oper = 0x10
	JUMP Oper = 0x12
	LOAD Oper = 0x13

	// Input/output operators.
	DUMP Oper = 0x20
	ECHO Oper = 0x21

	// Calculation and logic operators.
	ADDI Oper = 0x30
	ISEQ Oper = 0x31
)

// OperNOOP (NOOP) does nothing.
func OperNOOP(c *Core, elems []byte) error {
	return nil
}

// OperCOND (COND R I) jumps to index I if R is not zero.
func OperCOND(c *Core, elems []byte) error {
	if c.Array[elems[0]] > 0 {
		c.Index = elems[1]
	}

	return nil
}

// OperJUMP (COND I) jumps to index I.
func OperJUMP(c *Core, elems []byte) error {
	c.Index = elems[0]
	return nil
}

// OperLOAD (LOAD R I) loads integer I into register R.
func OperLOAD(c *Core, elems []byte) error {
	c.Array[elems[0]] = elems[1]
	return nil
}

// OperDUMP (DUMP R) prints R to STDOUT as a numerical value.
func OperDUMP(c *Core, elems []byte) error {
	elem := c.Array[elems[0]]
	fmt.Print(strconv.FormatInt(int64(elem), 10))
	return nil
}

// OperECHO (ECHO R) prints R to STDOUT as an ASCII character.
func OperECHO(c *Core, elems []byte) error {
	fmt.Printf("%c", c.Array[elems[0]])
	return nil
}

// OperADDI (ADDI R R) loads R + R into register 7.
func OperADDI(c *Core, elems []byte) error {
	c.Array[7] = c.Array[elems[0]] + c.Array[elems[1]]
	return nil
}

// OperISEQ (ISEQ R R) loads R == R into register 7.
func OperISEQ(c *Core, elems []byte) error {
	if c.Array[elems[0]] == c.Array[elems[1]] {
		c.Array[7] = 1
	} else {
		c.Array[7] = 0
	}

	return nil
}
