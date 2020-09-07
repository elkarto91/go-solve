package main

import "fmt"

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

	fmt.Println("Adding things to empty stack", stack)

	stack = push(stack, "1")
	stack = push(stack, "2")
	stack = push(stack, "3")
	stack = push(stack, "4")
	stack = push(stack, "5")

	fmt.Println("New stack", stack)

	T := topValue(stack)
	fmt.Println("top element", T)

	stack = pop(stack)
	fmt.Println("Popped stack", stack)

	T = topValue(stack)
	fmt.Println("top element", T)

	stack = pop(stack)
	fmt.Println("Popped stack", stack)

	T = topValue(stack)
	fmt.Println("top element", T)

	stack = pop(stack)
	fmt.Println("Popped stack", stack)
}
