package main

import "fmt"

func main() {

	val := 121
	x := isPalindrome(val)
	fmt.Println("X--> ", val, x)
}
func isPalindrome(x int) bool {

	original := x

	reverse := 0
	for x > 0 {

		reverse = reverse*10 + x%10
		x = x / 10
	}

	fmt.Println("reverse and original ", reverse, original)

	if original == reverse {
		return true
	} else {
		return false
	}
}
