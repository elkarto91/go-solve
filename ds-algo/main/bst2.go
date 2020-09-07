package main

import "fmt"

type binaryNode struct {
	left  *binaryNode
	right *binaryNode
	value int
}
type binaryTree struct {
	root *binaryNode
}

func (bst *binaryTree) InsertBST(val int) {
	bst.insertRecord(bst.root, val)
}
func (bst *binaryTree) insertRecord(node *binaryNode, val int) *binaryNode {

	if bst.root == nil {
		n := &binaryNode{
			value: val,
		}
		bst.root = n
		return bst.root
	}
	if node == nil {
		n := &binaryNode{
			value: val,
		}
		return n
	}

	if val <= node.value {
		node.left = bst.insertRecord(node.left, val)
	}
	if val > node.value {
		node.right = bst.insertRecord(node.right, val)
	}
	return node
}
func (bst *binaryTree) inorderCall() {
	bst.InorderFun(bst.root)
}
func (bst *binaryTree) InorderFun(node *binaryNode) {
	if node != nil {
		bst.InorderFun(node.left)
		fmt.Println(node.value)
		bst.InorderFun(node.right)
	}
}
func main() {

	bst := &binaryTree{
		root: nil,
	}
	fmt.Println("BST")
	intAr := []int{2, 5, 7, -1, -1, 5, 5}

	for k, v := range intAr {
		fmt.Println("Insert ", v, " from position ", k)
		bst.InsertBST(v)
	}

	fmt.Println("Inorder")
	bst.inorderCall()
}
