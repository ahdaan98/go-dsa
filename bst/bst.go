package main

import "fmt"

type Node struct {
	key   int
	left  *Node
	Right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) insert(data int) {
	if t.root == nil {
		t.root = &Node{key: data}
	} else {
		t.root.insert(data)
	}
}

func (n *Node) insert(data int) {
	if data <= n.key {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{key: data}
		} else {
			n.Right.insert(data)
		}
	}
}

func (t *Tree) Search(key int) bool {
	var found bool
	if t.root == nil {
		found = false
	} else if t.root.key != key {
		found = t.root.search(key)
	} else {
		found = true
	}

	return found
}

func (n *Node) search(key int) bool{
	var found bool
	if n == nil{
		found = false
	} else if key == n.key {
		found = true
	} else {
		if n.key > key {
			found = n.left.search(key)
		} else {
			found = n.Right.search(key)
		}
	}
	return found
}

func (t *Tree) validbst() bool {
	if t == nil {
		return true
	}
	return t.root.validbst(-99999,99999)
}

func (n *Node) validbst(min,max int) bool {
	if n == nil {
		return true
	}

	if n.key <= min || n.key >= max {
		return false
	}

	return n.left.validbst(min,n.key) && n.Right.validbst(n.key,max)
}

func PrintPreOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.key)
		PrintPreOrder(n.left)
		PrintPreOrder(n.Right)
	}
}

func PrintPostOrder(n *Node) {
	if n == nil {
		return
	} else {
		PrintPostOrder(n.left)
		PrintPostOrder(n.Right)
		fmt.Printf("%d ", n.key)
	}
}

func PrintInOrder(n *Node) {
	if n == nil {
		return
	} else {
		PrintPostOrder(n.left)
		fmt.Printf("%d ", n.key)
		PrintPostOrder(n.Right)
	}
}

func main() {
	t := Tree{}
	nums := []int{1, 3, 9, 7, 5}

	for _, n := range nums {
		t.insert(n)
	}

	fmt.Println("Pre Order:")
	PrintPreOrder(t.root)
	fmt.Println("\nPost Order:")
	PrintPostOrder(t.root)
	fmt.Println("\nIn Order:")
	PrintInOrder(t.root)

	fmt.Printf("valid bst: %v",t.validbst())

	fmt.Println("\nKey exist:",t.Search(7))
}