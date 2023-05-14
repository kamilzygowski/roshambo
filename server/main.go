package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients = []websocket.Conn{}
var allPlayers = []player{}
var allGames = []Games{}
var allNames string
var incrementingId uint16 = 0

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is online"))
}

func removeFromSlice(s []player, index int) []player {
	return append(s[:index], s[index+1:]...)
}

func sendRoomsToAllClients() {
	// Reset allNames string
	allNames = ""
	for _, p := range allPlayers {
		allNames += "###" + p.name
	}
	// SEND TO ALL - rooms info
	for i := 0; i < len(clients); i++ {
		clients[i].WriteMessage(1, []byte("r"+allNames))
	}
}

func socketReader(conn *websocket.Conn) {
	for {
		for i := range allGames {
			if allGames[i].players[0].isReady == true && allGames[i].players[1].isReady == true {
				var winner string
				if (allGames[i].players[0].choice == 0 && allGames[i].players[1].choice == 1) || (allGames[i].players[0].choice == 1 && allGames[i].players[1].choice == 2) || (allGames[i].players[0].choice == 2 && allGames[i].players[1].choice == 0) {
					winner = allGames[i].players[0].name + " WON"
					allGames[i].scores[0] = allGames[i].scores[0] + 1
				} else if (allGames[i].players[1].choice == 0 && allGames[i].players[0].choice == 1) || (allGames[i].players[1].choice == 1 && allGames[i].players[0].choice == 2) || (allGames[i].players[1].choice == 2 && allGames[i].players[0].choice == 0) {
					winner = allGames[i].players[1].name + " WON"
					allGames[i].scores[1] = allGames[i].scores[1] + 1
				} else {
					winner = "DRAW"
				}
				for index := range allGames[i].players {
					allGames[i].players[index].isReady = false
					allGames[i].players[index].conn.WriteMessage(1, []byte("w"+winner))
					allGames[i].players[index].conn.WriteMessage(1, []byte("s"+allGames[i].players[0].name+" "+strconv.Itoa(int(allGames[i].scores[0]))+" : "+strconv.Itoa(int(allGames[i].scores[1]))+" "+allGames[i].players[1].name))
				}
			}
		}

		msgType, msg, err := (*conn).ReadMessage()
		if err != nil {
			// Handling error / disconnect
			fmt.Printf("User %s has disconnected\n", (*conn).RemoteAddr())
			// Remove this player from allPlayers slice
			for i := range allPlayers {
				if (allPlayers[i]).remoteAddress == (*conn).RemoteAddr().String() {
					allPlayers = removeFromSlice(allPlayers, i)
				}
			}
			// Removing client
			for index, client := range clients {
				if client.RemoteAddr() == (*conn).RemoteAddr() {
					clients = append(clients[:index], clients[index+1:]...)
				}
			}
			sendRoomsToAllClients()
			return
		}

		fmt.Printf("%s send: %s\n", (*conn).RemoteAddr(), string(msg))
		fmt.Println("Number of clients: ", len(clients))
		for _, client := range clients {
			if msg[0] == []byte("m")[0] {
				// Chat message
				client.WriteMessage(msgType, msg) // Populate the message to other clients
			} else if msg[0] == []byte("n")[0] {
				// Name
				msgContent := msg[1:]
				for i := range allPlayers {
					if (allPlayers[i]).remoteAddress == (*conn).RemoteAddr().String() {
						(allPlayers[i]).name = string(msgContent)
					}
				}
				sendRoomsToAllClients()

			} else if msg[0] == []byte("r")[0] {
				// Room
				msgContent := msg[1:]
				for i := range allPlayers {
					if (allPlayers[i]).name == string(msgContent) {
						// IF PLAYER DIDNT CLICK ON HIS ROOM
						if client.RemoteAddr().String() != allPlayers[i].remoteAddress {
							if client.RemoteAddr().String() == (*conn).RemoteAddr().String() {
								var playerFromClient *player
								for i := range allPlayers {
									if allPlayers[i].remoteAddress == client.RemoteAddr().String() {
										playerFromClient = &allPlayers[i]
									}
								}
								allGames = append(allGames, Games{players: []player{*playerFromClient, allPlayers[i]}, isDone: false, scores: []uint16{0, 0}})
								(*allPlayers[i].conn).WriteMessage(1, []byte("g"))
								(*conn).WriteMessage(1, []byte("g"))
							}
						}
					}
				}
			} else if msg[0] == []byte("g")[0] {
				// Player is ready
				for i := 0; i < len(allPlayers); i++ {
					msgContent := msg[1:]
					if (allPlayers[i]).remoteAddress == (*conn).RemoteAddr().String() {
						msgInt, err := strconv.Atoi(string(msgContent))
						if err != nil {
							fmt.Println("ERROR [Msg string to int conversion] ERROR")
						}
						allPlayers[i].setReady(uint8(msgInt))
						for i := range allGames {
							if allGames[i].players[0].remoteAddress == (*conn).RemoteAddr().String() {
								allGames[i].players[0].setReady(uint8(msgInt))
							}
							if allGames[i].players[1].remoteAddress == (*conn).RemoteAddr().String() {
								allGames[i].players[1].setReady(uint8(msgInt))
							}
						}
					}
				}
			}
		}
	}
}

func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // Very important
	conn, _ := upgrader.Upgrade(w, r, nil)

	clients = append(clients, *conn)
	//(*conn).WriteMessage(1, []byte("r"+allNames))
	// Add and init player
	incrementingId++
	newPlayer := player{id: incrementingId, name: "Name", remoteAddress: (*conn).RemoteAddr().String(), isReady: false, conn: conn}
	createNewPlayer(newPlayer, &allPlayers, conn)
	sendRoomsToAllClients()
	socketReader(conn)
}

func main() {
	http.HandleFunc("/test", hello)
	http.HandleFunc("/", webSocketHandler)
	log.Println("Serving at localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
