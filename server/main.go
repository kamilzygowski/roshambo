package main

import (
	"fmt"
	"log"
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
			if msg[0] == []byte("m")[0] {
				client.WriteMessage(msgType, msg) // Populate the message to other clients
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
	newPlayer := player{incrementingId, 15, 15, "Name", 1, (*conn).RemoteAddr().String()}
	createNewPlayer(newPlayer, &allPlayers)
	socketReader(conn)
}

func main() {
	http.HandleFunc("/test", hello)
	http.HandleFunc("/", webSocketHandler)
	log.Println("Serving at localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
