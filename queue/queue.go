package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	arr []int
}

func (q *Queue) enqueue(val int) {
	q.arr = append(q.arr, val)
}

func (q *Queue) dequeue() {
	if len(q.arr) == 0{
		fmt.Println("Queue is empty")
	}
	q.arr = q.arr[1:]
}

func (q *Queue) peek() (int, error) {
    if len(q.arr) == 0 {
        return 0, errors.New("queue is empty")
    }
    return q.arr[0], nil
}

func (q *Queue) isEmpty() bool {
    return len(q.arr) == 0
}

func (q *Queue) size() int {
    return len(q.arr)
}

func main() {
	q := Queue{}

	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	q.enqueue(5)

	q.dequeue()

	fmt.Println(q.arr)

	if val, err := q.peek(); err == nil {
        fmt.Println("Peeked:", val)
    } else {
        fmt.Println("Error peeking:", err)
    }

	fmt.Println("Is queue empty?", q.isEmpty())
    fmt.Println("Queue size:", q.size())	
}