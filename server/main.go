package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

var allowOriginFunc = func(r *http.Request) bool {
	return true
}

func main() {
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})

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
