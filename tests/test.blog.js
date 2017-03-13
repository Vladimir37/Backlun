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

function get_all_posts() { //{{{
  fetch(`${URL}/api/get/all`)
    .then((res) => res.json())
    .then((json) => console.log('get posts: ', json.body))
    .catch(error => console.log('get posts error: ', error));
} //}}}

function create_post(auth) { //{{{
  const post = {
    token: auth.token,
    title: gen_token(8),
    tags: `${gen_token(4)}, ${gen_token(4)}`,
    text: gen_token(10),
  }
  fetch(`${URL}/api/posts/create`, {
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'POST',
      body: JSON.stringify(post),
    })
    .then((res) => res.json())
    .then((json) => {
      console.log('post a post: ', json)
    })
    .catch(error => console.log('post a post error: ', error));
} //}}}

function login(user) { //{{{
  return fetch(`${URL}/api/auth/login`, {
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'POST',
      body: JSON.stringify(user),
    })
    .then((res) => res.json())
    .then((json) => {
      console.log('login: ', json)
    })
    .catch(error => console.log('login error: ', error));
} //}}}

function get_auth_data() { //{{{
  fetch(`${URL}/api/auth/get`)
    .then((res) => res.json())
    .then((json) => console.log('get auth: ', json.body))
    .catch(error => console.log('get auth error: ', error));
} //}}}

function get_auth_and_login() { //{{{
  const auth = {
    login: '',
    password: '',
    token: ''
  };
  return fetch(`${URL}/api/auth/get`)
    .then((res) => res.json())
    .then((json) => {
      // console.log('auth data: ', json.body);
      auth.login = json.body.Login;
      auth.password = json.body.Password;
      return fetch(`${URL}/api/auth/login`, {
        headers: {
          'Content-Type': 'application/json'
        },
        method: 'POST',
        body: JSON.stringify(json.body),
      })
    })
    .then((res) => res.json())
    .then((json) => {
      auth.token = json.body;
      console.log('login: ', auth);
      return auth;
    })
    .catch(error => console.log('get auth error: ', error));
} //}}}

export default function blog() {
  get_auth_and_login().then(auth => {
    create_post(auth);
    get_all_posts();
  });
}
