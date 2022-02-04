package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readString() string {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Enter string: ")
	input, _ := scanner.ReadString('\n')
	return input
}

func printStringCap(input string) {
	fmt.Print(strings.ToUpper(input))
}

func main() {

	text := readString()

	printStringCap(text)

}
