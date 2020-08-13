package main

import "fmt"

type Player struct {
	PlayerName   string
	PlayerNumber int
}

func main() {

	playerVariable := Player{
		PlayerName:   "Luka",
		PlayerNumber: 7,
	}

	fmt.Println("Player Name", playerVariable.PlayerName)

}
