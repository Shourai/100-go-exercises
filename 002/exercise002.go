package main

import "fmt"

func main() {
	fmt.Print("Enter number: ")
	var number int

	fmt.Scanln(&number)
	fmt.Println(factorial(number))
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}
