package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)

	input, _ := scanner.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	array := strings.Split(input, " ")
	array = unique(array)
	sort.Strings(array)
	fmt.Println(strings.Join(array, " "))

}

func unique(array []string) []string {
	list := []string{}

	keys := make(map[string]bool)

	for _, word := range array {
		if !keys[word] {
			keys[word] = true
			list = append(list, word)
		}
	}
	return list
}
