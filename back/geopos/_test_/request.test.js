'use strict';
const fetch = require('node-fetch');

const SERVER = "localhost";
const PORT = "8000";
const URL = `http://${SERVER}:${PORT}`

const static_point = {
  Type: 'Point',
  Token: gen_token(),
  Coordinates: [gen_num(), gen_num()],
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
  fetch(`${URL}/api/points/get`)
    .then((res) => res.json())
    .then((json) => console.log('get points: ', json.body))
    .catch(error => console.log('get points error: ', error));
} //}}}

function get_point_from_token(token) { //{{{
  fetch(`${URL}/api/points/get_from_token?token=${token}`)
    .then((res) => res.json())
    .then((json) => console.log('get_from_token:', json))
    .catch(error => console.log('get_from_token error: ', error));
} //}}}

function post_point(point, sample) {//{{{
  fetch(`${URL}/api/points/post`, {
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'POST',
      body: JSON.stringify(point),
    })
    .then((res) => res.json())
    .then((json) => {
      if (Object.keys(point).length == 0) {
        console.log('here is null');
        console.assert(json.body == null);
      }
      console.log('post point: ', json)
    })
    .catch(error => console.log('post point error: ', error));
}//}}}

function get_rnd_point() { //{{{
  fetch(`${URL}/api/random_point/get`)
    .then((res) => res.json())
    .then((json) => console.log('get random_point:', json))
    .catch(error => console.log('get random_point error: ', error));
} //}}}

function post_rnd_point() {//{{{
  let rnd_pos = {
    Type: 'Point',
    Token: gen_token(),
    Coordinates: [gen_num(), gen_num()],
  };

  fetch(`${URL}/api/random_point/post`, {
      method: 'POST',
      body: JSON.stringify(rnd_pos),
    })
    .then((res) => res.json())
    .then((json) => console.log('post random_point: ', json))
    .catch(error => console.log('post random_point error: ', error));
}//}}}

function get_check_point() { //{{{
  fetch(`${URL}/api/check_point/get`)
    .then((res) => res.json())
    .then((json) => console.log('getCheckPoint:', json.body))
    .catch(error => console.log('getCheckPoint error: ', error));
} //}}}

function post_check_point(point) { //{{{
  fetch(`${URL}/api/check_point/post`, {
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'POST',
      body: JSON.stringify(point),
    })
    .then((res) => res.json())
    .then((json) => console.log('post check_point: ', json))
    .catch(error => console.log('post check_point error: ', error));
} //}}}

function put_distance_point(point) { //{{{
  console.log('distance point: ', point)
  fetch(`${URL}/api/check_point/put_distance`, {
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'PUT',
      body: JSON.stringify(point),
    })
    .then((res) => res.json())
    .then((json) => console.log('Distance', json))
    .catch(error => console.log('putDistance error: ', error));
} //}}}

function get_full_response() { //{{{
  fetch(`${URL}/api/random_point/get`)
    .then((res) => res.json())
    .then((json) => console.log('get response: ', json))
    .catch(error => console.log('get response error: ', error));
} //}}}

get_full_response();
post_point({});
post_point({test: "test"});
post_point(static_point);
post_rnd_point();
get_points();
