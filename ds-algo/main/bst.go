package main

import "fmt"

type bstnode struct {
	value int
	left  *bstnode
	right *bstnode
}

type bst struct {
	root *bstnode
}

func (b *bst) reset() {
	b.root = nil
}

func (b *bst) insert(data int) {

	b.insertRec(b.root, data)

}

func (b *bst) insertRec(node *bstnode, value int) *bstnode {

	//If root is nil insert value at root
	if b.root == nil {
		b.root = &bstnode{
			value: value,
		}
		return b.root
	}

	//If node is nil , insert value at the location
	if node == nil {
		return &bstnode{value: value}
	}

	//If value is less than node value go to left
	if value <= node.value {
		node.left = b.insertRec(node.left, value)
	}
	//If value is greater than node go right
	if value > node.value {
		node.right = b.insertRec(node.right, value)
	}

	return node
}
func (b *bst) find(value int) error {
	node := b.findRec(b.root, value)
	if node == nil {
		return fmt.Errorf("Value: %d not found in tree", value)
	}
	return nil
}

func (b *bst) findRec(node *bstnode, value int) *bstnode {
	if node == nil {
		return nil
	}
	if node.value == value {
		return b.root
	}
	if value < node.value {
		return b.findRec(node.left, value)
	}
	return b.findRec(node.right, value)
}

func (b *bst) inorder() {
	b.inorderRec(b.root)
}

func (b *bst) inorderRec(node *bstnode) {
	if node != nil {
		b.inorderRec(node.left)
		fmt.Println(node.value)
		b.inorderRec(node.right)
	}
}
func (b *bst) preorder() {
	b.preorderRec(b.root)
}

func (b *bst) preorderRec(node *bstnode) {
	if node != nil {
		fmt.Println(node.value)
		b.inorderRec(node.left)
		b.inorderRec(node.right)
	}
}

func (b *bst) postorder() {
	b.postorderRec(b.root)
}
func (b *bst) postorderRec(node *bstnode) {
	if node != nil {
		b.inorderRec(node.left)
		b.inorderRec(node.right)
		fmt.Println(node.value)
	}
}

func main() {

	fmt.Println("Binary Search Tree")

	bst := &bst{}
	eg := []int{2, 5, 7, -1, -1, 5, 5}
	for _, val := range eg {
		bst.insert(val)
	}

	fmt.Printf("\nPrinting Inorder:\n")
	bst.inorder()

	fmt.Printf("\nPrinting Preorder:\n")
	bst.preorder()

	fmt.Printf("\nPrinting Postorder:\n")
	bst.postorder()
}
