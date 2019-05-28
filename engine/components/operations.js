/**
*  An array of functions that hold instructions for the CPU
**/
const Operations = [
  // 0x00 NOP
  function() {
    return;
  },
  // 0x01
  function() {
    console.log("LOAD BC NN");
    Registers.BC = ROM[++Registers.PC] << 8 | ROM[++Registers.PC];
  },
  // 0x02
  function() {
    console.log("LOAD BC A");
    Registers.B = Registers.A;
  },
  // 0x03
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x04
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x05
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x06
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x07
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x08
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x09
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x0A
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x0B
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x0C
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x0D
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x0E
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x0F
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x10
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x11
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x12
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x13
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x14
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x15
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x16
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x17
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x18
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x19
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x1A
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x1B
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x1C
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x1D
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x1E
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x1F
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x20
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x21
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x22
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x23
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x24
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x25
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x26
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x27
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x28
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x29
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x2A
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x2B
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x2C
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x2D
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x2E
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x2F
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x30
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x31
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x32
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x33
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x34
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x35
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x36
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x37
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x38
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x39
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x3A
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x3B
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x3C
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x3D
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x3E
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x3F
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x40
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x41
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x42
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x43
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x44
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x45
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x46
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x47
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x48
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x49
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x4A
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x4B
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x4C
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x4D
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x4E
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x4F
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x50
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x51
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x52
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x53
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x54
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x55
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x56
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x57
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x58
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x59
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x5A
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x5B
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x5C
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x5D
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x5E
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x5F
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x60
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x61
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x62
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x63
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x64
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x65
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x66
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x67
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x68
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x69
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x6A
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x6B
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x6C
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x6D
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x6E
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x6F
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x70
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x71
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x72
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x73
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x74
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x75
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x76
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x77
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x78
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x79
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x7A
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x7B
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x7C
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x7D
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x7E
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x7F
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x80
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x81
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x82
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x83
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x84
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x85
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x86
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x87
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x88
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x89
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x8A
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x8B
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x8C
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x8D
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x8E
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x8F
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x90
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x91
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x92
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x93
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x94
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x95
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x96
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x97
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x98
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x99
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x9A
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x9B
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x9C
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x9D
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x9E
  function() {
    throw Error("Instruction not implemented");
  },
  // 0x9F
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xA0
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xA1
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xA2
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xA3
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xA4
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xA5
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xA6
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xA7
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xA8
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xA9
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xAA
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xAB
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xAC
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xAD
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xAE
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xAF
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xB0
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xB1
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xB2
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xB3
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xB4
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xB5
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xB6
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xB7
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xB8
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xB9
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xBA
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xBB
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xBC
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xBD
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xBE
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xBF
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xC0
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xC1
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xC2
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xC3
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xC4
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xC5
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xC6
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xC7
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xC8
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xC9
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xCA
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xCB
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xCC
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xCD
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xCE
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xCF
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xD0
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xD1
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xD2
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xD3
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xD4
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xD5
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xD6
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xD7
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xD8
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xD9
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xDA
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xDB
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xDC
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xDD
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xDE
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xDF
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xE0
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xE1
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xE2
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xE3
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xE4
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xE5
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xE6
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xE7
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xE8
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xE9
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xEA
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xEB
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xEC
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xED
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xEE
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xEF
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xF0
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xF1
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xF2
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xF3
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xF4
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xF5
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xF6
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xF7
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xF8
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xF9
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xFA
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xFB
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xFC
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xFD
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xFE
  function() {
    throw Error("Instruction not implemented");
  },
  // 0xFF
  function() {
    throw Error("Instruction not implemented");
  }
];

module.exports = Operations;
