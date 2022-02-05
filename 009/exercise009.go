package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Enter input: ")
	lines := make([]string, 0)
	for {
		input, _ := scanner.ReadString('\n')
		if input == "\n" {
			break
		} else {
			input = strings.TrimSuffix(input, "\n")
			lines = append(lines, strings.ToUpper(input))
		}
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
