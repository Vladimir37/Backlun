'use strict';
const fetch = require('node-fetch');

const SERVER = "localhost";
const PORT = "8000";
const URL = `http://${SERVER}:${PORT}`

const task = {
  title: gen_token(4),
  text: gen_token(20),
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

function get_all_tasks() { //{{{
  fetch(`${URL}/api/tasks/get_all`)
    .then((res) => res.json())
    .then((json) => console.log('tasks: ', json.body))
    .catch(error => console.log('tasks error: ', error));
} //}}}

function add_new_task(task) { //{{{
  fetch(`${URL}/api/tasks/add`, {
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'POST',
      body: JSON.stringify(task),
    })
    .then((res) => res.json())
    .then((json) => {
      console.log('add task: ', json)
    })
    .catch(error => console.log('add task error: ', error));
} //}}}

export default function todo() {
  add_new_task(task);
  get_all_tasks();
}
