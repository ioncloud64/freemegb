package core

import (
	"time"
	"runtime"
	"fmt"
	"github.com/gotk3/gotk3/glib"
	"github.com/ioncloud64/freemegb/core/components"
)

/*
	Interrupts
*/
type INTERRUPTSType struct {
	master byte
	enable byte
	flags  byte
}

var INTERRUPTS INTERRUPTSType = INTERRUPTSType{
	master: 0x00,
	enable: 0x00,
	flags:  0x00,
}

/*
	CPU Structure
	================
	---> Instructions Array
	---> Registers Structure
*/
type CPUType struct {
	INSTRUCTIONS []components.InstructionType
	REGISTERS    *components.RegistersType
	DEBUG        bool
}

var CPU = CPUType{
	INSTRUCTIONS: components.INSTRUCTIONS,
	REGISTERS:    &components.REGISTERS,
	DEBUG:        false,
}

func (cpu *CPUType) Run(debug bool) {
	cpu.DEBUG = debug
	for {
		if cpu.INSTRUCTIONS[ROM.data[cpu.REGISTERS.PC]].Name == "UNKNOWN" {
			var PCString = cpu.REGISTERS.Register16toString(cpu.REGISTERS.PC)
			components.Logger.Printf("UNKNOWN INSTRUCTION:\n\tINSTRUCTION: 0x%02X\n\tAt ROM Offset: %s\n",
				cpu.INSTRUCTIONS[ROM.data[cpu.REGISTERS.PC]].Opcode, PCString)
			// TODO: Add OS Switch for glib notifications in linux
			if runtime.GOOS == "linux" {
				notif := glib.NotificationNew("FreeMe!GB: Unknown Instruction")
				notif.SetBody(fmt.Sprintf("UNKNOWN INSTRUCTION:\nINSTRUCTION: 0x%02X\nAt ROM Offset: %s",
					cpu.INSTRUCTIONS[ROM.data[cpu.REGISTERS.PC]].Opcode, PCString))
				notif.SetIcon("ui/freemegb.png")
				components.AppRef.SendNotification("com.ioncloud64.freemegb", notif)
			}
			break
		}
		components.Logger.Println("CPU Step")
		if cpu.DEBUG {
			cpu.REGISTERS.Print()
			// time.Sleep(time.Second)
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
