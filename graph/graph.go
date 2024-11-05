package main

import (
	"fmt"
)

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("vertext %v already exist", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v-->%v)",from,to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent,to){
		err := fmt.Errorf("existing edge (%v-->%v)",from,to)
		fmt.Println(err.Error())
	} else {
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

func (g *Graph) getVertex(k int) *Vertex {
	for i,v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}

	return false
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("vertex %v :", v.key)
		for _, adj := range v.adjacent {
			fmt.Printf(" %d  ", adj.key)
		}
		fmt.Println()
	}
}

func main() {
	t := &Graph{}

	for i := 0; i < 5; i++ {
		t.AddVertex(i)
	}

	t.AddEdge(1,4)
	t.AddEdge(5,2)
	t.AddEdge(1,4)
	t.AddEdge(2,3)
	t.AddEdge(2,4)
	t.AddEdge(2,1)
	t.Print()
}
