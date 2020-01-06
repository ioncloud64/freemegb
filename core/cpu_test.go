package core

import (
  "testing"
)

func TestCPU(t *testing.T)  {
  if CPU.REGISTERS.AF != 0x01B0 {
    t.Errorf("CPU Register: AF is not initialized properly")
  }
  if CPU.REGISTERS.BC != 0x0013 {
    t.Errorf("CPU Register: BC is not initialized properly")
  }
  if CPU.REGISTERS.DE != 0x00D8 {
    t.Errorf("CPU Register: DE is not initialized properly")
  }
  if CPU.REGISTERS.HL != 0x014D {
    t.Errorf("CPU Register: HL is not initialized properly")
  }
  if CPU.REGISTERS.SP != 0xFFFE {
    t.Errorf("CPU Register: SP is not initialized properly")
  }
  if CPU.REGISTERS.PC != 0x0100 {
    t.Errorf("CPU Register: PC is not initialized properly")
  }
}
