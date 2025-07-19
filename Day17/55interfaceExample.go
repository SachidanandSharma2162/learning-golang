package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.radius
}

type Rectangle struct {
	Height, Width float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}
func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}

type Triangle struct {
	Height, Width float32
}

func (t Triangle) Area() float32 {
	return 0.5 * t.Height * t.Width
}
func (t Triangle) Perimeter() float32 {
	return 3 * t.Width // assuming an equileteral triangle
}

func printShapeDetails(s Shape) {
	fmt.Println("Area", s.Area())
	fmt.Println("Perimeter", s.Perimeter())
}
func main() {
	c := Circle{radius: 4}
	r := Rectangle{Height: 5, Width: 10}
	t := Triangle{Height: 3, Width: 4}
	printShapeDetails(c)
	printShapeDetails(r)
	printShapeDetails(t)
}
