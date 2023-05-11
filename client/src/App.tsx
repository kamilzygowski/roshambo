import React, { useEffect, useRef, useState } from 'react';
import './App.css';
import { socket } from "./socket"
import dateFormat from "dateformat"

function App() {
  const [connected, setConnected] = useState<boolean>(false);
  const [rooms, setRooms] = useState<string[]>([])
  const [name, setName] = useState<string>()
  const [draftingName, setDraftingName] = useState<string>()
  const [chatMessages, setChatMessages] = useState<string[]>(["#### GLOBAL CHAT ####"])
  const chatInputRef = useRef<any>()
  const chatWrapperRef = useRef<any>();
  const [duringMatch, setDuringMatch] = useState<boolean>(false)
  const [oneOfThreeChoice, setOneOfThreeChoice] = useState<number>()
  enum OneOfThree {
    Paper,
    Stone,
    Scissors
  }
  useEffect(() => {
    chatWrapperRef.current?.scrollIntoView({ behavior: "smooth" })
  }, [chatMessages])

  useEffect(() => {
    socket.onopen = () => {
      setConnected(true);
    }
    socket.onmessage = (e) => {
      if (e.data[0] === 'm') {
        const datePrefix = new Date()
        let dataAsString = e.data.toString().substring(1)
        setChatMessages(oldArray => [...oldArray, dateFormat(datePrefix, "HH:MM") + dataAsString])
      } else if (e.data[0] === "r") {
        const charsToSplit = "###"
        const dataAsString = e.data.toString().substring(charsToSplit.length + 1)
        const resultArr = dataAsString.split(charsToSplit)
        console.log(dataAsString)
        setRooms(resultArr)
      }else if (e.data[0] === "g"){
        setDuringMatch(true)
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
    socket.send(`m [${name}]: ` + chatInputRef.current.value)
    chatInputRef.current.value = ""
  }

  const handleRoomEnter = (e: any) => {
    socket.send("r" + e.target.value)
  }

  const keyDownHandler = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.code === "Enter") {
      handleMessage()
    }
  }
  const sendNameToSock = () => {
    setName(draftingName)
    socket.send("n" + draftingName)
  }

  return (
    <div className="App" onKeyDown={keyDownHandler}>
      {connected ? <p className='connectionStatus'>Connected</p> : <p className='connectionStatus'>Not connected</p>}
      {name === undefined ? <div className='setName'><input className='inputName' onChange={(elem) => setDraftingName(elem.target.value)} />
        <button className='acceptName' onClick={(e) => sendNameToSock()}>Accept</button></div> : null}
        {duringMatch ? 
        <div className='game'>
          <input type="button" className='oneOfThree' value={OneOfThree.Paper} onClick={(e: any) => setOneOfThreeChoice(e.target.value)}/>
          <input type="button" className='oneOfThree' value={OneOfThree.Stone} onClick={(e: any) => setOneOfThreeChoice(e.target.value)}/>
          <input type="button" className='oneOfThree' value={OneOfThree.Scissors} onClick={(e: any) => setOneOfThreeChoice(e.target.value)}/>
          </div>
      :
      <div className='main'>
      {rooms.map((element, index) => {
        if (element === ""){
          return null;
        }
        return <input className='room' type="button" value={element} key={index} onClick={handleRoomEnter}></input>
      })}
    </div>
      }

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
