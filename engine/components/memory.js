/**
*  Array for ROM file Data
**/
ROM = [];

/**
*  MMU controls memory in RAM banks and provides functions
*  for writing/reading bytes.
**/
MMU = new class {
  flags = {
      SUBTRACT: 0x01,
      ADD: 0x08
  }

  writeByte(address, value) {
    switch (address & 0xF000) {
      case 0x0000: case 0x1000:

      break;
      case 0x2000: case 0x3000:

      break;
      case 0x4000: case 0x5000:

      break;
      case 0x6000: case 0x7000:

      break;
      case 0x8000: case 0x9000:

      break;
      case 0xA000: case 0xB000:

      break;
      case 0xC000: case 0xD000: case 0xE000:

      break;
      case 0xF000:

      break;
      default:

    }
  }
  constructor() {

  }
}();

module.exports = { MMU, ROM };
