package main

import (
	"strconv"
)

type player struct {
	id            uint16
	x             int32
	y             int32
	name          string
	vocation      uint8
	remoteAddress string
}

const (
	North = 0
	South = 1
	East  = 2
	West  = 3
)

func createNewPlayer(_player player, allPlayers *[]player) {
	newPlayer := player{id: _player.id, x: _player.x, y: _player.y, name: _player.generateName(), vocation: _player.vocation, remoteAddress: _player.remoteAddress}
	*allPlayers = append(*allPlayers, newPlayer)
}

func (p *player) generateName() string {
	return "Player" + strconv.Itoa(int((*p).id))
}

func (p *player) move(direction uint8) {
	(*p).x++
}
