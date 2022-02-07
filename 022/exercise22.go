package main

import (
	"100-golang-exercises/utils"
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := utils.GetSingleLineInput()
	slice := strings.Split(input, " ")
	countFrequency(slice)
}

func countFrequency(slice []string) {
	count := make(map[string]int)
	for _, w := range slice {

		count[w]++
	}

	var sliceOfKeys []string
	for k := range count {
		sliceOfKeys = append(sliceOfKeys, k)
	}

	sortKeys(&sliceOfKeys)

	for _, v := range sliceOfKeys {
		fmt.Printf("%s: %d\n", v, count[v])
	}

}

func sortKeys(slice *[]string) {
	sort.Strings(*slice)
}
