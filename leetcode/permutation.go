package main

import "fmt"

var count = 0

func main() {
	str := "abc"
	strRune := []rune(str)

	fmt.Println("The rune now is : ", string(strRune))
	fmt.Println("Generated Perms ")

	generatePermutation(strRune, 0, len(strRune)-1)

	fmt.Println(count)

}
func generatePermutation(str []rune, left, right int) {
	if left == right {
		count++
		fmt.Println("left ==right ", string(str))
	} else {

		for i := left; i <= right; i++ {
			fmt.Println("before switch 1 ", string(str))
			str[left], str[i] = str[i], str[left]
			fmt.Println("after switch 1", string(str))
			generatePermutation(str, left+1, right)
			fmt.Println("before switch 2", string(str))
			str[left], str[i] = str[i], str[left]
			fmt.Println("after switch 2", string(str))

		}
	}
}
