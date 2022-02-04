package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func formula(input []string) string {
	const C float64 = 50
	const H float64 = 30
	var Q []string

	for _, x := range input {
		D, _ := strconv.ParseFloat(x, 64)
		ans := math.RoundToEven(math.Sqrt((2 * C * D) / H))
		ansString := strconv.FormatFloat(ans, 'f', 0, 64)
		Q = append(Q, ansString)
	}
	return strings.Join(Q, ",")
}

func main() {
	fmt.Print("Enter sequence: ")
	var input string
	fmt.Scanln(&input)

	splitString := strings.Split(input, ",")

	fmt.Println(formula(splitString))
}
