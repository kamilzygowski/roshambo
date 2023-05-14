import React, { useEffect, useRef, useState } from 'react';
import './App.css';
import { socket } from "./socket"
import dateFormat from "dateformat"
import { Oval } from 'react-loader-spinner';

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
  const [gameStatus, setGameStatus] = useState<string>("")
  const [winner, setWinner] = useState<string>("")
  const [scores, setScores] = useState<string>(" ")
  const buttonsRef = useRef<any>()
  const secondsRef = useRef<any>()
  const [seconds, setSeconds] = useState(1)
  const [timerActiv, setTimerActiv] = useState(false)
  const [isLoaderActive, setLoaderActive] = useState(false)
  enum OneOfThree {
    Paper,
    Stone,
    Scissors
  }
  useEffect(() => {
    chatWrapperRef.current?.scrollIntoView({ behavior: "smooth" })
  }, [chatMessages])

  const threeSecondsTimer = (winnerStr: string) => {
    //console.log(secondsRef.current.style)
    // secondsRef.current.style.color = "#f6e58d"
    setTimerActiv(true)
    setTimeout(() => {
      clearInterval(interval)
      setWinner(winnerStr)
      setTimerActiv(false)
      setSeconds(1)
    }, 2750)
    const interval = setInterval(() => {
      // secondsRef.current.style.color = "#ffbe76"

      setSeconds(oldVal => oldVal = oldVal + 1)
    }, 1000)
  }

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
        setRooms(resultArr)
      } else if (e.data[0] === "g") {
        setDuringMatch(true)
      } else if (e.data[0] === "w") {
        let dataAsString = e.data.toString().substring(1)
        threeSecondsTimer(dataAsString);
      } else if (e.data[0] === "s") {
        let dataAsString = e.data.toString().substring(1)
        setScores(dataAsString)
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
    if (draftingName === "" || draftingName === undefined) {
      return
    }
    setName(draftingName)
    socket.send("n" + draftingName)
  }
  const oneOfThreeHandler = (e: any) => {
    setOneOfThreeChoice(e.target.value)
    const arr = [].slice.call(buttonsRef.current.children);
    arr.forEach((element: HTMLElement) => {
      element.classList.remove("chosen")
      element.style.opacity = "0.6";
    })
    e.target.classList.add("chosen")
    e.target.style.opacity = 1;
  }

  return (
    <div className="App" onKeyDown={keyDownHandler}>
      {/*connected ? <p className='connectionStatus'>Connected</p> : <p className='connectionStatus'>Not connected</p>*/}
      {name === undefined ? <div className='setName'>
        <p className='paragraphName'>Provide your name</p>
        <input className='inputName' onChange={(elem) => setDraftingName(elem.target.value)} />
        <button className='acceptName' onClick={() => sendNameToSock()}>Accept</button></div> : null}
      {duringMatch ?
        <div className='game'>
          {timerActiv ? null : <p className='scores'>{scores}</p>}
          <div className='loadingDiv'>
            {timerActiv ? null : winner !== "" ? <p className='winner'>{winner}</p> : <p className='loading'>{gameStatus}</p>}
            {timerActiv ? <p className='seconds' ref={secondsRef}>{seconds}</p> : winner === "" ?
            isLoaderActive ?
                <Oval
                height={40}
                width={40}
                color="#f9ca24"
                visible={true}
                ariaLabel='oval-loading'
                secondaryColor="#f6e58d"
                strokeWidth={8}
                strokeWidthSecondary={8} /> : null
            : null }
          </div>
          <div ref={buttonsRef}>
            <input type="button" className='oneOfThree' value={OneOfThree.Paper} onClick={oneOfThreeHandler} />
            <input type="button" className='oneOfThree' value={OneOfThree.Stone} onClick={oneOfThreeHandler} />
            <input type="button" className='oneOfThree' value={OneOfThree.Scissors} onClick={oneOfThreeHandler} />
          </div>
          <input className='ready' value={"READY"} type="button" onClick={() => {
            socket.send("g" + oneOfThreeChoice)
            setLoaderActive(true)
            setGameStatus("Waiting for opponent")
            setWinner("");
          }} />
        </div>
        :
        <div className='main'>
          {rooms.map((element, index) => {
            if (element === "") {
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
