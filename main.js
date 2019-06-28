/**
*  Required imports
**/
const electron = require('electron');
const { execSync } = require('child_process');
const System = require('./engine/system');
var win;

/**
*  Creates the Electron Window
**/
function createWindow() {
  let GTK_THEME = execSync("gsettings get org.gnome.desktop.interface gtk-theme",
    {encoding: 'utf8'});

  // Create the browser window.
  win = new electron.BrowserWindow({
    title: "FreeMe!GB",
    titlebarstyle: 'hidden',
    width: 800,
    height: 600,
    icon: __dirname + '/resources/ioncloud64_icon.png',
    webPreferences: {
      nodeIntegration: true
    }
  });

  /**
  *  Build Application Menu
  **/
  var appmenu = electron.Menu.buildFromTemplate([
    {
    label: 'File',
    submenu: [
      {
        label: 'Open...',
        role: 'openFile',
        click() {
          const { dialog } = require('electron');
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
                loadROM(fileNames[0]);
                var columns = [
                  {id: "offset", name: "Offset", field: "offset"},
                  {id: "value", name: "Value", field: "value"}
                ];
                var options = {
                  enableColumnReorder: false,
                  forceFitColumns: true
                };
                var rows = [];
                for (var i = 0; i < System.cpu.ROM.length; i++) {
                  rows.push({
                    offset: i.toString(16).toUpperCase().padStart(6, '0'),
                    value: "0x" + System.cpu.ROM[i].toString(16).toUpperCase().padStart(4, '0')
                  });
                }
                console.log(rows);
                // var grid = new SlickGrid.Grid("#romGrid", data, columns, options);
                win.send('update-ROM_TABLE', rows, columns, options);
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
    },
    {
      label: 'Debug',
      submenu: [
        {
          label: 'Start/Resume',
          accelerator: 'F5',
          click() {
            System.emit('start');
          }
        },
        {
          label: 'Stop',
          accelerator: 'CmdOrCtrl+F5',
          click() {
            console.log("Stop");
          }
        },
        {
          label: 'Restart',
          accelerator: 'Shift+F5',
          click() {
            console.log("Restart");
          }
        },
        {
          label: 'Step',
          accelerator: 'F10',
          click() {

          }
        },
        {
          label: 'Dev Tools',
          accelerator: 'Shift+F12',
          click() {
            win.webContents.openDevTools();
          }
        }
      ]
    },
    {
      label: 'System',
      submenu: [
        {
          label: 'Run',
          accelerator: 'CmdOrCtrl+R',
          click() {
            System.emit('start');
          }
        }
      ]
    },
    {
    label: 'Tools',
    submenu: [
      {
        label: 'FreeMe!GB Settings',
        accelerator: 'CmdOrCtrl+S',
        click() {
          win.webContents.send('freemegb-settings');
        }
      }
    ]
  }
  ]);

  electron.Menu.setApplicationMenu(appmenu);
  // and load the index.html of the app.
  win.loadFile('main.html');
  win.webContents.on('did-finish-load', () => {
    win.webContents.send('GTK_THEME', GTK_THEME);
  });
  System.cpu.Registers.win = win;
}

/**
*  Loads the ROM file into the memory of the cpu's ROM
**/
function loadROM(fileName) {
  const fs = require('fs');
  var buffer = fs.readFileSync(fileName);
  for (const value of buffer) {
    System.cpu.ROM.push(value);
  }
}

/**
*  Creates the window when Electron.js is ready
**/
electron.app.on('ready', createWindow);
