import React, { useCallback, useEffect, useState } from 'react';
import logo from './logo.svg';
import './App.css';
import { io } from "socket.io-client"

function App() {
  const [socket, setSocket] = useState();
 useEffect(() => {
    const socket = io("localhost:8000/socket.io", {
      withCredentials: true,
      //transports: ["websocket"]
    })
    socket.on("connect", () => {
      console.log("Connected!")
    })
    socket.on("render", () => {
      console.log("Render")
    })
 }, [])
  
  return (
    <div className="App">

    </div>
  );
}

export default App;
