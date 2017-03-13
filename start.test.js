// 'use strict';
import * as reqs from './tests'

const spawn = require('child_process').spawn;

console.log(reqs);

function runcmd(cmd, args, func) {
  return new Promise(function(resolve, reject) {
    const child = spawn(cmd, args);
    let resp = "";
    func();
    setTimeout(() => {
      child.kill();
      resolve(args[1]);
    }, 3000);
  });
}

function start_one(backend) {
  console.log('platforma: ', backend);
  runcmd("./Backlun", ["start", backend], reqs[backend]).then((res) => {
    console.log('complete: ', res);
  });
}

function start_all() {
  Object.keys(reqs).forEach((backend, i) => {
    setTimeout(()=>{start_one(backend)}, 4000*i);
  });
}


// start_one("geopos");
start_all();
