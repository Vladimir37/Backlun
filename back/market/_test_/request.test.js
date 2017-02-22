'use strict';
const fetch = require('node-fetch');

const SERVER = "localhost";
const PORT = "8000";
const URL = `http://${SERVER}:${PORT}`

const user = {
  login: gen_token(4),
  password: gen_token(8),
  FullName: gen_token(10),
  address: gen_token(20),
};

const credit = {
  token: '',
  credit: 1000,
};

//gen data {{{
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

function get_all_products() { //{{{
  fetch(`${URL}/api/get/products`)
    .then((res) => res.json())
    .then((json) => console.log('products: ', json.body))
    .catch(error => console.log('products error: ', error));
} //}}}

function add_credits(credit) {//{{{
  console.log('thread will be created: ', credit);
  fetch(`${URL}/api/market/credits`, {
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'POST',
      body: JSON.stringify(credit),
    })
    .then((res) => res.json())
    .then((json) => {
      console.log('add credit: ', json)
    })
    .catch(error => console.log('add credit error: ', error));
}//}}}

function registration(user) {//{{{
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
}//}}}

// get_all_categories();
registration(user).then(user => {
  credit.token = user.token;
  add_credits(credit);
  get_all_products();
  // console.log(user);
});

