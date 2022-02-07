package main

import (
	"100-golang-exercises/utils"
	"fmt"
	"math"
	"strconv"
)

func main() {
	input := utils.GetSingleLineInput()
	number, _ := strconv.ParseFloat(input, 64)
	squared(number)
}

func squared(number float64) {
	fmt.Println(math.Pow(number, 2))
}
