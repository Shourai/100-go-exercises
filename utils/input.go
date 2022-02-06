package utils

import (
	"bufio"
	"os"
	"strings"
)

func GetUserInput() string {
	scanner := bufio.NewReader(os.Stdin)
	input, _ := scanner.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	return input
}
