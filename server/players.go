package main

import (
	"strconv"

	"github.com/gorilla/websocket"
)

type player struct {
	id            uint16
	name          string
	remoteAddress string
	inGame        bool
	conn          *websocket.Conn
}

type Games struct {
	isDone  bool
	players []string
	names   []string
}

const (
	North = 0
	South = 1
	East  = 2
	West  = 3
)

func createNewPlayer(_player player, allPlayers *[]player, conn *websocket.Conn) {
	newPlayer := player{id: _player.id, name: _player.generateName(), remoteAddress: _player.remoteAddress, conn: conn}
	*allPlayers = append(*allPlayers, newPlayer)
}

func (p *player) generateName() string {
	return "Player" + strconv.Itoa(int((*p).id))
}
