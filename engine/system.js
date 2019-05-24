const EventEmitter = require('events');
class System extends EventEmitter {
  cpu = require('./cpu');
  gpu = require('./gpu');

  running = false;
  paused = false;
  next = false;

  constructor() {
    super();
    this.on('start', () => {
      this.cpu.init();
      this.start();
    });
    this.on('pause', () => {
      this.start();
    });
    this.on('resume', () => {
      this.start();
    });
  }

  start() {
    while (!this.paused) {
      try {
        this.cpu.step();
        if (this.next) {
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
