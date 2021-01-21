const electron = require('electron');
const { app, BrowserWindow } = require('electron')
const ipcMain = require('electron').ipcMain



function loginWindow() {
    let win = new BrowserWindow({
        width: 400,
        height: 600,
        maxWidth: 400,
        maxHeight: 600,
        webPreferences: {
            nodeIntegration: true
        }
    });
    win.loadFile('login.html');
    ipcMain.on('loginSuccess', function (event, callerIP) {
        createWindow();
        win.close();
    });
    ipcMain.on('startSignUp', function (event, request) {
        win.loadFile('signup.html');
    });
    ipcMain.on('back', function (event, request) {
        win.loadFile('login.html');
    });
}

function createWindow() {
    let win = new BrowserWindow({
        width: 800,
        height: 600,
        maxWidth: 1600,
        maxHeight: 900,
        webPreferences: {
            nodeIntegration: true
        }
    })
    win.loadFile('index.html');
    win.on('close', () => {
        app.quit();
    });
    ipcMain.on('incomingcall', function (event, callerIP) {
        event.sender.send('playback', "success");
        win.webContents.send('incomingDivert', callerIP);
    });
    backgroundWindow();
}

function backgroundWindow() {
    let backgroundWindow = new BrowserWindow({
        show: false,
        webPreferences: {
            nodeIntegration: true
        }
    });
    backgroundWindow.loadFile("background.html");
}


app.setAppUserModelId("com.VCApp")

app.whenReady().then(loginWindow)

app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') {
        app.quit()
    }
})

app.on('activate', () => {
    if (BrowserWindow.getAllWindows().length === 0) {
        createWindow();
        // loginWindow();
    }
})

app.on('ready', async function () {
    //pass
});