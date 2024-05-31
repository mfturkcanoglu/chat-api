import Message from "./models/Message";

var socket = new WebSocket("ws://localhost:8080/ws");

let connect = (callback: (n: Message) => any) => {
  console.log("Attempting connecting...");

  socket.onopen = () => {
    console.log("Successfuly connected");
  };

  socket.onmessage = (msg) => {
    console.log(msg);
    callback(msg.data);
  };

  socket.onclose = (event) => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = (error) => {
    console.log("Socket Error: ", error);
  };
};

let sendMsg = (msg: Message) => {
  console.log("sending msg: ", msg);
  socket.send(JSON.stringify(msg));
};

let onMsg = () => {
  return "";
};

export { connect, sendMsg, onMsg };
