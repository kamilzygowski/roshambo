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

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is online"))
}

func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // Very important
	conn, _ := upgrader.Upgrade(w, r, nil)

	clients = append(clients, *conn)

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		fmt.Printf("%s send: %s\n", conn.RemoteAddr(), string(msg))

		/*for _, client := range clients {
			if err = client.WriteMessage(msgType, msg); err != nil {
				return
			}
		}*/
		clients[0].WriteMessage(msgType, msg) // Test back message
	}
}

func main() {
	allPlayers := []player{}
	createNewPlayer(player{0, 5, 5, "Newplayer", 1}, &allPlayers)
	createNewPlayer(player{0, 10, 10, "SecondOne", 2}, &allPlayers)
	fmt.Println(allPlayers)

	http.HandleFunc("/test", hello)
	http.HandleFunc("/", webSocketHandler)
	log.Println("Serving at localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
