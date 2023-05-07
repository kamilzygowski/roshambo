package main

import "fmt"

func main() {
	allPlayers := []player{}
	fmt.Println(allPlayers)
	createNewPlayer(player{5, 5, "Newplayer", 1}, &allPlayers)
	fmt.Println(allPlayers)
	fmt.Println("Running")
	for {
		fmt.Println(allPlayers)
		allPlayers[0].move()
	}
}
