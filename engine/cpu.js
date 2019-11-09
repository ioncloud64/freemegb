
class CPU {
  ROM = require("./components/memory").ROM;
  MMU = require("./components/memory").MMU;
  REGISTERS = require('./components/registers').REGISTERS;
  Operations = require('./components/operations');

  constructor() {}

  init() {
    // Initialize registers to standard values
    this.REGISTERS.AF = 0x01B0;
    this.REGISTERS.BC = 0x0013;
    this.REGISTERS.DE = 0x00D8;
    this.REGISTERS.HL = 0x014D;
    this.REGISTERS.SP = 0xFFFE;
    this.REGISTERS.PC = 0x0100;

  }

  step(debug) {
    console.log("CPU STEP");
    // console.log(this.REGISTERS);
    if (debug) {
      this.REGISTERS.print();
    }
    this.Operations[ROM[this.REGISTERS.PC]](this.REGISTERS, this.ROM, this.MMU);
    this.REGISTERS.PC++;
  }
}

module.exports = new CPU();
