package main

type player struct {
	id       uint8
	x        int32
	y        int32
	name     string
	vocation uint8
}

const (
	North = 0
	South = 1
	East  = 2
	West  = 3
)

func createNewPlayer(_player player, allPlayers *[]player) {
	newPlayer := player{id: 0, x: _player.x, y: _player.y, name: _player.name, vocation: _player.vocation}
	*allPlayers = append(*allPlayers, newPlayer)
}

func (p *player) move(direction uint8) {
	(*p).x++
}
