package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("Connected: ", s.ID())
		return nil
	})

	allPlayers := []player{}
	fmt.Println(allPlayers)
	createNewPlayer(player{0, 5, 5, "Newplayer", 1}, &allPlayers)
	createNewPlayer(player{0, 10, 10, "SecondOne", 2}, &allPlayers)
	fmt.Println(allPlayers)
	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io", server)
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
	/*for {
		//fmt.Println(allPlayers)
		//allPlayers[0].move(0)
	}*/
}
