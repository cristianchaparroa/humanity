// api/index.js


let connect = (socket,callback) => {
  console.log("Attempting connection");

  socket.onopen = () => {
    console.log("Successfukky connected");
  };

  socket.onmessage = msg => {
    console.log("onmessage")
    console.log(msg);
    callback(msg);
  }

  socket.onclose = event => {
    console.log("socket closed connection: ", event);
  };

  socket.onerror = error => {
    console.log("Socket error: ", error);
  };
}

let sendMessage =(socket,msg) => {
  console.log("sending message: ", msg);
  socket.send(msg);
}



export { connect, sendMessage}
