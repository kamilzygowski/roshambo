import React, { useEffect, useRef, useState } from 'react';
import './App.css';
import { socket } from "./socket"
import dateFormat from "dateformat"

function App() {
  const [connected, setConnected] = useState<boolean>(false);
  const [chatMessages, setChatMessages] = useState<string[]>(["#### GLOBAL CHAT ####"])
  const chatInputRef = useRef<any>()
  const chatWrapperRef = useRef<any>();

  useEffect(() => {
    chatWrapperRef.current?.scrollIntoView({ behavior: "smooth" })
  }, [chatMessages])

  useEffect(() => {
    socket.onopen = () => {
      setConnected(true);
    }
    socket.onmessage = (e) => {
      if (e.data[0] === 'm'){
        const datePrefix = new Date()
        let dataAsString = e.data.toString().substring(1)
        setChatMessages(oldArray => [...oldArray, dateFormat(datePrefix, "HH:MM") + dataAsString])
      }
    }
    socket.onclose = () => {
      setConnected(false);
    }
  }, [])

  const handleMessage = () => {
    if (chatInputRef.current.value === "") {
      return
    }
    socket.send("m [Name]: " + chatInputRef.current.value)
    chatInputRef.current.value = ""
  }

  const handleRoomEnter = () => {
    socket.send("rXD")
  }

  const keyDownHandler = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.code === "Enter"){
      handleMessage()
    }
  } 

  return (
    <div className="App" onKeyDown={keyDownHandler}>
      {connected ? <p className='connectionStatus'>Connected</p> : <p className='connectionStatus'>Not connected</p>}
      <div className='main'>
        <button className='startButton' onClick={handleRoomEnter}>Start</button>
      </div>
      <div className='chat'>
        {chatMessages.map((element, index) => {
          return <p key={index}>{element}</p>;
        })}
        <div ref={chatWrapperRef}></div>
      </div>
      <div className='wrapper'>
        <input type="text" ref={chatInputRef} className="chatInput"></input>
        <button onClick={handleMessage} className="chatButton">Send</button>
      </div>
    </div>
  );
}

export default App;
