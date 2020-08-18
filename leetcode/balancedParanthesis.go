package main

import (
	"fmt"
	"strings"
)

var stack []string
var tops = 0

func push(stack []string, element string) []string {
	stack = append(stack, element)
	return stack
}

func top(stack []string) int {
	return len(stack)
}

func topValue(stack []string) string {
	val := stack[len(stack)-1]
	return val
}

func pop(stack []string) []string {
	stack = stack[:len(stack)-1]
	return stack
}
func main() {

	brackets := "{[()]}(("

	s := strings.Split(brackets, "")
	for k, v := range s {
		fmt.Println("String split to individual characters", k, v)

		tops = top(stack)
		if tops < 0 {
			fmt.Println("empty")
			return
		} else {
			if v == "{" || v == "[" || v == "(" {
				stack = push(stack, v)
				fmt.Println("pushing", v)

			} else {
				topVal := topValue(stack)
				if v == "}" && topVal == "{" {
					stack = pop(stack)
					fmt.Println("popping", v)

				} else if v == ")" && topVal == "(" {
					stack = pop(stack)
					fmt.Println("popping", v)

				} else if v == "]" && topVal == "[" {
					stack = pop(stack)
					fmt.Println("popping", v)

				} else {
					stack = push(stack, v)
					fmt.Println("pushing", v)

				}
			}
		}
	}
	tops = top(stack)
	fmt.Println("tops", tops)

	if tops > 0 {
		fmt.Println(" Unbalanced")
	}
}
