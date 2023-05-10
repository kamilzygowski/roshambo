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

func socketReader(conn *websocket.Conn) {
	for {
		msgType, msg, err := (*conn).ReadMessage()
		if err != nil {
			// Handling error / disconnect
			fmt.Printf("User %s has disconnected\n", (*conn).RemoteAddr())
			// Removing client
			for index, client := range clients {
				if client.RemoteAddr() == (*conn).RemoteAddr() {
					clients = append(clients[:index], clients[index+1:]...)
				}
			}
			return
		}

		fmt.Printf("%s send: %s\n", (*conn).RemoteAddr(), string(msg))
		fmt.Println("Number of clients: ", len(clients))

		for _, client := range clients {
			fmt.Println(client.RemoteAddr())
			if msg[0] == []byte("m")[0] {
				// Chat message
				client.WriteMessage(msgType, msg) // Populate the message to other clients
			} else if msg[0] == []byte("n")[0] { //&& ((*getPlayerFromClient(client.RemoteAddr(), &allPlayers)).name == client.RemoteAddr().String()) {
				// Name
				msgContent := msg[1:]
				for i := range allPlayers {
					if (allPlayers[i]).remoteAddress == (*conn).RemoteAddr().String() {
						(allPlayers[i]).name = string(msgContent)
						fmt.Println(allPlayers[i])
						fmt.Println("JD")
					}
				}

			} else if msg[0] == []byte("r")[0] {
				// Room
				fmt.Println(client.RemoteAddr())
				for i := range allPlayers {
					if (allPlayers[i]).remoteAddress == client.RemoteAddr().String() {
						fmt.Println(allPlayers[i])
						fmt.Println("JD2")
					}
					//p := getPlayerFromClient(client.RemoteAddr(), &allPlayers)
				}
			}
		}
	}
}

func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // Very important
	conn, _ := upgrader.Upgrade(w, r, nil)

	clients = append(clients, *conn)
	// Add and init player
	incrementingId++
	newPlayer := player{id: incrementingId, name: "Name", remoteAddress: (*conn).RemoteAddr().String(), inGame: false}
	createNewPlayer(newPlayer, &allPlayers)
	socketReader(conn)
}

func main() {
	http.HandleFunc("/test", hello)
	http.HandleFunc("/", webSocketHandler)
	log.Println("Serving at localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
