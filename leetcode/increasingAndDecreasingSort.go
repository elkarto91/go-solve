package main

import "fmt"

func main() {
	fmt.Println(`Enter the number of integers`)
	var size, num int

	fmt.Scanf("%d%d", &size, &num)

	if m, err := Scan(&size); m != 1 {
		panic(err)
	}
	fmt.Println(`Enter the integers`)
	all := make([]int, size)
	ReadN(all, 0, size)
	fmt.Println(all)
}

func ReadN(all []int, i, n int) {
	if n == 0 {
		return
	}
	if m, err := Scan(&all[i]); m != 1 {
		panic(err)
	}
	ReadN(all, i+1, n-1)
}

func Scan(a *int) (int, error) {
	return fmt.Scan(a)
}
