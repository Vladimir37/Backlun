### backun - universal backend frontend

### installation/dependencies
  * Nothing. Backun fully all-sufficient.

### tests
  fetch-requests. Start: `npm install` inside `__test__` and:
  * `node request.test.js`
  * `babel-node request.test.js` (if it didn't work)

### What available now:

#### #blog
  Usable blog.
  
#### #forum
  Awesome forum.
  
#### #geopos
  Geoposition. Give a point to server and go.
  
#### #market
  Trade market.

#### #oauth2
  Oauth2 server with use google api. Start:
  * console: `backun start oauth "port" "host" "key"`
  * key - key file. name: key.json, `{ "cid": "-", "csecret": "-"}`. Key have a default value.
  
#### #chat
  websocket chat. Start:
  * console: `backun start chat "port"`
  
#### #todo
  Easy todo
