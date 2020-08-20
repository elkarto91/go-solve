package main

import "fmt"

var counter = 0

func main() {
	str := "abcdefgh"
	runes := []rune(str)

	fmt.Println("The rune now is : ", string(runes))
	fmt.Println("Generated Perms ")

	producePermutation(runes, 0, len(runes)-1)

	fmt.Println(counter)

}

func producePermutation(arr []rune, left, right int) {

	if left == right {
		counter++
		fmt.Println("Combination ", counter, string(arr))
	}
	for i := left; i <= right; i++ {
		ar := swap(arr, left, i)
		producePermutation(ar, left+1, right)
	}
}
func swap(str []rune, i, j int) []rune {

	temp := str[i]
	str[i] = str[j]
	str[j] = temp
	return str
}
