// api/index.js

var socket = new WebSocket("ws://localhost:8080/ws/room");

let connect = callback => {
  console.log("Attempting connection");

  socket.onopen = () => {
    console.log("Successfukky connected");
  };

  socket.onmessage = msg => {
    console.log(msg);
    console.log(callback);
    callback(msg);
  }

  socket.onclose = event => {
    console.log("socket closed connection: ", event);
  };

  socket.onerror = error => {
    console.log("Socket error: ", error);
  };
}

let sendMessage = msg => {
  console.log("sending message: ", msg);
  socket.send(msg);
}

export { connect, sendMessage}
