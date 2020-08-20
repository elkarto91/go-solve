package main

import (
	"fmt"
	"math"
)

func main() {

	x := 123
	y := reverse(x)
	fmt.Println("Original ", x, "Reverse ", y)
}
func reverse(x int) int {

	var num int
	var exponent, base float64
	base = 2
	exponent = 31
	MAX := int(math.Pow(base, exponent) - 1)

	for x != 0 {
		num = num*10 + x%10
		if num > MAX || num < -MAX {
			return 0
		}
		x = x / 10
	}
	return num
}
