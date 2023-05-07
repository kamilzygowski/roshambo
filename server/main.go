package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is online"))
}

var allowOriginFunc = func(r *http.Request) bool {
	return true
}

func main() {
	server := socketio.NewServer(nil)

	server.OnConnect("/socket.io/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("Connected: ", s.ID())
		return nil
	})

	server.OnDisconnect("/socket.io/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	server.BroadcastToRoom("", "bcast", "render", "")

	allPlayers := []player{}
	createNewPlayer(player{0, 5, 5, "Newplayer", 1}, &allPlayers)
	createNewPlayer(player{0, 10, 10, "SecondOne", 2}, &allPlayers)
	fmt.Println(allPlayers)

	//go server.Serve()
	//defer server.Close()

	http.HandleFunc("/", hello)
	//http.HandleFunc("/socket.io", socketHandler)
	http.Handle("/socket.io/", server)
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
	/*for {
		//fmt.Println(allPlayers)
		//allPlayers[0].move(0)
	}*/
}
