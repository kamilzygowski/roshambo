package main

import (
	"strconv"

	"github.com/gorilla/websocket"
)

type Player struct {
	id            uint16
	name          string
	remoteAddress string
	isReady       bool
	conn          *websocket.Conn
	choice        uint8
}

type Games struct {
	isDone  bool
	players []Player
	scores  []uint16
}

const (
	North = 0
	South = 1
	East  = 2
	West  = 3
)

func createNewPlayer(_player Player, allPlayers *[]Player, conn *websocket.Conn) {
	newPlayer := Player{id: _player.id, name: _player.generateName(), remoteAddress: _player.remoteAddress, conn: conn, isReady: false}
	*allPlayers = append(*allPlayers, newPlayer)
}

func (p *Player) generateName() string {
	return "Player" + strconv.Itoa(int((*p).id))
}

func (p *Player) setReady(choice uint8) {
	(*p).isReady = true
	(*p).choice = choice
}
