'use strict';
const fetch = require('node-fetch');

const SERVER = "localhost";
const PORT = "8000";
const URL = `http://${SERVER}:${PORT}`


//gen data {{{
function gen_num() {
  const min = 0.0000;
  const max = 1.9000;
  return Math.random() * (max - min) + min;
}

function gen_token(len_token) {
  const possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
  let text = "";
  for (let i = 0; i < len_token; i++) {
    text += possible.charAt(Math.floor(Math.random() * possible.length));
  };
  return text;
} //}}}

export default function calendar () {
}
