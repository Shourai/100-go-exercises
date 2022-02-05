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

	array := strings.Split(input, ",")
	sort.Strings(array)
	fmt.Println(strings.Join(array, ","))
}
