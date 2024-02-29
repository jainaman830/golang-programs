package main

import "fmt"

const ArraySize = 7

type Hash struct {
	bucket [ArraySize]*bucket
}

type bucket struct {
	head *node
}

type node struct {
	key  string
	next *node
}

func InitHash() *Hash {
	hash := &Hash{}
	for i := 0; i < ArraySize; i++ {
		hash.bucket[i] = &bucket{}
	}
	return hash
}
func (h *Hash) Insert(key string) {
	index := getHash(key)
	h.bucket[index].Insert(key)
}
func (h *Hash) Search(key string) bool {
	index := getHash(key)
	return h.bucket[index].Search(key)
}
func (h *Hash) Delete(key string) {
	index := getHash(key)
	h.bucket[index].Delete(key)
}
func (b *bucket) Insert(key string) {
	if !b.Search(key) {
		newNode := &node{key: key}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("Key already exists")
	}
}
func (b *bucket) Search(key string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}
func (b *bucket) Delete(key string) {
	if b.head == nil {
		return
	}
	if b.head.key == key {
		b.head = b.head.next
		return
	}
	currentNode := b.head
	for currentNode.next == nil {
		if currentNode.key == key {
			//delete
			currentNode.next = currentNode.next.next
		}
		currentNode = currentNode.next
	}
}
func getHash(key string) int {
	sum := 0
	for _, char := range key {
		sum += int(char)
	}
	return sum % ArraySize
}
func main() {
	hash := InitHash()
	hash.Insert("aman")
	hash.Insert("jain")

	fmt.Println(hash.Search("aman"))
	fmt.Println(hash.Search("aman2"))
	hash.Delete("aman")
	fmt.Println(hash.Search("aman"))
}
