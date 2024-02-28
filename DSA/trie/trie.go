package main

import "fmt"

const AlphabetSize = 26

type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	t := &Trie{&Node{}}
	return t
}
func (t *Trie) Insert(w string) {
	length := len(w)
	currentNode := t.root
	for i := 0; i < length; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}
func (t *Trie) Search(w string) bool {
	length := len(w)
	currentNode := t.root
	for i := 0; i < length; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}

func main() {
	trie := InitTrie()
	words := []string{"foo", "aman", "jain", "test", "testword"}
	for _, val := range words {
		trie.Insert(val)
	}
	fmt.Println(trie.Search("foo"))
	fmt.Println(trie.Search("aman"))
	fmt.Println(trie.Search("ama"))
	fmt.Println(trie.Search("tes"))
}
