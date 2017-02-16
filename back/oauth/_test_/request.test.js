'use strict';
const fetch = require('node-fetch');

const SERVER = "localhost";
const PORT = "8000";
const URL = `http://${SERVER}:${PORT}`

const static_auth = {
  state: gen_token(),
  code: gen_token()
};

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

function get_login() { //{{{
  fetch(`${URL}/api/login`)
    .then((res) => res.json())
    .then((json) => console.log('login: ', json))
    .catch(error => console.log('login error: ', error));
} //}}}

function get_auth(auth) { //{{{
  fetch(`${URL}/api/auth?state=${auth.state}`)
    .then((res) => res.json())
    .then((json) => console.log('auth:', json))
    .catch(error => console.log('auth error: ', error));
} //}}}

get_login();
