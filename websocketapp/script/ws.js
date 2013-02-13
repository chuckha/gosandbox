var sock = new WebSocket("ws://localhost:8080/echo");
console.log(sock);

sock.onopen = function (event) {
  sock.send("Hello websocket!");
}

sock.onmessage = function (event) {
  document.write(event.data);
}
