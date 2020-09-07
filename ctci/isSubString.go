package main

import (
	"fmt"
	"strings"
)

func main() {

	mainString := "waterbottle"
	fmt.Println("length ", len(mainString))
	fmt.Println("string ", mainString)

	subString := "erbottlewat"
	fmt.Println("Length substring ", len(subString))
	fmt.Println("string ", subString)

	s2 := mainString + mainString
	x := strings.Contains(s2, subString)
	fmt.Println("Answer ", x)

}
