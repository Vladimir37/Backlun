'use strict';
const fetch = require('node-fetch');

const SERVER = "localhost";
const PORT = "8000";
const URL = `http://${SERVER}:${PORT}`

const user = {
  login: gen_token(4),
  password: gen_token(8),
};

const thread = {
  token: '',
  category: 1,
  title: 'yes!',
  text: gen_token(25),
};

//gen data {{{

function obj_to_param(obj) {
  return Object.keys(obj).map(function(key) {
    return key + '=' + obj[key];
  }).join('&');
}

function gen_num() {
  const min = 1;
  const max = 10000;
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

function get_all_threads(category) { //{{{
  fetch(`${URL}/api/get/threads?${obj_to_param(category)}`)
    .then((res) => res.json())
    .then((json) => console.log('threads: ', json.body))
    .catch(error => console.log('threads error: ', error));
} //}}}

function get_all_categories() { //{{{
  fetch(`${URL}/api/get/categories`)
    .then((res) => res.json())
    .then((json) => console.log('categories: ', json.body))
    .catch(error => console.log('threads error: ', error));
} //}}}

function create_thread(thread) { //{{{
  console.log('thread will be created: ', thread);
  fetch(`${URL}/api/forum/create`, {
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'POST',
      body: JSON.stringify(thread),
    })
    .then((res) => res.json())
    .then((json) => {
      console.log('thread: ', json)
    })
    .catch(error => console.log('post a post error: ', error));
} //}}}

function registration(user) { //{{{
  return fetch(`${URL}/api/auth/registration`, {
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'POST',
      body: JSON.stringify(user),
    })
    .then((res) => res.json())
    .then((json) => {
      console.log(json.body);
      user.token = json.body.Token;
      return user;
    })
    .catch(error => console.log('get auth error: ', error));
} //}}}

// get_all_categories();
export default function forum () {
  registration(user).then(user => {
    thread.token = user.token;
    create_thread(thread);
    get_all_threads({
      category: 3
    });
    // console.log(user);
  });
}

forum();
