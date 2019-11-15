package engine

import (
	"engine/components"
	"time"
	"fmt"
)

/*
	CPU Structure
	================
	---> Instructions Array
	---> Registers Structure
*/
type CPUType struct {
	INSTRUCTIONS	[]components.InstructionType
	REGISTERS		components.RegistersType
}

var CPU = CPUType {
	INSTRUCTIONS:	components.INSTRUCTIONS,
	REGISTERS:		components.REGISTERS,
}

func (cpu *CPUType) Run(finished chan bool) {
	for i := 0; i < 20; i++ {
		fmt.Println("CPU Step")
		time.Sleep(1500 * time.Millisecond)
		cpu.REGISTERS.Print()
	}
	finished <- true
}