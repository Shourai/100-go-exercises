package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Enter comma separated numbers: ")
	var input string
	fmt.Scanln(&input)

	output := strings.Split(input, ",")
	fmt.Println(output)
}
