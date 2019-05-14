const electron = require('electron');

function createWindow() {
  // Create the browser window.
  let win = new electron.BrowserWindow({
    titlebarstyle: 'hidden',
    width: 800,
    height: 600,
    icon: __dirname + '/img/ioncloud64_icon.png',
    webPreferences: {
      nodeIntegration: true
    }
  })

  var appmenu = electron.Menu.buildFromTemplate([{
    label: 'File',
    submenu: [{
        label: 'Open...',
        role: 'openFile',
        click() {
          const {
            dialog
          } = require('electron')
          const fs = require('fs')
          dialog.showOpenDialog(win, {
              filters: [{
                name: 'GameBoy Game',
                extensions: ['gb']
              }]
            },
            function(fileNames) {

              // fileNames is an array that contains all the selected
              if (fileNames === undefined) {
                return;
              } else {
                readFile(fileNames[0]);
              }
            });
        },
        accelerator: 'CmdOrCtrl+O'
      },
      {
        label: 'Open Recent'
      },
      {
        label: 'Quit',
        role: 'quit'

      }
    ]
  }])

  electron.Menu.setApplicationMenu(appmenu);
  // and load the index.html of the app.
  win.loadFile('main.html')
}

function readFile(fileName) {
  const fs = require('fs');
  var buffer = fs.readFileSync(fileName);
  for (var i = 0x100; i < 0x100 + 20; i++) {
    console.log(buffer[i].toString(16).toUpperCase());
  }
}

electron.app.on('ready', createWindow)

const cpu = require('./engine/cpu.js');

cpu.step();

// console.log('A: 0x%s', registers.AF.A.toString(16));
// console.log('F: 0x%s', registers.AF.F.toString(16));
// console.log('F: 0x%s', registers.F.toString(16));
// registers.AF = registers.A.toString(16) + registers.F.toString(16);
// console.log('AF: 0x%s', registers.AF.padStart(4, '0').toUpperCase());
