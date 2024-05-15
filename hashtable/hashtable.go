package main

import "fmt"

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	Next *bucketNode
}

func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

func (b *bucket) insert(key string) {
	if !b.search(key){
	newNode := &bucketNode{key: key}
	newNode.Next = b.head
	b.head = newNode
	}
}

func (b *bucket) search(key string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.Next
	}
	return false
}

func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.Next
		return
	}

	previousNode := b.head
	for previousNode.Next != nil {
		if previousNode.Next.key == k {
			previousNode.Next = previousNode.Next.Next
		}
		previousNode = previousNode.Next
	}
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum+=int(v)
	}
	return sum % ArraySize
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func (h *HashTable) Print() string {
	var result string
	for i, b := range h.array {
		result += fmt.Sprintf("Bucket %d: ", i)
		currentNode := b.head
		for currentNode != nil {
			result += fmt.Sprintf("%s -> ", currentNode.key)
			currentNode = currentNode.Next
		}
		result += "nil\n"
	}
	return result
}

func main() {
	h := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"RANDY",
	}

	for _,v := range list {
		h.Insert(v)
	}

	h.Delete("RANDY")
	fmt.Println(h.Print())
}