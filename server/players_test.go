package main

import (
	"fmt"
	"testing"

	"github.com/gorilla/websocket"
)

func TestCreateNewPlayer(t *testing.T) {
	testPlayers := []Player{}
	newTestPlayer := Player{id: 1, name: "TestName", remoteAddress: "ANY_ADRESS", isReady: false}
	createNewPlayer(newTestPlayer, &testPlayers, &websocket.Conn{})
	if len(testPlayers) <= 0 {
		t.Errorf("Player creation error")
	}
}

func TestRemoveFromSlice(t *testing.T) {
	testSlice := []Player{
		{id: 1, name: "TestName", remoteAddress: "ANY_ADRESS", isReady: false},
		{id: 2, name: "TestName", remoteAddress: "ANY_ADRESS", isReady: false},
		{id: 3, name: "TestName", remoteAddress: "ANY_ADRESS", isReady: false}}
	testSlice = removeFromSlice(testSlice, 1)
	if len(testSlice) != 2 || (testSlice[1].id != 3 && testSlice[0].id != 1) {
		t.Errorf("removeFromSlice() function error")
	}
}

func TestSetReady(t *testing.T) {
	var testPlayer Player = Player{id: 1, name: "TestName", remoteAddress: "ANY_ADRESS", isReady: false, choice: 0}
	var testChoice uint8 = 2
	testPlayer.setReady(testChoice)
	if testPlayer.choice != 2 || testPlayer.isReady != true {
		t.Errorf("setReady() receiver function error")
	}
}

func TestGenerateName(t *testing.T) {
	var testPlayerId uint16 = 1
	var testPlayer Player = Player{id: testPlayerId, name: "Test name", remoteAddress: "ANY_ADRESS", isReady: false, choice: 0}
	generatedName := testPlayer.generateName()
	testPlayer.name = generatedName
	if testPlayer.name != fmt.Sprintf("Player%d", testPlayerId) {
		t.Errorf("generateName() receiver function error")
	}
}
