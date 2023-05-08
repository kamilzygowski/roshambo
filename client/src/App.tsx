import React, { useCallback, useEffect, useState } from 'react';
import logo from './logo.svg';
import './App.css';
import { socket } from "./socket"

function App() {
  //const [socket, setSocket] = useState<any>();
 useEffect(() => {
   /* const sock = io("localhost:8000/socket.io", {
      withCredentials: true,
      //transports: ["websocket"]
    })*/
    //setSocket(sock)
    console.log(socket)
    socket.on("connect", () => {
      console.log("Connected!")
    })
    socket.on("connecting", () => {
      console.log("Trying to connect")
    })
    socket.on("render", () => {
      console.log("Render")
    })
    return () => {
      socket.close()
    }
   
 }, [])
 
  return (
    <div className="App">

    </div>
  );
}

export default App;
