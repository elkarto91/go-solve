package main

import (
	"fmt"
)

type Player struct {
	PlayerName   string
	PlayerNumber int
}

type Rectangle struct {
	Width  int
	Height int
}
type Circle struct {
	Radius int
}

type Shapes interface {
	Area()
}

func PlayerWages(pl Player) int {
	return 1000 * pl.PlayerNumber
}

func PlayerWagePtr(pl *Player) int {
	return 5000 * pl.PlayerNumber
}

func (r Rectangle) Area() {

	fmt.Println(r.Height * r.Width)
}

func (c Circle) Area() {

	fmt.Println(c.Radius * c.Radius)
}

func calculateArea(s Shapes) {
	s.Area()
}

func main() {

	playerVariable := Player{
		PlayerName:   "Luka",
		PlayerNumber: 7,
	}

	fmt.Println("Player Name", playerVariable.PlayerName)

	x := PlayerWages(playerVariable)
	fmt.Println(" wage ", x)

	y := PlayerWagePtr(&playerVariable)
	fmt.Println("wage ", y)

	r := Rectangle{
		Width:  4,
		Height: 343,
	}
	calculateArea(r)

	c := Circle{Radius: 9}
	calculateArea(c)

}
