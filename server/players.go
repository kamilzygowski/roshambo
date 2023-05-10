package main

import (
	"strconv"
)

type player struct {
	id            uint16
	name          string
	remoteAddress string
	inGame        bool
}

const (
	North = 0
	South = 1
	East  = 2
	West  = 3
)

func createNewPlayer(_player player, allPlayers *[]player) {
	newPlayer := player{id: _player.id, name: _player.generateName(), remoteAddress: _player.remoteAddress}
	*allPlayers = append(*allPlayers, newPlayer)
}

func (p *player) generateName() string {
	return "Player" + strconv.Itoa(int((*p).id))
}
