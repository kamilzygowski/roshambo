package main

import (
	"testing"
)

func TestCreateNewPlayer(t *testing.T) {
	testPlayers := []player{}
	newTestPlayer := player{3, 5, 5, "TestName", 3}
	createNewPlayer(newTestPlayer, &testPlayers)
	if len(testPlayers) <= 0 {
		t.Errorf("Player creation error")
	}
}
