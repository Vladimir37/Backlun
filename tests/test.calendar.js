'use strict';
const fetch = require('node-fetch');

const SERVER = "localhost";
const PORT = "8000";
const URL = `http://${SERVER}:${PORT}`

const default_category = {
  name: "first",
  color: "red"
};

const default_event = {
  category: 1,
  title: gen_token(4),
  description: gen_token(10),
  time: "2006-01-02T15:04:05.000Z"
};

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

function categories_get_all() { //{{{
  fetch(`${URL}/api/get/all_categories`)
    .then((res) => res.json())
    .then((json) => console.log('all cats: ', json.body))
    .catch(error => console.log('all cats error: ', error));
} //}}}

function categories_create(cat) { //{{{
  fetch(`${URL}/api/categories/create`, {
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'POST',
      body: JSON.stringify(cat),
    })
    .then((res) => res.json())
    .then((json) => console.log('create cat event: ', json))
    .catch(error => console.log('create cat error: ', error));
} //}}}

function events_get_all() { //{{{
  fetch(`${URL}/api/get/all`)
    .then((res) => res.json())
    .then((json) => console.log('all events: ', json.body))
    .catch(error => console.log('all events error: ', error));
} //}}}

function create_short_event(event) { //{{{
  fetch(`${URL}/api/short/create`, {
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'POST',
      body: JSON.stringify(event),
    })
    .then((res) => res.json())
    .then((json) => console.log('create short event: ', json))
    .catch(error => console.log('create short event error: ', error));
} //}}}

export default function calendar() {
  categories_create(default_category);
  categories_get_all();
  create_short_event(default_event);
  events_get_all();
}

categories_create(default_category);
categories_get_all();
create_short_event(default_event);
events_get_all();
