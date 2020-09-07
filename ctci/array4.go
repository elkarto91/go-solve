package main

import (
	"fmt"
	"os"
)

//two strings; three operations
//add,remove and replace
//check if they are one or 0 edit away

var stringTwo = "bale"

func main() {

	stringOne := "pale"

	fmt.Println("String One ", stringOne)
	fmt.Println("String Two ", stringTwo)

	matchFlag := matchFinder(stringOne, stringTwo)
	if matchFlag == true {
		fmt.Println("Its a match ", stringOne, stringTwo)
	} else {

		makeChange(stringOne, stringTwo)

	}

}
func matchFinder(s1, s2 string) bool {

	if s1 == s2 {
		fmt.Println("Match")
		return true
	}
	return false
}

func makeChange(stringOne, stringTwo string) {

	s1 := len(stringOne)
	s2 := len(stringTwo)

	if s1 > s2 {

		fmt.Println("First Change Add to Two ", stringTwo)

	} else if s1 < s2 {

		fmt.Println("First Change Remove From Two ")

	} else {
		fmt.Println("First Change Replace In Two ", stringTwo)

	}
}
func removeFunc(string2, string1 string) {

	charSet := []rune(string1)
	stringSet := []rune(string2)

	for _, v := range charSet {

		x := isIn(string(v), stringSet)
		if x == false {
			var slic []rune
			for _, u := range stringSet {
				if u != v {
					slic = append(slic, u)
					l := len(slic)
					PermuationSystem(slic, 0, l-1)
				} else {
					continue
				}
			}
		} else {
			continue
		}
	}
	return
}

func addFunc(string2, string1 string) {

	charSet := []rune(string1)
	stringS := []rune(string2)

	for _, v := range charSet {
		stringSet := stringS
		stringSet = append(stringSet, v)
		l := len(stringSet)
		PermuationSystem(stringSet, 0, l-1)
	}
	fmt.Println("No Match")
	return
}

func PermuationSystem(base []rune, leftIndex, rightIndex int) {

	if leftIndex == rightIndex {
		counter++
		fmt.Println("The Combination Now ", string(base))
		if string(base) == stringTwo {
			fmt.Println("Combination Found", string(base), targetString)
			os.Exit(0)
		}
		return
	}
	for i := leftIndex; i <= rightIndex; i++ {
		sswap(base, leftIndex, i)
		Permuations(base, leftIndex+1, rightIndex)
		sswap(base, leftIndex, i)
	}
}
func sswap(str []rune, xIndex, yIndex int) []rune {
	temp := str[xIndex]
	str[xIndex] = str[yIndex]
	str[yIndex] = temp
	return str
}

func isIn(a string, list []rune) bool {

	for _, b := range list {
		if string(b) == a {
			return true
		}
	}
	return false
}
