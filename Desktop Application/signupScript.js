function signup() {
    var apigClient = apigClientFactory.newClient({
        apiKey: 'API_KEY_FOR_AWS_LAMBDA',
        region: 'LAMBDA_REGION'
    });
    var username = document.getElementById('username').value;
    var password = document.getElementById('password').value;
    var retypedPassword = document.getElementById('retypePassword').value;
    var body = {
        "username": username,
        "password": password
    }
    if (password == retypedPassword && username.length <= 16) {
        apigClient.signupPost({}, body, {})
            .then(function (result) {
                response = JSON.parse(result.data.body);
                if (response.result == "success") {
                    userConfig = require('./session.json');
                    userConfig.username = document.getElementById('username').value;
                    const fs = require('fs');
                    fs.writeFileSync("session.json", JSON.stringify(userConfig));
                    const ipcRenderer = require('electron').ipcRenderer
                    ipcRenderer.send('loginSuccess', "success");
                } else {
                    document.getElementById('username').value = "";
                    document.getElementById('password').value = "";
                    document.getElementById('retypePassword').value = "";
                    document.getElementById('username').style.borderColor = "red";
                    document.getElementById('password').style.borderColor = "red";
                    document.getElementById('retypepassword').style.borderColor = "red";

                }
            }).catch(function (result) {
                console.log(result);
            });
    }
    else {
        document.getElementById('username').value = "";
        document.getElementById('password').value = "";
        document.getElementById('retypePassword').value = "";
        document.getElementById('username').style.borderColor = "red";
        document.getElementById('password').style.borderColor = "red";
        document.getElementById('retypePassword').style.borderColor = "red";
    }

}
function back() {
    const ipcRenderer = require('electron').ipcRenderer
    ipcRenderer.send('back', "login");
}