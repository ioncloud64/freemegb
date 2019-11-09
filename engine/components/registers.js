/**
 *  Registers hold data in 8-bit values and are
 *  combined respectively for the real-world system
 **/
REGISTERS = new class {


  /**
   *  16-bit Register AF
   *  A: 8-bit
   *  F: 8-bit - Flag Register
   **/
  set AF(val) {
    if (this.win !== undefined) {
      this.win.webContents.
      send('update-AF', "0x" + val.toString(16).toUpperCase().padStart(4, '0'));
    }
    this.A = val >> 8 & 0xFF;
    this.F = val & 0x00FF;
  }
  get AF() {
    return this.A << 8 | this.F;
  }

  /**
   *  16-bit Register BC
   *  B: 8-bit
   *  C: 8-bit
   **/
  set BC(val) {
    if (this.win !== undefined) {
      this.win.webContents.
      send('update-BC', "0x" + val.toString(16).toUpperCase().padStart(4, '0'));
    }
    this.B = val >> 8 & 0xFF;
    this.C = val & 0x00FF;
  }
  get BC() {
    return this.B << 8 | this.C;
  }

  /**
   *  16-bit Register DE
   *  D: 8-bit
   *  E: 8-bit
   **/
  set DE(val) {
    if (this.win !== undefined) {
      this.win.webContents.
      send('update-DE', "0x" + val.toString(16).toUpperCase().padStart(4, '0'));
    }
    this.D = val >> 8 & 0xFF;
    this.E = val & 0x00FF;
  }
  get DE() {
    return this.D << 8 | this.E;
  }

  /**
   *  16-bit Register HL
   *  H: 8-bit
   *  L: 8-bit
   **/
  set HL(val) {
    if (this.win !== undefined) {
      this.win.webContents.
      send('update-HL', "0x" + val.toString(16).toUpperCase().padStart(4, '0'));
    }
    this.H = val >> 8 & 0xFF;
    this.L = val & 0x00FF;
  }
  get HL() {
    return this.H << 8 | this.L;
  }

  /**
   *  16-bit Register SP - Stack Pointer
   **/
  set SP(val) {
    if (this.win !== undefined) {
      this.win.webContents.
      send('update-SP', "0x" + val.toString(16).toUpperCase().padStart(4, '0'));
    }
    this.sp = val;
  }
  get SP() {
    return this.sp;
  }

  /**
   *  16-bit Register PC - Program Counter
   **/
  set PC(val) {
    if (this.win !== undefined) {
      this.win.webContents.
      send('update-PC', "0x" + val.toString(16).toUpperCase().padStart(4, '0'));
    }
    this.pc = val;
  }
  get PC() {
    return this.pc;
  }

  /**
   *  Data initialization
   **/
  constructor() {
    this.AF = 0;
    this.BC = 0;
    this.DE = 0;
    this.HL = 0;
    this.SP = 0;
    this.PC = 0;
    this.win = undefined;

    /**
     *  Prints the registers to the console preformatted
     **/
    this.print = function() {
      console.log("AF: 0x%s", this.AF.toString(16).padStart(4, '0').toUpperCase());
      console.log("BC: 0x%s", this.BC.toString(16).padStart(4, '0').toUpperCase());
      console.log("DE: 0x%s", this.DE.toString(16).padStart(4, '0').toUpperCase());
      console.log("HL: 0x%s", this.HL.toString(16).padStart(4, '0').toUpperCase());
      console.log("SP: 0x%s", this.SP.toString(16).padStart(4, '0').toUpperCase());
      console.log("PC: 0x%s", this.PC.toString(16).padStart(4, '0').toUpperCase());
    };
  }

}

module.exports = {
  REGISTERS
};
