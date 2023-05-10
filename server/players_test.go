package main

import (
	"testing"
)

func TestCreateNewPlayer(t *testing.T) {
	testPlayers := []player{}
	newTestPlayer := player{id: 1, name: "TestName", remoteAddress: "ANY_ADRESS", inGame: false}
	createNewPlayer(newTestPlayer, &testPlayers)
	if len(testPlayers) <= 0 {
		t.Errorf("Player creation error")
	}
}
