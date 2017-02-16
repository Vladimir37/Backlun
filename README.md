### backun - univeral static backend for your frontend project!

This project is important for any front-end developer, who whant to build a app with no details about backend programming.

### dependencies
  * Nothing. This project fully all-sufficient.

### installation
  * Nothing. Just start your backlun and use it.

### tests
  We make in free form javascript tests with fetch requests. You have examples for all request to backend, just copy-paste it in you application. If you want to start it, do command: `npm install` inside folder. Start `test` in `__test__`:
  * `node request.test.js`
  * `babel-node request.test.js` (if it didn't work)
  
### What backend functions available now:

#### #blog
  Usable blog
  
#### #forum
  Awesome forum
  
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
  
