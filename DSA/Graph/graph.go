package main

import "fmt"

type Graph struct {
	vertices []*vertex
}
type vertex struct {
	key      int
	adjecent []*vertex
}

func (g *Graph) AddVertex(key int) {
	if contains(g.vertices, key) {
		fmt.Println("Key already exists")
	} else {
		g.vertices = append(g.vertices, &vertex{key: key})
	}
}
func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		fmt.Printf("Invalid edge %v --> %v\n", from, to)
	} else if contains(fromVertex.adjecent, to) {
		fmt.Printf("Existing edge %v --> %v\n", from, to)
	} else {
		fromVertex.adjecent = append(fromVertex.adjecent, toVertex)
	}
}
func (g *Graph) Print() {
	for _, vertex := range g.vertices {
		fmt.Printf("Vertex %v", vertex.key)
		for _, ad := range vertex.adjecent {
			fmt.Printf(" : %v ", ad.key)
		}
		fmt.Println()
	}
}
func (g *Graph) getVertex(key int) *vertex {
	for _, vertex := range g.vertices {
		if vertex.key == key {
			return vertex
		}
	}
	return nil
}
func contains(list []*vertex, key int) bool {
	for _, v := range list {
		if v.key == key {
			return true
		}
	}
	return false
}
func main() {
	g := &Graph{}
	for i := 0; i < 5; i++ {
		g.AddVertex(i)
	}
	g.AddEdge(1, 2)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(6, 2)
	g.AddEdge(5, 2)
	g.Print()
}
