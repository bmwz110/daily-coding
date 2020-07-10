let http = require("http");
let url = require('url');

function start(route, handle) {
  http.createServer((req, res) => {
    let pathname = url.parse(req.url).pathname;
    console.log("Request for " + pathname + " received.");

    route(handle, pathname);

    res.writeHead(200, {"Content-Type": "text/plain"});
    res.write("Hello World");
    res.end();
  }).listen(8888);
  console.log("Server has started.");
}

exports.start = start;
 
