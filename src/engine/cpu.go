package engine

import (
	"engine/components"
	"time"
	"fmt"
//	"log"
)

type INTERRUPTSType struct {
	master	byte
	enable	byte
	flags	byte
}

var INTERRUPTS INTERRUPTSType = INTERRUPTSType {
	master:		0x00,
	enable:		0x00,
	flags:		0x00,
}

/*
	CPU Structure
	================
	---> Instructions Array
	---> Registers Structure
*/
type CPUType struct {
	INSTRUCTIONS	[]components.InstructionType
	REGISTERS		*components.RegistersType
	DEBUG			bool
}

var CPU = CPUType {
	INSTRUCTIONS:	components.INSTRUCTIONS,
	REGISTERS:		&components.REGISTERS,
	DEBUG:			false,
}

func (cpu *CPUType) Run(debug bool) {
	cpu.DEBUG = debug
	for ;; {
		if cpu.INSTRUCTIONS[ROM.data[cpu.REGISTERS.PC]].Name == "UNKNOWN" {
			break
		}
		fmt.Println("CPU Step")
		if cpu.DEBUG {
			cpu.REGISTERS.Print()
			time.Sleep(time.Second)
		} else {
			time.Sleep(4 * time.Microsecond)
		}
		cpu.INSTRUCTIONS[ROM.data[cpu.REGISTERS.PC]].Exec()
		cpu.REGISTERS.PC++
		
	}
//	finished <- true
}

func (cpu *CPUType) Reset() {
	
}