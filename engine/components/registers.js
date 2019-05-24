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
  get SP() {
    return this.sp;
  }
  get PC() {
    return this.pc;
  }

  /*
   * Set block for combined registers
   */
  set AF(val) {
    if (this.win !== undefined) {
      this.win.webContents.send('update-AF', "0x" + val.toString(16).toUpperCase().padStart(4, '0'));
    }
    this.A = val >> 8 & 0xFF;
    this.F = val & 0x00FF;
  }
  set BC(val) {
    if (this.win !== undefined) {
      this.win.webContents.send('update-BC', "0x" + val.toString(16).toUpperCase().padStart(4, '0'));
    }
    this.B = val >> 8 & 0xFF;
    this.C = val & 0x00FF;
  }
  set DE(val) {
    if (this.win !== undefined) {
      this.win.webContents.send('update-DE', "0x" + val.toString(16).toUpperCase().padStart(4, '0'));
    }
    this.D = val >> 8 & 0xFF;
    this.E = val & 0x00FF;
  }
  set HL(val) {
    if (this.win !== undefined) {
      this.win.webContents.send('update-HL', "0x" + val.toString(16).toUpperCase().padStart(4, '0'));
    }
    this.H = val >> 8 & 0xFF;
    this.L = val & 0x00FF;
  }
  set SP(val) {
    if (this.win !== undefined) {
      this.win.webContents.send('update-SP', "0x" + val.toString(16).toUpperCase().padStart(4, '0'));
    }
    this.sp = val;
  }
  set PC(val) {
    if (this.win !== undefined) {
      this.win.webContents.send('update-PC', "0x" + val.toString(16).toUpperCase().padStart(4, '0'));
    }
    this.pc = val;
  }

  constructor() {
    // Initialize registers to 0 values
    this.AF = 0;
    this.BC = 0;
    this.DE = 0;
    this.HL = 0;
    this.SP = 0;
    this.PC = 0;
    this.win = undefined;
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
