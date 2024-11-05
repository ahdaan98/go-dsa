package main

import "fmt"

type stack struct {
	arr []int
}

func (s *stack) push(val int) {
	s.arr = append(s.arr, val)
}

func (s *stack) pop() {
	if len(s.arr) == 0 {
		return // stack is empty, nothing to pop
	}
	last := len(s.arr)-1
	s.arr = s.arr[:last]
}

func (s *stack) peek() (int,bool) {
	if len(s.arr) == 0{
		return 0,false
	}

	return s.arr[len(s.arr)-1],true
}

func main() {
	s := stack{}

	s.push(2)
	s.push(4)
	s.push(1)
	s.push(3)
	s.push(5)

	s.pop()

	fmt.Println(s.arr)

	if top, ok := s.peek(); ok{
		fmt.Println("peek:",top)
	} else {
		fmt.Println("stack is empty")
	}
}