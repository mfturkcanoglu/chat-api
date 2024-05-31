import { useEffect, useState } from "react";
import "./App.css";
import { connect, sendMsg, onMsg } from "./api";
import Message from "./api/models/Message";
import Header from "./components/Header";

function App() {
  const [currentMessage, setCurrentMessage] = useState("");
  const [messages, setMessages] = useState<Message[]>([]);
  useEffect(() => {
    connect((msg) => {
      setMessages([msg, ...messages]);
    });
  }, []);

  useEffect(() => {
    console.log("message", messages);
  }, [messages]);

  const sendMessage = () => {
    let msg: Message = {
      message: currentMessage,
    };

    sendMsg(msg);
    setCurrentMessage("");
  };

  return (
    <div className="App">
      <Header />
      <div>
        <input
          value={currentMessage}
          onChange={({ target }) => setCurrentMessage(target.value)}
        />
        <button onClick={sendMessage}>Send Message</button>
      </div>
    </div>
  );
}

export default App;
