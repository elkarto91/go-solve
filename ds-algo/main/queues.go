package main

import "fmt"

func main() {

	var queue []string

	fmt.Println("Add element")
	queue = append(queue, "one")
	queue = append(queue, "two")
	queue = append(queue, "three")
	queue = append(queue, "four")

	fmt.Println("queue ---> ", queue)

	fmt.Println("dequeue")
	queue[0] = ""
	queue = queue[1:]
	fmt.Println("new queue", queue)
}
