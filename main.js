const electron = require('electron');
const { execSync } = require('child_process');
const System = require('./engine/system');
var win;

function createWindow() {
  let GTK_THEME = execSync("gsettings get org.gnome.desktop.interface gtk-theme", {encoding: 'utf8'});


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

  var appmenu = electron.Menu.buildFromTemplate([
    {
    label: 'File',
    submenu: [
      {
        label: 'Open...',
        role: 'openFile',
        click() {
          const { dialog } = require('electron');
          const fs = require('fs');
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
                tData = [];
                tDataString = "";
                for (var i = 0; i < System.cpu.ROM.length; i++) {
                  tDataString += "<tr>\n";
                  tDataString += "<th scope='row'>0x" + i.toString(16).toUpperCase().padStart(6, '0') + "</th>";
                  tDataString += "<td>0x" + System.cpu.ROM[i].toString(16).toUpperCase().padStart(4, '0') + "</td>\n";
                  tDataString += "</tr>\n";
                  tData.push(tDataString);
                  tDataString = "";
                }
                fs.writeFile("temp.txt", tData, (err) => {
                  if (err) console.log(err);
                    console.log("Successfully Written to File.");
                });
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
            console.log("Start/Resume");
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
          label: 'Dev Tools',
          accelerator: 'Shift+F12',
          click() {
            win.webContents.openDevTools();
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
          console.log("FreeMe!GB Settings");
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

function readFile(fileName) {
  const fs = require('fs');
  var buffer = fs.readFileSync(fileName);
  System.cpu.ROM = buffer;
}

electron.app.on('ready', createWindow);
