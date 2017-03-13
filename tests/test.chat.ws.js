'use strict'
const wsClient = require('websocket').client;
// const ws_client = (require('websocket').client)();
const fetch = require('node-fetch');

const SERVER = "localhost";
const PORT = "8000";
const URL = `http://${SERVER}:${PORT}`

const ws_client = new wsClient();

//gen data {{{
function gen_num() {
  const min = 0.0000;
  const max = 1.9000;
  return Math.random() * (max - min) + min;
}

function gen_token() {
  const possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
  const len_token = 5;
  let text = "";
  for (let i = 0; i < len_token; i++) {
    text += possible.charAt(Math.floor(Math.random() * possible.length));
  };
  return text;
} //}}}

ws_client.on('connectFailed', function(error) { //{{{
  console.log('Connect Error: ' + error.toString());
}); //}}}

ws_client.on('connect', function(connection) { //{{{
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

  function send_msg() {
    if (connection.connected) {
      const msg = {
        "message": gen_token()
      };
      connection.sendUTF(JSON.stringify(msg));
    }
  }

  function start_sending(num) {
    for (let i = 1; i < num; i++) {
      setTimeout(send_msg, 3000 * i);
    }
  }
  start_sending(10);
}); //}}}

function get_error_request() { //{{{
  fetch(`${URL}/api/error_request`)
    .then((res) => res.json())
    .then((json) => console.log('error_request:', json))
    .catch(error => console.log('error_request error: ', error));
} //}}}
export default function chat () {
  get_error_request();
  ws_client.connect(`ws://${SERVER}:${PORT}/api/ws`);
}
