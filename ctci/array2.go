package main

import (
	"fmt"
	"os"
)

//Given two strings find if one string is permutation of another

var targetString = "acdfbe"
var counter = 0

//var targetString ="acdc"

func main() {

	baseString := "abcdef"

	fmt.Println("The base string is ", baseString)
	fmt.Println("The target string is ", targetString)

	fmt.Println("Premutations ...")
	baseStringRune := []rune(baseString)
	startIndex := 0
	endIndex := len(baseStringRune) - 1
	Permuations(baseStringRune, startIndex, endIndex)
	fmt.Println("counter ...", counter)

}
func swap(str []rune, xIndex, yIndex int) []rune {
	temp := str[xIndex]
	str[xIndex] = str[yIndex]
	str[yIndex] = temp
	return str
}
func Permuations(base []rune, leftIndex, rightIndex int) {

	if leftIndex == rightIndex {
		counter++
		fmt.Println("The Combination Now ", string(base))
		if string(base) == targetString {
			fmt.Println("Combination Found", string(base), targetString)
			os.Exit(0)
		}
		return
	}
	for i := leftIndex; i <= rightIndex; i++ {
		swap(base, leftIndex, i)
		Permuations(base, leftIndex+1, rightIndex)
		swap(base, leftIndex, i)
	}
}
