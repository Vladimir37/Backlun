'use strict';
const fetch = require('node-fetch');

const SERVER = "localhost";
const PORT = "8000";
const URL = `http://${SERVER}:${PORT}`

const default_game = {
  players: 2,
  decks: 0,
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

function game_get_all() { //{{{
  fetch(`${URL}/api/get/all`)
    .then((res) => res.json())
    .then((json) => console.log('all game: ', json.body))
    .catch(error => console.log('all game error: ', error));
} //}}}

function game_start(game) { //{{{
  fetch(`${URL}/api/game/start`, {
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'POST',
      body: JSON.stringify(game),
    })
    .then((res) => res.json())
    .then((json) => console.log('game start: ', json.body))
    .catch(error => console.log('game start error: ', error));
} //}}}

export default function blackjack() {
  game_start(default_game);
  game_get_all();
}

game_start(default_game);
game_get_all();
