package main

type player struct {
	x        int32
	y        int32
	name     string
	vocation int8
}

func createNewPlayer(_player player, allPlayers *[]player) {
	newPlayer := player{x: _player.x, y: _player.y, name: _player.name, vocation: _player.vocation}
	*allPlayers = append(*allPlayers, newPlayer)
}

func (p *player) move() {
	(*p).x++
}
