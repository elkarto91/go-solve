package main

import (
	"fmt"
	"strings"
)

func romanToDecimal(roman string) {

	romanArray := strings.Split(roman, "")
	fmt.Println("Roman Array ", romanArray)

}
func getValue(r string) int {

	if r == "I" {
		return 1

	} else if r == "V" {
		return 5

	} else if r == "X" {
		return 10

	} else if r == "L" {

		return 50

	} else if r == "C" {

		return 100

	} else if r == "D" {

		return 500

	} else if r == "M" {

		return 1000

	}
	return -1
}
func main() {

	str := "MCMIV"
	fmt.Println("Integer form of Roman Numeral is ")
	romanToDecimal(str)
}
