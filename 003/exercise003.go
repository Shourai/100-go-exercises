package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter number: ")
	var number int
	fmt.Scanln(&number)

	fmt.Println(squares(number))
}

func squares(number int) map[int]int {
	myMap := make(map[int]int)

	for i := 1; i <= number; i++ {
		myMap[i] = i * i
	}
	return myMap
}
