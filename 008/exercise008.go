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

	arr := strings.Split(input, ",")
	sort.Strings(arr)
	fmt.Print(arr)

}
