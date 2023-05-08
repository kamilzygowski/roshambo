package main

import (
	"testing"
)

func TestCreateNewPlayer(t *testing.T) {
	testPlayers := []player{}
	newTestPlayer := player{1, 5, 5, "TestName", 3, "ANY_ADRESS"}
	createNewPlayer(newTestPlayer, &testPlayers)
	if len(testPlayers) <= 0 {
		t.Errorf("Player creation error")
	}
}
