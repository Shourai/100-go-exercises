package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetSingleLineInput() string {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Enter input: ")
	input, _ := scanner.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	return input
}

func GetMultiLineInput() []string {
	var inputList []string
	fmt.Print("Enter input: ")
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
