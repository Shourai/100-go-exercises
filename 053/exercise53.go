package main

import "fmt"

func main() {
	rectangle := shape{width: 2, height: 10}
	fmt.Println(rectangle.area())
	fmt.Println(rectangle.perimiter())
}

type shape struct {
	width, height int
}

func (rect shape) area() int {
	return rect.height * rect.width
}

func (rect shape) perimiter() int {
	return 2*rect.height + 2*rect.width
}
