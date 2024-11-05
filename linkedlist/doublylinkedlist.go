package main

// import "fmt"

// type Node struct {
// 	Val  int
// 	Prev *Node
// 	Next *Node
// }

// type DoublyLinkedList struct {
// 	head *Node
// 	tail *Node
// }

// func (dl *DoublyLinkedList) Append(val int) {
// 	newNode := &Node{Val: val}

// 	if dl.head == nil {
// 		dl.head = newNode
// 		dl.tail = newNode
// 	} else {
// 		newNode.Prev = dl.tail
// 		dl.tail.Next = newNode
// 		dl.tail = newNode
// 	}
// }

// func (dl *DoublyLinkedList) Prepend(val int) {
// 	newNode := &Node{Val: val}

// 	if dl.head == nil {
// 		dl.head = newNode
// 		dl.tail = newNode
// 	} else {
// 		newNode.Next = dl.head
// 		dl.head.Prev = newNode
// 		dl.head = newNode
// 	}
// }

// func (dl *DoublyLinkedList) Delete(val int) {
// 	if dl.head == nil {
// 		return
// 	}

// 	// If the node to delete is the head node
// 	if dl.head.Val == val {
// 		dl.head = dl.head.Next
// 		if dl.head != nil {
// 			dl.head.Prev = nil
// 		} else {
// 			dl.tail = nil // Update tail pointer if the list becomes empty
// 		}
// 		return
// 	}

// 	// If the node to delete is the tail node
// 	if dl.tail.Val == val {
// 		dl.tail = dl.tail.Prev
// 		if dl.tail != nil {
// 			dl.tail.Next = nil
// 		} else {
// 			dl.head = nil // Update head pointer if the list becomes empty
// 		}
// 		return
// 	}

// 	// Deleting a node in the middle of the list
// 	current := dl.head.Next
// 	for current != nil && current != dl.tail {
// 		if current.Val == val {
// 			current.Prev.Next = current.Next
// 			current.Next.Prev = current.Prev
// 			return
// 		}
// 		current = current.Next
// 	}
// }

// func (dl *DoublyLinkedList) AfterValue(after, val int) {
// 	newNode := &Node{Val: val}
// 	current := dl.head
// 	for current != nil {
// 		if current.Val == after {
// 			newNode.Next = current.Next
// 			newNode.Prev = current
// 			if current == dl.tail {
// 				dl.tail = newNode // Update tail pointer if 'after' is the tail node
// 			} else {
// 				current.Next.Prev = newNode
// 			}
// 			current.Next = newNode
// 			return
// 		}
// 		current = current.Next
// 	}
// }

// func (dl *DoublyLinkedList) BeforeValue(before, val int) {
// 	newNode := &Node{Val: val}
// 	current := dl.head
// 	for current != nil {
// 		if current.Val == before {
//             newNode.Next = current
//             newNode.Prev = current.Prev
//             if current == dl.head {
//                 dl.head = newNode // Update head pointer if 'before' is the head node
//             } else {
//                 current.Prev.Next = newNode
//             }
//             current.Prev = newNode
//             return
//         }
// 		current = current.Next
// 	}
// }

// func (dl *DoublyLinkedList) PrintForward() {
// 	current := dl.head
// 	fmt.Println("Orginal Order:")
// 	for current != nil {
// 		fmt.Print(current.Val, " -> ")
// 		current = current.Next
// 	}
// 	fmt.Print("nil")
// 	fmt.Println()
// }

// func (dl *DoublyLinkedList) PrintReverse() {
// 	current := dl.tail
// 	fmt.Println("Reverse Order:")

// 	for current != nil {
// 		fmt.Print(current.Val, " <- ")
// 		current = current.Prev
// 	}
// 	fmt.Print("nil")
// 	fmt.Println()
// }

// func main() {
// 	dl := &DoublyLinkedList{}
// 	nums := []int{2, 1, 3, 5, 4}

// 	for _, n := range nums {
// 		dl.Append(n)
// 	}

// 	dl.Delete(5)
// 	dl.Prepend(9)
// 	dl.AfterValue(3, 45)
// 	dl.BeforeValue(3,8)
// 	dl.PrintForward()
// 	dl.PrintReverse()
// }
