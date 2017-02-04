'use strict'
const WebSocketClient = require('websocket').client;

let first_client = new WebSocketClient();

function generateRandomNumber() {
    const   min = 0.0000,
            max = 1.9000;
    return Math.random() * (max - min) + min;
};

first_client.on('connectFailed', function(error) {
    console.log('Connect Error: ' + error.toString());
});

first_client.on('connect', function(connection) {
    console.log('WebSocket Client Connected');
    connection.on('error', function(error) {
        console.log("Connection Error: " + error.toString());
    });

    connection.on('close', function() {
        console.log('echo-protocol Connection Closed');
    });

    connection.on('message', function(message) {
        if (message.type === 'utf8') {
            console.log("Received: '" + message.utf8Data + "'");
        }
    });
  
    function sendGeoPos() {
        if (connection.connected) {
            const geo_pos = {
                "qW43GFvDghGG" : {
                "lat"   : generateRandomNumber(),
                "lng"   : generateRandomNumber()}
            };
            //first_client.send(JSON.stringify(geo_pos));
            connection.sendUTF(JSON.stringify(geo_pos));
            setTimeout(sendGeoPos, 10000);
        }
    }
    sendGeoPos();
});

first_client.connect('ws://localhost:8080/ws');
