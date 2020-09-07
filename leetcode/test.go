package main

import (
	"fmt"
	"strings"
)

var stack1 []string

func isValid(s string) bool {

	fmt.Println("Original ", s)
	if s == "" {
		return true
	}

	set := strings.Split(s, "")
	fmt.Println("String ", set)
	if len(set) < 2 {
		return false
	}
	for _, v := range set {

		if v == "{" || v == "[" || v == "(" {

			push1(v)
			fmt.Println("stack now ", stack1)

		} else {

			x := pop1()
			fmt.Println("Char ", x)

			if v == "}" && x == "{" {

				continue

			} else if v == "]" && x == "[" {

				continue

			} else if v == ")" && x == "(" {

				continue

			} else {

				return false
			}
		}
	}
	if len(stack1) > 0 {
		return false
	}
	return true
}

func push1(e string) {
	stack1 = append(stack1, e)
	return
}

func pop1() string {
	l := len(stack1)
	x := stack1[l-1]
	stack1 = stack1[:l-1]
	return x
}

func main() {
	x := "{[]}"
	t := isValid(x)
	fmt.Println(t)
}
