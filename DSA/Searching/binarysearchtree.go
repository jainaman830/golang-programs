package main

import "fmt"

type Node struct {
	key   int
	right *Node
	left  *Node
}

var TotalIterations int

func (n *Node) Insert(key int) {
	if key > n.key {
		if n.right == nil {
			n.right = &Node{key: key}
		} else {
			n.right.Insert(key)
		}
	} else if key < n.key {
		if n.left == nil {
			n.left = &Node{key: key}
		} else {
			n.left.Insert(key)
		}
	}
}
func (n *Node) Search(key int) bool {
	TotalIterations++
	if n == nil {
		return false
	} else if n.key == key {
		return true
	}
	if key < n.key {
		return n.left.Search(key)
	} else if key > n.key {
		return n.right.Search(key)
	}
	return true
}
func main() {
	node := &Node{key: 100}
	node.Insert(110)
	fmt.Println(node.Search(1000))
	fmt.Println(TotalIterations)
}
