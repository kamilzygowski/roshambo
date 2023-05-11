package main

import (
	"strconv"

	"github.com/gorilla/websocket"
)

type player struct {
	id            uint16
	name          string
	remoteAddress string
	isReady       bool
	conn          *websocket.Conn
	choice        uint8
}

type Games struct {
	isDone  bool
	players []player
}

const (
	North = 0
	South = 1
	East  = 2
	West  = 3
)

func createNewPlayer(_player player, allPlayers *[]player, conn *websocket.Conn) {
	newPlayer := player{id: _player.id, name: _player.generateName(), remoteAddress: _player.remoteAddress, conn: conn, isReady: false}
	*allPlayers = append(*allPlayers, newPlayer)
}

func (p *player) generateName() string {
	return "Player" + strconv.Itoa(int((*p).id))
}

func (p *player) setReady(choice uint8) {
	p.isReady = true
	p.choice = choice
}
