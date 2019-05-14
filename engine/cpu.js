const {
  Operations,
  Registers
} = require('./components/operations.js');
const {
  Memory,
  ROM
} = require("./components/memory.js");
class CPU {

  constructor() {}

  step() {

    // while (true) {
    Operations[0x00]();
    Operations
    Registers.PC++;
    // }
  }
}

module.exports = new CPU();
