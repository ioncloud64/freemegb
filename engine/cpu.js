class CPU {
  MMU = require("./components/memory").MMU;
  ROM = require("./components/memory").ROM;

  REGISTERS = require('./components/registers').REGISTERS;
  Operations = require('./components/operations');

  constructor() {}

  init() {


    // Initialize registers to standard values
    REGISTERS.AF = 0x01B0;
    REGISTERS.BC = 0x0013;
    REGISTERS.DE = 0x00D8;
    REGISTERS.HL = 0x014D;
    REGISTERS.SP = 0xFFFE;
    REGISTERS.PC = 0x0100;


    MMU.readByte(REGISTERS.F);
    console.log(REGISTERS);
    // REGISTERS.print();
    MMU.flags.SET(MMU.flags.ZERO);
    // REGISTERS.print();
    MMU.flags.SET(MMU.flags.CARRY);
    // REGISTERS.print();
    MMU.flags.CLEAR(MMU.flags.ZERO);
    // REGISTERS.print();

    // REGISTERS.print();
  }

  step() {
    this.Operations[ROM[REGISTERS.PC++]]();

    setTimeout(function() {
      console.log('hello world!');
    }, 5000);
  }
}

module.exports = new CPU();
