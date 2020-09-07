package main

import "fmt"

func main() {

	var numOfProducts = 0
	fmt.Scanf("%d", &numOfProducts)

	var prodID = ""
	fmt.Scanf("%s", &prodID)

	x := 10 * 10 * 10 * 10 * 10 * 10
	if numOfProducts >= 0 && numOfProducts < x {

		var laptopcount = 0
		var desktopcount = 0

		runeVal := []rune(prodID)

		if numOfProducts > 0 {

			for i := range runeVal {

				if string(runeVal[i]) == "a" || string(runeVal[i]) == "e" || string(runeVal[i]) == "i" || string(runeVal[i]) == "o" || string(runeVal[i]) == "u" || string(runeVal[i]) == "A" || string(runeVal[i]) == "E" || string(runeVal[i]) == "I" || string(runeVal[i]) == "O" || string(runeVal[i]) == "U" {

					laptopcount++
				} else {
					desktopcount++

				}
			}

		}

		fmt.Println(desktopcount)

	}

}
