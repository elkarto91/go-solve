package main

import (
	"fmt"
	"strings"
)

func romanToDecimal(roman string) {

	romanArray := strings.Split(roman, "")
	fmt.Println("Roman Array ", romanArray)
	length := len(romanArray)
	res := 0
	for i := 0; i < length; i++ {

		k := getValue(romanArray[i])

		if i+1 < length {
			v := getValue(romanArray[i+1])
			if k >= v {
				res = res + k
			} else {
				res = res + v - k
				i++
			}
		} else {
			res = res + k
			i++
		}

	}
	fmt.Println("The Roman Numeral and Decimal Number are ---- > ", roman, res)
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
