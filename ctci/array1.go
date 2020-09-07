package main

import "fmt"

//Check if string has unique characters, no additional data structure( golang needs rune)

func main() {

	myString := "abcdeafghijkl"
	myString2 := "abcde"

	fmt.Println("First string ", myString)

	ms := []rune(myString)
	fmt.Println("Rune ", string(ms))

	duplicateFlag := DuplicateChecker(ms)
	if duplicateFlag == true {
		fmt.Println("Duplicate found in string")
	} else {
		fmt.Println("No Duplicates in that one")
	}

	fmt.Println("Second string ", myString2)
	ms2 := []rune(myString2)
	fmt.Println("Rune ", string(ms2))

	duplicateFlag = DuplicateChecker(ms2)
	if duplicateFlag == true {
		fmt.Println("Duplicate found in string")
	} else {
		fmt.Println("No Duplicates in that one")
	}

	fmt.Println("Done")

}
func DuplicateChecker(ms []rune) bool {
	var slowPointer int
	var fastPointer int
	duplicateFlag := false

	slowPointer = 0
	fastPointer = len(ms) - 1

	for slowPointer = 0; slowPointer < len(ms)-1; slowPointer++ {
		for fastPointer = 0; fastPointer < len(ms)-1; fastPointer++ {
			if slowPointer != fastPointer {
				fmt.Println("Slow Pointer ", string(ms[slowPointer]))
				fmt.Println("Fast Pointer ", string(ms[fastPointer]))

				if string(ms[slowPointer]) == string(ms[fastPointer]) {
					fmt.Println("Dupilcates ", string(ms[slowPointer]), string(ms[fastPointer]))
					duplicateFlag = true
					return duplicateFlag
				}
			}
		}
	}
	return duplicateFlag
}
