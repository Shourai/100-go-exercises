package utils

import (
	"bufio"
	"os"
	"strings"
)

func GetSingleLineInput() string {
	scanner := bufio.NewReader(os.Stdin)
	input, _ := scanner.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	return input
}

func GetMultiLineInput() []string {
	var inputList []string
	scanner := bufio.NewReader(os.Stdin)
	for {
		input, _ := scanner.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		if input == "" {
			break
		}
		inputList = append(inputList, input)
	}
	return inputList
}
