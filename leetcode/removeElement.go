package main

import "fmt"

func removeElement(nums []int, val int) int {

	l := len(nums)
	if l == 0 {
		return l
	}
	if l == 1 {

		if nums[0] == val {
			return 0
		} else {
			return 1
		}
	}

	i := 0
	j := 0

	for j < len(nums) {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}

		j++
	}
	return i

}

func main() {

	i := []int{3, 2, 2, 3}
	x := 3
	y := removeElement(i, x)
	fmt.Println("Len ", i, x, y)
}
