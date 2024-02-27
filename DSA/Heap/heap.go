package main

import "fmt"

type Maxheap struct {
	array []int
}

func (h *Maxheap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}
func (h *Maxheap) maxHeapifyUp(i int) {
	for h.array[parent(i)] < h.array[i] {
		h.swap(parent(i), i)
		i = parent(i)
	}
}
func (h *Maxheap) Extract() int {
	if len(h.array) == 0 {
		return -1
	}
	l := len(h.array) - 1
	extracted := h.array[l]
	h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.maxHeapifyDown(0)
	return extracted
}
func (h *Maxheap) maxHeapifyDown(i int) {
	last := len(h.array) - 1
	l, r := left(i), right(i)
	childToCompare := 0
	for l <= last {
		if l == last {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[i] < h.array[childToCompare] {
			h.swap(i, childToCompare)
			i = childToCompare
			l, r = left(i), right(i)
		} else {
			return
		}
	}
}
func parent(i int) int {
	return (i - 1) / 2
}
func right(i int) int {
	return 2*i + 2
}
func left(i int) int {
	return 2*i + 1
}
func (h *Maxheap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}
func main() {
	h := &Maxheap{}
	arr := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, val := range arr {
		h.Insert(val)
		fmt.Println(h)
	}
	for i := 0; i < 5; i++ {
		h.Extract()
		fmt.Println(h)
	}
}
