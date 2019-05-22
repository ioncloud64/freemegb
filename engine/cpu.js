
class CPU {
  MMU = require("./components/memory").MMU;
  ROM = require("./components/memory").ROM;

  Operations = require('./components/operations');
  Registers = require('./components/registers');


  constructor() {}

  step() {
    Operations[ROM[Registers.PC++]]();
  }
}

module.exports = new CPU();
