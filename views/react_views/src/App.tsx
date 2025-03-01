import { useState } from "react";
import ChatHistory from "./components/ChatHistory";
import ChatForm from "./components/ChatForm";
import "./App.css";

export default function App() {
  const [messages, setMessages] = useState<string[]>([]);
  const [refresh, setRefresh] = useState(false);

  const handleNewMessage = (message: string) => {
    setMessages((prev) => [...prev, message]);
    setRefresh(!refresh);
  };

  return (
    <div className="container">
      <div className="chat_history">
        <div className="header">PublicChatter</div>
        <ChatHistory refresh={refresh} setRefresh={setRefresh}/>
      </div>
      <ChatForm onNewMessage={handleNewMessage} />
    </div>
  );
}