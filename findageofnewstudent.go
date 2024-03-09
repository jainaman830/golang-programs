/*
there are n students and if the student who is aged k years left the band and a new student joined the musical class then the average age of  entire class is lowered by m months. your task is to find the age of the new student
*/

package main

import "fmt"

func GetAge(n, k, m int) int {
	return k - int(float64(n)*float64(m)/12)
}
func main() {
	var n, k, m int
	fmt.Scanln(&n, &k)
	fmt.Scanln(&m)
	fmt.Println(GetAge(n, k, m))
}
