package main

import "fmt"

type Shape interface {
	Area() float64
	PrintName()
}
type Circle struct {
	Radius float64
	Name   string
}
type Rectangle struct {
	Width, Height float64
	Name          string
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Name string
type PrintText interface {
	PrintName()
}

func (c Circle) PrintName() {
	fmt.Println("Name : ", c.Name)
}
func (r Rectangle) PrintName() {
	fmt.Println("Name : ", r.Name)
}
func (s Name) PrintName() {
	fmt.Println("Name : ", s)
}
func main() {
	circle := Circle{
		Radius: 4,
		Name:   "circle",
	}
	Rectangle := Rectangle{
		Width:  10,
		Height: 20,
		Name:   "Rectangle",
	}
	Shapes := []Shape{circle, Rectangle}
	for _, shape := range Shapes {
		shape.PrintName()
		fmt.Println(shape.Area())
	}
	name := Name("Aman Jain")
	Print := PrintText(name)
	Print.PrintName()
}
