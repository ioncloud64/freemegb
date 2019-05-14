class Registers {
  /*
   * Get block for combined registers
   */
  get AF() {
    return this.A << 8 | this.F;
  }
  get BC() {
    return this.B << 8 | this.C;
  }
  get DE() {
    return this.D << 8 | this.E;
  }
  get HL() {
    return this.H << 8 | this.L;
  }

  /*
   * Set block for combined registers
   */
  set AF(val) {
    this.A = val >> 8 & 0xFF;
    this.F = val & 0x00FF;
  }
  set BC(val) {
    this.B = val >> 8 & 0xFF;
    this.C = val & 0x00FF;
  }
  set DE(val) {
    this.D = val >> 8 & 0xFF;
    this.E = val & 0x00FF;
  }
  set HL(val) {
    this.H = val >> 8 & 0xFF;
    this.L = val & 0x00FF;
  }

  constructor() {
    // Initialize registers to standard values
    this.AF = 0x01B0
    this.BC = 0x0013
    this.DE = 0x00D8
    this.HL = 0x014D
    this.SP = 0xFFFE;
    this.PC = 0x0100;
  }

  print() {
    console.log("AF: 0x%s", this.AF.toString(16).padStart(4, '0').toUpperCase());
    console.log("BC: 0x%s", this.BC.toString(16).padStart(4, '0').toUpperCase());
    console.log("DE: 0x%s", this.DE.toString(16).padStart(4, '0').toUpperCase());
    console.log("HL: 0x%s", this.HL.toString(16).padStart(4, '0').toUpperCase());
    console.log("SP: 0x%s", this.SP.toString(16).padStart(4, '0').toUpperCase());
    console.log("PC: 0x%s", this.PC.toString(16).padStart(4, '0').toUpperCase());
  }
}

module.exports = new Registers();
