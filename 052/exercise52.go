package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func main() {
	c := circle{radius: 2}
	fmt.Println(c.area())
	fmt.Println(c.circumference())
}

func (c circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func (c circle) circumference() float64 {
	return 2 * math.Pi * c.radius
}
