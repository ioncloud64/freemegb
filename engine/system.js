const EventEmitter = require('events');
const ipcMain = require('electron').ipcMain;

/**
 *  System controls the CPU and GPU while
 *  also providing a platform for debugging.
 **/
class System extends EventEmitter {
  cpu = require('./cpu');
  gpu = require('./gpu');

  isDebug = false;
  running = false;
  paused = false;
  next = false;
  breakpoints = [];

  constructor() {
    super();
    this.on('start', (debug = false) => {
      this.cpu.init();
      this.start(debug);
    });
    this.on('pause', () => {
      paused = true;
    });
    this.on('resume', () => {
      paused = false;
      this.start();
    });
    ipcMain.on('add-breakpoint', (address) => {
      if (!this.breakpoints.includes(address)) {
        this.breakpoints.push(address);
        console.log(this.breakpoints);
        return true;
      } else {
        return false;
      }
    });
  }

  start(debug = false) {
    while (!this.paused) {
      try {
        this.cpu.step(debug);
        if (this.next && debug) {
          this.next = false;
          break;
        }

      } catch (e) {
        this.cpu.REGISTERS.print();
        console.log(this.cpu.ROM[this.cpu.REGISTERS.PC].toString(16).toUpperCase().padStart(4, '0'));
        console.error(e);
        break;
      } finally {}
    }
  }

  info() {
    // TODO: build a formatted string to display information about the system to the user.
    var infoStr = "";

    infoStr += "{0}{1}".format("8KB".padStart(4, ' '), "Main RAM:");
    infoStr += "{0}{1}".format("8KB".padStart(4, ' '), "Video RAM:");
    infoStr += "{0}{1}".format("8KB".padStart(4, ' '), "Screen Size:");
    infoStr += "{0}{1}".format("8KB".padStart(4, ' '), "Main RAM:");
    infoStr += "{0}{1}".format("8KB".padStart(4, ' '), "Main RAM:");
    infoStr += "{0}{1}".format("8KB".padStart(4, ' '), "Main RAM:");
    infoStr += "{0}{1}".format("8KB".padStart(4, ' '), "Main RAM:");
    infoStr += "{0}{1}".format("8KB".padStart(4, ' '), "Main RAM:");
    infoStr += "{0}{1}".format("8KB".padStart(4, ' '), "Main RAM:");
    infoStr += "{0}{1}".format("8KB".padStart(4, ' '), "Main RAM:");
    infoStr += "{0}{1}".format("8KB".padStart(4, ' '), "Main RAM:");
  }

}

module.exports = new System();
