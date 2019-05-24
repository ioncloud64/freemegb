class CPU {
  MMU = require("./components/memory").MMU;
  ROM = require("./components/memory").ROM;

  Operations = require('./components/operations');
  Registers = require('./components/registers');


  constructor() {}

  init() {
    // Initialize registers to standard values
    this.Registers.AF = 0x01B0;
    this.Registers.BC = 0x0013;
    this.Registers.DE = 0x00D8;
    this.Registers.HL = 0x014D;
    this.Registers.SP = 0xFFFE;
    this.Registers.PC = 0x0100;
  }

  step() {
    this.Operations[ROM[this.Registers.PC++]]();
  }
}

module.exports = new CPU();
