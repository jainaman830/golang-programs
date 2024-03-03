package main

import "fmt"

func Add(a int) {
	a = a + 1
	fmt.Println("Inside function : ", a)
}
func AddWithPointer(a *int) {
	*a = *a + 1
	fmt.Println("Inside pointer function : ", *a)
}
func main() {
	a := 1
	Add(a)
	fmt.Println("After add : ", a)
	AddWithPointer(&a)
	fmt.Println("After add with pointer : ", a)
}
