package main

import "fmt"

type maxHeap struct {
	array []int
}

func (h *maxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *maxHeap) Extract() int {
	if len(h.array) == 0 {
		fmt.Println("array length is zero")
		return -1
	}
	last := len(h.array) - 1
	extracted := h.array[0]

	h.array[0] = h.array[last]
	h.array = h.array[:last]
	h.maxHeapifyDown(0)
	return extracted
}

func (h *maxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array)-1
	l,r := left(index),right(index)

	c2x := 0

	for l<=lastIndex{
		if l == lastIndex {
			c2x = l
		} else if h.array[l] > h.array[r] {
			c2x = l
		} else {
			c2x = r
		}

		if h.array[index] < h.array[c2x] {
			h.swap(index,c2x)
			index = c2x
			l,r = left(index),right(index)
		} else {
			return
		}
	}
}

func (h *maxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func main() {
	m := &maxHeap{}

	buildHeap := []int{10, 20, 30,5,7,9,1}
	for _, v := range buildHeap {
		m.Insert(v)
	}

	for i:=0;i<2;i++{
		m.Extract()
	}

	fmt.Println(m)
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *maxHeap) swap(a1, a2 int) {
	h.array[a1], h.array[a2] = h.array[a2], h.array[a1]
}
