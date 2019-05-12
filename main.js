const electron = require('electron');

function createWindow () {
  var appmenu = electron.Menu.buildFromTemplate([
    {
      label: 'Menu',
      submenu: [
        {label:'Adjust Notification Value'},
        {label:'CoinMarketCap'},
        {label:'Exit'}
      ]
    }
  ])
  electron.Menu.setApplicationMenu(appmenu);
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

  // and load the index.html of the app.
  win.loadFile('main.html')
}

electron.app.on('ready', createWindow)

const cpu = require('./engine/cpu.js');

cpu.step();

// console.log('A: 0x%s', registers.AF.A.toString(16));
// console.log('F: 0x%s', registers.AF.F.toString(16));
// console.log('F: 0x%s', registers.F.toString(16));
// registers.AF = registers.A.toString(16) + registers.F.toString(16);
// console.log('AF: 0x%s', registers.AF.padStart(4, '0').toUpperCase());
