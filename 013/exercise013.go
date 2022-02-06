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
	runeSlice := []rune(input)

	calculateNumbersAndDigits(runeSlice)
}

func calculateNumbersAndDigits(input []rune) {
	letters := 0
	digits := 0

	for _, i := range input {
		if (i >= 'a' && i <= 'z') || (i >= 'A' && i <= 'Z') {
			letters++
		} else if i >= '0' && i <= '9' {
			digits++
		}
	}

	fmt.Printf("LETTERS %d DIGITS %d", letters, digits)
}
