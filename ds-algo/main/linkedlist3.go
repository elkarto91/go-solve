package main

import "fmt"

type ele struct {
	name string
	next *ele
}

type singleList struct {
	len  int
	head *ele
}

func initList() *singleList {
	return &singleList{}
}

func (s *singleList) AddFront(data string) {
	e := &ele{
		name: data,
		next: nil,
	}

	if s.head == nil {
		s.head = e
	} else {
		e.next = s.head
		s.head = e
	}
	s.len++
	return
}
func (s *singleList) AddN(data string, n int) {

	e := &ele{
		name: data,
		next: nil,
	}
	if s.len == 1 {
		s.head = e
	} else {

		current := s.head
		for i := 0; i < n-1; i++ {
			current = current.next
		}
		e.next = current.next
		current.next = e
	}
	s.len++
	return
}
func (s *singleList) AddBack(data string) {

	e := &ele{
		name: data,
		next: nil,
	}
	if s.head == nil {
		s.head = e
	} else {

		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = e
	}
	s.len++
	return
}
func (s *singleList) Traverse() error {
	if s.head == nil {
		return fmt.Errorf("TranverseError: List is empty")
	}
	current := s.head
	for current != nil {
		fmt.Println(current.name)
		current = current.next
	}
	return nil
}
func (s *singleList) Size() int {
	return s.len
}
func (s *singleList) RemoveFront() error {
	if s.head == nil {
		return fmt.Errorf("List is empty")
	}
	s.head = s.head.next
	s.len--
	return nil
}
func (s *singleList) RemoveBack() error {
	if s.head == nil {
		return fmt.Errorf("removeBack: List is empty")
	}
	var prev *ele
	current := s.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev != nil {
		prev.next = nil
	} else {
		s.head = nil
	}
	s.len--
	return nil
}
func (s *singleList) RemoveN(n int) error {

	if s.head == nil {
		return fmt.Errorf("removeBack: List is empty")
	}

	current := s.head
	for i := 0; i < n-2; i++ {
		current = current.next
	}
	current.next = current.next.next
	s.len--
	return nil
}

func RecursivePrintList(node *ele) {

	if node == nil {
		return
	}
	fmt.Println("Node Value ", node.name)
	RecursivePrintList(node.next)
}

func ReversePrintList(node *ele) {

	if node == nil {
		return
	}
	ReversePrintList(node.next)
	fmt.Println("Node Value ", node.name)
}

func (s *singleList) ReverseLinkedList(node *ele) {
	if node.next == nil {
		s.head = node
		return
	}
	s.ReverseLinkedList(node.next)
	prev := node.next
	prev.next = node
	node.next = nil
}
func main() {

	fmt.Println("Adding Front")
	singleList := initList()
	fmt.Printf("AddFront: A\n")
	singleList.AddFront("A")
	fmt.Printf("AddFront: B\n")
	singleList.AddFront("B")
	fmt.Printf("AddBack: C\n")
	singleList.AddBack("C")

	fmt.Printf("AddBack: D\n")
	singleList.AddBack("D")

	fmt.Printf("AddBack: E\n")
	singleList.AddBack("E")

	fmt.Printf("Size: %d\n", singleList.Size())
	err := singleList.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Remove Front: \n")
	err = singleList.RemoveFront()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Size: %d\n", singleList.Size())
	err = singleList.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Remove Back: \n")
	err = singleList.RemoveBack()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Size: %d\n", singleList.Size())
	err = singleList.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	n := 2
	fmt.Println("Add n: ", n)
	singleList.AddN("W", n)
	fmt.Println("Added in nth position")

	fmt.Printf("Size: %d\n", singleList.Size())
	err = singleList.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	n = 4
	fmt.Println("Add n: ", n)
	singleList.AddN("Z", n)
	fmt.Println("Added in nth position")

	fmt.Printf("Size: %d\n", singleList.Size())
	err = singleList.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	n = 3
	fmt.Println("Add n: ", n)
	singleList.AddN("X", n)
	fmt.Println("Added in nth position")

	fmt.Printf("Size: %d\n", singleList.Size())
	err = singleList.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	m := 2
	fmt.Println("Remove at m", m)
	err = singleList.RemoveN(m)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("removed in nth position", m)

	fmt.Printf("Size: %d\n", singleList.Size())
	err = singleList.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Recursive Print")

	headNode := singleList.head
	RecursivePrintList(headNode)

	fmt.Println("Reverse Print")
	ReversePrintList(headNode)

	fmt.Println("Reverse Linked List")
	singleList.ReverseLinkedList(headNode)

	newHead := singleList.head
	RecursivePrintList(newHead)

}
