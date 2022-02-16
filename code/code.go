package code

import (
	"encoding/binary"
	"fmt"
)

type Instructions []byte

// Defines the operator part of an instruction (e.g. for mnemonics PUSH, POP).
type Opcode byte

const (
	OpConstant Opcode = iota // Pushes one integer parameter onto the stack.
)

// Definition for an Opcode (just for debugging purposes).
// `Name` helps to make an Opcode more readable and `OperandWidths` contains the number
// of bytes each operand takes up.
type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
}

func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}

	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1
	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2: // Two byte operand
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}

	return instruction
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}
