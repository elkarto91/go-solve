package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Longest Prefix")
	list := []string{"flower", "flow", "flight"}
	fmt.Println("List ", list)

	x := longestCommonPrefix(list)
	fmt.Println("Longest Sub ", x)
}

func longestCommonPrefix(list []string) string {

	l := len(list)

	if l == 0 {
		return ""
	}

	if l == 1 {
		return list[0]
	}

	prefix := list[0]
	fmt.Println("Prefix init and length ", prefix, l)

	for i := 1; i < l; i++ {

		current := list[i]
		fmt.Println("Current ", current)

		j := 0

		p := strings.Split(prefix, "")
		q := strings.Split(current, "")
		fmt.Println("P and Q ", p, q)

		for j < len(prefix) && j < len(current) && p[j] == q[j] {
			fmt.Println("J and Prefix Length  ", j, len(prefix))
			fmt.Println("P[j] and Q[j] ", p[j], q[j])
			j++
		}
		if j == 0 {
			fmt.Println("Array ended ", j)

			return ""
		}
		fmt.Println("Prefix changes to -> ", prefix, current[:j])
		prefix = current[:j]
	}
	return prefix
}
