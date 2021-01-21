function initTURNconfig() {
    var apigClient = apigClientFactory.newClient({
        apiKey: 'API_KEY_FOR_AWS_LAMBDA',
        region: 'LAMBDA_REGION'
    });
    var body = {
        "ques": "value"
    }
    apigClient.turnconfigPost({}, body, {})
        .then(function (result) {
            // console.log("making post calls to api gateway result = ", result);
            temp = JSON.parse(result.data.body);
            var config = require('./turnConfig.json');
            config.username = temp.username;
            config.password = temp.password;
            config.ip = temp.ip;
            config.port = temp.port;
            const fs = require('fs');
            fs.writeFileSync("turnConfig.json", JSON.stringify(config));
        }).catch(function (result) {
            console.log(result);
        });
}
initTURNconfig();
checkIncomingCalls();
function checkIncomingCalls() {
    setTimeout(function () {

        var apigClient = apigClientFactory.newClient({
            apiKey: 'API_KEY_FOR_AWS_LAMBDA',
            region: 'LAMBDA_REGION'
        });
        var body = {
            "callee": require('./session.json').username
        }
        apigClient.checkincomingPost({}, body, {})
            .then(function (result) {
                // console.log("making post calls to api gateway result = ", result);
                data = (JSON.parse(result.data.body));
                console.log(data)
                if (data.message != "NoIncomingCall") {
                    connectToChannel(data["callerIP"]);
                    console.log(data.callerIP);
                } else {
                    checkIncomingCalls();
                }
            }).catch(function (result) {
                console.log(result);
            });
        // fetch("http://LAPTOP-JQN82GT4:8080/isBeingCalled/indiasummers").then(
        //     function (response) {
        //         return response.text();
        //     }).then(function (data) {
        //         data = (JSON.parse(data));
        //         if (data["status"] == "success") {
        //             connectToChannel(data["callerIP"]);
        //         } else {
        //             checkIncomingCalls();
        //         }
        //     }).catch(function (err) {
        //         console.log(err);
        //     });
    }, 5000);
}
async function connectToChannel(callerIP) {
    const ipcRenderer = require('electron').ipcRenderer
    ipcRenderer.send('incomingcall', callerIP);
    // checkIncomingCalls();
}