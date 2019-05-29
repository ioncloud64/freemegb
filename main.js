/**
*  Setup console for GUI and logs
**/
console.clear();
const fs = require('fs');
const home = require('os').homedir();
const prgmdir = home + '/freemegb';
const prgmlogs = prgmdir + '/logs';
const logfile = prgmlogs + '/' + new Date().toISOString() + '.log';
if (!fs.existsSync(prgmdir)){
    fs.mkdirSync(prgmdir);
}
if (!fs.existsSync(prgmlogs)){
    fs.mkdirSync(prgmlogs);
}
fs.writeFile(logfile, "", function (err) {
  if (err) throw err;
  console.log(logfile + ' is created successfully.');
});

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
                tData = [];
                tDataString = "";
                for (var i = 0; i < System.cpu.ROM.length; i++) {
                  romLocData = System.cpu.ROM[i].toString(16).toUpperCase().padStart(4, '0');
                  tDataString += "<tr onclick='addBreakpoint(this);'>\n";
                  tDataString += "<th scope='row'>0x" + i.toString(16).toUpperCase().padStart(6, '0') + "</th>";
                  tDataString += "<td id='" + romLocData + "'>0x" + romLocData + "</td>\n";
                  tDataString += "</tr>\n";
                  tData.push(tDataString);
                  tDataString = "";
                }
                win.send('update-ROM_TABLE', tData);
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
