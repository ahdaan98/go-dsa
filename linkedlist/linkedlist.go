package main

import "fmt"

type Operations interface {
	AddNodeAtBeginning(val int)
	AddNodeAtEnd(val int)
	AddNodeBeforeAValue(before,new int)
	AddNodeAfterAValue(after,new int)
	DeleteNode(val int)
	Sort()
	DeleteDuplicate()
	Print()
	PrintReverse()
}

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedListOperations() Operations {
	return &LinkedList{}
}

func (li *LinkedList) AddNodeAtBeginning(val int) {
	newNode := &Node{Val: val, Next: nil}
	if li.head == nil {
		li.head = newNode
		return
	}

	newNode.Next = li.head
	li.head = newNode
}

func (li *LinkedList) AddNodeAtEnd(val int) {
	if li.head == nil{
		li.head = &Node{Val: val,Next: nil}
		return
	}

	current := li.head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = &Node{Val: val,Next: nil}
}

func (li *LinkedList) DeleteNode(val int) {
	if li.head == nil {
		fmt.Println("List is empty")
		return
	}

	if li.head.Val == val {
		li.head = li.head.Next
		return
	}

	var prev *Node
	current := li.head
	for current != nil{
		if current.Val == val {
			prev.Next = current.Next
			return
		}
		prev = current
		current = current.Next
	}
}

func (li *LinkedList) AddNodeBeforeAValue(before,new int){
	if li.head == nil {
		fmt.Println("List is empty")
		return
	}

	if li.head.Val == before {
		newNode := &Node{Val: new,Next: li.head}
		li.head = newNode
		return
	}

	var prev *Node
	current := li.head
	for current != nil {
		if current.Val == before{
			newNode := &Node{Val: new,Next: current}
			prev.Next = newNode
			return
		}
		prev = current
		current = current.Next
	}
}

func (li *LinkedList) AddNodeAfterAValue(after,new int){
	if li.head == nil {
		fmt.Println("List is empty")
		return
	}

	current := li.head
	for current != nil {
		if current.Val == after {
			newNode := &Node{Val: new,Next: current.Next}
			current.Next = newNode
			return
		}
		current = current.Next
	}
}

func (li * LinkedList) Sort() {
	if li.head == nil || li.head.Next == nil {
		return
	}

	swapped := true
	for swapped {
		swapped = false
		prev := li.head
		current:= li.head.Next

		for current != nil {
			if prev.Val > current.Val {
				prev.Val,current.Val = current.Val,prev.Val
				swapped = true
			} 
			prev = current
			current = current.Next
		}
	}
}

func (li *LinkedList) DeleteDuplicate() {
	if li.head == nil {
		return
	}

	current := li.head
	for current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
}

func (li *LinkedList) Print() {
	current := li.head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
}

func (li *LinkedList) PrintReverse() {
	li.PrintReverseHelper(li.head)
}

func (li *LinkedList) PrintReverseHelper(node *Node){
	if node == nil{
		return
	}

	li.PrintReverseHelper(node.Next)
	fmt.Print(node.Val," ")
}

func main() {
	li := NewLinkedListOperations()

	nums := []int{1, 9, 10, 4, 5, 9, 5}
	for _, n := range nums {
		li.AddNodeAtBeginning(n)
	}

	li.AddNodeAtEnd(6)

	li.DeleteNode(1)

	li.AddNodeBeforeAValue(8,13)

	li.Sort()
	li.DeleteDuplicate()

	fmt.Println("Orginal Linked List:")
	li.Print()
	fmt.Println()
	fmt.Println("Reversed Linked List:")
	li.PrintReverse()
}
