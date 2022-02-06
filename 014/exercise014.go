package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	input, _ := scanner.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	calculateUpperAndLowercase(input)
}

func calculateUpperAndLowercase(input string) {
	uppercase := 0
	lowercase := 0

	for _, ch := range input {
		if ch > 'a' && ch < 'z' {
			lowercase++
		} else if ch > 'A' && ch < 'Z' {
			uppercase++
		}
	}

	fmt.Printf("UPPERCASE %d LOWERCASE %d", uppercase, lowercase)
}
