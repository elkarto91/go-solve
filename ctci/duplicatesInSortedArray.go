package main

import "fmt"

func removeDuplicates(nums []int) int {

	l := len(nums)

	fmt.Println("L= ", l)
	if l == 0 || l == 1 {
		return l
	}

	slowPointer := 0
	var fastPointer int
	for v := range nums {
		fmt.Println("V ", v, " nums[v] ", nums[v])
	}
	for fastPointer = 0; fastPointer <= l-1; fastPointer++ {

		if fastPointer+1 <= l-1 {
			fmt.Println("SP= ", slowPointer, " FP ", fastPointer, " nums at SP and FP and FP+1 ", nums[slowPointer], nums[fastPointer], nums[fastPointer+1])

			if nums[fastPointer] != nums[fastPointer+1] {
				fmt.Println("Array Before ", nums)
				nums[slowPointer] = nums[fastPointer]
				slowPointer = slowPointer + 1
				fmt.Println("LP Now ", slowPointer)
				fmt.Println("Array Now ", nums)
			}
		} else if nums[slowPointer] != nums[fastPointer] {
			nums[slowPointer] = nums[fastPointer]
		} else {
			fmt.Println("Continuing Loop")
		}
	}

	fmt.Println("Array >> ", nums[:slowPointer+1])
	return slowPointer + 1
}

func main() {

	i := []int{1, 1, 2, 2, 2, 3, 4, 5, 6, 7, 7, 7, 7, 7, 8, 9}
	fmt.Println("length ", len(i))
	x := removeDuplicates(i)
	fmt.Println("Length now ", x)
}
