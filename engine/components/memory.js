/**
 *  Array for ROM file Data
 **/
ROM = [];

/**
 *  MMU controls memory in RAM banks and provides functions
 *  for writing/reading bytes.
 **/
MMU = {
  flags: {
    CLEAR(flag, REGISTERS) {
      REGISTERS.F &= ~flag;
    },
    SET(flag, REGISTERS) {
      REGISTERS.F |= flag;
    },
    ISSET(flag, REGISTERS) {
      return Boolean(REGISTERS.F & flag);
    },

    ZERO: 0x80,
    SUBTRACT: 0x40,
    HALF_CARRY: 0x20,
    CARRY: 0x10
  },

  ADD(destination, source) {
    destination += source;
  },

  ADD_C(destination, source) {
    destination += source;
  },

  writeByte(address, value) {
    switch (address & 0xF000) {
      //0x0000 to 0x1000
      case 0x0000:
      case 0x1000:

        break;
        //0x2000 to 0x3000
      case 0x2000:
      case 0x3000:

        break;
        //0x4000 to 0x5000
      case 0x4000:
      case 0x5000:

        break;
        //0x6000 to 0x7000
      case 0x6000:
      case 0x7000:

        break;
        //0x8000 to 0x9000
      case 0x8000:
      case 0x9000:

        break;
        //0xA000 to 0xB000
      case 0xA000:
      case 0xB000:

        break;
        //0xC000 to 0xE000
      case 0xC000:
      case 0xD000:
      case 0xE000:

        break;
        //0xF000
      case 0xF000:

        break;
      default:

    }
  },

  readByte(address, value) {
    switch (address & 0xF000) {
      //0x0000 to 0x1000
      case 0x0000:
      case 0x1000:

        break;
        //0x2000 to 0x3000
      case 0x2000:
      case 0x3000:

        break;
        //0x4000 to 0x5000
      case 0x4000:
      case 0x5000:

        break;
        //0x6000 to 0x7000
      case 0x6000:
      case 0x7000:

        break;
        //0x8000 to 0x9000
      case 0x8000:
      case 0x9000:

        break;
        //0xA000 to 0xB000
      case 0xA000:
      case 0xB000:

        break;
        //0xC000 to 0xE000
      case 0xC000:
      case 0xD000:
      case 0xE000:

        break;
        //0xF000
      case 0xF000:

        break;
      default:

    }
  }
};

module.exports = {
  MMU,
  ROM
};
