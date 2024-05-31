interface ChatHistoryProps {
  chatHistory: any[];
}

const ChatHistory = (props: ChatHistoryProps) => {
  const messages = props.chatHistory.map((msg, index) => (
    <p key={index}>{msg.data}</p>
  ));
  return (
    <div className="ChatHistory">
      <h2>Chat History</h2>
      {messages}
    </div>
  );
};

export default ChatHistory;
