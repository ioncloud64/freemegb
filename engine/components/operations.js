/**
 *  An array of functions that hold instructions for the CPU
 **/
const Operations = [
  // 0x00 NOP
  function(REGISTERS, ROM, MMU) {
    console.log("NOP");
  },
  // 0x01
  function(REGISTERS, ROM, MMU) {
    console.log("LOAD BC NN");
    REGISTERS.BC = ROM[++REGISTERS.PC] << 8 | ROM[++REGISTERS.PC];
  },
  // 0x02
  function(REGISTERS, ROM, MMU) {
    console.log("LOAD BC A");
    REGISTERS.B = REGISTERS.A;
  },
  // 0x03
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x04
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x05
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x06
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x07
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x08
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x09
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x0A
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x0B
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x0C
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x0D
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x0E
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x0F
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x10
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x11
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x12
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x13
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x14
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x15
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x16
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x17
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x18
  function(REGISTERS, ROM, MMU) {
    console.log("JUMP 0x" + (REGISTERS.PC + ROM[REGISTERS.PC + 1]).toString(16).toUpperCase().padStart(4, '0'));
    REGISTERS.PC = (REGISTERS.PC + ROM[REGISTERS.PC + 1]) - 1;
  },
  // 0x19
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x1A
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x1B
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x1C
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x1D
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x1E
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x1F
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x20
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x21
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x22
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x23
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x24
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x25
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x26
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x27
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x28
  function(REGISTERS, ROM, MMU) {
    console.log("JUMP Z 0x" + (REGISTERS.PC + ROM[REGISTERS.PC + 1]).toString(16).toUpperCase().padStart(4, '0'));
    if (MMU.flags.ISSET(MMU.flags.ZERO, REGISTERS)) {
      REGISTERS.PC = (REGISTERS.PC + ROM[REGISTERS.PC + 1] - 1);
    }
  },
  // 0x29
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x2A
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x2B
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x2C
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x2D
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x2E
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x2F
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x30
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x31
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x32
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x33
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x34
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x35
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x36
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x37
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x38
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x39
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x3A
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x3B
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x3C
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x3D
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x3E
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x3F
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x40
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x41
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x42
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x43
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x44
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x45
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x46
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x47
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x48
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x49
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x4A
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x4B
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x4C
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x4D
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x4E
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x4F
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x50
  function(REGISTERS, ROM, MMU) {
    console.log("LOAD D B");
    REGISTERS.D = REGISTERS.B;
  },
  // 0x51
  function(REGISTERS, ROM, MMU) {
    console.log("LOAD D C");
    REGISTERS.D = REGISTERS.C;
  },
  // 0x52
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x53
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x54
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x55
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x56
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x57
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x58
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x59
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x5A
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x5B
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x5C
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x5D
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x5E
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x5F
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x60
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x61
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x62
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x63
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x64
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x65
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x66
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x67
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x68
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x69
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x6A
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x6B
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x6C
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x6D
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x6E
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x6F
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x70
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x71
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x72
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x73
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x74
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x75
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x76
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x77
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x78
  function(REGISTERS, ROM, MMU) {
    console.log("LOAD A B");
    REGISTERS.A = REGISTERS.B;
  },
  // 0x79
  function(REGISTERS, ROM, MMU) {
    console.log("LOAD A C");
    REGISTERS.A = REGISTERS.C;
  },
  // 0x7A
  function(REGISTERS, ROM, MMU) {
    console.log("LOAD A D");
    REGISTERS.A = REGISTERS.D;
  },
  // 0x7B
  function(REGISTERS, ROM, MMU) {
    console.log("LOAD A E");
    REGISTERS.A = REGISTERS.E;
  },
  // 0x7C
  function(REGISTERS, ROM, MMU) {
    console.log("LOAD A H");
    REGISTERS.A = REGISTERS.H;
  },
  // 0x7D
  function(REGISTERS, ROM, MMU) {
    console.log("LOAD A L");
    REGISTERS.A = REGISTERS.L;
  },
  // 0x7E
  function(REGISTERS, ROM, MMU) {
    console.log("LOAD A HL");
    REGISTERS.A = MMU.readByte(REGISTERS.HL);
  },
  // 0x7F
  function(REGISTERS, ROM, MMU) {
    console.log("LOAD A A");
  },
  // 0x80
  function(REGISTERS, ROM, MMU) {
    console.log("ADD A B");
    MMU.ADD(REGISTERS.A, REGISTERS.B);
  },
  // 0x81
  function(REGISTERS, ROM, MMU) {
    console.log("ADD A C");
    MMU.ADD(REGISTERS.A, REGISTERS.C);
  },
  // 0x82
  function(REGISTERS, ROM, MMU) {
    console.log("ADD A D");
    MMU.ADD(REGISTERS.A, REGISTERS.D);
  },
  // 0x83
  function(REGISTERS, ROM, MMU) {
    console.log("ADD A E");
    MMU.ADD(REGISTERS.A, REGISTERS.E);
  },
  // 0x84
  function(REGISTERS, ROM, MMU) {
    console.log("ADD A H");
    MMU.ADD(REGISTERS.A, REGISTERS.H);
  },
  // 0x85
  function(REGISTERS, ROM, MMU) {
    console.log("ADD A L");
    MMU.ADD(REGISTERS.A, REGISTERS.L);
  },
  // 0x86
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x87
  function(REGISTERS, ROM, MMU) {
    console.log("ADD A A");
    MMU.ADD(REGISTERS.A, REGISTERS.A);
  },
  // 0x88
  function(REGISTERS, ROM, MMU) {
    console.log("ADC A B");
    MMU.ADD_C(REGISTERS.A, REGISTERS.B);
  },
  // 0x89
  function(REGISTERS, ROM, MMU) {
    console.log("ADC A C");
    MMU.ADD_C(REGISTERS.A, REGISTERS.C);
  },
  // 0x8A
  function(REGISTERS, ROM, MMU) {
    console.log("ADC A D");
    MMU.ADD_C(REGISTERS.A, REGISTERS.D);
  },
  // 0x8B
  function(REGISTERS, ROM, MMU) {
    console.log("ADC A E");
    MMU.ADD_C(REGISTERS.A, REGISTERS.E);
  },
  // 0x8C
  function(REGISTERS, ROM, MMU) {
    console.log("ADC A H");
    MMU.ADD_C(REGISTERS.A, REGISTERS.H);
  },
  // 0x8D
  function(REGISTERS, ROM, MMU) {
    console.log("ADC A L");
    MMU.ADD_C(REGISTERS.A, REGISTERS.L);
  },
  // 0x8E
  function(REGISTERS, ROM, MMU) {
    console.log("ADC A HL");
    MMU.ADD_C(REGISTERS.A, MMU.readByte(REGISTERS.HL));
  },
  // 0x8F
  function(REGISTERS, ROM, MMU) {
    console.log("ADC A A");
    MMU.ADD_C(REGISTERS.A, REGISTERS.A);
  },
  // 0x90
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x91
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x92
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x93
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x94
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x95
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x96
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x97
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x98
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x99
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x9A
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x9B
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x9C
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x9D
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x9E
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0x9F
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xA0
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xA1
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xA2
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xA3
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xA4
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xA5
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xA6
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xA7
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xA8
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xA9
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xAA
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xAB
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xAC
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xAD
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xAE
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xAF
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xB0
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xB1
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xB2
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xB3
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xB4
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xB5
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xB6
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xB7
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xB8
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xB9
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xBA
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xBB
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xBC
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xBD
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xBE
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xBF
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xC0
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xC1
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xC2
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xC3
  function(REGISTERS, ROM, MMU) {
    REGISTERS.PC = (ROM[REGISTERS.PC + 1] | ROM[REGISTERS.PC + 2] << 8) - 1;
    console.log("JUMP " + "0x" + (REGISTERS.PC + 1).toString(16).toUpperCase().padStart(4, '0'));
  },
  // 0xC4
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xC5
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xC6
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xC7
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xC8
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xC9
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xCA
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xCB
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xCC
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xCD
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xCE
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xCF
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xD0
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xD1
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xD2
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xD3
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xD4
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xD5
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xD6
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xD7
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xD8
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xD9
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xDA
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xDB
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xDC
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xDD
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xDE
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xDF
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xE0
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xE1
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xE2
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xE3
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xE4
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xE5
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xE6
  function(REGISTERS, ROM, MMU) {
    console.log(REGISTERS.PC);
    var num = ROM[++REGISTERS.PC]
    console.log(num);
    console.log("AND A 0x" + num.toString(16).toUpperCase().padStart(4, '0'));
    REGISTERS.A &= num;
  },
  // 0xE7
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xE8
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xE9
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xEA
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xEB
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xEC
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xED
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xEE
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xEF
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xF0
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xF1
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xF2
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xF3
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xF4
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xF5
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xF6
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xF7
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xF8
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xF9
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xFA
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xFB
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xFC
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xFD
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  },
  // 0xFE
  function(REGISTERS, ROM, MMU) {
    var operand = ROM[++REGISTERS.PC];
    console.log("CP A 0x" + operand.toString(16).toUpperCase().padStart(2, '0'));
    if (REGISTERS.A == operand) {
      MMU.flags.SET(MMU.flags.ZERO, REGISTERS);
    } else if (REGISTERS.A < operand) {
      MMU.flags.SET(MMU.flags.CARRY, REGISTERS);
    }
  },
  // 0xFF
  function(REGISTERS, ROM, MMU) {
    throw Error("Instruction not implemented 0x" + ROM[REGISTERS.PC].toString(16).toUpperCase().padStart(2, '0'));
  }
];

module.exports = Operations;
