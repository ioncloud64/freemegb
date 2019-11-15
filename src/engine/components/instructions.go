package components

import (
	"fmt"
)

type InstructionType struct {
	Exec		func()
	Opcode		uint8
	Name		string
	NumOperands	byte
	Operands	[]byte
	Cycles		uint8
	
}

var INSTRUCTIONS = []InstructionType {
	// 0x00 - NOP
	{
		func() {fmt.Println("NOP")},				// executed code
		0x00,										// opcode
		"NOP",										// name
		0,											// number of operands
		[]byte {},									// operands
		4,											// cpu cycles
	},
	// 0x01 - NOP
	{
		func() {fmt.Println("LOAD BC NN")},			// executed code
		0x00,										// opcode
		"LOAD BC NN",								// name
		0,											// number of operands
		[]byte {},									// operands
		4,											// cpu cycles
	},
}