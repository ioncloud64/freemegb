class CPU {
  MMU = require("./components/memory").MMU;
  ROM = require("./components/memory").ROM;

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


    this.MMU.readByte(REGISTERS.F);
    // console.log(REGISTERS);
    // REGISTERS.print();
    this.MMU.flags.SET(this.MMU.flags.ZERO);
    // REGISTERS.print();
    this.MMU.flags.SET(this.MMU.flags.CARRY);
    // REGISTERS.print();
    this.MMU.flags.CLEAR(this.MMU.flags.ZERO);
    // REGISTERS.print();

    // REGISTERS.print();
  }

  step(debug) {
    console.log("CPU STEP");
    // console.log(this.REGISTERS);
    if (debug) {
      this.REGISTERS.print();
    }
    this.Operations[ROM[this.REGISTERS.PC]](this.REGISTERS, this.ROM);
    this.REGISTERS.PC++;
  }
}

module.exports = new CPU();
