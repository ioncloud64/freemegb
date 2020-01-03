package components

import (
	"fmt"
)

type InstructionType struct {
	Exec        func() // executed code
	Opcode      uint8  // opcode
	Name        string // name
	NumOperands byte   // number of operands
	Operands    []byte // operands
	Cycles      uint8  // cpu cycles

}

var INSTRUCTIONS = []InstructionType{
	// 0x00 - NOP
	{
		Exec:        func() {},
		Opcode:      0x00,
		Name:        "NOP",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x01 - LOAD BC NN
	{
		Exec: func() {
			// When transferring the code from JS to Go
			// prefix ++ does not work on type uint16

			REGISTERS.PC++
			REGISTERS.BC = REGISTERS.CombineTo16(ROMref[REGISTERS.PC], ROMref[REGISTERS.PC+1])
			REGISTERS.PC++
		},
		Opcode:      0x01,         // opcode
		Name:        "LOAD BC NN", // name
		NumOperands: 2,            // number of operands
		Operands:    []byte{},     // operands
		Cycles:      12,           // cpu cycles
	},
	// 0x02 - LOAD BC NN
	{
		Exec:        func() { fmt.Println("LOAD BC NN") }, // executed code
		Opcode:      0x02,                                 // opcode
		Name:        "LOAD BC NN",                         // name
		NumOperands: 0,                                    // number of operands
		Operands:    []byte{},                             // operands
		Cycles:      4,                                    // cpu cycles
	},
	// 0x03 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x03,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x04 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x04,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x05 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x05,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x06 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x06,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x07 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x07,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x08 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x08,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x09 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x09,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x0A - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x0A,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x0B - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x0B,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x0C - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x0C,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x0D - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x0D,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x0E - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x0E,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x0F - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x0F,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x10 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x10,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x11 - UNKNOWN
	{
		Exec: func() {
			REGISTERS.PC++
			REGISTERS.DE = REGISTERS.CombineTo16(ROMref[REGISTERS.PC], ROMref[REGISTERS.PC+1])
			REGISTERS.PC++
		},
		Opcode:      0x11,
		Name:        "LOAD DE NN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x12 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x12,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x13 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x13,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x14 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x14,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x15 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x15,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x16 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x16,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x17 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x17,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x18 - UNKNOWN
	{
		// Set PC to PC + Operand
		Exec:        func() { REGISTERS.PC = REGISTERS.PC + uint16(ROMref[REGISTERS.PC+1]-1) },
		Opcode:      0x18,
		Name:        "JUMP PC+N",
		NumOperands: 1,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x19 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x19,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x1A - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x1A,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x1B - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x1B,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x1C - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x1C,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x1D - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x1D,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x1E - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x1E,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x1F - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x1F,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x20 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x20,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x21 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x21,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x22 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x22,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x23 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x23,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x24 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x24,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x25 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x25,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x26 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x26,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x27 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x27,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x28 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x28,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x29 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x29,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x2A - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x2A,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x2B - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x2B,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x2C - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x2C,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x2D - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x2D,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x2E - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x2E,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x2F - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x2F,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x30 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x30,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x31 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x31,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x32 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x32,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x33 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x33,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x34 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x34,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x35 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x35,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x36 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x36,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x37 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x37,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x38 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x38,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x39 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x39,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x3A - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x3A,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x3B - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x3B,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x3C - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x3C,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x3D - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x3D,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x3E - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x3E,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x3F - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x3F,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x40 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x40,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x41 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x41,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x42 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x42,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x43 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x43,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x44 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x44,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x45 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x45,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x46 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x46,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x47 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x47,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x48 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x48,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x49 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x49,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x4A - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x4A,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x4B - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x4B,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x4C - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x4C,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x4D - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x4D,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x4E - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x4E,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x4F - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x4F,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x50 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x50,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x51 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x51,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x52 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x52,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x53 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x53,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x54 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x54,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x55 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x55,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x56 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x56,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x57 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x57,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x58 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x58,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x59 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x59,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x5A - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x5A,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x5B - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x5B,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x5C - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x5C,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x5D - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x5D,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x5E - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x5E,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x5F - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x5F,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x60 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x60,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x61 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x61,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x62 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x62,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x63 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x63,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x64 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x64,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x65 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x65,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x66 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x66,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x67 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x67,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x68 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x68,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x69 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x69,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x6A - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x6A,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x6B - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x6B,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x6C - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x6C,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x6D - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x6D,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x6E - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x6E,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x6F - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x6F,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x70 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x70,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x71 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x71,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x72 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x72,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x73 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x73,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x74 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x74,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x75 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x75,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x76 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x76,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x77 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x77,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x78 - LOAD A B
	{
		Exec: func() {
			REGISTERS.SetA(REGISTERS.B())
		},
		Opcode:      0x78,
		Name:        "LOAD A B",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x79 - LOAD A C
	{
		Exec: func() {
			REGISTERS.SetA(REGISTERS.C())
		},
		Opcode:      0x79,
		Name:        "LOAD A C",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x7A - LOAD A D
	{
		Exec: func() {
			REGISTERS.SetA(REGISTERS.D())
		},
		Opcode:      0x7A,
		Name:        "LOAD A D",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x7B - LOAD A E
	{
		Exec: func() {
			REGISTERS.SetA(REGISTERS.E())
		},
		Opcode:      0x7B,
		Name:        "LOAD A E",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x7C - LOAD A H
	{
		Exec: func() {
			REGISTERS.SetA(REGISTERS.H())
		},
		Opcode:      0x7C,
		Name:        "LOAD A H",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x7D - LOAD A L
	{
		Exec: func() {
			REGISTERS.SetA(REGISTERS.L())
		},
		Opcode:      0x7D,
		Name:        "LOAD A L",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x7E - LOAD A HL
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x7E,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x7F - LOAD A A
	{
		Exec:        func() {},
		Opcode:      0x7F,
		Name:        "LOAD A A",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x80 - ADD A B
	{
		Exec: func() {
			REGISTERS.SetA(REGISTERS.Add8(REGISTERS.A(), REGISTERS.B()))
		},
		Opcode:      0x80,
		Name:        "ADD A B",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x81 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x81,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x82 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x82,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x83 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x83,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x84 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x84,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x85 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x85,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x86 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x86,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x87 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x87,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x88 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x88,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x89 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x89,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x8A - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x8A,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x8B - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x8B,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x8C - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x8C,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x8D - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x8D,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x8E - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x8E,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x8F - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x8F,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x90 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x90,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x91 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x91,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x92 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x92,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x93 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x93,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x94 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x94,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x95 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x95,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x96 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x96,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x97 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x97,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x98 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x98,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x99 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x99,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x9A - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x9A,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x9B - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x9B,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x9C - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x9C,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x9D - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x9D,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x9E - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x9E,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0x9F - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x9F,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xA0 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xA0,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xA1 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xA1,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xA2 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xA2,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xA3 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xA3,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xA4 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xA4,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xA5 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xA5,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xA6 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xA6,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xA7 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xA7,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xA8 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xA8,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xA9 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xA9,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xAA - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xAA,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xAB - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xAB,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xAC - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xAC,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xAD - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xAD,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xAE - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xAE,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xAF - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xAF,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xB0 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xB0,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xB1 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xB1,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xB2 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xB2,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xB3 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xB3,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xB4 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xB4,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xB5 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xB5,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xB6 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xB6,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xB7 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xB7,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xB8 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xB8,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xB9 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xB9,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xBA - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xBA,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xBB - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xBB,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xBC - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xBC,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xBD - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xBD,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xBE - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xBE,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xBF - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xBF,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xC0 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xC0,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xC1 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xC1,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xC2 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xC2,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xC3 - JUMP
	{
		Exec: func() {
			REGISTERS.PC = REGISTERS.CombineTo16(ROMref[REGISTERS.PC+2], ROMref[REGISTERS.PC+1])
			REGISTERS.PC--
		},
		Opcode:      0xC3,
		Name:        "JUMP",
		NumOperands: 2,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xC4 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xC4,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xC5 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xC5,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xC6 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xC6,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xC7 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xC7,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xC8 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xC8,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xC9 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xC9,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xCA - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xCA,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xCB - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xCB,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xCC - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xCC,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xCD - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xCD,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xCE - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xCE,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xCF - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xCF,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xD0 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xD0,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xD1 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xD1,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xD2 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xD2,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xD3 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xD3,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xD4 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xD4,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xD5 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xD5,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xD6 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xD6,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xD7 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xD7,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xD8 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xD8,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xD9 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xD9,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xDA - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xDA,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xDB - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xDB,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xDC - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xDC,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xDD - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xDD,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xDE - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xDE,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xDF - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xDF,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xE0 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xE0,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xE1 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xE1,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xE2 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xE2,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xE3 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xE3,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xE4 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xE4,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xE5 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xE5,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xE6 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xE6,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xE7 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xE7,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xE8 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xE8,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xE9 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xE9,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xEA - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xEA,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xEB - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xEB,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xEC - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xEC,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xED - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xED,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xEE - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xEE,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xEF - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xEF,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xF0 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xF0,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xF1 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xF1,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xF2 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xF2,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xF3 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xF3,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xF4 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xF4,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xF5 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xF5,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xF6 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xF6,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xF7 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xF7,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xF8 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xF8,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xF9 - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xF9,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xFA - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xFA,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xFB - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xFB,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xFC - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xFC,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xFD - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xFD,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xFE - UNKNOWN
	{
		Exec: func() {
			operand := ROMref[REGISTERS.PC+1]
			if REGISTERS.A() == operand {
				// MMU.flags.SET(MMU.flags.ZERO, REGISTERS)
			} else if REGISTERS.A() < operand {
				// MMU.flags.SET(MMU.flags.CARRY, REGISTERS)
			}
		},
		Opcode:      0xFE,
		Name:        "CP A N",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
	// 0xFF - UNKNOWN
	{
		Exec:        func() { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xFF,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operands:    []byte{},
		Cycles:      4,
	},
}
