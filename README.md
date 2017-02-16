### backun - universal backend frontend

This project is important for any front-end developer, who whant to build a app with no details about backend programming.

### installation/dependencies
  * Nothing. Backun fully all-sufficient.

### tests
  Tests with fetch requests: examples/templates of requests to backend. Start: `npm install` inside `__test__` and:
  * `node request.test.js`
  * `babel-node request.test.js` (if it didn't work)

### What available now:

#### #blog
  Usable blog.
  
#### #forum
  Awesome forum.
  
#### #geopos
  Geoposition! Give a point on server and go.
  
#### #market
  Trade a little things.

#### #oauth2
  Oauth server with use google api. Start:
  * in console: `backun start oauth "port" "host" "key"`
  * key - your key file. Looks like: key.json, `{ "cid": "-", "csecret": "-"}`. That key have a default values (It may be out of date).
  
#### #chat
  Chat server with use websocket. Start:
  * in console: `backun start chat "port"`
  
#### #todo
  You don't know what you should do? Do todo!
  

  
