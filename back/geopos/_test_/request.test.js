'use strict';
const fetch = require('node-fetch');

const SERVER = "localhost";
const PORT = "8000";

const static_point = {
  token: gen_token(),
  coordinates: [gen_num(), gen_num()],
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

function get_points() { //{{{
  fetch(`http://${SERVER}:${PORT}/getPoints`)
    .then((res) => res.json())
    .then((json) => console.log('getPoint:', json.body))
    .catch(error => console.log('get error: ', error));
} //}}}

function post_point(point) { //{{{
  fetch(`http://${SERVER}:${PORT}/postPoint`, {
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'POST',
      body: JSON.stringify(point),
    })
    .then((res) => res.json())
    .then((json) => console.log('postPoint: ', json.body))
    .catch(error => console.log('postPoint error: ', error));
} //}}}

function get_rnd_point() { //{{{
  fetch(`http://${SERVER}:${PORT}/getRndPoint`)
    .then((res) => res.json())
    .then((json) => console.log('getRndPoint:', json.body))
    .catch(error => console.log('getRndPoint error: ', error));
} //}}}

function post_rnd_point() { //{{{
  let rnd_pos = {
    token: gen_token(),
    coordinates: [gen_num(), gen_num()],
  };

  fetch(`http://${SERVER}:${PORT}/postRndPoint`, {
      method: 'POST',
      body: JSON.stringify(rnd_pos),
    })
    .then((res) => res.json())
    .then((json) => console.log('postRndPoint: ', json.body))
    .catch(error => console.log('postRndPoint error: ', error));
} //}}}

function get_point_on_token(token) { //{{{
  fetch(`http://${SERVER}:${PORT}/getPointOnToken?token=${token}`)
    .then((res) => res.json())
    .then((json) => console.log('getPointOnToken:', json.body))
    .catch(error => console.log('getPointOnToken error: ', error));
} //}}}

function get_check_point() {
  fetch(`http://${SERVER}:${PORT}/getCheckPoint`)
    .then((res) => res.json())
    .then((json) => console.log('getCheckPoint:', json.body))
    .catch(error => console.log('getCheckPoint error: ', error));
}

function put_distance_point(point) {
  console.log('point: ', point)
  fetch(`http://${SERVER}:${PORT}/putDistance`, {
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'PUT',
      body: JSON.stringify(point),
    })
    .then((res) => res.json())
    .then((json) => console.log('Distance', json.body))
    .catch(error => console.log('putDistance error: ', error));
}

post_rnd_point();
post_rnd_point();
post_rnd_point();
get_points();
post_point(static_point);
get_point_on_token(static_point.token);

get_check_point()
put_distance_point(static_point);
