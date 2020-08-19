package core

import (
	"fmt"
	"runtime"
	"time"

	// "log"
	// "github.com/gotk3/gotk3/glib"
	"github.com/esiqveland/notify"
	"github.com/godbus/dbus"
	// "gopkg.in/toast.v1"
)

// INTERRUPTSType is the structure to define constant values used to identify an interrupt
type INTERRUPTSType struct {
	master byte
	enable byte
	flags  byte
}

// INTERRUPTS is the exported object used in the system
//
// INTERRUPTS is exported for value setting in other files
var INTERRUPTS INTERRUPTSType = INTERRUPTSType{
	master: 0x00,
	enable: 0x00,
	flags:  0x00,
}

// CPUType is the structure to define what's inside a CPU
//  CPU Structure
//  ================
//  ---> Instructions Array
//  ---> Registers Structure
//  ---> DEBUG boolean value set with CPU.Run()
//  ================
type CPUType struct {
	INSTRUCTIONS []InstructionType
	REGISTERS    *RegistersType
	DEBUG        bool
}

// CPU is the exported object used in the system
// CPU is exported to become a shared variable in the System object
var CPU = CPUType{
	INSTRUCTIONS: INSTRUCTIONS,
	REGISTERS:    &REGISTERS,
	DEBUG:        false,
}

// Run is the thread loop function for the CPU
func (cpu *CPUType) Run(debug bool) {
	// TODO: Proper CPU control flow with stepping

	// TODO: Proper Breakpoint insertion using an array of addresses
	cpu.DEBUG = debug
	for {
		if cpu.INSTRUCTIONS[ROM.data[cpu.REGISTERS.PC]].Name == "UNKNOWN" {
			var PCString = cpu.REGISTERS.Register16toString(cpu.REGISTERS.PC)
			Logger.Logf(LogTypes.ERROR, "UNKNOWN INSTRUCTION:\n\t\t\t\tINSTRUCTION: 0x%02X\n\t\t\t\tAt ROM Offset: %s\n",
				cpu.INSTRUCTIONS[ROM.data[cpu.REGISTERS.PC]].Opcode, PCString)
			if runtime.GOOS == "linux" {
				conn, err := dbus.SessionBus()
				if err != nil {
					panic(err)
				}
				iconName := "ioncloud64-freemegb"
				notif := notify.Notification{
					AppName:    "FreeMe!GB",
					ReplacesID: uint32(0),
					AppIcon:    iconName,
					Summary:    "Unknown Instruction",
					Body: fmt.Sprintf("INSTRUCTION: 0x%02X\nAt ROM Offset: %s",
						cpu.INSTRUCTIONS[ROM.data[cpu.REGISTERS.PC]].Opcode, PCString),
					Actions:       []string{"cancel", "Cancel", "open", "Open"}, // tuples of (action_key, label)
					Hints:         map[string]dbus.Variant{"desktop-entry": dbus.MakeVariant("freemegb")},
					ExpireTimeout: int32(5000),
				}

				notify.SendNotification(conn, notif)
			} else if runtime.GOOS == "windows" {
				// file_location, _ := filepath.Abs("ui/freemegb.png")
				// notification := &toast.Notification{
				// 	AppID: "FreeMe!GB",
				// 	Title: "Unknown Instruction",
				// 	Message: fmt.Sprintf("INSTRUCTION: 0x%02X\nAt ROM Offset: %s",
				// 		cpu.INSTRUCTIONS[ROM.data[cpu.REGISTERS.PC]].Opcode, PCString),
				// 	Icon: file_location,
				// 	Actions: []toast.Action{
				// 		{"protocol", "Dismiss", ""},
				// 	},
				// }
				// err := notification.Push()
				// if err != nil {
				// 	Logger.Log(LogTypes.ERROR, err)
				// }
			}
			break
		}
		Logger.Logf(LogTypes.INFO, "Instruction: %s\n", cpu.INSTRUCTIONS[ROM.data[cpu.REGISTERS.PC]].Name)
		if cpu.DEBUG {
			cpu.REGISTERS.Print()
			time.Sleep(1 * time.Second)
		} else {
			time.Sleep(4 * time.Microsecond)
		}
		cpu.INSTRUCTIONS[ROM.data[cpu.REGISTERS.PC]].Exec()
		cpu.REGISTERS.PC++

	}
	//	finished <- true
}

// Reset will reset the CPU, INTERRUPTS and REGISTERS to their default values
func (cpu *CPUType) Reset() {

}
