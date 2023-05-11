package main

import (
	"fmt"
	"testing"

	"github.com/gorilla/websocket"
)

func TestCreateNewPlayer(t *testing.T) {
	testPlayers := []player{}
	newTestPlayer := player{id: 1, name: "TestName", remoteAddress: "ANY_ADRESS", inGame: false}
	createNewPlayer(newTestPlayer, &testPlayers, &websocket.Conn{})
	if len(testPlayers) <= 0 {
		t.Errorf("Player creation error")
	}
}

func TestRemoveFromSlice(t *testing.T) {
	testSlice := []player{
		{id: 1, name: "TestName", remoteAddress: "ANY_ADRESS", inGame: false},
		{id: 2, name: "TestName", remoteAddress: "ANY_ADRESS", inGame: false},
		{id: 3, name: "TestName", remoteAddress: "ANY_ADRESS", inGame: false}}
	testSlice = removeFromSlice(testSlice, 1)
	fmt.Println(testSlice)
	if len(testSlice) != 2 || (testSlice[1].id != 3 && testSlice[0].id != 1) {
		t.Errorf("removeFromSlice() function error")
	}
}
