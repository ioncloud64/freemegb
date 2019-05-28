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
        this.cpu.Registers.print();
        console.error(e);
        break;
      } finally {
      }
    }
  }

}

module.exports = new System();
