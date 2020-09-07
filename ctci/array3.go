package main

import (
	"fmt"
	"os"
)

//Given a string write a code to check if its permutation is a paliendrome

var counterVariable = 0

//var targetString ="acdc"

func main() {

	baseString := "paccpa"
	//baseString:="artur"

	fmt.Println("The base string is ", baseString)

	fmt.Println("Premutations ...")
	baseStringRune := []rune(baseString)
	startIndex := 0
	endIndex := len(baseStringRune) - 1
	PermuationForPaliendrome(baseStringRune, startIndex, endIndex)
	fmt.Println("counter ...", counterVariable)

}
func PermuationForPaliendrome(s []rune, lIndex, rIndex int) {

	if lIndex == rIndex {

		counterVariable++
		fmt.Println("The Permutation is ", string(s))
		fmt.Println("Checking Paliendrome ")
		flag := PaliendromeChecker(s)
		if flag == false {
			fmt.Println("No Paliendrome on this one")
			return
		} else {
			fmt.Println("Paliendrome !!!!!!!!")
			fmt.Println("Paliendrome ->", string(s))

			os.Exit(0)
			//return
		}

	}

	for i := lIndex; i <= rIndex; i++ {

		SwapForPaliendrome(s, lIndex, i)
		PermuationForPaliendrome(s, lIndex+1, rIndex)
		SwapForPaliendrome(s, lIndex, i)
	}
}

func PaliendromeChecker(s []rune) bool {

	length := len(s)
	startIndex := 0
	endIndex := length - 1
	median := length / 2

	fmt.Println("Length Start End ", length, startIndex, endIndex)

	for startIndex <= median {

		fmt.Println("Inner Length Start End ", length, startIndex, endIndex)

		if string(s[startIndex]) != string(s[endIndex]) {
			fmt.Println("No match", string(s[startIndex]), string(s[endIndex]))
			return false

		} else {
			startIndex++
			endIndex--
			continue
		}
	}
	return true
}
func SwapForPaliendrome(str []rune, xIndex, yIndex int) []rune {
	temp := str[xIndex]
	str[xIndex] = str[yIndex]
	str[yIndex] = temp
	return str
}
