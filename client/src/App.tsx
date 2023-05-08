import React, { useCallback, useEffect, useState } from 'react';
import logo from './logo.svg';
import './App.css';
import { socket } from "./socket"

function App() {
  const [message, setMessage] = useState<string>("");
 useEffect(() => {
  socket.onopen = () => {
    console.log("Connected")
  }
  socket.onmessage = (e) => {
    console.log(e.data)
  }
   
 }, [])

const handleMessage = () => {
  if (message === ""){
    return
  }
  try {
    socket.send(message);
  } catch (error) {
    console.error(error)
  }
  
}
 
  return (
    <div className="App">
      <input type="text" onChange={(event) => setMessage(event.target.value)}></input>
      <button onClick={handleMessage}>Send message</button>
    </div>
  );
}

export default App;
