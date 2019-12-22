package engine

import (
	"fmt"
)

type ROMType struct {
	data			[]byte
	model			[]interface{}
	modelColumns	[]string
}

func (rom *ROMType) BuildROMModel () {
	model := []interface{}{}
	
	for i := 0x0000; i < len(rom.data); i++ {
		row := []string{}
		if CPU.INSTRUCTIONS[rom.data[i]].NumOperands == 1 {
			i += int(CPU.INSTRUCTIONS[rom.data[i]].NumOperands)
			row = append(row, fmt.Sprintf("0x%04X", i))
			row = append(row, fmt.Sprintf("%s, 0x%02X", CPU.INSTRUCTIONS[rom.data[i]].Name, rom.data[i+1]))
		} else if CPU.INSTRUCTIONS[rom.data[i]].NumOperands == 2 {
			i += int(CPU.INSTRUCTIONS[rom.data[i]].NumOperands)
			row = append(row, fmt.Sprintf("0x%04X", i))
			row = append(row, fmt.Sprintf("%s 0x%02X 0x%02X", CPU.INSTRUCTIONS[rom.data[i]].Name, rom.data[i+1], rom.data[i+2]))
		} else {
			if CPU.INSTRUCTIONS[rom.data[i]].Name == "UNKNOWN" {
				row = append(row, fmt.Sprintf("0x%04X", i))
				row = append(row, fmt.Sprintf("%s: 0x%02X", CPU.INSTRUCTIONS[rom.data[i]].Name, rom.data[i]))
			} else {
				row = append(row, fmt.Sprintf("0x%04X", i))
				row = append(row, fmt.Sprintf("%s", CPU.INSTRUCTIONS[rom.data[i]].Name))
			}
		}
		model = append(model, row)
	}
	rom.model = model
}

var ROM = ROMType {
	data:			[]byte{},
	model:			nil,
	modelColumns:	[]string{"Offset", "Instruction"},
}
