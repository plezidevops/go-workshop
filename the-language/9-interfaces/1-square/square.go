package main

import (
	"fmt"
	"math"
)

type Square struct {
	Length float64
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// sumArea return the sum of all areas in the slice
func sumArea(shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}

type Shape interface {
	Area() float64
}

func main() {
	s := Square{20}
	fmt.Println(s.Area())

	c := Circle{10}
	fmt.Println(c.Area())

	shapes := []Shape{s, c}
	total := sumArea(shapes)
	fmt.Println(total)

}
