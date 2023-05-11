package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

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

func getPlayerFromClient(conn net.Addr, _allPlayers *[]player) *player {
	var resultPlayer *player
	for _, p := range *_allPlayers {
		if (conn).String() == p.remoteAddress {
			resultPlayer = &p
		}
	}
	return resultPlayer
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
				fmt.Println(client.RemoteAddr())
				for i := range allPlayers {
					if (allPlayers[i]).name == string(msgContent) {
						// IF PLAYER CLICKED ON HIS ROOM
						if client.RemoteAddr().String() != allPlayers[i].remoteAddress {
							if client.RemoteAddr().String() == (*conn).RemoteAddr().String() {
								fmt.Println(allPlayers[i])
								fmt.Println("JD2")
								allGames = append(allGames, Games{players: []string{client.RemoteAddr().String(), allPlayers[i].remoteAddress}, names: []string{}, isDone: false})
								(*allPlayers[i].conn).WriteMessage(1, []byte("g"))
								(*conn).WriteMessage(1, []byte("g"))
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
	newPlayer := player{id: incrementingId, name: "Name", remoteAddress: (*conn).RemoteAddr().String(), inGame: false, conn: conn}
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
