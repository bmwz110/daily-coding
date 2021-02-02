function route(handle, pathname, res, req) {
  if (typeof handle[pathname] === 'function') {
    handle[pathname](res, req);
  } else {
    res.writeHead(404, {"Content-Type": "text/plain"});
    res.write("404 Not Found");
    res.end();
  }
}

exports.route = route;
 
