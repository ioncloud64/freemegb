/**
*  Required imports
**/
const electron = require('electron');
const { execSync } = require('child_process');
const System = require('./engine/system');
const os = require('os');

/**
*  Required file globals
**/
var win;
var THEME = "";

/**
*  Start gathering System Information for Theming
**/
setTheme();

// TODO: Disable menu items by default

// TODO: Read ROM file asynchronously

// TODO: Show progressbar based upon how much of the file is read

// TODO: Create a poll every focus/minute for theme changes

/**
*  Creates the Electron Window
**/
function createWindow() {
  // Create the browser window.
  win = new electron.BrowserWindow({
    title: "FreeMe!GB",
    titlebarstyle: 'hidden',
    width: 800,
    height: 600,
    icon: __dirname + '/resources/ioncloud64_512.png',
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
          try {
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
                    offset: "0x" + i.toString(16).toUpperCase().padStart(6, '0'),
                    value: "0x" + System.cpu.ROM[i].toString(16).toUpperCase().padStart(4, '0')
                  });
                }
                console.log("Sending ROM Data to UI");
                win.send('update-ROM_TABLE', rows, columns, options);
              }
            });
            } catch (err) {
              console.log(err);
            }
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
    win.webContents.send('Process_Theme', THEME);
    console.log("THEME:", THEME);
  });
  System.cpu.REGISTERS.win = win;
}

/**
*  Sets the theme based upon OS
*  Works on Windows (10; older versions default to light theme), Linux,
*  and newer versions of Mac OS X
**/
async function setTheme() {
  var os_type = os.platform();
  console.log(os.release());
  if (os_type === "darwin") {
    THEME = systemPreferences.isDarkMode() ? 'dark' : 'light';
  } else if (os_type === "win32") {
    var [dwMajorVersion, dwMinorVersion, dwBuildNumber] = os.release().split(".").map(Number);
    if (dwBuildNumber >= 16299) {
      var Registry = require('winreg');
      var regKey = new Registry({
        hive: Registry.HKCU,
        key: '\\Software\\Microsoft\\Windows\\CurrentVersion\\Themes\\Personalize'
      })
      regKey.values((err, items) => {
        if (err) {
          console.error("Error:", err);
        } else {
          for (var i = 0; i < items.length; i++) {
            console.log('ITEM: '+items[i].name+'\t'+items[i].type+'\t'+items[i].value);
            if (items[i].name === "AppsUseLightTheme") {
              THEME = items[i].value | 0 ? 'light' : 'dark';
            }
          }
        }
      });
    } else {
      THEME = "light";
    }
  } else if (os_type === "linux") {
    var OS_THEME = execSync("gsettings get org.gnome.desktop.interface gtk-theme", {encoding: 'utf8'});
    THEME = OS_THEME.toLowerCase().includes("dark") ? "dark" : "light";
  } else {
    console.error("Undetectable Operating System: Defaulting to Light Theme");
    THEME = "light";
  }
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
